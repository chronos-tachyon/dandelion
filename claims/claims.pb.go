// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.22.3
// source: dandelion/claims/claims.proto

package claims

import (
	core "github.com/chronos-tachyon/dandelion/core"
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

type Retired_Reason int32

const (
	Retired_MANUAL_ROTATION     Retired_Reason = 0
	Retired_SCHEDULED_ROTATION  Retired_Reason = 1
	Retired_POSSIBLE_COMPROMISE Retired_Reason = 2
	Retired_KNOWN_COMPROMISE    Retired_Reason = 3
)

// Enum value maps for Retired_Reason.
var (
	Retired_Reason_name = map[int32]string{
		0: "MANUAL_ROTATION",
		1: "SCHEDULED_ROTATION",
		2: "POSSIBLE_COMPROMISE",
		3: "KNOWN_COMPROMISE",
	}
	Retired_Reason_value = map[string]int32{
		"MANUAL_ROTATION":     0,
		"SCHEDULED_ROTATION":  1,
		"POSSIBLE_COMPROMISE": 2,
		"KNOWN_COMPROMISE":    3,
	}
)

func (x Retired_Reason) Enum() *Retired_Reason {
	p := new(Retired_Reason)
	*p = x
	return p
}

func (x Retired_Reason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Retired_Reason) Descriptor() protoreflect.EnumDescriptor {
	return file_dandelion_claims_claims_proto_enumTypes[0].Descriptor()
}

func (Retired_Reason) Type() protoreflect.EnumType {
	return &file_dandelion_claims_claims_proto_enumTypes[0]
}

func (x Retired_Reason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Retired_Reason.Descriptor instead.
func (Retired_Reason) EnumDescriptor() ([]byte, []int) {
	return file_dandelion_claims_claims_proto_rawDescGZIP(), []int{4, 0}
}

type Confirmed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Inner *core.SignedAttestation `protobuf:"bytes,1,opt,name=inner,proto3" json:"inner,omitempty"`
}

func (x *Confirmed) Reset() {
	*x = Confirmed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dandelion_claims_claims_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Confirmed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Confirmed) ProtoMessage() {}

func (x *Confirmed) ProtoReflect() protoreflect.Message {
	mi := &file_dandelion_claims_claims_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Confirmed.ProtoReflect.Descriptor instead.
func (*Confirmed) Descriptor() ([]byte, []int) {
	return file_dandelion_claims_claims_proto_rawDescGZIP(), []int{0}
}

func (x *Confirmed) GetInner() *core.SignedAttestation {
	if x != nil {
		return x.Inner
	}
	return nil
}

type Disputed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Inner *core.SignedAttestation `protobuf:"bytes,1,opt,name=inner,proto3" json:"inner,omitempty"`
}

func (x *Disputed) Reset() {
	*x = Disputed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dandelion_claims_claims_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Disputed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Disputed) ProtoMessage() {}

func (x *Disputed) ProtoReflect() protoreflect.Message {
	mi := &file_dandelion_claims_claims_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Disputed.ProtoReflect.Descriptor instead.
func (*Disputed) Descriptor() ([]byte, []int) {
	return file_dandelion_claims_claims_proto_rawDescGZIP(), []int{1}
}

func (x *Disputed) GetInner() *core.SignedAttestation {
	if x != nil {
		return x.Inner
	}
	return nil
}

type MemberOf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupEntity *core.Identity `protobuf:"bytes,1,opt,name=group_entity,json=groupEntity,proto3" json:"group_entity,omitempty"`
}

func (x *MemberOf) Reset() {
	*x = MemberOf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dandelion_claims_claims_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberOf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberOf) ProtoMessage() {}

func (x *MemberOf) ProtoReflect() protoreflect.Message {
	mi := &file_dandelion_claims_claims_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberOf.ProtoReflect.Descriptor instead.
func (*MemberOf) Descriptor() ([]byte, []int) {
	return file_dandelion_claims_claims_proto_rawDescGZIP(), []int{2}
}

func (x *MemberOf) GetGroupEntity() *core.Identity {
	if x != nil {
		return x.GroupEntity
	}
	return nil
}

type FormerlyKnownAs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FormerIdentity *core.Identity `protobuf:"bytes,1,opt,name=former_identity,json=formerIdentity,proto3" json:"former_identity,omitempty"`
}

func (x *FormerlyKnownAs) Reset() {
	*x = FormerlyKnownAs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dandelion_claims_claims_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FormerlyKnownAs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FormerlyKnownAs) ProtoMessage() {}

func (x *FormerlyKnownAs) ProtoReflect() protoreflect.Message {
	mi := &file_dandelion_claims_claims_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FormerlyKnownAs.ProtoReflect.Descriptor instead.
func (*FormerlyKnownAs) Descriptor() ([]byte, []int) {
	return file_dandelion_claims_claims_proto_rawDescGZIP(), []int{3}
}

func (x *FormerlyKnownAs) GetFormerIdentity() *core.Identity {
	if x != nil {
		return x.FormerIdentity
	}
	return nil
}

type Retired struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reason Retired_Reason `protobuf:"varint,1,opt,name=reason,proto3,enum=dandelion.claims.Retired_Reason" json:"reason,omitempty"`
}

func (x *Retired) Reset() {
	*x = Retired{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dandelion_claims_claims_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Retired) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Retired) ProtoMessage() {}

func (x *Retired) ProtoReflect() protoreflect.Message {
	mi := &file_dandelion_claims_claims_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Retired.ProtoReflect.Descriptor instead.
func (*Retired) Descriptor() ([]byte, []int) {
	return file_dandelion_claims_claims_proto_rawDescGZIP(), []int{4}
}

func (x *Retired) GetReason() Retired_Reason {
	if x != nil {
		return x.Reason
	}
	return Retired_MANUAL_ROTATION
}

var File_dandelion_claims_claims_proto protoreflect.FileDescriptor

var file_dandelion_claims_claims_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x64, 0x61, 0x6e, 0x64, 0x65, 0x6c, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x6c, 0x61, 0x69,
	0x6d, 0x73, 0x2f, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x64, 0x61, 0x6e, 0x64, 0x65, 0x6c, 0x69, 0x6f, 0x6e, 0x2e, 0x63, 0x6c, 0x61, 0x69, 0x6d,
	0x73, 0x1a, 0x14, 0x64, 0x61, 0x6e, 0x64, 0x65, 0x6c, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3f, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x72, 0x6d, 0x65, 0x64, 0x12, 0x32, 0x0a, 0x05, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x64, 0x61, 0x6e, 0x64, 0x65, 0x6c, 0x69, 0x6f, 0x6e, 0x2e,
	0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x05, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x22, 0x3e, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x70,
	0x75, 0x74, 0x65, 0x64, 0x12, 0x32, 0x0a, 0x05, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x64, 0x61, 0x6e, 0x64, 0x65, 0x6c, 0x69, 0x6f, 0x6e, 0x2e,
	0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x05, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x22, 0x42, 0x0a, 0x08, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x4f, 0x66, 0x12, 0x36, 0x0a, 0x0c, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x64, 0x61, 0x6e,
	0x64, 0x65, 0x6c, 0x69, 0x6f, 0x6e, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52,
	0x0b, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x4f, 0x0a, 0x0f,
	0x46, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x6c, 0x79, 0x4b, 0x6e, 0x6f, 0x77, 0x6e, 0x41, 0x73, 0x12,
	0x3c, 0x0a, 0x0f, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x64, 0x61, 0x6e, 0x64, 0x65,
	0x6c, 0x69, 0x6f, 0x6e, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x0e, 0x66,
	0x6f, 0x72, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0xa9, 0x01,
	0x0a, 0x07, 0x52, 0x65, 0x74, 0x69, 0x72, 0x65, 0x64, 0x12, 0x38, 0x0a, 0x06, 0x72, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x64, 0x61, 0x6e, 0x64,
	0x65, 0x6c, 0x69, 0x6f, 0x6e, 0x2e, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x2e, 0x52, 0x65, 0x74,
	0x69, 0x72, 0x65, 0x64, 0x2e, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x72, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x22, 0x64, 0x0a, 0x06, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x13, 0x0a,
	0x0f, 0x4d, 0x41, 0x4e, 0x55, 0x41, 0x4c, 0x5f, 0x52, 0x4f, 0x54, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x43, 0x48, 0x45, 0x44, 0x55, 0x4c, 0x45, 0x44, 0x5f,
	0x52, 0x4f, 0x54, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x50, 0x4f,
	0x53, 0x53, 0x49, 0x42, 0x4c, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x52, 0x4f, 0x4d, 0x49, 0x53,
	0x45, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x43, 0x4f, 0x4d,
	0x50, 0x52, 0x4f, 0x4d, 0x49, 0x53, 0x45, 0x10, 0x03, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x6f, 0x73, 0x2d,
	0x74, 0x61, 0x63, 0x68, 0x79, 0x6f, 0x6e, 0x2f, 0x64, 0x61, 0x6e, 0x64, 0x65, 0x6c, 0x69, 0x6f,
	0x6e, 0x2f, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dandelion_claims_claims_proto_rawDescOnce sync.Once
	file_dandelion_claims_claims_proto_rawDescData = file_dandelion_claims_claims_proto_rawDesc
)

func file_dandelion_claims_claims_proto_rawDescGZIP() []byte {
	file_dandelion_claims_claims_proto_rawDescOnce.Do(func() {
		file_dandelion_claims_claims_proto_rawDescData = protoimpl.X.CompressGZIP(file_dandelion_claims_claims_proto_rawDescData)
	})
	return file_dandelion_claims_claims_proto_rawDescData
}

var file_dandelion_claims_claims_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_dandelion_claims_claims_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_dandelion_claims_claims_proto_goTypes = []interface{}{
	(Retired_Reason)(0),            // 0: dandelion.claims.Retired.Reason
	(*Confirmed)(nil),              // 1: dandelion.claims.Confirmed
	(*Disputed)(nil),               // 2: dandelion.claims.Disputed
	(*MemberOf)(nil),               // 3: dandelion.claims.MemberOf
	(*FormerlyKnownAs)(nil),        // 4: dandelion.claims.FormerlyKnownAs
	(*Retired)(nil),                // 5: dandelion.claims.Retired
	(*core.SignedAttestation)(nil), // 6: dandelion.SignedAttestation
	(*core.Identity)(nil),          // 7: dandelion.Identity
}
var file_dandelion_claims_claims_proto_depIdxs = []int32{
	6, // 0: dandelion.claims.Confirmed.inner:type_name -> dandelion.SignedAttestation
	6, // 1: dandelion.claims.Disputed.inner:type_name -> dandelion.SignedAttestation
	7, // 2: dandelion.claims.MemberOf.group_entity:type_name -> dandelion.Identity
	7, // 3: dandelion.claims.FormerlyKnownAs.former_identity:type_name -> dandelion.Identity
	0, // 4: dandelion.claims.Retired.reason:type_name -> dandelion.claims.Retired.Reason
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_dandelion_claims_claims_proto_init() }
func file_dandelion_claims_claims_proto_init() {
	if File_dandelion_claims_claims_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dandelion_claims_claims_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Confirmed); i {
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
		file_dandelion_claims_claims_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Disputed); i {
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
		file_dandelion_claims_claims_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberOf); i {
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
		file_dandelion_claims_claims_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FormerlyKnownAs); i {
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
		file_dandelion_claims_claims_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Retired); i {
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
			RawDescriptor: file_dandelion_claims_claims_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_dandelion_claims_claims_proto_goTypes,
		DependencyIndexes: file_dandelion_claims_claims_proto_depIdxs,
		EnumInfos:         file_dandelion_claims_claims_proto_enumTypes,
		MessageInfos:      file_dandelion_claims_claims_proto_msgTypes,
	}.Build()
	File_dandelion_claims_claims_proto = out.File
	file_dandelion_claims_claims_proto_rawDesc = nil
	file_dandelion_claims_claims_proto_goTypes = nil
	file_dandelion_claims_claims_proto_depIdxs = nil
}
