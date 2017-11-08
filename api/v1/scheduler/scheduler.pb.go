// Code generated by protoc-gen-go. DO NOT EDIT.
// source: scheduler.proto

/*
Package scheduler is a generated protocol buffer package.

It is generated from these files:
	scheduler.proto

It has these top-level messages:
	Event
	FrameworkID
*/
package scheduler

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

type Event struct {
	Subscribed       *Event_Subscribed `protobuf:"bytes,1,opt,name=subscribed" json:"subscribed,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Event) GetSubscribed() *Event_Subscribed {
	if m != nil {
		return m.Subscribed
	}
	return nil
}

type Event_Subscribed struct {
	FrameworkId      *FrameworkID `protobuf:"bytes,1,req,name=framework_id,json=frameworkId" json:"framework_id,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Event_Subscribed) Reset()                    { *m = Event_Subscribed{} }
func (m *Event_Subscribed) String() string            { return proto.CompactTextString(m) }
func (*Event_Subscribed) ProtoMessage()               {}
func (*Event_Subscribed) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *Event_Subscribed) GetFrameworkId() *FrameworkID {
	if m != nil {
		return m.FrameworkId
	}
	return nil
}

// *
// A unique ID assigned to a framework. A framework can reuse this ID
// in order to do failover (see MesosSchedulerDriver).
type FrameworkID struct {
	Value            *string `protobuf:"bytes,1,req,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *FrameworkID) Reset()                    { *m = FrameworkID{} }
func (m *FrameworkID) String() string            { return proto.CompactTextString(m) }
func (*FrameworkID) ProtoMessage()               {}
func (*FrameworkID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *FrameworkID) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*Event)(nil), "scheduler.Event")
	proto.RegisterType((*Event_Subscribed)(nil), "scheduler.Event.Subscribed")
	proto.RegisterType((*FrameworkID)(nil), "scheduler.FrameworkID")
}

func init() { proto.RegisterFile("scheduler.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 151 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x4e, 0xce, 0x48,
	0x4d, 0x29, 0xcd, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x84, 0x0b, 0x28,
	0xf5, 0x32, 0x72, 0xb1, 0xba, 0x96, 0xa5, 0xe6, 0x95, 0x08, 0x59, 0x73, 0x71, 0x15, 0x97, 0x26,
	0x15, 0x27, 0x17, 0x65, 0x26, 0xa5, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x1b, 0x49, 0xeb,
	0x21, 0xb4, 0x82, 0x55, 0xe9, 0x05, 0xc3, 0x95, 0x04, 0x21, 0x29, 0x97, 0x72, 0xe7, 0xe2, 0x42,
	0xc8, 0x08, 0x59, 0x72, 0xf1, 0xa4, 0x15, 0x25, 0xe6, 0xa6, 0x96, 0xe7, 0x17, 0x65, 0xc7, 0x67,
	0x82, 0x0c, 0x63, 0xd2, 0xe0, 0x36, 0x12, 0x43, 0x32, 0xcc, 0x0d, 0x26, 0xed, 0xe9, 0x12, 0xc4,
	0x0d, 0x57, 0xeb, 0x99, 0xa2, 0xa4, 0xcc, 0xc5, 0x8d, 0x24, 0x27, 0x24, 0xc2, 0xc5, 0x5a, 0x96,
	0x98, 0x53, 0x9a, 0x0a, 0x36, 0x82, 0x33, 0x08, 0xc2, 0x01, 0x04, 0x00, 0x00, 0xff, 0xff, 0x81,
	0xeb, 0xb2, 0xa0, 0xd1, 0x00, 0x00, 0x00,
}
