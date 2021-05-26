// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.1
// source: phoneservice.proto

// user 包

package phoneservice

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

// client 服务的各个接口的请求/响应结构
type PhoneRegRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone     string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`         //手机号码
	Phonecode string `protobuf:"bytes,2,opt,name=phonecode,proto3" json:"phonecode,omitempty"` //验证码
}

func (x *PhoneRegRequest) Reset() {
	*x = PhoneRegRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_phoneservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PhoneRegRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PhoneRegRequest) ProtoMessage() {}

func (x *PhoneRegRequest) ProtoReflect() protoreflect.Message {
	mi := &file_phoneservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PhoneRegRequest.ProtoReflect.Descriptor instead.
func (*PhoneRegRequest) Descriptor() ([]byte, []int) {
	return file_phoneservice_proto_rawDescGZIP(), []int{0}
}

func (x *PhoneRegRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *PhoneRegRequest) GetPhonecode() string {
	if x != nil {
		return x.Phonecode
	}
	return ""
}

type PhoneRegResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Err int32  `protobuf:"varint,1,opt,name=err,proto3" json:"err,omitempty"`
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *PhoneRegResponse) Reset() {
	*x = PhoneRegResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_phoneservice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PhoneRegResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PhoneRegResponse) ProtoMessage() {}

func (x *PhoneRegResponse) ProtoReflect() protoreflect.Message {
	mi := &file_phoneservice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PhoneRegResponse.ProtoReflect.Descriptor instead.
func (*PhoneRegResponse) Descriptor() ([]byte, []int) {
	return file_phoneservice_proto_rawDescGZIP(), []int{1}
}

func (x *PhoneRegResponse) GetErr() int32 {
	if x != nil {
		return x.Err
	}
	return 0
}

func (x *PhoneRegResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_phoneservice_proto protoreflect.FileDescriptor

var file_phoneservice_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x22, 0x45, 0x0a, 0x0f, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x36, 0x0a, 0x10, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x52, 0x65, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x65, 0x72, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x65, 0x72, 0x72, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73,
	0x67, 0x32, 0x5b, 0x0a, 0x0c, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x4b, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1d, 0x2e,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x52, 0x65, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x68, 0x6f, 0x6e,
	0x65, 0x52, 0x65, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0f,
	0x5a, 0x0d, 0x2f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_phoneservice_proto_rawDescOnce sync.Once
	file_phoneservice_proto_rawDescData = file_phoneservice_proto_rawDesc
)

func file_phoneservice_proto_rawDescGZIP() []byte {
	file_phoneservice_proto_rawDescOnce.Do(func() {
		file_phoneservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_phoneservice_proto_rawDescData)
	})
	return file_phoneservice_proto_rawDescData
}

var file_phoneservice_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_phoneservice_proto_goTypes = []interface{}{
	(*PhoneRegRequest)(nil),  // 0: phoneservice.PhoneRegRequest
	(*PhoneRegResponse)(nil), // 1: phoneservice.PhoneRegResponse
}
var file_phoneservice_proto_depIdxs = []int32{
	0, // 0: phoneservice.PhoneService.RegPhone:input_type -> phoneservice.PhoneRegRequest
	1, // 1: phoneservice.PhoneService.RegPhone:output_type -> phoneservice.PhoneRegResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_phoneservice_proto_init() }
func file_phoneservice_proto_init() {
	if File_phoneservice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_phoneservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PhoneRegRequest); i {
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
		file_phoneservice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PhoneRegResponse); i {
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
			RawDescriptor: file_phoneservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_phoneservice_proto_goTypes,
		DependencyIndexes: file_phoneservice_proto_depIdxs,
		MessageInfos:      file_phoneservice_proto_msgTypes,
	}.Build()
	File_phoneservice_proto = out.File
	file_phoneservice_proto_rawDesc = nil
	file_phoneservice_proto_goTypes = nil
	file_phoneservice_proto_depIdxs = nil
}
