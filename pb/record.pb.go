// Code generated by protoc-gen-gogo.
// source: record.proto
// DO NOT EDIT!

/*
Package record_pb is a generated protocol buffer package.

It is generated from these files:
	record.proto

It has these top-level messages:
	Record
*/
package record_pb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Record represents a dht record that contains a value
// for a key value pair
type Record struct {
	// The key that references this record
	Key *string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	// The actual value this record is storing
	Value []byte `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	// Time the record was received, set by receiver
	TimeReceived     *string `protobuf:"bytes,5,opt,name=timeReceived" json:"timeReceived,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Record) Reset()         { *m = Record{} }
func (m *Record) String() string { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()    {}

func (m *Record) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *Record) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Record) GetTimeReceived() string {
	if m != nil && m.TimeReceived != nil {
		return *m.TimeReceived
	}
	return ""
}

func init() {
	proto.RegisterType((*Record)(nil), "record.pb.Record")
}