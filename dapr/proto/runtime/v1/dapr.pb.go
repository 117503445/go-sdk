// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dapr/proto/runtime/v1/dapr.proto

package runtime

import (
	context "context"
	fmt "fmt"
	v1 "github.com/dapr/go-sdk/dapr/proto/common/v1"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// InvokeServiceRequest represents the request message for Service invocation.
type InvokeServiceRequest struct {
	// Required. Callee's app id.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Required. message which will be delivered to callee.
	Message              *v1.InvokeRequest `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *InvokeServiceRequest) Reset()         { *m = InvokeServiceRequest{} }
func (m *InvokeServiceRequest) String() string { return proto.CompactTextString(m) }
func (*InvokeServiceRequest) ProtoMessage()    {}
func (*InvokeServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_da511bac0105b1e5, []int{0}
}

func (m *InvokeServiceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvokeServiceRequest.Unmarshal(m, b)
}
func (m *InvokeServiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvokeServiceRequest.Marshal(b, m, deterministic)
}
func (m *InvokeServiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvokeServiceRequest.Merge(m, src)
}
func (m *InvokeServiceRequest) XXX_Size() int {
	return xxx_messageInfo_InvokeServiceRequest.Size(m)
}
func (m *InvokeServiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InvokeServiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InvokeServiceRequest proto.InternalMessageInfo

func (m *InvokeServiceRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *InvokeServiceRequest) GetMessage() *v1.InvokeRequest {
	if m != nil {
		return m.Message
	}
	return nil
}

// GetStateRequest is the message to get key-value states from specific state store.
type GetStateRequest struct {
	// The name of state store.
	StoreName string `protobuf:"bytes,1,opt,name=store_name,json=storeName,proto3" json:"store_name,omitempty"`
	// The key of the desired state
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	// The read consistency of the state store.
	Consistency          v1.StateOptions_StateConsistency `protobuf:"varint,3,opt,name=consistency,proto3,enum=dapr.proto.common.v1.StateOptions_StateConsistency" json:"consistency,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *GetStateRequest) Reset()         { *m = GetStateRequest{} }
func (m *GetStateRequest) String() string { return proto.CompactTextString(m) }
func (*GetStateRequest) ProtoMessage()    {}
func (*GetStateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_da511bac0105b1e5, []int{1}
}

func (m *GetStateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStateRequest.Unmarshal(m, b)
}
func (m *GetStateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStateRequest.Marshal(b, m, deterministic)
}
func (m *GetStateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStateRequest.Merge(m, src)
}
func (m *GetStateRequest) XXX_Size() int {
	return xxx_messageInfo_GetStateRequest.Size(m)
}
func (m *GetStateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetStateRequest proto.InternalMessageInfo

func (m *GetStateRequest) GetStoreName() string {
	if m != nil {
		return m.StoreName
	}
	return ""
}

func (m *GetStateRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *GetStateRequest) GetConsistency() v1.StateOptions_StateConsistency {
	if m != nil {
		return m.Consistency
	}
	return v1.StateOptions_CONSISTENCY_UNSPECIFIED
}

// GetStateResponse is the response conveying the state value and etag.
type GetStateResponse struct {
	// The byte array data
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	// The entity tag which represents the specific version of data.
	// ETag format is defined by the corresponding data store.
	Etag                 string   `protobuf:"bytes,2,opt,name=etag,proto3" json:"etag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetStateResponse) Reset()         { *m = GetStateResponse{} }
func (m *GetStateResponse) String() string { return proto.CompactTextString(m) }
func (*GetStateResponse) ProtoMessage()    {}
func (*GetStateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_da511bac0105b1e5, []int{2}
}

func (m *GetStateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStateResponse.Unmarshal(m, b)
}
func (m *GetStateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStateResponse.Marshal(b, m, deterministic)
}
func (m *GetStateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStateResponse.Merge(m, src)
}
func (m *GetStateResponse) XXX_Size() int {
	return xxx_messageInfo_GetStateResponse.Size(m)
}
func (m *GetStateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetStateResponse proto.InternalMessageInfo

func (m *GetStateResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *GetStateResponse) GetEtag() string {
	if m != nil {
		return m.Etag
	}
	return ""
}

// DeleteStateRequest is the message to delete key-value states in the specific state store.
type DeleteStateRequest struct {
	// The name of state store.
	StoreName string `protobuf:"bytes,1,opt,name=store_name,json=storeName,proto3" json:"store_name,omitempty"`
	// The key of the desired state
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	// The entity tag which represents the specific version of data.
	// The exact ETag format is defined by the corresponding data store.
	Etag string `protobuf:"bytes,3,opt,name=etag,proto3" json:"etag,omitempty"`
	// State operation options which includes concurrency/
	// consistency/retry_policy.
	Options              *v1.StateOptions `protobuf:"bytes,4,opt,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *DeleteStateRequest) Reset()         { *m = DeleteStateRequest{} }
func (m *DeleteStateRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteStateRequest) ProtoMessage()    {}
func (*DeleteStateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_da511bac0105b1e5, []int{3}
}

func (m *DeleteStateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteStateRequest.Unmarshal(m, b)
}
func (m *DeleteStateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteStateRequest.Marshal(b, m, deterministic)
}
func (m *DeleteStateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteStateRequest.Merge(m, src)
}
func (m *DeleteStateRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteStateRequest.Size(m)
}
func (m *DeleteStateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteStateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteStateRequest proto.InternalMessageInfo

func (m *DeleteStateRequest) GetStoreName() string {
	if m != nil {
		return m.StoreName
	}
	return ""
}

func (m *DeleteStateRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *DeleteStateRequest) GetEtag() string {
	if m != nil {
		return m.Etag
	}
	return ""
}

func (m *DeleteStateRequest) GetOptions() *v1.StateOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

// SaveStateRequest is the message to save multiple states into state store.
type SaveStateRequest struct {
	// The name of state store.
	StoreName string `protobuf:"bytes,1,opt,name=store_name,json=storeName,proto3" json:"store_name,omitempty"`
	// The array of the state key values.
	States               []*v1.StateItem `protobuf:"bytes,2,rep,name=states,proto3" json:"states,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *SaveStateRequest) Reset()         { *m = SaveStateRequest{} }
func (m *SaveStateRequest) String() string { return proto.CompactTextString(m) }
func (*SaveStateRequest) ProtoMessage()    {}
func (*SaveStateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_da511bac0105b1e5, []int{4}
}

func (m *SaveStateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveStateRequest.Unmarshal(m, b)
}
func (m *SaveStateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveStateRequest.Marshal(b, m, deterministic)
}
func (m *SaveStateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveStateRequest.Merge(m, src)
}
func (m *SaveStateRequest) XXX_Size() int {
	return xxx_messageInfo_SaveStateRequest.Size(m)
}
func (m *SaveStateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveStateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SaveStateRequest proto.InternalMessageInfo

func (m *SaveStateRequest) GetStoreName() string {
	if m != nil {
		return m.StoreName
	}
	return ""
}

func (m *SaveStateRequest) GetStates() []*v1.StateItem {
	if m != nil {
		return m.States
	}
	return nil
}

// PublishEventRequest is the message to publish event data to pubsub topic
type PublishEventRequest struct {
	// The pubsub topic
	Topic string `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	// The data which will be published to topic.
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublishEventRequest) Reset()         { *m = PublishEventRequest{} }
func (m *PublishEventRequest) String() string { return proto.CompactTextString(m) }
func (*PublishEventRequest) ProtoMessage()    {}
func (*PublishEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_da511bac0105b1e5, []int{5}
}

func (m *PublishEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishEventRequest.Unmarshal(m, b)
}
func (m *PublishEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishEventRequest.Marshal(b, m, deterministic)
}
func (m *PublishEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishEventRequest.Merge(m, src)
}
func (m *PublishEventRequest) XXX_Size() int {
	return xxx_messageInfo_PublishEventRequest.Size(m)
}
func (m *PublishEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PublishEventRequest proto.InternalMessageInfo

func (m *PublishEventRequest) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *PublishEventRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// InvokeBindingRequest is the message to send data to output bindings
type InvokeBindingRequest struct {
	// The name of the output binding to invoke.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The data which will be sent to output binding.
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	// The metadata passing to output binding components
	//
	// Common metadata property:
	// - ttlInSeconds : the time to live in seconds for the message.
	// If set in the binding definition will cause all messages to
	// have a default time to live. The message ttl overrides any value
	// in the binding definition.
	Metadata map[string]string `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The name of the operation type for the binding to invoke
	Operation            string   `protobuf:"bytes,4,opt,name=operation,proto3" json:"operation,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InvokeBindingRequest) Reset()         { *m = InvokeBindingRequest{} }
func (m *InvokeBindingRequest) String() string { return proto.CompactTextString(m) }
func (*InvokeBindingRequest) ProtoMessage()    {}
func (*InvokeBindingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_da511bac0105b1e5, []int{6}
}

func (m *InvokeBindingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvokeBindingRequest.Unmarshal(m, b)
}
func (m *InvokeBindingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvokeBindingRequest.Marshal(b, m, deterministic)
}
func (m *InvokeBindingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvokeBindingRequest.Merge(m, src)
}
func (m *InvokeBindingRequest) XXX_Size() int {
	return xxx_messageInfo_InvokeBindingRequest.Size(m)
}
func (m *InvokeBindingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InvokeBindingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InvokeBindingRequest proto.InternalMessageInfo

func (m *InvokeBindingRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InvokeBindingRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *InvokeBindingRequest) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *InvokeBindingRequest) GetOperation() string {
	if m != nil {
		return m.Operation
	}
	return ""
}

// InvokeBindingResponse is the message returned from an output binding invocation
type InvokeBindingResponse struct {
	// The data which will be sent to output binding.
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	// The metadata returned from an external system
	Metadata             map[string]string `protobuf:"bytes,2,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *InvokeBindingResponse) Reset()         { *m = InvokeBindingResponse{} }
func (m *InvokeBindingResponse) String() string { return proto.CompactTextString(m) }
func (*InvokeBindingResponse) ProtoMessage()    {}
func (*InvokeBindingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_da511bac0105b1e5, []int{7}
}

func (m *InvokeBindingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvokeBindingResponse.Unmarshal(m, b)
}
func (m *InvokeBindingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvokeBindingResponse.Marshal(b, m, deterministic)
}
func (m *InvokeBindingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvokeBindingResponse.Merge(m, src)
}
func (m *InvokeBindingResponse) XXX_Size() int {
	return xxx_messageInfo_InvokeBindingResponse.Size(m)
}
func (m *InvokeBindingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InvokeBindingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InvokeBindingResponse proto.InternalMessageInfo

func (m *InvokeBindingResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *InvokeBindingResponse) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// GetSecretRequest is the message to get secret from secret store.
type GetSecretRequest struct {
	// The name of secret store.
	StoreName string `protobuf:"bytes,1,opt,name=store_name,json=storeName,proto3" json:"store_name,omitempty"`
	// The name of secret key.
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	// The metadata which will be sent to secret store components.
	Metadata             map[string]string `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GetSecretRequest) Reset()         { *m = GetSecretRequest{} }
func (m *GetSecretRequest) String() string { return proto.CompactTextString(m) }
func (*GetSecretRequest) ProtoMessage()    {}
func (*GetSecretRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_da511bac0105b1e5, []int{8}
}

func (m *GetSecretRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSecretRequest.Unmarshal(m, b)
}
func (m *GetSecretRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSecretRequest.Marshal(b, m, deterministic)
}
func (m *GetSecretRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSecretRequest.Merge(m, src)
}
func (m *GetSecretRequest) XXX_Size() int {
	return xxx_messageInfo_GetSecretRequest.Size(m)
}
func (m *GetSecretRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSecretRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetSecretRequest proto.InternalMessageInfo

func (m *GetSecretRequest) GetStoreName() string {
	if m != nil {
		return m.StoreName
	}
	return ""
}

func (m *GetSecretRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *GetSecretRequest) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// GetSecretResponse is the response mesage to convey the requested secret.
type GetSecretResponse struct {
	// data is the secret value. Some secret store, such as kubernetes secret
	// store, can save multiple secrets for single secret key.
	Data                 map[string]string `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GetSecretResponse) Reset()         { *m = GetSecretResponse{} }
func (m *GetSecretResponse) String() string { return proto.CompactTextString(m) }
func (*GetSecretResponse) ProtoMessage()    {}
func (*GetSecretResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_da511bac0105b1e5, []int{9}
}

func (m *GetSecretResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSecretResponse.Unmarshal(m, b)
}
func (m *GetSecretResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSecretResponse.Marshal(b, m, deterministic)
}
func (m *GetSecretResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSecretResponse.Merge(m, src)
}
func (m *GetSecretResponse) XXX_Size() int {
	return xxx_messageInfo_GetSecretResponse.Size(m)
}
func (m *GetSecretResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSecretResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetSecretResponse proto.InternalMessageInfo

func (m *GetSecretResponse) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*InvokeServiceRequest)(nil), "dapr.proto.runtime.v1.InvokeServiceRequest")
	proto.RegisterType((*GetStateRequest)(nil), "dapr.proto.runtime.v1.GetStateRequest")
	proto.RegisterType((*GetStateResponse)(nil), "dapr.proto.runtime.v1.GetStateResponse")
	proto.RegisterType((*DeleteStateRequest)(nil), "dapr.proto.runtime.v1.DeleteStateRequest")
	proto.RegisterType((*SaveStateRequest)(nil), "dapr.proto.runtime.v1.SaveStateRequest")
	proto.RegisterType((*PublishEventRequest)(nil), "dapr.proto.runtime.v1.PublishEventRequest")
	proto.RegisterType((*InvokeBindingRequest)(nil), "dapr.proto.runtime.v1.InvokeBindingRequest")
	proto.RegisterMapType((map[string]string)(nil), "dapr.proto.runtime.v1.InvokeBindingRequest.MetadataEntry")
	proto.RegisterType((*InvokeBindingResponse)(nil), "dapr.proto.runtime.v1.InvokeBindingResponse")
	proto.RegisterMapType((map[string]string)(nil), "dapr.proto.runtime.v1.InvokeBindingResponse.MetadataEntry")
	proto.RegisterType((*GetSecretRequest)(nil), "dapr.proto.runtime.v1.GetSecretRequest")
	proto.RegisterMapType((map[string]string)(nil), "dapr.proto.runtime.v1.GetSecretRequest.MetadataEntry")
	proto.RegisterType((*GetSecretResponse)(nil), "dapr.proto.runtime.v1.GetSecretResponse")
	proto.RegisterMapType((map[string]string)(nil), "dapr.proto.runtime.v1.GetSecretResponse.DataEntry")
}

func init() { proto.RegisterFile("dapr/proto/runtime/v1/dapr.proto", fileDescriptor_da511bac0105b1e5) }

var fileDescriptor_da511bac0105b1e5 = []byte{
	// 764 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0xeb, 0x6e, 0xd3, 0x4a,
	0x10, 0x8e, 0x93, 0xf4, 0x92, 0x49, 0xdb, 0xd3, 0xb3, 0xa7, 0x3d, 0x8a, 0xdc, 0x73, 0x44, 0x30,
	0x88, 0x86, 0x8b, 0x36, 0x4a, 0x2a, 0x54, 0x68, 0x41, 0x88, 0x36, 0xa5, 0xea, 0x0f, 0x4a, 0x71,
	0x29, 0x3f, 0x90, 0x10, 0x38, 0xc9, 0xe0, 0x9a, 0xc6, 0x5e, 0x63, 0x6f, 0x2c, 0xe5, 0x3d, 0x40,
	0x88, 0x57, 0xe0, 0x29, 0x78, 0x07, 0x5e, 0x84, 0x47, 0x40, 0xde, 0xb5, 0x1d, 0xa7, 0xcd, 0xa5,
	0x17, 0xf1, 0x27, 0x1a, 0x8f, 0x66, 0xbe, 0xf9, 0x66, 0x67, 0xf6, 0xdb, 0x40, 0xb9, 0x6d, 0xb8,
	0x5e, 0xd5, 0xf5, 0x18, 0x67, 0x55, 0xaf, 0xeb, 0x70, 0xcb, 0xc6, 0x6a, 0x50, 0xab, 0x86, 0x5e,
	0x2a, 0xbc, 0x64, 0xb9, 0x6f, 0xd3, 0x28, 0x82, 0x06, 0x35, 0x75, 0xc5, 0x64, 0xcc, 0xec, 0xa0,
	0x4c, 0x6d, 0x76, 0x3f, 0x54, 0xd1, 0x76, 0x79, 0x4f, 0xc6, 0xa9, 0xd7, 0x53, 0xa8, 0x2d, 0x66,
	0xdb, 0xcc, 0x09, 0x41, 0xa5, 0x25, 0x43, 0x34, 0x84, 0xa5, 0x3d, 0x27, 0x60, 0x27, 0x78, 0x88,
	0x5e, 0x60, 0xb5, 0x50, 0xc7, 0x4f, 0x5d, 0xf4, 0x39, 0x59, 0x80, 0xac, 0xd5, 0x2e, 0x29, 0x65,
	0xa5, 0x52, 0xd0, 0xb3, 0x56, 0x9b, 0x3c, 0x86, 0x19, 0x1b, 0x7d, 0xdf, 0x30, 0xb1, 0x94, 0x2b,
	0x2b, 0x95, 0x62, 0xfd, 0x06, 0x4d, 0x11, 0x8a, 0x20, 0x83, 0x1a, 0x95, 0x60, 0x11, 0x8a, 0x1e,
	0xe7, 0x68, 0xdf, 0x14, 0xf8, 0x6b, 0x17, 0xf9, 0x21, 0x37, 0x78, 0x52, 0xe2, 0x7f, 0x00, 0x9f,
	0x33, 0x0f, 0xdf, 0x39, 0x86, 0x8d, 0x51, 0xa9, 0x82, 0xf0, 0xec, 0x1b, 0x36, 0x92, 0x45, 0xc8,
	0x9d, 0x60, 0xaf, 0x94, 0x15, 0xfe, 0xd0, 0x24, 0x47, 0x50, 0x6c, 0x31, 0xc7, 0xb7, 0x7c, 0x8e,
	0x4e, 0xab, 0x27, 0x78, 0x2c, 0xd4, 0xd7, 0x86, 0xf3, 0x10, 0x95, 0x5e, 0xb8, 0xdc, 0x62, 0x8e,
	0x2f, 0x3f, 0xb6, 0xfb, 0xa9, 0x7a, 0x1a, 0x47, 0xdb, 0x80, 0xc5, 0x3e, 0x35, 0xdf, 0x65, 0x8e,
	0x8f, 0x84, 0x40, 0xbe, 0x6d, 0x70, 0x43, 0xb0, 0x9a, 0xd3, 0x85, 0x1d, 0xfa, 0x90, 0x1b, 0x66,
	0xc4, 0x48, 0xd8, 0xda, 0x57, 0x05, 0x48, 0x03, 0x3b, 0xc8, 0xf1, 0x6a, 0xad, 0xc5, 0xd8, 0xb9,
	0x3e, 0x36, 0x79, 0x04, 0x33, 0x4c, 0x36, 0x50, 0xca, 0x8b, 0x23, 0xd7, 0x26, 0xb7, 0xaa, 0xc7,
	0x29, 0xda, 0x47, 0x58, 0x3c, 0x34, 0x82, 0x0b, 0xd1, 0x5a, 0x87, 0x69, 0x3f, 0x0c, 0xf7, 0x4b,
	0xd9, 0x72, 0xae, 0x52, 0xac, 0x5f, 0x1b, 0x53, 0x6f, 0x8f, 0xa3, 0xad, 0x47, 0xe1, 0xda, 0x13,
	0xf8, 0xe7, 0xa0, 0xdb, 0xec, 0x58, 0xfe, 0xf1, 0x4e, 0x80, 0x0e, 0x8f, 0xcb, 0x2d, 0xc1, 0x14,
	0x67, 0xae, 0xd5, 0x8a, 0x2a, 0xc9, 0x8f, 0xe4, 0x68, 0xb3, 0xfd, 0xa3, 0xd5, 0x7e, 0x29, 0xf1,
	0x1a, 0x6e, 0x59, 0x4e, 0xdb, 0x72, 0xcc, 0x18, 0x82, 0x40, 0x3e, 0xc5, 0x55, 0xd8, 0xc3, 0x00,
	0xc8, 0x11, 0xcc, 0xda, 0xc8, 0x0d, 0xe1, 0xcf, 0x09, 0xf2, 0x0f, 0xe9, 0xd0, 0x0b, 0x43, 0x87,
	0x95, 0xa1, 0xcf, 0xa3, 0xdc, 0x1d, 0x87, 0x7b, 0x3d, 0x3d, 0x81, 0x22, 0xff, 0x41, 0x81, 0xb9,
	0xe8, 0x19, 0xe1, 0x91, 0x8a, 0x21, 0x14, 0xf4, 0xbe, 0x43, 0xdd, 0x84, 0xf9, 0x81, 0xc4, 0x78,
	0xae, 0x4a, 0x7f, 0xae, 0x4b, 0x30, 0x15, 0x18, 0x9d, 0x2e, 0x46, 0xb3, 0x96, 0x1f, 0x1b, 0xd9,
	0x07, 0x8a, 0xf6, 0x43, 0x81, 0xe5, 0x53, 0x5c, 0xc6, 0xec, 0xde, 0xeb, 0x54, 0x7f, 0x72, 0x38,
	0x1b, 0xe7, 0xeb, 0x4f, 0x62, 0x8e, 0x6a, 0xf0, 0x6a, 0x2d, 0xfc, 0x54, 0xe4, 0xcd, 0xc1, 0x96,
	0x87, 0xfc, 0xd2, 0xab, 0xff, 0xf2, 0xcc, 0xe8, 0xee, 0x8f, 0x68, 0xed, 0x74, 0xad, 0x3f, 0xd3,
	0xd5, 0x67, 0x05, 0xfe, 0x4e, 0x55, 0x8a, 0x86, 0xf2, 0x2c, 0x19, 0x4a, 0xc8, 0xb0, 0x3e, 0x99,
	0x61, 0x74, 0xf0, 0x8d, 0x84, 0x9e, 0xc8, 0x57, 0xd7, 0xa1, 0xd0, 0xb8, 0x0c, 0xad, 0xfa, 0x97,
	0x29, 0xc8, 0x37, 0x0c, 0xd7, 0x23, 0x6d, 0x98, 0x1f, 0x50, 0x6c, 0x72, 0x77, 0xec, 0x26, 0x0c,
	0xea, 0xba, 0x7a, 0x73, 0xbc, 0x6c, 0x4b, 0xd6, 0x5a, 0x86, 0xbc, 0x85, 0xd9, 0x58, 0x14, 0xc9,
	0xad, 0x31, 0xdd, 0xa6, 0xe4, 0x45, 0x5d, 0x9d, 0x18, 0x97, 0xc0, 0xef, 0x43, 0x21, 0x51, 0x27,
	0x32, 0x2a, 0xef, 0xb4, 0x7e, 0xa9, 0xff, 0x52, 0xf9, 0xda, 0xd1, 0xf8, 0xb5, 0xa3, 0x3b, 0xe1,
	0x6b, 0xa7, 0x65, 0x88, 0x0e, 0xc5, 0x94, 0x0c, 0x93, 0xdb, 0x23, 0x10, 0xcf, 0x4a, 0xf5, 0x18,
	0xcc, 0x57, 0x30, 0x97, 0x56, 0x35, 0x72, 0x67, 0x04, 0xe8, 0x10, 0xe9, 0x1b, 0x83, 0xda, 0x89,
	0xc7, 0x17, 0x5d, 0xd1, 0x09, 0xe3, 0x1b, 0x14, 0x2a, 0xf5, 0xde, 0x45, 0x6e, 0xbd, 0x96, 0x21,
	0xef, 0xa1, 0x90, 0xec, 0x24, 0x59, 0x3d, 0xe7, 0xbd, 0x52, 0x2b, 0xe7, 0x5d, 0x6f, 0x2d, 0xb3,
	0x65, 0x01, 0x58, 0x4c, 0xc6, 0x07, 0xb5, 0x2d, 0x08, 0x57, 0xf4, 0x20, 0xcc, 0xf3, 0xdf, 0xd4,
	0x4c, 0x8b, 0x1f, 0x77, 0x9b, 0xe1, 0x96, 0x89, 0x3f, 0x32, 0xf2, 0xc7, 0x3d, 0x31, 0xcf, 0xfc,
	0xcf, 0xd9, 0x8c, 0xcc, 0xef, 0xd9, 0x95, 0x30, 0x9f, 0x6e, 0x77, 0x2c, 0x74, 0x38, 0x7d, 0xda,
	0xe5, 0xcc, 0x44, 0x87, 0xee, 0x7a, 0x6e, 0x8b, 0x06, 0xb5, 0xe6, 0xb4, 0xc8, 0x5b, 0xfb, 0x1d,
	0x00, 0x00, 0xff, 0xff, 0xe9, 0x06, 0x92, 0xcb, 0x2d, 0x09, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DaprClient is the client API for Dapr service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DaprClient interface {
	// Invokes a method on a remote Dapr app.
	InvokeService(ctx context.Context, in *InvokeServiceRequest, opts ...grpc.CallOption) (*v1.InvokeResponse, error)
	// Gets the state for a specific key.
	GetState(ctx context.Context, in *GetStateRequest, opts ...grpc.CallOption) (*GetStateResponse, error)
	// Saves the state for a specific key.
	SaveState(ctx context.Context, in *SaveStateRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Deletes the state for a specific key.
	DeleteState(ctx context.Context, in *DeleteStateRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Publishes events to the specific topic.
	PublishEvent(ctx context.Context, in *PublishEventRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Invokes binding data to specific output bindings
	InvokeBinding(ctx context.Context, in *InvokeBindingRequest, opts ...grpc.CallOption) (*InvokeBindingResponse, error)
	// Gets secrets from secret stores.
	GetSecret(ctx context.Context, in *GetSecretRequest, opts ...grpc.CallOption) (*GetSecretResponse, error)
}

type daprClient struct {
	cc *grpc.ClientConn
}

func NewDaprClient(cc *grpc.ClientConn) DaprClient {
	return &daprClient{cc}
}

func (c *daprClient) InvokeService(ctx context.Context, in *InvokeServiceRequest, opts ...grpc.CallOption) (*v1.InvokeResponse, error) {
	out := new(v1.InvokeResponse)
	err := c.cc.Invoke(ctx, "/dapr.proto.runtime.v1.Dapr/InvokeService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daprClient) GetState(ctx context.Context, in *GetStateRequest, opts ...grpc.CallOption) (*GetStateResponse, error) {
	out := new(GetStateResponse)
	err := c.cc.Invoke(ctx, "/dapr.proto.runtime.v1.Dapr/GetState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daprClient) SaveState(ctx context.Context, in *SaveStateRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/dapr.proto.runtime.v1.Dapr/SaveState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daprClient) DeleteState(ctx context.Context, in *DeleteStateRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/dapr.proto.runtime.v1.Dapr/DeleteState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daprClient) PublishEvent(ctx context.Context, in *PublishEventRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/dapr.proto.runtime.v1.Dapr/PublishEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daprClient) InvokeBinding(ctx context.Context, in *InvokeBindingRequest, opts ...grpc.CallOption) (*InvokeBindingResponse, error) {
	out := new(InvokeBindingResponse)
	err := c.cc.Invoke(ctx, "/dapr.proto.runtime.v1.Dapr/InvokeBinding", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daprClient) GetSecret(ctx context.Context, in *GetSecretRequest, opts ...grpc.CallOption) (*GetSecretResponse, error) {
	out := new(GetSecretResponse)
	err := c.cc.Invoke(ctx, "/dapr.proto.runtime.v1.Dapr/GetSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DaprServer is the server API for Dapr service.
type DaprServer interface {
	// Invokes a method on a remote Dapr app.
	InvokeService(context.Context, *InvokeServiceRequest) (*v1.InvokeResponse, error)
	// Gets the state for a specific key.
	GetState(context.Context, *GetStateRequest) (*GetStateResponse, error)
	// Saves the state for a specific key.
	SaveState(context.Context, *SaveStateRequest) (*empty.Empty, error)
	// Deletes the state for a specific key.
	DeleteState(context.Context, *DeleteStateRequest) (*empty.Empty, error)
	// Publishes events to the specific topic.
	PublishEvent(context.Context, *PublishEventRequest) (*empty.Empty, error)
	// Invokes binding data to specific output bindings
	InvokeBinding(context.Context, *InvokeBindingRequest) (*InvokeBindingResponse, error)
	// Gets secrets from secret stores.
	GetSecret(context.Context, *GetSecretRequest) (*GetSecretResponse, error)
}

// UnimplementedDaprServer can be embedded to have forward compatible implementations.
type UnimplementedDaprServer struct {
}

func (*UnimplementedDaprServer) InvokeService(ctx context.Context, req *InvokeServiceRequest) (*v1.InvokeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InvokeService not implemented")
}
func (*UnimplementedDaprServer) GetState(ctx context.Context, req *GetStateRequest) (*GetStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetState not implemented")
}
func (*UnimplementedDaprServer) SaveState(ctx context.Context, req *SaveStateRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveState not implemented")
}
func (*UnimplementedDaprServer) DeleteState(ctx context.Context, req *DeleteStateRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteState not implemented")
}
func (*UnimplementedDaprServer) PublishEvent(ctx context.Context, req *PublishEventRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishEvent not implemented")
}
func (*UnimplementedDaprServer) InvokeBinding(ctx context.Context, req *InvokeBindingRequest) (*InvokeBindingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InvokeBinding not implemented")
}
func (*UnimplementedDaprServer) GetSecret(ctx context.Context, req *GetSecretRequest) (*GetSecretResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSecret not implemented")
}

func RegisterDaprServer(s *grpc.Server, srv DaprServer) {
	s.RegisterService(&_Dapr_serviceDesc, srv)
}

func _Dapr_InvokeService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvokeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaprServer).InvokeService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dapr.proto.runtime.v1.Dapr/InvokeService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaprServer).InvokeService(ctx, req.(*InvokeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dapr_GetState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaprServer).GetState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dapr.proto.runtime.v1.Dapr/GetState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaprServer).GetState(ctx, req.(*GetStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dapr_SaveState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaprServer).SaveState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dapr.proto.runtime.v1.Dapr/SaveState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaprServer).SaveState(ctx, req.(*SaveStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dapr_DeleteState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaprServer).DeleteState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dapr.proto.runtime.v1.Dapr/DeleteState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaprServer).DeleteState(ctx, req.(*DeleteStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dapr_PublishEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaprServer).PublishEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dapr.proto.runtime.v1.Dapr/PublishEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaprServer).PublishEvent(ctx, req.(*PublishEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dapr_InvokeBinding_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvokeBindingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaprServer).InvokeBinding(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dapr.proto.runtime.v1.Dapr/InvokeBinding",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaprServer).InvokeBinding(ctx, req.(*InvokeBindingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dapr_GetSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaprServer).GetSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dapr.proto.runtime.v1.Dapr/GetSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaprServer).GetSecret(ctx, req.(*GetSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Dapr_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dapr.proto.runtime.v1.Dapr",
	HandlerType: (*DaprServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InvokeService",
			Handler:    _Dapr_InvokeService_Handler,
		},
		{
			MethodName: "GetState",
			Handler:    _Dapr_GetState_Handler,
		},
		{
			MethodName: "SaveState",
			Handler:    _Dapr_SaveState_Handler,
		},
		{
			MethodName: "DeleteState",
			Handler:    _Dapr_DeleteState_Handler,
		},
		{
			MethodName: "PublishEvent",
			Handler:    _Dapr_PublishEvent_Handler,
		},
		{
			MethodName: "InvokeBinding",
			Handler:    _Dapr_InvokeBinding_Handler,
		},
		{
			MethodName: "GetSecret",
			Handler:    _Dapr_GetSecret_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dapr/proto/runtime/v1/dapr.proto",
}
