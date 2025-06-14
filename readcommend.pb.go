// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: readcommend.proto

package readcommend

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Book struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	YearPublished uint32                 `protobuf:"varint,3,opt,name=yearPublished,proto3" json:"yearPublished,omitempty"`
	Rating        float64                `protobuf:"fixed64,4,opt,name=rating,proto3" json:"rating,omitempty"`
	Pages         uint32                 `protobuf:"varint,5,opt,name=pages,proto3" json:"pages,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Book) Reset() {
	*x = Book{}
	mi := &file_readcommend_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_readcommend_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book.ProtoReflect.Descriptor instead.
func (*Book) Descriptor() ([]byte, []int) {
	return file_readcommend_proto_rawDescGZIP(), []int{0}
}

func (x *Book) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Book) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Book) GetYearPublished() uint32 {
	if x != nil {
		return x.YearPublished
	}
	return 0
}

func (x *Book) GetRating() float64 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *Book) GetPages() uint32 {
	if x != nil {
		return x.Pages
	}
	return 0
}

type Author struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName     string                 `protobuf:"bytes,2,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName      string                 `protobuf:"bytes,3,opt,name=lastName,proto3" json:"lastName,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Author) Reset() {
	*x = Author{}
	mi := &file_readcommend_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Author) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Author) ProtoMessage() {}

func (x *Author) ProtoReflect() protoreflect.Message {
	mi := &file_readcommend_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Author.ProtoReflect.Descriptor instead.
func (*Author) Descriptor() ([]byte, []int) {
	return file_readcommend_proto_rawDescGZIP(), []int{1}
}

func (x *Author) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Author) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Author) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

type Genre struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Genre) Reset() {
	*x = Genre{}
	mi := &file_readcommend_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Genre) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Genre) ProtoMessage() {}

func (x *Genre) ProtoReflect() protoreflect.Message {
	mi := &file_readcommend_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Genre.ProtoReflect.Descriptor instead.
func (*Genre) Descriptor() ([]byte, []int) {
	return file_readcommend_proto_rawDescGZIP(), []int{2}
}

func (x *Genre) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Genre) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type Era struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	MinYear       uint32                 `protobuf:"varint,3,opt,name=minYear,proto3" json:"minYear,omitempty"`
	MaxYear       uint32                 `protobuf:"varint,4,opt,name=maxYear,proto3" json:"maxYear,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Era) Reset() {
	*x = Era{}
	mi := &file_readcommend_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Era) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Era) ProtoMessage() {}

func (x *Era) ProtoReflect() protoreflect.Message {
	mi := &file_readcommend_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Era.ProtoReflect.Descriptor instead.
func (*Era) Descriptor() ([]byte, []int) {
	return file_readcommend_proto_rawDescGZIP(), []int{3}
}

func (x *Era) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Era) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Era) GetMinYear() uint32 {
	if x != nil {
		return x.MinYear
	}
	return 0
}

func (x *Era) GetMaxYear() uint32 {
	if x != nil {
		return x.MaxYear
	}
	return 0
}

type Size struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	MinPages      uint32                 `protobuf:"varint,3,opt,name=minPages,proto3" json:"minPages,omitempty"`
	MaxPages      uint32                 `protobuf:"varint,4,opt,name=maxPages,proto3" json:"maxPages,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Size) Reset() {
	*x = Size{}
	mi := &file_readcommend_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Size) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Size) ProtoMessage() {}

func (x *Size) ProtoReflect() protoreflect.Message {
	mi := &file_readcommend_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Size.ProtoReflect.Descriptor instead.
func (*Size) Descriptor() ([]byte, []int) {
	return file_readcommend_proto_rawDescGZIP(), []int{4}
}

func (x *Size) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Size) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Size) GetMinPages() uint32 {
	if x != nil {
		return x.MinPages
	}
	return 0
}

func (x *Size) GetMaxPages() uint32 {
	if x != nil {
		return x.MaxPages
	}
	return 0
}

type GetGenresParams struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetGenresParams) Reset() {
	*x = GetGenresParams{}
	mi := &file_readcommend_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetGenresParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGenresParams) ProtoMessage() {}

func (x *GetGenresParams) ProtoReflect() protoreflect.Message {
	mi := &file_readcommend_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGenresParams.ProtoReflect.Descriptor instead.
func (*GetGenresParams) Descriptor() ([]byte, []int) {
	return file_readcommend_proto_rawDescGZIP(), []int{5}
}

type GetGenresResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Genre         []*Genre               `protobuf:"bytes,1,rep,name=genre,proto3" json:"genre,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetGenresResponse) Reset() {
	*x = GetGenresResponse{}
	mi := &file_readcommend_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetGenresResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGenresResponse) ProtoMessage() {}

func (x *GetGenresResponse) ProtoReflect() protoreflect.Message {
	mi := &file_readcommend_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGenresResponse.ProtoReflect.Descriptor instead.
func (*GetGenresResponse) Descriptor() ([]byte, []int) {
	return file_readcommend_proto_rawDescGZIP(), []int{6}
}

func (x *GetGenresResponse) GetGenre() []*Genre {
	if x != nil {
		return x.Genre
	}
	return nil
}

type GetErasParams struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetErasParams) Reset() {
	*x = GetErasParams{}
	mi := &file_readcommend_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetErasParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetErasParams) ProtoMessage() {}

func (x *GetErasParams) ProtoReflect() protoreflect.Message {
	mi := &file_readcommend_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetErasParams.ProtoReflect.Descriptor instead.
func (*GetErasParams) Descriptor() ([]byte, []int) {
	return file_readcommend_proto_rawDescGZIP(), []int{7}
}

type GetErasResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Eras          []*Era                 `protobuf:"bytes,1,rep,name=eras,proto3" json:"eras,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetErasResponse) Reset() {
	*x = GetErasResponse{}
	mi := &file_readcommend_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetErasResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetErasResponse) ProtoMessage() {}

func (x *GetErasResponse) ProtoReflect() protoreflect.Message {
	mi := &file_readcommend_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetErasResponse.ProtoReflect.Descriptor instead.
func (*GetErasResponse) Descriptor() ([]byte, []int) {
	return file_readcommend_proto_rawDescGZIP(), []int{8}
}

func (x *GetErasResponse) GetEras() []*Era {
	if x != nil {
		return x.Eras
	}
	return nil
}

type GetSizesParams struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetSizesParams) Reset() {
	*x = GetSizesParams{}
	mi := &file_readcommend_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSizesParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSizesParams) ProtoMessage() {}

func (x *GetSizesParams) ProtoReflect() protoreflect.Message {
	mi := &file_readcommend_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSizesParams.ProtoReflect.Descriptor instead.
func (*GetSizesParams) Descriptor() ([]byte, []int) {
	return file_readcommend_proto_rawDescGZIP(), []int{9}
}

type GetSizesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Sizes         []*Size                `protobuf:"bytes,1,rep,name=sizes,proto3" json:"sizes,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetSizesResponse) Reset() {
	*x = GetSizesResponse{}
	mi := &file_readcommend_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSizesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSizesResponse) ProtoMessage() {}

func (x *GetSizesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_readcommend_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSizesResponse.ProtoReflect.Descriptor instead.
func (*GetSizesResponse) Descriptor() ([]byte, []int) {
	return file_readcommend_proto_rawDescGZIP(), []int{10}
}

func (x *GetSizesResponse) GetSizes() []*Size {
	if x != nil {
		return x.Sizes
	}
	return nil
}

type GetAuthorsParams struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAuthorsParams) Reset() {
	*x = GetAuthorsParams{}
	mi := &file_readcommend_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAuthorsParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthorsParams) ProtoMessage() {}

func (x *GetAuthorsParams) ProtoReflect() protoreflect.Message {
	mi := &file_readcommend_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthorsParams.ProtoReflect.Descriptor instead.
func (*GetAuthorsParams) Descriptor() ([]byte, []int) {
	return file_readcommend_proto_rawDescGZIP(), []int{11}
}

type GetAuthorsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Authors       []*Author              `protobuf:"bytes,1,rep,name=authors,proto3" json:"authors,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAuthorsResponse) Reset() {
	*x = GetAuthorsResponse{}
	mi := &file_readcommend_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAuthorsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthorsResponse) ProtoMessage() {}

func (x *GetAuthorsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_readcommend_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthorsResponse.ProtoReflect.Descriptor instead.
func (*GetAuthorsResponse) Descriptor() ([]byte, []int) {
	return file_readcommend_proto_rawDescGZIP(), []int{12}
}

func (x *GetAuthorsResponse) GetAuthors() []*Author {
	if x != nil {
		return x.Authors
	}
	return nil
}

type GetBooksParams struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Authors       *string                `protobuf:"bytes,1,opt,name=authors,proto3,oneof" json:"authors,omitempty"`
	Genres        *string                `protobuf:"bytes,2,opt,name=genres,proto3,oneof" json:"genres,omitempty"`
	MinPages      *uint32                `protobuf:"varint,3,opt,name=minPages,json=min-pages,proto3,oneof" json:"minPages,omitempty"`
	MaxPages      *uint32                `protobuf:"varint,4,opt,name=maxPages,json=max-pages,proto3,oneof" json:"maxPages,omitempty"`
	MinYear       *uint32                `protobuf:"varint,5,opt,name=minYear,json=min-year,proto3,oneof" json:"minYear,omitempty"`
	MaxYear       *uint32                `protobuf:"varint,6,opt,name=maxYear,json=max-year,proto3,oneof" json:"maxYear,omitempty"`
	Limit         *uint32                `protobuf:"varint,7,opt,name=limit,proto3,oneof" json:"limit,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetBooksParams) Reset() {
	*x = GetBooksParams{}
	mi := &file_readcommend_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBooksParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBooksParams) ProtoMessage() {}

func (x *GetBooksParams) ProtoReflect() protoreflect.Message {
	mi := &file_readcommend_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBooksParams.ProtoReflect.Descriptor instead.
func (*GetBooksParams) Descriptor() ([]byte, []int) {
	return file_readcommend_proto_rawDescGZIP(), []int{13}
}

func (x *GetBooksParams) GetAuthors() string {
	if x != nil && x.Authors != nil {
		return *x.Authors
	}
	return ""
}

func (x *GetBooksParams) GetGenres() string {
	if x != nil && x.Genres != nil {
		return *x.Genres
	}
	return ""
}

func (x *GetBooksParams) GetMinPages() uint32 {
	if x != nil && x.MinPages != nil {
		return *x.MinPages
	}
	return 0
}

func (x *GetBooksParams) GetMaxPages() uint32 {
	if x != nil && x.MaxPages != nil {
		return *x.MaxPages
	}
	return 0
}

func (x *GetBooksParams) GetMinYear() uint32 {
	if x != nil && x.MinYear != nil {
		return *x.MinYear
	}
	return 0
}

func (x *GetBooksParams) GetMaxYear() uint32 {
	if x != nil && x.MaxYear != nil {
		return *x.MaxYear
	}
	return 0
}

func (x *GetBooksParams) GetLimit() uint32 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

type GetBooksResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Books         []*Book                `protobuf:"bytes,1,rep,name=books,proto3" json:"books,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetBooksResponse) Reset() {
	*x = GetBooksResponse{}
	mi := &file_readcommend_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBooksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBooksResponse) ProtoMessage() {}

func (x *GetBooksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_readcommend_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBooksResponse.ProtoReflect.Descriptor instead.
func (*GetBooksResponse) Descriptor() ([]byte, []int) {
	return file_readcommend_proto_rawDescGZIP(), []int{14}
}

func (x *GetBooksResponse) GetBooks() []*Book {
	if x != nil {
		return x.Books
	}
	return nil
}

var File_readcommend_proto protoreflect.FileDescriptor

const file_readcommend_proto_rawDesc = "" +
	"\n" +
	"\x11readcommend.proto\x12\vreadcommend\x1a\x1cgoogle/api/annotations.proto\"\x80\x01\n" +
	"\x04Book\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\rR\x02id\x12\x14\n" +
	"\x05title\x18\x02 \x01(\tR\x05title\x12$\n" +
	"\ryearPublished\x18\x03 \x01(\rR\ryearPublished\x12\x16\n" +
	"\x06rating\x18\x04 \x01(\x01R\x06rating\x12\x14\n" +
	"\x05pages\x18\x05 \x01(\rR\x05pages\"R\n" +
	"\x06Author\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\rR\x02id\x12\x1c\n" +
	"\tfirstName\x18\x02 \x01(\tR\tfirstName\x12\x1a\n" +
	"\blastName\x18\x03 \x01(\tR\blastName\"-\n" +
	"\x05Genre\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\rR\x02id\x12\x14\n" +
	"\x05title\x18\x02 \x01(\tR\x05title\"_\n" +
	"\x03Era\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\rR\x02id\x12\x14\n" +
	"\x05title\x18\x02 \x01(\tR\x05title\x12\x18\n" +
	"\aminYear\x18\x03 \x01(\rR\aminYear\x12\x18\n" +
	"\amaxYear\x18\x04 \x01(\rR\amaxYear\"d\n" +
	"\x04Size\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\rR\x02id\x12\x14\n" +
	"\x05title\x18\x02 \x01(\tR\x05title\x12\x1a\n" +
	"\bminPages\x18\x03 \x01(\rR\bminPages\x12\x1a\n" +
	"\bmaxPages\x18\x04 \x01(\rR\bmaxPages\"\x11\n" +
	"\x0fGetGenresParams\"=\n" +
	"\x11GetGenresResponse\x12(\n" +
	"\x05genre\x18\x01 \x03(\v2\x12.readcommend.GenreR\x05genre\"\x0f\n" +
	"\rGetErasParams\"7\n" +
	"\x0fGetErasResponse\x12$\n" +
	"\x04eras\x18\x01 \x03(\v2\x10.readcommend.EraR\x04eras\"\x10\n" +
	"\x0eGetSizesParams\";\n" +
	"\x10GetSizesResponse\x12'\n" +
	"\x05sizes\x18\x01 \x03(\v2\x11.readcommend.SizeR\x05sizes\"\x12\n" +
	"\x10GetAuthorsParams\"C\n" +
	"\x12GetAuthorsResponse\x12-\n" +
	"\aauthors\x18\x01 \x03(\v2\x13.readcommend.AuthorR\aauthors\"\xbe\x02\n" +
	"\x0eGetBooksParams\x12\x1d\n" +
	"\aauthors\x18\x01 \x01(\tH\x00R\aauthors\x88\x01\x01\x12\x1b\n" +
	"\x06genres\x18\x02 \x01(\tH\x01R\x06genres\x88\x01\x01\x12 \n" +
	"\bminPages\x18\x03 \x01(\rH\x02R\tmin-pages\x88\x01\x01\x12 \n" +
	"\bmaxPages\x18\x04 \x01(\rH\x03R\tmax-pages\x88\x01\x01\x12\x1e\n" +
	"\aminYear\x18\x05 \x01(\rH\x04R\bmin-year\x88\x01\x01\x12\x1e\n" +
	"\amaxYear\x18\x06 \x01(\rH\x05R\bmax-year\x88\x01\x01\x12\x19\n" +
	"\x05limit\x18\a \x01(\rH\x06R\x05limit\x88\x01\x01B\n" +
	"\n" +
	"\b_authorsB\t\n" +
	"\a_genresB\v\n" +
	"\t_minPagesB\v\n" +
	"\t_maxPagesB\n" +
	"\n" +
	"\b_minYearB\n" +
	"\n" +
	"\b_maxYearB\b\n" +
	"\x06_limit\";\n" +
	"\x10GetBooksResponse\x12'\n" +
	"\x05books\x18\x01 \x03(\v2\x11.readcommend.BookR\x05books2\xf7\x03\n" +
	"\x12ReadCommendService\x12]\n" +
	"\bGetBooks\x12\x1b.readcommend.GetBooksParams\x1a\x1d.readcommend.GetBooksResponse\"\x15\x82\xd3\xe4\x93\x02\x0f\x12\r/api/v1/books\x12a\n" +
	"\tGetGenres\x12\x1c.readcommend.GetGenresParams\x1a\x1e.readcommend.GetGenresResponse\"\x16\x82\xd3\xe4\x93\x02\x10\x12\x0e/api/v1/genres\x12e\n" +
	"\n" +
	"GetAuthors\x12\x1d.readcommend.GetAuthorsParams\x1a\x1f.readcommend.GetAuthorsResponse\"\x17\x82\xd3\xe4\x93\x02\x11\x12\x0f/api/v1/authors\x12]\n" +
	"\bGetSizes\x12\x1b.readcommend.GetSizesParams\x1a\x1d.readcommend.GetSizesResponse\"\x15\x82\xd3\xe4\x93\x02\x0f\x12\r/api/v1/sizes\x12Y\n" +
	"\aGetEras\x12\x1a.readcommend.GetErasParams\x1a\x1c.readcommend.GetErasResponse\"\x14\x82\xd3\xe4\x93\x02\x0e\x12\f/api/v1/erasB'Z%github.com/BastianEscaich/readcommendb\x06proto3"

var (
	file_readcommend_proto_rawDescOnce sync.Once
	file_readcommend_proto_rawDescData []byte
)

func file_readcommend_proto_rawDescGZIP() []byte {
	file_readcommend_proto_rawDescOnce.Do(func() {
		file_readcommend_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_readcommend_proto_rawDesc), len(file_readcommend_proto_rawDesc)))
	})
	return file_readcommend_proto_rawDescData
}

var file_readcommend_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_readcommend_proto_goTypes = []any{
	(*Book)(nil),               // 0: readcommend.Book
	(*Author)(nil),             // 1: readcommend.Author
	(*Genre)(nil),              // 2: readcommend.Genre
	(*Era)(nil),                // 3: readcommend.Era
	(*Size)(nil),               // 4: readcommend.Size
	(*GetGenresParams)(nil),    // 5: readcommend.GetGenresParams
	(*GetGenresResponse)(nil),  // 6: readcommend.GetGenresResponse
	(*GetErasParams)(nil),      // 7: readcommend.GetErasParams
	(*GetErasResponse)(nil),    // 8: readcommend.GetErasResponse
	(*GetSizesParams)(nil),     // 9: readcommend.GetSizesParams
	(*GetSizesResponse)(nil),   // 10: readcommend.GetSizesResponse
	(*GetAuthorsParams)(nil),   // 11: readcommend.GetAuthorsParams
	(*GetAuthorsResponse)(nil), // 12: readcommend.GetAuthorsResponse
	(*GetBooksParams)(nil),     // 13: readcommend.GetBooksParams
	(*GetBooksResponse)(nil),   // 14: readcommend.GetBooksResponse
}
var file_readcommend_proto_depIdxs = []int32{
	2,  // 0: readcommend.GetGenresResponse.genre:type_name -> readcommend.Genre
	3,  // 1: readcommend.GetErasResponse.eras:type_name -> readcommend.Era
	4,  // 2: readcommend.GetSizesResponse.sizes:type_name -> readcommend.Size
	1,  // 3: readcommend.GetAuthorsResponse.authors:type_name -> readcommend.Author
	0,  // 4: readcommend.GetBooksResponse.books:type_name -> readcommend.Book
	13, // 5: readcommend.ReadCommendService.GetBooks:input_type -> readcommend.GetBooksParams
	5,  // 6: readcommend.ReadCommendService.GetGenres:input_type -> readcommend.GetGenresParams
	11, // 7: readcommend.ReadCommendService.GetAuthors:input_type -> readcommend.GetAuthorsParams
	9,  // 8: readcommend.ReadCommendService.GetSizes:input_type -> readcommend.GetSizesParams
	7,  // 9: readcommend.ReadCommendService.GetEras:input_type -> readcommend.GetErasParams
	14, // 10: readcommend.ReadCommendService.GetBooks:output_type -> readcommend.GetBooksResponse
	6,  // 11: readcommend.ReadCommendService.GetGenres:output_type -> readcommend.GetGenresResponse
	12, // 12: readcommend.ReadCommendService.GetAuthors:output_type -> readcommend.GetAuthorsResponse
	10, // 13: readcommend.ReadCommendService.GetSizes:output_type -> readcommend.GetSizesResponse
	8,  // 14: readcommend.ReadCommendService.GetEras:output_type -> readcommend.GetErasResponse
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_readcommend_proto_init() }
func file_readcommend_proto_init() {
	if File_readcommend_proto != nil {
		return
	}
	file_readcommend_proto_msgTypes[13].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_readcommend_proto_rawDesc), len(file_readcommend_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_readcommend_proto_goTypes,
		DependencyIndexes: file_readcommend_proto_depIdxs,
		MessageInfos:      file_readcommend_proto_msgTypes,
	}.Build()
	File_readcommend_proto = out.File
	file_readcommend_proto_goTypes = nil
	file_readcommend_proto_depIdxs = nil
}
