// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: greet/greet.proto

// The namespace when other proto file imports.

package greet

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

type Greet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"` // must be number > 0;
	Message string `protobuf:"bytes,10,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Greet) Reset() {
	*x = Greet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greet_greet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Greet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Greet) ProtoMessage() {}

func (x *Greet) ProtoReflect() protoreflect.Message {
	mi := &file_greet_greet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Greet.ProtoReflect.Descriptor instead.
func (*Greet) Descriptor() ([]byte, []int) {
	return file_greet_greet_proto_rawDescGZIP(), []int{0}
}

func (x *Greet) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Greet) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_greet_greet_proto protoreflect.FileDescriptor

var file_greet_greet_proto_rawDesc = []byte{
	0x0a, 0x11, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2f, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x67, 0x72, 0x65, 0x65, 0x74, 0x22, 0x35, 0x0a, 0x05, 0x47, 0x72,
	0x65, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x42, 0x0d, 0x5a, 0x0b, 0x67, 0x72, 0x70, 0x63, 0x62, 0x2f, 0x67, 0x72, 0x65, 0x65, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_greet_greet_proto_rawDescOnce sync.Once
	file_greet_greet_proto_rawDescData = file_greet_greet_proto_rawDesc
)

func file_greet_greet_proto_rawDescGZIP() []byte {
	file_greet_greet_proto_rawDescOnce.Do(func() {
		file_greet_greet_proto_rawDescData = protoimpl.X.CompressGZIP(file_greet_greet_proto_rawDescData)
	})
	return file_greet_greet_proto_rawDescData
}

var file_greet_greet_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_greet_greet_proto_goTypes = []interface{}{
	(*Greet)(nil), // 0: greet.Greet
}
var file_greet_greet_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_greet_greet_proto_init() }
func file_greet_greet_proto_init() {
	if File_greet_greet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_greet_greet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Greet); i {
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
			RawDescriptor: file_greet_greet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_greet_greet_proto_goTypes,
		DependencyIndexes: file_greet_greet_proto_depIdxs,
		MessageInfos:      file_greet_greet_proto_msgTypes,
	}.Build()
	File_greet_greet_proto = out.File
	file_greet_greet_proto_rawDesc = nil
	file_greet_greet_proto_goTypes = nil
	file_greet_greet_proto_depIdxs = nil
}
