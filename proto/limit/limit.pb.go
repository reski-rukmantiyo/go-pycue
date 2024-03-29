// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: OpenCue/proto/limit.proto

package limit

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

type Limit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name           string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	MaxValue       int32  `protobuf:"varint,3,opt,name=max_value,json=maxValue,proto3" json:"max_value,omitempty"`
	CurrentRunning int32  `protobuf:"varint,4,opt,name=current_running,json=currentRunning,proto3" json:"current_running,omitempty"`
}

func (x *Limit) Reset() {
	*x = Limit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_OpenCue_proto_limit_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Limit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Limit) ProtoMessage() {}

func (x *Limit) ProtoReflect() protoreflect.Message {
	mi := &file_OpenCue_proto_limit_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Limit.ProtoReflect.Descriptor instead.
func (*Limit) Descriptor() ([]byte, []int) {
	return file_OpenCue_proto_limit_proto_rawDescGZIP(), []int{0}
}

func (x *Limit) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Limit) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Limit) GetMaxValue() int32 {
	if x != nil {
		return x.MaxValue
	}
	return 0
}

func (x *Limit) GetCurrentRunning() int32 {
	if x != nil {
		return x.CurrentRunning
	}
	return 0
}

// Create
type LimitCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	MaxValue int32  `protobuf:"varint,2,opt,name=max_value,json=maxValue,proto3" json:"max_value,omitempty"`
}

func (x *LimitCreateRequest) Reset() {
	*x = LimitCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_OpenCue_proto_limit_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LimitCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LimitCreateRequest) ProtoMessage() {}

func (x *LimitCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_OpenCue_proto_limit_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LimitCreateRequest.ProtoReflect.Descriptor instead.
func (*LimitCreateRequest) Descriptor() ([]byte, []int) {
	return file_OpenCue_proto_limit_proto_rawDescGZIP(), []int{1}
}

func (x *LimitCreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LimitCreateRequest) GetMaxValue() int32 {
	if x != nil {
		return x.MaxValue
	}
	return 0
}

type LimitCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit *Limit `protobuf:"bytes,1,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *LimitCreateResponse) Reset() {
	*x = LimitCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_OpenCue_proto_limit_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LimitCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LimitCreateResponse) ProtoMessage() {}

func (x *LimitCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_OpenCue_proto_limit_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LimitCreateResponse.ProtoReflect.Descriptor instead.
func (*LimitCreateResponse) Descriptor() ([]byte, []int) {
	return file_OpenCue_proto_limit_proto_rawDescGZIP(), []int{2}
}

func (x *LimitCreateResponse) GetLimit() *Limit {
	if x != nil {
		return x.Limit
	}
	return nil
}

// Delete
type LimitDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *LimitDeleteRequest) Reset() {
	*x = LimitDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_OpenCue_proto_limit_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LimitDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LimitDeleteRequest) ProtoMessage() {}

func (x *LimitDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_OpenCue_proto_limit_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LimitDeleteRequest.ProtoReflect.Descriptor instead.
func (*LimitDeleteRequest) Descriptor() ([]byte, []int) {
	return file_OpenCue_proto_limit_proto_rawDescGZIP(), []int{3}
}

func (x *LimitDeleteRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type LimitDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LimitDeleteResponse) Reset() {
	*x = LimitDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_OpenCue_proto_limit_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LimitDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LimitDeleteResponse) ProtoMessage() {}

func (x *LimitDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_OpenCue_proto_limit_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LimitDeleteResponse.ProtoReflect.Descriptor instead.
func (*LimitDeleteResponse) Descriptor() ([]byte, []int) {
	return file_OpenCue_proto_limit_proto_rawDescGZIP(), []int{4}
}

// Find
type LimitFindRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *LimitFindRequest) Reset() {
	*x = LimitFindRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_OpenCue_proto_limit_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LimitFindRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LimitFindRequest) ProtoMessage() {}

func (x *LimitFindRequest) ProtoReflect() protoreflect.Message {
	mi := &file_OpenCue_proto_limit_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LimitFindRequest.ProtoReflect.Descriptor instead.
func (*LimitFindRequest) Descriptor() ([]byte, []int) {
	return file_OpenCue_proto_limit_proto_rawDescGZIP(), []int{5}
}

func (x *LimitFindRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type LimitFindResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit *Limit `protobuf:"bytes,1,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *LimitFindResponse) Reset() {
	*x = LimitFindResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_OpenCue_proto_limit_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LimitFindResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LimitFindResponse) ProtoMessage() {}

func (x *LimitFindResponse) ProtoReflect() protoreflect.Message {
	mi := &file_OpenCue_proto_limit_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LimitFindResponse.ProtoReflect.Descriptor instead.
func (*LimitFindResponse) Descriptor() ([]byte, []int) {
	return file_OpenCue_proto_limit_proto_rawDescGZIP(), []int{6}
}

func (x *LimitFindResponse) GetLimit() *Limit {
	if x != nil {
		return x.Limit
	}
	return nil
}

// Get
type LimitGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *LimitGetRequest) Reset() {
	*x = LimitGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_OpenCue_proto_limit_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LimitGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LimitGetRequest) ProtoMessage() {}

func (x *LimitGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_OpenCue_proto_limit_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LimitGetRequest.ProtoReflect.Descriptor instead.
func (*LimitGetRequest) Descriptor() ([]byte, []int) {
	return file_OpenCue_proto_limit_proto_rawDescGZIP(), []int{7}
}

func (x *LimitGetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type LimitGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit *Limit `protobuf:"bytes,1,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *LimitGetResponse) Reset() {
	*x = LimitGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_OpenCue_proto_limit_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LimitGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LimitGetResponse) ProtoMessage() {}

func (x *LimitGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_OpenCue_proto_limit_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LimitGetResponse.ProtoReflect.Descriptor instead.
func (*LimitGetResponse) Descriptor() ([]byte, []int) {
	return file_OpenCue_proto_limit_proto_rawDescGZIP(), []int{8}
}

func (x *LimitGetResponse) GetLimit() *Limit {
	if x != nil {
		return x.Limit
	}
	return nil
}

// GetAll
type LimitGetAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LimitGetAllRequest) Reset() {
	*x = LimitGetAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_OpenCue_proto_limit_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LimitGetAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LimitGetAllRequest) ProtoMessage() {}

func (x *LimitGetAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_OpenCue_proto_limit_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LimitGetAllRequest.ProtoReflect.Descriptor instead.
func (*LimitGetAllRequest) Descriptor() ([]byte, []int) {
	return file_OpenCue_proto_limit_proto_rawDescGZIP(), []int{9}
}

type LimitGetAllResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limits []*Limit `protobuf:"bytes,1,rep,name=limits,proto3" json:"limits,omitempty"`
}

func (x *LimitGetAllResponse) Reset() {
	*x = LimitGetAllResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_OpenCue_proto_limit_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LimitGetAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LimitGetAllResponse) ProtoMessage() {}

func (x *LimitGetAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_OpenCue_proto_limit_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LimitGetAllResponse.ProtoReflect.Descriptor instead.
func (*LimitGetAllResponse) Descriptor() ([]byte, []int) {
	return file_OpenCue_proto_limit_proto_rawDescGZIP(), []int{10}
}

func (x *LimitGetAllResponse) GetLimits() []*Limit {
	if x != nil {
		return x.Limits
	}
	return nil
}

// Rename
type LimitRenameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OldName string `protobuf:"bytes,1,opt,name=old_name,json=oldName,proto3" json:"old_name,omitempty"`
	NewName string `protobuf:"bytes,2,opt,name=new_name,json=newName,proto3" json:"new_name,omitempty"`
}

func (x *LimitRenameRequest) Reset() {
	*x = LimitRenameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_OpenCue_proto_limit_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LimitRenameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LimitRenameRequest) ProtoMessage() {}

func (x *LimitRenameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_OpenCue_proto_limit_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LimitRenameRequest.ProtoReflect.Descriptor instead.
func (*LimitRenameRequest) Descriptor() ([]byte, []int) {
	return file_OpenCue_proto_limit_proto_rawDescGZIP(), []int{11}
}

func (x *LimitRenameRequest) GetOldName() string {
	if x != nil {
		return x.OldName
	}
	return ""
}

func (x *LimitRenameRequest) GetNewName() string {
	if x != nil {
		return x.NewName
	}
	return ""
}

type LimitRenameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LimitRenameResponse) Reset() {
	*x = LimitRenameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_OpenCue_proto_limit_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LimitRenameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LimitRenameResponse) ProtoMessage() {}

func (x *LimitRenameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_OpenCue_proto_limit_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LimitRenameResponse.ProtoReflect.Descriptor instead.
func (*LimitRenameResponse) Descriptor() ([]byte, []int) {
	return file_OpenCue_proto_limit_proto_rawDescGZIP(), []int{12}
}

// SetMaxValue
type LimitSetMaxValueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	MaxValue int32  `protobuf:"varint,2,opt,name=max_value,json=maxValue,proto3" json:"max_value,omitempty"`
}

func (x *LimitSetMaxValueRequest) Reset() {
	*x = LimitSetMaxValueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_OpenCue_proto_limit_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LimitSetMaxValueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LimitSetMaxValueRequest) ProtoMessage() {}

func (x *LimitSetMaxValueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_OpenCue_proto_limit_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LimitSetMaxValueRequest.ProtoReflect.Descriptor instead.
func (*LimitSetMaxValueRequest) Descriptor() ([]byte, []int) {
	return file_OpenCue_proto_limit_proto_rawDescGZIP(), []int{13}
}

func (x *LimitSetMaxValueRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LimitSetMaxValueRequest) GetMaxValue() int32 {
	if x != nil {
		return x.MaxValue
	}
	return 0
}

type LimitSetMaxValueResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LimitSetMaxValueResponse) Reset() {
	*x = LimitSetMaxValueResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_OpenCue_proto_limit_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LimitSetMaxValueResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LimitSetMaxValueResponse) ProtoMessage() {}

func (x *LimitSetMaxValueResponse) ProtoReflect() protoreflect.Message {
	mi := &file_OpenCue_proto_limit_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LimitSetMaxValueResponse.ProtoReflect.Descriptor instead.
func (*LimitSetMaxValueResponse) Descriptor() ([]byte, []int) {
	return file_OpenCue_proto_limit_proto_rawDescGZIP(), []int{14}
}

var File_OpenCue_proto_limit_proto protoreflect.FileDescriptor

var file_OpenCue_proto_limit_proto_rawDesc = []byte{
	0x0a, 0x19, 0x4f, 0x70, 0x65, 0x6e, 0x43, 0x75, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x22, 0x71, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x27, 0x0a, 0x0f,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x75,
	0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x22, 0x45, 0x0a, 0x12, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x39, 0x0a, 0x13,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x28, 0x0a, 0x12, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x15, 0x0a, 0x13, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26, 0x0a, 0x10, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x37, 0x0a, 0x11, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x21, 0x0a, 0x0f, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x36, 0x0a, 0x10,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x22, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x22, 0x14, 0x0a, 0x12, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3b, 0x0a, 0x13, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x24, 0x0a, 0x06, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52,
	0x06, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x22, 0x4a, 0x0a, 0x12, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a,
	0x08, 0x6f, 0x6c, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6f, 0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x65, 0x77, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x77, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x15, 0x0a, 0x13, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x6e, 0x61,
	0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4a, 0x0a, 0x17, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x53, 0x65, 0x74, 0x4d, 0x61, 0x78, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x78,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6d, 0x61,
	0x78, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x1a, 0x0a, 0x18, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x53,
	0x65, 0x74, 0x4d, 0x61, 0x78, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x32, 0xd7, 0x03, 0x0a, 0x0e, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x19, 0x2e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x2e, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x12, 0x19, 0x2e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x2e, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x04, 0x46, 0x69, 0x6e, 0x64, 0x12,
	0x17, 0x2e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x46, 0x69, 0x6e,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x2e, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x36, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x16, 0x2e, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x2e, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x17, 0x2e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x06, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x12, 0x19, 0x2e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x06, 0x52,
	0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x2e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65,
	0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x0b,
	0x53, 0x65, 0x74, 0x4d, 0x61, 0x78, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1e, 0x2e, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x2e, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x53, 0x65, 0x74, 0x4d, 0x61, 0x78, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x2e, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x53, 0x65, 0x74, 0x4d, 0x61, 0x78, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x30, 0x0a, 0x1f,
	0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2e, 0x73,
	0x70, 0x63, 0x75, 0x65, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x50,
	0x01, 0x5a, 0x0b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_OpenCue_proto_limit_proto_rawDescOnce sync.Once
	file_OpenCue_proto_limit_proto_rawDescData = file_OpenCue_proto_limit_proto_rawDesc
)

func file_OpenCue_proto_limit_proto_rawDescGZIP() []byte {
	file_OpenCue_proto_limit_proto_rawDescOnce.Do(func() {
		file_OpenCue_proto_limit_proto_rawDescData = protoimpl.X.CompressGZIP(file_OpenCue_proto_limit_proto_rawDescData)
	})
	return file_OpenCue_proto_limit_proto_rawDescData
}

var file_OpenCue_proto_limit_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_OpenCue_proto_limit_proto_goTypes = []interface{}{
	(*Limit)(nil),                    // 0: limit.Limit
	(*LimitCreateRequest)(nil),       // 1: limit.LimitCreateRequest
	(*LimitCreateResponse)(nil),      // 2: limit.LimitCreateResponse
	(*LimitDeleteRequest)(nil),       // 3: limit.LimitDeleteRequest
	(*LimitDeleteResponse)(nil),      // 4: limit.LimitDeleteResponse
	(*LimitFindRequest)(nil),         // 5: limit.LimitFindRequest
	(*LimitFindResponse)(nil),        // 6: limit.LimitFindResponse
	(*LimitGetRequest)(nil),          // 7: limit.LimitGetRequest
	(*LimitGetResponse)(nil),         // 8: limit.LimitGetResponse
	(*LimitGetAllRequest)(nil),       // 9: limit.LimitGetAllRequest
	(*LimitGetAllResponse)(nil),      // 10: limit.LimitGetAllResponse
	(*LimitRenameRequest)(nil),       // 11: limit.LimitRenameRequest
	(*LimitRenameResponse)(nil),      // 12: limit.LimitRenameResponse
	(*LimitSetMaxValueRequest)(nil),  // 13: limit.LimitSetMaxValueRequest
	(*LimitSetMaxValueResponse)(nil), // 14: limit.LimitSetMaxValueResponse
}
var file_OpenCue_proto_limit_proto_depIdxs = []int32{
	0,  // 0: limit.LimitCreateResponse.limit:type_name -> limit.Limit
	0,  // 1: limit.LimitFindResponse.limit:type_name -> limit.Limit
	0,  // 2: limit.LimitGetResponse.limit:type_name -> limit.Limit
	0,  // 3: limit.LimitGetAllResponse.limits:type_name -> limit.Limit
	1,  // 4: limit.LimitInterface.Create:input_type -> limit.LimitCreateRequest
	3,  // 5: limit.LimitInterface.Delete:input_type -> limit.LimitDeleteRequest
	5,  // 6: limit.LimitInterface.Find:input_type -> limit.LimitFindRequest
	7,  // 7: limit.LimitInterface.Get:input_type -> limit.LimitGetRequest
	9,  // 8: limit.LimitInterface.GetAll:input_type -> limit.LimitGetAllRequest
	11, // 9: limit.LimitInterface.Rename:input_type -> limit.LimitRenameRequest
	13, // 10: limit.LimitInterface.SetMaxValue:input_type -> limit.LimitSetMaxValueRequest
	2,  // 11: limit.LimitInterface.Create:output_type -> limit.LimitCreateResponse
	4,  // 12: limit.LimitInterface.Delete:output_type -> limit.LimitDeleteResponse
	6,  // 13: limit.LimitInterface.Find:output_type -> limit.LimitFindResponse
	8,  // 14: limit.LimitInterface.Get:output_type -> limit.LimitGetResponse
	10, // 15: limit.LimitInterface.GetAll:output_type -> limit.LimitGetAllResponse
	12, // 16: limit.LimitInterface.Rename:output_type -> limit.LimitRenameResponse
	14, // 17: limit.LimitInterface.SetMaxValue:output_type -> limit.LimitSetMaxValueResponse
	11, // [11:18] is the sub-list for method output_type
	4,  // [4:11] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_OpenCue_proto_limit_proto_init() }
func file_OpenCue_proto_limit_proto_init() {
	if File_OpenCue_proto_limit_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_OpenCue_proto_limit_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Limit); i {
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
		file_OpenCue_proto_limit_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LimitCreateRequest); i {
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
		file_OpenCue_proto_limit_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LimitCreateResponse); i {
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
		file_OpenCue_proto_limit_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LimitDeleteRequest); i {
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
		file_OpenCue_proto_limit_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LimitDeleteResponse); i {
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
		file_OpenCue_proto_limit_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LimitFindRequest); i {
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
		file_OpenCue_proto_limit_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LimitFindResponse); i {
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
		file_OpenCue_proto_limit_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LimitGetRequest); i {
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
		file_OpenCue_proto_limit_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LimitGetResponse); i {
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
		file_OpenCue_proto_limit_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LimitGetAllRequest); i {
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
		file_OpenCue_proto_limit_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LimitGetAllResponse); i {
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
		file_OpenCue_proto_limit_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LimitRenameRequest); i {
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
		file_OpenCue_proto_limit_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LimitRenameResponse); i {
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
		file_OpenCue_proto_limit_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LimitSetMaxValueRequest); i {
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
		file_OpenCue_proto_limit_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LimitSetMaxValueResponse); i {
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
			RawDescriptor: file_OpenCue_proto_limit_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_OpenCue_proto_limit_proto_goTypes,
		DependencyIndexes: file_OpenCue_proto_limit_proto_depIdxs,
		MessageInfos:      file_OpenCue_proto_limit_proto_msgTypes,
	}.Build()
	File_OpenCue_proto_limit_proto = out.File
	file_OpenCue_proto_limit_proto_rawDesc = nil
	file_OpenCue_proto_limit_proto_goTypes = nil
	file_OpenCue_proto_limit_proto_depIdxs = nil
}
