// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: longmsg.proto

package longmsg

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

type LongMsgDeleteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsgResid []byte `protobuf:"bytes,1,opt,name=msgResid,proto3" json:"msgResid,omitempty"`
	MsgType  int32  `protobuf:"varint,2,opt,name=msgType,proto3" json:"msgType,omitempty"`
}

func (x *LongMsgDeleteReq) Reset() {
	*x = LongMsgDeleteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_longmsg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LongMsgDeleteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LongMsgDeleteReq) ProtoMessage() {}

func (x *LongMsgDeleteReq) ProtoReflect() protoreflect.Message {
	mi := &file_longmsg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LongMsgDeleteReq.ProtoReflect.Descriptor instead.
func (*LongMsgDeleteReq) Descriptor() ([]byte, []int) {
	return file_longmsg_proto_rawDescGZIP(), []int{0}
}

func (x *LongMsgDeleteReq) GetMsgResid() []byte {
	if x != nil {
		return x.MsgResid
	}
	return nil
}

func (x *LongMsgDeleteReq) GetMsgType() int32 {
	if x != nil {
		return x.MsgType
	}
	return 0
}

type LongMsgDeleteRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result   int32  `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	MsgResid []byte `protobuf:"bytes,2,opt,name=msgResid,proto3" json:"msgResid,omitempty"`
}

func (x *LongMsgDeleteRsp) Reset() {
	*x = LongMsgDeleteRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_longmsg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LongMsgDeleteRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LongMsgDeleteRsp) ProtoMessage() {}

func (x *LongMsgDeleteRsp) ProtoReflect() protoreflect.Message {
	mi := &file_longmsg_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LongMsgDeleteRsp.ProtoReflect.Descriptor instead.
func (*LongMsgDeleteRsp) Descriptor() ([]byte, []int) {
	return file_longmsg_proto_rawDescGZIP(), []int{1}
}

func (x *LongMsgDeleteRsp) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

func (x *LongMsgDeleteRsp) GetMsgResid() []byte {
	if x != nil {
		return x.MsgResid
	}
	return nil
}

type LongMsgDownReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SrcUin    int32  `protobuf:"varint,1,opt,name=srcUin,proto3" json:"srcUin,omitempty"`
	MsgResid  []byte `protobuf:"bytes,2,opt,name=msgResid,proto3" json:"msgResid,omitempty"`
	MsgType   int32  `protobuf:"varint,3,opt,name=msgType,proto3" json:"msgType,omitempty"`
	NeedCache int32  `protobuf:"varint,4,opt,name=needCache,proto3" json:"needCache,omitempty"`
}

func (x *LongMsgDownReq) Reset() {
	*x = LongMsgDownReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_longmsg_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LongMsgDownReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LongMsgDownReq) ProtoMessage() {}

func (x *LongMsgDownReq) ProtoReflect() protoreflect.Message {
	mi := &file_longmsg_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LongMsgDownReq.ProtoReflect.Descriptor instead.
func (*LongMsgDownReq) Descriptor() ([]byte, []int) {
	return file_longmsg_proto_rawDescGZIP(), []int{2}
}

func (x *LongMsgDownReq) GetSrcUin() int32 {
	if x != nil {
		return x.SrcUin
	}
	return 0
}

func (x *LongMsgDownReq) GetMsgResid() []byte {
	if x != nil {
		return x.MsgResid
	}
	return nil
}

func (x *LongMsgDownReq) GetMsgType() int32 {
	if x != nil {
		return x.MsgType
	}
	return 0
}

func (x *LongMsgDownReq) GetNeedCache() int32 {
	if x != nil {
		return x.NeedCache
	}
	return 0
}

type LongMsgDownRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result     int32  `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	MsgResid   []byte `protobuf:"bytes,2,opt,name=msgResid,proto3" json:"msgResid,omitempty"`
	MsgContent []byte `protobuf:"bytes,3,opt,name=msgContent,proto3" json:"msgContent,omitempty"`
}

func (x *LongMsgDownRsp) Reset() {
	*x = LongMsgDownRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_longmsg_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LongMsgDownRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LongMsgDownRsp) ProtoMessage() {}

func (x *LongMsgDownRsp) ProtoReflect() protoreflect.Message {
	mi := &file_longmsg_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LongMsgDownRsp.ProtoReflect.Descriptor instead.
func (*LongMsgDownRsp) Descriptor() ([]byte, []int) {
	return file_longmsg_proto_rawDescGZIP(), []int{3}
}

func (x *LongMsgDownRsp) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

func (x *LongMsgDownRsp) GetMsgResid() []byte {
	if x != nil {
		return x.MsgResid
	}
	return nil
}

func (x *LongMsgDownRsp) GetMsgContent() []byte {
	if x != nil {
		return x.MsgContent
	}
	return nil
}

type LongMsgUpReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsgType    int32  `protobuf:"varint,1,opt,name=msgType,proto3" json:"msgType,omitempty"`
	DstUin     int64  `protobuf:"varint,2,opt,name=dstUin,proto3" json:"dstUin,omitempty"`
	MsgId      int32  `protobuf:"varint,3,opt,name=msgId,proto3" json:"msgId,omitempty"`
	MsgContent []byte `protobuf:"bytes,4,opt,name=msgContent,proto3" json:"msgContent,omitempty"`
	StoreType  int32  `protobuf:"varint,5,opt,name=storeType,proto3" json:"storeType,omitempty"`
	MsgUkey    []byte `protobuf:"bytes,6,opt,name=msgUkey,proto3" json:"msgUkey,omitempty"`
	NeedCache  int32  `protobuf:"varint,7,opt,name=needCache,proto3" json:"needCache,omitempty"`
}

func (x *LongMsgUpReq) Reset() {
	*x = LongMsgUpReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_longmsg_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LongMsgUpReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LongMsgUpReq) ProtoMessage() {}

func (x *LongMsgUpReq) ProtoReflect() protoreflect.Message {
	mi := &file_longmsg_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LongMsgUpReq.ProtoReflect.Descriptor instead.
func (*LongMsgUpReq) Descriptor() ([]byte, []int) {
	return file_longmsg_proto_rawDescGZIP(), []int{4}
}

func (x *LongMsgUpReq) GetMsgType() int32 {
	if x != nil {
		return x.MsgType
	}
	return 0
}

func (x *LongMsgUpReq) GetDstUin() int64 {
	if x != nil {
		return x.DstUin
	}
	return 0
}

func (x *LongMsgUpReq) GetMsgId() int32 {
	if x != nil {
		return x.MsgId
	}
	return 0
}

func (x *LongMsgUpReq) GetMsgContent() []byte {
	if x != nil {
		return x.MsgContent
	}
	return nil
}

func (x *LongMsgUpReq) GetStoreType() int32 {
	if x != nil {
		return x.StoreType
	}
	return 0
}

func (x *LongMsgUpReq) GetMsgUkey() []byte {
	if x != nil {
		return x.MsgUkey
	}
	return nil
}

func (x *LongMsgUpReq) GetNeedCache() int32 {
	if x != nil {
		return x.NeedCache
	}
	return 0
}

type LongMsgUpRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result   int32  `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	MsgId    int32  `protobuf:"varint,2,opt,name=msgId,proto3" json:"msgId,omitempty"`
	MsgResid []byte `protobuf:"bytes,3,opt,name=msgResid,proto3" json:"msgResid,omitempty"`
}

func (x *LongMsgUpRsp) Reset() {
	*x = LongMsgUpRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_longmsg_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LongMsgUpRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LongMsgUpRsp) ProtoMessage() {}

func (x *LongMsgUpRsp) ProtoReflect() protoreflect.Message {
	mi := &file_longmsg_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LongMsgUpRsp.ProtoReflect.Descriptor instead.
func (*LongMsgUpRsp) Descriptor() ([]byte, []int) {
	return file_longmsg_proto_rawDescGZIP(), []int{5}
}

func (x *LongMsgUpRsp) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

func (x *LongMsgUpRsp) GetMsgId() int32 {
	if x != nil {
		return x.MsgId
	}
	return 0
}

func (x *LongMsgUpRsp) GetMsgResid() []byte {
	if x != nil {
		return x.MsgResid
	}
	return nil
}

type LongReqBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subcmd       int32               `protobuf:"varint,1,opt,name=subcmd,proto3" json:"subcmd,omitempty"`
	TermType     int32               `protobuf:"varint,2,opt,name=termType,proto3" json:"termType,omitempty"`
	PlatformType int32               `protobuf:"varint,3,opt,name=platformType,proto3" json:"platformType,omitempty"`
	MsgUpReq     []*LongMsgUpReq     `protobuf:"bytes,4,rep,name=msgUpReq,proto3" json:"msgUpReq,omitempty"`
	MsgDownReq   []*LongMsgDownReq   `protobuf:"bytes,5,rep,name=msgDownReq,proto3" json:"msgDownReq,omitempty"`
	MsgDelReq    []*LongMsgDeleteReq `protobuf:"bytes,6,rep,name=msgDelReq,proto3" json:"msgDelReq,omitempty"`
	AgentType    int32               `protobuf:"varint,10,opt,name=agentType,proto3" json:"agentType,omitempty"`
}

func (x *LongReqBody) Reset() {
	*x = LongReqBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_longmsg_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LongReqBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LongReqBody) ProtoMessage() {}

func (x *LongReqBody) ProtoReflect() protoreflect.Message {
	mi := &file_longmsg_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LongReqBody.ProtoReflect.Descriptor instead.
func (*LongReqBody) Descriptor() ([]byte, []int) {
	return file_longmsg_proto_rawDescGZIP(), []int{6}
}

func (x *LongReqBody) GetSubcmd() int32 {
	if x != nil {
		return x.Subcmd
	}
	return 0
}

func (x *LongReqBody) GetTermType() int32 {
	if x != nil {
		return x.TermType
	}
	return 0
}

func (x *LongReqBody) GetPlatformType() int32 {
	if x != nil {
		return x.PlatformType
	}
	return 0
}

func (x *LongReqBody) GetMsgUpReq() []*LongMsgUpReq {
	if x != nil {
		return x.MsgUpReq
	}
	return nil
}

func (x *LongReqBody) GetMsgDownReq() []*LongMsgDownReq {
	if x != nil {
		return x.MsgDownReq
	}
	return nil
}

func (x *LongReqBody) GetMsgDelReq() []*LongMsgDeleteReq {
	if x != nil {
		return x.MsgDelReq
	}
	return nil
}

func (x *LongReqBody) GetAgentType() int32 {
	if x != nil {
		return x.AgentType
	}
	return 0
}

type LongRspBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subcmd     int32               `protobuf:"varint,1,opt,name=subcmd,proto3" json:"subcmd,omitempty"`
	MsgUpRsp   []*LongMsgUpRsp     `protobuf:"bytes,2,rep,name=msgUpRsp,proto3" json:"msgUpRsp,omitempty"`
	MsgDownRsp []*LongMsgDownRsp   `protobuf:"bytes,3,rep,name=msgDownRsp,proto3" json:"msgDownRsp,omitempty"`
	MsgDelRsp  []*LongMsgDeleteRsp `protobuf:"bytes,4,rep,name=msgDelRsp,proto3" json:"msgDelRsp,omitempty"`
}

func (x *LongRspBody) Reset() {
	*x = LongRspBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_longmsg_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LongRspBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LongRspBody) ProtoMessage() {}

func (x *LongRspBody) ProtoReflect() protoreflect.Message {
	mi := &file_longmsg_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LongRspBody.ProtoReflect.Descriptor instead.
func (*LongRspBody) Descriptor() ([]byte, []int) {
	return file_longmsg_proto_rawDescGZIP(), []int{7}
}

func (x *LongRspBody) GetSubcmd() int32 {
	if x != nil {
		return x.Subcmd
	}
	return 0
}

func (x *LongRspBody) GetMsgUpRsp() []*LongMsgUpRsp {
	if x != nil {
		return x.MsgUpRsp
	}
	return nil
}

func (x *LongRspBody) GetMsgDownRsp() []*LongMsgDownRsp {
	if x != nil {
		return x.MsgDownRsp
	}
	return nil
}

func (x *LongRspBody) GetMsgDelRsp() []*LongMsgDeleteRsp {
	if x != nil {
		return x.MsgDelRsp
	}
	return nil
}

var File_longmsg_proto protoreflect.FileDescriptor

var file_longmsg_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6c, 0x6f, 0x6e, 0x67, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x48, 0x0a, 0x10, 0x4c, 0x6f, 0x6e, 0x67, 0x4d, 0x73, 0x67, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x73, 0x67, 0x52, 0x65, 0x73, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6d, 0x73, 0x67, 0x52, 0x65, 0x73, 0x69, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x6d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x22, 0x46, 0x0a, 0x10, 0x4c, 0x6f, 0x6e,
	0x67, 0x4d, 0x73, 0x67, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x73, 0x70, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x73, 0x67, 0x52, 0x65, 0x73, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6d, 0x73, 0x67, 0x52, 0x65, 0x73, 0x69,
	0x64, 0x22, 0x7c, 0x0a, 0x0e, 0x4c, 0x6f, 0x6e, 0x67, 0x4d, 0x73, 0x67, 0x44, 0x6f, 0x77, 0x6e,
	0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x72, 0x63, 0x55, 0x69, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x72, 0x63, 0x55, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6d,
	0x73, 0x67, 0x52, 0x65, 0x73, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6d,
	0x73, 0x67, 0x52, 0x65, 0x73, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x73, 0x67, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d, 0x73, 0x67, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x65, 0x65, 0x64, 0x43, 0x61, 0x63, 0x68, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6e, 0x65, 0x65, 0x64, 0x43, 0x61, 0x63, 0x68, 0x65, 0x22,
	0x64, 0x0a, 0x0e, 0x4c, 0x6f, 0x6e, 0x67, 0x4d, 0x73, 0x67, 0x44, 0x6f, 0x77, 0x6e, 0x52, 0x73,
	0x70, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x73, 0x67,
	0x52, 0x65, 0x73, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6d, 0x73, 0x67,
	0x52, 0x65, 0x73, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x73, 0x67, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x6d, 0x73, 0x67, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0xcc, 0x01, 0x0a, 0x0c, 0x4c, 0x6f, 0x6e, 0x67, 0x4d, 0x73,
	0x67, 0x55, 0x70, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x73, 0x67, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x64, 0x73, 0x74, 0x55, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x64, 0x73, 0x74, 0x55, 0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x73, 0x67, 0x49,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6d, 0x73, 0x67, 0x49, 0x64, 0x12, 0x1e,
	0x0a, 0x0a, 0x6d, 0x73, 0x67, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0a, 0x6d, 0x73, 0x67, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x73, 0x67, 0x55, 0x6b, 0x65, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x6d,
	0x73, 0x67, 0x55, 0x6b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x65, 0x65, 0x64, 0x43, 0x61,
	0x63, 0x68, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6e, 0x65, 0x65, 0x64, 0x43,
	0x61, 0x63, 0x68, 0x65, 0x22, 0x58, 0x0a, 0x0c, 0x4c, 0x6f, 0x6e, 0x67, 0x4d, 0x73, 0x67, 0x55,
	0x70, 0x52, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x6d, 0x73, 0x67, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6d, 0x73, 0x67,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x73, 0x67, 0x52, 0x65, 0x73, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6d, 0x73, 0x67, 0x52, 0x65, 0x73, 0x69, 0x64, 0x22, 0x90,
	0x02, 0x0a, 0x0b, 0x4c, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x75, 0x62, 0x63, 0x6d, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x73, 0x75, 0x62, 0x63, 0x6d, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x72, 0x6d, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x74, 0x65, 0x72, 0x6d, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x29, 0x0a, 0x08, 0x6d, 0x73, 0x67, 0x55, 0x70, 0x52,
	0x65, 0x71, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x4c, 0x6f, 0x6e, 0x67, 0x4d,
	0x73, 0x67, 0x55, 0x70, 0x52, 0x65, 0x71, 0x52, 0x08, 0x6d, 0x73, 0x67, 0x55, 0x70, 0x52, 0x65,
	0x71, 0x12, 0x2f, 0x0a, 0x0a, 0x6d, 0x73, 0x67, 0x44, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x71, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x4c, 0x6f, 0x6e, 0x67, 0x4d, 0x73, 0x67, 0x44,
	0x6f, 0x77, 0x6e, 0x52, 0x65, 0x71, 0x52, 0x0a, 0x6d, 0x73, 0x67, 0x44, 0x6f, 0x77, 0x6e, 0x52,
	0x65, 0x71, 0x12, 0x2f, 0x0a, 0x09, 0x6d, 0x73, 0x67, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x4c, 0x6f, 0x6e, 0x67, 0x4d, 0x73, 0x67, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x52, 0x09, 0x6d, 0x73, 0x67, 0x44, 0x65, 0x6c,
	0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x22, 0xb2, 0x01, 0x0a, 0x0b, 0x4c, 0x6f, 0x6e, 0x67, 0x52, 0x73, 0x70, 0x42, 0x6f, 0x64,
	0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x75, 0x62, 0x63, 0x6d, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x73, 0x75, 0x62, 0x63, 0x6d, 0x64, 0x12, 0x29, 0x0a, 0x08, 0x6d, 0x73, 0x67,
	0x55, 0x70, 0x52, 0x73, 0x70, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x4c, 0x6f,
	0x6e, 0x67, 0x4d, 0x73, 0x67, 0x55, 0x70, 0x52, 0x73, 0x70, 0x52, 0x08, 0x6d, 0x73, 0x67, 0x55,
	0x70, 0x52, 0x73, 0x70, 0x12, 0x2f, 0x0a, 0x0a, 0x6d, 0x73, 0x67, 0x44, 0x6f, 0x77, 0x6e, 0x52,
	0x73, 0x70, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x4c, 0x6f, 0x6e, 0x67, 0x4d,
	0x73, 0x67, 0x44, 0x6f, 0x77, 0x6e, 0x52, 0x73, 0x70, 0x52, 0x0a, 0x6d, 0x73, 0x67, 0x44, 0x6f,
	0x77, 0x6e, 0x52, 0x73, 0x70, 0x12, 0x2f, 0x0a, 0x09, 0x6d, 0x73, 0x67, 0x44, 0x65, 0x6c, 0x52,
	0x73, 0x70, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x4c, 0x6f, 0x6e, 0x67, 0x4d,
	0x73, 0x67, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x73, 0x70, 0x52, 0x09, 0x6d, 0x73, 0x67,
	0x44, 0x65, 0x6c, 0x52, 0x73, 0x70, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x3b, 0x6c, 0x6f, 0x6e, 0x67,
	0x6d, 0x73, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_longmsg_proto_rawDescOnce sync.Once
	file_longmsg_proto_rawDescData = file_longmsg_proto_rawDesc
)

func file_longmsg_proto_rawDescGZIP() []byte {
	file_longmsg_proto_rawDescOnce.Do(func() {
		file_longmsg_proto_rawDescData = protoimpl.X.CompressGZIP(file_longmsg_proto_rawDescData)
	})
	return file_longmsg_proto_rawDescData
}

var file_longmsg_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_longmsg_proto_goTypes = []interface{}{
	(*LongMsgDeleteReq)(nil), // 0: LongMsgDeleteReq
	(*LongMsgDeleteRsp)(nil), // 1: LongMsgDeleteRsp
	(*LongMsgDownReq)(nil),   // 2: LongMsgDownReq
	(*LongMsgDownRsp)(nil),   // 3: LongMsgDownRsp
	(*LongMsgUpReq)(nil),     // 4: LongMsgUpReq
	(*LongMsgUpRsp)(nil),     // 5: LongMsgUpRsp
	(*LongReqBody)(nil),      // 6: LongReqBody
	(*LongRspBody)(nil),      // 7: LongRspBody
}
var file_longmsg_proto_depIdxs = []int32{
	4, // 0: LongReqBody.msgUpReq:type_name -> LongMsgUpReq
	2, // 1: LongReqBody.msgDownReq:type_name -> LongMsgDownReq
	0, // 2: LongReqBody.msgDelReq:type_name -> LongMsgDeleteReq
	5, // 3: LongRspBody.msgUpRsp:type_name -> LongMsgUpRsp
	3, // 4: LongRspBody.msgDownRsp:type_name -> LongMsgDownRsp
	1, // 5: LongRspBody.msgDelRsp:type_name -> LongMsgDeleteRsp
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_longmsg_proto_init() }
func file_longmsg_proto_init() {
	if File_longmsg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_longmsg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LongMsgDeleteReq); i {
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
		file_longmsg_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LongMsgDeleteRsp); i {
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
		file_longmsg_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LongMsgDownReq); i {
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
		file_longmsg_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LongMsgDownRsp); i {
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
		file_longmsg_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LongMsgUpReq); i {
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
		file_longmsg_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LongMsgUpRsp); i {
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
		file_longmsg_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LongReqBody); i {
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
		file_longmsg_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LongRspBody); i {
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
			RawDescriptor: file_longmsg_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_longmsg_proto_goTypes,
		DependencyIndexes: file_longmsg_proto_depIdxs,
		MessageInfos:      file_longmsg_proto_msgTypes,
	}.Build()
	File_longmsg_proto = out.File
	file_longmsg_proto_rawDesc = nil
	file_longmsg_proto_goTypes = nil
	file_longmsg_proto_depIdxs = nil
}
