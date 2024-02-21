// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: com/coralogixapis/aaa/rbac/v2/permissions.proto

package __

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

type ListAllPermissionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListAllPermissionsRequest) Reset() {
	*x = ListAllPermissionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_aaa_rbac_v2_permissions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAllPermissionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAllPermissionsRequest) ProtoMessage() {}

func (x *ListAllPermissionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_aaa_rbac_v2_permissions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAllPermissionsRequest.ProtoReflect.Descriptor instead.
func (*ListAllPermissionsRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_aaa_rbac_v2_permissions_proto_rawDescGZIP(), []int{0}
}

type Permission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Expression  *string `protobuf:"bytes,1,opt,name=expression,proto3,oneof" json:"expression,omitempty"`
	Description *string `protobuf:"bytes,2,opt,name=description,proto3,oneof" json:"description,omitempty"`
	DocLink     *string `protobuf:"bytes,3,opt,name=doc_link,json=docLink,proto3,oneof" json:"doc_link,omitempty"`
	Explanation *string `protobuf:"bytes,4,opt,name=explanation,proto3,oneof" json:"explanation,omitempty"`
}

func (x *Permission) Reset() {
	*x = Permission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_aaa_rbac_v2_permissions_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Permission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Permission) ProtoMessage() {}

func (x *Permission) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_aaa_rbac_v2_permissions_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Permission.ProtoReflect.Descriptor instead.
func (*Permission) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_aaa_rbac_v2_permissions_proto_rawDescGZIP(), []int{1}
}

func (x *Permission) GetExpression() string {
	if x != nil && x.Expression != nil {
		return *x.Expression
	}
	return ""
}

func (x *Permission) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *Permission) GetDocLink() string {
	if x != nil && x.DocLink != nil {
		return *x.DocLink
	}
	return ""
}

func (x *Permission) GetExplanation() string {
	if x != nil && x.Explanation != nil {
		return *x.Explanation
	}
	return ""
}

type ListAllPermissionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Permissions []*Permission `protobuf:"bytes,1,rep,name=permissions,proto3" json:"permissions,omitempty"`
}

func (x *ListAllPermissionsResponse) Reset() {
	*x = ListAllPermissionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_aaa_rbac_v2_permissions_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAllPermissionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAllPermissionsResponse) ProtoMessage() {}

func (x *ListAllPermissionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_aaa_rbac_v2_permissions_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAllPermissionsResponse.ProtoReflect.Descriptor instead.
func (*ListAllPermissionsResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_aaa_rbac_v2_permissions_proto_rawDescGZIP(), []int{2}
}

func (x *ListAllPermissionsResponse) GetPermissions() []*Permission {
	if x != nil {
		return x.Permissions
	}
	return nil
}

var File_com_coralogixapis_aaa_rbac_v2_permissions_proto protoreflect.FileDescriptor

var file_com_coralogixapis_aaa_rbac_v2_permissions_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x61, 0x61, 0x2f, 0x72, 0x62, 0x61, 0x63, 0x2f, 0x76, 0x32, 0x2f,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x61, 0x61, 0x2e, 0x72, 0x62, 0x61, 0x63, 0x2e, 0x76, 0x32,
	0x22, 0x1b, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xdb, 0x01,
	0x0a, 0x0a, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0a,
	0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x88, 0x01,
	0x01, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x1e, 0x0a, 0x08, 0x64, 0x6f, 0x63, 0x5f,
	0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x07, 0x64, 0x6f,
	0x63, 0x4c, 0x69, 0x6e, 0x6b, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x6c,
	0x61, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52,
	0x0b, 0x65, 0x78, 0x70, 0x6c, 0x61, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42,
	0x0d, 0x0a, 0x0b, 0x5f, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x0e,
	0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0b,
	0x0a, 0x09, 0x5f, 0x64, 0x6f, 0x63, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x42, 0x0e, 0x0a, 0x0c, 0x5f,
	0x65, 0x78, 0x70, 0x6c, 0x61, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x69, 0x0a, 0x1a, 0x4c,
	0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x0b, 0x70, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x61, 0x61, 0x61, 0x2e, 0x72, 0x62, 0x61, 0x63, 0x2e, 0x76, 0x32, 0x2e, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0xa0, 0x01, 0x0a, 0x12, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x89, 0x01,
	0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x38, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x61, 0x61, 0x2e, 0x72, 0x62, 0x61,
	0x63, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x39,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x61, 0x61, 0x61, 0x2e, 0x72, 0x62, 0x61, 0x63, 0x2e, 0x76, 0x32, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_aaa_rbac_v2_permissions_proto_rawDescOnce sync.Once
	file_com_coralogixapis_aaa_rbac_v2_permissions_proto_rawDescData = file_com_coralogixapis_aaa_rbac_v2_permissions_proto_rawDesc
)

func file_com_coralogixapis_aaa_rbac_v2_permissions_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_aaa_rbac_v2_permissions_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_aaa_rbac_v2_permissions_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_aaa_rbac_v2_permissions_proto_rawDescData)
	})
	return file_com_coralogixapis_aaa_rbac_v2_permissions_proto_rawDescData
}

var file_com_coralogixapis_aaa_rbac_v2_permissions_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_com_coralogixapis_aaa_rbac_v2_permissions_proto_goTypes = []interface{}{
	(*ListAllPermissionsRequest)(nil),  // 0: com.coralogixapis.aaa.rbac.v2.ListAllPermissionsRequest
	(*Permission)(nil),                 // 1: com.coralogixapis.aaa.rbac.v2.Permission
	(*ListAllPermissionsResponse)(nil), // 2: com.coralogixapis.aaa.rbac.v2.ListAllPermissionsResponse
}
var file_com_coralogixapis_aaa_rbac_v2_permissions_proto_depIdxs = []int32{
	1, // 0: com.coralogixapis.aaa.rbac.v2.ListAllPermissionsResponse.permissions:type_name -> com.coralogixapis.aaa.rbac.v2.Permission
	0, // 1: com.coralogixapis.aaa.rbac.v2.PermissionsService.ListAllPermissions:input_type -> com.coralogixapis.aaa.rbac.v2.ListAllPermissionsRequest
	2, // 2: com.coralogixapis.aaa.rbac.v2.PermissionsService.ListAllPermissions:output_type -> com.coralogixapis.aaa.rbac.v2.ListAllPermissionsResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_aaa_rbac_v2_permissions_proto_init() }
func file_com_coralogixapis_aaa_rbac_v2_permissions_proto_init() {
	if File_com_coralogixapis_aaa_rbac_v2_permissions_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_com_coralogixapis_aaa_rbac_v2_permissions_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAllPermissionsRequest); i {
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
		file_com_coralogixapis_aaa_rbac_v2_permissions_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Permission); i {
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
		file_com_coralogixapis_aaa_rbac_v2_permissions_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAllPermissionsResponse); i {
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
	file_com_coralogixapis_aaa_rbac_v2_permissions_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_aaa_rbac_v2_permissions_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_coralogixapis_aaa_rbac_v2_permissions_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_aaa_rbac_v2_permissions_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_aaa_rbac_v2_permissions_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_aaa_rbac_v2_permissions_proto = out.File
	file_com_coralogixapis_aaa_rbac_v2_permissions_proto_rawDesc = nil
	file_com_coralogixapis_aaa_rbac_v2_permissions_proto_goTypes = nil
	file_com_coralogixapis_aaa_rbac_v2_permissions_proto_depIdxs = nil
}
