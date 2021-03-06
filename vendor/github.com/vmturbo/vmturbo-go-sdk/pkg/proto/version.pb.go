// Code generated by protoc-gen-go.
// source: version.proto
// DO NOT EDIT!

/*
Package version_dto is a generated protocol buffer package.

It is generated from these files:
	version.proto

It has these top-level messages:
	NegotiationRequest
	NegotiationAnswer
*/
package proto

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type NegotiationAnswer_NegotiationStatus int32

const (
	// Server accepts the version of client's protocol, which is fully functional.
	NegotiationAnswer_ACCEPTED NegotiationAnswer_NegotiationStatus = 0
	// Server either do not know about this version (server is obsolete)
	// or the version is no longer supported
	NegotiationAnswer_REJECTED NegotiationAnswer_NegotiationStatus = 1
	// Used for specific cases, where protocol version can be used, but some another conditions failed
	NegotiationAnswer_OTHER NegotiationAnswer_NegotiationStatus = 2
)

var NegotiationAnswer_NegotiationStatus_name = map[int32]string{
	0: "ACCEPTED",
	1: "REJECTED",
	2: "OTHER",
}
var NegotiationAnswer_NegotiationStatus_value = map[string]int32{
	"ACCEPTED": 0,
	"REJECTED": 1,
	"OTHER":    2,
}

func (x NegotiationAnswer_NegotiationStatus) Enum() *NegotiationAnswer_NegotiationStatus {
	p := new(NegotiationAnswer_NegotiationStatus)
	*p = x
	return p
}
func (x NegotiationAnswer_NegotiationStatus) String() string {
	return proto.EnumName(NegotiationAnswer_NegotiationStatus_name, int32(x))
}
func (x *NegotiationAnswer_NegotiationStatus) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(NegotiationAnswer_NegotiationStatus_value, data, "NegotiationAnswer_NegotiationStatus")
	if err != nil {
		return err
	}
	*x = NegotiationAnswer_NegotiationStatus(value)
	return nil
}

// Version negotiation request. This should be sent by client to server
type NegotiationRequest struct {
	// This is the version of protocol at the client side
	ProtocolVersion  *string `protobuf:"bytes,1,req,name=protocol_version" json:"protocol_version,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *NegotiationRequest) Reset()         { *m = NegotiationRequest{} }
func (m *NegotiationRequest) String() string { return proto.CompactTextString(m) }
func (*NegotiationRequest) ProtoMessage()    {}

func (m *NegotiationRequest) GetProtocolVersion() string {
	if m != nil && m.ProtocolVersion != nil {
		return *m.ProtocolVersion
	}
	return ""
}

// Negotiation result. Server answers with this message to a request by NegotiationRequest
type NegotiationAnswer struct {
	// The status, that shows, whether this certain version of protocol can be used to
	// communicate between client and server
	NegotiationResult *NegotiationAnswer_NegotiationStatus `protobuf:"varint,1,req,name=negotiation_result,enum=version_dto.NegotiationAnswer_NegotiationStatus" json:"negotiation_result,omitempty"`
	// Description section, explaining the reasons and details of the certain status
	Description      *string `protobuf:"bytes,2,req,name=description" json:"description,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *NegotiationAnswer) Reset()         { *m = NegotiationAnswer{} }
func (m *NegotiationAnswer) String() string { return proto.CompactTextString(m) }
func (*NegotiationAnswer) ProtoMessage()    {}

func (m *NegotiationAnswer) GetNegotiationResult() NegotiationAnswer_NegotiationStatus {
	if m != nil && m.NegotiationResult != nil {
		return *m.NegotiationResult
	}
	return NegotiationAnswer_ACCEPTED
}

func (m *NegotiationAnswer) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return ""
}

func init() {
	proto.RegisterEnum("version_dto.NegotiationAnswer_NegotiationStatus", NegotiationAnswer_NegotiationStatus_name, NegotiationAnswer_NegotiationStatus_value)
}
