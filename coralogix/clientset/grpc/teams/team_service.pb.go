// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: com/coralogixapis/aaa/organisations/v2/team_service.proto

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

type CreateTeamInOrgRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamName        string   `protobuf:"bytes,1,opt,name=team_name,json=teamName,proto3" json:"team_name,omitempty"`
	TeamAdminsEmail []string `protobuf:"bytes,3,rep,name=team_admins_email,json=teamAdminsEmail,proto3" json:"team_admins_email,omitempty"`
	Retention       int32    `protobuf:"varint,5,opt,name=retention,proto3" json:"retention,omitempty"`
}

func (x *CreateTeamInOrgRequest) Reset() {
	*x = CreateTeamInOrgRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_aaa_organisations_v2_team_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTeamInOrgRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTeamInOrgRequest) ProtoMessage() {}

func (x *CreateTeamInOrgRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_aaa_organisations_v2_team_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTeamInOrgRequest.ProtoReflect.Descriptor instead.
func (*CreateTeamInOrgRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_aaa_organisations_v2_team_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTeamInOrgRequest) GetTeamName() string {
	if x != nil {
		return x.TeamName
	}
	return ""
}

func (x *CreateTeamInOrgRequest) GetTeamAdminsEmail() []string {
	if x != nil {
		return x.TeamAdminsEmail
	}
	return nil
}

func (x *CreateTeamInOrgRequest) GetRetention() int32 {
	if x != nil {
		return x.Retention
	}
	return 0
}

type CreateTeamInOrgResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamId      *TeamId `protobuf:"bytes,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	SendDataKey string  `protobuf:"bytes,2,opt,name=send_data_key,json=sendDataKey,proto3" json:"send_data_key,omitempty"`
}

func (x *CreateTeamInOrgResponse) Reset() {
	*x = CreateTeamInOrgResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_aaa_organisations_v2_team_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTeamInOrgResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTeamInOrgResponse) ProtoMessage() {}

func (x *CreateTeamInOrgResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_aaa_organisations_v2_team_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTeamInOrgResponse.ProtoReflect.Descriptor instead.
func (*CreateTeamInOrgResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_aaa_organisations_v2_team_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTeamInOrgResponse) GetTeamId() *TeamId {
	if x != nil {
		return x.TeamId
	}
	return nil
}

func (x *CreateTeamInOrgResponse) GetSendDataKey() string {
	if x != nil {
		return x.SendDataKey
	}
	return ""
}

type MoveQuotaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SourceTeam      *TeamId `protobuf:"bytes,1,opt,name=source_team,json=sourceTeam,proto3" json:"source_team,omitempty"`
	DestinationTeam *TeamId `protobuf:"bytes,2,opt,name=destination_team,json=destinationTeam,proto3" json:"destination_team,omitempty"`
	UnitsToMove     float32 `protobuf:"fixed32,3,opt,name=units_to_move,json=unitsToMove,proto3" json:"units_to_move,omitempty"`
}

func (x *MoveQuotaRequest) Reset() {
	*x = MoveQuotaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_aaa_organisations_v2_team_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MoveQuotaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MoveQuotaRequest) ProtoMessage() {}

func (x *MoveQuotaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_aaa_organisations_v2_team_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MoveQuotaRequest.ProtoReflect.Descriptor instead.
func (*MoveQuotaRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_aaa_organisations_v2_team_service_proto_rawDescGZIP(), []int{2}
}

func (x *MoveQuotaRequest) GetSourceTeam() *TeamId {
	if x != nil {
		return x.SourceTeam
	}
	return nil
}

func (x *MoveQuotaRequest) GetDestinationTeam() *TeamId {
	if x != nil {
		return x.DestinationTeam
	}
	return nil
}

func (x *MoveQuotaRequest) GetUnitsToMove() float32 {
	if x != nil {
		return x.UnitsToMove
	}
	return 0
}

type MoveQuotaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SourceTeamQuota      float32 `protobuf:"fixed32,1,opt,name=source_team_quota,json=sourceTeamQuota,proto3" json:"source_team_quota,omitempty"`
	DestinationTeamQuota float32 `protobuf:"fixed32,2,opt,name=destination_team_quota,json=destinationTeamQuota,proto3" json:"destination_team_quota,omitempty"`
}

func (x *MoveQuotaResponse) Reset() {
	*x = MoveQuotaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_aaa_organisations_v2_team_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MoveQuotaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MoveQuotaResponse) ProtoMessage() {}

func (x *MoveQuotaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_aaa_organisations_v2_team_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MoveQuotaResponse.ProtoReflect.Descriptor instead.
func (*MoveQuotaResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_aaa_organisations_v2_team_service_proto_rawDescGZIP(), []int{3}
}

func (x *MoveQuotaResponse) GetSourceTeamQuota() float32 {
	if x != nil {
		return x.SourceTeamQuota
	}
	return 0
}

func (x *MoveQuotaResponse) GetDestinationTeamQuota() float32 {
	if x != nil {
		return x.DestinationTeamQuota
	}
	return 0
}

type InviteUsersToTeamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamId     *TeamId  `protobuf:"bytes,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	Emails     []string `protobuf:"bytes,2,rep,name=emails,proto3" json:"emails,omitempty"`
	GroupNames []string `protobuf:"bytes,3,rep,name=group_names,json=groupNames,proto3" json:"group_names,omitempty"`
}

func (x *InviteUsersToTeamRequest) Reset() {
	*x = InviteUsersToTeamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_aaa_organisations_v2_team_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InviteUsersToTeamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InviteUsersToTeamRequest) ProtoMessage() {}

func (x *InviteUsersToTeamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_aaa_organisations_v2_team_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InviteUsersToTeamRequest.ProtoReflect.Descriptor instead.
func (*InviteUsersToTeamRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_aaa_organisations_v2_team_service_proto_rawDescGZIP(), []int{4}
}

func (x *InviteUsersToTeamRequest) GetTeamId() *TeamId {
	if x != nil {
		return x.TeamId
	}
	return nil
}

func (x *InviteUsersToTeamRequest) GetEmails() []string {
	if x != nil {
		return x.Emails
	}
	return nil
}

func (x *InviteUsersToTeamRequest) GetGroupNames() []string {
	if x != nil {
		return x.GroupNames
	}
	return nil
}

type InviteUsersToTeamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *InviteUsersToTeamResponse) Reset() {
	*x = InviteUsersToTeamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_aaa_organisations_v2_team_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InviteUsersToTeamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InviteUsersToTeamResponse) ProtoMessage() {}

func (x *InviteUsersToTeamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_aaa_organisations_v2_team_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InviteUsersToTeamResponse.ProtoReflect.Descriptor instead.
func (*InviteUsersToTeamResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_aaa_organisations_v2_team_service_proto_rawDescGZIP(), []int{5}
}

type GetTeamQuotaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamId *TeamId `protobuf:"bytes,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
}

func (x *GetTeamQuotaRequest) Reset() {
	*x = GetTeamQuotaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_aaa_organisations_v2_team_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTeamQuotaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTeamQuotaRequest) ProtoMessage() {}

func (x *GetTeamQuotaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_aaa_organisations_v2_team_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTeamQuotaRequest.ProtoReflect.Descriptor instead.
func (*GetTeamQuotaRequest) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_aaa_organisations_v2_team_service_proto_rawDescGZIP(), []int{6}
}

func (x *GetTeamQuotaRequest) GetTeamId() *TeamId {
	if x != nil {
		return x.TeamId
	}
	return nil
}

type GetTeamQuotaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamId    *TeamId `protobuf:"bytes,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	Quota     float32 `protobuf:"fixed32,2,opt,name=quota,proto3" json:"quota,omitempty"`
	Retention int32   `protobuf:"varint,3,opt,name=retention,proto3" json:"retention,omitempty"`
}

func (x *GetTeamQuotaResponse) Reset() {
	*x = GetTeamQuotaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_aaa_organisations_v2_team_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTeamQuotaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTeamQuotaResponse) ProtoMessage() {}

func (x *GetTeamQuotaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_aaa_organisations_v2_team_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTeamQuotaResponse.ProtoReflect.Descriptor instead.
func (*GetTeamQuotaResponse) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_aaa_organisations_v2_team_service_proto_rawDescGZIP(), []int{7}
}

func (x *GetTeamQuotaResponse) GetTeamId() *TeamId {
	if x != nil {
		return x.TeamId
	}
	return nil
}

func (x *GetTeamQuotaResponse) GetQuota() float32 {
	if x != nil {
		return x.Quota
	}
	return 0
}

func (x *GetTeamQuotaResponse) GetRetention() int32 {
	if x != nil {
		return x.Retention
	}
	return 0
}

var File_com_coralogixapis_aaa_organisations_v2_team_service_proto protoreflect.FileDescriptor

var file_com_coralogixapis_aaa_organisations_v2_team_service_proto_rawDesc = []byte{
	0x0a, 0x39, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x61, 0x61, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x73, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x26, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61,
	0x61, 0x61, 0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x76, 0x32, 0x1a, 0x32, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67,
	0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x61, 0x61, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7f, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x6e, 0x4f, 0x72, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2a,
	0x0a, 0x11, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x73, 0x5f, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x74, 0x65, 0x61, 0x6d, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x73, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65,
	0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x72,
	0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x86, 0x01, 0x0a, 0x17, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x6e, 0x4f, 0x72, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x61, 0x61, 0x2e, 0x6f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x54,
	0x65, 0x61, 0x6d, 0x49, 0x64, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x22, 0x0a,
	0x0d, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x61, 0x4b, 0x65,
	0x79, 0x22, 0xe2, 0x01, 0x0a, 0x10, 0x4d, 0x6f, 0x76, 0x65, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4f, 0x0a, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x74, 0x65, 0x61, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x61, 0x61, 0x61, 0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x76, 0x32, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x52, 0x0a, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x12, 0x59, 0x0a, 0x10, 0x64, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x65, 0x61, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x61, 0x61, 0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x49,
	0x64, 0x52, 0x0f, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x65,
	0x61, 0x6d, 0x12, 0x22, 0x0a, 0x0d, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x5f, 0x74, 0x6f, 0x5f, 0x6d,
	0x6f, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x75, 0x6e, 0x69, 0x74, 0x73,
	0x54, 0x6f, 0x4d, 0x6f, 0x76, 0x65, 0x22, 0x75, 0x0a, 0x11, 0x4d, 0x6f, 0x76, 0x65, 0x51, 0x75,
	0x6f, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x71, 0x75, 0x6f, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x65,
	0x61, 0x6d, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x12, 0x34, 0x0a, 0x16, 0x64, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x71, 0x75, 0x6f, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x14, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x65, 0x61, 0x6d, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x22, 0x9c, 0x01,
	0x0a, 0x18, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x54, 0x6f, 0x54,
	0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x47, 0x0a, 0x07, 0x74, 0x65,
	0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x61, 0x61, 0x61, 0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x76, 0x32, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x52, 0x06, 0x74, 0x65, 0x61,
	0x6d, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x06, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0a, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x1b, 0x0a, 0x19,
	0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x54, 0x6f, 0x54, 0x65, 0x61,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5e, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x54, 0x65, 0x61, 0x6d, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x47, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x61, 0x61, 0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x49,
	0x64, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x22, 0x93, 0x01, 0x0a, 0x14, 0x47, 0x65,
	0x74, 0x54, 0x65, 0x61, 0x6d, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x47, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x61, 0x61, 0x2e, 0x6f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x54, 0x65, 0x61,
	0x6d, 0x49, 0x64, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x71,
	0x75, 0x6f, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x71, 0x75, 0x6f, 0x74,
	0x61, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x32,
	0xd4, 0x04, 0x0a, 0x0b, 0x54, 0x65, 0x61, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x94, 0x01, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x6e,
	0x4f, 0x72, 0x67, 0x12, 0x3e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x61, 0x61, 0x2e, 0x6f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x6e, 0x4f, 0x72, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x3f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x61, 0x61, 0x2e, 0x6f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x6e, 0x4f, 0x72, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x82, 0x01, 0x0a, 0x09, 0x4d, 0x6f, 0x76, 0x65, 0x51,
	0x75, 0x6f, 0x74, 0x61, 0x12, 0x38, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c,
	0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x61, 0x61, 0x2e, 0x6f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x4d, 0x6f,
	0x76, 0x65, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x39,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x61, 0x61, 0x61, 0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x73, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x4d, 0x6f, 0x76, 0x65, 0x51, 0x75, 0x6f, 0x74,
	0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x8b, 0x01, 0x0a, 0x0c,
	0x47, 0x65, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x12, 0x3b, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x61, 0x61, 0x61, 0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x73, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x51, 0x75, 0x6f,
	0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x61,
	0x61, 0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x9a, 0x01, 0x0a, 0x11, 0x49, 0x6e,
	0x76, 0x69, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x54, 0x6f, 0x54, 0x65, 0x61, 0x6d, 0x12,
	0x40, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x61, 0x61, 0x61, 0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x73, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x54, 0x6f, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x41, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x61, 0x61, 0x61, 0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x49, 0x6e, 0x76, 0x69, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x54, 0x6f, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_aaa_organisations_v2_team_service_proto_rawDescOnce sync.Once
	file_com_coralogixapis_aaa_organisations_v2_team_service_proto_rawDescData = file_com_coralogixapis_aaa_organisations_v2_team_service_proto_rawDesc
)

func file_com_coralogixapis_aaa_organisations_v2_team_service_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_aaa_organisations_v2_team_service_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_aaa_organisations_v2_team_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_aaa_organisations_v2_team_service_proto_rawDescData)
	})
	return file_com_coralogixapis_aaa_organisations_v2_team_service_proto_rawDescData
}

var file_com_coralogixapis_aaa_organisations_v2_team_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_com_coralogixapis_aaa_organisations_v2_team_service_proto_goTypes = []interface{}{
	(*CreateTeamInOrgRequest)(nil),    // 0: com.coralogixapis.aaa.organisations.v2.CreateTeamInOrgRequest
	(*CreateTeamInOrgResponse)(nil),   // 1: com.coralogixapis.aaa.organisations.v2.CreateTeamInOrgResponse
	(*MoveQuotaRequest)(nil),          // 2: com.coralogixapis.aaa.organisations.v2.MoveQuotaRequest
	(*MoveQuotaResponse)(nil),         // 3: com.coralogixapis.aaa.organisations.v2.MoveQuotaResponse
	(*InviteUsersToTeamRequest)(nil),  // 4: com.coralogixapis.aaa.organisations.v2.InviteUsersToTeamRequest
	(*InviteUsersToTeamResponse)(nil), // 5: com.coralogixapis.aaa.organisations.v2.InviteUsersToTeamResponse
	(*GetTeamQuotaRequest)(nil),       // 6: com.coralogixapis.aaa.organisations.v2.GetTeamQuotaRequest
	(*GetTeamQuotaResponse)(nil),      // 7: com.coralogixapis.aaa.organisations.v2.GetTeamQuotaResponse
	(*TeamId)(nil),                    // 8: com.coralogixapis.aaa.organisations.v2.TeamId
}
var file_com_coralogixapis_aaa_organisations_v2_team_service_proto_depIdxs = []int32{
	8,  // 0: com.coralogixapis.aaa.organisations.v2.CreateTeamInOrgResponse.team_id:type_name -> com.coralogixapis.aaa.organisations.v2.TeamId
	8,  // 1: com.coralogixapis.aaa.organisations.v2.MoveQuotaRequest.source_team:type_name -> com.coralogixapis.aaa.organisations.v2.TeamId
	8,  // 2: com.coralogixapis.aaa.organisations.v2.MoveQuotaRequest.destination_team:type_name -> com.coralogixapis.aaa.organisations.v2.TeamId
	8,  // 3: com.coralogixapis.aaa.organisations.v2.InviteUsersToTeamRequest.team_id:type_name -> com.coralogixapis.aaa.organisations.v2.TeamId
	8,  // 4: com.coralogixapis.aaa.organisations.v2.GetTeamQuotaRequest.team_id:type_name -> com.coralogixapis.aaa.organisations.v2.TeamId
	8,  // 5: com.coralogixapis.aaa.organisations.v2.GetTeamQuotaResponse.team_id:type_name -> com.coralogixapis.aaa.organisations.v2.TeamId
	0,  // 6: com.coralogixapis.aaa.organisations.v2.TeamService.CreateTeamInOrg:input_type -> com.coralogixapis.aaa.organisations.v2.CreateTeamInOrgRequest
	2,  // 7: com.coralogixapis.aaa.organisations.v2.TeamService.MoveQuota:input_type -> com.coralogixapis.aaa.organisations.v2.MoveQuotaRequest
	6,  // 8: com.coralogixapis.aaa.organisations.v2.TeamService.GetTeamQuota:input_type -> com.coralogixapis.aaa.organisations.v2.GetTeamQuotaRequest
	4,  // 9: com.coralogixapis.aaa.organisations.v2.TeamService.InviteUsersToTeam:input_type -> com.coralogixapis.aaa.organisations.v2.InviteUsersToTeamRequest
	1,  // 10: com.coralogixapis.aaa.organisations.v2.TeamService.CreateTeamInOrg:output_type -> com.coralogixapis.aaa.organisations.v2.CreateTeamInOrgResponse
	3,  // 11: com.coralogixapis.aaa.organisations.v2.TeamService.MoveQuota:output_type -> com.coralogixapis.aaa.organisations.v2.MoveQuotaResponse
	7,  // 12: com.coralogixapis.aaa.organisations.v2.TeamService.GetTeamQuota:output_type -> com.coralogixapis.aaa.organisations.v2.GetTeamQuotaResponse
	5,  // 13: com.coralogixapis.aaa.organisations.v2.TeamService.InviteUsersToTeam:output_type -> com.coralogixapis.aaa.organisations.v2.InviteUsersToTeamResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_aaa_organisations_v2_team_service_proto_init() }
func file_com_coralogixapis_aaa_organisations_v2_team_service_proto_init() {
	if File_com_coralogixapis_aaa_organisations_v2_team_service_proto != nil {
		return
	}
	file_com_coralogixapis_aaa_organisations_v2_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_com_coralogixapis_aaa_organisations_v2_team_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTeamInOrgRequest); i {
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
		file_com_coralogixapis_aaa_organisations_v2_team_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTeamInOrgResponse); i {
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
		file_com_coralogixapis_aaa_organisations_v2_team_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MoveQuotaRequest); i {
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
		file_com_coralogixapis_aaa_organisations_v2_team_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MoveQuotaResponse); i {
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
		file_com_coralogixapis_aaa_organisations_v2_team_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InviteUsersToTeamRequest); i {
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
		file_com_coralogixapis_aaa_organisations_v2_team_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InviteUsersToTeamResponse); i {
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
		file_com_coralogixapis_aaa_organisations_v2_team_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTeamQuotaRequest); i {
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
		file_com_coralogixapis_aaa_organisations_v2_team_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTeamQuotaResponse); i {
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
			RawDescriptor: file_com_coralogixapis_aaa_organisations_v2_team_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_coralogixapis_aaa_organisations_v2_team_service_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_aaa_organisations_v2_team_service_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_aaa_organisations_v2_team_service_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_aaa_organisations_v2_team_service_proto = out.File
	file_com_coralogixapis_aaa_organisations_v2_team_service_proto_rawDesc = nil
	file_com_coralogixapis_aaa_organisations_v2_team_service_proto_goTypes = nil
	file_com_coralogixapis_aaa_organisations_v2_team_service_proto_depIdxs = nil
}