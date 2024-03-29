// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: rpc_images.proto

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
type CreateImagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionID string `protobuf:"bytes,1,opt,name=SessionID,proto3" json:"SessionID,omitempty"`
	Img1      []byte `protobuf:"bytes,2,opt,name=Img1,proto3" json:"Img1,omitempty"`
	Img2      []byte `protobuf:"bytes,3,opt,name=Img2,proto3" json:"Img2,omitempty"`
	Img3      []byte `protobuf:"bytes,4,opt,name=Img3,proto3" json:"Img3,omitempty"`
	Img4      []byte `protobuf:"bytes,5,opt,name=Img4,proto3" json:"Img4,omitempty"`
	Img5      []byte `protobuf:"bytes,6,opt,name=img5,proto3" json:"img5,omitempty"`
}

func (x *CreateImagesRequest) Reset() {
	*x = CreateImagesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_images_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateImagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateImagesRequest) ProtoMessage() {}

func (x *CreateImagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_images_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateImagesRequest.ProtoReflect.Descriptor instead.
func (*CreateImagesRequest) Descriptor() ([]byte, []int) {
	return file_rpc_images_proto_rawDescGZIP(), []int{0}
}

func (x *CreateImagesRequest) GetSessionID() string {
	if x != nil {
		return x.SessionID
	}
	return ""
}

func (x *CreateImagesRequest) GetImg1() []byte {
	if x != nil {
		return x.Img1
	}
	return nil
}

func (x *CreateImagesRequest) GetImg2() []byte {
	if x != nil {
		return x.Img2
	}
	return nil
}

func (x *CreateImagesRequest) GetImg3() []byte {
	if x != nil {
		return x.Img3
	}
	return nil
}

func (x *CreateImagesRequest) GetImg4() []byte {
	if x != nil {
		return x.Img4
	}
	return nil
}

func (x *CreateImagesRequest) GetImg5() []byte {
	if x != nil {
		return x.Img5
	}
	return nil
}

type CreateImagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Img *Images `protobuf:"bytes,1,opt,name=Img,proto3" json:"Img,omitempty"`
}

func (x *CreateImagesResponse) Reset() {
	*x = CreateImagesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_images_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateImagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateImagesResponse) ProtoMessage() {}

func (x *CreateImagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_images_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateImagesResponse.ProtoReflect.Descriptor instead.
func (*CreateImagesResponse) Descriptor() ([]byte, []int) {
	return file_rpc_images_proto_rawDescGZIP(), []int{1}
}

func (x *CreateImagesResponse) GetImg() *Images {
	if x != nil {
		return x.Img
	}
	return nil
}

// Get
type GetImagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionID string `protobuf:"bytes,1,opt,name=SessionID,proto3" json:"SessionID,omitempty"`
	UserID    int32  `protobuf:"varint,2,opt,name=UserID,proto3" json:"UserID,omitempty"`
}

func (x *GetImagesRequest) Reset() {
	*x = GetImagesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_images_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetImagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetImagesRequest) ProtoMessage() {}

func (x *GetImagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_images_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetImagesRequest.ProtoReflect.Descriptor instead.
func (*GetImagesRequest) Descriptor() ([]byte, []int) {
	return file_rpc_images_proto_rawDescGZIP(), []int{2}
}

func (x *GetImagesRequest) GetSessionID() string {
	if x != nil {
		return x.SessionID
	}
	return ""
}

func (x *GetImagesRequest) GetUserID() int32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

type GetImagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Img *Images `protobuf:"bytes,1,opt,name=Img,proto3" json:"Img,omitempty"`
}

func (x *GetImagesResponse) Reset() {
	*x = GetImagesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_images_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetImagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetImagesResponse) ProtoMessage() {}

func (x *GetImagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_images_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetImagesResponse.ProtoReflect.Descriptor instead.
func (*GetImagesResponse) Descriptor() ([]byte, []int) {
	return file_rpc_images_proto_rawDescGZIP(), []int{3}
}

func (x *GetImagesResponse) GetImg() *Images {
	if x != nil {
		return x.Img
	}
	return nil
}

// Updat
type UpdateImagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionID string `protobuf:"bytes,1,opt,name=SessionID,proto3" json:"SessionID,omitempty"`
	Img1      []byte `protobuf:"bytes,2,opt,name=Img1,proto3" json:"Img1,omitempty"`
	Img2      []byte `protobuf:"bytes,3,opt,name=Img2,proto3" json:"Img2,omitempty"`
	Img3      []byte `protobuf:"bytes,4,opt,name=Img3,proto3" json:"Img3,omitempty"`
	Img4      []byte `protobuf:"bytes,5,opt,name=Img4,proto3" json:"Img4,omitempty"`
	Img5      []byte `protobuf:"bytes,6,opt,name=img5,proto3" json:"img5,omitempty"`
}

func (x *UpdateImagesRequest) Reset() {
	*x = UpdateImagesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_images_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateImagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateImagesRequest) ProtoMessage() {}

func (x *UpdateImagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_images_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateImagesRequest.ProtoReflect.Descriptor instead.
func (*UpdateImagesRequest) Descriptor() ([]byte, []int) {
	return file_rpc_images_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateImagesRequest) GetSessionID() string {
	if x != nil {
		return x.SessionID
	}
	return ""
}

func (x *UpdateImagesRequest) GetImg1() []byte {
	if x != nil {
		return x.Img1
	}
	return nil
}

func (x *UpdateImagesRequest) GetImg2() []byte {
	if x != nil {
		return x.Img2
	}
	return nil
}

func (x *UpdateImagesRequest) GetImg3() []byte {
	if x != nil {
		return x.Img3
	}
	return nil
}

func (x *UpdateImagesRequest) GetImg4() []byte {
	if x != nil {
		return x.Img4
	}
	return nil
}

func (x *UpdateImagesRequest) GetImg5() []byte {
	if x != nil {
		return x.Img5
	}
	return nil
}

type UpdateImagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Img *Images `protobuf:"bytes,1,opt,name=Img,proto3" json:"Img,omitempty"`
}

func (x *UpdateImagesResponse) Reset() {
	*x = UpdateImagesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_images_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateImagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateImagesResponse) ProtoMessage() {}

func (x *UpdateImagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_images_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateImagesResponse.ProtoReflect.Descriptor instead.
func (*UpdateImagesResponse) Descriptor() ([]byte, []int) {
	return file_rpc_images_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateImagesResponse) GetImg() *Images {
	if x != nil {
		return x.Img
	}
	return nil
}

// Delete
type DeleteImagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionID string `protobuf:"bytes,1,opt,name=SessionID,proto3" json:"SessionID,omitempty"`
}

func (x *DeleteImagesRequest) Reset() {
	*x = DeleteImagesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_images_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteImagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteImagesRequest) ProtoMessage() {}

func (x *DeleteImagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_images_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteImagesRequest.ProtoReflect.Descriptor instead.
func (*DeleteImagesRequest) Descriptor() ([]byte, []int) {
	return file_rpc_images_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteImagesRequest) GetSessionID() string {
	if x != nil {
		return x.SessionID
	}
	return ""
}

type DeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_images_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_images_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_rpc_images_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_rpc_images_proto protoreflect.FileDescriptor

var file_rpc_images_proto_rawDesc = []byte{
	0x0a, 0x10, 0x72, 0x70, 0x63, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x0c, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x01, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x49, 0x6d,
	0x67, 0x31, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x49, 0x6d, 0x67, 0x31, 0x12, 0x12,
	0x0a, 0x04, 0x49, 0x6d, 0x67, 0x32, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x49, 0x6d,
	0x67, 0x32, 0x12, 0x12, 0x0a, 0x04, 0x49, 0x6d, 0x67, 0x33, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x49, 0x6d, 0x67, 0x33, 0x12, 0x12, 0x0a, 0x04, 0x49, 0x6d, 0x67, 0x34, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x49, 0x6d, 0x67, 0x34, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6d,
	0x67, 0x35, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x69, 0x6d, 0x67, 0x35, 0x22, 0x34,
	0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x03, 0x49, 0x6d, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x52,
	0x03, 0x49, 0x6d, 0x67, 0x22, 0x48, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x31,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x03, 0x49, 0x6d, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x52, 0x03, 0x49, 0x6d,
	0x67, 0x22, 0x97, 0x01, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x49, 0x6d, 0x67, 0x31, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x49, 0x6d, 0x67, 0x31, 0x12, 0x12, 0x0a, 0x04, 0x49,
	0x6d, 0x67, 0x32, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x49, 0x6d, 0x67, 0x32, 0x12,
	0x12, 0x0a, 0x04, 0x49, 0x6d, 0x67, 0x33, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x49,
	0x6d, 0x67, 0x33, 0x12, 0x12, 0x0a, 0x04, 0x49, 0x6d, 0x67, 0x34, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x49, 0x6d, 0x67, 0x34, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6d, 0x67, 0x35, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x69, 0x6d, 0x67, 0x35, 0x22, 0x34, 0x0a, 0x14, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x03, 0x49, 0x6d, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x52, 0x03, 0x49, 0x6d,
	0x67, 0x22, 0x33, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x22, 0x22, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x42, 0x0c, 0x5a, 0x0a, 0x42, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_images_proto_rawDescOnce sync.Once
	file_rpc_images_proto_rawDescData = file_rpc_images_proto_rawDesc
)

func file_rpc_images_proto_rawDescGZIP() []byte {
	file_rpc_images_proto_rawDescOnce.Do(func() {
		file_rpc_images_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_images_proto_rawDescData)
	})
	return file_rpc_images_proto_rawDescData
}

var file_rpc_images_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_rpc_images_proto_goTypes = []interface{}{
	(*CreateImagesRequest)(nil),  // 0: pb.CreateImagesRequest
	(*CreateImagesResponse)(nil), // 1: pb.CreateImagesResponse
	(*GetImagesRequest)(nil),     // 2: pb.GetImagesRequest
	(*GetImagesResponse)(nil),    // 3: pb.GetImagesResponse
	(*UpdateImagesRequest)(nil),  // 4: pb.UpdateImagesRequest
	(*UpdateImagesResponse)(nil), // 5: pb.UpdateImagesResponse
	(*DeleteImagesRequest)(nil),  // 6: pb.DeleteImagesRequest
	(*DeleteResponse)(nil),       // 7: pb.DeleteResponse
	(*Images)(nil),               // 8: pb.Images
}
var file_rpc_images_proto_depIdxs = []int32{
	8, // 0: pb.CreateImagesResponse.Img:type_name -> pb.Images
	8, // 1: pb.GetImagesResponse.Img:type_name -> pb.Images
	8, // 2: pb.UpdateImagesResponse.Img:type_name -> pb.Images
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_rpc_images_proto_init() }
func file_rpc_images_proto_init() {
	if File_rpc_images_proto != nil {
		return
	}
	file_images_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rpc_images_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateImagesRequest); i {
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
		file_rpc_images_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateImagesResponse); i {
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
		file_rpc_images_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetImagesRequest); i {
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
		file_rpc_images_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetImagesResponse); i {
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
		file_rpc_images_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateImagesRequest); i {
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
		file_rpc_images_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateImagesResponse); i {
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
		file_rpc_images_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteImagesRequest); i {
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
		file_rpc_images_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResponse); i {
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
			RawDescriptor: file_rpc_images_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_images_proto_goTypes,
		DependencyIndexes: file_rpc_images_proto_depIdxs,
		MessageInfos:      file_rpc_images_proto_msgTypes,
	}.Build()
	File_rpc_images_proto = out.File
	file_rpc_images_proto_rawDesc = nil
	file_rpc_images_proto_goTypes = nil
	file_rpc_images_proto_depIdxs = nil
}
