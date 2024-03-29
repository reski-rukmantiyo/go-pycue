// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: OpenCue/proto/subscription.proto

package subscription

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

type Subscription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name           string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ShowName       string `protobuf:"bytes,3,opt,name=show_name,json=showName,proto3" json:"show_name,omitempty"`
	Facility       string `protobuf:"bytes,4,opt,name=facility,proto3" json:"facility,omitempty"`
	AllocationName string `protobuf:"bytes,5,opt,name=allocation_name,json=allocationName,proto3" json:"allocation_name,omitempty"`
	Size           int32  `protobuf:"varint,6,opt,name=size,proto3" json:"size,omitempty"`
	Burst          int32  `protobuf:"varint,7,opt,name=burst,proto3" json:"burst,omitempty"`
	ReservedCores  int32  `protobuf:"varint,8,opt,name=reserved_cores,json=reservedCores,proto3" json:"reserved_cores,omitempty"`
	ReservedGpus   int32  `protobuf:"varint,9,opt,name=reserved_gpus,json=reservedGpus,proto3" json:"reserved_gpus,omitempty"`
}

func (x *Subscription) Reset() {
	*x = Subscription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_OpenCue_proto_subscription_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subscription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subscription) ProtoMessage() {}

func (x *Subscription) ProtoReflect() protoreflect.Message {
	mi := &file_OpenCue_proto_subscription_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subscription.ProtoReflect.Descriptor instead.
func (*Subscription) Descriptor() ([]byte, []int) {
	return file_OpenCue_proto_subscription_proto_rawDescGZIP(), []int{0}
}

func (x *Subscription) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Subscription) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Subscription) GetShowName() string {
	if x != nil {
		return x.ShowName
	}
	return ""
}

func (x *Subscription) GetFacility() string {
	if x != nil {
		return x.Facility
	}
	return ""
}

func (x *Subscription) GetAllocationName() string {
	if x != nil {
		return x.AllocationName
	}
	return ""
}

func (x *Subscription) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Subscription) GetBurst() int32 {
	if x != nil {
		return x.Burst
	}
	return 0
}

func (x *Subscription) GetReservedCores() int32 {
	if x != nil {
		return x.ReservedCores
	}
	return 0
}

func (x *Subscription) GetReservedGpus() int32 {
	if x != nil {
		return x.ReservedGpus
	}
	return 0
}

type SubscriptionSeq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subscriptions []*Subscription `protobuf:"bytes,1,rep,name=subscriptions,proto3" json:"subscriptions,omitempty"`
}

func (x *SubscriptionSeq) Reset() {
	*x = SubscriptionSeq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_OpenCue_proto_subscription_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscriptionSeq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscriptionSeq) ProtoMessage() {}

func (x *SubscriptionSeq) ProtoReflect() protoreflect.Message {
	mi := &file_OpenCue_proto_subscription_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscriptionSeq.ProtoReflect.Descriptor instead.
func (*SubscriptionSeq) Descriptor() ([]byte, []int) {
	return file_OpenCue_proto_subscription_proto_rawDescGZIP(), []int{1}
}

func (x *SubscriptionSeq) GetSubscriptions() []*Subscription {
	if x != nil {
		return x.Subscriptions
	}
	return nil
}

// Delete
type SubscriptionDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subscription *Subscription `protobuf:"bytes,1,opt,name=subscription,proto3" json:"subscription,omitempty"`
}

func (x *SubscriptionDeleteRequest) Reset() {
	*x = SubscriptionDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_OpenCue_proto_subscription_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscriptionDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscriptionDeleteRequest) ProtoMessage() {}

func (x *SubscriptionDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_OpenCue_proto_subscription_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscriptionDeleteRequest.ProtoReflect.Descriptor instead.
func (*SubscriptionDeleteRequest) Descriptor() ([]byte, []int) {
	return file_OpenCue_proto_subscription_proto_rawDescGZIP(), []int{2}
}

func (x *SubscriptionDeleteRequest) GetSubscription() *Subscription {
	if x != nil {
		return x.Subscription
	}
	return nil
}

type SubscriptionDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SubscriptionDeleteResponse) Reset() {
	*x = SubscriptionDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_OpenCue_proto_subscription_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscriptionDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscriptionDeleteResponse) ProtoMessage() {}

func (x *SubscriptionDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_OpenCue_proto_subscription_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscriptionDeleteResponse.ProtoReflect.Descriptor instead.
func (*SubscriptionDeleteResponse) Descriptor() ([]byte, []int) {
	return file_OpenCue_proto_subscription_proto_rawDescGZIP(), []int{3}
}

// Find
type SubscriptionFindRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SubscriptionFindRequest) Reset() {
	*x = SubscriptionFindRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_OpenCue_proto_subscription_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscriptionFindRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscriptionFindRequest) ProtoMessage() {}

func (x *SubscriptionFindRequest) ProtoReflect() protoreflect.Message {
	mi := &file_OpenCue_proto_subscription_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscriptionFindRequest.ProtoReflect.Descriptor instead.
func (*SubscriptionFindRequest) Descriptor() ([]byte, []int) {
	return file_OpenCue_proto_subscription_proto_rawDescGZIP(), []int{4}
}

func (x *SubscriptionFindRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type SubscriptionFindResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subscription *Subscription `protobuf:"bytes,1,opt,name=subscription,proto3" json:"subscription,omitempty"`
}

func (x *SubscriptionFindResponse) Reset() {
	*x = SubscriptionFindResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_OpenCue_proto_subscription_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscriptionFindResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscriptionFindResponse) ProtoMessage() {}

func (x *SubscriptionFindResponse) ProtoReflect() protoreflect.Message {
	mi := &file_OpenCue_proto_subscription_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscriptionFindResponse.ProtoReflect.Descriptor instead.
func (*SubscriptionFindResponse) Descriptor() ([]byte, []int) {
	return file_OpenCue_proto_subscription_proto_rawDescGZIP(), []int{5}
}

func (x *SubscriptionFindResponse) GetSubscription() *Subscription {
	if x != nil {
		return x.Subscription
	}
	return nil
}

// Get
type SubscriptionGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SubscriptionGetRequest) Reset() {
	*x = SubscriptionGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_OpenCue_proto_subscription_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscriptionGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscriptionGetRequest) ProtoMessage() {}

func (x *SubscriptionGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_OpenCue_proto_subscription_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscriptionGetRequest.ProtoReflect.Descriptor instead.
func (*SubscriptionGetRequest) Descriptor() ([]byte, []int) {
	return file_OpenCue_proto_subscription_proto_rawDescGZIP(), []int{6}
}

func (x *SubscriptionGetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type SubscriptionGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subscription *Subscription `protobuf:"bytes,1,opt,name=subscription,proto3" json:"subscription,omitempty"`
}

func (x *SubscriptionGetResponse) Reset() {
	*x = SubscriptionGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_OpenCue_proto_subscription_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscriptionGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscriptionGetResponse) ProtoMessage() {}

func (x *SubscriptionGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_OpenCue_proto_subscription_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscriptionGetResponse.ProtoReflect.Descriptor instead.
func (*SubscriptionGetResponse) Descriptor() ([]byte, []int) {
	return file_OpenCue_proto_subscription_proto_rawDescGZIP(), []int{7}
}

func (x *SubscriptionGetResponse) GetSubscription() *Subscription {
	if x != nil {
		return x.Subscription
	}
	return nil
}

// SetBurst
type SubscriptionSetBurstRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subscription *Subscription `protobuf:"bytes,1,opt,name=subscription,proto3" json:"subscription,omitempty"`
	Burst        int32         `protobuf:"varint,2,opt,name=burst,proto3" json:"burst,omitempty"`
}

func (x *SubscriptionSetBurstRequest) Reset() {
	*x = SubscriptionSetBurstRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_OpenCue_proto_subscription_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscriptionSetBurstRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscriptionSetBurstRequest) ProtoMessage() {}

func (x *SubscriptionSetBurstRequest) ProtoReflect() protoreflect.Message {
	mi := &file_OpenCue_proto_subscription_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscriptionSetBurstRequest.ProtoReflect.Descriptor instead.
func (*SubscriptionSetBurstRequest) Descriptor() ([]byte, []int) {
	return file_OpenCue_proto_subscription_proto_rawDescGZIP(), []int{8}
}

func (x *SubscriptionSetBurstRequest) GetSubscription() *Subscription {
	if x != nil {
		return x.Subscription
	}
	return nil
}

func (x *SubscriptionSetBurstRequest) GetBurst() int32 {
	if x != nil {
		return x.Burst
	}
	return 0
}

type SubscriptionSetBurstResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SubscriptionSetBurstResponse) Reset() {
	*x = SubscriptionSetBurstResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_OpenCue_proto_subscription_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscriptionSetBurstResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscriptionSetBurstResponse) ProtoMessage() {}

func (x *SubscriptionSetBurstResponse) ProtoReflect() protoreflect.Message {
	mi := &file_OpenCue_proto_subscription_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscriptionSetBurstResponse.ProtoReflect.Descriptor instead.
func (*SubscriptionSetBurstResponse) Descriptor() ([]byte, []int) {
	return file_OpenCue_proto_subscription_proto_rawDescGZIP(), []int{9}
}

// SetSize
type SubscriptionSetSizeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subscription *Subscription `protobuf:"bytes,1,opt,name=subscription,proto3" json:"subscription,omitempty"`
	NewSize      int32         `protobuf:"varint,2,opt,name=new_size,json=newSize,proto3" json:"new_size,omitempty"`
}

func (x *SubscriptionSetSizeRequest) Reset() {
	*x = SubscriptionSetSizeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_OpenCue_proto_subscription_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscriptionSetSizeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscriptionSetSizeRequest) ProtoMessage() {}

func (x *SubscriptionSetSizeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_OpenCue_proto_subscription_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscriptionSetSizeRequest.ProtoReflect.Descriptor instead.
func (*SubscriptionSetSizeRequest) Descriptor() ([]byte, []int) {
	return file_OpenCue_proto_subscription_proto_rawDescGZIP(), []int{10}
}

func (x *SubscriptionSetSizeRequest) GetSubscription() *Subscription {
	if x != nil {
		return x.Subscription
	}
	return nil
}

func (x *SubscriptionSetSizeRequest) GetNewSize() int32 {
	if x != nil {
		return x.NewSize
	}
	return 0
}

type SubscriptionSetSizeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SubscriptionSetSizeResponse) Reset() {
	*x = SubscriptionSetSizeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_OpenCue_proto_subscription_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscriptionSetSizeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscriptionSetSizeResponse) ProtoMessage() {}

func (x *SubscriptionSetSizeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_OpenCue_proto_subscription_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscriptionSetSizeResponse.ProtoReflect.Descriptor instead.
func (*SubscriptionSetSizeResponse) Descriptor() ([]byte, []int) {
	return file_OpenCue_proto_subscription_proto_rawDescGZIP(), []int{11}
}

var File_OpenCue_proto_subscription_proto protoreflect.FileDescriptor

var file_OpenCue_proto_subscription_proto_rawDesc = []byte{
	0x0a, 0x20, 0x4f, 0x70, 0x65, 0x6e, 0x43, 0x75, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x8a, 0x02, 0x0a, 0x0c, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x77, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x77, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x61, 0x63, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x61, 0x63, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x27,
	0x0a, 0x0f, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x62,
	0x75, 0x72, 0x73, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x62, 0x75, 0x72, 0x73,
	0x74, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x5f, 0x63, 0x6f,
	0x72, 0x65, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x64, 0x43, 0x6f, 0x72, 0x65, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x64, 0x5f, 0x67, 0x70, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0c, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x47, 0x70, 0x75, 0x73, 0x22, 0x53, 0x0a,
	0x0f, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x71,
	0x12, 0x40, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x22, 0x5b, 0x0a, 0x19, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x3e, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x1c, 0x0a, 0x1a, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2d, 0x0a,
	0x17, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6e,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x5a, 0x0a, 0x18,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6e, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x28, 0x0a, 0x16, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x59, 0x0a, 0x17, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a,
	0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x73, 0x0a,
	0x1b, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74,
	0x42, 0x75, 0x72, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x0c,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x62, 0x75, 0x72, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x62, 0x75, 0x72,
	0x73, 0x74, 0x22, 0x1e, 0x0a, 0x1c, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x65, 0x74, 0x42, 0x75, 0x72, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x77, 0x0a, 0x1a, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x65, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x3e, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x19, 0x0a, 0x08, 0x6e, 0x65, 0x77, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x6e, 0x65, 0x77, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x1d, 0x0a, 0x1b, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x53, 0x69,
	0x7a, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xe2, 0x03, 0x0a, 0x15, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x12, 0x5b, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x27,
	0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x55, 0x0a, 0x04, 0x46, 0x69, 0x6e, 0x64, 0x12, 0x25, 0x2e, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x26, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6e, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12,
	0x24, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x61, 0x0a, 0x08,
	0x53, 0x65, 0x74, 0x42, 0x75, 0x72, 0x73, 0x74, 0x12, 0x29, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x42, 0x75, 0x72, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x74, 0x42, 0x75, 0x72, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x5e, 0x0a, 0x07, 0x53, 0x65, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x28, 0x2e, 0x73, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x3e, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x77, 0x6f, 0x72, 0x6b,
	0x73, 0x2e, 0x73, 0x70, 0x63, 0x75, 0x65, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x01, 0x5a, 0x12, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_OpenCue_proto_subscription_proto_rawDescOnce sync.Once
	file_OpenCue_proto_subscription_proto_rawDescData = file_OpenCue_proto_subscription_proto_rawDesc
)

func file_OpenCue_proto_subscription_proto_rawDescGZIP() []byte {
	file_OpenCue_proto_subscription_proto_rawDescOnce.Do(func() {
		file_OpenCue_proto_subscription_proto_rawDescData = protoimpl.X.CompressGZIP(file_OpenCue_proto_subscription_proto_rawDescData)
	})
	return file_OpenCue_proto_subscription_proto_rawDescData
}

var file_OpenCue_proto_subscription_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_OpenCue_proto_subscription_proto_goTypes = []interface{}{
	(*Subscription)(nil),                 // 0: subscription.Subscription
	(*SubscriptionSeq)(nil),              // 1: subscription.SubscriptionSeq
	(*SubscriptionDeleteRequest)(nil),    // 2: subscription.SubscriptionDeleteRequest
	(*SubscriptionDeleteResponse)(nil),   // 3: subscription.SubscriptionDeleteResponse
	(*SubscriptionFindRequest)(nil),      // 4: subscription.SubscriptionFindRequest
	(*SubscriptionFindResponse)(nil),     // 5: subscription.SubscriptionFindResponse
	(*SubscriptionGetRequest)(nil),       // 6: subscription.SubscriptionGetRequest
	(*SubscriptionGetResponse)(nil),      // 7: subscription.SubscriptionGetResponse
	(*SubscriptionSetBurstRequest)(nil),  // 8: subscription.SubscriptionSetBurstRequest
	(*SubscriptionSetBurstResponse)(nil), // 9: subscription.SubscriptionSetBurstResponse
	(*SubscriptionSetSizeRequest)(nil),   // 10: subscription.SubscriptionSetSizeRequest
	(*SubscriptionSetSizeResponse)(nil),  // 11: subscription.SubscriptionSetSizeResponse
}
var file_OpenCue_proto_subscription_proto_depIdxs = []int32{
	0,  // 0: subscription.SubscriptionSeq.subscriptions:type_name -> subscription.Subscription
	0,  // 1: subscription.SubscriptionDeleteRequest.subscription:type_name -> subscription.Subscription
	0,  // 2: subscription.SubscriptionFindResponse.subscription:type_name -> subscription.Subscription
	0,  // 3: subscription.SubscriptionGetResponse.subscription:type_name -> subscription.Subscription
	0,  // 4: subscription.SubscriptionSetBurstRequest.subscription:type_name -> subscription.Subscription
	0,  // 5: subscription.SubscriptionSetSizeRequest.subscription:type_name -> subscription.Subscription
	2,  // 6: subscription.SubscriptionInterface.Delete:input_type -> subscription.SubscriptionDeleteRequest
	4,  // 7: subscription.SubscriptionInterface.Find:input_type -> subscription.SubscriptionFindRequest
	6,  // 8: subscription.SubscriptionInterface.Get:input_type -> subscription.SubscriptionGetRequest
	8,  // 9: subscription.SubscriptionInterface.SetBurst:input_type -> subscription.SubscriptionSetBurstRequest
	10, // 10: subscription.SubscriptionInterface.SetSize:input_type -> subscription.SubscriptionSetSizeRequest
	3,  // 11: subscription.SubscriptionInterface.Delete:output_type -> subscription.SubscriptionDeleteResponse
	5,  // 12: subscription.SubscriptionInterface.Find:output_type -> subscription.SubscriptionFindResponse
	7,  // 13: subscription.SubscriptionInterface.Get:output_type -> subscription.SubscriptionGetResponse
	9,  // 14: subscription.SubscriptionInterface.SetBurst:output_type -> subscription.SubscriptionSetBurstResponse
	11, // 15: subscription.SubscriptionInterface.SetSize:output_type -> subscription.SubscriptionSetSizeResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_OpenCue_proto_subscription_proto_init() }
func file_OpenCue_proto_subscription_proto_init() {
	if File_OpenCue_proto_subscription_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_OpenCue_proto_subscription_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Subscription); i {
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
		file_OpenCue_proto_subscription_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscriptionSeq); i {
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
		file_OpenCue_proto_subscription_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscriptionDeleteRequest); i {
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
		file_OpenCue_proto_subscription_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscriptionDeleteResponse); i {
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
		file_OpenCue_proto_subscription_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscriptionFindRequest); i {
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
		file_OpenCue_proto_subscription_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscriptionFindResponse); i {
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
		file_OpenCue_proto_subscription_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscriptionGetRequest); i {
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
		file_OpenCue_proto_subscription_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscriptionGetResponse); i {
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
		file_OpenCue_proto_subscription_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscriptionSetBurstRequest); i {
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
		file_OpenCue_proto_subscription_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscriptionSetBurstResponse); i {
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
		file_OpenCue_proto_subscription_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscriptionSetSizeRequest); i {
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
		file_OpenCue_proto_subscription_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscriptionSetSizeResponse); i {
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
			RawDescriptor: file_OpenCue_proto_subscription_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_OpenCue_proto_subscription_proto_goTypes,
		DependencyIndexes: file_OpenCue_proto_subscription_proto_depIdxs,
		MessageInfos:      file_OpenCue_proto_subscription_proto_msgTypes,
	}.Build()
	File_OpenCue_proto_subscription_proto = out.File
	file_OpenCue_proto_subscription_proto_rawDesc = nil
	file_OpenCue_proto_subscription_proto_goTypes = nil
	file_OpenCue_proto_subscription_proto_depIdxs = nil
}
