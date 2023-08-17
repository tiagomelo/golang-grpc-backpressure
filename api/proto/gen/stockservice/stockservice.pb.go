// Copyright (c) 2023 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.3
// source: stockservice.proto

package stockservice

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

// Empty message sent by the client to request stock updates.
type EmptyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyRequest) Reset() {
	*x = EmptyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stockservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyRequest) ProtoMessage() {}

func (x *EmptyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stockservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyRequest.ProtoReflect.Descriptor instead.
func (*EmptyRequest) Descriptor() ([]byte, []int) {
	return file_stockservice_proto_rawDescGZIP(), []int{0}
}

// Message containing detailed stock update information.
type StockUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ticker        string  `protobuf:"bytes,1,opt,name=ticker,proto3" json:"ticker,omitempty"`                 // Stock ticker symbol, e.g., "AAPL" for Apple Inc.
	Price         float64 `protobuf:"fixed64,2,opt,name=price,proto3" json:"price,omitempty"`                 // Current stock price.
	Change        float64 `protobuf:"fixed64,3,opt,name=change,proto3" json:"change,omitempty"`               // Price change since the last update.
	ChangePercent float64 `protobuf:"fixed64,4,opt,name=changePercent,proto3" json:"changePercent,omitempty"` // Price change percentage since the last update.
	Volume        int64   `protobuf:"varint,5,opt,name=volume,proto3" json:"volume,omitempty"`                // Trading volume for the current day.
	OpenPrice     float64 `protobuf:"fixed64,6,opt,name=openPrice,proto3" json:"openPrice,omitempty"`         // Opening price for the current trading session.
	HighPrice     float64 `protobuf:"fixed64,7,opt,name=highPrice,proto3" json:"highPrice,omitempty"`         // Highest price reached during the current trading session.
	LowPrice      float64 `protobuf:"fixed64,8,opt,name=lowPrice,proto3" json:"lowPrice,omitempty"`           // Lowest price reached during the current trading session.
	MarketCap     int64   `protobuf:"varint,9,opt,name=marketCap,proto3" json:"marketCap,omitempty"`          // Market capitalization.
	Timestamp     string  `protobuf:"bytes,10,opt,name=timestamp,proto3" json:"timestamp,omitempty"`          // Timestamp of the update, e.g., "2023-08-16T15:04:05Z".
}

func (x *StockUpdate) Reset() {
	*x = StockUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stockservice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StockUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockUpdate) ProtoMessage() {}

func (x *StockUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_stockservice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockUpdate.ProtoReflect.Descriptor instead.
func (*StockUpdate) Descriptor() ([]byte, []int) {
	return file_stockservice_proto_rawDescGZIP(), []int{1}
}

func (x *StockUpdate) GetTicker() string {
	if x != nil {
		return x.Ticker
	}
	return ""
}

func (x *StockUpdate) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *StockUpdate) GetChange() float64 {
	if x != nil {
		return x.Change
	}
	return 0
}

func (x *StockUpdate) GetChangePercent() float64 {
	if x != nil {
		return x.ChangePercent
	}
	return 0
}

func (x *StockUpdate) GetVolume() int64 {
	if x != nil {
		return x.Volume
	}
	return 0
}

func (x *StockUpdate) GetOpenPrice() float64 {
	if x != nil {
		return x.OpenPrice
	}
	return 0
}

func (x *StockUpdate) GetHighPrice() float64 {
	if x != nil {
		return x.HighPrice
	}
	return 0
}

func (x *StockUpdate) GetLowPrice() float64 {
	if x != nil {
		return x.LowPrice
	}
	return 0
}

func (x *StockUpdate) GetMarketCap() int64 {
	if x != nil {
		return x.MarketCap
	}
	return 0
}

func (x *StockUpdate) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

var File_stockservice_proto protoreflect.FileDescriptor

var file_stockservice_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x22, 0x0e, 0x0a, 0x0c, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0xa5, 0x02, 0x0a, 0x0b, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x06, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0d, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x6e, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x6e, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x68, 0x69, 0x67, 0x68, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x68, 0x69, 0x67, 0x68, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x77, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x6f, 0x77, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x43, 0x61, 0x70, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x43, 0x61, 0x70, 0x12, 0x1c, 0x0a, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x32, 0x55, 0x0a, 0x0c, 0x53, 0x74,
	0x6f, 0x63, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x12, 0x1a, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x30,
	0x01, 0x42, 0x4a, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x74, 0x69, 0x61, 0x67, 0x6f, 0x6d, 0x65, 0x6c, 0x6f, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x70, 0x72, 0x65, 0x73, 0x73, 0x75,
	0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stockservice_proto_rawDescOnce sync.Once
	file_stockservice_proto_rawDescData = file_stockservice_proto_rawDesc
)

func file_stockservice_proto_rawDescGZIP() []byte {
	file_stockservice_proto_rawDescOnce.Do(func() {
		file_stockservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_stockservice_proto_rawDescData)
	})
	return file_stockservice_proto_rawDescData
}

var file_stockservice_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_stockservice_proto_goTypes = []interface{}{
	(*EmptyRequest)(nil), // 0: stockservice.EmptyRequest
	(*StockUpdate)(nil),  // 1: stockservice.StockUpdate
}
var file_stockservice_proto_depIdxs = []int32{
	0, // 0: stockservice.StockService.GetUpdates:input_type -> stockservice.EmptyRequest
	1, // 1: stockservice.StockService.GetUpdates:output_type -> stockservice.StockUpdate
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_stockservice_proto_init() }
func file_stockservice_proto_init() {
	if File_stockservice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stockservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyRequest); i {
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
		file_stockservice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StockUpdate); i {
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
			RawDescriptor: file_stockservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stockservice_proto_goTypes,
		DependencyIndexes: file_stockservice_proto_depIdxs,
		MessageInfos:      file_stockservice_proto_msgTypes,
	}.Build()
	File_stockservice_proto = out.File
	file_stockservice_proto_rawDesc = nil
	file_stockservice_proto_goTypes = nil
	file_stockservice_proto_depIdxs = nil
}
