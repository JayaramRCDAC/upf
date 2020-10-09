// Copyright (c) 2016-2017, Nefeli Networks, Inc.
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// * Redistributions of source code must retain the above copyright notice, this
// list of conditions and the following disclaimer.
//
// * Redistributions in binary form must reproduce the above copyright notice,
// this list of conditions and the following disclaimer in the documentation
// and/or other materials provided with the distribution.
//
// * Neither the names of the copyright holders nor the names of their
// contributors may be used to endorse or promote products derived from this
// software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
// ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE
// LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
// CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
// SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
// INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
// CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
// ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
// POSSIBILITY OF SUCH DAMAGE.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        (unknown)
// source: util_msg.proto

package bess_pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

/// The Field message represents one field in a packet -- either stored in metadata or in the packet body.
type Field struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Position:
	//	*Field_AttrName
	//	*Field_Offset
	Position isField_Position `protobuf_oneof:"position"`
	NumBytes uint32           `protobuf:"varint,3,opt,name=num_bytes,json=numBytes,proto3" json:"num_bytes,omitempty"` /// The size of the data in bytes
}

func (x *Field) Reset() {
	*x = Field{}
	if protoimpl.UnsafeEnabled {
		mi := &file_util_msg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Field) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Field) ProtoMessage() {}

func (x *Field) ProtoReflect() protoreflect.Message {
	mi := &file_util_msg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Field.ProtoReflect.Descriptor instead.
func (*Field) Descriptor() ([]byte, []int) {
	return file_util_msg_proto_rawDescGZIP(), []int{0}
}

func (m *Field) GetPosition() isField_Position {
	if m != nil {
		return m.Position
	}
	return nil
}

func (x *Field) GetAttrName() string {
	if x, ok := x.GetPosition().(*Field_AttrName); ok {
		return x.AttrName
	}
	return ""
}

func (x *Field) GetOffset() uint32 {
	if x, ok := x.GetPosition().(*Field_Offset); ok {
		return x.Offset
	}
	return 0
}

func (x *Field) GetNumBytes() uint32 {
	if x != nil {
		return x.NumBytes
	}
	return 0
}

type isField_Position interface {
	isField_Position()
}

type Field_AttrName struct {
	AttrName string `protobuf:"bytes,1,opt,name=attr_name,json=attrName,proto3,oneof"` /// The metadata attribute assigned to store the data
}

type Field_Offset struct {
	Offset uint32 `protobuf:"varint,2,opt,name=offset,proto3,oneof"` /// The offset in bytes to store the data into
}

func (*Field_AttrName) isField_Position() {}

func (*Field_Offset) isField_Position() {}

/// The FieldData message encodes a value to insert into a packet; the value can be supplied as either an int or a bytestring.
type FieldData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Encoding:
	//	*FieldData_ValueBin
	//	*FieldData_ValueInt
	Encoding isFieldData_Encoding `protobuf_oneof:"encoding"`
}

func (x *FieldData) Reset() {
	*x = FieldData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_util_msg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldData) ProtoMessage() {}

func (x *FieldData) ProtoReflect() protoreflect.Message {
	mi := &file_util_msg_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldData.ProtoReflect.Descriptor instead.
func (*FieldData) Descriptor() ([]byte, []int) {
	return file_util_msg_proto_rawDescGZIP(), []int{1}
}

func (m *FieldData) GetEncoding() isFieldData_Encoding {
	if m != nil {
		return m.Encoding
	}
	return nil
}

func (x *FieldData) GetValueBin() []byte {
	if x, ok := x.GetEncoding().(*FieldData_ValueBin); ok {
		return x.ValueBin
	}
	return nil
}

func (x *FieldData) GetValueInt() uint64 {
	if x, ok := x.GetEncoding().(*FieldData_ValueInt); ok {
		return x.ValueInt
	}
	return 0
}

type isFieldData_Encoding interface {
	isFieldData_Encoding()
}

type FieldData_ValueBin struct {
	ValueBin []byte `protobuf:"bytes,1,opt,name=value_bin,json=valueBin,proto3,oneof"` /// The value as a bytestring
}

type FieldData_ValueInt struct {
	ValueInt uint64 `protobuf:"varint,2,opt,name=value_int,json=valueInt,proto3,oneof"` /// The value in integer format
}

func (*FieldData_ValueBin) isFieldData_Encoding() {}

func (*FieldData_ValueInt) isFieldData_Encoding() {}

var File_util_msg_proto protoreflect.FileDescriptor

var file_util_msg_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x75, 0x74, 0x69, 0x6c, 0x5f, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x62, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x62, 0x22, 0x69, 0x0a, 0x05, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x12, 0x1d, 0x0a, 0x09, 0x61, 0x74, 0x74, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x61, 0x74, 0x74, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x48, 0x00, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6e,
	0x75, 0x6d, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x6e, 0x75, 0x6d, 0x42, 0x79, 0x74, 0x65, 0x73, 0x42, 0x0a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x55, 0x0a, 0x09, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x1d, 0x0a, 0x09, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x62, 0x69, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x08, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x69, 0x6e,
	0x12, 0x1d, 0x0a, 0x09, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x48, 0x00, 0x52, 0x08, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x49, 0x6e, 0x74, 0x42,
	0x0a, 0x0a, 0x08, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_util_msg_proto_rawDescOnce sync.Once
	file_util_msg_proto_rawDescData = file_util_msg_proto_rawDesc
)

func file_util_msg_proto_rawDescGZIP() []byte {
	file_util_msg_proto_rawDescOnce.Do(func() {
		file_util_msg_proto_rawDescData = protoimpl.X.CompressGZIP(file_util_msg_proto_rawDescData)
	})
	return file_util_msg_proto_rawDescData
}

var file_util_msg_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_util_msg_proto_goTypes = []interface{}{
	(*Field)(nil),     // 0: bess.pb.Field
	(*FieldData)(nil), // 1: bess.pb.FieldData
}
var file_util_msg_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_util_msg_proto_init() }
func file_util_msg_proto_init() {
	if File_util_msg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_util_msg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Field); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_util_msg_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FieldData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_util_msg_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Field_AttrName)(nil),
		(*Field_Offset)(nil),
	}
	file_util_msg_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*FieldData_ValueBin)(nil),
		(*FieldData_ValueInt)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_util_msg_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_util_msg_proto_goTypes,
		DependencyIndexes: file_util_msg_proto_depIdxs,
		MessageInfos:      file_util_msg_proto_msgTypes,
	}.Build()
	File_util_msg_proto = out.File
	file_util_msg_proto_rawDesc = nil
	file_util_msg_proto_goTypes = nil
	file_util_msg_proto_depIdxs = nil
}
