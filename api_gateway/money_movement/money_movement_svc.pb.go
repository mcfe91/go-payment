// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: proto/money_movement_svc.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AuthorizePayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerWalletUserId string `protobuf:"bytes,1,opt,name=customerWalletUserId,proto3" json:"customerWalletUserId,omitempty"`
	MerchantWalletUserId string `protobuf:"bytes,2,opt,name=merchantWalletUserId,proto3" json:"merchantWalletUserId,omitempty"`
	Cents                int64  `protobuf:"varint,3,opt,name=cents,proto3" json:"cents,omitempty"`
	Currency             string `protobuf:"bytes,4,opt,name=currency,proto3" json:"currency,omitempty"`
}

func (x *AuthorizePayload) Reset() {
	*x = AuthorizePayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_money_movement_svc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorizePayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizePayload) ProtoMessage() {}

func (x *AuthorizePayload) ProtoReflect() protoreflect.Message {
	mi := &file_proto_money_movement_svc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizePayload.ProtoReflect.Descriptor instead.
func (*AuthorizePayload) Descriptor() ([]byte, []int) {
	return file_proto_money_movement_svc_proto_rawDescGZIP(), []int{0}
}

func (x *AuthorizePayload) GetCustomerWalletUserId() string {
	if x != nil {
		return x.CustomerWalletUserId
	}
	return ""
}

func (x *AuthorizePayload) GetMerchantWalletUserId() string {
	if x != nil {
		return x.MerchantWalletUserId
	}
	return ""
}

func (x *AuthorizePayload) GetCents() int64 {
	if x != nil {
		return x.Cents
	}
	return 0
}

func (x *AuthorizePayload) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

type CapturePayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid string `protobuf:"bytes,1,opt,name=pid,proto3" json:"pid,omitempty"`
}

func (x *CapturePayload) Reset() {
	*x = CapturePayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_money_movement_svc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CapturePayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CapturePayload) ProtoMessage() {}

func (x *CapturePayload) ProtoReflect() protoreflect.Message {
	mi := &file_proto_money_movement_svc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CapturePayload.ProtoReflect.Descriptor instead.
func (*CapturePayload) Descriptor() ([]byte, []int) {
	return file_proto_money_movement_svc_proto_rawDescGZIP(), []int{1}
}

func (x *CapturePayload) GetPid() string {
	if x != nil {
		return x.Pid
	}
	return ""
}

type AuthorizeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid string `protobuf:"bytes,1,opt,name=pid,proto3" json:"pid,omitempty"`
}

func (x *AuthorizeResponse) Reset() {
	*x = AuthorizeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_money_movement_svc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorizeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizeResponse) ProtoMessage() {}

func (x *AuthorizeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_money_movement_svc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizeResponse.ProtoReflect.Descriptor instead.
func (*AuthorizeResponse) Descriptor() ([]byte, []int) {
	return file_proto_money_movement_svc_proto_rawDescGZIP(), []int{2}
}

func (x *AuthorizeResponse) GetPid() string {
	if x != nil {
		return x.Pid
	}
	return ""
}

var File_proto_money_movement_svc_proto protoreflect.FileDescriptor

var file_proto_money_movement_svc_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x5f, 0x6d, 0x6f,
	0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x76, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xac, 0x01,
	0x0a, 0x10, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x12, 0x32, 0x0a, 0x14, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x57, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x14, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x14, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61,
	0x6e, 0x74, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x57, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x65,
	0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x65, 0x6e, 0x74, 0x73,
	0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x22, 0x22, 0x0a, 0x0e,
	0x43, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x69, 0x64,
	0x22, 0x25, 0x0a, 0x11, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x70, 0x69, 0x64, 0x32, 0x82, 0x01, 0x0a, 0x14, 0x4d, 0x6f, 0x6e, 0x65,
	0x79, 0x4d, 0x6f, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x34, 0x0a, 0x09, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x12, 0x11, 0x2e,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x1a, 0x12, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x07, 0x43, 0x61, 0x70, 0x74, 0x75, 0x72,
	0x65, 0x12, 0x0f, 0x2e, 0x43, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x37, 0x5a, 0x35,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x6d, 0x63, 0x66, 0x65,
	0x72, 0x72, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2d, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2f, 0x6d, 0x6f,
	0x6e, 0x65, 0x79, 0x5f, 0x6d, 0x6f, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_money_movement_svc_proto_rawDescOnce sync.Once
	file_proto_money_movement_svc_proto_rawDescData = file_proto_money_movement_svc_proto_rawDesc
)

func file_proto_money_movement_svc_proto_rawDescGZIP() []byte {
	file_proto_money_movement_svc_proto_rawDescOnce.Do(func() {
		file_proto_money_movement_svc_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_money_movement_svc_proto_rawDescData)
	})
	return file_proto_money_movement_svc_proto_rawDescData
}

var file_proto_money_movement_svc_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_money_movement_svc_proto_goTypes = []interface{}{
	(*AuthorizePayload)(nil),  // 0: AuthorizePayload
	(*CapturePayload)(nil),    // 1: CapturePayload
	(*AuthorizeResponse)(nil), // 2: AuthorizeResponse
	(*emptypb.Empty)(nil),     // 3: google.protobuf.Empty
}
var file_proto_money_movement_svc_proto_depIdxs = []int32{
	0, // 0: MoneyMovementService.Authorize:input_type -> AuthorizePayload
	1, // 1: MoneyMovementService.Capture:input_type -> CapturePayload
	2, // 2: MoneyMovementService.Authorize:output_type -> AuthorizeResponse
	3, // 3: MoneyMovementService.Capture:output_type -> google.protobuf.Empty
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_money_movement_svc_proto_init() }
func file_proto_money_movement_svc_proto_init() {
	if File_proto_money_movement_svc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_money_movement_svc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorizePayload); i {
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
		file_proto_money_movement_svc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CapturePayload); i {
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
		file_proto_money_movement_svc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorizeResponse); i {
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
			RawDescriptor: file_proto_money_movement_svc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_money_movement_svc_proto_goTypes,
		DependencyIndexes: file_proto_money_movement_svc_proto_depIdxs,
		MessageInfos:      file_proto_money_movement_svc_proto_msgTypes,
	}.Build()
	File_proto_money_movement_svc_proto = out.File
	file_proto_money_movement_svc_proto_rawDesc = nil
	file_proto_money_movement_svc_proto_goTypes = nil
	file_proto_money_movement_svc_proto_depIdxs = nil
}