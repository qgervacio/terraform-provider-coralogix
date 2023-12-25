// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: com/coralogixapis/dashboards/v1/common/time_frame.proto

package __

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TimeFrame struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To   *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
}

func (x *TimeFrame) Reset() {
	*x = TimeFrame{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_dashboards_v1_common_time_frame_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeFrame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeFrame) ProtoMessage() {}

func (x *TimeFrame) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_common_time_frame_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeFrame.ProtoReflect.Descriptor instead.
func (*TimeFrame) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_common_time_frame_proto_rawDescGZIP(), []int{0}
}

func (x *TimeFrame) GetFrom() *timestamppb.Timestamp {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *TimeFrame) GetTo() *timestamppb.Timestamp {
	if x != nil {
		return x.To
	}
	return nil
}

var File_com_coralogixapis_dashboards_v1_common_time_frame_proto protoreflect.FileDescriptor

var file_com_coralogixapis_dashboards_v1_common_time_frame_proto_rawDesc = []byte{
	0x0a, 0x37, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x66, 0x72,
	0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73,
	0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x67, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x12,
	0x2e, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12,
	0x2a, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x6f, 0x42, 0x03, 0x5a, 0x01, 0x2e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_dashboards_v1_common_time_frame_proto_rawDescOnce sync.Once
	file_com_coralogixapis_dashboards_v1_common_time_frame_proto_rawDescData = file_com_coralogixapis_dashboards_v1_common_time_frame_proto_rawDesc
)

func file_com_coralogixapis_dashboards_v1_common_time_frame_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_dashboards_v1_common_time_frame_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_dashboards_v1_common_time_frame_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_dashboards_v1_common_time_frame_proto_rawDescData)
	})
	return file_com_coralogixapis_dashboards_v1_common_time_frame_proto_rawDescData
}

var file_com_coralogixapis_dashboards_v1_common_time_frame_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogixapis_dashboards_v1_common_time_frame_proto_goTypes = []interface{}{
	(*TimeFrame)(nil),             // 0: com.coralogixapis.dashboards.v1.common.TimeFrame
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_com_coralogixapis_dashboards_v1_common_time_frame_proto_depIdxs = []int32{
	1, // 0: com.coralogixapis.dashboards.v1.common.TimeFrame.from:type_name -> google.protobuf.Timestamp
	1, // 1: com.coralogixapis.dashboards.v1.common.TimeFrame.to:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_dashboards_v1_common_time_frame_proto_init() }
func file_com_coralogixapis_dashboards_v1_common_time_frame_proto_init() {
	if File_com_coralogixapis_dashboards_v1_common_time_frame_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_com_coralogixapis_dashboards_v1_common_time_frame_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeFrame); i {
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
			RawDescriptor: file_com_coralogixapis_dashboards_v1_common_time_frame_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_dashboards_v1_common_time_frame_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_dashboards_v1_common_time_frame_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_dashboards_v1_common_time_frame_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_dashboards_v1_common_time_frame_proto = out.File
	file_com_coralogixapis_dashboards_v1_common_time_frame_proto_rawDesc = nil
	file_com_coralogixapis_dashboards_v1_common_time_frame_proto_goTypes = nil
	file_com_coralogixapis_dashboards_v1_common_time_frame_proto_depIdxs = nil
}
