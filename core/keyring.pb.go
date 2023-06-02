// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.22.3
// source: dandelion/keyring.proto

package core

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

type KeyRing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*KeyRing_Secret `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *KeyRing) Reset() {
	*x = KeyRing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dandelion_keyring_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyRing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyRing) ProtoMessage() {}

func (x *KeyRing) ProtoReflect() protoreflect.Message {
	mi := &file_dandelion_keyring_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyRing.ProtoReflect.Descriptor instead.
func (*KeyRing) Descriptor() ([]byte, []int) {
	return file_dandelion_keyring_proto_rawDescGZIP(), []int{0}
}

func (x *KeyRing) GetList() []*KeyRing_Secret {
	if x != nil {
		return x.List
	}
	return nil
}

type KeyRing_Secret struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Raw []byte `protobuf:"bytes,2,opt,name=raw,proto3" json:"raw,omitempty"`
}

func (x *KeyRing_Secret) Reset() {
	*x = KeyRing_Secret{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dandelion_keyring_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyRing_Secret) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyRing_Secret) ProtoMessage() {}

func (x *KeyRing_Secret) ProtoReflect() protoreflect.Message {
	mi := &file_dandelion_keyring_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyRing_Secret.ProtoReflect.Descriptor instead.
func (*KeyRing_Secret) Descriptor() ([]byte, []int) {
	return file_dandelion_keyring_proto_rawDescGZIP(), []int{0, 0}
}

func (x *KeyRing_Secret) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *KeyRing_Secret) GetRaw() []byte {
	if x != nil {
		return x.Raw
	}
	return nil
}

var File_dandelion_keyring_proto protoreflect.FileDescriptor

var file_dandelion_keyring_proto_rawDesc = []byte{
	0x0a, 0x17, 0x64, 0x61, 0x6e, 0x64, 0x65, 0x6c, 0x69, 0x6f, 0x6e, 0x2f, 0x6b, 0x65, 0x79, 0x72,
	0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x64, 0x61, 0x6e, 0x64, 0x65,
	0x6c, 0x69, 0x6f, 0x6e, 0x22, 0x64, 0x0a, 0x07, 0x4b, 0x65, 0x79, 0x52, 0x69, 0x6e, 0x67, 0x12,
	0x2d, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x64, 0x61, 0x6e, 0x64, 0x65, 0x6c, 0x69, 0x6f, 0x6e, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x69, 0x6e,
	0x67, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x1a, 0x2a,
	0x0a, 0x06, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x61, 0x77, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x72, 0x61, 0x77, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x6f, 0x73,
	0x2d, 0x74, 0x61, 0x63, 0x68, 0x79, 0x6f, 0x6e, 0x2f, 0x64, 0x61, 0x6e, 0x64, 0x65, 0x6c, 0x69,
	0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dandelion_keyring_proto_rawDescOnce sync.Once
	file_dandelion_keyring_proto_rawDescData = file_dandelion_keyring_proto_rawDesc
)

func file_dandelion_keyring_proto_rawDescGZIP() []byte {
	file_dandelion_keyring_proto_rawDescOnce.Do(func() {
		file_dandelion_keyring_proto_rawDescData = protoimpl.X.CompressGZIP(file_dandelion_keyring_proto_rawDescData)
	})
	return file_dandelion_keyring_proto_rawDescData
}

var file_dandelion_keyring_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_dandelion_keyring_proto_goTypes = []interface{}{
	(*KeyRing)(nil),        // 0: dandelion.KeyRing
	(*KeyRing_Secret)(nil), // 1: dandelion.KeyRing.Secret
}
var file_dandelion_keyring_proto_depIdxs = []int32{
	1, // 0: dandelion.KeyRing.list:type_name -> dandelion.KeyRing.Secret
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_dandelion_keyring_proto_init() }
func file_dandelion_keyring_proto_init() {
	if File_dandelion_keyring_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dandelion_keyring_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyRing); i {
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
		file_dandelion_keyring_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyRing_Secret); i {
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
			RawDescriptor: file_dandelion_keyring_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_dandelion_keyring_proto_goTypes,
		DependencyIndexes: file_dandelion_keyring_proto_depIdxs,
		MessageInfos:      file_dandelion_keyring_proto_msgTypes,
	}.Build()
	File_dandelion_keyring_proto = out.File
	file_dandelion_keyring_proto_rawDesc = nil
	file_dandelion_keyring_proto_goTypes = nil
	file_dandelion_keyring_proto_depIdxs = nil
}