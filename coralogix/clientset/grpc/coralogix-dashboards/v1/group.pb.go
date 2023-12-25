// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: com/coralogixapis/dashboards/v1/common/group.proto

package __

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Group struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field  *FieldGroup             `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	Value  *wrapperspb.DoubleValue `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Groups []*Group                `protobuf:"bytes,3,rep,name=groups,proto3" json:"groups,omitempty"`
}

func (x *Group) Reset() {
	*x = Group{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_dashboards_v1_common_group_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Group) ProtoMessage() {}

func (x *Group) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_common_group_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Group.ProtoReflect.Descriptor instead.
func (*Group) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_common_group_proto_rawDescGZIP(), []int{0}
}

func (x *Group) GetField() *FieldGroup {
	if x != nil {
		return x.Field
	}
	return nil
}

func (x *Group) GetValue() *wrapperspb.DoubleValue {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *Group) GetGroups() []*Group {
	if x != nil {
		return x.Groups
	}
	return nil
}

type MultiGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fields []*FieldGroup             `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields,omitempty"`
	Values []*wrapperspb.DoubleValue `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *MultiGroup) Reset() {
	*x = MultiGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_dashboards_v1_common_group_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiGroup) ProtoMessage() {}

func (x *MultiGroup) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_common_group_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiGroup.ProtoReflect.Descriptor instead.
func (*MultiGroup) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_common_group_proto_rawDescGZIP(), []int{1}
}

func (x *MultiGroup) GetFields() []*FieldGroup {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *MultiGroup) GetValues() []*wrapperspb.DoubleValue {
	if x != nil {
		return x.Values
	}
	return nil
}

type FieldGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *FieldGroup) Reset() {
	*x = FieldGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_dashboards_v1_common_group_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldGroup) ProtoMessage() {}

func (x *FieldGroup) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_common_group_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldGroup.ProtoReflect.Descriptor instead.
func (*FieldGroup) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_common_group_proto_rawDescGZIP(), []int{2}
}

func (x *FieldGroup) GetName() *wrapperspb.StringValue {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *FieldGroup) GetValue() *wrapperspb.StringValue {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_com_coralogixapis_dashboards_v1_common_group_proto protoreflect.FileDescriptor

var file_com_coralogixapis_dashboards_v1_common_group_proto_rawDesc = []byte{
	0x0a, 0x32, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72,
	0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcc, 0x01, 0x0a,
	0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x48, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x12, 0x32, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x45, 0x0a, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x52, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x22, 0x8e, 0x01, 0x0a, 0x0a,
	0x4d, 0x75, 0x6c, 0x74, 0x69, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x4a, 0x0a, 0x06, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64,
	0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x06,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x34, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x72, 0x0a, 0x0a,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x30, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x42, 0x03, 0x5a, 0x01, 0x2e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_dashboards_v1_common_group_proto_rawDescOnce sync.Once
	file_com_coralogixapis_dashboards_v1_common_group_proto_rawDescData = file_com_coralogixapis_dashboards_v1_common_group_proto_rawDesc
)

func file_com_coralogixapis_dashboards_v1_common_group_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_dashboards_v1_common_group_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_dashboards_v1_common_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_dashboards_v1_common_group_proto_rawDescData)
	})
	return file_com_coralogixapis_dashboards_v1_common_group_proto_rawDescData
}

var file_com_coralogixapis_dashboards_v1_common_group_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_com_coralogixapis_dashboards_v1_common_group_proto_goTypes = []interface{}{
	(*Group)(nil),                  // 0: com.coralogixapis.dashboards.v1.common.Group
	(*MultiGroup)(nil),             // 1: com.coralogixapis.dashboards.v1.common.MultiGroup
	(*FieldGroup)(nil),             // 2: com.coralogixapis.dashboards.v1.common.FieldGroup
	(*wrapperspb.DoubleValue)(nil), // 3: google.protobuf.DoubleValue
	(*wrapperspb.StringValue)(nil), // 4: google.protobuf.StringValue
}
var file_com_coralogixapis_dashboards_v1_common_group_proto_depIdxs = []int32{
	2, // 0: com.coralogixapis.dashboards.v1.common.Group.field:type_name -> com.coralogixapis.dashboards.v1.common.FieldGroup
	3, // 1: com.coralogixapis.dashboards.v1.common.Group.value:type_name -> google.protobuf.DoubleValue
	0, // 2: com.coralogixapis.dashboards.v1.common.Group.groups:type_name -> com.coralogixapis.dashboards.v1.common.Group
	2, // 3: com.coralogixapis.dashboards.v1.common.MultiGroup.fields:type_name -> com.coralogixapis.dashboards.v1.common.FieldGroup
	3, // 4: com.coralogixapis.dashboards.v1.common.MultiGroup.values:type_name -> google.protobuf.DoubleValue
	4, // 5: com.coralogixapis.dashboards.v1.common.FieldGroup.name:type_name -> google.protobuf.StringValue
	4, // 6: com.coralogixapis.dashboards.v1.common.FieldGroup.value:type_name -> google.protobuf.StringValue
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_dashboards_v1_common_group_proto_init() }
func file_com_coralogixapis_dashboards_v1_common_group_proto_init() {
	if File_com_coralogixapis_dashboards_v1_common_group_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_com_coralogixapis_dashboards_v1_common_group_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Group); i {
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
		file_com_coralogixapis_dashboards_v1_common_group_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiGroup); i {
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
		file_com_coralogixapis_dashboards_v1_common_group_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FieldGroup); i {
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
			RawDescriptor: file_com_coralogixapis_dashboards_v1_common_group_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_dashboards_v1_common_group_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_dashboards_v1_common_group_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_dashboards_v1_common_group_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_dashboards_v1_common_group_proto = out.File
	file_com_coralogixapis_dashboards_v1_common_group_proto_rawDesc = nil
	file_com_coralogixapis_dashboards_v1_common_group_proto_goTypes = nil
	file_com_coralogixapis_dashboards_v1_common_group_proto_depIdxs = nil
}
