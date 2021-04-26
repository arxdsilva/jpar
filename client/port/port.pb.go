// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.15.8
// source: port/port.proto

package port

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

type Port struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	City        string    `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	Country     string    `protobuf:"bytes,3,opt,name=country,proto3" json:"country,omitempty"`
	Alias       []string  `protobuf:"bytes,4,rep,name=alias,proto3" json:"alias,omitempty"`
	Regions     []string  `protobuf:"bytes,5,rep,name=regions,proto3" json:"regions,omitempty"`
	Coordinates []float32 `protobuf:"fixed32,6,rep,packed,name=coordinates,proto3" json:"coordinates,omitempty"`
	Province    string    `protobuf:"bytes,7,opt,name=province,proto3" json:"province,omitempty"`
	Timezone    string    `protobuf:"bytes,8,opt,name=timezone,proto3" json:"timezone,omitempty"`
	Unlocs      []string  `protobuf:"bytes,9,rep,name=unlocs,proto3" json:"unlocs,omitempty"`
	Code        string    `protobuf:"bytes,10,opt,name=code,proto3" json:"code,omitempty"`
	Id          string    `protobuf:"bytes,11,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Port) Reset() {
	*x = Port{}
	if protoimpl.UnsafeEnabled {
		mi := &file_port_port_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Port) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Port) ProtoMessage() {}

func (x *Port) ProtoReflect() protoreflect.Message {
	mi := &file_port_port_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Port.ProtoReflect.Descriptor instead.
func (*Port) Descriptor() ([]byte, []int) {
	return file_port_port_proto_rawDescGZIP(), []int{0}
}

func (x *Port) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Port) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Port) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Port) GetAlias() []string {
	if x != nil {
		return x.Alias
	}
	return nil
}

func (x *Port) GetRegions() []string {
	if x != nil {
		return x.Regions
	}
	return nil
}

func (x *Port) GetCoordinates() []float32 {
	if x != nil {
		return x.Coordinates
	}
	return nil
}

func (x *Port) GetProvince() string {
	if x != nil {
		return x.Province
	}
	return ""
}

func (x *Port) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

func (x *Port) GetUnlocs() []string {
	if x != nil {
		return x.Unlocs
	}
	return nil
}

func (x *Port) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Port) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type PortResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *PortResponse) Reset() {
	*x = PortResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_port_port_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PortResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PortResponse) ProtoMessage() {}

func (x *PortResponse) ProtoReflect() protoreflect.Message {
	mi := &file_port_port_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PortResponse.ProtoReflect.Descriptor instead.
func (*PortResponse) Descriptor() ([]byte, []int) {
	return file_port_port_proto_rawDescGZIP(), []int{1}
}

func (x *PortResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type List struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *List) Reset() {
	*x = List{}
	if protoimpl.UnsafeEnabled {
		mi := &file_port_port_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *List) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*List) ProtoMessage() {}

func (x *List) ProtoReflect() protoreflect.Message {
	mi := &file_port_port_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use List.ProtoReflect.Descriptor instead.
func (*List) Descriptor() ([]byte, []int) {
	return file_port_port_proto_rawDescGZIP(), []int{2}
}

func (x *List) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_port_port_proto protoreflect.FileDescriptor

var file_port_port_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0a, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x22, 0x8e, 0x02,
	0x0a, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69,
	0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x61,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x07, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6f, 0x72,
	0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x02, 0x52, 0x0b, 0x63,
	0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f,
	0x6e, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f,
	0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x73, 0x18, 0x09, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x24,
	0x0a, 0x0c, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x22, 0x20, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x84, 0x01, 0x0a, 0x11, 0x50, 0x6f, 0x72, 0x74, 0x44,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x0a,
	0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x10, 0x2e, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x1a, 0x18, 0x2e, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x10, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69,
	0x64, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67,
	0x75, 0x69, 0x64, 0x65, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x22, 0x00, 0x30, 0x01, 0x42, 0x35, 0x0a,
	0x04, 0x70, 0x6f, 0x72, 0x74, 0x42, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x01, 0x5a, 0x25, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x72, 0x78, 0x64, 0x73, 0x69,
	0x6c, 0x76, 0x61, 0x2f, 0x6a, 0x70, 0x61, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f,
	0x70, 0x6f, 0x72, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_port_port_proto_rawDescOnce sync.Once
	file_port_port_proto_rawDescData = file_port_port_proto_rawDesc
)

func file_port_port_proto_rawDescGZIP() []byte {
	file_port_port_proto_rawDescOnce.Do(func() {
		file_port_port_proto_rawDescData = protoimpl.X.CompressGZIP(file_port_port_proto_rawDescData)
	})
	return file_port_port_proto_rawDescData
}

var file_port_port_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_port_port_proto_goTypes = []interface{}{
	(*Port)(nil),         // 0: routeguide.Port
	(*PortResponse)(nil), // 1: routeguide.PortResponse
	(*List)(nil),         // 2: routeguide.List
}
var file_port_port_proto_depIdxs = []int32{
	0, // 0: routeguide.PortDomainService.UpsertPort:input_type -> routeguide.Port
	2, // 1: routeguide.PortDomainService.ListPorts:input_type -> routeguide.List
	1, // 2: routeguide.PortDomainService.UpsertPort:output_type -> routeguide.PortResponse
	0, // 3: routeguide.PortDomainService.ListPorts:output_type -> routeguide.Port
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_port_port_proto_init() }
func file_port_port_proto_init() {
	if File_port_port_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_port_port_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Port); i {
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
		file_port_port_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PortResponse); i {
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
		file_port_port_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*List); i {
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
			RawDescriptor: file_port_port_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_port_port_proto_goTypes,
		DependencyIndexes: file_port_port_proto_depIdxs,
		MessageInfos:      file_port_port_proto_msgTypes,
	}.Build()
	File_port_port_proto = out.File
	file_port_port_proto_rawDesc = nil
	file_port_port_proto_goTypes = nil
	file_port_port_proto_depIdxs = nil
}