// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.12.4
// source: searching/searching.proto

package searching

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

// Request for searching videos
type VideoSearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query         string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`                                        // Search query
	MaxResults    int32  `protobuf:"varint,2,opt,name=max_results,json=maxResults,proto3" json:"max_results,omitempty"`           // Maximum results per request
	NextPageToken string `protobuf:"bytes,3,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"` // Token for pagination
}

func (x *VideoSearchRequest) Reset() {
	*x = VideoSearchRequest{}
	mi := &file_searching_searching_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VideoSearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoSearchRequest) ProtoMessage() {}

func (x *VideoSearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_searching_searching_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoSearchRequest.ProtoReflect.Descriptor instead.
func (*VideoSearchRequest) Descriptor() ([]byte, []int) {
	return file_searching_searching_proto_rawDescGZIP(), []int{0}
}

func (x *VideoSearchRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *VideoSearchRequest) GetMaxResults() int32 {
	if x != nil {
		return x.MaxResults
	}
	return 0
}

func (x *VideoSearchRequest) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

// Generic query request
type QueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query      string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`                              // Search query
	PageNumber int32  `protobuf:"varint,2,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"` // Page number for pagination
	PageSize   int32  `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`       // Number of items per page
}

func (x *QueryRequest) Reset() {
	*x = QueryRequest{}
	mi := &file_searching_searching_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *QueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRequest) ProtoMessage() {}

func (x *QueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_searching_searching_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRequest.ProtoReflect.Descriptor instead.
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return file_searching_searching_proto_rawDescGZIP(), []int{1}
}

func (x *QueryRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *QueryRequest) GetPageNumber() int32 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *QueryRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

// Generic search response
type SearchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalItems int32     `protobuf:"varint,2,opt,name=total_items,json=totalItems,proto3" json:"total_items,omitempty"`
	Results    []*Result `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"` // List of generic search results
}

func (x *SearchResponse) Reset() {
	*x = SearchResponse{}
	mi := &file_searching_searching_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResponse) ProtoMessage() {}

func (x *SearchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_searching_searching_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchResponse.ProtoReflect.Descriptor instead.
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return file_searching_searching_proto_rawDescGZIP(), []int{2}
}

func (x *SearchResponse) GetTotalItems() int32 {
	if x != nil {
		return x.TotalItems
	}
	return 0
}

func (x *SearchResponse) GetResults() []*Result {
	if x != nil {
		return x.Results
	}
	return nil
}

// Response for video search
type VideoSearchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalItems    int32          `protobuf:"varint,3,opt,name=total_items,json=totalItems,proto3" json:"total_items,omitempty"`
	Videos        []*VideoResult `protobuf:"bytes,1,rep,name=videos,proto3" json:"videos,omitempty"`                                      // List of video results
	NextPageToken string         `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"` // Token for next page
}

func (x *VideoSearchResponse) Reset() {
	*x = VideoSearchResponse{}
	mi := &file_searching_searching_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VideoSearchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoSearchResponse) ProtoMessage() {}

func (x *VideoSearchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_searching_searching_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoSearchResponse.ProtoReflect.Descriptor instead.
func (*VideoSearchResponse) Descriptor() ([]byte, []int) {
	return file_searching_searching_proto_rawDescGZIP(), []int{3}
}

func (x *VideoSearchResponse) GetTotalItems() int32 {
	if x != nil {
		return x.TotalItems
	}
	return 0
}

func (x *VideoSearchResponse) GetVideos() []*VideoResult {
	if x != nil {
		return x.Videos
	}
	return nil
}

func (x *VideoSearchResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

// Response for image search
type ImageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalItems int32          `protobuf:"varint,2,opt,name=total_items,json=totalItems,proto3" json:"total_items,omitempty"`
	Images     []*ImageResult `protobuf:"bytes,1,rep,name=images,proto3" json:"images,omitempty"` // List of image results
}

func (x *ImageResponse) Reset() {
	*x = ImageResponse{}
	mi := &file_searching_searching_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ImageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageResponse) ProtoMessage() {}

func (x *ImageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_searching_searching_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageResponse.ProtoReflect.Descriptor instead.
func (*ImageResponse) Descriptor() ([]byte, []int) {
	return file_searching_searching_proto_rawDescGZIP(), []int{4}
}

func (x *ImageResponse) GetTotalItems() int32 {
	if x != nil {
		return x.TotalItems
	}
	return 0
}

func (x *ImageResponse) GetImages() []*ImageResult {
	if x != nil {
		return x.Images
	}
	return nil
}

// Structure for generic result
type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title           string             `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`                                              // Title of the result
	Content         string             `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`                                          // Description/content snippet
	DirectLink      string             `protobuf:"bytes,3,opt,name=direct_link,json=directLink,proto3" json:"direct_link,omitempty"`                  // Direct link to the result
	PrimaryImageUrl string             `protobuf:"bytes,4,opt,name=primary_image_url,json=primaryImageUrl,proto3" json:"primary_image_url,omitempty"` // URL of the primary image
	FavIconUrl      string             `protobuf:"bytes,5,opt,name=fav_icon_url,json=favIconUrl,proto3" json:"fav_icon_url,omitempty"`                // URL of the favicon
	DisplayLink     string             `protobuf:"bytes,6,opt,name=display_link,json=displayLink,proto3" json:"display_link,omitempty"`               // Display link (e.g., domain)
	Thumbnails      []*SearchThumbnail `protobuf:"bytes,7,rep,name=thumbnails,proto3" json:"thumbnails,omitempty"`                                    // List of thumbnails
}

func (x *Result) Reset() {
	*x = Result{}
	mi := &file_searching_searching_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_searching_searching_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_searching_searching_proto_rawDescGZIP(), []int{5}
}

func (x *Result) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Result) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Result) GetDirectLink() string {
	if x != nil {
		return x.DirectLink
	}
	return ""
}

func (x *Result) GetPrimaryImageUrl() string {
	if x != nil {
		return x.PrimaryImageUrl
	}
	return ""
}

func (x *Result) GetFavIconUrl() string {
	if x != nil {
		return x.FavIconUrl
	}
	return ""
}

func (x *Result) GetDisplayLink() string {
	if x != nil {
		return x.DisplayLink
	}
	return ""
}

func (x *Result) GetThumbnails() []*SearchThumbnail {
	if x != nil {
		return x.Thumbnails
	}
	return nil
}

// Structure for video result
type VideoResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title        string           `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`                       // Video title
	VideoUrl     string           `protobuf:"bytes,2,opt,name=video_url,json=videoUrl,proto3" json:"video_url,omitempty"` // URL to the video
	Description  string           `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
	ThumbnailUrl string           `protobuf:"bytes,3,opt,name=thumbnail_url,json=thumbnailUrl,proto3" json:"thumbnail_url,omitempty"` // Thumbnail URL
	ChannelTitle string           `protobuf:"bytes,4,opt,name=channel_title,json=channelTitle,proto3" json:"channel_title,omitempty"` // Channel name
	ChannelUrl   string           `protobuf:"bytes,5,opt,name=channel_url,json=channelUrl,proto3" json:"channel_url,omitempty"`       // Channel URL
	PublishTime  string           `protobuf:"bytes,6,opt,name=publish_time,json=publishTime,proto3" json:"publish_time,omitempty"`    // Publish time of the video
	Thumbnails   *VideoThumbnails `protobuf:"bytes,7,opt,name=thumbnails,proto3" json:"thumbnails,omitempty"`                         // Thumbnails of the video
}

func (x *VideoResult) Reset() {
	*x = VideoResult{}
	mi := &file_searching_searching_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VideoResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoResult) ProtoMessage() {}

func (x *VideoResult) ProtoReflect() protoreflect.Message {
	mi := &file_searching_searching_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoResult.ProtoReflect.Descriptor instead.
func (*VideoResult) Descriptor() ([]byte, []int) {
	return file_searching_searching_proto_rawDescGZIP(), []int{6}
}

func (x *VideoResult) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *VideoResult) GetVideoUrl() string {
	if x != nil {
		return x.VideoUrl
	}
	return ""
}

func (x *VideoResult) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *VideoResult) GetThumbnailUrl() string {
	if x != nil {
		return x.ThumbnailUrl
	}
	return ""
}

func (x *VideoResult) GetChannelTitle() string {
	if x != nil {
		return x.ChannelTitle
	}
	return ""
}

func (x *VideoResult) GetChannelUrl() string {
	if x != nil {
		return x.ChannelUrl
	}
	return ""
}

func (x *VideoResult) GetPublishTime() string {
	if x != nil {
		return x.PublishTime
	}
	return ""
}

func (x *VideoResult) GetThumbnails() *VideoThumbnails {
	if x != nil {
		return x.Thumbnails
	}
	return nil
}

// Structure for video thumbnails
type VideoThumbnails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Default *SearchThumbnail `protobuf:"bytes,1,opt,name=default,proto3" json:"default,omitempty"` // Default quality thumbnail
	Medium  *SearchThumbnail `protobuf:"bytes,2,opt,name=medium,proto3" json:"medium,omitempty"`   // Medium quality thumbnail
	High    *SearchThumbnail `protobuf:"bytes,3,opt,name=high,proto3" json:"high,omitempty"`       // High quality thumbnail
}

func (x *VideoThumbnails) Reset() {
	*x = VideoThumbnails{}
	mi := &file_searching_searching_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VideoThumbnails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoThumbnails) ProtoMessage() {}

func (x *VideoThumbnails) ProtoReflect() protoreflect.Message {
	mi := &file_searching_searching_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoThumbnails.ProtoReflect.Descriptor instead.
func (*VideoThumbnails) Descriptor() ([]byte, []int) {
	return file_searching_searching_proto_rawDescGZIP(), []int{7}
}

func (x *VideoThumbnails) GetDefault() *SearchThumbnail {
	if x != nil {
		return x.Default
	}
	return nil
}

func (x *VideoThumbnails) GetMedium() *SearchThumbnail {
	if x != nil {
		return x.Medium
	}
	return nil
}

func (x *VideoThumbnails) GetHigh() *SearchThumbnail {
	if x != nil {
		return x.High
	}
	return nil
}

// Structure for image result
type ImageResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title    string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`                       // Title of the image
	ImageUrl string `protobuf:"bytes,2,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"` // URL of the image
	Width    int32  `protobuf:"varint,3,opt,name=width,proto3" json:"width,omitempty"`                      // Image width
	Height   int32  `protobuf:"varint,4,opt,name=height,proto3" json:"height,omitempty"`                    // Image height
}

func (x *ImageResult) Reset() {
	*x = ImageResult{}
	mi := &file_searching_searching_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ImageResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageResult) ProtoMessage() {}

func (x *ImageResult) ProtoReflect() protoreflect.Message {
	mi := &file_searching_searching_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageResult.ProtoReflect.Descriptor instead.
func (*ImageResult) Descriptor() ([]byte, []int) {
	return file_searching_searching_proto_rawDescGZIP(), []int{8}
}

func (x *ImageResult) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ImageResult) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *ImageResult) GetWidth() int32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *ImageResult) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

// Structure for a single thumbnail
type SearchThumbnail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Src    string `protobuf:"bytes,1,opt,name=src,proto3" json:"src,omitempty"`        // Thumbnail URL
	Width  int64  `protobuf:"varint,2,opt,name=width,proto3" json:"width,omitempty"`   // Thumbnail width
	Height int64  `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"` // Thumbnail height
}

func (x *SearchThumbnail) Reset() {
	*x = SearchThumbnail{}
	mi := &file_searching_searching_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchThumbnail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchThumbnail) ProtoMessage() {}

func (x *SearchThumbnail) ProtoReflect() protoreflect.Message {
	mi := &file_searching_searching_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchThumbnail.ProtoReflect.Descriptor instead.
func (*SearchThumbnail) Descriptor() ([]byte, []int) {
	return file_searching_searching_proto_rawDescGZIP(), []int{9}
}

func (x *SearchThumbnail) GetSrc() string {
	if x != nil {
		return x.Src
	}
	return ""
}

func (x *SearchThumbnail) GetWidth() int64 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *SearchThumbnail) GetHeight() int64 {
	if x != nil {
		return x.Height
	}
	return 0
}

var File_searching_searching_proto protoreflect.FileDescriptor

var file_searching_searching_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x2f, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x73,
	0x0a, 0x12, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x61,
	0x78, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x6d, 0x61, 0x78, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e,
	0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0x62, 0x0a, 0x0c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70,
	0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x66, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x33, 0x0a, 0x07, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x22,
	0x96, 0x01, 0x0a, 0x13, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x36, 0x0a, 0x06, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x73,
	0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50,
	0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x68, 0x0a, 0x0d, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x36, 0x0a, 0x06, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x73, 0x22, 0x8e, 0x02, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x0a,
	0x0b, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x2a,
	0x0a, 0x11, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x72, 0x69, 0x6d, 0x61,
	0x72, 0x79, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x20, 0x0a, 0x0c, 0x66, 0x61,
	0x76, 0x5f, 0x69, 0x63, 0x6f, 0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x66, 0x61, 0x76, 0x49, 0x63, 0x6f, 0x6e, 0x55, 0x72, 0x6c, 0x12, 0x21, 0x0a, 0x0c,
	0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4c, 0x69, 0x6e, 0x6b, 0x12,
	0x42, 0x0a, 0x0a, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x07, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x68,
	0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x52, 0x0a, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61,
	0x69, 0x6c, 0x73, 0x22, 0xb4, 0x02, 0x0a, 0x0b, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x55, 0x72, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x68, 0x75, 0x6d,
	0x62, 0x6e, 0x61, 0x69, 0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x55, 0x72, 0x6c, 0x12, 0x23, 0x0a,
	0x0d, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x54, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x55, 0x72, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x42, 0x0a, 0x0a, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e,
	0x61, 0x69, 0x6c, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x0a,
	0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x73, 0x22, 0xc3, 0x01, 0x0a, 0x0f, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x3c,
	0x0a, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e,
	0x61, 0x69, 0x6c, 0x52, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x3a, 0x0a, 0x06,
	0x6d, 0x65, 0x64, 0x69, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c,
	0x52, 0x06, 0x6d, 0x65, 0x64, 0x69, 0x75, 0x6d, 0x12, 0x36, 0x0a, 0x04, 0x68, 0x69, 0x67, 0x68,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x69,
	0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x52, 0x04, 0x68, 0x69, 0x67, 0x68,
	0x22, 0x6e, 0x0a, 0x0b, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55,
	0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x22, 0x51, 0x0a, 0x0f, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e,
	0x61, 0x69, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x72, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x73, 0x72, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x68, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x32, 0x92, 0x02, 0x0a, 0x10, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x69, 0x6e,
	0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x12, 0x1f, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5d, 0x0a, 0x0c, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x56, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x12, 0x25, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x69,
	0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x0c, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x1f, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x69, 0x6e,
	0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x69,
	0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x14, 0x5a, 0x12, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_searching_searching_proto_rawDescOnce sync.Once
	file_searching_searching_proto_rawDescData = file_searching_searching_proto_rawDesc
)

func file_searching_searching_proto_rawDescGZIP() []byte {
	file_searching_searching_proto_rawDescOnce.Do(func() {
		file_searching_searching_proto_rawDescData = protoimpl.X.CompressGZIP(file_searching_searching_proto_rawDescData)
	})
	return file_searching_searching_proto_rawDescData
}

var file_searching_searching_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_searching_searching_proto_goTypes = []any{
	(*VideoSearchRequest)(nil),  // 0: searching_service.VideoSearchRequest
	(*QueryRequest)(nil),        // 1: searching_service.QueryRequest
	(*SearchResponse)(nil),      // 2: searching_service.SearchResponse
	(*VideoSearchResponse)(nil), // 3: searching_service.VideoSearchResponse
	(*ImageResponse)(nil),       // 4: searching_service.ImageResponse
	(*Result)(nil),              // 5: searching_service.Result
	(*VideoResult)(nil),         // 6: searching_service.VideoResult
	(*VideoThumbnails)(nil),     // 7: searching_service.VideoThumbnails
	(*ImageResult)(nil),         // 8: searching_service.ImageResult
	(*SearchThumbnail)(nil),     // 9: searching_service.SearchThumbnail
}
var file_searching_searching_proto_depIdxs = []int32{
	5,  // 0: searching_service.SearchResponse.results:type_name -> searching_service.Result
	6,  // 1: searching_service.VideoSearchResponse.videos:type_name -> searching_service.VideoResult
	8,  // 2: searching_service.ImageResponse.images:type_name -> searching_service.ImageResult
	9,  // 3: searching_service.Result.thumbnails:type_name -> searching_service.SearchThumbnail
	7,  // 4: searching_service.VideoResult.thumbnails:type_name -> searching_service.VideoThumbnails
	9,  // 5: searching_service.VideoThumbnails.default:type_name -> searching_service.SearchThumbnail
	9,  // 6: searching_service.VideoThumbnails.medium:type_name -> searching_service.SearchThumbnail
	9,  // 7: searching_service.VideoThumbnails.high:type_name -> searching_service.SearchThumbnail
	1,  // 8: searching_service.SearchingService.Search:input_type -> searching_service.QueryRequest
	0,  // 9: searching_service.SearchingService.SearchVideos:input_type -> searching_service.VideoSearchRequest
	1,  // 10: searching_service.SearchingService.SearchImages:input_type -> searching_service.QueryRequest
	2,  // 11: searching_service.SearchingService.Search:output_type -> searching_service.SearchResponse
	3,  // 12: searching_service.SearchingService.SearchVideos:output_type -> searching_service.VideoSearchResponse
	4,  // 13: searching_service.SearchingService.SearchImages:output_type -> searching_service.ImageResponse
	11, // [11:14] is the sub-list for method output_type
	8,  // [8:11] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_searching_searching_proto_init() }
func file_searching_searching_proto_init() {
	if File_searching_searching_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_searching_searching_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_searching_searching_proto_goTypes,
		DependencyIndexes: file_searching_searching_proto_depIdxs,
		MessageInfos:      file_searching_searching_proto_msgTypes,
	}.Build()
	File_searching_searching_proto = out.File
	file_searching_searching_proto_rawDesc = nil
	file_searching_searching_proto_goTypes = nil
	file_searching_searching_proto_depIdxs = nil
}
