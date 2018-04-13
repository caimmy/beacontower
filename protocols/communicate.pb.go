// Code generated by protoc-gen-go. DO NOT EDIT.
// source: communicate.proto

/*
Package beacontower is a generated protocol buffer package.

It is generated from these files:
	communicate.proto

It has these top-level messages:
	ComMessage
	DirectMessage
	BroadcaseMessage
*/
package beacontower

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// 基础通信过程中的消息信息载体
type ComMessage struct {
	Version   int32  `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	MsgId     string `protobuf:"bytes,2,opt,name=msgId" json:"msgId,omitempty"`
	MsgType   string `protobuf:"bytes,3,opt,name=msgType" json:"msgType,omitempty"`
	Timestamp int32  `protobuf:"varint,4,opt,name=timestamp" json:"timestamp,omitempty"`
	Payload   string `protobuf:"bytes,5,opt,name=payload" json:"payload,omitempty"`
}

func (m *ComMessage) Reset()                    { *m = ComMessage{} }
func (m *ComMessage) String() string            { return proto.CompactTextString(m) }
func (*ComMessage) ProtoMessage()               {}
func (*ComMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ComMessage) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *ComMessage) GetMsgId() string {
	if m != nil {
		return m.MsgId
	}
	return ""
}

func (m *ComMessage) GetMsgType() string {
	if m != nil {
		return m.MsgType
	}
	return ""
}

func (m *ComMessage) GetTimestamp() int32 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *ComMessage) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

// 直接消息
type DirectMessage struct {
	AppId        string      `protobuf:"bytes,1,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	AppKey       string      `protobuf:"bytes,2,opt,name=app_key,json=appKey" json:"app_key,omitempty"`
	AppSecret    string      `protobuf:"bytes,3,opt,name=app_secret,json=appSecret" json:"app_secret,omitempty"`
	FromAccount  string      `protobuf:"bytes,4,opt,name=fromAccount" json:"fromAccount,omitempty"`
	FromResource string      `protobuf:"bytes,5,opt,name=fromResource" json:"fromResource,omitempty"`
	Msg          *ComMessage `protobuf:"bytes,6,opt,name=msg" json:"msg,omitempty"`
}

func (m *DirectMessage) Reset()                    { *m = DirectMessage{} }
func (m *DirectMessage) String() string            { return proto.CompactTextString(m) }
func (*DirectMessage) ProtoMessage()               {}
func (*DirectMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DirectMessage) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *DirectMessage) GetAppKey() string {
	if m != nil {
		return m.AppKey
	}
	return ""
}

func (m *DirectMessage) GetAppSecret() string {
	if m != nil {
		return m.AppSecret
	}
	return ""
}

func (m *DirectMessage) GetFromAccount() string {
	if m != nil {
		return m.FromAccount
	}
	return ""
}

func (m *DirectMessage) GetFromResource() string {
	if m != nil {
		return m.FromResource
	}
	return ""
}

func (m *DirectMessage) GetMsg() *ComMessage {
	if m != nil {
		return m.Msg
	}
	return nil
}

// 广播消息
type BroadcaseMessage struct {
	AppId        string      `protobuf:"bytes,1,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	AppKey       string      `protobuf:"bytes,2,opt,name=app_key,json=appKey" json:"app_key,omitempty"`
	AppSecret    string      `protobuf:"bytes,3,opt,name=app_secret,json=appSecret" json:"app_secret,omitempty"`
	FromAccount  string      `protobuf:"bytes,4,opt,name=fromAccount" json:"fromAccount,omitempty"`
	FromResource string      `protobuf:"bytes,5,opt,name=fromResource" json:"fromResource,omitempty"`
	GroupId      string      `protobuf:"bytes,6,opt,name=group_id,json=groupId" json:"group_id,omitempty"`
	Msg          *ComMessage `protobuf:"bytes,7,opt,name=msg" json:"msg,omitempty"`
}

func (m *BroadcaseMessage) Reset()                    { *m = BroadcaseMessage{} }
func (m *BroadcaseMessage) String() string            { return proto.CompactTextString(m) }
func (*BroadcaseMessage) ProtoMessage()               {}
func (*BroadcaseMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *BroadcaseMessage) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *BroadcaseMessage) GetAppKey() string {
	if m != nil {
		return m.AppKey
	}
	return ""
}

func (m *BroadcaseMessage) GetAppSecret() string {
	if m != nil {
		return m.AppSecret
	}
	return ""
}

func (m *BroadcaseMessage) GetFromAccount() string {
	if m != nil {
		return m.FromAccount
	}
	return ""
}

func (m *BroadcaseMessage) GetFromResource() string {
	if m != nil {
		return m.FromResource
	}
	return ""
}

func (m *BroadcaseMessage) GetGroupId() string {
	if m != nil {
		return m.GroupId
	}
	return ""
}

func (m *BroadcaseMessage) GetMsg() *ComMessage {
	if m != nil {
		return m.Msg
	}
	return nil
}

func init() {
	proto.RegisterType((*ComMessage)(nil), "beacontower.ComMessage")
	proto.RegisterType((*DirectMessage)(nil), "beacontower.DirectMessage")
	proto.RegisterType((*BroadcaseMessage)(nil), "beacontower.BroadcaseMessage")
}

func init() { proto.RegisterFile("communicate.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x92, 0x4f, 0x4e, 0xeb, 0x30,
	0x10, 0x87, 0xe5, 0xd7, 0x97, 0xf4, 0x65, 0xfa, 0x90, 0xc0, 0x02, 0xd5, 0x48, 0x20, 0x45, 0x59,
	0x95, 0x4d, 0x17, 0x70, 0x02, 0xfe, 0x6c, 0x2a, 0xc4, 0x26, 0xb0, 0x47, 0xae, 0x33, 0x44, 0x11,
	0x38, 0x63, 0xd9, 0x0e, 0x28, 0x97, 0xe0, 0x6a, 0x9c, 0x84, 0x3b, 0x20, 0xbb, 0x89, 0x5a, 0x56,
	0x6c, 0x59, 0xfe, 0x3e, 0xcf, 0x58, 0xdf, 0x8c, 0x0d, 0x07, 0x8a, 0xb4, 0xee, 0xda, 0x46, 0x49,
	0x8f, 0x4b, 0x63, 0xc9, 0x13, 0x9f, 0xad, 0x51, 0x2a, 0x6a, 0x3d, 0xbd, 0xa1, 0x2d, 0xde, 0x19,
	0xc0, 0x35, 0xe9, 0x3b, 0x74, 0x4e, 0xd6, 0xc8, 0x05, 0x4c, 0x5f, 0xd1, 0xba, 0x86, 0x5a, 0xc1,
	0x72, 0xb6, 0x48, 0xca, 0x31, 0xf2, 0x43, 0x48, 0xb4, 0xab, 0x57, 0x95, 0xf8, 0x93, 0xb3, 0x45,
	0x56, 0x6e, 0x42, 0xa8, 0xd7, 0xae, 0x7e, 0xe8, 0x0d, 0x8a, 0x49, 0xe4, 0x63, 0xe4, 0x27, 0x90,
	0xf9, 0x46, 0xa3, 0xf3, 0x52, 0x1b, 0xf1, 0x37, 0xde, 0xb5, 0x05, 0xa1, 0xcf, 0xc8, 0xfe, 0x85,
	0x64, 0x25, 0x92, 0x4d, 0xdf, 0x10, 0x8b, 0x0f, 0x06, 0x7b, 0x37, 0x8d, 0x45, 0xe5, 0x47, 0xa7,
	0x23, 0x48, 0xa5, 0x31, 0x8f, 0x4d, 0x15, 0x95, 0xb2, 0x32, 0x91, 0xc6, 0xac, 0x2a, 0x3e, 0x87,
	0x69, 0xc0, 0xcf, 0xd8, 0x0f, 0x4a, 0xa1, 0xea, 0x16, 0x7b, 0x7e, 0x0a, 0x10, 0x0e, 0x1c, 0x2a,
	0x8b, 0x7e, 0xd0, 0xca, 0xa4, 0x31, 0xf7, 0x11, 0xf0, 0x1c, 0x66, 0x4f, 0x96, 0xf4, 0xa5, 0x52,
	0xd4, 0xb5, 0x3e, 0xaa, 0x65, 0xe5, 0x2e, 0xe2, 0x05, 0xfc, 0x0f, 0xb1, 0x44, 0x47, 0x9d, 0x55,
	0x38, 0x18, 0x7e, 0x63, 0xfc, 0x0c, 0x26, 0xda, 0xd5, 0x22, 0xcd, 0xd9, 0x62, 0x76, 0x3e, 0x5f,
	0xee, 0xac, 0x74, 0xb9, 0x5d, 0x67, 0x19, 0x6a, 0x8a, 0x4f, 0x06, 0xfb, 0x57, 0x96, 0x64, 0xa5,
	0xa4, 0xc3, 0x5f, 0x3d, 0xd4, 0x31, 0xfc, 0xab, 0x2d, 0x75, 0x51, 0x2b, 0xdd, 0x3c, 0x4b, 0xcc,
	0xab, 0x6a, 0x9c, 0x77, 0xfa, 0xf3, 0xbc, 0xeb, 0x34, 0x7e, 0xb3, 0x8b, 0xaf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xdc, 0x6a, 0x34, 0x94, 0x7b, 0x02, 0x00, 0x00,
}