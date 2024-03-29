// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: hobby.proto

package pb

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

type Hobby struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID        int32                  `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Era           int32                  `protobuf:"varint,2,opt,name=Era,proto3" json:"Era,omitempty"`
	City          []string               `protobuf:"bytes,3,rep,name=City,proto3" json:"City,omitempty"`
	Gender        string                 `protobuf:"bytes,4,opt,name=Gender,proto3" json:"Gender,omitempty"`
	Speaklanguage []string               `protobuf:"bytes,5,rep,name=Speaklanguage,proto3" json:"Speaklanguage,omitempty"`
	FindType      string                 `protobuf:"bytes,6,opt,name=FindType,proto3" json:"FindType,omitempty"`
	Experience    int32                  `protobuf:"varint,7,opt,name=Experience,proto3" json:"Experience,omitempty"`
	InfoChangedAt *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=InfoChangedAt,proto3" json:"InfoChangedAt,omitempty"`
}

func (x *Hobby) Reset() {
	*x = Hobby{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hobby_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hobby) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hobby) ProtoMessage() {}

func (x *Hobby) ProtoReflect() protoreflect.Message {
	mi := &file_hobby_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hobby.ProtoReflect.Descriptor instead.
func (*Hobby) Descriptor() ([]byte, []int) {
	return file_hobby_proto_rawDescGZIP(), []int{0}
}

func (x *Hobby) GetUserID() int32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *Hobby) GetEra() int32 {
	if x != nil {
		return x.Era
	}
	return 0
}

func (x *Hobby) GetCity() []string {
	if x != nil {
		return x.City
	}
	return nil
}

func (x *Hobby) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *Hobby) GetSpeaklanguage() []string {
	if x != nil {
		return x.Speaklanguage
	}
	return nil
}

func (x *Hobby) GetFindType() string {
	if x != nil {
		return x.FindType
	}
	return ""
}

func (x *Hobby) GetExperience() int32 {
	if x != nil {
		return x.Experience
	}
	return 0
}

func (x *Hobby) GetInfoChangedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.InfoChangedAt
	}
	return nil
}

var File_hobby_proto protoreflect.FileDescriptor

var file_hobby_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x68, 0x6f, 0x62, 0x62, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x81, 0x02, 0x0a, 0x05, 0x48, 0x6f, 0x62, 0x62, 0x79, 0x12, 0x16, 0x0a, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x03, 0x45, 0x72, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x03, 0x45, 0x72, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x69, 0x74, 0x79, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x43, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x47, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x47, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x12, 0x24, 0x0a, 0x0d, 0x53, 0x70, 0x65, 0x61, 0x6b, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x53, 0x70, 0x65, 0x61, 0x6b,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x6e, 0x64,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x69, 0x6e, 0x64,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e,
	0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69,
	0x65, 0x6e, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x0d, 0x49, 0x6e, 0x66, 0x6f, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x64, 0x41, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x49, 0x6e, 0x66, 0x6f, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x64, 0x41, 0x74, 0x42, 0x0c, 0x5a, 0x0a, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hobby_proto_rawDescOnce sync.Once
	file_hobby_proto_rawDescData = file_hobby_proto_rawDesc
)

func file_hobby_proto_rawDescGZIP() []byte {
	file_hobby_proto_rawDescOnce.Do(func() {
		file_hobby_proto_rawDescData = protoimpl.X.CompressGZIP(file_hobby_proto_rawDescData)
	})
	return file_hobby_proto_rawDescData
}

var file_hobby_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_hobby_proto_goTypes = []interface{}{
	(*Hobby)(nil),                 // 0: pb.Hobby
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_hobby_proto_depIdxs = []int32{
	1, // 0: pb.Hobby.InfoChangedAt:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_hobby_proto_init() }
func file_hobby_proto_init() {
	if File_hobby_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hobby_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hobby); i {
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
			RawDescriptor: file_hobby_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_hobby_proto_goTypes,
		DependencyIndexes: file_hobby_proto_depIdxs,
		MessageInfos:      file_hobby_proto_msgTypes,
	}.Build()
	File_hobby_proto = out.File
	file_hobby_proto_rawDesc = nil
	file_hobby_proto_goTypes = nil
	file_hobby_proto_depIdxs = nil
}
