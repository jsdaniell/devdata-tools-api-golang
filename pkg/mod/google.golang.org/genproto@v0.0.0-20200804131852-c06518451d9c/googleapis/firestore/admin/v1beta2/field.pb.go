// Copyright 2019 Google LLC.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.12.3
// source: google/firestore/admin/v1beta2/field.proto

package admin

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Represents a single field in the database.
//
// Fields are grouped by their "Collection Group", which represent all
// collections in the database with the same id.
type Field struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A field name of the form
	// `projects/{project_id}/databases/{database_id}/collectionGroups/{collection_id}/fields/{field_path}`
	//
	// A field path may be a simple field name, e.g. `address` or a path to fields
	// within map_value , e.g. `address.city`,
	// or a special field path. The only valid special field is `*`, which
	// represents any field.
	//
	// Field paths may be quoted using ` (backtick). The only character that needs
	// to be escaped within a quoted field path is the backtick character itself,
	// escaped using a backslash. Special characters in field paths that
	// must be quoted include: `*`, `.`,
	// ``` (backtick), `[`, `]`, as well as any ascii symbolic characters.
	//
	// Examples:
	// (Note: Comments here are written in markdown syntax, so there is an
	//  additional layer of backticks to represent a code block)
	// `\`address.city\`` represents a field named `address.city`, not the map key
	// `city` in the field `address`.
	// `\`*\`` represents a field named `*`, not any field.
	//
	// A special `Field` contains the default indexing settings for all fields.
	// This field's resource name is:
	// `projects/{project_id}/databases/{database_id}/collectionGroups/__default__/fields/*`
	// Indexes defined on this `Field` will be applied to all fields which do not
	// have their own `Field` index configuration.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The index configuration for this field. If unset, field indexing will
	// revert to the configuration defined by the `ancestor_field`. To
	// explicitly remove all indexes for this field, specify an index config
	// with an empty list of indexes.
	IndexConfig *Field_IndexConfig `protobuf:"bytes,2,opt,name=index_config,json=indexConfig,proto3" json:"index_config,omitempty"`
}

func (x *Field) Reset() {
	*x = Field{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_firestore_admin_v1beta2_field_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Field) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Field) ProtoMessage() {}

func (x *Field) ProtoReflect() protoreflect.Message {
	mi := &file_google_firestore_admin_v1beta2_field_proto_msgTypes[0]
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
	return file_google_firestore_admin_v1beta2_field_proto_rawDescGZIP(), []int{0}
}

func (x *Field) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Field) GetIndexConfig() *Field_IndexConfig {
	if x != nil {
		return x.IndexConfig
	}
	return nil
}

// The index configuration for this field.
type Field_IndexConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The indexes supported for this field.
	Indexes []*Index `protobuf:"bytes,1,rep,name=indexes,proto3" json:"indexes,omitempty"`
	// Output only. When true, the `Field`'s index configuration is set from the
	// configuration specified by the `ancestor_field`.
	// When false, the `Field`'s index configuration is defined explicitly.
	UsesAncestorConfig bool `protobuf:"varint,2,opt,name=uses_ancestor_config,json=usesAncestorConfig,proto3" json:"uses_ancestor_config,omitempty"`
	// Output only. Specifies the resource name of the `Field` from which this field's
	// index configuration is set (when `uses_ancestor_config` is true),
	// or from which it *would* be set if this field had no index configuration
	// (when `uses_ancestor_config` is false).
	AncestorField string `protobuf:"bytes,3,opt,name=ancestor_field,json=ancestorField,proto3" json:"ancestor_field,omitempty"`
	// Output only
	// When true, the `Field`'s index configuration is in the process of being
	// reverted. Once complete, the index config will transition to the same
	// state as the field specified by `ancestor_field`, at which point
	// `uses_ancestor_config` will be `true` and `reverting` will be `false`.
	Reverting bool `protobuf:"varint,4,opt,name=reverting,proto3" json:"reverting,omitempty"`
}

func (x *Field_IndexConfig) Reset() {
	*x = Field_IndexConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_firestore_admin_v1beta2_field_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Field_IndexConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Field_IndexConfig) ProtoMessage() {}

func (x *Field_IndexConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_firestore_admin_v1beta2_field_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Field_IndexConfig.ProtoReflect.Descriptor instead.
func (*Field_IndexConfig) Descriptor() ([]byte, []int) {
	return file_google_firestore_admin_v1beta2_field_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Field_IndexConfig) GetIndexes() []*Index {
	if x != nil {
		return x.Indexes
	}
	return nil
}

func (x *Field_IndexConfig) GetUsesAncestorConfig() bool {
	if x != nil {
		return x.UsesAncestorConfig
	}
	return false
}

func (x *Field_IndexConfig) GetAncestorField() string {
	if x != nil {
		return x.AncestorField
	}
	return ""
}

func (x *Field_IndexConfig) GetReverting() bool {
	if x != nil {
		return x.Reverting
	}
	return false
}

var File_google_firestore_admin_v1beta2_field_proto protoreflect.FileDescriptor

var file_google_firestore_admin_v1beta2_field_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x66, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32,
	0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x1a, 0x2a, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x66, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2f, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb9, 0x02, 0x0a, 0x05, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x54, 0x0a, 0x0c, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x66, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2e, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0b, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0xc5, 0x01, 0x0a, 0x0b, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3f, 0x0a, 0x07, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2e, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x52, 0x07, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x12, 0x30, 0x0a, 0x14, 0x75,
	0x73, 0x65, 0x73, 0x5f, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x75, 0x73, 0x65, 0x73, 0x41,
	0x6e, 0x63, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x25, 0x0a,
	0x0e, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x76, 0x65, 0x72, 0x74, 0x69, 0x6e,
	0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x72, 0x65, 0x76, 0x65, 0x72, 0x74, 0x69,
	0x6e, 0x67, 0x42, 0xa5, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x66, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x42, 0x0a, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x66,
	0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x3b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0xa2, 0x02, 0x04, 0x47,
	0x43, 0x46, 0x53, 0xaa, 0x02, 0x24, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x46, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x56, 0x31, 0x42, 0x65, 0x74, 0x61, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_google_firestore_admin_v1beta2_field_proto_rawDescOnce sync.Once
	file_google_firestore_admin_v1beta2_field_proto_rawDescData = file_google_firestore_admin_v1beta2_field_proto_rawDesc
)

func file_google_firestore_admin_v1beta2_field_proto_rawDescGZIP() []byte {
	file_google_firestore_admin_v1beta2_field_proto_rawDescOnce.Do(func() {
		file_google_firestore_admin_v1beta2_field_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_firestore_admin_v1beta2_field_proto_rawDescData)
	})
	return file_google_firestore_admin_v1beta2_field_proto_rawDescData
}

var file_google_firestore_admin_v1beta2_field_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_firestore_admin_v1beta2_field_proto_goTypes = []interface{}{
	(*Field)(nil),             // 0: google.firestore.admin.v1beta2.Field
	(*Field_IndexConfig)(nil), // 1: google.firestore.admin.v1beta2.Field.IndexConfig
	(*Index)(nil),             // 2: google.firestore.admin.v1beta2.Index
}
var file_google_firestore_admin_v1beta2_field_proto_depIdxs = []int32{
	1, // 0: google.firestore.admin.v1beta2.Field.index_config:type_name -> google.firestore.admin.v1beta2.Field.IndexConfig
	2, // 1: google.firestore.admin.v1beta2.Field.IndexConfig.indexes:type_name -> google.firestore.admin.v1beta2.Index
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_firestore_admin_v1beta2_field_proto_init() }
func file_google_firestore_admin_v1beta2_field_proto_init() {
	if File_google_firestore_admin_v1beta2_field_proto != nil {
		return
	}
	file_google_firestore_admin_v1beta2_index_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_firestore_admin_v1beta2_field_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_google_firestore_admin_v1beta2_field_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Field_IndexConfig); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_firestore_admin_v1beta2_field_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_firestore_admin_v1beta2_field_proto_goTypes,
		DependencyIndexes: file_google_firestore_admin_v1beta2_field_proto_depIdxs,
		MessageInfos:      file_google_firestore_admin_v1beta2_field_proto_msgTypes,
	}.Build()
	File_google_firestore_admin_v1beta2_field_proto = out.File
	file_google_firestore_admin_v1beta2_field_proto_rawDesc = nil
	file_google_firestore_admin_v1beta2_field_proto_goTypes = nil
	file_google_firestore_admin_v1beta2_field_proto_depIdxs = nil
}
