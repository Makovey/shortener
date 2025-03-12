// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.29.3
// source: shortener.proto

package shortener

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PostURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LongUrl string `protobuf:"bytes,1,opt,name=long_url,json=longUrl,proto3" json:"long_url,omitempty"`
}

func (x *PostURLRequest) Reset() {
	*x = PostURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shortener_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostURLRequest) ProtoMessage() {}

func (x *PostURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostURLRequest.ProtoReflect.Descriptor instead.
func (*PostURLRequest) Descriptor() ([]byte, []int) {
	return file_shortener_proto_rawDescGZIP(), []int{0}
}

func (x *PostURLRequest) GetLongUrl() string {
	if x != nil {
		return x.LongUrl
	}
	return ""
}

type PostURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FullShortUrl string `protobuf:"bytes,1,opt,name=full_short_url,json=fullShortUrl,proto3" json:"full_short_url,omitempty"`
}

func (x *PostURLResponse) Reset() {
	*x = PostURLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shortener_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostURLResponse) ProtoMessage() {}

func (x *PostURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostURLResponse.ProtoReflect.Descriptor instead.
func (*PostURLResponse) Descriptor() ([]byte, []int) {
	return file_shortener_proto_rawDescGZIP(), []int{1}
}

func (x *PostURLResponse) GetFullShortUrl() string {
	if x != nil {
		return x.FullShortUrl
	}
	return ""
}

type GetURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortUrl string `protobuf:"bytes,1,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
}

func (x *GetURLRequest) Reset() {
	*x = GetURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shortener_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetURLRequest) ProtoMessage() {}

func (x *GetURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetURLRequest.ProtoReflect.Descriptor instead.
func (*GetURLRequest) Descriptor() ([]byte, []int) {
	return file_shortener_proto_rawDescGZIP(), []int{2}
}

func (x *GetURLRequest) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

type GetURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LongUrl string `protobuf:"bytes,1,opt,name=long_url,json=longUrl,proto3" json:"long_url,omitempty"`
}

func (x *GetURLResponse) Reset() {
	*x = GetURLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shortener_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetURLResponse) ProtoMessage() {}

func (x *GetURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetURLResponse.ProtoReflect.Descriptor instead.
func (*GetURLResponse) Descriptor() ([]byte, []int) {
	return file_shortener_proto_rawDescGZIP(), []int{3}
}

func (x *GetURLResponse) GetLongUrl() string {
	if x != nil {
		return x.LongUrl
	}
	return ""
}

type PostURLRequestV2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *PostURLRequestV2) Reset() {
	*x = PostURLRequestV2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shortener_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostURLRequestV2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostURLRequestV2) ProtoMessage() {}

func (x *PostURLRequestV2) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostURLRequestV2.ProtoReflect.Descriptor instead.
func (*PostURLRequestV2) Descriptor() ([]byte, []int) {
	return file_shortener_proto_rawDescGZIP(), []int{4}
}

func (x *PostURLRequestV2) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type PostURLResponseV2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *PostURLResponseV2) Reset() {
	*x = PostURLResponseV2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shortener_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostURLResponseV2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostURLResponseV2) ProtoMessage() {}

func (x *PostURLResponseV2) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostURLResponseV2.ProtoReflect.Descriptor instead.
func (*PostURLResponseV2) Descriptor() ([]byte, []int) {
	return file_shortener_proto_rawDescGZIP(), []int{5}
}

func (x *PostURLResponseV2) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type PostBatchURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Batch []*BatchURLRequest `protobuf:"bytes,1,rep,name=batch,proto3" json:"batch,omitempty"`
}

func (x *PostBatchURLRequest) Reset() {
	*x = PostBatchURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shortener_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostBatchURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostBatchURLRequest) ProtoMessage() {}

func (x *PostBatchURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostBatchURLRequest.ProtoReflect.Descriptor instead.
func (*PostBatchURLRequest) Descriptor() ([]byte, []int) {
	return file_shortener_proto_rawDescGZIP(), []int{6}
}

func (x *PostBatchURLRequest) GetBatch() []*BatchURLRequest {
	if x != nil {
		return x.Batch
	}
	return nil
}

type PostBatchURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Batch []*BatchURLResponse `protobuf:"bytes,2,rep,name=batch,proto3" json:"batch,omitempty"`
}

func (x *PostBatchURLResponse) Reset() {
	*x = PostBatchURLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shortener_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostBatchURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostBatchURLResponse) ProtoMessage() {}

func (x *PostBatchURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostBatchURLResponse.ProtoReflect.Descriptor instead.
func (*PostBatchURLResponse) Descriptor() ([]byte, []int) {
	return file_shortener_proto_rawDescGZIP(), []int{7}
}

func (x *PostBatchURLResponse) GetBatch() []*BatchURLResponse {
	if x != nil {
		return x.Batch
	}
	return nil
}

type BatchURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CorrelationID string `protobuf:"bytes,1,opt,name=correlationID,proto3" json:"correlationID,omitempty"`
	OriginalURL   string `protobuf:"bytes,2,opt,name=originalURL,proto3" json:"originalURL,omitempty"`
}

func (x *BatchURLRequest) Reset() {
	*x = BatchURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shortener_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchURLRequest) ProtoMessage() {}

func (x *BatchURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchURLRequest.ProtoReflect.Descriptor instead.
func (*BatchURLRequest) Descriptor() ([]byte, []int) {
	return file_shortener_proto_rawDescGZIP(), []int{8}
}

func (x *BatchURLRequest) GetCorrelationID() string {
	if x != nil {
		return x.CorrelationID
	}
	return ""
}

func (x *BatchURLRequest) GetOriginalURL() string {
	if x != nil {
		return x.OriginalURL
	}
	return ""
}

type BatchURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CorrelationID string `protobuf:"bytes,1,opt,name=correlationID,proto3" json:"correlationID,omitempty"`
	ShortURL      string `protobuf:"bytes,2,opt,name=shortURL,proto3" json:"shortURL,omitempty"`
}

func (x *BatchURLResponse) Reset() {
	*x = BatchURLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shortener_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchURLResponse) ProtoMessage() {}

func (x *BatchURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchURLResponse.ProtoReflect.Descriptor instead.
func (*BatchURLResponse) Descriptor() ([]byte, []int) {
	return file_shortener_proto_rawDescGZIP(), []int{9}
}

func (x *BatchURLResponse) GetCorrelationID() string {
	if x != nil {
		return x.CorrelationID
	}
	return ""
}

func (x *BatchURLResponse) GetShortURL() string {
	if x != nil {
		return x.ShortURL
	}
	return ""
}

type GetUserURLsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserURLs []*UserURL `protobuf:"bytes,1,rep,name=userURLs,proto3" json:"userURLs,omitempty"`
}

func (x *GetUserURLsResponse) Reset() {
	*x = GetUserURLsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shortener_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserURLsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserURLsResponse) ProtoMessage() {}

func (x *GetUserURLsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserURLsResponse.ProtoReflect.Descriptor instead.
func (*GetUserURLsResponse) Descriptor() ([]byte, []int) {
	return file_shortener_proto_rawDescGZIP(), []int{10}
}

func (x *GetUserURLsResponse) GetUserURLs() []*UserURL {
	if x != nil {
		return x.UserURLs
	}
	return nil
}

type UserURL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CorrelationID string `protobuf:"bytes,1,opt,name=correlationID,proto3" json:"correlationID,omitempty"`
	OriginalURL   string `protobuf:"bytes,2,opt,name=originalURL,proto3" json:"originalURL,omitempty"`
	ShortURL      string `protobuf:"bytes,3,opt,name=shortURL,proto3" json:"shortURL,omitempty"`
}

func (x *UserURL) Reset() {
	*x = UserURL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shortener_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserURL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserURL) ProtoMessage() {}

func (x *UserURL) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserURL.ProtoReflect.Descriptor instead.
func (*UserURL) Descriptor() ([]byte, []int) {
	return file_shortener_proto_rawDescGZIP(), []int{11}
}

func (x *UserURL) GetCorrelationID() string {
	if x != nil {
		return x.CorrelationID
	}
	return ""
}

func (x *UserURL) GetOriginalURL() string {
	if x != nil {
		return x.OriginalURL
	}
	return ""
}

func (x *UserURL) GetShortURL() string {
	if x != nil {
		return x.ShortURL
	}
	return ""
}

type DeleteUserURLsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortURLs []string `protobuf:"bytes,1,rep,name=shortURLs,proto3" json:"shortURLs,omitempty"`
}

func (x *DeleteUserURLsRequest) Reset() {
	*x = DeleteUserURLsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shortener_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserURLsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserURLsRequest) ProtoMessage() {}

func (x *DeleteUserURLsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserURLsRequest.ProtoReflect.Descriptor instead.
func (*DeleteUserURLsRequest) Descriptor() ([]byte, []int) {
	return file_shortener_proto_rawDescGZIP(), []int{12}
}

func (x *DeleteUserURLsRequest) GetShortURLs() []string {
	if x != nil {
		return x.ShortURLs
	}
	return nil
}

var File_shortener_proto protoreflect.FileDescriptor

var file_shortener_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x75, 0x72, 0x6c, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x2b, 0x0a, 0x0e, 0x50, 0x6f, 0x73, 0x74, 0x55, 0x52, 0x4c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x6f, 0x6e, 0x67, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f, 0x6e, 0x67, 0x55, 0x72, 0x6c,
	0x22, 0x37, 0x0a, 0x0f, 0x50, 0x6f, 0x73, 0x74, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x75, 0x6c,
	0x6c, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x22, 0x2c, 0x0a, 0x0d, 0x47, 0x65, 0x74,
	0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73,
	0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x22, 0x2b, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x52,
	0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x6f, 0x6e,
	0x67, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f, 0x6e,
	0x67, 0x55, 0x72, 0x6c, 0x22, 0x24, 0x0a, 0x10, 0x50, 0x6f, 0x73, 0x74, 0x55, 0x52, 0x4c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x32, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x2b, 0x0a, 0x11, 0x50, 0x6f,
	0x73, 0x74, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x32, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x41, 0x0a, 0x13, 0x50, 0x6f, 0x73, 0x74, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a,
	0x0a, 0x05, 0x62, 0x61, 0x74, 0x63, 0x68, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x75, 0x72, 0x6c, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x05, 0x62, 0x61, 0x74, 0x63, 0x68, 0x22, 0x43, 0x0a, 0x14, 0x50, 0x6f,
	0x73, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2b, 0x0a, 0x05, 0x62, 0x61, 0x74, 0x63, 0x68, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x75, 0x72, 0x6c, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x55, 0x52, 0x4c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x62, 0x61, 0x74, 0x63, 0x68, 0x22,
	0x59, 0x0a, 0x0f, 0x42, 0x61, 0x74, 0x63, 0x68, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x61, 0x6c, 0x55, 0x52, 0x4c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x52, 0x4c, 0x22, 0x54, 0x0a, 0x10, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24,
	0x0a, 0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c,
	0x22, 0x3f, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x55, 0x52, 0x4c, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x55,
	0x52, 0x4c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x75, 0x72, 0x6c, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x55, 0x52, 0x4c, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x55, 0x52, 0x4c,
	0x73, 0x22, 0x6d, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x55, 0x52, 0x4c, 0x12, 0x24, 0x0a, 0x0d,
	0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x52,
	0x4c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61,
	0x6c, 0x55, 0x52, 0x4c, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c,
	0x22, 0x35, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x55, 0x52,
	0x4c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x55, 0x52, 0x4c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x73, 0x32, 0xfc, 0x02, 0x0a, 0x09, 0x53, 0x68, 0x6f, 0x72,
	0x74, 0x65, 0x6e, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x07, 0x50, 0x6f, 0x73, 0x74, 0x55, 0x52, 0x4c,
	0x12, 0x13, 0x2e, 0x75, 0x72, 0x6c, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x55, 0x52, 0x4c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x75, 0x72, 0x6c, 0x2e, 0x50, 0x6f, 0x73, 0x74,
	0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x06, 0x47,
	0x65, 0x74, 0x55, 0x52, 0x4c, 0x12, 0x12, 0x2e, 0x75, 0x72, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x55,
	0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x75, 0x72, 0x6c, 0x2e,
	0x47, 0x65, 0x74, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a,
	0x0a, 0x09, 0x50, 0x6f, 0x73, 0x74, 0x55, 0x52, 0x4c, 0x56, 0x32, 0x12, 0x15, 0x2e, 0x75, 0x72,
	0x6c, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x56, 0x32, 0x1a, 0x16, 0x2e, 0x75, 0x72, 0x6c, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x55, 0x52, 0x4c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x32, 0x12, 0x43, 0x0a, 0x0c, 0x50, 0x6f,
	0x73, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x55, 0x52, 0x4c, 0x12, 0x18, 0x2e, 0x75, 0x72, 0x6c,
	0x2e, 0x50, 0x6f, 0x73, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x75, 0x72, 0x6c, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3f, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x55, 0x52, 0x4c, 0x73, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x18, 0x2e, 0x75, 0x72, 0x6c, 0x2e, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x55, 0x52, 0x4c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x44, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x55, 0x52,
	0x4c, 0x73, 0x12, 0x1a, 0x2e, 0x75, 0x72, 0x6c, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x55, 0x52, 0x4c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x61, 0x6b, 0x6f, 0x76, 0x65, 0x79, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_shortener_proto_rawDescOnce sync.Once
	file_shortener_proto_rawDescData = file_shortener_proto_rawDesc
)

func file_shortener_proto_rawDescGZIP() []byte {
	file_shortener_proto_rawDescOnce.Do(func() {
		file_shortener_proto_rawDescData = protoimpl.X.CompressGZIP(file_shortener_proto_rawDescData)
	})
	return file_shortener_proto_rawDescData
}

var file_shortener_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_shortener_proto_goTypes = []interface{}{
	(*PostURLRequest)(nil),        // 0: url.PostURLRequest
	(*PostURLResponse)(nil),       // 1: url.PostURLResponse
	(*GetURLRequest)(nil),         // 2: url.GetURLRequest
	(*GetURLResponse)(nil),        // 3: url.GetURLResponse
	(*PostURLRequestV2)(nil),      // 4: url.PostURLRequestV2
	(*PostURLResponseV2)(nil),     // 5: url.PostURLResponseV2
	(*PostBatchURLRequest)(nil),   // 6: url.PostBatchURLRequest
	(*PostBatchURLResponse)(nil),  // 7: url.PostBatchURLResponse
	(*BatchURLRequest)(nil),       // 8: url.BatchURLRequest
	(*BatchURLResponse)(nil),      // 9: url.BatchURLResponse
	(*GetUserURLsResponse)(nil),   // 10: url.GetUserURLsResponse
	(*UserURL)(nil),               // 11: url.UserURL
	(*DeleteUserURLsRequest)(nil), // 12: url.DeleteUserURLsRequest
	(*emptypb.Empty)(nil),         // 13: google.protobuf.Empty
}
var file_shortener_proto_depIdxs = []int32{
	8,  // 0: url.PostBatchURLRequest.batch:type_name -> url.BatchURLRequest
	9,  // 1: url.PostBatchURLResponse.batch:type_name -> url.BatchURLResponse
	11, // 2: url.GetUserURLsResponse.userURLs:type_name -> url.UserURL
	0,  // 3: url.Shortener.PostURL:input_type -> url.PostURLRequest
	2,  // 4: url.Shortener.GetURL:input_type -> url.GetURLRequest
	4,  // 5: url.Shortener.PostURLV2:input_type -> url.PostURLRequestV2
	6,  // 6: url.Shortener.PostBatchURL:input_type -> url.PostBatchURLRequest
	13, // 7: url.Shortener.GetUserURLs:input_type -> google.protobuf.Empty
	12, // 8: url.Shortener.DeleteUserURLs:input_type -> url.DeleteUserURLsRequest
	1,  // 9: url.Shortener.PostURL:output_type -> url.PostURLResponse
	3,  // 10: url.Shortener.GetURL:output_type -> url.GetURLResponse
	5,  // 11: url.Shortener.PostURLV2:output_type -> url.PostURLResponseV2
	7,  // 12: url.Shortener.PostBatchURL:output_type -> url.PostBatchURLResponse
	10, // 13: url.Shortener.GetUserURLs:output_type -> url.GetUserURLsResponse
	13, // 14: url.Shortener.DeleteUserURLs:output_type -> google.protobuf.Empty
	9,  // [9:15] is the sub-list for method output_type
	3,  // [3:9] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_shortener_proto_init() }
func file_shortener_proto_init() {
	if File_shortener_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shortener_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostURLRequest); i {
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
		file_shortener_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostURLResponse); i {
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
		file_shortener_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetURLRequest); i {
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
		file_shortener_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetURLResponse); i {
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
		file_shortener_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostURLRequestV2); i {
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
		file_shortener_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostURLResponseV2); i {
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
		file_shortener_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostBatchURLRequest); i {
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
		file_shortener_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostBatchURLResponse); i {
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
		file_shortener_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchURLRequest); i {
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
		file_shortener_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchURLResponse); i {
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
		file_shortener_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserURLsResponse); i {
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
		file_shortener_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserURL); i {
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
		file_shortener_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserURLsRequest); i {
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
			RawDescriptor: file_shortener_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shortener_proto_goTypes,
		DependencyIndexes: file_shortener_proto_depIdxs,
		MessageInfos:      file_shortener_proto_msgTypes,
	}.Build()
	File_shortener_proto = out.File
	file_shortener_proto_rawDesc = nil
	file_shortener_proto_goTypes = nil
	file_shortener_proto_depIdxs = nil
}
