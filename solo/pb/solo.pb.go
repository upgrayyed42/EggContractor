// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: solo/pb/solo.proto

package pb

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type SoloContract struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                             string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                           string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	IsElite                        bool    `protobuf:"varint,3,opt,name=is_elite,json=isElite,proto3" json:"is_elite,omitempty"`
	UltimateGoal                   float64 `protobuf:"fixed64,4,opt,name=ultimate_goal,json=ultimateGoal,proto3" json:"ultimate_goal,omitempty"`
	EggsLaid                       float64 `protobuf:"fixed64,5,opt,name=eggs_laid,json=eggsLaid,proto3" json:"eggs_laid,omitempty"`
	EggsPerSecond                  float64 `protobuf:"fixed64,6,opt,name=eggs_per_second,json=eggsPerSecond,proto3" json:"eggs_per_second,omitempty"`
	SecondsUntilProductionDeadline float64 `protobuf:"fixed64,8,opt,name=seconds_until_production_deadline,json=secondsUntilProductionDeadline,proto3" json:"seconds_until_production_deadline,omitempty"`
	SecondsUntilCollectionDeadline float64 `protobuf:"fixed64,9,opt,name=seconds_until_collection_deadline,json=secondsUntilCollectionDeadline,proto3" json:"seconds_until_collection_deadline,omitempty"`
	ServerRefreshTimestamp         float64 `protobuf:"fixed64,10,opt,name=server_refresh_timestamp,json=serverRefreshTimestamp,proto3" json:"server_refresh_timestamp,omitempty"`
}

func (x *SoloContract) Reset() {
	*x = SoloContract{}
	if protoimpl.UnsafeEnabled {
		mi := &file_solo_pb_solo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SoloContract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SoloContract) ProtoMessage() {}

func (x *SoloContract) ProtoReflect() protoreflect.Message {
	mi := &file_solo_pb_solo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SoloContract.ProtoReflect.Descriptor instead.
func (*SoloContract) Descriptor() ([]byte, []int) {
	return file_solo_pb_solo_proto_rawDescGZIP(), []int{0}
}

func (x *SoloContract) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SoloContract) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SoloContract) GetIsElite() bool {
	if x != nil {
		return x.IsElite
	}
	return false
}

func (x *SoloContract) GetUltimateGoal() float64 {
	if x != nil {
		return x.UltimateGoal
	}
	return 0
}

func (x *SoloContract) GetEggsLaid() float64 {
	if x != nil {
		return x.EggsLaid
	}
	return 0
}

func (x *SoloContract) GetEggsPerSecond() float64 {
	if x != nil {
		return x.EggsPerSecond
	}
	return 0
}

func (x *SoloContract) GetSecondsUntilProductionDeadline() float64 {
	if x != nil {
		return x.SecondsUntilProductionDeadline
	}
	return 0
}

func (x *SoloContract) GetSecondsUntilCollectionDeadline() float64 {
	if x != nil {
		return x.SecondsUntilCollectionDeadline
	}
	return 0
}

func (x *SoloContract) GetServerRefreshTimestamp() float64 {
	if x != nil {
		return x.ServerRefreshTimestamp
	}
	return 0
}

var File_solo_pb_solo_proto protoreflect.FileDescriptor

var file_solo_pb_solo_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x6f, 0x6c, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x87, 0x03, 0x0a, 0x0c, 0x53, 0x6f, 0x6c, 0x6f, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f,
	0x65, 0x6c, 0x69, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x45,
	0x6c, 0x69, 0x74, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x75, 0x6c, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65,
	0x5f, 0x67, 0x6f, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x75, 0x6c, 0x74,
	0x69, 0x6d, 0x61, 0x74, 0x65, 0x47, 0x6f, 0x61, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x67, 0x67,
	0x73, 0x5f, 0x6c, 0x61, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x65, 0x67,
	0x67, 0x73, 0x4c, 0x61, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x65, 0x67, 0x67, 0x73, 0x5f, 0x70,
	0x65, 0x72, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0d, 0x65, 0x67, 0x67, 0x73, 0x50, 0x65, 0x72, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x12, 0x49,
	0x0a, 0x21, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x5f, 0x75, 0x6e, 0x74, 0x69, 0x6c, 0x5f,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x65, 0x61, 0x64, 0x6c,
	0x69, 0x6e, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x1e, 0x73, 0x65, 0x63, 0x6f, 0x6e,
	0x64, 0x73, 0x55, 0x6e, 0x74, 0x69, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x44, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x49, 0x0a, 0x21, 0x73, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x73, 0x5f, 0x75, 0x6e, 0x74, 0x69, 0x6c, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x1e, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x55, 0x6e, 0x74,
	0x69, 0x6c, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x61, 0x64,
	0x6c, 0x69, 0x6e, 0x65, 0x12, 0x38, 0x0a, 0x18, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x72,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x16, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65,
	0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x32,
	0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x61, 0x6e,
	0x61, 0x74, 0x69, 0x63, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x65, 0x72, 0x2f, 0x45, 0x67, 0x67,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_solo_pb_solo_proto_rawDescOnce sync.Once
	file_solo_pb_solo_proto_rawDescData = file_solo_pb_solo_proto_rawDesc
)

func file_solo_pb_solo_proto_rawDescGZIP() []byte {
	file_solo_pb_solo_proto_rawDescOnce.Do(func() {
		file_solo_pb_solo_proto_rawDescData = protoimpl.X.CompressGZIP(file_solo_pb_solo_proto_rawDescData)
	})
	return file_solo_pb_solo_proto_rawDescData
}

var (
	file_solo_pb_solo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
	file_solo_pb_solo_proto_goTypes  = []interface{}{
		(*SoloContract)(nil), // 0: SoloContract
	}
)

var file_solo_pb_solo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_solo_pb_solo_proto_init() }
func file_solo_pb_solo_proto_init() {
	if File_solo_pb_solo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_solo_pb_solo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SoloContract); i {
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
			RawDescriptor: file_solo_pb_solo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_solo_pb_solo_proto_goTypes,
		DependencyIndexes: file_solo_pb_solo_proto_depIdxs,
		MessageInfos:      file_solo_pb_solo_proto_msgTypes,
	}.Build()
	File_solo_pb_solo_proto = out.File
	file_solo_pb_solo_proto_rawDesc = nil
	file_solo_pb_solo_proto_goTypes = nil
	file_solo_pb_solo_proto_depIdxs = nil
}
