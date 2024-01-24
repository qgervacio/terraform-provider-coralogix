// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: com/coralogixapis/alerting/alert_scheduler_rule_protobuf/v1/alert_scheduler_rule.proto

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

type AlertSchedulerRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UniqueIdentifier *string      `protobuf:"bytes,1,opt,name=unique_identifier,json=uniqueIdentifier,proto3,oneof" json:"unique_identifier,omitempty"`
	Id               *string      `protobuf:"bytes,2,opt,name=id,proto3,oneof" json:"id,omitempty"`
	Name             string       `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description      *string      `protobuf:"bytes,4,opt,name=description,proto3,oneof" json:"description,omitempty"`
	MetaLabels       []*MetaLabel `protobuf:"bytes,5,rep,name=meta_labels,json=metaLabels,proto3" json:"meta_labels,omitempty"`
	Filter           *Filter      `protobuf:"bytes,6,opt,name=filter,proto3" json:"filter,omitempty"`
	Schedule         *Schedule    `protobuf:"bytes,7,opt,name=schedule,proto3" json:"schedule,omitempty"`
	Enabled          bool         `protobuf:"varint,8,opt,name=enabled,proto3" json:"enabled,omitempty"`
	CreatedAt        *string      `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3,oneof" json:"created_at,omitempty"`
	UpdatedAt        *string      `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3,oneof" json:"updated_at,omitempty"`
}

func (x *AlertSchedulerRule) Reset() {
	*x = AlertSchedulerRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlertSchedulerRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertSchedulerRule) ProtoMessage() {}

func (x *AlertSchedulerRule) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertSchedulerRule.ProtoReflect.Descriptor instead.
func (*AlertSchedulerRule) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_proto_rawDescGZIP(), []int{0}
}

func (x *AlertSchedulerRule) GetUniqueIdentifier() string {
	if x != nil && x.UniqueIdentifier != nil {
		return *x.UniqueIdentifier
	}
	return ""
}

func (x *AlertSchedulerRule) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *AlertSchedulerRule) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AlertSchedulerRule) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *AlertSchedulerRule) GetMetaLabels() []*MetaLabel {
	if x != nil {
		return x.MetaLabels
	}
	return nil
}

func (x *AlertSchedulerRule) GetFilter() *Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *AlertSchedulerRule) GetSchedule() *Schedule {
	if x != nil {
		return x.Schedule
	}
	return nil
}

func (x *AlertSchedulerRule) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *AlertSchedulerRule) GetCreatedAt() string {
	if x != nil && x.CreatedAt != nil {
		return *x.CreatedAt
	}
	return ""
}

func (x *AlertSchedulerRule) GetUpdatedAt() string {
	if x != nil && x.UpdatedAt != nil {
		return *x.UpdatedAt
	}
	return ""
}

var File_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_proto protoreflect.FileDescriptor

var file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_proto_rawDesc = []byte{
	0x0a, 0x56, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x75, 0x6c,
	0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x75,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x3b, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x76, 0x31, 0x1a, 0x44, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e,
	0x67, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x58, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x4a, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e,
	0x67, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x72, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xe3, 0x04, 0x0a, 0x12, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x72, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x30, 0x0a, 0x11, 0x75, 0x6e, 0x69, 0x71,
	0x75, 0x65, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x10, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x5e, 0x0a, 0x0b, 0x6d, 0x65,
	0x74, 0x61, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x3d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x65, 0x74,
	0x61, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x0a,
	0x6d, 0x65, 0x74, 0x61, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x5b, 0x0a, 0x06, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x43, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52,
	0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x61, 0x0a, 0x08, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x45, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x73, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x52, 0x08, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x12, 0x22, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x42, 0x14, 0x0a, 0x12,
	0x5f, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_proto_rawDescOnce sync.Once
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_proto_rawDescData = file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_proto_rawDesc
)

func file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_proto_rawDescData)
	})
	return file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_proto_rawDescData
}

var file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_proto_goTypes = []interface{}{
	(*AlertSchedulerRule)(nil), // 0: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.AlertSchedulerRule
	(*MetaLabel)(nil),          // 1: com.coralogixapis.alerting.meta_labels_protobuf.v1.MetaLabel
	(*Filter)(nil),             // 2: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.Filter
	(*Schedule)(nil),           // 3: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.Schedule
}
var file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_proto_depIdxs = []int32{
	1, // 0: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.AlertSchedulerRule.meta_labels:type_name -> com.coralogixapis.alerting.meta_labels_protobuf.v1.MetaLabel
	2, // 1: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.AlertSchedulerRule.filter:type_name -> com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.Filter
	3, // 2: com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.AlertSchedulerRule.schedule:type_name -> com.coralogixapis.alerting.alert_scheduler_rule_protobuf.v1.Schedule
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() {
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_proto_init()
}
func file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_proto_init() {
	if File_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_proto != nil {
		return
	}
	file_com_coralogixapis_alerting_meta_labels_protobuf_v1_meta_labels_proto_init()
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_filter_proto_init()
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_schedule_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlertSchedulerRule); i {
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
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_proto = out.File
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_proto_rawDesc = nil
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_proto_goTypes = nil
	file_com_coralogixapis_alerting_alert_scheduler_rule_protobuf_v1_alert_scheduler_rule_proto_depIdxs = nil
}
