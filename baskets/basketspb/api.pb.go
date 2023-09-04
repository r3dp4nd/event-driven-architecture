// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: basketspb/api.proto

package basketspb

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

type StartBasketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
}

func (x *StartBasketRequest) Reset() {
	*x = StartBasketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basketspb_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartBasketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartBasketRequest) ProtoMessage() {}

func (x *StartBasketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_basketspb_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartBasketRequest.ProtoReflect.Descriptor instead.
func (*StartBasketRequest) Descriptor() ([]byte, []int) {
	return file_basketspb_api_proto_rawDescGZIP(), []int{0}
}

func (x *StartBasketRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

type StartBasketResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *StartBasketResponse) Reset() {
	*x = StartBasketResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basketspb_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartBasketResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartBasketResponse) ProtoMessage() {}

func (x *StartBasketResponse) ProtoReflect() protoreflect.Message {
	mi := &file_basketspb_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartBasketResponse.ProtoReflect.Descriptor instead.
func (*StartBasketResponse) Descriptor() ([]byte, []int) {
	return file_basketspb_api_proto_rawDescGZIP(), []int{1}
}

func (x *StartBasketResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CancelBasketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CancelBasketRequest) Reset() {
	*x = CancelBasketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basketspb_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelBasketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelBasketRequest) ProtoMessage() {}

func (x *CancelBasketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_basketspb_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelBasketRequest.ProtoReflect.Descriptor instead.
func (*CancelBasketRequest) Descriptor() ([]byte, []int) {
	return file_basketspb_api_proto_rawDescGZIP(), []int{2}
}

func (x *CancelBasketRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CancelBasketResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CancelBasketResponse) Reset() {
	*x = CancelBasketResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basketspb_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelBasketResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelBasketResponse) ProtoMessage() {}

func (x *CancelBasketResponse) ProtoReflect() protoreflect.Message {
	mi := &file_basketspb_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelBasketResponse.ProtoReflect.Descriptor instead.
func (*CancelBasketResponse) Descriptor() ([]byte, []int) {
	return file_basketspb_api_proto_rawDescGZIP(), []int{3}
}

type CheckoutBasketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	PaymentId string `protobuf:"bytes,2,opt,name=payment_id,json=paymentId,proto3" json:"payment_id,omitempty"`
}

func (x *CheckoutBasketRequest) Reset() {
	*x = CheckoutBasketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basketspb_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckoutBasketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckoutBasketRequest) ProtoMessage() {}

func (x *CheckoutBasketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_basketspb_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckoutBasketRequest.ProtoReflect.Descriptor instead.
func (*CheckoutBasketRequest) Descriptor() ([]byte, []int) {
	return file_basketspb_api_proto_rawDescGZIP(), []int{4}
}

func (x *CheckoutBasketRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CheckoutBasketRequest) GetPaymentId() string {
	if x != nil {
		return x.PaymentId
	}
	return ""
}

type CheckoutBasketResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CheckoutBasketResponse) Reset() {
	*x = CheckoutBasketResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basketspb_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckoutBasketResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckoutBasketResponse) ProtoMessage() {}

func (x *CheckoutBasketResponse) ProtoReflect() protoreflect.Message {
	mi := &file_basketspb_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckoutBasketResponse.ProtoReflect.Descriptor instead.
func (*CheckoutBasketResponse) Descriptor() ([]byte, []int) {
	return file_basketspb_api_proto_rawDescGZIP(), []int{5}
}

type AddItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ProductId string `protobuf:"bytes,3,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Quantity  int32  `protobuf:"varint,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *AddItemRequest) Reset() {
	*x = AddItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basketspb_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddItemRequest) ProtoMessage() {}

func (x *AddItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_basketspb_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddItemRequest.ProtoReflect.Descriptor instead.
func (*AddItemRequest) Descriptor() ([]byte, []int) {
	return file_basketspb_api_proto_rawDescGZIP(), []int{6}
}

func (x *AddItemRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AddItemRequest) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *AddItemRequest) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type AddItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddItemResponse) Reset() {
	*x = AddItemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basketspb_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddItemResponse) ProtoMessage() {}

func (x *AddItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_basketspb_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddItemResponse.ProtoReflect.Descriptor instead.
func (*AddItemResponse) Descriptor() ([]byte, []int) {
	return file_basketspb_api_proto_rawDescGZIP(), []int{7}
}

type RemoveItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ProductId string `protobuf:"bytes,3,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Quantity  int32  `protobuf:"varint,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *RemoveItemRequest) Reset() {
	*x = RemoveItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basketspb_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveItemRequest) ProtoMessage() {}

func (x *RemoveItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_basketspb_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveItemRequest.ProtoReflect.Descriptor instead.
func (*RemoveItemRequest) Descriptor() ([]byte, []int) {
	return file_basketspb_api_proto_rawDescGZIP(), []int{8}
}

func (x *RemoveItemRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RemoveItemRequest) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *RemoveItemRequest) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type RemoveItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemoveItemResponse) Reset() {
	*x = RemoveItemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basketspb_api_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveItemResponse) ProtoMessage() {}

func (x *RemoveItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_basketspb_api_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveItemResponse.ProtoReflect.Descriptor instead.
func (*RemoveItemResponse) Descriptor() ([]byte, []int) {
	return file_basketspb_api_proto_rawDescGZIP(), []int{9}
}

type GetBasketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetBasketRequest) Reset() {
	*x = GetBasketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basketspb_api_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBasketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBasketRequest) ProtoMessage() {}

func (x *GetBasketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_basketspb_api_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBasketRequest.ProtoReflect.Descriptor instead.
func (*GetBasketRequest) Descriptor() ([]byte, []int) {
	return file_basketspb_api_proto_rawDescGZIP(), []int{10}
}

func (x *GetBasketRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetBasketResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Basket *Basket `protobuf:"bytes,1,opt,name=basket,proto3" json:"basket,omitempty"`
}

func (x *GetBasketResponse) Reset() {
	*x = GetBasketResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basketspb_api_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBasketResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBasketResponse) ProtoMessage() {}

func (x *GetBasketResponse) ProtoReflect() protoreflect.Message {
	mi := &file_basketspb_api_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBasketResponse.ProtoReflect.Descriptor instead.
func (*GetBasketResponse) Descriptor() ([]byte, []int) {
	return file_basketspb_api_proto_rawDescGZIP(), []int{11}
}

func (x *GetBasketResponse) GetBasket() *Basket {
	if x != nil {
		return x.Basket
	}
	return nil
}

var File_basketspb_api_proto protoreflect.FileDescriptor

var file_basketspb_api_proto_rawDesc = []byte{
	0x0a, 0x13, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x73, 0x70, 0x62, 0x2f, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x73, 0x70, 0x62,
	0x1a, 0x18, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x73, 0x70, 0x62, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x35, 0x0a, 0x12, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x25, 0x0a, 0x13, 0x53, 0x74, 0x61, 0x72, 0x74, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x25, 0x0a, 0x13, 0x43, 0x61, 0x6e, 0x63,
	0x65, 0x6c, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x16, 0x0a, 0x14, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x46, 0x0a, 0x15, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x6f, 0x75, 0x74, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22,
	0x18, 0x0a, 0x16, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x42, 0x61, 0x73, 0x6b, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5b, 0x0a, 0x0e, 0x41, 0x64, 0x64,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75,
	0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75,
	0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x11, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5e, 0x0a, 0x11, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x14, 0x0a, 0x12, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x22, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x3e, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x06, 0x62, 0x61, 0x73, 0x6b,
	0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x62, 0x61, 0x73, 0x6b, 0x65,
	0x74, 0x73, 0x70, 0x62, 0x2e, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x52, 0x06, 0x62, 0x61, 0x73,
	0x6b, 0x65, 0x74, 0x32, 0xe6, 0x03, 0x0a, 0x0d, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x72, 0x74, 0x42, 0x61,
	0x73, 0x6b, 0x65, 0x74, 0x12, 0x1d, 0x2e, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x73, 0x70, 0x62,
	0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x73, 0x70, 0x62, 0x2e,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x0c, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x42,
	0x61, 0x73, 0x6b, 0x65, 0x74, 0x12, 0x1e, 0x2e, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x73, 0x70,
	0x62, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x73, 0x70,
	0x62, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x57, 0x0a, 0x0e, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x6f, 0x75, 0x74, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x12, 0x20, 0x2e, 0x62, 0x61, 0x73,
	0x6b, 0x65, 0x74, 0x73, 0x70, 0x62, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x42,
	0x61, 0x73, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x62,
	0x61, 0x73, 0x6b, 0x65, 0x74, 0x73, 0x70, 0x62, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75,
	0x74, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x42, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x19, 0x2e, 0x62,
	0x61, 0x73, 0x6b, 0x65, 0x74, 0x73, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x64, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74,
	0x73, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x0a, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x12, 0x1c, 0x2e, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x73, 0x70, 0x62, 0x2e,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x73, 0x70, 0x62, 0x2e, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x48, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x12,
	0x1b, 0x2e, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x73, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x42,
	0x61, 0x73, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x62,
	0x61, 0x73, 0x6b, 0x65, 0x74, 0x73, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x61, 0x73, 0x6b,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xa7, 0x01, 0x0a,
	0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x73, 0x70, 0x62, 0x42, 0x08,
	0x41, 0x70, 0x69, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x72, 0x61, 0x69, 0x6b, 0x32, 0x35, 0x2f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x2d, 0x64, 0x72, 0x69, 0x76, 0x65, 0x6e, 0x2d, 0x61, 0x72, 0x63, 0x68,
	0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x73,
	0x2f, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x73, 0x70, 0x62, 0x2f, 0x62, 0x61, 0x73, 0x6b, 0x65,
	0x74, 0x73, 0x70, 0x62, 0xa2, 0x02, 0x03, 0x42, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x42, 0x61, 0x73,
	0x6b, 0x65, 0x74, 0x73, 0x70, 0x62, 0xca, 0x02, 0x09, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x73,
	0x70, 0x62, 0xe2, 0x02, 0x15, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x73, 0x70, 0x62, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x09, 0x42, 0x61, 0x73,
	0x6b, 0x65, 0x74, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_basketspb_api_proto_rawDescOnce sync.Once
	file_basketspb_api_proto_rawDescData = file_basketspb_api_proto_rawDesc
)

func file_basketspb_api_proto_rawDescGZIP() []byte {
	file_basketspb_api_proto_rawDescOnce.Do(func() {
		file_basketspb_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_basketspb_api_proto_rawDescData)
	})
	return file_basketspb_api_proto_rawDescData
}

var file_basketspb_api_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_basketspb_api_proto_goTypes = []interface{}{
	(*StartBasketRequest)(nil),     // 0: basketspb.StartBasketRequest
	(*StartBasketResponse)(nil),    // 1: basketspb.StartBasketResponse
	(*CancelBasketRequest)(nil),    // 2: basketspb.CancelBasketRequest
	(*CancelBasketResponse)(nil),   // 3: basketspb.CancelBasketResponse
	(*CheckoutBasketRequest)(nil),  // 4: basketspb.CheckoutBasketRequest
	(*CheckoutBasketResponse)(nil), // 5: basketspb.CheckoutBasketResponse
	(*AddItemRequest)(nil),         // 6: basketspb.AddItemRequest
	(*AddItemResponse)(nil),        // 7: basketspb.AddItemResponse
	(*RemoveItemRequest)(nil),      // 8: basketspb.RemoveItemRequest
	(*RemoveItemResponse)(nil),     // 9: basketspb.RemoveItemResponse
	(*GetBasketRequest)(nil),       // 10: basketspb.GetBasketRequest
	(*GetBasketResponse)(nil),      // 11: basketspb.GetBasketResponse
	(*Basket)(nil),                 // 12: basketspb.Basket
}
var file_basketspb_api_proto_depIdxs = []int32{
	12, // 0: basketspb.GetBasketResponse.basket:type_name -> basketspb.Basket
	0,  // 1: basketspb.BasketService.StartBasket:input_type -> basketspb.StartBasketRequest
	2,  // 2: basketspb.BasketService.CancelBasket:input_type -> basketspb.CancelBasketRequest
	4,  // 3: basketspb.BasketService.CheckoutBasket:input_type -> basketspb.CheckoutBasketRequest
	6,  // 4: basketspb.BasketService.AddItem:input_type -> basketspb.AddItemRequest
	8,  // 5: basketspb.BasketService.RemoveItem:input_type -> basketspb.RemoveItemRequest
	10, // 6: basketspb.BasketService.GetBasket:input_type -> basketspb.GetBasketRequest
	1,  // 7: basketspb.BasketService.StartBasket:output_type -> basketspb.StartBasketResponse
	3,  // 8: basketspb.BasketService.CancelBasket:output_type -> basketspb.CancelBasketResponse
	5,  // 9: basketspb.BasketService.CheckoutBasket:output_type -> basketspb.CheckoutBasketResponse
	7,  // 10: basketspb.BasketService.AddItem:output_type -> basketspb.AddItemResponse
	9,  // 11: basketspb.BasketService.RemoveItem:output_type -> basketspb.RemoveItemResponse
	11, // 12: basketspb.BasketService.GetBasket:output_type -> basketspb.GetBasketResponse
	7,  // [7:13] is the sub-list for method output_type
	1,  // [1:7] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_basketspb_api_proto_init() }
func file_basketspb_api_proto_init() {
	if File_basketspb_api_proto != nil {
		return
	}
	file_basketspb_messages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_basketspb_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartBasketRequest); i {
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
		file_basketspb_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartBasketResponse); i {
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
		file_basketspb_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelBasketRequest); i {
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
		file_basketspb_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelBasketResponse); i {
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
		file_basketspb_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckoutBasketRequest); i {
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
		file_basketspb_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckoutBasketResponse); i {
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
		file_basketspb_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddItemRequest); i {
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
		file_basketspb_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddItemResponse); i {
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
		file_basketspb_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveItemRequest); i {
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
		file_basketspb_api_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveItemResponse); i {
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
		file_basketspb_api_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBasketRequest); i {
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
		file_basketspb_api_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBasketResponse); i {
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
			RawDescriptor: file_basketspb_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_basketspb_api_proto_goTypes,
		DependencyIndexes: file_basketspb_api_proto_depIdxs,
		MessageInfos:      file_basketspb_api_proto_msgTypes,
	}.Build()
	File_basketspb_api_proto = out.File
	file_basketspb_api_proto_rawDesc = nil
	file_basketspb_api_proto_goTypes = nil
	file_basketspb_api_proto_depIdxs = nil
}
