// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: favorites.proto

package favorites

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

type CreateFavoritesReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PropertyId string `protobuf:"bytes,2,opt,name=property_id,json=propertyId,proto3" json:"property_id,omitempty"`
}

func (x *CreateFavoritesReq) Reset() {
	*x = CreateFavoritesReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favorites_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFavoritesReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFavoritesReq) ProtoMessage() {}

func (x *CreateFavoritesReq) ProtoReflect() protoreflect.Message {
	mi := &file_favorites_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFavoritesReq.ProtoReflect.Descriptor instead.
func (*CreateFavoritesReq) Descriptor() ([]byte, []int) {
	return file_favorites_proto_rawDescGZIP(), []int{0}
}

func (x *CreateFavoritesReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateFavoritesReq) GetPropertyId() string {
	if x != nil {
		return x.PropertyId
	}
	return ""
}

type CreateFavoritesRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId     string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PropertyId string `protobuf:"bytes,3,opt,name=property_id,json=propertyId,proto3" json:"property_id,omitempty"`
	CreatedAt  string `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *CreateFavoritesRes) Reset() {
	*x = CreateFavoritesRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favorites_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFavoritesRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFavoritesRes) ProtoMessage() {}

func (x *CreateFavoritesRes) ProtoReflect() protoreflect.Message {
	mi := &file_favorites_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFavoritesRes.ProtoReflect.Descriptor instead.
func (*CreateFavoritesRes) Descriptor() ([]byte, []int) {
	return file_favorites_proto_rawDescGZIP(), []int{1}
}

func (x *CreateFavoritesRes) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateFavoritesRes) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateFavoritesRes) GetPropertyId() string {
	if x != nil {
		return x.PropertyId
	}
	return ""
}

func (x *CreateFavoritesRes) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type GetAllFavoritesReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit int32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Page  int32 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *GetAllFavoritesReq) Reset() {
	*x = GetAllFavoritesReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favorites_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllFavoritesReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllFavoritesReq) ProtoMessage() {}

func (x *GetAllFavoritesReq) ProtoReflect() protoreflect.Message {
	mi := &file_favorites_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllFavoritesReq.ProtoReflect.Descriptor instead.
func (*GetAllFavoritesReq) Descriptor() ([]byte, []int) {
	return file_favorites_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllFavoritesReq) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetAllFavoritesReq) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

type Favorite struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId     string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PropertyId string `protobuf:"bytes,3,opt,name=property_id,json=propertyId,proto3" json:"property_id,omitempty"`
	CreatedAt  string `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Favorite) Reset() {
	*x = Favorite{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favorites_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Favorite) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Favorite) ProtoMessage() {}

func (x *Favorite) ProtoReflect() protoreflect.Message {
	mi := &file_favorites_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Favorite.ProtoReflect.Descriptor instead.
func (*Favorite) Descriptor() ([]byte, []int) {
	return file_favorites_proto_rawDescGZIP(), []int{3}
}

func (x *Favorite) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Favorite) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Favorite) GetPropertyId() string {
	if x != nil {
		return x.PropertyId
	}
	return ""
}

func (x *Favorite) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type GetAllFavoritesRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Favorites []*Favorite `protobuf:"bytes,1,rep,name=favorites,proto3" json:"favorites,omitempty"`
}

func (x *GetAllFavoritesRes) Reset() {
	*x = GetAllFavoritesRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favorites_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllFavoritesRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllFavoritesRes) ProtoMessage() {}

func (x *GetAllFavoritesRes) ProtoReflect() protoreflect.Message {
	mi := &file_favorites_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllFavoritesRes.ProtoReflect.Descriptor instead.
func (*GetAllFavoritesRes) Descriptor() ([]byte, []int) {
	return file_favorites_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllFavoritesRes) GetFavorites() []*Favorite {
	if x != nil {
		return x.Favorites
	}
	return nil
}

type GetByIdFavoritesReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetByIdFavoritesReq) Reset() {
	*x = GetByIdFavoritesReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favorites_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByIdFavoritesReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByIdFavoritesReq) ProtoMessage() {}

func (x *GetByIdFavoritesReq) ProtoReflect() protoreflect.Message {
	mi := &file_favorites_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByIdFavoritesReq.ProtoReflect.Descriptor instead.
func (*GetByIdFavoritesReq) Descriptor() ([]byte, []int) {
	return file_favorites_proto_rawDescGZIP(), []int{5}
}

func (x *GetByIdFavoritesReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetByIdFavoritesRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId     string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PropertyId string `protobuf:"bytes,3,opt,name=property_id,json=propertyId,proto3" json:"property_id,omitempty"`
	CreatedAt  string `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *GetByIdFavoritesRes) Reset() {
	*x = GetByIdFavoritesRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favorites_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByIdFavoritesRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByIdFavoritesRes) ProtoMessage() {}

func (x *GetByIdFavoritesRes) ProtoReflect() protoreflect.Message {
	mi := &file_favorites_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByIdFavoritesRes.ProtoReflect.Descriptor instead.
func (*GetByIdFavoritesRes) Descriptor() ([]byte, []int) {
	return file_favorites_proto_rawDescGZIP(), []int{6}
}

func (x *GetByIdFavoritesRes) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetByIdFavoritesRes) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetByIdFavoritesRes) GetPropertyId() string {
	if x != nil {
		return x.PropertyId
	}
	return ""
}

func (x *GetByIdFavoritesRes) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type DeleteFavoritesReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteFavoritesReq) Reset() {
	*x = DeleteFavoritesReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favorites_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFavoritesReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFavoritesReq) ProtoMessage() {}

func (x *DeleteFavoritesReq) ProtoReflect() protoreflect.Message {
	mi := &file_favorites_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFavoritesReq.ProtoReflect.Descriptor instead.
func (*DeleteFavoritesReq) Descriptor() ([]byte, []int) {
	return file_favorites_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteFavoritesReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteFavoritesRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteFavoritesRes) Reset() {
	*x = DeleteFavoritesRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favorites_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFavoritesRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFavoritesRes) ProtoMessage() {}

func (x *DeleteFavoritesRes) ProtoReflect() protoreflect.Message {
	mi := &file_favorites_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFavoritesRes.ProtoReflect.Descriptor instead.
func (*DeleteFavoritesRes) Descriptor() ([]byte, []int) {
	return file_favorites_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteFavoritesRes) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_favorites_proto protoreflect.FileDescriptor

var file_favorites_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x73, 0x22, 0x4e, 0x0a, 0x12,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70,
	0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x49, 0x64, 0x22, 0x7d, 0x0a, 0x12,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70,
	0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x3e, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x73, 0x0a, 0x08, 0x46,
	0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x22, 0x47, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69,
	0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x12, 0x31, 0x0a, 0x09, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69,
	0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x66, 0x61, 0x76, 0x6f,
	0x72, 0x69, 0x74, 0x65, 0x73, 0x2e, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x52, 0x09,
	0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x73, 0x22, 0x25, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x42, 0x79, 0x49, 0x64, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x7e, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x46, 0x61, 0x76, 0x6f, 0x72,
	0x69, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x22, 0x24, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69,
	0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2e, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xd2, 0x02, 0x0a, 0x09, 0x46, 0x61, 0x76, 0x6f, 0x72,
	0x69, 0x74, 0x65, 0x73, 0x12, 0x4f, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x61,
	0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x73, 0x12, 0x1d, 0x2e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69,
	0x74, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69,
	0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74,
	0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x12, 0x4f, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x46,
	0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x73, 0x12, 0x1d, 0x2e, 0x66, 0x61, 0x76, 0x6f, 0x72,
	0x69, 0x74, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x46, 0x61, 0x76, 0x6f, 0x72,
	0x69, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69,
	0x74, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69,
	0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x12, 0x52, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49,
	0x64, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x73, 0x12, 0x1e, 0x2e, 0x66, 0x61, 0x76,
	0x6f, 0x72, 0x69, 0x74, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x46, 0x61,
	0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x66, 0x61, 0x76,
	0x6f, 0x72, 0x69, 0x74, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x46, 0x61,
	0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x12, 0x4f, 0x0a, 0x0f, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x73, 0x12, 0x1d, 0x2e,
	0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x66,
	0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46,
	0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x42, 0x14, 0x5a, 0x12, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_favorites_proto_rawDescOnce sync.Once
	file_favorites_proto_rawDescData = file_favorites_proto_rawDesc
)

func file_favorites_proto_rawDescGZIP() []byte {
	file_favorites_proto_rawDescOnce.Do(func() {
		file_favorites_proto_rawDescData = protoimpl.X.CompressGZIP(file_favorites_proto_rawDescData)
	})
	return file_favorites_proto_rawDescData
}

var file_favorites_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_favorites_proto_goTypes = []any{
	(*CreateFavoritesReq)(nil),  // 0: favorites.CreateFavoritesReq
	(*CreateFavoritesRes)(nil),  // 1: favorites.CreateFavoritesRes
	(*GetAllFavoritesReq)(nil),  // 2: favorites.GetAllFavoritesReq
	(*Favorite)(nil),            // 3: favorites.Favorite
	(*GetAllFavoritesRes)(nil),  // 4: favorites.GetAllFavoritesRes
	(*GetByIdFavoritesReq)(nil), // 5: favorites.GetByIdFavoritesReq
	(*GetByIdFavoritesRes)(nil), // 6: favorites.GetByIdFavoritesRes
	(*DeleteFavoritesReq)(nil),  // 7: favorites.DeleteFavoritesReq
	(*DeleteFavoritesRes)(nil),  // 8: favorites.DeleteFavoritesRes
}
var file_favorites_proto_depIdxs = []int32{
	3, // 0: favorites.GetAllFavoritesRes.favorites:type_name -> favorites.Favorite
	0, // 1: favorites.Favorites.CreateFavorites:input_type -> favorites.CreateFavoritesReq
	2, // 2: favorites.Favorites.GetAllFavorites:input_type -> favorites.GetAllFavoritesReq
	5, // 3: favorites.Favorites.GetByIdFavorites:input_type -> favorites.GetByIdFavoritesReq
	7, // 4: favorites.Favorites.DeleteFavorites:input_type -> favorites.DeleteFavoritesReq
	1, // 5: favorites.Favorites.CreateFavorites:output_type -> favorites.CreateFavoritesRes
	4, // 6: favorites.Favorites.GetAllFavorites:output_type -> favorites.GetAllFavoritesRes
	6, // 7: favorites.Favorites.GetByIdFavorites:output_type -> favorites.GetByIdFavoritesRes
	8, // 8: favorites.Favorites.DeleteFavorites:output_type -> favorites.DeleteFavoritesRes
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_favorites_proto_init() }
func file_favorites_proto_init() {
	if File_favorites_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_favorites_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateFavoritesReq); i {
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
		file_favorites_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateFavoritesRes); i {
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
		file_favorites_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllFavoritesReq); i {
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
		file_favorites_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Favorite); i {
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
		file_favorites_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllFavoritesRes); i {
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
		file_favorites_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetByIdFavoritesReq); i {
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
		file_favorites_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*GetByIdFavoritesRes); i {
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
		file_favorites_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteFavoritesReq); i {
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
		file_favorites_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteFavoritesRes); i {
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
			RawDescriptor: file_favorites_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_favorites_proto_goTypes,
		DependencyIndexes: file_favorites_proto_depIdxs,
		MessageInfos:      file_favorites_proto_msgTypes,
	}.Build()
	File_favorites_proto = out.File
	file_favorites_proto_rawDesc = nil
	file_favorites_proto_goTypes = nil
	file_favorites_proto_depIdxs = nil
}
