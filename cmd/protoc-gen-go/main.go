// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The protoc-gen-go binary is a protoc plugin to generate a Go protocol
// buffer package.
package main

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/golang/protobuf/proto"
	descpb "github.com/golang/protobuf/protoc-gen-go/descriptor"
	"google.golang.org/proto/protogen"
	"google.golang.org/proto/reflect/protoreflect"
)

// generatedCodeVersion indicates a version of the generated code.
// It is incremented whenever an incompatibility between the generated code and
// proto package is introduced; the generated code references
// a constant, proto.ProtoPackageIsVersionN (where N is generatedCodeVersion).
const generatedCodeVersion = 2

const protoPackage = "github.com/golang/protobuf/proto"

func main() {
	var flags flag.FlagSet
	// TODO: Decide what to do for backwards compatibility with plugins=grpc.
	flags.String("plugins", "", "")
	opts := &protogen.Options{
		ParamFunc: flags.Set,
	}
	protogen.Run(opts, func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			genFile(gen, f)
		}
		return nil
	})
}

type File struct {
	*protogen.File
	locationMap   map[string][]*descpb.SourceCodeInfo_Location
	descriptorVar string // var containing the gzipped FileDescriptorProto
	allEnums      []*protogen.Enum
	allMessages   []*protogen.Message
}

func genFile(gen *protogen.Plugin, file *protogen.File) {
	f := &File{
		File:        file,
		locationMap: make(map[string][]*descpb.SourceCodeInfo_Location),
	}
	for _, loc := range file.Proto.GetSourceCodeInfo().GetLocation() {
		key := pathKey(loc.Path)
		f.locationMap[key] = append(f.locationMap[key], loc)
	}

	f.allEnums = append(f.allEnums, f.File.Enums...)
	f.allMessages = append(f.allMessages, f.File.Messages...)
	for _, message := range f.Messages {
		f.initMessage(message)
	}

	// Determine the name of the var holding the file descriptor:
	//
	//     fileDescriptor_<hash of filename>
	filenameHash := sha256.Sum256([]byte(f.Desc.Path()))
	f.descriptorVar = fmt.Sprintf("fileDescriptor_%s", hex.EncodeToString(filenameHash[:8]))

	g := gen.NewGeneratedFile(f.GeneratedFilenamePrefix+".pb.go", f.GoImportPath)
	g.P("// Code generated by protoc-gen-go. DO NOT EDIT.")
	g.P("// source: ", f.Desc.Path())
	g.P()
	const filePackageField = 2 // FileDescriptorProto.package
	genComment(g, f, []int32{filePackageField})
	g.P()
	g.P("package ", f.GoPackageName)
	g.P()

	// These references are not necessary, since we automatically add
	// all necessary imports before formatting the generated file.
	//
	// This section exists to generate output more consistent with
	// the previous version of protoc-gen-go, to make it easier to
	// detect unintended variations.
	//
	// TODO: Eventually remove this.
	g.P("// Reference imports to suppress errors if they are not otherwise used.")
	g.P("var _ = ", protogen.GoIdent{GoImportPath: protoPackage, GoName: "Marshal"})
	g.P("var _ = ", protogen.GoIdent{GoImportPath: "fmt", GoName: "Errorf"})
	g.P("var _ = ", protogen.GoIdent{GoImportPath: "math", GoName: "Inf"})
	g.P()

	g.P("// This is a compile-time assertion to ensure that this generated file")
	g.P("// is compatible with the proto package it is being compiled against.")
	g.P("// A compilation error at this line likely means your copy of the")
	g.P("// proto package needs to be updated.")
	g.P("const _ = ", protogen.GoIdent{
		GoImportPath: protoPackage,
		GoName:       fmt.Sprintf("ProtoPackageIsVersion%d", generatedCodeVersion),
	}, "// please upgrade the proto package")
	g.P()

	for _, enum := range f.allEnums {
		genEnum(gen, g, f, enum)
	}
	for _, message := range f.allMessages {
		genMessage(gen, g, f, message)
	}

	genInitFunction(gen, g, f)

	genFileDescriptor(gen, g, f)
}

func (f *File) initMessage(message *protogen.Message) {
	f.allEnums = append(f.allEnums, message.Enums...)
	f.allMessages = append(f.allMessages, message.Messages...)
	for _, m := range message.Messages {
		f.initMessage(m)
	}
}

func genFileDescriptor(gen *protogen.Plugin, g *protogen.GeneratedFile, f *File) {
	// Trim the source_code_info from the descriptor.
	// Marshal and gzip it.
	descProto := proto.Clone(f.Proto).(*descpb.FileDescriptorProto)
	descProto.SourceCodeInfo = nil
	b, err := proto.Marshal(descProto)
	if err != nil {
		gen.Error(err)
		return
	}
	var buf bytes.Buffer
	w, _ := gzip.NewWriterLevel(&buf, gzip.BestCompression)
	w.Write(b)
	w.Close()
	b = buf.Bytes()

	g.P("func init() { proto.RegisterFile(", strconv.Quote(f.Desc.Path()), ", ", f.descriptorVar, ") }")
	g.P()
	g.P("var ", f.descriptorVar, " = []byte{")
	g.P("// ", len(b), " bytes of a gzipped FileDescriptorProto")
	for len(b) > 0 {
		n := 16
		if n > len(b) {
			n = len(b)
		}

		s := ""
		for _, c := range b[:n] {
			s += fmt.Sprintf("0x%02x,", c)
		}
		g.P(s)

		b = b[n:]
	}
	g.P("}")
	g.P()
}

func genEnum(gen *protogen.Plugin, g *protogen.GeneratedFile, f *File, enum *protogen.Enum) {
	genComment(g, f, enum.Path)
	// TODO: deprecation
	g.P("type ", enum.GoIdent, " int32")
	g.P("const (")
	for _, value := range enum.Values {
		genComment(g, f, value.Path)
		// TODO: deprecation
		g.P(value.GoIdent, " ", enum.GoIdent, " = ", value.Desc.Number())
	}
	g.P(")")
	g.P()
	nameMap := enum.GoIdent.GoName + "_name"
	g.P("var ", nameMap, " = map[int32]string{")
	generated := make(map[protoreflect.EnumNumber]bool)
	for _, value := range enum.Values {
		duplicate := ""
		if _, present := generated[value.Desc.Number()]; present {
			duplicate = "// Duplicate value: "
		}
		g.P(duplicate, value.Desc.Number(), ": ", strconv.Quote(string(value.Desc.Name())), ",")
		generated[value.Desc.Number()] = true
	}
	g.P("}")
	g.P()
	valueMap := enum.GoIdent.GoName + "_value"
	g.P("var ", valueMap, " = map[string]int32{")
	for _, value := range enum.Values {
		g.P(strconv.Quote(string(value.Desc.Name())), ": ", value.Desc.Number(), ",")
	}
	g.P("}")
	g.P()
	if enum.Desc.Syntax() != protoreflect.Proto3 {
		g.P("func (x ", enum.GoIdent, ") Enum() *", enum.GoIdent, " {")
		g.P("p := new(", enum.GoIdent, ")")
		g.P("*p = x")
		g.P("return p")
		g.P("}")
		g.P()
	}
	g.P("func (x ", enum.GoIdent, ") String() string {")
	g.P("return ", protogen.GoIdent{GoImportPath: protoPackage, GoName: "EnumName"}, "(", enum.GoIdent, "_name, int32(x))")
	g.P("}")
	g.P()

	if enum.Desc.Syntax() != protoreflect.Proto3 {
		g.P("func (x *", enum.GoIdent, ") UnmarshalJSON(data []byte) error {")
		g.P("value, err := ", protogen.GoIdent{GoImportPath: protoPackage, GoName: "UnmarshalJSONEnum"}, "(", enum.GoIdent, `_value, data, "`, enum.GoIdent, `")`)
		g.P("if err != nil {")
		g.P("return err")
		g.P("}")
		g.P("*x = ", enum.GoIdent, "(value)")
		g.P("return nil")
		g.P("}")
		g.P()
	}

	var indexes []string
	for i := 1; i < len(enum.Path); i += 2 {
		indexes = append(indexes, strconv.Itoa(int(enum.Path[i])))
	}
	g.P("func (", enum.GoIdent, ") EnumDescriptor() ([]byte, []int) {")
	g.P("return ", f.descriptorVar, ", []int{", strings.Join(indexes, ","), "}")
	g.P("}")
	g.P()

	genWellKnownType(g, enum.GoIdent, enum.Desc)
}

// enumRegistryName returns the name used to register an enum with the proto
// package registry.
//
// Confusingly, this is <proto_package>.<go_ident>. This probably should have
// been the full name of the proto enum type instead, but changing it at this
// point would require thought.
func enumRegistryName(enum *protogen.Enum) string {
	// Find the FileDescriptor for this enum.
	var desc protoreflect.Descriptor = enum.Desc
	for {
		p, ok := desc.Parent()
		if !ok {
			break
		}
		desc = p
	}
	fdesc := desc.(protoreflect.FileDescriptor)
	return string(fdesc.Package()) + "." + enum.GoIdent.GoName
}

func genMessage(gen *protogen.Plugin, g *protogen.GeneratedFile, f *File, message *protogen.Message) {
	if message.Desc.IsMapEntry() {
		return
	}

	genComment(g, f, message.Path)
	// TODO: deprecation
	g.P("type ", message.GoIdent, " struct {")
	for _, field := range message.Fields {
		if field.Desc.OneofType() != nil {
			// TODO oneofs
			continue
		}
		genComment(g, f, field.Path)
		goType, pointer := fieldGoType(g, field)
		if pointer {
			goType = "*" + goType
		}
		tags := []string{
			fmt.Sprintf("protobuf:%q", fieldProtobufTag(field)),
			fmt.Sprintf("json:%q", fieldJSONTag(field)),
		}
		if field.Desc.IsMap() {
			key := field.MessageType.Fields[0]
			val := field.MessageType.Fields[1]
			tags = append(tags,
				fmt.Sprintf("protobuf_key:%q", fieldProtobufTag(key)),
				fmt.Sprintf("protobuf_val:%q", fieldProtobufTag(val)),
			)
		}
		g.P(field.GoIdent, " ", goType, " `", strings.Join(tags, " "), "`")
	}
	g.P("XXX_NoUnkeyedLiteral struct{} `json:\"-\"`")
	// TODO XXX_InternalExtensions
	g.P("XXX_unrecognized []byte `json:\"-\"`")
	g.P("XXX_sizecache int32 `json:\"-\"`")
	g.P("}")
	g.P()

	// Reset
	g.P("func (m *", message.GoIdent, ") Reset() { *m = ", message.GoIdent, "{} }")
	// String
	g.P("func (m *", message.GoIdent, ") String() string { return ", protogen.GoIdent{
		GoImportPath: protoPackage,
		GoName:       "CompactTextString",
	}, "(m) }")
	// ProtoMessage
	g.P("func (*", message.GoIdent, ") ProtoMessage() {}")
	// Descriptor
	var indexes []string
	for i := 1; i < len(message.Path); i += 2 {
		indexes = append(indexes, strconv.Itoa(int(message.Path[i])))
	}
	g.P("func (*", message.GoIdent, ") Descriptor() ([]byte, []int) {")
	g.P("return ", f.descriptorVar, ", []int{", strings.Join(indexes, ","), "}")
	g.P("}")
	// TODO: extension support methods

	// Table-driven proto support.
	//
	// TODO: It does not scale to keep adding another method for every
	// operation on protos that we want to switch over to using the
	// table-driven approach. Instead, we should only add a single method
	// that allows getting access to the *InternalMessageInfo struct and then
	// calling Unmarshal, Marshal, Merge, Size, and Discard directly on that.
	messageInfoVar := "xxx_messageInfo_" + message.GoIdent.GoName
	// XXX_Unmarshal
	g.P("func (m *", message.GoIdent, ") XXX_Unmarshal(b []byte) error {")
	g.P("return ", messageInfoVar, ".Unmarshal(m, b)")
	g.P("}")
	// XXX_Marshal
	g.P("func (m *", message.GoIdent, ") XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {")
	g.P("return ", messageInfoVar, ".Marshal(b, m, deterministic)")
	g.P("}")
	// XXX_Merge
	g.P("func (m *", message.GoIdent, ") XXX_Merge(src proto.Message) {")
	g.P(messageInfoVar, ".Merge(m, src)")
	g.P("}")
	// XXX_Size
	g.P("func (m *", message.GoIdent, ") XXX_Size() int {")
	g.P("return ", messageInfoVar, ".Size(m)")
	g.P("}")
	// XXX_DiscardUnknown
	g.P("func (m *", message.GoIdent, ") XXX_DiscardUnknown() {")
	g.P(messageInfoVar, ".DiscardUnknown(m)")
	g.P("}")
	g.P()
	g.P("var ", messageInfoVar, " ", protogen.GoIdent{
		GoImportPath: protoPackage,
		GoName:       "InternalMessageInfo",
	})
	g.P()

	// Constants and vars holding the default values of fields.
	for _, field := range message.Fields {
		if !field.Desc.HasDefault() {
			continue
		}
		defVarName := "Default_" + message.GoIdent.GoName + "_" + field.GoIdent.GoName
		def := field.Desc.Default()
		switch field.Desc.Kind() {
		case protoreflect.StringKind:
			g.P("const ", defVarName, " string = ", strconv.Quote(def.String()))
		case protoreflect.BytesKind:
			g.P("var ", defVarName, " []byte = []byte(", strconv.Quote(string(def.Bytes())), ")")
		case protoreflect.EnumKind:
			enum := field.EnumType
			evalue := enum.Values[enum.Desc.Values().ByNumber(def.Enum()).Index()]
			g.P("const ", defVarName, " ", field.EnumType.GoIdent, " = ", evalue.GoIdent)
		case protoreflect.FloatKind, protoreflect.DoubleKind:
			// Floating point numbers need extra handling for -Inf/Inf/NaN.
			f := field.Desc.Default().Float()
			goType := "float64"
			if field.Desc.Kind() == protoreflect.FloatKind {
				goType = "float32"
			}
			// funcCall returns a call to a function in the math package,
			// possibly converting the result to float32.
			funcCall := func(fn, param string) string {
				s := g.QualifiedGoIdent(protogen.GoIdent{
					GoImportPath: "math",
					GoName:       fn,
				}) + param
				if goType != "float64" {
					s = goType + "(" + s + ")"
				}
				return s
			}
			switch {
			case math.IsInf(f, -1):
				g.P("var ", defVarName, " ", goType, " = ", funcCall("Inf", "(-1)"))
			case math.IsInf(f, 1):
				g.P("var ", defVarName, " ", goType, " = ", funcCall("Inf", "(1)"))
			case math.IsNaN(f):
				g.P("var ", defVarName, " ", goType, " = ", funcCall("NaN", "()"))
			default:
				g.P("const ", defVarName, " ", goType, " = ", f)
			}
		default:
			goType, _ := fieldGoType(g, field)
			g.P("const ", defVarName, " ", goType, " = ", def.Interface())
		}
	}
	g.P()

	// Getters.
	for _, field := range message.Fields {
		goType, pointer := fieldGoType(g, field)
		defaultValue := fieldDefaultValue(g, message, field)
		g.P("func (m *", message.GoIdent, ") Get", field.GoIdent, "() ", goType, " {")
		if field.Desc.Syntax() == protoreflect.Proto3 || defaultValue == "nil" {
			g.P("if m != nil {")
		} else {
			g.P("if m != nil && m.", field.GoIdent, " != nil {")
		}
		star := ""
		if pointer {
			star = "*"
		}
		g.P("return ", star, " m.", field.GoIdent)
		g.P("}")
		g.P("return ", defaultValue)
		g.P("}")
		g.P()
	}

	genWellKnownType(g, message.GoIdent, message.Desc)
}

// fieldGoType returns the Go type used for a field.
//
// If it returns pointer=true, the struct field is a pointer to the type.
func fieldGoType(g *protogen.GeneratedFile, field *protogen.Field) (goType string, pointer bool) {
	pointer = true
	switch field.Desc.Kind() {
	case protoreflect.BoolKind:
		goType = "bool"
	case protoreflect.EnumKind:
		goType = g.QualifiedGoIdent(field.EnumType.GoIdent)
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		goType = "int32"
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		goType = "uint32"
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		goType = "int64"
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		goType = "uint64"
	case protoreflect.FloatKind:
		goType = "float32"
	case protoreflect.DoubleKind:
		goType = "float64"
	case protoreflect.StringKind:
		goType = "string"
	case protoreflect.BytesKind:
		goType = "[]byte"
		pointer = false
	case protoreflect.MessageKind, protoreflect.GroupKind:
		if field.Desc.IsMap() {
			keyType, _ := fieldGoType(g, field.MessageType.Fields[0])
			valType, _ := fieldGoType(g, field.MessageType.Fields[1])
			return fmt.Sprintf("map[%v]%v", keyType, valType), false
		}
		goType = "*" + g.QualifiedGoIdent(field.MessageType.GoIdent)
		pointer = false
	}
	if field.Desc.Cardinality() == protoreflect.Repeated {
		goType = "[]" + goType
		pointer = false
	}
	if field.Desc.Syntax() == protoreflect.Proto3 {
		pointer = false
	}
	return goType, pointer
}

func fieldProtobufTag(field *protogen.Field) string {
	var tag []string
	// wire type
	tag = append(tag, wireTypes[field.Desc.Kind()])
	// field number
	tag = append(tag, strconv.Itoa(int(field.Desc.Number())))
	// cardinality
	switch field.Desc.Cardinality() {
	case protoreflect.Optional:
		tag = append(tag, "opt")
	case protoreflect.Required:
		tag = append(tag, "req")
	case protoreflect.Repeated:
		tag = append(tag, "rep")
	}
	// TODO: packed
	// name
	name := string(field.Desc.Name())
	if field.Desc.Kind() == protoreflect.GroupKind {
		// The name of the FieldDescriptor for a group field is
		// lowercased. To find the original capitalization, we
		// look in the field's MessageType.
		name = string(field.MessageType.Desc.Name())
	}
	tag = append(tag, "name="+name)
	// JSON name
	if jsonName := field.Desc.JSONName(); jsonName != "" && jsonName != name {
		tag = append(tag, "json="+jsonName)
	}
	// proto3
	if field.Desc.Syntax() == protoreflect.Proto3 {
		tag = append(tag, "proto3")
	}
	// enum
	if field.Desc.Kind() == protoreflect.EnumKind {
		tag = append(tag, "enum="+enumRegistryName(field.EnumType))
	}
	// oneof
	if field.Desc.OneofType() != nil {
		tag = append(tag, "oneof")
	}
	// default value
	// This must appear last in the tag, since commas in strings aren't escaped.
	if field.Desc.HasDefault() {
		var def string
		switch field.Desc.Kind() {
		case protoreflect.BoolKind:
			if field.Desc.Default().Bool() {
				def = "1"
			} else {
				def = "0"
			}
		case protoreflect.BytesKind:
			def = string(field.Desc.Default().Bytes())
		case protoreflect.FloatKind, protoreflect.DoubleKind:
			f := field.Desc.Default().Float()
			switch {
			case math.IsInf(f, -1):
				def = "-inf"
			case math.IsInf(f, 1):
				def = "inf"
			case math.IsNaN(f):
				def = "nan"
			default:
				def = fmt.Sprint(f)
			}
		default:
			def = fmt.Sprint(field.Desc.Default().Interface())
		}
		tag = append(tag, "def="+def)
	}
	return strings.Join(tag, ",")
}

func fieldDefaultValue(g *protogen.GeneratedFile, message *protogen.Message, field *protogen.Field) string {
	if field.Desc.Cardinality() == protoreflect.Repeated {
		return "nil"
	}
	if field.Desc.HasDefault() {
		defVarName := "Default_" + message.GoIdent.GoName + "_" + field.GoIdent.GoName
		if field.Desc.Kind() == protoreflect.BytesKind {
			return "append([]byte(nil), " + defVarName + "...)"
		}
		return defVarName
	}
	switch field.Desc.Kind() {
	case protoreflect.BoolKind:
		return "false"
	case protoreflect.StringKind:
		return `""`
	case protoreflect.MessageKind, protoreflect.GroupKind, protoreflect.BytesKind:
		return "nil"
	case protoreflect.EnumKind:
		return g.QualifiedGoIdent(field.EnumType.Values[0].GoIdent)
	default:
		return "0"
	}
}

var wireTypes = map[protoreflect.Kind]string{
	protoreflect.BoolKind:     "varint",
	protoreflect.EnumKind:     "varint",
	protoreflect.Int32Kind:    "varint",
	protoreflect.Sint32Kind:   "zigzag32",
	protoreflect.Uint32Kind:   "varint",
	protoreflect.Int64Kind:    "varint",
	protoreflect.Sint64Kind:   "zigzag64",
	protoreflect.Uint64Kind:   "varint",
	protoreflect.Sfixed32Kind: "fixed32",
	protoreflect.Fixed32Kind:  "fixed32",
	protoreflect.FloatKind:    "fixed32",
	protoreflect.Sfixed64Kind: "fixed64",
	protoreflect.Fixed64Kind:  "fixed64",
	protoreflect.DoubleKind:   "fixed64",
	protoreflect.StringKind:   "bytes",
	protoreflect.BytesKind:    "bytes",
	protoreflect.MessageKind:  "bytes",
	protoreflect.GroupKind:    "group",
}

func fieldJSONTag(field *protogen.Field) string {
	return string(field.Desc.Name()) + ",omitempty"
}

// genInitFunction generates an init function that registers the types in the
// generated file with the proto package.
func genInitFunction(gen *protogen.Plugin, g *protogen.GeneratedFile, f *File) {
	if len(f.allMessages) == 0 && len(f.allEnums) == 0 {
		return
	}

	g.P("func init() {")
	for _, message := range f.allMessages {
		if message.Desc.IsMapEntry() {
			continue
		}

		name := message.GoIdent.GoName
		g.P(protogen.GoIdent{
			GoImportPath: protoPackage,
			GoName:       "RegisterType",
		}, fmt.Sprintf("((*%s)(nil), %q)", name, message.Desc.FullName()))

		// Types of map fields, sorted by the name of the field message type.
		var mapFields []*protogen.Field
		for _, field := range message.Fields {
			if field.Desc.IsMap() {
				mapFields = append(mapFields, field)
			}
		}
		sort.Slice(mapFields, func(i, j int) bool {
			ni := mapFields[i].MessageType.Desc.FullName()
			nj := mapFields[j].MessageType.Desc.FullName()
			return ni < nj
		})
		for _, field := range mapFields {
			typeName := string(field.MessageType.Desc.FullName())
			goType, _ := fieldGoType(g, field)
			g.P(protogen.GoIdent{
				GoImportPath: protoPackage,
				GoName:       "RegisterMapType",
			}, fmt.Sprintf("((%v)(nil), %q)", goType, typeName))
		}
	}
	for _, enum := range f.allEnums {
		name := enum.GoIdent.GoName
		g.P(protogen.GoIdent{
			GoImportPath: protoPackage,
			GoName:       "RegisterEnum",
		}, fmt.Sprintf("(%q, %s_name, %s_value)", enumRegistryName(enum), name, name))
	}
	g.P("}")
	g.P()
}

func genComment(g *protogen.GeneratedFile, f *File, path []int32) {
	for _, loc := range f.locationMap[pathKey(path)] {
		if loc.LeadingComments == nil {
			continue
		}
		for _, line := range strings.Split(strings.TrimSuffix(loc.GetLeadingComments(), "\n"), "\n") {
			g.P("//", line)
		}
		return
	}
}

// pathKey converts a location path to a string suitable for use as a map key.
func pathKey(path []int32) string {
	var buf []byte
	for i, x := range path {
		if i != 0 {
			buf = append(buf, ',')
		}
		buf = strconv.AppendInt(buf, int64(x), 10)
	}
	return string(buf)
}

func genWellKnownType(g *protogen.GeneratedFile, ident protogen.GoIdent, desc protoreflect.Descriptor) {
	if wellKnownTypes[desc.FullName()] {
		g.P("func (", ident, `) XXX_WellKnownType() string { return "`, desc.Name(), `" }`)
		g.P()
	}
}

// Names of messages and enums for which we will generate XXX_WellKnownType methods.
var wellKnownTypes = map[protoreflect.FullName]bool{
	"google.protobuf.NullValue": true,
}
