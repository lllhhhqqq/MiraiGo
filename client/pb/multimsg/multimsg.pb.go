// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.4
// source: multimsg.proto

package multimsg

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ExternMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChannelType int32 `protobuf:"varint,1,opt,name=channelType,proto3" json:"channelType,omitempty"`
}

func (x *ExternMsg) Reset() {
	*x = ExternMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_multimsg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExternMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExternMsg) ProtoMessage() {}

func (x *ExternMsg) ProtoReflect() protoreflect.Message {
	mi := &file_multimsg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExternMsg.ProtoReflect.Descriptor instead.
func (*ExternMsg) Descriptor() ([]byte, []int) {
	return file_multimsg_proto_rawDescGZIP(), []int{0}
}

func (x *ExternMsg) GetChannelType() int32 {
	if x != nil {
		return x.ChannelType
	}
	return 0
}

type MultiMsgApplyDownReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsgResid []byte `protobuf:"bytes,1,opt,name=msgResid,proto3" json:"msgResid,omitempty"`
	MsgType  int32  `protobuf:"varint,2,opt,name=msgType,proto3" json:"msgType,omitempty"`
	SrcUin   int64  `protobuf:"varint,3,opt,name=srcUin,proto3" json:"srcUin,omitempty"`
}

func (x *MultiMsgApplyDownReq) Reset() {
	*x = MultiMsgApplyDownReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_multimsg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiMsgApplyDownReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiMsgApplyDownReq) ProtoMessage() {}

func (x *MultiMsgApplyDownReq) ProtoReflect() protoreflect.Message {
	mi := &file_multimsg_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiMsgApplyDownReq.ProtoReflect.Descriptor instead.
func (*MultiMsgApplyDownReq) Descriptor() ([]byte, []int) {
	return file_multimsg_proto_rawDescGZIP(), []int{1}
}

func (x *MultiMsgApplyDownReq) GetMsgResid() []byte {
	if x != nil {
		return x.MsgResid
	}
	return nil
}

func (x *MultiMsgApplyDownReq) GetMsgType() int32 {
	if x != nil {
		return x.MsgType
	}
	return 0
}

func (x *MultiMsgApplyDownReq) GetSrcUin() int64 {
	if x != nil {
		return x.SrcUin
	}
	return 0
}

type MultiMsgApplyDownRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result           int32      `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	ThumbDownPara    []byte     `protobuf:"bytes,2,opt,name=thumbDownPara,proto3" json:"thumbDownPara,omitempty"`
	MsgKey           []byte     `protobuf:"bytes,3,opt,name=msgKey,proto3" json:"msgKey,omitempty"`
	Uint32DownIp     []int32    `protobuf:"varint,4,rep,packed,name=uint32DownIp,proto3" json:"uint32DownIp,omitempty"`
	Uint32DownPort   []int32    `protobuf:"varint,5,rep,packed,name=uint32DownPort,proto3" json:"uint32DownPort,omitempty"`
	MsgResid         []byte     `protobuf:"bytes,6,opt,name=msgResid,proto3" json:"msgResid,omitempty"`
	MsgExternInfo    *ExternMsg `protobuf:"bytes,7,opt,name=msgExternInfo,proto3" json:"msgExternInfo,omitempty"`
	BytesDownIpV6    [][]byte   `protobuf:"bytes,8,rep,name=bytesDownIpV6,proto3" json:"bytesDownIpV6,omitempty"`
	Uint32DownV6Port []int32    `protobuf:"varint,9,rep,packed,name=uint32DownV6Port,proto3" json:"uint32DownV6Port,omitempty"`
}

func (x *MultiMsgApplyDownRsp) Reset() {
	*x = MultiMsgApplyDownRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_multimsg_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiMsgApplyDownRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiMsgApplyDownRsp) ProtoMessage() {}

func (x *MultiMsgApplyDownRsp) ProtoReflect() protoreflect.Message {
	mi := &file_multimsg_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiMsgApplyDownRsp.ProtoReflect.Descriptor instead.
func (*MultiMsgApplyDownRsp) Descriptor() ([]byte, []int) {
	return file_multimsg_proto_rawDescGZIP(), []int{2}
}

func (x *MultiMsgApplyDownRsp) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

func (x *MultiMsgApplyDownRsp) GetThumbDownPara() []byte {
	if x != nil {
		return x.ThumbDownPara
	}
	return nil
}

func (x *MultiMsgApplyDownRsp) GetMsgKey() []byte {
	if x != nil {
		return x.MsgKey
	}
	return nil
}

func (x *MultiMsgApplyDownRsp) GetUint32DownIp() []int32 {
	if x != nil {
		return x.Uint32DownIp
	}
	return nil
}

func (x *MultiMsgApplyDownRsp) GetUint32DownPort() []int32 {
	if x != nil {
		return x.Uint32DownPort
	}
	return nil
}

func (x *MultiMsgApplyDownRsp) GetMsgResid() []byte {
	if x != nil {
		return x.MsgResid
	}
	return nil
}

func (x *MultiMsgApplyDownRsp) GetMsgExternInfo() *ExternMsg {
	if x != nil {
		return x.MsgExternInfo
	}
	return nil
}

func (x *MultiMsgApplyDownRsp) GetBytesDownIpV6() [][]byte {
	if x != nil {
		return x.BytesDownIpV6
	}
	return nil
}

func (x *MultiMsgApplyDownRsp) GetUint32DownV6Port() []int32 {
	if x != nil {
		return x.Uint32DownV6Port
	}
	return nil
}

type MultiMsgApplyUpReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DstUin  int64  `protobuf:"varint,1,opt,name=dstUin,proto3" json:"dstUin,omitempty"`
	MsgSize int64  `protobuf:"varint,2,opt,name=msgSize,proto3" json:"msgSize,omitempty"`
	MsgMd5  []byte `protobuf:"bytes,3,opt,name=msgMd5,proto3" json:"msgMd5,omitempty"`
	MsgType int32  `protobuf:"varint,4,opt,name=msgType,proto3" json:"msgType,omitempty"`
	ApplyId int32  `protobuf:"varint,5,opt,name=applyId,proto3" json:"applyId,omitempty"`
}

func (x *MultiMsgApplyUpReq) Reset() {
	*x = MultiMsgApplyUpReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_multimsg_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiMsgApplyUpReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiMsgApplyUpReq) ProtoMessage() {}

func (x *MultiMsgApplyUpReq) ProtoReflect() protoreflect.Message {
	mi := &file_multimsg_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiMsgApplyUpReq.ProtoReflect.Descriptor instead.
func (*MultiMsgApplyUpReq) Descriptor() ([]byte, []int) {
	return file_multimsg_proto_rawDescGZIP(), []int{3}
}

func (x *MultiMsgApplyUpReq) GetDstUin() int64 {
	if x != nil {
		return x.DstUin
	}
	return 0
}

func (x *MultiMsgApplyUpReq) GetMsgSize() int64 {
	if x != nil {
		return x.MsgSize
	}
	return 0
}

func (x *MultiMsgApplyUpReq) GetMsgMd5() []byte {
	if x != nil {
		return x.MsgMd5
	}
	return nil
}

func (x *MultiMsgApplyUpReq) GetMsgType() int32 {
	if x != nil {
		return x.MsgType
	}
	return 0
}

func (x *MultiMsgApplyUpReq) GetApplyId() int32 {
	if x != nil {
		return x.ApplyId
	}
	return 0
}

type MultiMsgApplyUpRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result         int32      `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	MsgResid       string     `protobuf:"bytes,2,opt,name=msgResid,proto3" json:"msgResid,omitempty"`
	MsgUkey        []byte     `protobuf:"bytes,3,opt,name=msgUkey,proto3" json:"msgUkey,omitempty"`
	Uint32UpIp     []int32    `protobuf:"varint,4,rep,packed,name=uint32UpIp,proto3" json:"uint32UpIp,omitempty"`
	Uint32UpPort   []int32    `protobuf:"varint,5,rep,packed,name=uint32UpPort,proto3" json:"uint32UpPort,omitempty"`
	BlockSize      int64      `protobuf:"varint,6,opt,name=blockSize,proto3" json:"blockSize,omitempty"`
	UpOffset       int64      `protobuf:"varint,7,opt,name=upOffset,proto3" json:"upOffset,omitempty"`
	ApplyId        int32      `protobuf:"varint,8,opt,name=applyId,proto3" json:"applyId,omitempty"`
	MsgKey         []byte     `protobuf:"bytes,9,opt,name=msgKey,proto3" json:"msgKey,omitempty"`
	MsgSig         []byte     `protobuf:"bytes,10,opt,name=msgSig,proto3" json:"msgSig,omitempty"`
	MsgExternInfo  *ExternMsg `protobuf:"bytes,11,opt,name=msgExternInfo,proto3" json:"msgExternInfo,omitempty"`
	BytesUpIpV6    [][]byte   `protobuf:"bytes,12,rep,name=bytesUpIpV6,proto3" json:"bytesUpIpV6,omitempty"`
	Uint32UpV6Port []int32    `protobuf:"varint,13,rep,packed,name=uint32UpV6Port,proto3" json:"uint32UpV6Port,omitempty"`
}

func (x *MultiMsgApplyUpRsp) Reset() {
	*x = MultiMsgApplyUpRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_multimsg_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiMsgApplyUpRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiMsgApplyUpRsp) ProtoMessage() {}

func (x *MultiMsgApplyUpRsp) ProtoReflect() protoreflect.Message {
	mi := &file_multimsg_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiMsgApplyUpRsp.ProtoReflect.Descriptor instead.
func (*MultiMsgApplyUpRsp) Descriptor() ([]byte, []int) {
	return file_multimsg_proto_rawDescGZIP(), []int{4}
}

func (x *MultiMsgApplyUpRsp) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

func (x *MultiMsgApplyUpRsp) GetMsgResid() string {
	if x != nil {
		return x.MsgResid
	}
	return ""
}

func (x *MultiMsgApplyUpRsp) GetMsgUkey() []byte {
	if x != nil {
		return x.MsgUkey
	}
	return nil
}

func (x *MultiMsgApplyUpRsp) GetUint32UpIp() []int32 {
	if x != nil {
		return x.Uint32UpIp
	}
	return nil
}

func (x *MultiMsgApplyUpRsp) GetUint32UpPort() []int32 {
	if x != nil {
		return x.Uint32UpPort
	}
	return nil
}

func (x *MultiMsgApplyUpRsp) GetBlockSize() int64 {
	if x != nil {
		return x.BlockSize
	}
	return 0
}

func (x *MultiMsgApplyUpRsp) GetUpOffset() int64 {
	if x != nil {
		return x.UpOffset
	}
	return 0
}

func (x *MultiMsgApplyUpRsp) GetApplyId() int32 {
	if x != nil {
		return x.ApplyId
	}
	return 0
}

func (x *MultiMsgApplyUpRsp) GetMsgKey() []byte {
	if x != nil {
		return x.MsgKey
	}
	return nil
}

func (x *MultiMsgApplyUpRsp) GetMsgSig() []byte {
	if x != nil {
		return x.MsgSig
	}
	return nil
}

func (x *MultiMsgApplyUpRsp) GetMsgExternInfo() *ExternMsg {
	if x != nil {
		return x.MsgExternInfo
	}
	return nil
}

func (x *MultiMsgApplyUpRsp) GetBytesUpIpV6() [][]byte {
	if x != nil {
		return x.BytesUpIpV6
	}
	return nil
}

func (x *MultiMsgApplyUpRsp) GetUint32UpV6Port() []int32 {
	if x != nil {
		return x.Uint32UpV6Port
	}
	return nil
}

type ReqBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subcmd               int32                   `protobuf:"varint,1,opt,name=subcmd,proto3" json:"subcmd,omitempty"`
	TermType             int32                   `protobuf:"varint,2,opt,name=termType,proto3" json:"termType,omitempty"`
	PlatformType         int32                   `protobuf:"varint,3,opt,name=platformType,proto3" json:"platformType,omitempty"`
	NetType              int32                   `protobuf:"varint,4,opt,name=netType,proto3" json:"netType,omitempty"`
	BuildVer             string                  `protobuf:"bytes,5,opt,name=buildVer,proto3" json:"buildVer,omitempty"`
	MultimsgApplyupReq   []*MultiMsgApplyUpReq   `protobuf:"bytes,6,rep,name=multimsgApplyupReq,proto3" json:"multimsgApplyupReq,omitempty"`
	MultimsgApplydownReq []*MultiMsgApplyDownReq `protobuf:"bytes,7,rep,name=multimsgApplydownReq,proto3" json:"multimsgApplydownReq,omitempty"`
	BuType               int32                   `protobuf:"varint,8,opt,name=buType,proto3" json:"buType,omitempty"`
	ReqChannelType       int32                   `protobuf:"varint,9,opt,name=reqChannelType,proto3" json:"reqChannelType,omitempty"`
}

func (x *ReqBody) Reset() {
	*x = ReqBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_multimsg_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqBody) ProtoMessage() {}

func (x *ReqBody) ProtoReflect() protoreflect.Message {
	mi := &file_multimsg_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqBody.ProtoReflect.Descriptor instead.
func (*ReqBody) Descriptor() ([]byte, []int) {
	return file_multimsg_proto_rawDescGZIP(), []int{5}
}

func (x *ReqBody) GetSubcmd() int32 {
	if x != nil {
		return x.Subcmd
	}
	return 0
}

func (x *ReqBody) GetTermType() int32 {
	if x != nil {
		return x.TermType
	}
	return 0
}

func (x *ReqBody) GetPlatformType() int32 {
	if x != nil {
		return x.PlatformType
	}
	return 0
}

func (x *ReqBody) GetNetType() int32 {
	if x != nil {
		return x.NetType
	}
	return 0
}

func (x *ReqBody) GetBuildVer() string {
	if x != nil {
		return x.BuildVer
	}
	return ""
}

func (x *ReqBody) GetMultimsgApplyupReq() []*MultiMsgApplyUpReq {
	if x != nil {
		return x.MultimsgApplyupReq
	}
	return nil
}

func (x *ReqBody) GetMultimsgApplydownReq() []*MultiMsgApplyDownReq {
	if x != nil {
		return x.MultimsgApplydownReq
	}
	return nil
}

func (x *ReqBody) GetBuType() int32 {
	if x != nil {
		return x.BuType
	}
	return 0
}

func (x *ReqBody) GetReqChannelType() int32 {
	if x != nil {
		return x.ReqChannelType
	}
	return 0
}

type RspBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subcmd               int32                   `protobuf:"varint,1,opt,name=subcmd,proto3" json:"subcmd,omitempty"`
	MultimsgApplyupRsp   []*MultiMsgApplyUpRsp   `protobuf:"bytes,2,rep,name=multimsgApplyupRsp,proto3" json:"multimsgApplyupRsp,omitempty"`
	MultimsgApplydownRsp []*MultiMsgApplyDownRsp `protobuf:"bytes,3,rep,name=multimsgApplydownRsp,proto3" json:"multimsgApplydownRsp,omitempty"`
}

func (x *RspBody) Reset() {
	*x = RspBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_multimsg_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RspBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RspBody) ProtoMessage() {}

func (x *RspBody) ProtoReflect() protoreflect.Message {
	mi := &file_multimsg_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RspBody.ProtoReflect.Descriptor instead.
func (*RspBody) Descriptor() ([]byte, []int) {
	return file_multimsg_proto_rawDescGZIP(), []int{6}
}

func (x *RspBody) GetSubcmd() int32 {
	if x != nil {
		return x.Subcmd
	}
	return 0
}

func (x *RspBody) GetMultimsgApplyupRsp() []*MultiMsgApplyUpRsp {
	if x != nil {
		return x.MultimsgApplyupRsp
	}
	return nil
}

func (x *RspBody) GetMultimsgApplydownRsp() []*MultiMsgApplyDownRsp {
	if x != nil {
		return x.MultimsgApplydownRsp
	}
	return nil
}

var File_multimsg_proto protoreflect.FileDescriptor

var file_multimsg_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x2d, 0x0a, 0x09, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x4d, 0x73, 0x67, 0x12, 0x20, 0x0a,
	0x0b, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x22,
	0x64, 0x0a, 0x14, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x4d, 0x73, 0x67, 0x41, 0x70, 0x70, 0x6c, 0x79,
	0x44, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x73, 0x67, 0x52, 0x65,
	0x73, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6d, 0x73, 0x67, 0x52, 0x65,
	0x73, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x72, 0x63, 0x55, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73,
	0x72, 0x63, 0x55, 0x69, 0x6e, 0x22, 0xd8, 0x02, 0x0a, 0x14, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x4d,
	0x73, 0x67, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x44, 0x6f, 0x77, 0x6e, 0x52, 0x73, 0x70, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x44,
	0x6f, 0x77, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x74,
	0x68, 0x75, 0x6d, 0x62, 0x44, 0x6f, 0x77, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x12, 0x16, 0x0a, 0x06,
	0x6d, 0x73, 0x67, 0x4b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x6d, 0x73,
	0x67, 0x4b, 0x65, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x44, 0x6f,
	0x77, 0x6e, 0x49, 0x70, 0x18, 0x04, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0c, 0x75, 0x69, 0x6e, 0x74,
	0x33, 0x32, 0x44, 0x6f, 0x77, 0x6e, 0x49, 0x70, 0x12, 0x26, 0x0a, 0x0e, 0x75, 0x69, 0x6e, 0x74,
	0x33, 0x32, 0x44, 0x6f, 0x77, 0x6e, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x05,
	0x52, 0x0e, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x44, 0x6f, 0x77, 0x6e, 0x50, 0x6f, 0x72, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x73, 0x67, 0x52, 0x65, 0x73, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x08, 0x6d, 0x73, 0x67, 0x52, 0x65, 0x73, 0x69, 0x64, 0x12, 0x30, 0x0a, 0x0d,
	0x6d, 0x73, 0x67, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x4d, 0x73, 0x67, 0x52,
	0x0d, 0x6d, 0x73, 0x67, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x24,
	0x0a, 0x0d, 0x62, 0x79, 0x74, 0x65, 0x73, 0x44, 0x6f, 0x77, 0x6e, 0x49, 0x70, 0x56, 0x36, 0x18,
	0x08, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0d, 0x62, 0x79, 0x74, 0x65, 0x73, 0x44, 0x6f, 0x77, 0x6e,
	0x49, 0x70, 0x56, 0x36, 0x12, 0x2a, 0x0a, 0x10, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x44, 0x6f,
	0x77, 0x6e, 0x56, 0x36, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x09, 0x20, 0x03, 0x28, 0x05, 0x52, 0x10,
	0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x44, 0x6f, 0x77, 0x6e, 0x56, 0x36, 0x50, 0x6f, 0x72, 0x74,
	0x22, 0x92, 0x01, 0x0a, 0x12, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x4d, 0x73, 0x67, 0x41, 0x70, 0x70,
	0x6c, 0x79, 0x55, 0x70, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x73, 0x74, 0x55, 0x69,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x64, 0x73, 0x74, 0x55, 0x69, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x73, 0x67, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x6d, 0x73, 0x67, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x73, 0x67,
	0x4d, 0x64, 0x35, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x6d, 0x73, 0x67, 0x4d, 0x64,
	0x35, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x6d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x70, 0x70, 0x6c, 0x79, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x61, 0x70,
	0x70, 0x6c, 0x79, 0x49, 0x64, 0x22, 0xa6, 0x03, 0x0a, 0x12, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x4d,
	0x73, 0x67, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x55, 0x70, 0x52, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x73, 0x67, 0x52, 0x65, 0x73, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x73, 0x67, 0x52, 0x65, 0x73, 0x69, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x73, 0x67, 0x55, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x07, 0x6d, 0x73, 0x67, 0x55, 0x6b, 0x65, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x69,
	0x6e, 0x74, 0x33, 0x32, 0x55, 0x70, 0x49, 0x70, 0x18, 0x04, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0a,
	0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x55, 0x70, 0x49, 0x70, 0x12, 0x22, 0x0a, 0x0c, 0x75, 0x69,
	0x6e, 0x74, 0x33, 0x32, 0x55, 0x70, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x05,
	0x52, 0x0c, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x55, 0x70, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x70, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x75, 0x70, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x70, 0x70, 0x6c,
	0x79, 0x49, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x61, 0x70, 0x70, 0x6c, 0x79,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x73, 0x67, 0x4b, 0x65, 0x79, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x06, 0x6d, 0x73, 0x67, 0x4b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x73,
	0x67, 0x53, 0x69, 0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x6d, 0x73, 0x67, 0x53,
	0x69, 0x67, 0x12, 0x30, 0x0a, 0x0d, 0x6d, 0x73, 0x67, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x45, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x4d, 0x73, 0x67, 0x52, 0x0d, 0x6d, 0x73, 0x67, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x62, 0x79, 0x74, 0x65, 0x73, 0x55, 0x70, 0x49,
	0x70, 0x56, 0x36, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0b, 0x62, 0x79, 0x74, 0x65, 0x73,
	0x55, 0x70, 0x49, 0x70, 0x56, 0x36, 0x12, 0x26, 0x0a, 0x0e, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32,
	0x55, 0x70, 0x56, 0x36, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0e,
	0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x55, 0x70, 0x56, 0x36, 0x50, 0x6f, 0x72, 0x74, 0x22, 0xe7,
	0x02, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x75,
	0x62, 0x63, 0x6d, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x75, 0x62, 0x63,
	0x6d, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x72, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x74, 0x65, 0x72, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22,
	0x0a, 0x0c, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x56, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x56, 0x65, 0x72, 0x12, 0x43, 0x0a, 0x12, 0x6d, 0x75, 0x6c, 0x74,
	0x69, 0x6d, 0x73, 0x67, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x75, 0x70, 0x52, 0x65, 0x71, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x4d, 0x73, 0x67, 0x41,
	0x70, 0x70, 0x6c, 0x79, 0x55, 0x70, 0x52, 0x65, 0x71, 0x52, 0x12, 0x6d, 0x75, 0x6c, 0x74, 0x69,
	0x6d, 0x73, 0x67, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x75, 0x70, 0x52, 0x65, 0x71, 0x12, 0x49, 0x0a,
	0x14, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x6d, 0x73, 0x67, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x64, 0x6f,
	0x77, 0x6e, 0x52, 0x65, 0x71, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x4d, 0x75,
	0x6c, 0x74, 0x69, 0x4d, 0x73, 0x67, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x44, 0x6f, 0x77, 0x6e, 0x52,
	0x65, 0x71, 0x52, 0x14, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x6d, 0x73, 0x67, 0x41, 0x70, 0x70, 0x6c,
	0x79, 0x64, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x62, 0x75, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x26, 0x0a, 0x0e, 0x72, 0x65, 0x71, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x72, 0x65, 0x71, 0x43, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x22, 0xb1, 0x01, 0x0a, 0x07, 0x52, 0x73, 0x70,
	0x42, 0x6f, 0x64, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x75, 0x62, 0x63, 0x6d, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x75, 0x62, 0x63, 0x6d, 0x64, 0x12, 0x43, 0x0a, 0x12,
	0x6d, 0x75, 0x6c, 0x74, 0x69, 0x6d, 0x73, 0x67, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x75, 0x70, 0x52,
	0x73, 0x70, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69,
	0x4d, 0x73, 0x67, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x55, 0x70, 0x52, 0x73, 0x70, 0x52, 0x12, 0x6d,
	0x75, 0x6c, 0x74, 0x69, 0x6d, 0x73, 0x67, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x75, 0x70, 0x52, 0x73,
	0x70, 0x12, 0x49, 0x0a, 0x14, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x6d, 0x73, 0x67, 0x41, 0x70, 0x70,
	0x6c, 0x79, 0x64, 0x6f, 0x77, 0x6e, 0x52, 0x73, 0x70, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x4d, 0x73, 0x67, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x44,
	0x6f, 0x77, 0x6e, 0x52, 0x73, 0x70, 0x52, 0x14, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x6d, 0x73, 0x67,
	0x41, 0x70, 0x70, 0x6c, 0x79, 0x64, 0x6f, 0x77, 0x6e, 0x52, 0x73, 0x70, 0x42, 0x0c, 0x5a, 0x0a,
	0x2e, 0x3b, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x6d, 0x73, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_multimsg_proto_rawDescOnce sync.Once
	file_multimsg_proto_rawDescData = file_multimsg_proto_rawDesc
)

func file_multimsg_proto_rawDescGZIP() []byte {
	file_multimsg_proto_rawDescOnce.Do(func() {
		file_multimsg_proto_rawDescData = protoimpl.X.CompressGZIP(file_multimsg_proto_rawDescData)
	})
	return file_multimsg_proto_rawDescData
}

var file_multimsg_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_multimsg_proto_goTypes = []interface{}{
	(*ExternMsg)(nil),            // 0: ExternMsg
	(*MultiMsgApplyDownReq)(nil), // 1: MultiMsgApplyDownReq
	(*MultiMsgApplyDownRsp)(nil), // 2: MultiMsgApplyDownRsp
	(*MultiMsgApplyUpReq)(nil),   // 3: MultiMsgApplyUpReq
	(*MultiMsgApplyUpRsp)(nil),   // 4: MultiMsgApplyUpRsp
	(*ReqBody)(nil),              // 5: ReqBody
	(*RspBody)(nil),              // 6: RspBody
}
var file_multimsg_proto_depIdxs = []int32{
	0, // 0: MultiMsgApplyDownRsp.msgExternInfo:type_name -> ExternMsg
	0, // 1: MultiMsgApplyUpRsp.msgExternInfo:type_name -> ExternMsg
	3, // 2: ReqBody.multimsgApplyupReq:type_name -> MultiMsgApplyUpReq
	1, // 3: ReqBody.multimsgApplydownReq:type_name -> MultiMsgApplyDownReq
	4, // 4: RspBody.multimsgApplyupRsp:type_name -> MultiMsgApplyUpRsp
	2, // 5: RspBody.multimsgApplydownRsp:type_name -> MultiMsgApplyDownRsp
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_multimsg_proto_init() }
func file_multimsg_proto_init() {
	if File_multimsg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_multimsg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExternMsg); i {
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
		file_multimsg_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiMsgApplyDownReq); i {
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
		file_multimsg_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiMsgApplyDownRsp); i {
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
		file_multimsg_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiMsgApplyUpReq); i {
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
		file_multimsg_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiMsgApplyUpRsp); i {
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
		file_multimsg_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqBody); i {
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
		file_multimsg_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RspBody); i {
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
			RawDescriptor: file_multimsg_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_multimsg_proto_goTypes,
		DependencyIndexes: file_multimsg_proto_depIdxs,
		MessageInfos:      file_multimsg_proto_msgTypes,
	}.Build()
	File_multimsg_proto = out.File
	file_multimsg_proto_rawDesc = nil
	file_multimsg_proto_goTypes = nil
	file_multimsg_proto_depIdxs = nil
}
