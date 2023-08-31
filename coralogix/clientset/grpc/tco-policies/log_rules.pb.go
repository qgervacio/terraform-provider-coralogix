// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: com/coralogix/quota/v1/log_rules.proto

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

type LogRules struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Severities []Severity `protobuf:"varint,1,rep,packed,name=severities,proto3,enum=com.coralogix.quota.v1.Severity" json:"severities,omitempty"`
}

func (x *LogRules) Reset() {
	*x = LogRules{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogix_quota_v1_log_rules_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogRules) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogRules) ProtoMessage() {}

func (x *LogRules) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogix_quota_v1_log_rules_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogRules.ProtoReflect.Descriptor instead.
func (*LogRules) Descriptor() ([]byte, []int) {
	return file_com_coralogix_quota_v1_log_rules_proto_rawDescGZIP(), []int{0}
}

func (x *LogRules) GetSeverities() []Severity {
	if x != nil {
		return x.Severities
	}
	return nil
}

var File_com_coralogix_quota_v1_log_rules_proto protoreflect.FileDescriptor

var file_com_coralogix_quota_v1_log_rules_proto_rawDesc = []byte{
	0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x71, 0x75, 0x6f, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x5f, 0x72, 0x75, 0x6c,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2e, 0x76, 0x31,
	0x1a, 0x22, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x2f,
	0x71, 0x75, 0x6f, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4c, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x52, 0x75, 0x6c, 0x65, 0x73,
	0x12, 0x40, 0x0a, 0x0a, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65,
	0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x52, 0x0a, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogix_quota_v1_log_rules_proto_rawDescOnce sync.Once
	file_com_coralogix_quota_v1_log_rules_proto_rawDescData = file_com_coralogix_quota_v1_log_rules_proto_rawDesc
)

func file_com_coralogix_quota_v1_log_rules_proto_rawDescGZIP() []byte {
	file_com_coralogix_quota_v1_log_rules_proto_rawDescOnce.Do(func() {
		file_com_coralogix_quota_v1_log_rules_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogix_quota_v1_log_rules_proto_rawDescData)
	})
	return file_com_coralogix_quota_v1_log_rules_proto_rawDescData
}

var file_com_coralogix_quota_v1_log_rules_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogix_quota_v1_log_rules_proto_goTypes = []interface{}{
	(*LogRules)(nil), // 0: com.coralogix.quota.v1.LogRules
	(Severity)(0),    // 1: com.coralogix.quota.v1.Severity
}
var file_com_coralogix_quota_v1_log_rules_proto_depIdxs = []int32{
	1, // 0: com.coralogix.quota.v1.LogRules.severities:type_name -> com.coralogix.quota.v1.Severity
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_com_coralogix_quota_v1_log_rules_proto_init() }
func file_com_coralogix_quota_v1_log_rules_proto_init() {
	if File_com_coralogix_quota_v1_log_rules_proto != nil {
		return
	}
	file_com_coralogix_quota_v1_enums_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_com_coralogix_quota_v1_log_rules_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogRules); i {
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
			RawDescriptor: file_com_coralogix_quota_v1_log_rules_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogix_quota_v1_log_rules_proto_goTypes,
		DependencyIndexes: file_com_coralogix_quota_v1_log_rules_proto_depIdxs,
		MessageInfos:      file_com_coralogix_quota_v1_log_rules_proto_msgTypes,
	}.Build()
	File_com_coralogix_quota_v1_log_rules_proto = out.File
	file_com_coralogix_quota_v1_log_rules_proto_rawDesc = nil
	file_com_coralogix_quota_v1_log_rules_proto_goTypes = nil
	file_com_coralogix_quota_v1_log_rules_proto_depIdxs = nil
}