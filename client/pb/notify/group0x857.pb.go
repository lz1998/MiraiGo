// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: group0x857.proto

package notify

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type NotifyMsgBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OptMsgGrayTips    *AIOGrayTipsInfo       `protobuf:"bytes,5,opt,name=optMsgGrayTips,proto3" json:"optMsgGrayTips,omitempty"`
	OptMsgRedTips     *RedGrayTipsInfo       `protobuf:"bytes,9,opt,name=optMsgRedTips,proto3" json:"optMsgRedTips,omitempty"`
	OptMsgRecall      *MessageRecallReminder `protobuf:"bytes,11,opt,name=optMsgRecall,proto3" json:"optMsgRecall,omitempty"`
	OptGeneralGrayTip *GeneralGrayTipInfo    `protobuf:"bytes,26,opt,name=optGeneralGrayTip,proto3" json:"optGeneralGrayTip,omitempty"`
	QqGroupDigestMsg  *QQGroupDigestMsg      `protobuf:"bytes,33,opt,name=qqGroupDigestMsg,proto3" json:"qqGroupDigestMsg,omitempty"`
	ServiceType       int32                  `protobuf:"varint,13,opt,name=serviceType,proto3" json:"serviceType,omitempty"`
}

func (x *NotifyMsgBody) Reset() {
	*x = NotifyMsgBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group0x857_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyMsgBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyMsgBody) ProtoMessage() {}

func (x *NotifyMsgBody) ProtoReflect() protoreflect.Message {
	mi := &file_group0x857_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyMsgBody.ProtoReflect.Descriptor instead.
func (*NotifyMsgBody) Descriptor() ([]byte, []int) {
	return file_group0x857_proto_rawDescGZIP(), []int{0}
}

func (x *NotifyMsgBody) GetOptMsgGrayTips() *AIOGrayTipsInfo {
	if x != nil {
		return x.OptMsgGrayTips
	}
	return nil
}

func (x *NotifyMsgBody) GetOptMsgRedTips() *RedGrayTipsInfo {
	if x != nil {
		return x.OptMsgRedTips
	}
	return nil
}

func (x *NotifyMsgBody) GetOptMsgRecall() *MessageRecallReminder {
	if x != nil {
		return x.OptMsgRecall
	}
	return nil
}

func (x *NotifyMsgBody) GetOptGeneralGrayTip() *GeneralGrayTipInfo {
	if x != nil {
		return x.OptGeneralGrayTip
	}
	return nil
}

func (x *NotifyMsgBody) GetQqGroupDigestMsg() *QQGroupDigestMsg {
	if x != nil {
		return x.QqGroupDigestMsg
	}
	return nil
}

func (x *NotifyMsgBody) GetServiceType() int32 {
	if x != nil {
		return x.ServiceType
	}
	return 0
}

type AIOGrayTipsInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShowLatest     uint32 `protobuf:"varint,1,opt,name=showLatest,proto3" json:"showLatest,omitempty"`
	Content        []byte `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Remind         uint32 `protobuf:"varint,3,opt,name=remind,proto3" json:"remind,omitempty"`
	Brief          []byte `protobuf:"bytes,4,opt,name=brief,proto3" json:"brief,omitempty"`
	ReceiverUin    uint64 `protobuf:"varint,5,opt,name=receiverUin,proto3" json:"receiverUin,omitempty"`
	ReliaoAdminOpt uint32 `protobuf:"varint,6,opt,name=reliaoAdminOpt,proto3" json:"reliaoAdminOpt,omitempty"`
}

func (x *AIOGrayTipsInfo) Reset() {
	*x = AIOGrayTipsInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group0x857_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AIOGrayTipsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AIOGrayTipsInfo) ProtoMessage() {}

func (x *AIOGrayTipsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_group0x857_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AIOGrayTipsInfo.ProtoReflect.Descriptor instead.
func (*AIOGrayTipsInfo) Descriptor() ([]byte, []int) {
	return file_group0x857_proto_rawDescGZIP(), []int{1}
}

func (x *AIOGrayTipsInfo) GetShowLatest() uint32 {
	if x != nil {
		return x.ShowLatest
	}
	return 0
}

func (x *AIOGrayTipsInfo) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *AIOGrayTipsInfo) GetRemind() uint32 {
	if x != nil {
		return x.Remind
	}
	return 0
}

func (x *AIOGrayTipsInfo) GetBrief() []byte {
	if x != nil {
		return x.Brief
	}
	return nil
}

func (x *AIOGrayTipsInfo) GetReceiverUin() uint64 {
	if x != nil {
		return x.ReceiverUin
	}
	return 0
}

func (x *AIOGrayTipsInfo) GetReliaoAdminOpt() uint32 {
	if x != nil {
		return x.ReliaoAdminOpt
	}
	return 0
}

type GeneralGrayTipInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BusiType      uint64        `protobuf:"varint,1,opt,name=busiType,proto3" json:"busiType,omitempty"`
	BusiId        uint64        `protobuf:"varint,2,opt,name=busiId,proto3" json:"busiId,omitempty"`
	CtrlFlag      uint32        `protobuf:"varint,3,opt,name=ctrlFlag,proto3" json:"ctrlFlag,omitempty"`
	C2CType       uint32        `protobuf:"varint,4,opt,name=c2cType,proto3" json:"c2cType,omitempty"`
	ServiceType   uint32        `protobuf:"varint,5,opt,name=serviceType,proto3" json:"serviceType,omitempty"`
	TemplId       uint64        `protobuf:"varint,6,opt,name=templId,proto3" json:"templId,omitempty"`
	MsgTemplParam []*TemplParam `protobuf:"bytes,7,rep,name=msgTemplParam,proto3" json:"msgTemplParam,omitempty"`
	Content       string        `protobuf:"bytes,8,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *GeneralGrayTipInfo) Reset() {
	*x = GeneralGrayTipInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group0x857_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeneralGrayTipInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeneralGrayTipInfo) ProtoMessage() {}

func (x *GeneralGrayTipInfo) ProtoReflect() protoreflect.Message {
	mi := &file_group0x857_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeneralGrayTipInfo.ProtoReflect.Descriptor instead.
func (*GeneralGrayTipInfo) Descriptor() ([]byte, []int) {
	return file_group0x857_proto_rawDescGZIP(), []int{2}
}

func (x *GeneralGrayTipInfo) GetBusiType() uint64 {
	if x != nil {
		return x.BusiType
	}
	return 0
}

func (x *GeneralGrayTipInfo) GetBusiId() uint64 {
	if x != nil {
		return x.BusiId
	}
	return 0
}

func (x *GeneralGrayTipInfo) GetCtrlFlag() uint32 {
	if x != nil {
		return x.CtrlFlag
	}
	return 0
}

func (x *GeneralGrayTipInfo) GetC2CType() uint32 {
	if x != nil {
		return x.C2CType
	}
	return 0
}

func (x *GeneralGrayTipInfo) GetServiceType() uint32 {
	if x != nil {
		return x.ServiceType
	}
	return 0
}

func (x *GeneralGrayTipInfo) GetTemplId() uint64 {
	if x != nil {
		return x.TemplId
	}
	return 0
}

func (x *GeneralGrayTipInfo) GetMsgTemplParam() []*TemplParam {
	if x != nil {
		return x.MsgTemplParam
	}
	return nil
}

func (x *GeneralGrayTipInfo) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type TemplParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *TemplParam) Reset() {
	*x = TemplParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group0x857_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplParam) ProtoMessage() {}

func (x *TemplParam) ProtoReflect() protoreflect.Message {
	mi := &file_group0x857_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplParam.ProtoReflect.Descriptor instead.
func (*TemplParam) Descriptor() ([]byte, []int) {
	return file_group0x857_proto_rawDescGZIP(), []int{3}
}

func (x *TemplParam) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TemplParam) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type MessageRecallReminder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uin             int64                  `protobuf:"varint,1,opt,name=uin,proto3" json:"uin,omitempty"`
	Nickname        []byte                 `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
	RecalledMsgList []*RecalledMessageMeta `protobuf:"bytes,3,rep,name=recalledMsgList,proto3" json:"recalledMsgList,omitempty"`
	ReminderContent []byte                 `protobuf:"bytes,4,opt,name=reminderContent,proto3" json:"reminderContent,omitempty"`
	Userdef         []byte                 `protobuf:"bytes,5,opt,name=userdef,proto3" json:"userdef,omitempty"`
	GroupType       int32                  `protobuf:"varint,6,opt,name=groupType,proto3" json:"groupType,omitempty"`
	OpType          int32                  `protobuf:"varint,7,opt,name=opType,proto3" json:"opType,omitempty"`
}

func (x *MessageRecallReminder) Reset() {
	*x = MessageRecallReminder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group0x857_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageRecallReminder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageRecallReminder) ProtoMessage() {}

func (x *MessageRecallReminder) ProtoReflect() protoreflect.Message {
	mi := &file_group0x857_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageRecallReminder.ProtoReflect.Descriptor instead.
func (*MessageRecallReminder) Descriptor() ([]byte, []int) {
	return file_group0x857_proto_rawDescGZIP(), []int{4}
}

func (x *MessageRecallReminder) GetUin() int64 {
	if x != nil {
		return x.Uin
	}
	return 0
}

func (x *MessageRecallReminder) GetNickname() []byte {
	if x != nil {
		return x.Nickname
	}
	return nil
}

func (x *MessageRecallReminder) GetRecalledMsgList() []*RecalledMessageMeta {
	if x != nil {
		return x.RecalledMsgList
	}
	return nil
}

func (x *MessageRecallReminder) GetReminderContent() []byte {
	if x != nil {
		return x.ReminderContent
	}
	return nil
}

func (x *MessageRecallReminder) GetUserdef() []byte {
	if x != nil {
		return x.Userdef
	}
	return nil
}

func (x *MessageRecallReminder) GetGroupType() int32 {
	if x != nil {
		return x.GroupType
	}
	return 0
}

func (x *MessageRecallReminder) GetOpType() int32 {
	if x != nil {
		return x.OpType
	}
	return 0
}

type RecalledMessageMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seq       int32 `protobuf:"varint,1,opt,name=seq,proto3" json:"seq,omitempty"`
	Time      int32 `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`
	MsgRandom int32 `protobuf:"varint,3,opt,name=msgRandom,proto3" json:"msgRandom,omitempty"`
	MsgType   int32 `protobuf:"varint,4,opt,name=msgType,proto3" json:"msgType,omitempty"`
	MsgFlag   int32 `protobuf:"varint,5,opt,name=msgFlag,proto3" json:"msgFlag,omitempty"`
	AuthorUin int64 `protobuf:"varint,6,opt,name=authorUin,proto3" json:"authorUin,omitempty"`
}

func (x *RecalledMessageMeta) Reset() {
	*x = RecalledMessageMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group0x857_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecalledMessageMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecalledMessageMeta) ProtoMessage() {}

func (x *RecalledMessageMeta) ProtoReflect() protoreflect.Message {
	mi := &file_group0x857_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecalledMessageMeta.ProtoReflect.Descriptor instead.
func (*RecalledMessageMeta) Descriptor() ([]byte, []int) {
	return file_group0x857_proto_rawDescGZIP(), []int{5}
}

func (x *RecalledMessageMeta) GetSeq() int32 {
	if x != nil {
		return x.Seq
	}
	return 0
}

func (x *RecalledMessageMeta) GetTime() int32 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *RecalledMessageMeta) GetMsgRandom() int32 {
	if x != nil {
		return x.MsgRandom
	}
	return 0
}

func (x *RecalledMessageMeta) GetMsgType() int32 {
	if x != nil {
		return x.MsgType
	}
	return 0
}

func (x *RecalledMessageMeta) GetMsgFlag() int32 {
	if x != nil {
		return x.MsgFlag
	}
	return 0
}

func (x *RecalledMessageMeta) GetAuthorUin() int64 {
	if x != nil {
		return x.AuthorUin
	}
	return 0
}

type RedGrayTipsInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShowLatest          uint32 `protobuf:"varint,1,opt,name=showLatest,proto3" json:"showLatest,omitempty"`
	SenderUin           uint64 `protobuf:"varint,2,opt,name=senderUin,proto3" json:"senderUin,omitempty"`
	ReceiverUin         uint64 `protobuf:"varint,3,opt,name=receiverUin,proto3" json:"receiverUin,omitempty"`
	SenderRichContent   string `protobuf:"bytes,4,opt,name=senderRichContent,proto3" json:"senderRichContent,omitempty"`
	ReceiverRichContent string `protobuf:"bytes,5,opt,name=receiverRichContent,proto3" json:"receiverRichContent,omitempty"`
	AuthKey             []byte `protobuf:"bytes,6,opt,name=authKey,proto3" json:"authKey,omitempty"`
	MsgType             int32  `protobuf:"zigzag32,7,opt,name=msgType,proto3" json:"msgType,omitempty"`
	LuckyFlag           uint32 `protobuf:"varint,8,opt,name=luckyFlag,proto3" json:"luckyFlag,omitempty"`
	HideFlag            uint32 `protobuf:"varint,9,opt,name=hideFlag,proto3" json:"hideFlag,omitempty"`
	LuckyUin            uint64 `protobuf:"varint,12,opt,name=luckyUin,proto3" json:"luckyUin,omitempty"`
}

func (x *RedGrayTipsInfo) Reset() {
	*x = RedGrayTipsInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group0x857_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RedGrayTipsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedGrayTipsInfo) ProtoMessage() {}

func (x *RedGrayTipsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_group0x857_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedGrayTipsInfo.ProtoReflect.Descriptor instead.
func (*RedGrayTipsInfo) Descriptor() ([]byte, []int) {
	return file_group0x857_proto_rawDescGZIP(), []int{6}
}

func (x *RedGrayTipsInfo) GetShowLatest() uint32 {
	if x != nil {
		return x.ShowLatest
	}
	return 0
}

func (x *RedGrayTipsInfo) GetSenderUin() uint64 {
	if x != nil {
		return x.SenderUin
	}
	return 0
}

func (x *RedGrayTipsInfo) GetReceiverUin() uint64 {
	if x != nil {
		return x.ReceiverUin
	}
	return 0
}

func (x *RedGrayTipsInfo) GetSenderRichContent() string {
	if x != nil {
		return x.SenderRichContent
	}
	return ""
}

func (x *RedGrayTipsInfo) GetReceiverRichContent() string {
	if x != nil {
		return x.ReceiverRichContent
	}
	return ""
}

func (x *RedGrayTipsInfo) GetAuthKey() []byte {
	if x != nil {
		return x.AuthKey
	}
	return nil
}

func (x *RedGrayTipsInfo) GetMsgType() int32 {
	if x != nil {
		return x.MsgType
	}
	return 0
}

func (x *RedGrayTipsInfo) GetLuckyFlag() uint32 {
	if x != nil {
		return x.LuckyFlag
	}
	return 0
}

func (x *RedGrayTipsInfo) GetHideFlag() uint32 {
	if x != nil {
		return x.HideFlag
	}
	return 0
}

func (x *RedGrayTipsInfo) GetLuckyUin() uint64 {
	if x != nil {
		return x.LuckyUin
	}
	return 0
}

type QQGroupDigestMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupCode     uint64 `protobuf:"varint,1,opt,name=groupCode,proto3" json:"groupCode,omitempty"`
	Seq           uint32 `protobuf:"varint,2,opt,name=seq,proto3" json:"seq,omitempty"`
	Random        uint32 `protobuf:"varint,3,opt,name=random,proto3" json:"random,omitempty"`
	OpType        int32  `protobuf:"varint,4,opt,name=opType,proto3" json:"opType,omitempty"`
	Sender        uint64 `protobuf:"varint,5,opt,name=sender,proto3" json:"sender,omitempty"`
	DigestOper    uint64 `protobuf:"varint,6,opt,name=digestOper,proto3" json:"digestOper,omitempty"`
	OpTime        uint32 `protobuf:"varint,7,opt,name=opTime,proto3" json:"opTime,omitempty"`
	LastestMsgSeq uint32 `protobuf:"varint,8,opt,name=lastestMsgSeq,proto3" json:"lastestMsgSeq,omitempty"`
	OperNick      []byte `protobuf:"bytes,9,opt,name=operNick,proto3" json:"operNick,omitempty"`
	SenderNick    []byte `protobuf:"bytes,10,opt,name=senderNick,proto3" json:"senderNick,omitempty"`
	ExtInfo       int32  `protobuf:"varint,11,opt,name=extInfo,proto3" json:"extInfo,omitempty"`
}

func (x *QQGroupDigestMsg) Reset() {
	*x = QQGroupDigestMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group0x857_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QQGroupDigestMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QQGroupDigestMsg) ProtoMessage() {}

func (x *QQGroupDigestMsg) ProtoReflect() protoreflect.Message {
	mi := &file_group0x857_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QQGroupDigestMsg.ProtoReflect.Descriptor instead.
func (*QQGroupDigestMsg) Descriptor() ([]byte, []int) {
	return file_group0x857_proto_rawDescGZIP(), []int{7}
}

func (x *QQGroupDigestMsg) GetGroupCode() uint64 {
	if x != nil {
		return x.GroupCode
	}
	return 0
}

func (x *QQGroupDigestMsg) GetSeq() uint32 {
	if x != nil {
		return x.Seq
	}
	return 0
}

func (x *QQGroupDigestMsg) GetRandom() uint32 {
	if x != nil {
		return x.Random
	}
	return 0
}

func (x *QQGroupDigestMsg) GetOpType() int32 {
	if x != nil {
		return x.OpType
	}
	return 0
}

func (x *QQGroupDigestMsg) GetSender() uint64 {
	if x != nil {
		return x.Sender
	}
	return 0
}

func (x *QQGroupDigestMsg) GetDigestOper() uint64 {
	if x != nil {
		return x.DigestOper
	}
	return 0
}

func (x *QQGroupDigestMsg) GetOpTime() uint32 {
	if x != nil {
		return x.OpTime
	}
	return 0
}

func (x *QQGroupDigestMsg) GetLastestMsgSeq() uint32 {
	if x != nil {
		return x.LastestMsgSeq
	}
	return 0
}

func (x *QQGroupDigestMsg) GetOperNick() []byte {
	if x != nil {
		return x.OperNick
	}
	return nil
}

func (x *QQGroupDigestMsg) GetSenderNick() []byte {
	if x != nil {
		return x.SenderNick
	}
	return nil
}

func (x *QQGroupDigestMsg) GetExtInfo() int32 {
	if x != nil {
		return x.ExtInfo
	}
	return 0
}

var File_group0x857_proto protoreflect.FileDescriptor

var file_group0x857_proto_rawDesc = []byte{
	0x0a, 0x10, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x30, 0x78, 0x38, 0x35, 0x37, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xe1, 0x02, 0x0a, 0x0d, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4d, 0x73, 0x67,
	0x42, 0x6f, 0x64, 0x79, 0x12, 0x38, 0x0a, 0x0e, 0x6f, 0x70, 0x74, 0x4d, 0x73, 0x67, 0x47, 0x72,
	0x61, 0x79, 0x54, 0x69, 0x70, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x41,
	0x49, 0x4f, 0x47, 0x72, 0x61, 0x79, 0x54, 0x69, 0x70, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0e,
	0x6f, 0x70, 0x74, 0x4d, 0x73, 0x67, 0x47, 0x72, 0x61, 0x79, 0x54, 0x69, 0x70, 0x73, 0x12, 0x36,
	0x0a, 0x0d, 0x6f, 0x70, 0x74, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x64, 0x54, 0x69, 0x70, 0x73, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x52, 0x65, 0x64, 0x47, 0x72, 0x61, 0x79, 0x54,
	0x69, 0x70, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0d, 0x6f, 0x70, 0x74, 0x4d, 0x73, 0x67, 0x52,
	0x65, 0x64, 0x54, 0x69, 0x70, 0x73, 0x12, 0x3a, 0x0a, 0x0c, 0x6f, 0x70, 0x74, 0x4d, 0x73, 0x67,
	0x52, 0x65, 0x63, 0x61, 0x6c, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x63, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x6d, 0x69,
	0x6e, 0x64, 0x65, 0x72, 0x52, 0x0c, 0x6f, 0x70, 0x74, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x63, 0x61,
	0x6c, 0x6c, 0x12, 0x41, 0x0a, 0x11, 0x6f, 0x70, 0x74, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c,
	0x47, 0x72, 0x61, 0x79, 0x54, 0x69, 0x70, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x47, 0x72, 0x61, 0x79, 0x54, 0x69, 0x70, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x11, 0x6f, 0x70, 0x74, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x47, 0x72,
	0x61, 0x79, 0x54, 0x69, 0x70, 0x12, 0x3d, 0x0a, 0x10, 0x71, 0x71, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x4d, 0x73, 0x67, 0x18, 0x21, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x51, 0x51, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x4d,
	0x73, 0x67, 0x52, 0x10, 0x71, 0x71, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x69, 0x67, 0x65, 0x73,
	0x74, 0x4d, 0x73, 0x67, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0xc3, 0x01, 0x0a, 0x0f, 0x41, 0x49, 0x4f, 0x47, 0x72,
	0x61, 0x79, 0x54, 0x69, 0x70, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x68,
	0x6f, 0x77, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a,
	0x73, 0x68, 0x6f, 0x77, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x62, 0x72, 0x69, 0x65, 0x66, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x62, 0x72, 0x69,
	0x65, 0x66, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x55, 0x69,
	0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x72, 0x55, 0x69, 0x6e, 0x12, 0x26, 0x0a, 0x0e, 0x72, 0x65, 0x6c, 0x69, 0x61, 0x6f, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x4f, 0x70, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x72, 0x65,
	0x6c, 0x69, 0x61, 0x6f, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x4f, 0x70, 0x74, 0x22, 0x87, 0x02, 0x0a,
	0x12, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x47, 0x72, 0x61, 0x79, 0x54, 0x69, 0x70, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x75, 0x73, 0x69, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x62, 0x75, 0x73, 0x69, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x62, 0x75, 0x73, 0x69, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x06, 0x62, 0x75, 0x73, 0x69, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x74, 0x72, 0x6c, 0x46,
	0x6c, 0x61, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x74, 0x72, 0x6c, 0x46,
	0x6c, 0x61, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x32, 0x63, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x63, 0x32, 0x63, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x07, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x0d, 0x6d, 0x73, 0x67,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0b, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x0d, 0x6d,
	0x73, 0x67, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x36, 0x0a, 0x0a, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xff,
	0x01, 0x0a, 0x15, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x63, 0x61, 0x6c, 0x6c,
	0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69,
	0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6e, 0x69,
	0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3e, 0x0a, 0x0f, 0x72, 0x65, 0x63, 0x61, 0x6c, 0x6c,
	0x65, 0x64, 0x4d, 0x73, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x52, 0x65, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x0f, 0x72, 0x65, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x4d,
	0x73, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x72, 0x65, 0x6d, 0x69, 0x6e, 0x64,
	0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0f, 0x72, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x64, 0x65, 0x66, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x64, 0x65, 0x66, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x54, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x70, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x70, 0x54, 0x79, 0x70, 0x65,
	0x22, 0xab, 0x01, 0x0a, 0x13, 0x52, 0x65, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x71, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x73, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x6d, 0x73, 0x67, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x6d, 0x73, 0x67, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d,
	0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x73, 0x67, 0x46, 0x6c, 0x61,
	0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d, 0x73, 0x67, 0x46, 0x6c, 0x61, 0x67,
	0x12, 0x1c, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x55, 0x69, 0x6e, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x55, 0x69, 0x6e, 0x22, 0xdb,
	0x02, 0x0a, 0x0f, 0x52, 0x65, 0x64, 0x47, 0x72, 0x61, 0x79, 0x54, 0x69, 0x70, 0x73, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x68, 0x6f, 0x77, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x73, 0x68, 0x6f, 0x77, 0x4c, 0x61, 0x74, 0x65,
	0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x55, 0x69, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x55, 0x69, 0x6e,
	0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x55, 0x69, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x55,
	0x69, 0x6e, 0x12, 0x2c, 0x0a, 0x11, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x69, 0x63, 0x68,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x73,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x69, 0x63, 0x68, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x12, 0x30, 0x0a, 0x13, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x52, 0x69, 0x63, 0x68,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x72,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x52, 0x69, 0x63, 0x68, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x07, 0x61, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x11, 0x52, 0x07, 0x6d,
	0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x75, 0x63, 0x6b, 0x79, 0x46,
	0x6c, 0x61, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x6c, 0x75, 0x63, 0x6b, 0x79,
	0x46, 0x6c, 0x61, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x69, 0x64, 0x65, 0x46, 0x6c, 0x61, 0x67,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x68, 0x69, 0x64, 0x65, 0x46, 0x6c, 0x61, 0x67,
	0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x75, 0x63, 0x6b, 0x79, 0x55, 0x69, 0x6e, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x08, 0x6c, 0x75, 0x63, 0x6b, 0x79, 0x55, 0x69, 0x6e, 0x22, 0xbe, 0x02, 0x0a,
	0x10, 0x51, 0x51, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x4d, 0x73,
	0x67, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x73, 0x65, 0x71, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x73, 0x65,
	0x71, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x06, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x70, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x70, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x69, 0x67,
	0x65, 0x73, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x64,
	0x69, 0x67, 0x65, 0x73, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x70, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6f, 0x70, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x24, 0x0a, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x65, 0x73, 0x74, 0x4d, 0x73, 0x67, 0x53,
	0x65, 0x71, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x65, 0x73,
	0x74, 0x4d, 0x73, 0x67, 0x53, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x4e,
	0x69, 0x63, 0x6b, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x4e,
	0x69, 0x63, 0x6b, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x4e, 0x69, 0x63,
	0x6b, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x4e,
	0x69, 0x63, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x78, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x65, 0x78, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x0b, 0x5a,
	0x09, 0x2e, 0x2f, 0x3b, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_group0x857_proto_rawDescOnce sync.Once
	file_group0x857_proto_rawDescData = file_group0x857_proto_rawDesc
)

func file_group0x857_proto_rawDescGZIP() []byte {
	file_group0x857_proto_rawDescOnce.Do(func() {
		file_group0x857_proto_rawDescData = protoimpl.X.CompressGZIP(file_group0x857_proto_rawDescData)
	})
	return file_group0x857_proto_rawDescData
}

var file_group0x857_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_group0x857_proto_goTypes = []interface{}{
	(*NotifyMsgBody)(nil),         // 0: NotifyMsgBody
	(*AIOGrayTipsInfo)(nil),       // 1: AIOGrayTipsInfo
	(*GeneralGrayTipInfo)(nil),    // 2: GeneralGrayTipInfo
	(*TemplParam)(nil),            // 3: TemplParam
	(*MessageRecallReminder)(nil), // 4: MessageRecallReminder
	(*RecalledMessageMeta)(nil),   // 5: RecalledMessageMeta
	(*RedGrayTipsInfo)(nil),       // 6: RedGrayTipsInfo
	(*QQGroupDigestMsg)(nil),      // 7: QQGroupDigestMsg
}
var file_group0x857_proto_depIdxs = []int32{
	1, // 0: NotifyMsgBody.optMsgGrayTips:type_name -> AIOGrayTipsInfo
	6, // 1: NotifyMsgBody.optMsgRedTips:type_name -> RedGrayTipsInfo
	4, // 2: NotifyMsgBody.optMsgRecall:type_name -> MessageRecallReminder
	2, // 3: NotifyMsgBody.optGeneralGrayTip:type_name -> GeneralGrayTipInfo
	7, // 4: NotifyMsgBody.qqGroupDigestMsg:type_name -> QQGroupDigestMsg
	3, // 5: GeneralGrayTipInfo.msgTemplParam:type_name -> TemplParam
	5, // 6: MessageRecallReminder.recalledMsgList:type_name -> RecalledMessageMeta
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_group0x857_proto_init() }
func file_group0x857_proto_init() {
	if File_group0x857_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_group0x857_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifyMsgBody); i {
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
		file_group0x857_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AIOGrayTipsInfo); i {
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
		file_group0x857_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeneralGrayTipInfo); i {
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
		file_group0x857_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplParam); i {
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
		file_group0x857_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageRecallReminder); i {
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
		file_group0x857_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecalledMessageMeta); i {
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
		file_group0x857_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RedGrayTipsInfo); i {
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
		file_group0x857_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QQGroupDigestMsg); i {
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
			RawDescriptor: file_group0x857_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_group0x857_proto_goTypes,
		DependencyIndexes: file_group0x857_proto_depIdxs,
		MessageInfos:      file_group0x857_proto_msgTypes,
	}.Build()
	File_group0x857_proto = out.File
	file_group0x857_proto_rawDesc = nil
	file_group0x857_proto_goTypes = nil
	file_group0x857_proto_depIdxs = nil
}
