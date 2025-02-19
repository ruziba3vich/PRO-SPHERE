// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.12.4
// source: feeds/models.proto

package feeds

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

// FeedCategory message
type FeedCategory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	IconUrl string `protobuf:"bytes,3,opt,name=icon_url,json=iconUrl,proto3" json:"icon_url,omitempty"`
}

func (x *FeedCategory) Reset() {
	*x = FeedCategory{}
	mi := &file_feeds_models_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FeedCategory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedCategory) ProtoMessage() {}

func (x *FeedCategory) ProtoReflect() protoreflect.Message {
	mi := &file_feeds_models_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedCategory.ProtoReflect.Descriptor instead.
func (*FeedCategory) Descriptor() ([]byte, []int) {
	return file_feeds_models_proto_rawDescGZIP(), []int{0}
}

func (x *FeedCategory) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FeedCategory) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FeedCategory) GetIconUrl() string {
	if x != nil {
		return x.IconUrl
	}
	return ""
}

// FeedItem message
type FeedItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FeedId      int32  `protobuf:"varint,1,opt,name=feed_id,json=feedId,proto3" json:"feed_id,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	ImageUrl    string `protobuf:"bytes,3,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	PublishedAt string `protobuf:"bytes,5,opt,name=published_at,json=publishedAt,proto3" json:"published_at,omitempty"`
}

func (x *FeedItem) Reset() {
	*x = FeedItem{}
	mi := &file_feeds_models_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FeedItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedItem) ProtoMessage() {}

func (x *FeedItem) ProtoReflect() protoreflect.Message {
	mi := &file_feeds_models_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedItem.ProtoReflect.Descriptor instead.
func (*FeedItem) Descriptor() ([]byte, []int) {
	return file_feeds_models_proto_rawDescGZIP(), []int{1}
}

func (x *FeedItem) GetFeedId() int32 {
	if x != nil {
		return x.FeedId
	}
	return 0
}

func (x *FeedItem) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *FeedItem) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *FeedItem) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *FeedItem) GetPublishedAt() string {
	if x != nil {
		return x.PublishedAt
	}
	return ""
}

// FeedItemResponse message
type FeedItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FeedId      int32  `protobuf:"varint,2,opt,name=feed_id,json=feedId,proto3" json:"feed_id,omitempty"`
	Title       string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	ImageUrl    string `protobuf:"bytes,5,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	PublishedAt string `protobuf:"bytes,6,opt,name=published_at,json=publishedAt,proto3" json:"published_at,omitempty"`
	CreatedAt   string `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   string `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *FeedItemResponse) Reset() {
	*x = FeedItemResponse{}
	mi := &file_feeds_models_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FeedItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedItemResponse) ProtoMessage() {}

func (x *FeedItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_feeds_models_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedItemResponse.ProtoReflect.Descriptor instead.
func (*FeedItemResponse) Descriptor() ([]byte, []int) {
	return file_feeds_models_proto_rawDescGZIP(), []int{2}
}

func (x *FeedItemResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FeedItemResponse) GetFeedId() int32 {
	if x != nil {
		return x.FeedId
	}
	return 0
}

func (x *FeedItemResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *FeedItemResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *FeedItemResponse) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *FeedItemResponse) GetPublishedAt() string {
	if x != nil {
		return x.PublishedAt
	}
	return ""
}

func (x *FeedItemResponse) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *FeedItemResponse) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

// UpdateItem message
type UpdateItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	PublishedAt string `protobuf:"bytes,4,opt,name=published_at,json=publishedAt,proto3" json:"published_at,omitempty"`
	ImageUrl    string `protobuf:"bytes,5,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
}

func (x *UpdateItem) Reset() {
	*x = UpdateItem{}
	mi := &file_feeds_models_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateItem) ProtoMessage() {}

func (x *UpdateItem) ProtoReflect() protoreflect.Message {
	mi := &file_feeds_models_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateItem.ProtoReflect.Descriptor instead.
func (*UpdateItem) Descriptor() ([]byte, []int) {
	return file_feeds_models_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateItem) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateItem) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateItem) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateItem) GetPublishedAt() string {
	if x != nil {
		return x.PublishedAt
	}
	return ""
}

func (x *UpdateItem) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

// Feed message
type Feed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Link        string `protobuf:"bytes,3,opt,name=link,proto3" json:"link,omitempty"`
	CategoryId  int32  `protobuf:"varint,4,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	LogoUrl     string `protobuf:"bytes,5,opt,name=logo_url,json=logoUrl,proto3" json:"logo_url,omitempty"`
	Priority    int32  `protobuf:"varint,6,opt,name=priority,proto3" json:"priority,omitempty"`
	MaxItems    int32  `protobuf:"varint,7,opt,name=max_items,json=maxItems,proto3" json:"max_items,omitempty"`
}

func (x *Feed) Reset() {
	*x = Feed{}
	mi := &file_feeds_models_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Feed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Feed) ProtoMessage() {}

func (x *Feed) ProtoReflect() protoreflect.Message {
	mi := &file_feeds_models_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Feed.ProtoReflect.Descriptor instead.
func (*Feed) Descriptor() ([]byte, []int) {
	return file_feeds_models_proto_rawDescGZIP(), []int{4}
}

func (x *Feed) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Feed) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Feed) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *Feed) GetCategoryId() int32 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

func (x *Feed) GetLogoUrl() string {
	if x != nil {
		return x.LogoUrl
	}
	return ""
}

func (x *Feed) GetPriority() int32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *Feed) GetMaxItems() int32 {
	if x != nil {
		return x.MaxItems
	}
	return 0
}

// FeedResponse message
type FeedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description   string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Link          string `protobuf:"bytes,4,opt,name=link,proto3" json:"link,omitempty"`
	CategoryId    int32  `protobuf:"varint,5,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	LogoUrl       string `protobuf:"bytes,6,opt,name=logo_url,json=logoUrl,proto3" json:"logo_url,omitempty"`
	Priority      int32  `protobuf:"varint,7,opt,name=priority,proto3" json:"priority,omitempty"`
	MaxItems      int32  `protobuf:"varint,8,opt,name=max_items,json=maxItems,proto3" json:"max_items,omitempty"`
	LastRefreshed string `protobuf:"bytes,9,opt,name=last_refreshed,json=lastRefreshed,proto3" json:"last_refreshed,omitempty"`
	CreatedAt     string `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *FeedResponse) Reset() {
	*x = FeedResponse{}
	mi := &file_feeds_models_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FeedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedResponse) ProtoMessage() {}

func (x *FeedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_feeds_models_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedResponse.ProtoReflect.Descriptor instead.
func (*FeedResponse) Descriptor() ([]byte, []int) {
	return file_feeds_models_proto_rawDescGZIP(), []int{5}
}

func (x *FeedResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FeedResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *FeedResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *FeedResponse) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *FeedResponse) GetCategoryId() int32 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

func (x *FeedResponse) GetLogoUrl() string {
	if x != nil {
		return x.LogoUrl
	}
	return ""
}

func (x *FeedResponse) GetPriority() int32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *FeedResponse) GetMaxItems() int32 {
	if x != nil {
		return x.MaxItems
	}
	return 0
}

func (x *FeedResponse) GetLastRefreshed() string {
	if x != nil {
		return x.LastRefreshed
	}
	return ""
}

func (x *FeedResponse) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *FeedResponse) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

// UpdateFeed message
type UpdateFeed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Link        string `protobuf:"bytes,4,opt,name=link,proto3" json:"link,omitempty"`
	CategoryId  int32  `protobuf:"varint,5,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	LogoUrl     string `protobuf:"bytes,6,opt,name=logo_url,json=logoUrl,proto3" json:"logo_url,omitempty"`
	Priority    int32  `protobuf:"varint,7,opt,name=priority,proto3" json:"priority,omitempty"`
	MaxItems    int32  `protobuf:"varint,8,opt,name=max_items,json=maxItems,proto3" json:"max_items,omitempty"`
}

func (x *UpdateFeed) Reset() {
	*x = UpdateFeed{}
	mi := &file_feeds_models_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateFeed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFeed) ProtoMessage() {}

func (x *UpdateFeed) ProtoReflect() protoreflect.Message {
	mi := &file_feeds_models_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFeed.ProtoReflect.Descriptor instead.
func (*UpdateFeed) Descriptor() ([]byte, []int) {
	return file_feeds_models_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateFeed) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateFeed) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateFeed) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateFeed) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *UpdateFeed) GetCategoryId() int32 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

func (x *UpdateFeed) GetLogoUrl() string {
	if x != nil {
		return x.LogoUrl
	}
	return ""
}

func (x *UpdateFeed) GetPriority() int32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *UpdateFeed) GetMaxItems() int32 {
	if x != nil {
		return x.MaxItems
	}
	return 0
}

var File_feeds_models_proto protoreflect.FileDescriptor

var file_feeds_models_proto_rawDesc = []byte{
	0x0a, 0x12, 0x66, 0x65, 0x65, 0x64, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x22, 0x4d, 0x0a, 0x0c,
	0x46, 0x65, 0x65, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x69, 0x63, 0x6f, 0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x69, 0x63, 0x6f, 0x6e, 0x55, 0x72, 0x6c, 0x22, 0x9b, 0x01, 0x0a, 0x08,
	0x46, 0x65, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x65, 0x65, 0x64,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x66, 0x65, 0x65, 0x64, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x55, 0x72, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x41, 0x74, 0x22, 0xf1, 0x01, 0x0a, 0x10, 0x46, 0x65,
	0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x66, 0x65, 0x65, 0x64, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x21, 0x0a, 0x0c,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x94, 0x01,
	0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x55, 0x72, 0x6c, 0x22, 0xc7, 0x01, 0x0a, 0x04, 0x46, 0x65, 0x65, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x6f,
	0x67, 0x6f, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f,
	0x67, 0x6f, 0x55, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x22, 0xc4,
	0x02, 0x0a, 0x0c, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x1f, 0x0a, 0x0b, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08,
	0x6c, 0x6f, 0x67, 0x6f, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6c, 0x6f, 0x67, 0x6f, 0x55, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72,
	0x69, 0x74, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72,
	0x69, 0x74, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x49, 0x74, 0x65, 0x6d, 0x73,
	0x12, 0x25, 0x0a, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x52, 0x65,
	0x66, 0x72, 0x65, 0x73, 0x68, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xdd, 0x01, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x46, 0x65, 0x65, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b,
	0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49,
	0x64, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x6f, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f, 0x67, 0x6f, 0x55, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6d, 0x61, 0x78,
	0x49, 0x74, 0x65, 0x6d, 0x73, 0x42, 0x10, 0x5a, 0x0e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_feeds_models_proto_rawDescOnce sync.Once
	file_feeds_models_proto_rawDescData = file_feeds_models_proto_rawDesc
)

func file_feeds_models_proto_rawDescGZIP() []byte {
	file_feeds_models_proto_rawDescOnce.Do(func() {
		file_feeds_models_proto_rawDescData = protoimpl.X.CompressGZIP(file_feeds_models_proto_rawDescData)
	})
	return file_feeds_models_proto_rawDescData
}

var file_feeds_models_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_feeds_models_proto_goTypes = []any{
	(*FeedCategory)(nil),     // 0: models.FeedCategory
	(*FeedItem)(nil),         // 1: models.FeedItem
	(*FeedItemResponse)(nil), // 2: models.FeedItemResponse
	(*UpdateItem)(nil),       // 3: models.UpdateItem
	(*Feed)(nil),             // 4: models.Feed
	(*FeedResponse)(nil),     // 5: models.FeedResponse
	(*UpdateFeed)(nil),       // 6: models.UpdateFeed
}
var file_feeds_models_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_feeds_models_proto_init() }
func file_feeds_models_proto_init() {
	if File_feeds_models_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_feeds_models_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_feeds_models_proto_goTypes,
		DependencyIndexes: file_feeds_models_proto_depIdxs,
		MessageInfos:      file_feeds_models_proto_msgTypes,
	}.Build()
	File_feeds_models_proto = out.File
	file_feeds_models_proto_rawDesc = nil
	file_feeds_models_proto_goTypes = nil
	file_feeds_models_proto_depIdxs = nil
}
