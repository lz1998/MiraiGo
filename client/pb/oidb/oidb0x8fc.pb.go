// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: oidb0x8fc.proto

package oidb

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

type D8FCReqBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupCode      *int64            `protobuf:"varint,1,opt,name=groupCode" json:"groupCode,omitempty"`
	ShowFlag       *int32            `protobuf:"varint,2,opt,name=showFlag" json:"showFlag,omitempty"`
	MemLevelInfo   []*D8FCMemberInfo `protobuf:"bytes,3,rep,name=memLevelInfo" json:"memLevelInfo,omitempty"`
	LevelName      []*D8FCLevelName  `protobuf:"bytes,4,rep,name=levelName" json:"levelName,omitempty"`
	UpdateTime     *int32            `protobuf:"varint,5,opt,name=updateTime" json:"updateTime,omitempty"`
	OfficeMode     *int32            `protobuf:"varint,6,opt,name=officeMode" json:"officeMode,omitempty"`
	GroupOpenAppid *int32            `protobuf:"varint,7,opt,name=groupOpenAppid" json:"groupOpenAppid,omitempty"`
	MsgClientInfo  *D8FCClientInfo   `protobuf:"bytes,8,opt,name=msgClientInfo" json:"msgClientInfo,omitempty"`
	AuthKey        []byte            `protobuf:"bytes,9,opt,name=authKey" json:"authKey,omitempty"`
}

func (x *D8FCReqBody) Reset() {
	*x = D8FCReqBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oidb0x8fc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *D8FCReqBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*D8FCReqBody) ProtoMessage() {}

func (x *D8FCReqBody) ProtoReflect() protoreflect.Message {
	mi := &file_oidb0x8fc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use D8FCReqBody.ProtoReflect.Descriptor instead.
func (*D8FCReqBody) Descriptor() ([]byte, []int) {
	return file_oidb0x8fc_proto_rawDescGZIP(), []int{0}
}

func (x *D8FCReqBody) GetGroupCode() int64 {
	if x != nil && x.GroupCode != nil {
		return *x.GroupCode
	}
	return 0
}

func (x *D8FCReqBody) GetShowFlag() int32 {
	if x != nil && x.ShowFlag != nil {
		return *x.ShowFlag
	}
	return 0
}

func (x *D8FCReqBody) GetMemLevelInfo() []*D8FCMemberInfo {
	if x != nil {
		return x.MemLevelInfo
	}
	return nil
}

func (x *D8FCReqBody) GetLevelName() []*D8FCLevelName {
	if x != nil {
		return x.LevelName
	}
	return nil
}

func (x *D8FCReqBody) GetUpdateTime() int32 {
	if x != nil && x.UpdateTime != nil {
		return *x.UpdateTime
	}
	return 0
}

func (x *D8FCReqBody) GetOfficeMode() int32 {
	if x != nil && x.OfficeMode != nil {
		return *x.OfficeMode
	}
	return 0
}

func (x *D8FCReqBody) GetGroupOpenAppid() int32 {
	if x != nil && x.GroupOpenAppid != nil {
		return *x.GroupOpenAppid
	}
	return 0
}

func (x *D8FCReqBody) GetMsgClientInfo() *D8FCClientInfo {
	if x != nil {
		return x.MsgClientInfo
	}
	return nil
}

func (x *D8FCReqBody) GetAuthKey() []byte {
	if x != nil {
		return x.AuthKey
	}
	return nil
}

type D8FCMemberInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uin                    *int64              `protobuf:"varint,1,opt,name=uin" json:"uin,omitempty"`
	Point                  *int32              `protobuf:"varint,2,opt,name=point" json:"point,omitempty"`
	ActiveDay              *int32              `protobuf:"varint,3,opt,name=activeDay" json:"activeDay,omitempty"`
	Level                  *int32              `protobuf:"varint,4,opt,name=level" json:"level,omitempty"`
	SpecialTitle           []byte              `protobuf:"bytes,5,opt,name=specialTitle" json:"specialTitle,omitempty"`
	SpecialTitleExpireTime *int32              `protobuf:"varint,6,opt,name=specialTitleExpireTime" json:"specialTitleExpireTime,omitempty"`
	UinName                []byte              `protobuf:"bytes,7,opt,name=uinName" json:"uinName,omitempty"`
	MemberCardName         []byte              `protobuf:"bytes,8,opt,name=memberCardName" json:"memberCardName,omitempty"`
	Phone                  []byte              `protobuf:"bytes,9,opt,name=phone" json:"phone,omitempty"`
	Email                  []byte              `protobuf:"bytes,10,opt,name=email" json:"email,omitempty"`
	Remark                 []byte              `protobuf:"bytes,11,opt,name=remark" json:"remark,omitempty"`
	Gender                 *int32              `protobuf:"varint,12,opt,name=gender" json:"gender,omitempty"`
	Job                    []byte              `protobuf:"bytes,13,opt,name=job" json:"job,omitempty"`
	TribeLevel             *int32              `protobuf:"varint,14,opt,name=tribeLevel" json:"tribeLevel,omitempty"`
	TribePoint             *int32              `protobuf:"varint,15,opt,name=tribePoint" json:"tribePoint,omitempty"`
	RichCardName           []*D8FCCardNameElem `protobuf:"bytes,16,rep,name=richCardName" json:"richCardName,omitempty"`
	CommRichCardName       []byte              `protobuf:"bytes,17,opt,name=commRichCardName" json:"commRichCardName,omitempty"`
}

func (x *D8FCMemberInfo) Reset() {
	*x = D8FCMemberInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oidb0x8fc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *D8FCMemberInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*D8FCMemberInfo) ProtoMessage() {}

func (x *D8FCMemberInfo) ProtoReflect() protoreflect.Message {
	mi := &file_oidb0x8fc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use D8FCMemberInfo.ProtoReflect.Descriptor instead.
func (*D8FCMemberInfo) Descriptor() ([]byte, []int) {
	return file_oidb0x8fc_proto_rawDescGZIP(), []int{1}
}

func (x *D8FCMemberInfo) GetUin() int64 {
	if x != nil && x.Uin != nil {
		return *x.Uin
	}
	return 0
}

func (x *D8FCMemberInfo) GetPoint() int32 {
	if x != nil && x.Point != nil {
		return *x.Point
	}
	return 0
}

func (x *D8FCMemberInfo) GetActiveDay() int32 {
	if x != nil && x.ActiveDay != nil {
		return *x.ActiveDay
	}
	return 0
}

func (x *D8FCMemberInfo) GetLevel() int32 {
	if x != nil && x.Level != nil {
		return *x.Level
	}
	return 0
}

func (x *D8FCMemberInfo) GetSpecialTitle() []byte {
	if x != nil {
		return x.SpecialTitle
	}
	return nil
}

func (x *D8FCMemberInfo) GetSpecialTitleExpireTime() int32 {
	if x != nil && x.SpecialTitleExpireTime != nil {
		return *x.SpecialTitleExpireTime
	}
	return 0
}

func (x *D8FCMemberInfo) GetUinName() []byte {
	if x != nil {
		return x.UinName
	}
	return nil
}

func (x *D8FCMemberInfo) GetMemberCardName() []byte {
	if x != nil {
		return x.MemberCardName
	}
	return nil
}

func (x *D8FCMemberInfo) GetPhone() []byte {
	if x != nil {
		return x.Phone
	}
	return nil
}

func (x *D8FCMemberInfo) GetEmail() []byte {
	if x != nil {
		return x.Email
	}
	return nil
}

func (x *D8FCMemberInfo) GetRemark() []byte {
	if x != nil {
		return x.Remark
	}
	return nil
}

func (x *D8FCMemberInfo) GetGender() int32 {
	if x != nil && x.Gender != nil {
		return *x.Gender
	}
	return 0
}

func (x *D8FCMemberInfo) GetJob() []byte {
	if x != nil {
		return x.Job
	}
	return nil
}

func (x *D8FCMemberInfo) GetTribeLevel() int32 {
	if x != nil && x.TribeLevel != nil {
		return *x.TribeLevel
	}
	return 0
}

func (x *D8FCMemberInfo) GetTribePoint() int32 {
	if x != nil && x.TribePoint != nil {
		return *x.TribePoint
	}
	return 0
}

func (x *D8FCMemberInfo) GetRichCardName() []*D8FCCardNameElem {
	if x != nil {
		return x.RichCardName
	}
	return nil
}

func (x *D8FCMemberInfo) GetCommRichCardName() []byte {
	if x != nil {
		return x.CommRichCardName
	}
	return nil
}

type D8FCCardNameElem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EnumCardType *int32 `protobuf:"varint,1,opt,name=enumCardType" json:"enumCardType,omitempty"`
	Value        []byte `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (x *D8FCCardNameElem) Reset() {
	*x = D8FCCardNameElem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oidb0x8fc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *D8FCCardNameElem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*D8FCCardNameElem) ProtoMessage() {}

func (x *D8FCCardNameElem) ProtoReflect() protoreflect.Message {
	mi := &file_oidb0x8fc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use D8FCCardNameElem.ProtoReflect.Descriptor instead.
func (*D8FCCardNameElem) Descriptor() ([]byte, []int) {
	return file_oidb0x8fc_proto_rawDescGZIP(), []int{2}
}

func (x *D8FCCardNameElem) GetEnumCardType() int32 {
	if x != nil && x.EnumCardType != nil {
		return *x.EnumCardType
	}
	return 0
}

func (x *D8FCCardNameElem) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type D8FCLevelName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Level *int32  `protobuf:"varint,1,opt,name=level" json:"level,omitempty"`
	Name  *string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (x *D8FCLevelName) Reset() {
	*x = D8FCLevelName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oidb0x8fc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *D8FCLevelName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*D8FCLevelName) ProtoMessage() {}

func (x *D8FCLevelName) ProtoReflect() protoreflect.Message {
	mi := &file_oidb0x8fc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use D8FCLevelName.ProtoReflect.Descriptor instead.
func (*D8FCLevelName) Descriptor() ([]byte, []int) {
	return file_oidb0x8fc_proto_rawDescGZIP(), []int{3}
}

func (x *D8FCLevelName) GetLevel() int32 {
	if x != nil && x.Level != nil {
		return *x.Level
	}
	return 0
}

func (x *D8FCLevelName) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

type D8FCClientInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Implat       *int32  `protobuf:"varint,1,opt,name=implat" json:"implat,omitempty"`
	IngClientver *string `protobuf:"bytes,2,opt,name=ingClientver" json:"ingClientver,omitempty"`
}

func (x *D8FCClientInfo) Reset() {
	*x = D8FCClientInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oidb0x8fc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *D8FCClientInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*D8FCClientInfo) ProtoMessage() {}

func (x *D8FCClientInfo) ProtoReflect() protoreflect.Message {
	mi := &file_oidb0x8fc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use D8FCClientInfo.ProtoReflect.Descriptor instead.
func (*D8FCClientInfo) Descriptor() ([]byte, []int) {
	return file_oidb0x8fc_proto_rawDescGZIP(), []int{4}
}

func (x *D8FCClientInfo) GetImplat() int32 {
	if x != nil && x.Implat != nil {
		return *x.Implat
	}
	return 0
}

func (x *D8FCClientInfo) GetIngClientver() string {
	if x != nil && x.IngClientver != nil {
		return *x.IngClientver
	}
	return ""
}

type D8FCCommCardNameBuf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RichCardName []*D8FCRichCardNameElem `protobuf:"bytes,1,rep,name=richCardName" json:"richCardName,omitempty"`
}

func (x *D8FCCommCardNameBuf) Reset() {
	*x = D8FCCommCardNameBuf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oidb0x8fc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *D8FCCommCardNameBuf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*D8FCCommCardNameBuf) ProtoMessage() {}

func (x *D8FCCommCardNameBuf) ProtoReflect() protoreflect.Message {
	mi := &file_oidb0x8fc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use D8FCCommCardNameBuf.ProtoReflect.Descriptor instead.
func (*D8FCCommCardNameBuf) Descriptor() ([]byte, []int) {
	return file_oidb0x8fc_proto_rawDescGZIP(), []int{5}
}

func (x *D8FCCommCardNameBuf) GetRichCardName() []*D8FCRichCardNameElem {
	if x != nil {
		return x.RichCardName
	}
	return nil
}

type D8FCRichCardNameElem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ctrl []byte `protobuf:"bytes,1,opt,name=ctrl" json:"ctrl,omitempty"`
	Text []byte `protobuf:"bytes,2,opt,name=text" json:"text,omitempty"`
}

func (x *D8FCRichCardNameElem) Reset() {
	*x = D8FCRichCardNameElem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oidb0x8fc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *D8FCRichCardNameElem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*D8FCRichCardNameElem) ProtoMessage() {}

func (x *D8FCRichCardNameElem) ProtoReflect() protoreflect.Message {
	mi := &file_oidb0x8fc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use D8FCRichCardNameElem.ProtoReflect.Descriptor instead.
func (*D8FCRichCardNameElem) Descriptor() ([]byte, []int) {
	return file_oidb0x8fc_proto_rawDescGZIP(), []int{6}
}

func (x *D8FCRichCardNameElem) GetCtrl() []byte {
	if x != nil {
		return x.Ctrl
	}
	return nil
}

func (x *D8FCRichCardNameElem) GetText() []byte {
	if x != nil {
		return x.Text
	}
	return nil
}

var File_oidb0x8fc_proto protoreflect.FileDescriptor

var file_oidb0x8fc_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6f, 0x69, 0x64, 0x62, 0x30, 0x78, 0x38, 0x66, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xe3, 0x02, 0x0a, 0x0b, 0x44, 0x38, 0x46, 0x43, 0x52, 0x65, 0x71, 0x42, 0x6f, 0x64,
	0x79, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x73, 0x68, 0x6f, 0x77, 0x46, 0x6c, 0x61, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x77, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x33, 0x0a, 0x0c, 0x6d,
	0x65, 0x6d, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x44, 0x38, 0x46, 0x43, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x0c, 0x6d, 0x65, 0x6d, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x2c, 0x0a, 0x09, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x44, 0x38, 0x46, 0x43, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4e,
	0x61, 0x6d, 0x65, 0x52, 0x09, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x26,
	0x0a, 0x0e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4f, 0x70, 0x65, 0x6e, 0x41, 0x70, 0x70, 0x69, 0x64,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4f, 0x70, 0x65,
	0x6e, 0x41, 0x70, 0x70, 0x69, 0x64, 0x12, 0x35, 0x0a, 0x0d, 0x6d, 0x73, 0x67, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x44, 0x38, 0x46, 0x43, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0d,
	0x6d, 0x73, 0x67, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07,
	0x61, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x22, 0x9b, 0x04, 0x0a, 0x0e, 0x44, 0x38, 0x46, 0x43,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x44, 0x61, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x44, 0x61, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x70, 0x65, 0x63, 0x69, 0x61,
	0x6c, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x73, 0x70,
	0x65, 0x63, 0x69, 0x61, 0x6c, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x36, 0x0a, 0x16, 0x73, 0x70,
	0x65, 0x63, 0x69, 0x61, 0x6c, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x16, 0x73, 0x70, 0x65, 0x63,
	0x69, 0x61, 0x6c, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x07, 0x75, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x43, 0x61, 0x72, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x43, 0x61, 0x72, 0x64,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x12, 0x10, 0x0a, 0x03, 0x6a, 0x6f, 0x62, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6a,
	0x6f, 0x62, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x72, 0x69, 0x62, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x72, 0x69, 0x62, 0x65, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x72, 0x69, 0x62, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x72, 0x69, 0x62, 0x65, 0x50, 0x6f, 0x69,
	0x6e, 0x74, 0x12, 0x35, 0x0a, 0x0c, 0x72, 0x69, 0x63, 0x68, 0x43, 0x61, 0x72, 0x64, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x10, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x44, 0x38, 0x46, 0x43, 0x43,
	0x61, 0x72, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x45, 0x6c, 0x65, 0x6d, 0x52, 0x0c, 0x72, 0x69, 0x63,
	0x68, 0x43, 0x61, 0x72, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x63, 0x6f, 0x6d,
	0x6d, 0x52, 0x69, 0x63, 0x68, 0x43, 0x61, 0x72, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x11, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x10, 0x63, 0x6f, 0x6d, 0x6d, 0x52, 0x69, 0x63, 0x68, 0x43, 0x61, 0x72,
	0x64, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x4c, 0x0a, 0x10, 0x44, 0x38, 0x46, 0x43, 0x43, 0x61, 0x72,
	0x64, 0x4e, 0x61, 0x6d, 0x65, 0x45, 0x6c, 0x65, 0x6d, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x6e, 0x75,
	0x6d, 0x43, 0x61, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0c, 0x65, 0x6e, 0x75, 0x6d, 0x43, 0x61, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x39, 0x0a, 0x0d, 0x44, 0x38, 0x46, 0x43, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x4c,
	0x0a, 0x0e, 0x44, 0x38, 0x46, 0x43, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x16, 0x0a, 0x06, 0x69, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x69, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x6e, 0x67, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x69, 0x6e, 0x67, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x76, 0x65, 0x72, 0x22, 0x50, 0x0a, 0x13,
	0x44, 0x38, 0x46, 0x43, 0x43, 0x6f, 0x6d, 0x6d, 0x43, 0x61, 0x72, 0x64, 0x4e, 0x61, 0x6d, 0x65,
	0x42, 0x75, 0x66, 0x12, 0x39, 0x0a, 0x0c, 0x72, 0x69, 0x63, 0x68, 0x43, 0x61, 0x72, 0x64, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x44, 0x38, 0x46, 0x43,
	0x52, 0x69, 0x63, 0x68, 0x43, 0x61, 0x72, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x45, 0x6c, 0x65, 0x6d,
	0x52, 0x0c, 0x72, 0x69, 0x63, 0x68, 0x43, 0x61, 0x72, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x3e,
	0x0a, 0x14, 0x44, 0x38, 0x46, 0x43, 0x52, 0x69, 0x63, 0x68, 0x43, 0x61, 0x72, 0x64, 0x4e, 0x61,
	0x6d, 0x65, 0x45, 0x6c, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x74, 0x72, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x63, 0x74, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65,
	0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x42, 0x09,
	0x5a, 0x07, 0x2e, 0x2f, 0x3b, 0x6f, 0x69, 0x64, 0x62,
}

var (
	file_oidb0x8fc_proto_rawDescOnce sync.Once
	file_oidb0x8fc_proto_rawDescData = file_oidb0x8fc_proto_rawDesc
)

func file_oidb0x8fc_proto_rawDescGZIP() []byte {
	file_oidb0x8fc_proto_rawDescOnce.Do(func() {
		file_oidb0x8fc_proto_rawDescData = protoimpl.X.CompressGZIP(file_oidb0x8fc_proto_rawDescData)
	})
	return file_oidb0x8fc_proto_rawDescData
}

var file_oidb0x8fc_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_oidb0x8fc_proto_goTypes = []interface{}{
	(*D8FCReqBody)(nil),          // 0: D8FCReqBody
	(*D8FCMemberInfo)(nil),       // 1: D8FCMemberInfo
	(*D8FCCardNameElem)(nil),     // 2: D8FCCardNameElem
	(*D8FCLevelName)(nil),        // 3: D8FCLevelName
	(*D8FCClientInfo)(nil),       // 4: D8FCClientInfo
	(*D8FCCommCardNameBuf)(nil),  // 5: D8FCCommCardNameBuf
	(*D8FCRichCardNameElem)(nil), // 6: D8FCRichCardNameElem
}
var file_oidb0x8fc_proto_depIdxs = []int32{
	1, // 0: D8FCReqBody.memLevelInfo:type_name -> D8FCMemberInfo
	3, // 1: D8FCReqBody.levelName:type_name -> D8FCLevelName
	4, // 2: D8FCReqBody.msgClientInfo:type_name -> D8FCClientInfo
	2, // 3: D8FCMemberInfo.richCardName:type_name -> D8FCCardNameElem
	6, // 4: D8FCCommCardNameBuf.richCardName:type_name -> D8FCRichCardNameElem
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_oidb0x8fc_proto_init() }
func file_oidb0x8fc_proto_init() {
	if File_oidb0x8fc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_oidb0x8fc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*D8FCReqBody); i {
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
		file_oidb0x8fc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*D8FCMemberInfo); i {
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
		file_oidb0x8fc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*D8FCCardNameElem); i {
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
		file_oidb0x8fc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*D8FCLevelName); i {
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
		file_oidb0x8fc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*D8FCClientInfo); i {
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
		file_oidb0x8fc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*D8FCCommCardNameBuf); i {
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
		file_oidb0x8fc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*D8FCRichCardNameElem); i {
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
			RawDescriptor: file_oidb0x8fc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_oidb0x8fc_proto_goTypes,
		DependencyIndexes: file_oidb0x8fc_proto_depIdxs,
		MessageInfos:      file_oidb0x8fc_proto_msgTypes,
	}.Build()
	File_oidb0x8fc_proto = out.File
	file_oidb0x8fc_proto_rawDesc = nil
	file_oidb0x8fc_proto_goTypes = nil
	file_oidb0x8fc_proto_depIdxs = nil
}
