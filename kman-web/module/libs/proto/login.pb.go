// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: login.proto

package proto

import (
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

type LoginReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Namespace     string                 `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	ProjectId     int32                  `protobuf:"varint,4,opt,name=projectId,proto3" json:"projectId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginReq) Reset() {
	*x = LoginReq{}
	mi := &file_login_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReq) ProtoMessage() {}

func (x *LoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[0]
	if x != nil {
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
	return file_login_proto_rawDescGZIP(), []int{0}
}

func (x *LoginReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginReq) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *LoginReq) GetProjectId() int32 {
	if x != nil {
		return x.ProjectId
	}
	return 0
}

type LoginResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int64                  `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Permissions   []int64                `protobuf:"varint,2,rep,packed,name=permissions,proto3" json:"permissions,omitempty"`
	ProjectId     int32                  `protobuf:"varint,3,opt,name=projectId,proto3" json:"projectId,omitempty"`
	Namespace     string                 `protobuf:"bytes,4,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Password      string                 `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginResp) Reset() {
	*x = LoginResp{}
	mi := &file_login_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResp) ProtoMessage() {}

func (x *LoginResp) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[1]
	if x != nil {
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
	return file_login_proto_rawDescGZIP(), []int{1}
}

func (x *LoginResp) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *LoginResp) GetPermissions() []int64 {
	if x != nil {
		return x.Permissions
	}
	return nil
}

func (x *LoginResp) GetProjectId() int32 {
	if x != nil {
		return x.ProjectId
	}
	return 0
}

func (x *LoginResp) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *LoginResp) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

var File_login_proto protoreflect.FileDescriptor

const file_login_proto_rawDesc = "" +
	"\n" +
	"\vlogin.proto\"b\n" +
	"\bLoginReq\x12\x1a\n" +
	"\busername\x18\x01 \x01(\tR\busername\x12\x1c\n" +
	"\tnamespace\x18\x03 \x01(\tR\tnamespace\x12\x1c\n" +
	"\tprojectId\x18\x04 \x01(\x05R\tprojectId\"\x9d\x01\n" +
	"\tLoginResp\x12\x16\n" +
	"\x06userId\x18\x01 \x01(\x03R\x06userId\x12 \n" +
	"\vpermissions\x18\x02 \x03(\x03R\vpermissions\x12\x1c\n" +
	"\tprojectId\x18\x03 \x01(\x05R\tprojectId\x12\x1c\n" +
	"\tnamespace\x18\x04 \x01(\tR\tnamespace\x12\x1a\n" +
	"\bpassword\x18\x05 \x01(\tR\bpassword2)\n" +
	"\x05Login\x12 \n" +
	"\x05Login\x12\t.LoginReq\x1a\n" +
	".LoginResp\"\x00B\tZ\a./protob\x06proto3"

var (
	file_login_proto_rawDescOnce sync.Once
	file_login_proto_rawDescData []byte
)

func file_login_proto_rawDescGZIP() []byte {
	file_login_proto_rawDescOnce.Do(func() {
		file_login_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_login_proto_rawDesc), len(file_login_proto_rawDesc)))
	})
	return file_login_proto_rawDescData
}

var file_login_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_login_proto_goTypes = []any{
	(*LoginReq)(nil),  // 0: LoginReq
	(*LoginResp)(nil), // 1: LoginResp
}
var file_login_proto_depIdxs = []int32{
	0, // 0: Login.Login:input_type -> LoginReq
	1, // 1: Login.Login:output_type -> LoginResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_login_proto_init() }
func file_login_proto_init() {
	if File_login_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_login_proto_rawDesc), len(file_login_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_login_proto_goTypes,
		DependencyIndexes: file_login_proto_depIdxs,
		MessageInfos:      file_login_proto_msgTypes,
	}.Build()
	File_login_proto = out.File
	file_login_proto_goTypes = nil
	file_login_proto_depIdxs = nil
}
