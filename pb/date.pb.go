// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: date.proto

package pb

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

type Date struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Year  int32 `protobuf:"varint,1,opt,name=year,proto3" json:"year,omitempty"`
	Month int32 `protobuf:"varint,2,opt,name=month,proto3" json:"month,omitempty"`
	Day   int32 `protobuf:"varint,3,opt,name=day,proto3" json:"day,omitempty"`
}

func (x *Date) Reset() {
	*x = Date{}
	if protoimpl.UnsafeEnabled {
		mi := &file_date_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Date) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Date) ProtoMessage() {}

func (x *Date) ProtoReflect() protoreflect.Message {
	mi := &file_date_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Date.ProtoReflect.Descriptor instead.
func (*Date) Descriptor() ([]byte, []int) {
	return file_date_proto_rawDescGZIP(), []int{0}
}

func (x *Date) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *Date) GetMonth() int32 {
	if x != nil {
		return x.Month
	}
	return 0
}

func (x *Date) GetDay() int32 {
	if x != nil {
		return x.Day
	}
	return 0
}

var File_date_proto protoreflect.FileDescriptor

var file_date_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x42, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x79,
	0x65, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x61, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x03, 0x64, 0x61, 0x79, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_date_proto_rawDescOnce sync.Once
	file_date_proto_rawDescData = file_date_proto_rawDesc
)

func file_date_proto_rawDescGZIP() []byte {
	file_date_proto_rawDescOnce.Do(func() {
		file_date_proto_rawDescData = protoimpl.X.CompressGZIP(file_date_proto_rawDescData)
	})
	return file_date_proto_rawDescData
}

var file_date_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_date_proto_goTypes = []interface{}{
	(*Date)(nil), // 0: proto.Date
}
var file_date_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_date_proto_init() }
func file_date_proto_init() {
	if File_date_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_date_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Date); i {
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
			RawDescriptor: file_date_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_date_proto_goTypes,
		DependencyIndexes: file_date_proto_depIdxs,
		MessageInfos:      file_date_proto_msgTypes,
	}.Build()
	File_date_proto = out.File
	file_date_proto_rawDesc = nil
	file_date_proto_goTypes = nil
	file_date_proto_depIdxs = nil
}
