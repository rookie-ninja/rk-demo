// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.1
// source: v1/server.proto

package greeter

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

type LoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LoginReq) Reset() {
	*x = LoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReq) ProtoMessage() {}

func (x *LoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_v1_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReq.ProtoReflect.Descriptor instead.
func (*LoginReq) Descriptor() ([]byte, []int) {
	return file_v1_server_proto_rawDescGZIP(), []int{0}
}

type LoginResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LoginResp) Reset() {
	*x = LoginResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResp) ProtoMessage() {}

func (x *LoginResp) ProtoReflect() protoreflect.Message {
	mi := &file_v1_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResp.ProtoReflect.Descriptor instead.
func (*LoginResp) Descriptor() ([]byte, []int) {
	return file_v1_server_proto_rawDescGZIP(), []int{1}
}

type CallAReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CallAReq) Reset() {
	*x = CallAReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallAReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallAReq) ProtoMessage() {}

func (x *CallAReq) ProtoReflect() protoreflect.Message {
	mi := &file_v1_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallAReq.ProtoReflect.Descriptor instead.
func (*CallAReq) Descriptor() ([]byte, []int) {
	return file_v1_server_proto_rawDescGZIP(), []int{2}
}

type CallAResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CallAResp) Reset() {
	*x = CallAResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallAResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallAResp) ProtoMessage() {}

func (x *CallAResp) ProtoReflect() protoreflect.Message {
	mi := &file_v1_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallAResp.ProtoReflect.Descriptor instead.
func (*CallAResp) Descriptor() ([]byte, []int) {
	return file_v1_server_proto_rawDescGZIP(), []int{3}
}

type CallBReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CallBReq) Reset() {
	*x = CallBReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_server_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallBReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallBReq) ProtoMessage() {}

func (x *CallBReq) ProtoReflect() protoreflect.Message {
	mi := &file_v1_server_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallBReq.ProtoReflect.Descriptor instead.
func (*CallBReq) Descriptor() ([]byte, []int) {
	return file_v1_server_proto_rawDescGZIP(), []int{4}
}

type CallBResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CallBResp) Reset() {
	*x = CallBResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_server_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallBResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallBResp) ProtoMessage() {}

func (x *CallBResp) ProtoReflect() protoreflect.Message {
	mi := &file_v1_server_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallBResp.ProtoReflect.Descriptor instead.
func (*CallBResp) Descriptor() ([]byte, []int) {
	return file_v1_server_proto_rawDescGZIP(), []int{5}
}

var File_v1_server_proto protoreflect.FileDescriptor

var file_v1_server_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x22, 0x0a, 0x0a, 0x08, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x22, 0x0b, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x22, 0x0a, 0x0a, 0x08, 0x43, 0x61, 0x6c, 0x6c, 0x41, 0x52, 0x65, 0x71, 0x22, 0x0b,
	0x0a, 0x09, 0x43, 0x61, 0x6c, 0x6c, 0x41, 0x52, 0x65, 0x73, 0x70, 0x22, 0x0a, 0x0a, 0x08, 0x43,
	0x61, 0x6c, 0x6c, 0x42, 0x52, 0x65, 0x71, 0x22, 0x0b, 0x0a, 0x09, 0x43, 0x61, 0x6c, 0x6c, 0x42,
	0x52, 0x65, 0x73, 0x70, 0x32, 0x69, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x12,
	0x2e, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12,
	0x2e, 0x0a, 0x05, 0x43, 0x61, 0x6c, 0x6c, 0x41, 0x12, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x41, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x41, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x32,
	0x39, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x42, 0x12, 0x2e, 0x0a, 0x05, 0x43, 0x61,
	0x6c, 0x6c, 0x42, 0x12, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x6c,
	0x6c, 0x42, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x61, 0x6c, 0x6c, 0x42, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x10, 0x5a, 0x0e, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_server_proto_rawDescOnce sync.Once
	file_v1_server_proto_rawDescData = file_v1_server_proto_rawDesc
)

func file_v1_server_proto_rawDescGZIP() []byte {
	file_v1_server_proto_rawDescOnce.Do(func() {
		file_v1_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_server_proto_rawDescData)
	})
	return file_v1_server_proto_rawDescData
}

var file_v1_server_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_v1_server_proto_goTypes = []interface{}{
	(*LoginReq)(nil),  // 0: api.v1.LoginReq
	(*LoginResp)(nil), // 1: api.v1.LoginResp
	(*CallAReq)(nil),  // 2: api.v1.CallAReq
	(*CallAResp)(nil), // 3: api.v1.CallAResp
	(*CallBReq)(nil),  // 4: api.v1.CallBReq
	(*CallBResp)(nil), // 5: api.v1.CallBResp
}
var file_v1_server_proto_depIdxs = []int32{
	0, // 0: api.v1.ServerA.Login:input_type -> api.v1.LoginReq
	2, // 1: api.v1.ServerA.CallA:input_type -> api.v1.CallAReq
	4, // 2: api.v1.ServerB.CallB:input_type -> api.v1.CallBReq
	1, // 3: api.v1.ServerA.Login:output_type -> api.v1.LoginResp
	3, // 4: api.v1.ServerA.CallA:output_type -> api.v1.CallAResp
	5, // 5: api.v1.ServerB.CallB:output_type -> api.v1.CallBResp
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1_server_proto_init() }
func file_v1_server_proto_init() {
	if File_v1_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReq); i {
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
		file_v1_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResp); i {
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
		file_v1_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallAReq); i {
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
		file_v1_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallAResp); i {
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
		file_v1_server_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallBReq); i {
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
		file_v1_server_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallBResp); i {
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
			RawDescriptor: file_v1_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_v1_server_proto_goTypes,
		DependencyIndexes: file_v1_server_proto_depIdxs,
		MessageInfos:      file_v1_server_proto_msgTypes,
	}.Build()
	File_v1_server_proto = out.File
	file_v1_server_proto_rawDesc = nil
	file_v1_server_proto_goTypes = nil
	file_v1_server_proto_depIdxs = nil
}
