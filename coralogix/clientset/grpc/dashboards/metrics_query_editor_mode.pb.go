// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: com/coralogixapis/dashboards/v1/ast/widgets/common/metrics_query_editor_mode.proto

package v1

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

type MetricsQueryEditorMode int32

const (
	MetricsQueryEditorMode_METRICS_QUERY_EDITOR_MODE_UNSPECIFIED MetricsQueryEditorMode = 0
	MetricsQueryEditorMode_METRICS_QUERY_EDITOR_MODE_TEXT        MetricsQueryEditorMode = 1
	MetricsQueryEditorMode_METRICS_QUERY_EDITOR_MODE_BUILDER     MetricsQueryEditorMode = 2
)

// Enum value maps for MetricsQueryEditorMode.
var (
	MetricsQueryEditorMode_name = map[int32]string{
		0: "METRICS_QUERY_EDITOR_MODE_UNSPECIFIED",
		1: "METRICS_QUERY_EDITOR_MODE_TEXT",
		2: "METRICS_QUERY_EDITOR_MODE_BUILDER",
	}
	MetricsQueryEditorMode_value = map[string]int32{
		"METRICS_QUERY_EDITOR_MODE_UNSPECIFIED": 0,
		"METRICS_QUERY_EDITOR_MODE_TEXT":        1,
		"METRICS_QUERY_EDITOR_MODE_BUILDER":     2,
	}
)

func (x MetricsQueryEditorMode) Enum() *MetricsQueryEditorMode {
	p := new(MetricsQueryEditorMode)
	*p = x
	return p
}

func (x MetricsQueryEditorMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MetricsQueryEditorMode) Descriptor() protoreflect.EnumDescriptor {
	return file_com_coralogixapis_dashboards_v1_ast_widgets_common_metrics_query_editor_mode_proto_enumTypes[0].Descriptor()
}

func (MetricsQueryEditorMode) Type() protoreflect.EnumType {
	return &file_com_coralogixapis_dashboards_v1_ast_widgets_common_metrics_query_editor_mode_proto_enumTypes[0]
}

func (x MetricsQueryEditorMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MetricsQueryEditorMode.Descriptor instead.
func (MetricsQueryEditorMode) EnumDescriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_ast_widgets_common_metrics_query_editor_mode_proto_rawDescGZIP(), []int{0}
}

var File_com_coralogixapis_dashboards_v1_ast_widgets_common_metrics_query_editor_mode_proto protoreflect.FileDescriptor

var file_com_coralogixapis_dashboards_v1_ast_widgets_common_metrics_query_editor_mode_proto_rawDesc = []byte{
	0x0a, 0x52, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x73, 0x74, 0x2f, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x5f, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x5f, 0x65, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x32, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x73, 0x74, 0x2e, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2a, 0x8e, 0x01, 0x0a, 0x16, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x45, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x4d,
	0x6f, 0x64, 0x65, 0x12, 0x29, 0x0a, 0x25, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x53, 0x5f, 0x51,
	0x55, 0x45, 0x52, 0x59, 0x5f, 0x45, 0x44, 0x49, 0x54, 0x4f, 0x52, 0x5f, 0x4d, 0x4f, 0x44, 0x45,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x22,
	0x0a, 0x1e, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x53, 0x5f, 0x51, 0x55, 0x45, 0x52, 0x59, 0x5f,
	0x45, 0x44, 0x49, 0x54, 0x4f, 0x52, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x54, 0x45, 0x58, 0x54,
	0x10, 0x01, 0x12, 0x25, 0x0a, 0x21, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x53, 0x5f, 0x51, 0x55,
	0x45, 0x52, 0x59, 0x5f, 0x45, 0x44, 0x49, 0x54, 0x4f, 0x52, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f,
	0x42, 0x55, 0x49, 0x4c, 0x44, 0x45, 0x52, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_com_coralogixapis_dashboards_v1_ast_widgets_common_metrics_query_editor_mode_proto_rawDescOnce sync.Once
	file_com_coralogixapis_dashboards_v1_ast_widgets_common_metrics_query_editor_mode_proto_rawDescData = file_com_coralogixapis_dashboards_v1_ast_widgets_common_metrics_query_editor_mode_proto_rawDesc
)

func file_com_coralogixapis_dashboards_v1_ast_widgets_common_metrics_query_editor_mode_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_dashboards_v1_ast_widgets_common_metrics_query_editor_mode_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_dashboards_v1_ast_widgets_common_metrics_query_editor_mode_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_dashboards_v1_ast_widgets_common_metrics_query_editor_mode_proto_rawDescData)
	})
	return file_com_coralogixapis_dashboards_v1_ast_widgets_common_metrics_query_editor_mode_proto_rawDescData
}

var file_com_coralogixapis_dashboards_v1_ast_widgets_common_metrics_query_editor_mode_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_com_coralogixapis_dashboards_v1_ast_widgets_common_metrics_query_editor_mode_proto_goTypes = []any{
	(MetricsQueryEditorMode)(0), // 0: com.coralogixapis.dashboards.v1.ast.widgets.common.MetricsQueryEditorMode
}
var file_com_coralogixapis_dashboards_v1_ast_widgets_common_metrics_query_editor_mode_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_com_coralogixapis_dashboards_v1_ast_widgets_common_metrics_query_editor_mode_proto_init()
}
func file_com_coralogixapis_dashboards_v1_ast_widgets_common_metrics_query_editor_mode_proto_init() {
	if File_com_coralogixapis_dashboards_v1_ast_widgets_common_metrics_query_editor_mode_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_dashboards_v1_ast_widgets_common_metrics_query_editor_mode_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_dashboards_v1_ast_widgets_common_metrics_query_editor_mode_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_dashboards_v1_ast_widgets_common_metrics_query_editor_mode_proto_depIdxs,
		EnumInfos:         file_com_coralogixapis_dashboards_v1_ast_widgets_common_metrics_query_editor_mode_proto_enumTypes,
	}.Build()
	File_com_coralogixapis_dashboards_v1_ast_widgets_common_metrics_query_editor_mode_proto = out.File
	file_com_coralogixapis_dashboards_v1_ast_widgets_common_metrics_query_editor_mode_proto_rawDesc = nil
	file_com_coralogixapis_dashboards_v1_ast_widgets_common_metrics_query_editor_mode_proto_goTypes = nil
	file_com_coralogixapis_dashboards_v1_ast_widgets_common_metrics_query_editor_mode_proto_depIdxs = nil
}