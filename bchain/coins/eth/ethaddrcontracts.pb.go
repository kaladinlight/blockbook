// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: bchain/coins/eth/ethaddrcontracts.proto

package eth

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

type ProtoAddrContracts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalTxs       uint64                             `protobuf:"varint,1,opt,name=TotalTxs,proto3" json:"TotalTxs,omitempty"`
	NonContractTxs uint64                             `protobuf:"varint,2,opt,name=NonContractTxs,proto3" json:"NonContractTxs,omitempty"`
	InternalTxs    uint64                             `protobuf:"varint,3,opt,name=InternalTxs,proto3" json:"InternalTxs,omitempty"`
	Contracts      []*ProtoAddrContracts_AddrContract `protobuf:"bytes,4,rep,name=Contracts,proto3" json:"Contracts,omitempty"`
}

func (x *ProtoAddrContracts) Reset() {
	*x = ProtoAddrContracts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bchain_coins_eth_ethaddrcontracts_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoAddrContracts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoAddrContracts) ProtoMessage() {}

func (x *ProtoAddrContracts) ProtoReflect() protoreflect.Message {
	mi := &file_bchain_coins_eth_ethaddrcontracts_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoAddrContracts.ProtoReflect.Descriptor instead.
func (*ProtoAddrContracts) Descriptor() ([]byte, []int) {
	return file_bchain_coins_eth_ethaddrcontracts_proto_rawDescGZIP(), []int{0}
}

func (x *ProtoAddrContracts) GetTotalTxs() uint64 {
	if x != nil {
		return x.TotalTxs
	}
	return 0
}

func (x *ProtoAddrContracts) GetNonContractTxs() uint64 {
	if x != nil {
		return x.NonContractTxs
	}
	return 0
}

func (x *ProtoAddrContracts) GetInternalTxs() uint64 {
	if x != nil {
		return x.InternalTxs
	}
	return 0
}

func (x *ProtoAddrContracts) GetContracts() []*ProtoAddrContracts_AddrContract {
	if x != nil {
		return x.Contracts
	}
	return nil
}

type ProtoAddrContracts_MultiTokenValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    []byte `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *ProtoAddrContracts_MultiTokenValue) Reset() {
	*x = ProtoAddrContracts_MultiTokenValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bchain_coins_eth_ethaddrcontracts_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoAddrContracts_MultiTokenValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoAddrContracts_MultiTokenValue) ProtoMessage() {}

func (x *ProtoAddrContracts_MultiTokenValue) ProtoReflect() protoreflect.Message {
	mi := &file_bchain_coins_eth_ethaddrcontracts_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoAddrContracts_MultiTokenValue.ProtoReflect.Descriptor instead.
func (*ProtoAddrContracts_MultiTokenValue) Descriptor() ([]byte, []int) {
	return file_bchain_coins_eth_ethaddrcontracts_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ProtoAddrContracts_MultiTokenValue) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *ProtoAddrContracts_MultiTokenValue) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type ProtoAddrContracts_AddrContract struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type             int64                                 `protobuf:"varint,1,opt,name=Type,proto3" json:"Type,omitempty"`
	Contract         []byte                                `protobuf:"bytes,2,opt,name=Contract,proto3" json:"Contract,omitempty"`
	Txs              uint64                                `protobuf:"varint,3,opt,name=Txs,proto3" json:"Txs,omitempty"`
	Value            []byte                                `protobuf:"bytes,4,opt,name=Value,proto3" json:"Value,omitempty"`
	Ids              [][]byte                              `protobuf:"bytes,5,rep,name=Ids,proto3" json:"Ids,omitempty"`
	MultiTokenValues []*ProtoAddrContracts_MultiTokenValue `protobuf:"bytes,6,rep,name=MultiTokenValues,proto3" json:"MultiTokenValues,omitempty"`
}

func (x *ProtoAddrContracts_AddrContract) Reset() {
	*x = ProtoAddrContracts_AddrContract{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bchain_coins_eth_ethaddrcontracts_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoAddrContracts_AddrContract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoAddrContracts_AddrContract) ProtoMessage() {}

func (x *ProtoAddrContracts_AddrContract) ProtoReflect() protoreflect.Message {
	mi := &file_bchain_coins_eth_ethaddrcontracts_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoAddrContracts_AddrContract.ProtoReflect.Descriptor instead.
func (*ProtoAddrContracts_AddrContract) Descriptor() ([]byte, []int) {
	return file_bchain_coins_eth_ethaddrcontracts_proto_rawDescGZIP(), []int{0, 1}
}

func (x *ProtoAddrContracts_AddrContract) GetType() int64 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *ProtoAddrContracts_AddrContract) GetContract() []byte {
	if x != nil {
		return x.Contract
	}
	return nil
}

func (x *ProtoAddrContracts_AddrContract) GetTxs() uint64 {
	if x != nil {
		return x.Txs
	}
	return 0
}

func (x *ProtoAddrContracts_AddrContract) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *ProtoAddrContracts_AddrContract) GetIds() [][]byte {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *ProtoAddrContracts_AddrContract) GetMultiTokenValues() []*ProtoAddrContracts_MultiTokenValue {
	if x != nil {
		return x.MultiTokenValues
	}
	return nil
}

var File_bchain_coins_eth_ethaddrcontracts_proto protoreflect.FileDescriptor

var file_bchain_coins_eth_ethaddrcontracts_proto_rawDesc = []byte{
	0x0a, 0x27, 0x62, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x73, 0x2f, 0x65,
	0x74, 0x68, 0x2f, 0x65, 0x74, 0x68, 0x61, 0x64, 0x64, 0x72, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbf, 0x03, 0x0a, 0x12, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73,
	0x12, 0x1a, 0x0a, 0x08, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x54, 0x78, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x08, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x54, 0x78, 0x73, 0x12, 0x26, 0x0a, 0x0e,
	0x4e, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x54, 0x78, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x4e, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x54, 0x78, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x54, 0x78, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x54, 0x78, 0x73, 0x12, 0x3e, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x41, 0x64, 0x64, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x41,
	0x64, 0x64, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x09, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x1a, 0x37, 0x0a, 0x0f, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a,
	0xc9, 0x01, 0x0a, 0x0c, 0x41, 0x64, 0x64, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x54, 0x78, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x54,
	0x78, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x49, 0x64, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x03, 0x49, 0x64, 0x73, 0x12, 0x4f, 0x0a, 0x10, 0x4d, 0x75,
	0x6c, 0x74, 0x69, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x41, 0x64, 0x64, 0x72,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x10, 0x4d, 0x75, 0x6c, 0x74, 0x69,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x42, 0x2e, 0x5a, 0x2c, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x72, 0x65, 0x7a, 0x6f, 0x72,
	0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x62, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x73, 0x2f, 0x65, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_bchain_coins_eth_ethaddrcontracts_proto_rawDescOnce sync.Once
	file_bchain_coins_eth_ethaddrcontracts_proto_rawDescData = file_bchain_coins_eth_ethaddrcontracts_proto_rawDesc
)

func file_bchain_coins_eth_ethaddrcontracts_proto_rawDescGZIP() []byte {
	file_bchain_coins_eth_ethaddrcontracts_proto_rawDescOnce.Do(func() {
		file_bchain_coins_eth_ethaddrcontracts_proto_rawDescData = protoimpl.X.CompressGZIP(file_bchain_coins_eth_ethaddrcontracts_proto_rawDescData)
	})
	return file_bchain_coins_eth_ethaddrcontracts_proto_rawDescData
}

var file_bchain_coins_eth_ethaddrcontracts_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_bchain_coins_eth_ethaddrcontracts_proto_goTypes = []interface{}{
	(*ProtoAddrContracts)(nil),                 // 0: ProtoAddrContracts
	(*ProtoAddrContracts_MultiTokenValue)(nil), // 1: ProtoAddrContracts.MultiTokenValue
	(*ProtoAddrContracts_AddrContract)(nil),    // 2: ProtoAddrContracts.AddrContract
}
var file_bchain_coins_eth_ethaddrcontracts_proto_depIdxs = []int32{
	2, // 0: ProtoAddrContracts.Contracts:type_name -> ProtoAddrContracts.AddrContract
	1, // 1: ProtoAddrContracts.AddrContract.MultiTokenValues:type_name -> ProtoAddrContracts.MultiTokenValue
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_bchain_coins_eth_ethaddrcontracts_proto_init() }
func file_bchain_coins_eth_ethaddrcontracts_proto_init() {
	if File_bchain_coins_eth_ethaddrcontracts_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bchain_coins_eth_ethaddrcontracts_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoAddrContracts); i {
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
		file_bchain_coins_eth_ethaddrcontracts_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoAddrContracts_MultiTokenValue); i {
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
		file_bchain_coins_eth_ethaddrcontracts_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoAddrContracts_AddrContract); i {
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
			RawDescriptor: file_bchain_coins_eth_ethaddrcontracts_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bchain_coins_eth_ethaddrcontracts_proto_goTypes,
		DependencyIndexes: file_bchain_coins_eth_ethaddrcontracts_proto_depIdxs,
		MessageInfos:      file_bchain_coins_eth_ethaddrcontracts_proto_msgTypes,
	}.Build()
	File_bchain_coins_eth_ethaddrcontracts_proto = out.File
	file_bchain_coins_eth_ethaddrcontracts_proto_rawDesc = nil
	file_bchain_coins_eth_ethaddrcontracts_proto_goTypes = nil
	file_bchain_coins_eth_ethaddrcontracts_proto_depIdxs = nil
}