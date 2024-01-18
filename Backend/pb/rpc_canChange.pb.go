// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: rpc_canChange.proto

package pb

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

// Create
type CreateCanChangeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionID     string   `protobuf:"bytes,1,opt,name=SessionID,proto3" json:"SessionID,omitempty"`
	NickName      string   `protobuf:"bytes,2,opt,name=NickName,proto3" json:"NickName,omitempty"`
	City          string   `protobuf:"bytes,3,opt,name=City,proto3" json:"City,omitempty"`
	Sexual        string   `protobuf:"bytes,4,opt,name=Sexual,proto3" json:"Sexual,omitempty"`
	Height        int32    `protobuf:"varint,5,opt,name=Height,proto3" json:"Height,omitempty"`
	Weight        int32    `protobuf:"varint,6,opt,name=Weight,proto3" json:"Weight,omitempty"`
	Speaklanguage []string `protobuf:"bytes,7,rep,name=Speaklanguage,proto3" json:"Speaklanguage,omitempty"`
	Education     string   `protobuf:"bytes,8,opt,name=Education,proto3" json:"Education,omitempty"`
	Job           string   `protobuf:"bytes,9,opt,name=Job,proto3" json:"Job,omitempty"`
	AnnualSalary  int32    `protobuf:"varint,10,opt,name=AnnualSalary,proto3" json:"AnnualSalary,omitempty"`
	Sociability   string   `protobuf:"bytes,11,opt,name=Sociability,proto3" json:"Sociability,omitempty"`
	HobbyType     string   `protobuf:"bytes,12,opt,name=hobbyType,proto3" json:"hobbyType,omitempty"`
	Experience    int32    `protobuf:"varint,13,opt,name=Experience,proto3" json:"Experience,omitempty"`
	AccompanyType string   `protobuf:"bytes,14,opt,name=AccompanyType,proto3" json:"AccompanyType,omitempty"`
	Religious     string   `protobuf:"bytes,15,opt,name=Religious,proto3" json:"Religious,omitempty"`
	Introduce     string   `protobuf:"bytes,16,opt,name=Introduce,proto3" json:"Introduce,omitempty"`
}

func (x *CreateCanChangeRequest) Reset() {
	*x = CreateCanChangeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_canChange_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCanChangeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCanChangeRequest) ProtoMessage() {}

func (x *CreateCanChangeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_canChange_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCanChangeRequest.ProtoReflect.Descriptor instead.
func (*CreateCanChangeRequest) Descriptor() ([]byte, []int) {
	return file_rpc_canChange_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCanChangeRequest) GetSessionID() string {
	if x != nil {
		return x.SessionID
	}
	return ""
}

func (x *CreateCanChangeRequest) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *CreateCanChangeRequest) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *CreateCanChangeRequest) GetSexual() string {
	if x != nil {
		return x.Sexual
	}
	return ""
}

func (x *CreateCanChangeRequest) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *CreateCanChangeRequest) GetWeight() int32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *CreateCanChangeRequest) GetSpeaklanguage() []string {
	if x != nil {
		return x.Speaklanguage
	}
	return nil
}

func (x *CreateCanChangeRequest) GetEducation() string {
	if x != nil {
		return x.Education
	}
	return ""
}

func (x *CreateCanChangeRequest) GetJob() string {
	if x != nil {
		return x.Job
	}
	return ""
}

func (x *CreateCanChangeRequest) GetAnnualSalary() int32 {
	if x != nil {
		return x.AnnualSalary
	}
	return 0
}

func (x *CreateCanChangeRequest) GetSociability() string {
	if x != nil {
		return x.Sociability
	}
	return ""
}

func (x *CreateCanChangeRequest) GetHobbyType() string {
	if x != nil {
		return x.HobbyType
	}
	return ""
}

func (x *CreateCanChangeRequest) GetExperience() int32 {
	if x != nil {
		return x.Experience
	}
	return 0
}

func (x *CreateCanChangeRequest) GetAccompanyType() string {
	if x != nil {
		return x.AccompanyType
	}
	return ""
}

func (x *CreateCanChangeRequest) GetReligious() string {
	if x != nil {
		return x.Religious
	}
	return ""
}

func (x *CreateCanChangeRequest) GetIntroduce() string {
	if x != nil {
		return x.Introduce
	}
	return ""
}

type CreateCanChangeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CanChangeInfo *CanChange `protobuf:"bytes,1,opt,name=CanChangeInfo,proto3" json:"CanChangeInfo,omitempty"`
}

func (x *CreateCanChangeResponse) Reset() {
	*x = CreateCanChangeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_canChange_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCanChangeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCanChangeResponse) ProtoMessage() {}

func (x *CreateCanChangeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_canChange_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCanChangeResponse.ProtoReflect.Descriptor instead.
func (*CreateCanChangeResponse) Descriptor() ([]byte, []int) {
	return file_rpc_canChange_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCanChangeResponse) GetCanChangeInfo() *CanChange {
	if x != nil {
		return x.CanChangeInfo
	}
	return nil
}

// Get
type GetCanChangeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionID string `protobuf:"bytes,1,opt,name=SessionID,proto3" json:"SessionID,omitempty"`
	UserID    int32  `protobuf:"varint,2,opt,name=UserID,proto3" json:"UserID,omitempty"`
}

func (x *GetCanChangeRequest) Reset() {
	*x = GetCanChangeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_canChange_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCanChangeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCanChangeRequest) ProtoMessage() {}

func (x *GetCanChangeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_canChange_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCanChangeRequest.ProtoReflect.Descriptor instead.
func (*GetCanChangeRequest) Descriptor() ([]byte, []int) {
	return file_rpc_canChange_proto_rawDescGZIP(), []int{2}
}

func (x *GetCanChangeRequest) GetSessionID() string {
	if x != nil {
		return x.SessionID
	}
	return ""
}

func (x *GetCanChangeRequest) GetUserID() int32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

type GetCanChangeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CanChangeInfo *CanChange `protobuf:"bytes,1,opt,name=CanChangeInfo,proto3" json:"CanChangeInfo,omitempty"`
}

func (x *GetCanChangeResponse) Reset() {
	*x = GetCanChangeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_canChange_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCanChangeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCanChangeResponse) ProtoMessage() {}

func (x *GetCanChangeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_canChange_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCanChangeResponse.ProtoReflect.Descriptor instead.
func (*GetCanChangeResponse) Descriptor() ([]byte, []int) {
	return file_rpc_canChange_proto_rawDescGZIP(), []int{3}
}

func (x *GetCanChangeResponse) GetCanChangeInfo() *CanChange {
	if x != nil {
		return x.CanChangeInfo
	}
	return nil
}

// Update
type UpdateCanChangeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionID     string   `protobuf:"bytes,1,opt,name=SessionID,proto3" json:"SessionID,omitempty"`
	NickName      string   `protobuf:"bytes,2,opt,name=NickName,proto3" json:"NickName,omitempty"`
	City          string   `protobuf:"bytes,3,opt,name=City,proto3" json:"City,omitempty"`
	Sexual        string   `protobuf:"bytes,4,opt,name=Sexual,proto3" json:"Sexual,omitempty"`
	Height        int32    `protobuf:"varint,5,opt,name=Height,proto3" json:"Height,omitempty"`
	Weight        int32    `protobuf:"varint,6,opt,name=Weight,proto3" json:"Weight,omitempty"`
	Speaklanguage []string `protobuf:"bytes,7,rep,name=Speaklanguage,proto3" json:"Speaklanguage,omitempty"`
	Education     string   `protobuf:"bytes,8,opt,name=Education,proto3" json:"Education,omitempty"`
	Job           string   `protobuf:"bytes,9,opt,name=Job,proto3" json:"Job,omitempty"`
	AnnualSalary  int32    `protobuf:"varint,10,opt,name=AnnualSalary,proto3" json:"AnnualSalary,omitempty"`
	Sociability   string   `protobuf:"bytes,11,opt,name=Sociability,proto3" json:"Sociability,omitempty"`
	HobbyType     string   `protobuf:"bytes,12,opt,name=hobbyType,proto3" json:"hobbyType,omitempty"`
	Experience    int32    `protobuf:"varint,13,opt,name=Experience,proto3" json:"Experience,omitempty"`
	AccompanyType string   `protobuf:"bytes,14,opt,name=AccompanyType,proto3" json:"AccompanyType,omitempty"`
	Religious     string   `protobuf:"bytes,15,opt,name=Religious,proto3" json:"Religious,omitempty"`
	Introduce     string   `protobuf:"bytes,16,opt,name=Introduce,proto3" json:"Introduce,omitempty"`
}

func (x *UpdateCanChangeRequest) Reset() {
	*x = UpdateCanChangeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_canChange_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCanChangeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCanChangeRequest) ProtoMessage() {}

func (x *UpdateCanChangeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_canChange_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCanChangeRequest.ProtoReflect.Descriptor instead.
func (*UpdateCanChangeRequest) Descriptor() ([]byte, []int) {
	return file_rpc_canChange_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateCanChangeRequest) GetSessionID() string {
	if x != nil {
		return x.SessionID
	}
	return ""
}

func (x *UpdateCanChangeRequest) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *UpdateCanChangeRequest) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *UpdateCanChangeRequest) GetSexual() string {
	if x != nil {
		return x.Sexual
	}
	return ""
}

func (x *UpdateCanChangeRequest) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *UpdateCanChangeRequest) GetWeight() int32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *UpdateCanChangeRequest) GetSpeaklanguage() []string {
	if x != nil {
		return x.Speaklanguage
	}
	return nil
}

func (x *UpdateCanChangeRequest) GetEducation() string {
	if x != nil {
		return x.Education
	}
	return ""
}

func (x *UpdateCanChangeRequest) GetJob() string {
	if x != nil {
		return x.Job
	}
	return ""
}

func (x *UpdateCanChangeRequest) GetAnnualSalary() int32 {
	if x != nil {
		return x.AnnualSalary
	}
	return 0
}

func (x *UpdateCanChangeRequest) GetSociability() string {
	if x != nil {
		return x.Sociability
	}
	return ""
}

func (x *UpdateCanChangeRequest) GetHobbyType() string {
	if x != nil {
		return x.HobbyType
	}
	return ""
}

func (x *UpdateCanChangeRequest) GetExperience() int32 {
	if x != nil {
		return x.Experience
	}
	return 0
}

func (x *UpdateCanChangeRequest) GetAccompanyType() string {
	if x != nil {
		return x.AccompanyType
	}
	return ""
}

func (x *UpdateCanChangeRequest) GetReligious() string {
	if x != nil {
		return x.Religious
	}
	return ""
}

func (x *UpdateCanChangeRequest) GetIntroduce() string {
	if x != nil {
		return x.Introduce
	}
	return ""
}

type UpdateCanChangeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CanChangeInfo *CanChange `protobuf:"bytes,1,opt,name=CanChangeInfo,proto3" json:"CanChangeInfo,omitempty"`
}

func (x *UpdateCanChangeResponse) Reset() {
	*x = UpdateCanChangeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_canChange_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCanChangeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCanChangeResponse) ProtoMessage() {}

func (x *UpdateCanChangeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_canChange_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCanChangeResponse.ProtoReflect.Descriptor instead.
func (*UpdateCanChangeResponse) Descriptor() ([]byte, []int) {
	return file_rpc_canChange_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateCanChangeResponse) GetCanChangeInfo() *CanChange {
	if x != nil {
		return x.CanChangeInfo
	}
	return nil
}

// Delete
type DeleteCanChangeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionID string `protobuf:"bytes,1,opt,name=SessionID,proto3" json:"SessionID,omitempty"`
}

func (x *DeleteCanChangeRequest) Reset() {
	*x = DeleteCanChangeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_canChange_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCanChangeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCanChangeRequest) ProtoMessage() {}

func (x *DeleteCanChangeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_canChange_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCanChangeRequest.ProtoReflect.Descriptor instead.
func (*DeleteCanChangeRequest) Descriptor() ([]byte, []int) {
	return file_rpc_canChange_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteCanChangeRequest) GetSessionID() string {
	if x != nil {
		return x.SessionID
	}
	return ""
}

var File_rpc_canChange_proto protoreflect.FileDescriptor

var file_rpc_canChange_proto_rawDesc = []byte{
	0x0a, 0x13, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x61, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x0f, 0x63, 0x61, 0x6e, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xea, 0x03, 0x0a, 0x16, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x4e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x43, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43,
	0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x65, 0x78, 0x75, 0x61, 0x6c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x65, 0x78, 0x75, 0x61, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x48,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x48, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x53,
	0x70, 0x65, 0x61, 0x6b, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0d, 0x53, 0x70, 0x65, 0x61, 0x6b, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x45, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x45, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x10, 0x0a, 0x03, 0x4a, 0x6f, 0x62, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4a, 0x6f,
	0x62, 0x12, 0x22, 0x0a, 0x0c, 0x41, 0x6e, 0x6e, 0x75, 0x61, 0x6c, 0x53, 0x61, 0x6c, 0x61, 0x72,
	0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x41, 0x6e, 0x6e, 0x75, 0x61, 0x6c, 0x53,
	0x61, 0x6c, 0x61, 0x72, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x53, 0x6f, 0x63, 0x69,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x68, 0x6f, 0x62, 0x62, 0x79,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x68, 0x6f, 0x62, 0x62,
	0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65,
	0x6e, 0x63, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x45, 0x78, 0x70, 0x65, 0x72,
	0x69, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x41, 0x63,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x52,
	0x65, 0x6c, 0x69, 0x67, 0x69, 0x6f, 0x75, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x52, 0x65, 0x6c, 0x69, 0x67, 0x69, 0x6f, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x6e, 0x74,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x49, 0x6e,
	0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x22, 0x4e, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x61, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x33, 0x0a, 0x0d, 0x43, 0x61, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x43,
	0x61, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x0d, 0x43, 0x61, 0x6e, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x4b, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x61,
	0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x22, 0x4b, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x61, 0x6e, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x0d,
	0x43, 0x61, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x6e, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x52, 0x0d, 0x43, 0x61, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x22, 0xea, 0x03, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x6e, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x4e, 0x69,
	0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4e, 0x69,
	0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x69, 0x74, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x65,
	0x78, 0x75, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x65, 0x78, 0x75,
	0x61, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x57, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x57, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x53, 0x70, 0x65, 0x61, 0x6b, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x53, 0x70, 0x65, 0x61, 0x6b,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x45, 0x64, 0x75, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x45, 0x64, 0x75,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x4a, 0x6f, 0x62, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x4a, 0x6f, 0x62, 0x12, 0x22, 0x0a, 0x0c, 0x41, 0x6e, 0x6e, 0x75,
	0x61, 0x6c, 0x53, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c,
	0x41, 0x6e, 0x6e, 0x75, 0x61, 0x6c, 0x53, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x12, 0x20, 0x0a, 0x0b,
	0x53, 0x6f, 0x63, 0x69, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x1c,
	0x0a, 0x09, 0x68, 0x6f, 0x62, 0x62, 0x79, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x68, 0x6f, 0x62, 0x62, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x24, 0x0a, 0x0d,
	0x41, 0x63, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x52, 0x65, 0x6c, 0x69, 0x67, 0x69, 0x6f, 0x75, 0x73, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x52, 0x65, 0x6c, 0x69, 0x67, 0x69, 0x6f, 0x75, 0x73,
	0x12, 0x1c, 0x0a, 0x09, 0x49, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x18, 0x10, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x49, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x22, 0x4e,
	0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x0d, 0x43, 0x61, 0x6e,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52,
	0x0d, 0x43, 0x61, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x36,
	0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x42, 0x0c, 0x5a, 0x0a, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_canChange_proto_rawDescOnce sync.Once
	file_rpc_canChange_proto_rawDescData = file_rpc_canChange_proto_rawDesc
)

func file_rpc_canChange_proto_rawDescGZIP() []byte {
	file_rpc_canChange_proto_rawDescOnce.Do(func() {
		file_rpc_canChange_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_canChange_proto_rawDescData)
	})
	return file_rpc_canChange_proto_rawDescData
}

var file_rpc_canChange_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_rpc_canChange_proto_goTypes = []interface{}{
	(*CreateCanChangeRequest)(nil),  // 0: pb.CreateCanChangeRequest
	(*CreateCanChangeResponse)(nil), // 1: pb.CreateCanChangeResponse
	(*GetCanChangeRequest)(nil),     // 2: pb.GetCanChangeRequest
	(*GetCanChangeResponse)(nil),    // 3: pb.GetCanChangeResponse
	(*UpdateCanChangeRequest)(nil),  // 4: pb.UpdateCanChangeRequest
	(*UpdateCanChangeResponse)(nil), // 5: pb.UpdateCanChangeResponse
	(*DeleteCanChangeRequest)(nil),  // 6: pb.DeleteCanChangeRequest
	(*CanChange)(nil),               // 7: pb.CanChange
}
var file_rpc_canChange_proto_depIdxs = []int32{
	7, // 0: pb.CreateCanChangeResponse.CanChangeInfo:type_name -> pb.CanChange
	7, // 1: pb.GetCanChangeResponse.CanChangeInfo:type_name -> pb.CanChange
	7, // 2: pb.UpdateCanChangeResponse.CanChangeInfo:type_name -> pb.CanChange
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_rpc_canChange_proto_init() }
func file_rpc_canChange_proto_init() {
	if File_rpc_canChange_proto != nil {
		return
	}
	file_canChange_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rpc_canChange_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCanChangeRequest); i {
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
		file_rpc_canChange_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCanChangeResponse); i {
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
		file_rpc_canChange_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCanChangeRequest); i {
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
		file_rpc_canChange_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCanChangeResponse); i {
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
		file_rpc_canChange_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCanChangeRequest); i {
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
		file_rpc_canChange_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCanChangeResponse); i {
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
		file_rpc_canChange_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCanChangeRequest); i {
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
			RawDescriptor: file_rpc_canChange_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_canChange_proto_goTypes,
		DependencyIndexes: file_rpc_canChange_proto_depIdxs,
		MessageInfos:      file_rpc_canChange_proto_msgTypes,
	}.Build()
	File_rpc_canChange_proto = out.File
	file_rpc_canChange_proto_rawDesc = nil
	file_rpc_canChange_proto_goTypes = nil
	file_rpc_canChange_proto_depIdxs = nil
}
