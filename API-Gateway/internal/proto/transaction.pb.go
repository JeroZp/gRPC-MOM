// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: transaction.proto

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

// Solicitud de transferencia
type TransferRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FromId        string                 `protobuf:"bytes,1,opt,name=from_id,json=fromId,proto3" json:"from_id,omitempty"`
	ToId          string                 `protobuf:"bytes,2,opt,name=to_id,json=toId,proto3" json:"to_id,omitempty"`
	Amount        int32                  `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TransferRequest) Reset() {
	*x = TransferRequest{}
	mi := &file_transaction_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TransferRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferRequest) ProtoMessage() {}

func (x *TransferRequest) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferRequest.ProtoReflect.Descriptor instead.
func (*TransferRequest) Descriptor() ([]byte, []int) {
	return file_transaction_proto_rawDescGZIP(), []int{0}
}

func (x *TransferRequest) GetFromId() string {
	if x != nil {
		return x.FromId
	}
	return ""
}

func (x *TransferRequest) GetToId() string {
	if x != nil {
		return x.ToId
	}
	return ""
}

func (x *TransferRequest) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

// Respuesta, con los dos usuarios actualizados
type TransferResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	FromUser      *User                  `protobuf:"bytes,3,opt,name=from_user,json=fromUser,proto3" json:"from_user,omitempty"`
	ToUser        *User                  `protobuf:"bytes,4,opt,name=to_user,json=toUser,proto3" json:"to_user,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TransferResponse) Reset() {
	*x = TransferResponse{}
	mi := &file_transaction_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TransferResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferResponse) ProtoMessage() {}

func (x *TransferResponse) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferResponse.ProtoReflect.Descriptor instead.
func (*TransferResponse) Descriptor() ([]byte, []int) {
	return file_transaction_proto_rawDescGZIP(), []int{1}
}

func (x *TransferResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *TransferResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *TransferResponse) GetFromUser() *User {
	if x != nil {
		return x.FromUser
	}
	return nil
}

func (x *TransferResponse) GetToUser() *User {
	if x != nil {
		return x.ToUser
	}
	return nil
}

var File_transaction_proto protoreflect.FileDescriptor

const file_transaction_proto_rawDesc = "" +
	"\n" +
	"\x11transaction.proto\x12\vtransaction\x1a\n" +
	"user.proto\"W\n" +
	"\x0fTransferRequest\x12\x17\n" +
	"\afrom_id\x18\x01 \x01(\tR\x06fromId\x12\x13\n" +
	"\x05to_id\x18\x02 \x01(\tR\x04toId\x12\x16\n" +
	"\x06amount\x18\x03 \x01(\x05R\x06amount\"\x94\x01\n" +
	"\x10TransferResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\x12'\n" +
	"\tfrom_user\x18\x03 \x01(\v2\n" +
	".user.UserR\bfromUser\x12#\n" +
	"\ato_user\x18\x04 \x01(\v2\n" +
	".user.UserR\x06toUser2]\n" +
	"\x12TransactionService\x12G\n" +
	"\bTransfer\x12\x1c.transaction.TransferRequest\x1a\x1d.transaction.TransferResponseB=Z;github.com/JeroZp/gRPC-MOM/API-Gateway/internal/proto;protob\x06proto3"

var (
	file_transaction_proto_rawDescOnce sync.Once
	file_transaction_proto_rawDescData []byte
)

func file_transaction_proto_rawDescGZIP() []byte {
	file_transaction_proto_rawDescOnce.Do(func() {
		file_transaction_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_transaction_proto_rawDesc), len(file_transaction_proto_rawDesc)))
	})
	return file_transaction_proto_rawDescData
}

var file_transaction_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_transaction_proto_goTypes = []any{
	(*TransferRequest)(nil),  // 0: transaction.TransferRequest
	(*TransferResponse)(nil), // 1: transaction.TransferResponse
	(*User)(nil),             // 2: user.User
}
var file_transaction_proto_depIdxs = []int32{
	2, // 0: transaction.TransferResponse.from_user:type_name -> user.User
	2, // 1: transaction.TransferResponse.to_user:type_name -> user.User
	0, // 2: transaction.TransactionService.Transfer:input_type -> transaction.TransferRequest
	1, // 3: transaction.TransactionService.Transfer:output_type -> transaction.TransferResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_transaction_proto_init() }
func file_transaction_proto_init() {
	if File_transaction_proto != nil {
		return
	}
	file_user_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_transaction_proto_rawDesc), len(file_transaction_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_transaction_proto_goTypes,
		DependencyIndexes: file_transaction_proto_depIdxs,
		MessageInfos:      file_transaction_proto_msgTypes,
	}.Build()
	File_transaction_proto = out.File
	file_transaction_proto_goTypes = nil
	file_transaction_proto_depIdxs = nil
}
