// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: api/ads.proto

package api

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

type Ad struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdID         string `protobuf:"bytes,1,opt,name=AdID,proto3" json:"AdID,omitempty"`
	AdvertiserID string `protobuf:"bytes,2,opt,name=AdvertiserID,proto3" json:"AdvertiserID,omitempty"`
	Title        string `protobuf:"bytes,3,opt,name=Title,proto3" json:"Title,omitempty"`
	AdURL        string `protobuf:"bytes,4,opt,name=AdURL,proto3" json:"AdURL,omitempty"`
	ImpressionId string `protobuf:"bytes,5,opt,name=ImpressionId,proto3" json:"ImpressionId,omitempty"`
}

func (x *Ad) Reset() {
	*x = Ad{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ads_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ad) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ad) ProtoMessage() {}

func (x *Ad) ProtoReflect() protoreflect.Message {
	mi := &file_api_ads_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ad.ProtoReflect.Descriptor instead.
func (*Ad) Descriptor() ([]byte, []int) {
	return file_api_ads_proto_rawDescGZIP(), []int{0}
}

func (x *Ad) GetAdID() string {
	if x != nil {
		return x.AdID
	}
	return ""
}

func (x *Ad) GetAdvertiserID() string {
	if x != nil {
		return x.AdvertiserID
	}
	return ""
}

func (x *Ad) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Ad) GetAdURL() string {
	if x != nil {
		return x.AdURL
	}
	return ""
}

func (x *Ad) GetImpressionId() string {
	if x != nil {
		return x.ImpressionId
	}
	return ""
}

type CreateAdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdvertiserID string `protobuf:"bytes,1,opt,name=AdvertiserID,proto3" json:"AdvertiserID,omitempty"`
	Title        string `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	AdURL        string `protobuf:"bytes,3,opt,name=AdURL,proto3" json:"AdURL,omitempty"`
}

func (x *CreateAdRequest) Reset() {
	*x = CreateAdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ads_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAdRequest) ProtoMessage() {}

func (x *CreateAdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_ads_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAdRequest.ProtoReflect.Descriptor instead.
func (*CreateAdRequest) Descriptor() ([]byte, []int) {
	return file_api_ads_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAdRequest) GetAdvertiserID() string {
	if x != nil {
		return x.AdvertiserID
	}
	return ""
}

func (x *CreateAdRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateAdRequest) GetAdURL() string {
	if x != nil {
		return x.AdURL
	}
	return ""
}

type GetAdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdvertiserID string `protobuf:"bytes,1,opt,name=AdvertiserID,proto3" json:"AdvertiserID,omitempty"`
	AdID         string `protobuf:"bytes,2,opt,name=AdID,proto3" json:"AdID,omitempty"`
}

func (x *GetAdRequest) Reset() {
	*x = GetAdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ads_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAdRequest) ProtoMessage() {}

func (x *GetAdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_ads_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAdRequest.ProtoReflect.Descriptor instead.
func (*GetAdRequest) Descriptor() ([]byte, []int) {
	return file_api_ads_proto_rawDescGZIP(), []int{2}
}

func (x *GetAdRequest) GetAdvertiserID() string {
	if x != nil {
		return x.AdvertiserID
	}
	return ""
}

func (x *GetAdRequest) GetAdID() string {
	if x != nil {
		return x.AdID
	}
	return ""
}

type Click struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdID       string `protobuf:"bytes,1,opt,name=AdID,proto3" json:"AdID,omitempty"`
	Timestamp  int64  `protobuf:"varint,2,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	IsAccepted bool   `protobuf:"varint,3,opt,name=IsAccepted,proto3" json:"IsAccepted,omitempty"`
}

func (x *Click) Reset() {
	*x = Click{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ads_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Click) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Click) ProtoMessage() {}

func (x *Click) ProtoReflect() protoreflect.Message {
	mi := &file_api_ads_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Click.ProtoReflect.Descriptor instead.
func (*Click) Descriptor() ([]byte, []int) {
	return file_api_ads_proto_rawDescGZIP(), []int{3}
}

func (x *Click) GetAdID() string {
	if x != nil {
		return x.AdID
	}
	return ""
}

func (x *Click) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Click) GetIsAccepted() bool {
	if x != nil {
		return x.IsAccepted
	}
	return false
}

type ClickCounter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdID        string `protobuf:"bytes,1,opt,name=AdID,proto3" json:"AdID,omitempty"`
	TotalClicks int32  `protobuf:"varint,2,opt,name=TotalClicks,proto3" json:"TotalClicks,omitempty"`
}

func (x *ClickCounter) Reset() {
	*x = ClickCounter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ads_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClickCounter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClickCounter) ProtoMessage() {}

func (x *ClickCounter) ProtoReflect() protoreflect.Message {
	mi := &file_api_ads_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClickCounter.ProtoReflect.Descriptor instead.
func (*ClickCounter) Descriptor() ([]byte, []int) {
	return file_api_ads_proto_rawDescGZIP(), []int{4}
}

func (x *ClickCounter) GetAdID() string {
	if x != nil {
		return x.AdID
	}
	return ""
}

func (x *ClickCounter) GetTotalClicks() int32 {
	if x != nil {
		return x.TotalClicks
	}
	return 0
}

type SendClickRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdID         string `protobuf:"bytes,1,opt,name=AdID,proto3" json:"AdID,omitempty"`
	ImpressionID string `protobuf:"bytes,2,opt,name=ImpressionID,proto3" json:"ImpressionID,omitempty"`
}

func (x *SendClickRequest) Reset() {
	*x = SendClickRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ads_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendClickRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendClickRequest) ProtoMessage() {}

func (x *SendClickRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_ads_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendClickRequest.ProtoReflect.Descriptor instead.
func (*SendClickRequest) Descriptor() ([]byte, []int) {
	return file_api_ads_proto_rawDescGZIP(), []int{5}
}

func (x *SendClickRequest) GetAdID() string {
	if x != nil {
		return x.AdID
	}
	return ""
}

func (x *SendClickRequest) GetImpressionID() string {
	if x != nil {
		return x.ImpressionID
	}
	return ""
}

type GetClicksCounterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdId string `protobuf:"bytes,1,opt,name=AdId,proto3" json:"AdId,omitempty"`
}

func (x *GetClicksCounterRequest) Reset() {
	*x = GetClicksCounterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ads_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClicksCounterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClicksCounterRequest) ProtoMessage() {}

func (x *GetClicksCounterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_ads_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClicksCounterRequest.ProtoReflect.Descriptor instead.
func (*GetClicksCounterRequest) Descriptor() ([]byte, []int) {
	return file_api_ads_proto_rawDescGZIP(), []int{6}
}

func (x *GetClicksCounterRequest) GetAdId() string {
	if x != nil {
		return x.AdId
	}
	return ""
}

var File_api_ads_proto protoreflect.FileDescriptor

var file_api_ads_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x61, 0x70, 0x69, 0x22, 0x8c, 0x01, 0x0a, 0x02, 0x41, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x41,
	0x64, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x41, 0x64, 0x49, 0x44, 0x12,
	0x22, 0x0a, 0x0c, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x64, 0x55,
	0x52, 0x4c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x64, 0x55, 0x52, 0x4c, 0x12,
	0x22, 0x0a, 0x0c, 0x49, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x49, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x22, 0x61, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74,
	0x69, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x41, 0x64,
	0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x41, 0x64, 0x55, 0x52, 0x4c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x41, 0x64, 0x55, 0x52, 0x4c, 0x22, 0x46, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74,
	0x69, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x41, 0x64,
	0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x41, 0x64,
	0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x41, 0x64, 0x49, 0x44, 0x22, 0x59,
	0x0a, 0x05, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x41, 0x64, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x41, 0x64, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x49, 0x73, 0x41,
	0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x49,
	0x73, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x22, 0x44, 0x0a, 0x0c, 0x43, 0x6c, 0x69,
	0x63, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x41, 0x64, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x41, 0x64, 0x49, 0x44, 0x12, 0x20, 0x0a,
	0x0b, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0b, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x73, 0x22,
	0x4a, 0x0a, 0x10, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x41, 0x64, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x41, 0x64, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c, 0x49, 0x6d, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x49,
	0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x22, 0x2d, 0x0a, 0x17, 0x47,
	0x65, 0x74, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x41, 0x64, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x41, 0x64, 0x49, 0x64, 0x32, 0x5c, 0x0a, 0x0a, 0x41, 0x64,
	0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x29, 0x0a, 0x08, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x64, 0x12, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x07, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x41, 0x64, 0x12, 0x23, 0x0a, 0x05, 0x47, 0x65, 0x74, 0x41, 0x64, 0x12, 0x11, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x07, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x64, 0x32, 0x87, 0x01, 0x0a, 0x11, 0x41, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2e,
	0x0a, 0x09, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x12, 0x15, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x12, 0x42,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65,
	0x72, 0x12, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x63, 0x6b,
	0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x65, 0x72, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x7a, 0x6f, 0x6e, 0x69, 0x6e, 0x6e, 0x69, 0x6b, 0x38, 0x39, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_ads_proto_rawDescOnce sync.Once
	file_api_ads_proto_rawDescData = file_api_ads_proto_rawDesc
)

func file_api_ads_proto_rawDescGZIP() []byte {
	file_api_ads_proto_rawDescOnce.Do(func() {
		file_api_ads_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_ads_proto_rawDescData)
	})
	return file_api_ads_proto_rawDescData
}

var file_api_ads_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_ads_proto_goTypes = []any{
	(*Ad)(nil),                      // 0: api.Ad
	(*CreateAdRequest)(nil),         // 1: api.CreateAdRequest
	(*GetAdRequest)(nil),            // 2: api.GetAdRequest
	(*Click)(nil),                   // 3: api.Click
	(*ClickCounter)(nil),            // 4: api.ClickCounter
	(*SendClickRequest)(nil),        // 5: api.SendClickRequest
	(*GetClicksCounterRequest)(nil), // 6: api.GetClicksCounterRequest
}
var file_api_ads_proto_depIdxs = []int32{
	1, // 0: api.AdsService.CreateAd:input_type -> api.CreateAdRequest
	2, // 1: api.AdsService.GetAd:input_type -> api.GetAdRequest
	5, // 2: api.AggregatorService.SendClick:input_type -> api.SendClickRequest
	6, // 3: api.AggregatorService.GetClickCounter:input_type -> api.GetClicksCounterRequest
	0, // 4: api.AdsService.CreateAd:output_type -> api.Ad
	0, // 5: api.AdsService.GetAd:output_type -> api.Ad
	3, // 6: api.AggregatorService.SendClick:output_type -> api.Click
	4, // 7: api.AggregatorService.GetClickCounter:output_type -> api.ClickCounter
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_ads_proto_init() }
func file_api_ads_proto_init() {
	if File_api_ads_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_ads_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Ad); i {
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
		file_api_ads_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateAdRequest); i {
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
		file_api_ads_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetAdRequest); i {
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
		file_api_ads_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Click); i {
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
		file_api_ads_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ClickCounter); i {
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
		file_api_ads_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*SendClickRequest); i {
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
		file_api_ads_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*GetClicksCounterRequest); i {
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
			RawDescriptor: file_api_ads_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_api_ads_proto_goTypes,
		DependencyIndexes: file_api_ads_proto_depIdxs,
		MessageInfos:      file_api_ads_proto_msgTypes,
	}.Build()
	File_api_ads_proto = out.File
	file_api_ads_proto_rawDesc = nil
	file_api_ads_proto_goTypes = nil
	file_api_ads_proto_depIdxs = nil
}
