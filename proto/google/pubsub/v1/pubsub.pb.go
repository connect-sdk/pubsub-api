// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: google/pubsub/v1/pubsub.proto

package pubsubv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// PushPubsubMessageRequest represents a request for google.pubsub.v1.PubsubService.PushPubsubMessage method.
type PushPubsubMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The subscription from which messages should be pulled.
	Subscription string `protobuf:"bytes,1,opt,name=subscription,proto3" json:"subscription,omitempty"`
	// The message.
	Message *PubsubMessage `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PushPubsubMessageRequest) Reset() {
	*x = PushPubsubMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_pubsub_v1_pubsub_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushPubsubMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushPubsubMessageRequest) ProtoMessage() {}

func (x *PushPubsubMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_pubsub_v1_pubsub_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushPubsubMessageRequest.ProtoReflect.Descriptor instead.
func (*PushPubsubMessageRequest) Descriptor() ([]byte, []int) {
	return file_google_pubsub_v1_pubsub_proto_rawDescGZIP(), []int{0}
}

func (x *PushPubsubMessageRequest) GetSubscription() string {
	if x != nil {
		return x.Subscription
	}
	return ""
}

func (x *PushPubsubMessageRequest) GetMessage() *PubsubMessage {
	if x != nil {
		return x.Message
	}
	return nil
}

// PushPubsubMessageResponse represents a response for google.pubsub.v1.PubsubService.PushPubsubMessage method.
type PushPubsubMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PushPubsubMessageResponse) Reset() {
	*x = PushPubsubMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_pubsub_v1_pubsub_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushPubsubMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushPubsubMessageResponse) ProtoMessage() {}

func (x *PushPubsubMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_pubsub_v1_pubsub_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushPubsubMessageResponse.ProtoReflect.Descriptor instead.
func (*PushPubsubMessageResponse) Descriptor() ([]byte, []int) {
	return file_google_pubsub_v1_pubsub_proto_rawDescGZIP(), []int{1}
}

// A message data and its attributes.
type PubsubMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The message payload. For JSON requests, the value of this field must be
	// base64-encoded.
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	// Attributes for this message.
	Attributes map[string]string `protobuf:"bytes,2,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *PubsubMessage) Reset() {
	*x = PubsubMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_pubsub_v1_pubsub_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PubsubMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PubsubMessage) ProtoMessage() {}

func (x *PubsubMessage) ProtoReflect() protoreflect.Message {
	mi := &file_google_pubsub_v1_pubsub_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PubsubMessage.ProtoReflect.Descriptor instead.
func (*PubsubMessage) Descriptor() ([]byte, []int) {
	return file_google_pubsub_v1_pubsub_proto_rawDescGZIP(), []int{2}
}

func (x *PubsubMessage) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *PubsubMessage) GetAttributes() map[string]string {
	if x != nil {
		return x.Attributes
	}
	return nil
}

var File_google_pubsub_v1_pubsub_proto protoreflect.FileDescriptor

var file_google_pubsub_v1_pubsub_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x76,
	0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x01, 0x0a, 0x18,
	0x50, 0x75, 0x73, 0x68, 0x50, 0x75, 0x62, 0x73, 0x75, 0x62, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x02, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x46, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x75, 0x62, 0x73,
	0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x62, 0x73, 0x75, 0x62, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x42, 0x0b, 0xe0, 0x41, 0x02, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x1b, 0x0a, 0x19, 0x50, 0x75, 0x73,
	0x68, 0x50, 0x75, 0x62, 0x73, 0x75, 0x62, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xc3, 0x01, 0x0a, 0x0d, 0x50, 0x75, 0x62, 0x73, 0x75,
	0x62, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x17, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x54, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x75, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x62, 0x73, 0x75, 0x62, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0a, 0x61, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x1a, 0x3d, 0x0a, 0x0f, 0x41, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x4a, 0x04, 0x08, 0x03, 0x10, 0x04, 0x32, 0x7d, 0x0a, 0x0d,
	0x50, 0x75, 0x62, 0x73, 0x75, 0x62, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6c, 0x0a,
	0x11, 0x50, 0x75, 0x73, 0x68, 0x50, 0x75, 0x62, 0x73, 0x75, 0x62, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x75, 0x62, 0x73,
	0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x50, 0x75, 0x62, 0x73, 0x75, 0x62,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x50, 0x75, 0x62, 0x73, 0x75, 0x62, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xc8, 0x01, 0x0a, 0x14,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x75, 0x62, 0x73, 0x75,
	0x62, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x50, 0x75, 0x62, 0x73, 0x75, 0x62, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x75, 0x62, 0x73,
	0x75, 0x62, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x75,
	0x62, 0x73, 0x75, 0x62, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x47, 0x50, 0x58, 0xaa, 0x02, 0x10, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x50, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x10, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x50, 0x75, 0x62, 0x73, 0x75, 0x62, 0x5c,
	0x56, 0x31, 0xe2, 0x02, 0x1c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x50, 0x75, 0x62, 0x73,
	0x75, 0x62, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x12, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x50, 0x75, 0x62, 0x73,
	0x75, 0x62, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_pubsub_v1_pubsub_proto_rawDescOnce sync.Once
	file_google_pubsub_v1_pubsub_proto_rawDescData = file_google_pubsub_v1_pubsub_proto_rawDesc
)

func file_google_pubsub_v1_pubsub_proto_rawDescGZIP() []byte {
	file_google_pubsub_v1_pubsub_proto_rawDescOnce.Do(func() {
		file_google_pubsub_v1_pubsub_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_pubsub_v1_pubsub_proto_rawDescData)
	})
	return file_google_pubsub_v1_pubsub_proto_rawDescData
}

var file_google_pubsub_v1_pubsub_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_pubsub_v1_pubsub_proto_goTypes = []interface{}{
	(*PushPubsubMessageRequest)(nil),  // 0: google.pubsub.v1.PushPubsubMessageRequest
	(*PushPubsubMessageResponse)(nil), // 1: google.pubsub.v1.PushPubsubMessageResponse
	(*PubsubMessage)(nil),             // 2: google.pubsub.v1.PubsubMessage
	nil,                               // 3: google.pubsub.v1.PubsubMessage.AttributesEntry
}
var file_google_pubsub_v1_pubsub_proto_depIdxs = []int32{
	2, // 0: google.pubsub.v1.PushPubsubMessageRequest.message:type_name -> google.pubsub.v1.PubsubMessage
	3, // 1: google.pubsub.v1.PubsubMessage.attributes:type_name -> google.pubsub.v1.PubsubMessage.AttributesEntry
	0, // 2: google.pubsub.v1.PubsubService.PushPubsubMessage:input_type -> google.pubsub.v1.PushPubsubMessageRequest
	1, // 3: google.pubsub.v1.PubsubService.PushPubsubMessage:output_type -> google.pubsub.v1.PushPubsubMessageResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_pubsub_v1_pubsub_proto_init() }
func file_google_pubsub_v1_pubsub_proto_init() {
	if File_google_pubsub_v1_pubsub_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_pubsub_v1_pubsub_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushPubsubMessageRequest); i {
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
		file_google_pubsub_v1_pubsub_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushPubsubMessageResponse); i {
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
		file_google_pubsub_v1_pubsub_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PubsubMessage); i {
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
			RawDescriptor: file_google_pubsub_v1_pubsub_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_pubsub_v1_pubsub_proto_goTypes,
		DependencyIndexes: file_google_pubsub_v1_pubsub_proto_depIdxs,
		MessageInfos:      file_google_pubsub_v1_pubsub_proto_msgTypes,
	}.Build()
	File_google_pubsub_v1_pubsub_proto = out.File
	file_google_pubsub_v1_pubsub_proto_rawDesc = nil
	file_google_pubsub_v1_pubsub_proto_goTypes = nil
	file_google_pubsub_v1_pubsub_proto_depIdxs = nil
}
