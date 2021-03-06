// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/enums/keyword_plan_forecast_interval.proto

package enums

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Forecast intervals.
type KeywordPlanForecastIntervalEnum_KeywordPlanForecastInterval int32

const (
	// Not specified.
	KeywordPlanForecastIntervalEnum_UNSPECIFIED KeywordPlanForecastIntervalEnum_KeywordPlanForecastInterval = 0
	// The value is unknown in this version.
	KeywordPlanForecastIntervalEnum_UNKNOWN KeywordPlanForecastIntervalEnum_KeywordPlanForecastInterval = 1
	// The next week date range for keyword plan. The next week is based
	// on the default locale of the user's account and is mostly SUN-SAT or
	// MON-SUN.
	// This can be different from next-7 days.
	KeywordPlanForecastIntervalEnum_NEXT_WEEK KeywordPlanForecastIntervalEnum_KeywordPlanForecastInterval = 3
	// The next month date range for keyword plan.
	KeywordPlanForecastIntervalEnum_NEXT_MONTH KeywordPlanForecastIntervalEnum_KeywordPlanForecastInterval = 4
	// The next quarter date range for keyword plan.
	KeywordPlanForecastIntervalEnum_NEXT_QUARTER KeywordPlanForecastIntervalEnum_KeywordPlanForecastInterval = 5
)

var KeywordPlanForecastIntervalEnum_KeywordPlanForecastInterval_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	3: "NEXT_WEEK",
	4: "NEXT_MONTH",
	5: "NEXT_QUARTER",
}

var KeywordPlanForecastIntervalEnum_KeywordPlanForecastInterval_value = map[string]int32{
	"UNSPECIFIED":  0,
	"UNKNOWN":      1,
	"NEXT_WEEK":    3,
	"NEXT_MONTH":   4,
	"NEXT_QUARTER": 5,
}

func (x KeywordPlanForecastIntervalEnum_KeywordPlanForecastInterval) String() string {
	return proto.EnumName(KeywordPlanForecastIntervalEnum_KeywordPlanForecastInterval_name, int32(x))
}

func (KeywordPlanForecastIntervalEnum_KeywordPlanForecastInterval) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0edc75f1de38c675, []int{0, 0}
}

// Container for enumeration of forecast intervals.
type KeywordPlanForecastIntervalEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeywordPlanForecastIntervalEnum) Reset()         { *m = KeywordPlanForecastIntervalEnum{} }
func (m *KeywordPlanForecastIntervalEnum) String() string { return proto.CompactTextString(m) }
func (*KeywordPlanForecastIntervalEnum) ProtoMessage()    {}
func (*KeywordPlanForecastIntervalEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_0edc75f1de38c675, []int{0}
}

func (m *KeywordPlanForecastIntervalEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeywordPlanForecastIntervalEnum.Unmarshal(m, b)
}
func (m *KeywordPlanForecastIntervalEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeywordPlanForecastIntervalEnum.Marshal(b, m, deterministic)
}
func (m *KeywordPlanForecastIntervalEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeywordPlanForecastIntervalEnum.Merge(m, src)
}
func (m *KeywordPlanForecastIntervalEnum) XXX_Size() int {
	return xxx_messageInfo_KeywordPlanForecastIntervalEnum.Size(m)
}
func (m *KeywordPlanForecastIntervalEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_KeywordPlanForecastIntervalEnum.DiscardUnknown(m)
}

var xxx_messageInfo_KeywordPlanForecastIntervalEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.enums.KeywordPlanForecastIntervalEnum_KeywordPlanForecastInterval", KeywordPlanForecastIntervalEnum_KeywordPlanForecastInterval_name, KeywordPlanForecastIntervalEnum_KeywordPlanForecastInterval_value)
	proto.RegisterType((*KeywordPlanForecastIntervalEnum)(nil), "google.ads.googleads.v2.enums.KeywordPlanForecastIntervalEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/enums/keyword_plan_forecast_interval.proto", fileDescriptor_0edc75f1de38c675)
}

var fileDescriptor_0edc75f1de38c675 = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0x4b, 0x4e, 0xc3, 0x30,
	0x10, 0x25, 0x2d, 0x1f, 0xe1, 0xf2, 0x89, 0xb2, 0x04, 0x2a, 0x68, 0x0f, 0xe0, 0x48, 0x61, 0x67,
	0x56, 0x29, 0xb8, 0xa5, 0xaa, 0x48, 0x43, 0xe9, 0x07, 0xa1, 0x48, 0x91, 0x69, 0x8c, 0x15, 0xe1,
	0xda, 0x51, 0x9c, 0x16, 0x71, 0x0a, 0xee, 0xc0, 0x92, 0xa3, 0x70, 0x14, 0xf6, 0xec, 0x51, 0xec,
	0xb6, 0x3b, 0xba, 0xb1, 0xde, 0x78, 0x66, 0xde, 0x9b, 0xf7, 0x40, 0x8b, 0x49, 0xc9, 0x38, 0x75,
	0x49, 0xa2, 0x5c, 0x03, 0x4b, 0xb4, 0xf0, 0x5c, 0x2a, 0xe6, 0x33, 0xe5, 0xbe, 0xd2, 0xf7, 0x37,
	0x99, 0x27, 0x71, 0xc6, 0x89, 0x88, 0x5f, 0x64, 0x4e, 0xa7, 0x44, 0x15, 0x71, 0x2a, 0x0a, 0x9a,
	0x2f, 0x08, 0x87, 0x59, 0x2e, 0x0b, 0xe9, 0xd4, 0xcd, 0x22, 0x24, 0x89, 0x82, 0x6b, 0x0e, 0xb8,
	0xf0, 0xa0, 0xe6, 0x38, 0x39, 0x5b, 0x49, 0x64, 0xa9, 0x4b, 0x84, 0x90, 0x05, 0x29, 0x52, 0x29,
	0x94, 0x59, 0x6e, 0x7e, 0x58, 0xe0, 0xbc, 0x67, 0x54, 0x42, 0x4e, 0x44, 0x7b, 0xa9, 0xd1, 0x5d,
	0x4a, 0x60, 0x31, 0x9f, 0x35, 0x39, 0x38, 0xdd, 0x30, 0xe2, 0x1c, 0x83, 0xda, 0x28, 0x78, 0x08,
	0xf1, 0x75, 0xb7, 0xdd, 0xc5, 0x37, 0xf6, 0x96, 0x53, 0x03, 0x7b, 0xa3, 0xa0, 0x17, 0xf4, 0x27,
	0x81, 0x6d, 0x39, 0x87, 0x60, 0x3f, 0xc0, 0x8f, 0xc3, 0x78, 0x82, 0x71, 0xcf, 0xae, 0x3a, 0x47,
	0x00, 0xe8, 0xf2, 0xae, 0x1f, 0x0c, 0x6f, 0xed, 0x6d, 0xc7, 0x06, 0x07, 0xba, 0xbe, 0x1f, 0xf9,
	0x83, 0x21, 0x1e, 0xd8, 0x3b, 0xad, 0x5f, 0x0b, 0x34, 0xa6, 0x72, 0x06, 0x37, 0xba, 0x6a, 0x5d,
	0x6c, 0xb8, 0x28, 0x2c, 0x9d, 0x85, 0xd6, 0xd3, 0x32, 0x5c, 0xc8, 0x24, 0x27, 0x82, 0x41, 0x99,
	0x33, 0x97, 0x51, 0xa1, 0x7d, 0xaf, 0xc2, 0xce, 0x52, 0xf5, 0x4f, 0xf6, 0x57, 0xfa, 0xfd, 0xac,
	0x54, 0x3b, 0xbe, 0xff, 0x55, 0xa9, 0x77, 0x0c, 0x95, 0x9f, 0x28, 0x68, 0x60, 0x89, 0xc6, 0x1e,
	0x2c, 0x03, 0x52, 0xdf, 0xab, 0x7e, 0xe4, 0x27, 0x2a, 0x5a, 0xf7, 0xa3, 0xb1, 0x17, 0xe9, 0xfe,
	0x4f, 0xa5, 0x61, 0x3e, 0x11, 0xf2, 0x13, 0x85, 0xd0, 0x7a, 0x02, 0xa1, 0xb1, 0x87, 0x90, 0x9e,
	0x79, 0xde, 0xd5, 0x87, 0x5d, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0xe3, 0x9a, 0x2b, 0x78, 0x13,
	0x02, 0x00, 0x00,
}
