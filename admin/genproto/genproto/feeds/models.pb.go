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

type Translation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lang string `protobuf:"bytes,1,opt,name=lang,proto3" json:"lang,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Translation) Reset() {
	*x = Translation{}
	mi := &file_feeds_models_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Translation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Translation) ProtoMessage() {}

func (x *Translation) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Translation.ProtoReflect.Descriptor instead.
func (*Translation) Descriptor() ([]byte, []int) {
	return file_feeds_models_proto_rawDescGZIP(), []int{0}
}

func (x *Translation) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

func (x *Translation) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type FeedCategory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int32          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	IconUrl      string         `protobuf:"bytes,2,opt,name=icon_url,json=iconUrl,proto3" json:"icon_url,omitempty"`
	IconId       string         `protobuf:"bytes,3,opt,name=icon_id,json=iconId,proto3" json:"icon_id,omitempty"`
	Translations []*Translation `protobuf:"bytes,4,rep,name=translations,proto3" json:"translations,omitempty"`
}

func (x *FeedCategory) Reset() {
	*x = FeedCategory{}
	mi := &file_feeds_models_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FeedCategory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedCategory) ProtoMessage() {}

func (x *FeedCategory) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use FeedCategory.ProtoReflect.Descriptor instead.
func (*FeedCategory) Descriptor() ([]byte, []int) {
	return file_feeds_models_proto_rawDescGZIP(), []int{1}
}

func (x *FeedCategory) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FeedCategory) GetIconUrl() string {
	if x != nil {
		return x.IconUrl
	}
	return ""
}

func (x *FeedCategory) GetIconId() string {
	if x != nil {
		return x.IconId
	}
	return ""
}

func (x *FeedCategory) GetTranslations() []*Translation {
	if x != nil {
		return x.Translations
	}
	return nil
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
	mi := &file_feeds_models_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FeedItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedItem) ProtoMessage() {}

func (x *FeedItem) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use FeedItem.ProtoReflect.Descriptor instead.
func (*FeedItem) Descriptor() ([]byte, []int) {
	return file_feeds_models_proto_rawDescGZIP(), []int{2}
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
	mi := &file_feeds_models_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FeedItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedItemResponse) ProtoMessage() {}

func (x *FeedItemResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use FeedItemResponse.ProtoReflect.Descriptor instead.
func (*FeedItemResponse) Descriptor() ([]byte, []int) {
	return file_feeds_models_proto_rawDescGZIP(), []int{3}
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
	mi := &file_feeds_models_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateItem) ProtoMessage() {}

func (x *UpdateItem) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use UpdateItem.ProtoReflect.Descriptor instead.
func (*UpdateItem) Descriptor() ([]byte, []int) {
	return file_feeds_models_proto_rawDescGZIP(), []int{4}
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

	Id          int64               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	LogoUrl     string              `protobuf:"bytes,2,opt,name=logo_url,json=logoUrl,proto3" json:"logo_url,omitempty"`
	LogoUrlId   string              `protobuf:"bytes,3,opt,name=logo_url_id,json=logoUrlId,proto3" json:"logo_url_id,omitempty"`
	Priority    int32               `protobuf:"varint,4,opt,name=priority,proto3" json:"priority,omitempty"`
	MaxItems    int32               `protobuf:"varint,5,opt,name=max_items,json=maxItems,proto3" json:"max_items,omitempty"`
	BaseUrl     string              `protobuf:"bytes,6,opt,name=base_url,json=baseUrl,proto3" json:"base_url,omitempty"`
	CreatedAt   string              `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   string              `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Translation []*FeedTransalation `protobuf:"bytes,9,rep,name=translation,proto3" json:"translation,omitempty"`
}

func (x *Feed) Reset() {
	*x = Feed{}
	mi := &file_feeds_models_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Feed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Feed) ProtoMessage() {}

func (x *Feed) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Feed.ProtoReflect.Descriptor instead.
func (*Feed) Descriptor() ([]byte, []int) {
	return file_feeds_models_proto_rawDescGZIP(), []int{5}
}

func (x *Feed) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Feed) GetLogoUrl() string {
	if x != nil {
		return x.LogoUrl
	}
	return ""
}

func (x *Feed) GetLogoUrlId() string {
	if x != nil {
		return x.LogoUrlId
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

func (x *Feed) GetBaseUrl() string {
	if x != nil {
		return x.BaseUrl
	}
	return ""
}

func (x *Feed) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Feed) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Feed) GetTranslation() []*FeedTransalation {
	if x != nil {
		return x.Translation
	}
	return nil
}

type FeedTransalation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FeedId      int64  `protobuf:"varint,2,opt,name=feed_id,json=feedId,proto3" json:"feed_id,omitempty"`
	Title       string `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
	Lang        string `protobuf:"bytes,3,opt,name=lang,proto3" json:"lang,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *FeedTransalation) Reset() {
	*x = FeedTransalation{}
	mi := &file_feeds_models_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FeedTransalation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedTransalation) ProtoMessage() {}

func (x *FeedTransalation) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use FeedTransalation.ProtoReflect.Descriptor instead.
func (*FeedTransalation) Descriptor() ([]byte, []int) {
	return file_feeds_models_proto_rawDescGZIP(), []int{6}
}

func (x *FeedTransalation) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FeedTransalation) GetFeedId() int64 {
	if x != nil {
		return x.FeedId
	}
	return 0
}

func (x *FeedTransalation) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *FeedTransalation) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

func (x *FeedTransalation) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type FeedContent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FeedId     int64  `protobuf:"varint,2,opt,name=feed_id,json=feedId,proto3" json:"feed_id,omitempty"`
	CategoryId int64  `protobuf:"varint,3,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	Lang       string `protobuf:"bytes,4,opt,name=lang,proto3" json:"lang,omitempty"`
	Link       string `protobuf:"bytes,5,opt,name=link,proto3" json:"link,omitempty"`
}

func (x *FeedContent) Reset() {
	*x = FeedContent{}
	mi := &file_feeds_models_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FeedContent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedContent) ProtoMessage() {}

func (x *FeedContent) ProtoReflect() protoreflect.Message {
	mi := &file_feeds_models_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedContent.ProtoReflect.Descriptor instead.
func (*FeedContent) Descriptor() ([]byte, []int) {
	return file_feeds_models_proto_rawDescGZIP(), []int{7}
}

func (x *FeedContent) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FeedContent) GetFeedId() int64 {
	if x != nil {
		return x.FeedId
	}
	return 0
}

func (x *FeedContent) GetCategoryId() int64 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

func (x *FeedContent) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

func (x *FeedContent) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

type CreateFeedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LogoUrl     string                    `protobuf:"bytes,1,opt,name=logo_url,json=logoUrl,proto3" json:"logo_url,omitempty"`
	LogoUrlId   string                    `protobuf:"bytes,2,opt,name=logo_url_id,json=logoUrlId,proto3" json:"logo_url_id,omitempty"`
	Priority    int32                     `protobuf:"varint,3,opt,name=priority,proto3" json:"priority,omitempty"`
	MaxItems    int32                     `protobuf:"varint,4,opt,name=max_items,json=maxItems,proto3" json:"max_items,omitempty"`
	BaseUrl     string                    `protobuf:"bytes,5,opt,name=base_url,json=baseUrl,proto3" json:"base_url,omitempty"`
	Translation []*CreateFeedTransalation `protobuf:"bytes,6,rep,name=translation,proto3" json:"translation,omitempty"`
}

func (x *CreateFeedRequest) Reset() {
	*x = CreateFeedRequest{}
	mi := &file_feeds_models_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateFeedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFeedRequest) ProtoMessage() {}

func (x *CreateFeedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_feeds_models_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFeedRequest.ProtoReflect.Descriptor instead.
func (*CreateFeedRequest) Descriptor() ([]byte, []int) {
	return file_feeds_models_proto_rawDescGZIP(), []int{8}
}

func (x *CreateFeedRequest) GetLogoUrl() string {
	if x != nil {
		return x.LogoUrl
	}
	return ""
}

func (x *CreateFeedRequest) GetLogoUrlId() string {
	if x != nil {
		return x.LogoUrlId
	}
	return ""
}

func (x *CreateFeedRequest) GetPriority() int32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *CreateFeedRequest) GetMaxItems() int32 {
	if x != nil {
		return x.MaxItems
	}
	return 0
}

func (x *CreateFeedRequest) GetBaseUrl() string {
	if x != nil {
		return x.BaseUrl
	}
	return ""
}

func (x *CreateFeedRequest) GetTranslation() []*CreateFeedTransalation {
	if x != nil {
		return x.Translation
	}
	return nil
}

type CreateFeedTransalation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
	Lang        string `protobuf:"bytes,3,opt,name=lang,proto3" json:"lang,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *CreateFeedTransalation) Reset() {
	*x = CreateFeedTransalation{}
	mi := &file_feeds_models_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateFeedTransalation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFeedTransalation) ProtoMessage() {}

func (x *CreateFeedTransalation) ProtoReflect() protoreflect.Message {
	mi := &file_feeds_models_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFeedTransalation.ProtoReflect.Descriptor instead.
func (*CreateFeedTransalation) Descriptor() ([]byte, []int) {
	return file_feeds_models_proto_rawDescGZIP(), []int{9}
}

func (x *CreateFeedTransalation) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateFeedTransalation) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

func (x *CreateFeedTransalation) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type FeedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FeedRequest) Reset() {
	*x = FeedRequest{}
	mi := &file_feeds_models_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FeedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedRequest) ProtoMessage() {}

func (x *FeedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_feeds_models_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedRequest.ProtoReflect.Descriptor instead.
func (*FeedRequest) Descriptor() ([]byte, []int) {
	return file_feeds_models_proto_rawDescGZIP(), []int{10}
}

func (x *FeedRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UpdateFeedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	LogoUrl     string              `protobuf:"bytes,2,opt,name=logo_url,json=logoUrl,proto3" json:"logo_url,omitempty"`
	LogoUrlId   string              `protobuf:"bytes,3,opt,name=logo_url_id,json=logoUrlId,proto3" json:"logo_url_id,omitempty"`
	Priority    int32               `protobuf:"varint,4,opt,name=priority,proto3" json:"priority,omitempty"`
	MaxItems    int32               `protobuf:"varint,5,opt,name=max_items,json=maxItems,proto3" json:"max_items,omitempty"`
	BaseUrl     string              `protobuf:"bytes,6,opt,name=base_url,json=baseUrl,proto3" json:"base_url,omitempty"`
	Translation []*FeedTransalation `protobuf:"bytes,7,rep,name=translation,proto3" json:"translation,omitempty"`
}

func (x *UpdateFeedRequest) Reset() {
	*x = UpdateFeedRequest{}
	mi := &file_feeds_models_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateFeedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFeedRequest) ProtoMessage() {}

func (x *UpdateFeedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_feeds_models_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFeedRequest.ProtoReflect.Descriptor instead.
func (*UpdateFeedRequest) Descriptor() ([]byte, []int) {
	return file_feeds_models_proto_rawDescGZIP(), []int{11}
}

func (x *UpdateFeedRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateFeedRequest) GetLogoUrl() string {
	if x != nil {
		return x.LogoUrl
	}
	return ""
}

func (x *UpdateFeedRequest) GetLogoUrlId() string {
	if x != nil {
		return x.LogoUrlId
	}
	return ""
}

func (x *UpdateFeedRequest) GetPriority() int32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *UpdateFeedRequest) GetMaxItems() int32 {
	if x != nil {
		return x.MaxItems
	}
	return 0
}

func (x *UpdateFeedRequest) GetBaseUrl() string {
	if x != nil {
		return x.BaseUrl
	}
	return ""
}

func (x *UpdateFeedRequest) GetTranslation() []*FeedTransalation {
	if x != nil {
		return x.Translation
	}
	return nil
}

var File_feeds_models_proto protoreflect.FileDescriptor

var file_feeds_models_proto_rawDesc = []byte{
	0x0a, 0x12, 0x66, 0x65, 0x65, 0x64, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x22, 0x35, 0x0a, 0x0b,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6c,
	0x61, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x8b, 0x01, 0x0a, 0x0c, 0x46, 0x65, 0x65, 0x64, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x63, 0x6f, 0x6e, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x63, 0x6f, 0x6e, 0x55, 0x72, 0x6c, 0x12,
	0x17, 0x0a, 0x07, 0x69, 0x63, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x69, 0x63, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x37, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x22, 0x9b, 0x01, 0x0a, 0x08, 0x46, 0x65, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x17,
	0x0a, 0x07, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x66, 0x65, 0x65, 0x64, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x41, 0x74, 0x22,
	0xf1, 0x01, 0x0a, 0x10, 0x46, 0x65, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x66, 0x65, 0x65, 0x64, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55,
	0x72, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0x94, 0x01, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x22, 0x9f, 0x02, 0x0a, 0x04, 0x46,
	0x65, 0x65, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x6f, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f, 0x67, 0x6f, 0x55, 0x72, 0x6c, 0x12, 0x1e,
	0x0a, 0x0b, 0x6c, 0x6f, 0x67, 0x6f, 0x5f, 0x75, 0x72, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x6f, 0x67, 0x6f, 0x55, 0x72, 0x6c, 0x49, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61,
	0x78, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6d,
	0x61, 0x78, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x61, 0x73, 0x65, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x61, 0x73, 0x65, 0x55,
	0x72, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x3a, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x46,
	0x65, 0x65, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x87, 0x01, 0x0a,
	0x10, 0x46, 0x65, 0x65, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x66, 0x65, 0x65, 0x64, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6c, 0x61, 0x6e, 0x67, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x7f, 0x0a, 0x0b, 0x46, 0x65, 0x65, 0x64, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x66, 0x65, 0x65, 0x64, 0x49, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c,
	0x61, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x22, 0xe4, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a,
	0x08, 0x6c, 0x6f, 0x67, 0x6f, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6c, 0x6f, 0x67, 0x6f, 0x55, 0x72, 0x6c, 0x12, 0x1e, 0x0a, 0x0b, 0x6c, 0x6f, 0x67, 0x6f,
	0x5f, 0x75, 0x72, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c,
	0x6f, 0x67, 0x6f, 0x55, 0x72, 0x6c, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f,
	0x72, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f,
	0x72, 0x69, 0x74, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x49, 0x74, 0x65, 0x6d,
	0x73, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x61, 0x73, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x40, 0x0a, 0x0b,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x46, 0x65, 0x65, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x64,
	0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x65, 0x65, 0x64, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x61,
	0x6e, 0x67, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x1d, 0x0a, 0x0b, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x22, 0xee, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x65,
	0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x6f, 0x67,
	0x6f, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f, 0x67,
	0x6f, 0x55, 0x72, 0x6c, 0x12, 0x1e, 0x0a, 0x0b, 0x6c, 0x6f, 0x67, 0x6f, 0x5f, 0x75, 0x72, 0x6c,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x6f, 0x67, 0x6f, 0x55,
	0x72, 0x6c, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79,
	0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x19, 0x0a,
	0x08, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x62, 0x61, 0x73, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x3a, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x42, 0x10, 0x5a, 0x0e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x66, 0x65, 0x65, 0x64, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_feeds_models_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_feeds_models_proto_goTypes = []any{
	(*Translation)(nil),            // 0: models.Translation
	(*FeedCategory)(nil),           // 1: models.FeedCategory
	(*FeedItem)(nil),               // 2: models.FeedItem
	(*FeedItemResponse)(nil),       // 3: models.FeedItemResponse
	(*UpdateItem)(nil),             // 4: models.UpdateItem
	(*Feed)(nil),                   // 5: models.Feed
	(*FeedTransalation)(nil),       // 6: models.FeedTransalation
	(*FeedContent)(nil),            // 7: models.FeedContent
	(*CreateFeedRequest)(nil),      // 8: models.CreateFeedRequest
	(*CreateFeedTransalation)(nil), // 9: models.CreateFeedTransalation
	(*FeedRequest)(nil),            // 10: models.FeedRequest
	(*UpdateFeedRequest)(nil),      // 11: models.UpdateFeedRequest
}
var file_feeds_models_proto_depIdxs = []int32{
	0, // 0: models.FeedCategory.translations:type_name -> models.Translation
	6, // 1: models.Feed.translation:type_name -> models.FeedTransalation
	9, // 2: models.CreateFeedRequest.translation:type_name -> models.CreateFeedTransalation
	6, // 3: models.UpdateFeedRequest.translation:type_name -> models.FeedTransalation
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
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
			NumMessages:   12,
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
