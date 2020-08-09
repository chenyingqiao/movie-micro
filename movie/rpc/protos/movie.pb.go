// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: movie.proto

package protos

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type MovieResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source             string   `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Title              string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Alias              string   `protobuf:"bytes,3,opt,name=alias,proto3" json:"alias,omitempty"`
	Director           string   `protobuf:"bytes,4,opt,name=director,proto3" json:"director,omitempty"`
	Actor              string   `protobuf:"bytes,5,opt,name=actor,proto3" json:"actor,omitempty"`
	Types              string   `protobuf:"bytes,6,opt,name=types,proto3" json:"types,omitempty"`
	Location           string   `protobuf:"bytes,7,opt,name=location,proto3" json:"location,omitempty"`
	Language           string   `protobuf:"bytes,8,opt,name=language,proto3" json:"language,omitempty"`
	ShowingTime        string   `protobuf:"bytes,9,opt,name=showingTime,proto3" json:"showingTime,omitempty"`
	Long               string   `protobuf:"bytes,10,opt,name=long,proto3" json:"long,omitempty"`
	UpdateTime         string   `protobuf:"bytes,11,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
	Introduce          string   `protobuf:"bytes,12,opt,name=introduce,proto3" json:"introduce,omitempty"`
	VideoM3U8Source    []string `protobuf:"bytes,13,rep,name=videoM3u8Source,proto3" json:"videoM3u8Source,omitempty"`
	VideoZuidallSource []string `protobuf:"bytes,14,rep,name=videoZuidallSource,proto3" json:"videoZuidallSource,omitempty"`
	VideoMp4Source     []string `protobuf:"bytes,15,rep,name=videoMp4Source,proto3" json:"videoMp4Source,omitempty"`
	ImageURL           string   `protobuf:"bytes,16,opt,name=imageURL,proto3" json:"imageURL,omitempty"`
	Hash               string   `protobuf:"bytes,17,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *MovieResponse) Reset() {
	*x = MovieResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieResponse) ProtoMessage() {}

func (x *MovieResponse) ProtoReflect() protoreflect.Message {
	mi := &file_movie_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieResponse.ProtoReflect.Descriptor instead.
func (*MovieResponse) Descriptor() ([]byte, []int) {
	return file_movie_proto_rawDescGZIP(), []int{0}
}

func (x *MovieResponse) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *MovieResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *MovieResponse) GetAlias() string {
	if x != nil {
		return x.Alias
	}
	return ""
}

func (x *MovieResponse) GetDirector() string {
	if x != nil {
		return x.Director
	}
	return ""
}

func (x *MovieResponse) GetActor() string {
	if x != nil {
		return x.Actor
	}
	return ""
}

func (x *MovieResponse) GetTypes() string {
	if x != nil {
		return x.Types
	}
	return ""
}

func (x *MovieResponse) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *MovieResponse) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *MovieResponse) GetShowingTime() string {
	if x != nil {
		return x.ShowingTime
	}
	return ""
}

func (x *MovieResponse) GetLong() string {
	if x != nil {
		return x.Long
	}
	return ""
}

func (x *MovieResponse) GetUpdateTime() string {
	if x != nil {
		return x.UpdateTime
	}
	return ""
}

func (x *MovieResponse) GetIntroduce() string {
	if x != nil {
		return x.Introduce
	}
	return ""
}

func (x *MovieResponse) GetVideoM3U8Source() []string {
	if x != nil {
		return x.VideoM3U8Source
	}
	return nil
}

func (x *MovieResponse) GetVideoZuidallSource() []string {
	if x != nil {
		return x.VideoZuidallSource
	}
	return nil
}

func (x *MovieResponse) GetVideoMp4Source() []string {
	if x != nil {
		return x.VideoMp4Source
	}
	return nil
}

func (x *MovieResponse) GetImageURL() string {
	if x != nil {
		return x.ImageURL
	}
	return ""
}

func (x *MovieResponse) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

// movie请求
type MovieRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash  string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	ObjId string `protobuf:"bytes,2,opt,name=obj_id,json=objId,proto3" json:"obj_id,omitempty"`
	Limit int64  `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *MovieRequest) Reset() {
	*x = MovieRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieRequest) ProtoMessage() {}

func (x *MovieRequest) ProtoReflect() protoreflect.Message {
	mi := &file_movie_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieRequest.ProtoReflect.Descriptor instead.
func (*MovieRequest) Descriptor() ([]byte, []int) {
	return file_movie_proto_rawDescGZIP(), []int{1}
}

func (x *MovieRequest) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *MovieRequest) GetObjId() string {
	if x != nil {
		return x.ObjId
	}
	return ""
}

func (x *MovieRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type MovieDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *MovieDeleteResponse) Reset() {
	*x = MovieDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieDeleteResponse) ProtoMessage() {}

func (x *MovieDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_movie_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieDeleteResponse.ProtoReflect.Descriptor instead.
func (*MovieDeleteResponse) Descriptor() ([]byte, []int) {
	return file_movie_proto_rawDescGZIP(), []int{2}
}

func (x *MovieDeleteResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_movie_proto protoreflect.FileDescriptor

var file_movie_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xf9, 0x03, 0x0a, 0x0d, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a,
	0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x68, 0x6f,
	0x77, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x73, 0x68, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c,
	0x6f, 0x6e, 0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x6f, 0x6e, 0x67, 0x12,
	0x1e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x12, 0x28, 0x0a,
	0x0f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x4d, 0x33, 0x75, 0x38, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x18, 0x0d, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x4d, 0x33, 0x75,
	0x38, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x5a, 0x75, 0x69, 0x64, 0x61, 0x6c, 0x6c, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x0e, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x12, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5a, 0x75, 0x69, 0x64, 0x61, 0x6c,
	0x6c, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x4d, 0x70, 0x34, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x4d, 0x70, 0x34, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x52, 0x4c, 0x18, 0x10, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x52, 0x4c, 0x12, 0x12, 0x0a, 0x04, 0x68,
	0x61, 0x73, 0x68, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x22,
	0x4f, 0x0a, 0x0c, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68,
	0x61, 0x73, 0x68, 0x12, 0x15, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x62, 0x6a, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x22, 0x2d, 0x0a, 0x13, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32,
	0xdc, 0x02, 0x0a, 0x05, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x4f, 0x0a, 0x06, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f,
	0x76, 0x69, 0x65, 0x2f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x4d, 0x0a, 0x04, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4d, 0x6f, 0x76, 0x69,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x76,
	0x69, 0x65, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x30, 0x01, 0x12, 0x58, 0x0a, 0x06, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x10,
	0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x3a, 0x01, 0x2a, 0x12, 0x59, 0x0a, 0x09, 0x52, 0x65, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72,
	0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65,
	0x2f, 0x72, 0x65, 0x5f, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x3a, 0x01, 0x2a, 0x42, 0x0a,
	0x5a, 0x08, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_movie_proto_rawDescOnce sync.Once
	file_movie_proto_rawDescData = file_movie_proto_rawDesc
)

func file_movie_proto_rawDescGZIP() []byte {
	file_movie_proto_rawDescOnce.Do(func() {
		file_movie_proto_rawDescData = protoimpl.X.CompressGZIP(file_movie_proto_rawDescData)
	})
	return file_movie_proto_rawDescData
}

var file_movie_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_movie_proto_goTypes = []interface{}{
	(*MovieResponse)(nil),       // 0: protos.MovieResponse
	(*MovieRequest)(nil),        // 1: protos.MovieRequest
	(*MovieDeleteResponse)(nil), // 2: protos.MovieDeleteResponse
}
var file_movie_proto_depIdxs = []int32{
	1, // 0: protos.Movie.Detail:input_type -> protos.MovieRequest
	1, // 1: protos.Movie.List:input_type -> protos.MovieRequest
	1, // 2: protos.Movie.Delete:input_type -> protos.MovieRequest
	1, // 3: protos.Movie.ReCrawler:input_type -> protos.MovieRequest
	0, // 4: protos.Movie.Detail:output_type -> protos.MovieResponse
	0, // 5: protos.Movie.List:output_type -> protos.MovieResponse
	2, // 6: protos.Movie.Delete:output_type -> protos.MovieDeleteResponse
	0, // 7: protos.Movie.ReCrawler:output_type -> protos.MovieResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_movie_proto_init() }
func file_movie_proto_init() {
	if File_movie_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_movie_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovieResponse); i {
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
		file_movie_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovieRequest); i {
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
		file_movie_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovieDeleteResponse); i {
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
			RawDescriptor: file_movie_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_movie_proto_goTypes,
		DependencyIndexes: file_movie_proto_depIdxs,
		MessageInfos:      file_movie_proto_msgTypes,
	}.Build()
	File_movie_proto = out.File
	file_movie_proto_rawDesc = nil
	file_movie_proto_goTypes = nil
	file_movie_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MovieClient is the client API for Movie service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MovieClient interface {
	Detail(ctx context.Context, in *MovieRequest, opts ...grpc.CallOption) (*MovieResponse, error)
	List(ctx context.Context, in *MovieRequest, opts ...grpc.CallOption) (Movie_ListClient, error)
	Delete(ctx context.Context, in *MovieRequest, opts ...grpc.CallOption) (*MovieDeleteResponse, error)
	ReCrawler(ctx context.Context, in *MovieRequest, opts ...grpc.CallOption) (*MovieResponse, error)
}

type movieClient struct {
	cc grpc.ClientConnInterface
}

func NewMovieClient(cc grpc.ClientConnInterface) MovieClient {
	return &movieClient{cc}
}

func (c *movieClient) Detail(ctx context.Context, in *MovieRequest, opts ...grpc.CallOption) (*MovieResponse, error) {
	out := new(MovieResponse)
	err := c.cc.Invoke(ctx, "/protos.Movie/Detail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieClient) List(ctx context.Context, in *MovieRequest, opts ...grpc.CallOption) (Movie_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Movie_serviceDesc.Streams[0], "/protos.Movie/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &movieListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Movie_ListClient interface {
	Recv() (*MovieResponse, error)
	grpc.ClientStream
}

type movieListClient struct {
	grpc.ClientStream
}

func (x *movieListClient) Recv() (*MovieResponse, error) {
	m := new(MovieResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *movieClient) Delete(ctx context.Context, in *MovieRequest, opts ...grpc.CallOption) (*MovieDeleteResponse, error) {
	out := new(MovieDeleteResponse)
	err := c.cc.Invoke(ctx, "/protos.Movie/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieClient) ReCrawler(ctx context.Context, in *MovieRequest, opts ...grpc.CallOption) (*MovieResponse, error) {
	out := new(MovieResponse)
	err := c.cc.Invoke(ctx, "/protos.Movie/ReCrawler", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MovieServer is the server API for Movie service.
type MovieServer interface {
	Detail(context.Context, *MovieRequest) (*MovieResponse, error)
	List(*MovieRequest, Movie_ListServer) error
	Delete(context.Context, *MovieRequest) (*MovieDeleteResponse, error)
	ReCrawler(context.Context, *MovieRequest) (*MovieResponse, error)
}

// UnimplementedMovieServer can be embedded to have forward compatible implementations.
type UnimplementedMovieServer struct {
}

func (*UnimplementedMovieServer) Detail(context.Context, *MovieRequest) (*MovieResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Detail not implemented")
}
func (*UnimplementedMovieServer) List(*MovieRequest, Movie_ListServer) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedMovieServer) Delete(context.Context, *MovieRequest) (*MovieDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedMovieServer) ReCrawler(context.Context, *MovieRequest) (*MovieResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReCrawler not implemented")
}

func RegisterMovieServer(s *grpc.Server, srv MovieServer) {
	s.RegisterService(&_Movie_serviceDesc, srv)
}

func _Movie_Detail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MovieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieServer).Detail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Movie/Detail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieServer).Detail(ctx, req.(*MovieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Movie_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MovieRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MovieServer).List(m, &movieListServer{stream})
}

type Movie_ListServer interface {
	Send(*MovieResponse) error
	grpc.ServerStream
}

type movieListServer struct {
	grpc.ServerStream
}

func (x *movieListServer) Send(m *MovieResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Movie_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MovieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Movie/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieServer).Delete(ctx, req.(*MovieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Movie_ReCrawler_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MovieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieServer).ReCrawler(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Movie/ReCrawler",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieServer).ReCrawler(ctx, req.(*MovieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Movie_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Movie",
	HandlerType: (*MovieServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Detail",
			Handler:    _Movie_Detail_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Movie_Delete_Handler,
		},
		{
			MethodName: "ReCrawler",
			Handler:    _Movie_ReCrawler_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _Movie_List_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "movie.proto",
}