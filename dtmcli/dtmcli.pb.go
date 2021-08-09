// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: dtmcli/dtmcli.proto

package dtmcli

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

type DtmTransInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gid       string `protobuf:"bytes,1,opt,name=Gid,proto3" json:"Gid,omitempty"`
	TransType string `protobuf:"bytes,2,opt,name=TransType,proto3" json:"TransType,omitempty"`
	BranchID  string `protobuf:"bytes,3,opt,name=BranchID,proto3" json:"BranchID,omitempty"`
	Dtm       string `protobuf:"bytes,4,opt,name=Dtm,proto3" json:"Dtm,omitempty"`
}

func (x *DtmTransInfo) Reset() {
	*x = DtmTransInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dtmcli_dtmcli_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DtmTransInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DtmTransInfo) ProtoMessage() {}

func (x *DtmTransInfo) ProtoReflect() protoreflect.Message {
	mi := &file_dtmcli_dtmcli_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DtmTransInfo.ProtoReflect.Descriptor instead.
func (*DtmTransInfo) Descriptor() ([]byte, []int) {
	return file_dtmcli_dtmcli_proto_rawDescGZIP(), []int{0}
}

func (x *DtmTransInfo) GetGid() string {
	if x != nil {
		return x.Gid
	}
	return ""
}

func (x *DtmTransInfo) GetTransType() string {
	if x != nil {
		return x.TransType
	}
	return ""
}

func (x *DtmTransInfo) GetBranchID() string {
	if x != nil {
		return x.BranchID
	}
	return ""
}

func (x *DtmTransInfo) GetDtm() string {
	if x != nil {
		return x.Dtm
	}
	return ""
}

// The request message containing the user's name.
type DtmRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gid           string            `protobuf:"bytes,1,opt,name=Gid,proto3" json:"Gid,omitempty"`
	TransType     string            `protobuf:"bytes,2,opt,name=TransType,proto3" json:"TransType,omitempty"`
	QueryPrepared string            `protobuf:"bytes,3,opt,name=QueryPrepared,proto3" json:"QueryPrepared,omitempty"`
	Method        string            `protobuf:"bytes,4,opt,name=Method,proto3" json:"Method,omitempty"`
	Extra         map[string]string `protobuf:"bytes,5,rep,name=Extra,proto3" json:"Extra,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	AppData       []byte            `protobuf:"bytes,6,opt,name=AppData,proto3" json:"AppData,omitempty"`
}

func (x *DtmRequest) Reset() {
	*x = DtmRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dtmcli_dtmcli_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DtmRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DtmRequest) ProtoMessage() {}

func (x *DtmRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dtmcli_dtmcli_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DtmRequest.ProtoReflect.Descriptor instead.
func (*DtmRequest) Descriptor() ([]byte, []int) {
	return file_dtmcli_dtmcli_proto_rawDescGZIP(), []int{1}
}

func (x *DtmRequest) GetGid() string {
	if x != nil {
		return x.Gid
	}
	return ""
}

func (x *DtmRequest) GetTransType() string {
	if x != nil {
		return x.TransType
	}
	return ""
}

func (x *DtmRequest) GetQueryPrepared() string {
	if x != nil {
		return x.QueryPrepared
	}
	return ""
}

func (x *DtmRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *DtmRequest) GetExtra() map[string]string {
	if x != nil {
		return x.Extra
	}
	return nil
}

func (x *DtmRequest) GetAppData() []byte {
	if x != nil {
		return x.AppData
	}
	return nil
}

// The response message containing the greetings
type DtmReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DtmResult  string `protobuf:"bytes,1,opt,name=DtmResult,proto3" json:"DtmResult,omitempty"`
	DtmMessage string `protobuf:"bytes,2,opt,name=DtmMessage,proto3" json:"DtmMessage,omitempty"`
}

func (x *DtmReply) Reset() {
	*x = DtmReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dtmcli_dtmcli_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DtmReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DtmReply) ProtoMessage() {}

func (x *DtmReply) ProtoReflect() protoreflect.Message {
	mi := &file_dtmcli_dtmcli_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DtmReply.ProtoReflect.Descriptor instead.
func (*DtmReply) Descriptor() ([]byte, []int) {
	return file_dtmcli_dtmcli_proto_rawDescGZIP(), []int{2}
}

func (x *DtmReply) GetDtmResult() string {
	if x != nil {
		return x.DtmResult
	}
	return ""
}

func (x *DtmReply) GetDtmMessage() string {
	if x != nil {
		return x.DtmMessage
	}
	return ""
}

// The request message containing the user's name.
type BusiRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info    *DtmTransInfo     `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Extra   map[string]string `protobuf:"bytes,2,rep,name=Extra,proto3" json:"Extra,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	AppData []byte            `protobuf:"bytes,3,opt,name=AppData,proto3" json:"AppData,omitempty"`
}

func (x *BusiRequest) Reset() {
	*x = BusiRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dtmcli_dtmcli_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BusiRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BusiRequest) ProtoMessage() {}

func (x *BusiRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dtmcli_dtmcli_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BusiRequest.ProtoReflect.Descriptor instead.
func (*BusiRequest) Descriptor() ([]byte, []int) {
	return file_dtmcli_dtmcli_proto_rawDescGZIP(), []int{3}
}

func (x *BusiRequest) GetInfo() *DtmTransInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *BusiRequest) GetExtra() map[string]string {
	if x != nil {
		return x.Extra
	}
	return nil
}

func (x *BusiRequest) GetAppData() []byte {
	if x != nil {
		return x.AppData
	}
	return nil
}

// The response message containing the greetings
type BusiReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DtmResult  string `protobuf:"bytes,1,opt,name=DtmResult,proto3" json:"DtmResult,omitempty"`
	DtmMessage string `protobuf:"bytes,2,opt,name=DtmMessage,proto3" json:"DtmMessage,omitempty"`
}

func (x *BusiReply) Reset() {
	*x = BusiReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dtmcli_dtmcli_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BusiReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BusiReply) ProtoMessage() {}

func (x *BusiReply) ProtoReflect() protoreflect.Message {
	mi := &file_dtmcli_dtmcli_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BusiReply.ProtoReflect.Descriptor instead.
func (*BusiReply) Descriptor() ([]byte, []int) {
	return file_dtmcli_dtmcli_proto_rawDescGZIP(), []int{4}
}

func (x *BusiReply) GetDtmResult() string {
	if x != nil {
		return x.DtmResult
	}
	return ""
}

func (x *BusiReply) GetDtmMessage() string {
	if x != nil {
		return x.DtmMessage
	}
	return ""
}

var File_dtmcli_dtmcli_proto protoreflect.FileDescriptor

var file_dtmcli_dtmcli_proto_rawDesc = []byte{
	0x0a, 0x13, 0x64, 0x74, 0x6d, 0x63, 0x6c, 0x69, 0x2f, 0x64, 0x74, 0x6d, 0x63, 0x6c, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x64, 0x74, 0x6d, 0x63, 0x6c, 0x69, 0x22, 0x6c, 0x0a,
	0x0c, 0x44, 0x74, 0x6d, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x10, 0x0a,
	0x03, 0x47, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x47, 0x69, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x03, 0x44, 0x74, 0x6d,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x44, 0x74, 0x6d, 0x22, 0x83, 0x02, 0x0a, 0x0a,
	0x44, 0x74, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x47, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x47, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x33, 0x0a, 0x05, 0x45, 0x78, 0x74, 0x72,
	0x61, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x64, 0x74, 0x6d, 0x63, 0x6c, 0x69,
	0x2e, 0x44, 0x74, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x45, 0x78, 0x74, 0x72,
	0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x45, 0x78, 0x74, 0x72, 0x61, 0x12, 0x18, 0x0a,
	0x07, 0x41, 0x70, 0x70, 0x44, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07,
	0x41, 0x70, 0x70, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x38, 0x0a, 0x0a, 0x45, 0x78, 0x74, 0x72, 0x61,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x48, 0x0a, 0x08, 0x44, 0x74, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1c, 0x0a,
	0x09, 0x44, 0x74, 0x6d, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x44, 0x74, 0x6d, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x44,
	0x74, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x44, 0x74, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xc1, 0x01, 0x0a, 0x0b,
	0x42, 0x75, 0x73, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x04, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x64, 0x74, 0x6d, 0x63,
	0x6c, 0x69, 0x2e, 0x44, 0x74, 0x6d, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x34, 0x0a, 0x05, 0x45, 0x78, 0x74, 0x72, 0x61, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x64, 0x74, 0x6d, 0x63, 0x6c, 0x69, 0x2e, 0x42, 0x75,
	0x73, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x45, 0x78, 0x74, 0x72, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x41,
	0x70, 0x70, 0x44, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x41, 0x70,
	0x70, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x38, 0x0a, 0x0a, 0x45, 0x78, 0x74, 0x72, 0x61, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x49, 0x0a, 0x09, 0x42, 0x75, 0x73, 0x69, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1c, 0x0a, 0x09,
	0x44, 0x74, 0x6d, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x44, 0x74, 0x6d, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x44, 0x74,
	0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x44, 0x74, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x35, 0x0a, 0x03, 0x44, 0x74,
	0x6d, 0x12, 0x2e, 0x0a, 0x04, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x12, 0x2e, 0x64, 0x74, 0x6d, 0x63,
	0x6c, 0x69, 0x2e, 0x44, 0x74, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e,
	0x64, 0x74, 0x6d, 0x63, 0x6c, 0x69, 0x2e, 0x44, 0x74, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x42, 0x1c, 0x5a, 0x1a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x79, 0x65, 0x64, 0x66, 0x2f, 0x64, 0x74, 0x6d, 0x2f, 0x64, 0x74, 0x6d, 0x63, 0x6c, 0x69, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dtmcli_dtmcli_proto_rawDescOnce sync.Once
	file_dtmcli_dtmcli_proto_rawDescData = file_dtmcli_dtmcli_proto_rawDesc
)

func file_dtmcli_dtmcli_proto_rawDescGZIP() []byte {
	file_dtmcli_dtmcli_proto_rawDescOnce.Do(func() {
		file_dtmcli_dtmcli_proto_rawDescData = protoimpl.X.CompressGZIP(file_dtmcli_dtmcli_proto_rawDescData)
	})
	return file_dtmcli_dtmcli_proto_rawDescData
}

var file_dtmcli_dtmcli_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_dtmcli_dtmcli_proto_goTypes = []interface{}{
	(*DtmTransInfo)(nil), // 0: dtmcli.DtmTransInfo
	(*DtmRequest)(nil),   // 1: dtmcli.DtmRequest
	(*DtmReply)(nil),     // 2: dtmcli.DtmReply
	(*BusiRequest)(nil),  // 3: dtmcli.BusiRequest
	(*BusiReply)(nil),    // 4: dtmcli.BusiReply
	nil,                  // 5: dtmcli.DtmRequest.ExtraEntry
	nil,                  // 6: dtmcli.BusiRequest.ExtraEntry
}
var file_dtmcli_dtmcli_proto_depIdxs = []int32{
	5, // 0: dtmcli.DtmRequest.Extra:type_name -> dtmcli.DtmRequest.ExtraEntry
	0, // 1: dtmcli.BusiRequest.info:type_name -> dtmcli.DtmTransInfo
	6, // 2: dtmcli.BusiRequest.Extra:type_name -> dtmcli.BusiRequest.ExtraEntry
	1, // 3: dtmcli.Dtm.Call:input_type -> dtmcli.DtmRequest
	2, // 4: dtmcli.Dtm.Call:output_type -> dtmcli.DtmReply
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_dtmcli_dtmcli_proto_init() }
func file_dtmcli_dtmcli_proto_init() {
	if File_dtmcli_dtmcli_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dtmcli_dtmcli_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DtmTransInfo); i {
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
		file_dtmcli_dtmcli_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DtmRequest); i {
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
		file_dtmcli_dtmcli_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DtmReply); i {
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
		file_dtmcli_dtmcli_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BusiRequest); i {
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
		file_dtmcli_dtmcli_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BusiReply); i {
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
			RawDescriptor: file_dtmcli_dtmcli_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dtmcli_dtmcli_proto_goTypes,
		DependencyIndexes: file_dtmcli_dtmcli_proto_depIdxs,
		MessageInfos:      file_dtmcli_dtmcli_proto_msgTypes,
	}.Build()
	File_dtmcli_dtmcli_proto = out.File
	file_dtmcli_dtmcli_proto_rawDesc = nil
	file_dtmcli_dtmcli_proto_goTypes = nil
	file_dtmcli_dtmcli_proto_depIdxs = nil
}
