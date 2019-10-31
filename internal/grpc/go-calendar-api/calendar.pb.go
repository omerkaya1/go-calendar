// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/calendar.proto

package calendar

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Event struct {
	EventID              string               `protobuf:"bytes,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	UserName             string               `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	EventName            string               `protobuf:"bytes,3,opt,name=event_name,json=eventName,proto3" json:"event_name,omitempty"`
	Note                 string               `protobuf:"bytes,4,opt,name=note,proto3" json:"note,omitempty"`
	StartTime            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime              *timestamp.Timestamp `protobuf:"bytes,6,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f3e00302f51dd79, []int{0}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetEventID() string {
	if m != nil {
		return m.EventID
	}
	return ""
}

func (m *Event) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *Event) GetEventName() string {
	if m != nil {
		return m.EventName
	}
	return ""
}

func (m *Event) GetNote() string {
	if m != nil {
		return m.Note
	}
	return ""
}

func (m *Event) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *Event) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

type CreateEventRequest struct {
	EventName            string               `protobuf:"bytes,1,opt,name=event_name,json=eventName,proto3" json:"event_name,omitempty"`
	Text                 string               `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	UserName             string               `protobuf:"bytes,3,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	StartTime            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime              *timestamp.Timestamp `protobuf:"bytes,5,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CreateEventRequest) Reset()         { *m = CreateEventRequest{} }
func (m *CreateEventRequest) String() string { return proto.CompactTextString(m) }
func (*CreateEventRequest) ProtoMessage()    {}
func (*CreateEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f3e00302f51dd79, []int{1}
}

func (m *CreateEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateEventRequest.Unmarshal(m, b)
}
func (m *CreateEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateEventRequest.Marshal(b, m, deterministic)
}
func (m *CreateEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateEventRequest.Merge(m, src)
}
func (m *CreateEventRequest) XXX_Size() int {
	return xxx_messageInfo_CreateEventRequest.Size(m)
}
func (m *CreateEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateEventRequest proto.InternalMessageInfo

func (m *CreateEventRequest) GetEventName() string {
	if m != nil {
		return m.EventName
	}
	return ""
}

func (m *CreateEventRequest) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *CreateEventRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *CreateEventRequest) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *CreateEventRequest) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

type ResponseWithEvent struct {
	// Types that are valid to be assigned to Result:
	//	*ResponseWithEvent_Event
	//	*ResponseWithEvent_Error
	Result               isResponseWithEvent_Result `protobuf_oneof:"result"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ResponseWithEvent) Reset()         { *m = ResponseWithEvent{} }
func (m *ResponseWithEvent) String() string { return proto.CompactTextString(m) }
func (*ResponseWithEvent) ProtoMessage()    {}
func (*ResponseWithEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f3e00302f51dd79, []int{2}
}

func (m *ResponseWithEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseWithEvent.Unmarshal(m, b)
}
func (m *ResponseWithEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseWithEvent.Marshal(b, m, deterministic)
}
func (m *ResponseWithEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseWithEvent.Merge(m, src)
}
func (m *ResponseWithEvent) XXX_Size() int {
	return xxx_messageInfo_ResponseWithEvent.Size(m)
}
func (m *ResponseWithEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseWithEvent.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseWithEvent proto.InternalMessageInfo

type isResponseWithEvent_Result interface {
	isResponseWithEvent_Result()
}

type ResponseWithEvent_Event struct {
	Event *Event `protobuf:"bytes,1,opt,name=event,proto3,oneof"`
}

type ResponseWithEvent_Error struct {
	Error string `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

func (*ResponseWithEvent_Event) isResponseWithEvent_Result() {}

func (*ResponseWithEvent_Error) isResponseWithEvent_Result() {}

func (m *ResponseWithEvent) GetResult() isResponseWithEvent_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *ResponseWithEvent) GetEvent() *Event {
	if x, ok := m.GetResult().(*ResponseWithEvent_Event); ok {
		return x.Event
	}
	return nil
}

func (m *ResponseWithEvent) GetError() string {
	if x, ok := m.GetResult().(*ResponseWithEvent_Error); ok {
		return x.Error
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ResponseWithEvent) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ResponseWithEvent_Event)(nil),
		(*ResponseWithEvent_Error)(nil),
	}
}

type ResponseWithEventID struct {
	// Types that are valid to be assigned to Result:
	//	*ResponseWithEventID_EventID
	//	*ResponseWithEventID_Error
	Result               isResponseWithEventID_Result `protobuf_oneof:"result"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *ResponseWithEventID) Reset()         { *m = ResponseWithEventID{} }
func (m *ResponseWithEventID) String() string { return proto.CompactTextString(m) }
func (*ResponseWithEventID) ProtoMessage()    {}
func (*ResponseWithEventID) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f3e00302f51dd79, []int{3}
}

func (m *ResponseWithEventID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseWithEventID.Unmarshal(m, b)
}
func (m *ResponseWithEventID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseWithEventID.Marshal(b, m, deterministic)
}
func (m *ResponseWithEventID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseWithEventID.Merge(m, src)
}
func (m *ResponseWithEventID) XXX_Size() int {
	return xxx_messageInfo_ResponseWithEventID.Size(m)
}
func (m *ResponseWithEventID) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseWithEventID.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseWithEventID proto.InternalMessageInfo

type isResponseWithEventID_Result interface {
	isResponseWithEventID_Result()
}

type ResponseWithEventID_EventID struct {
	EventID string `protobuf:"bytes,1,opt,name=eventID,proto3,oneof"`
}

type ResponseWithEventID_Error struct {
	Error string `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

func (*ResponseWithEventID_EventID) isResponseWithEventID_Result() {}

func (*ResponseWithEventID_Error) isResponseWithEventID_Result() {}

func (m *ResponseWithEventID) GetResult() isResponseWithEventID_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *ResponseWithEventID) GetEventID() string {
	if x, ok := m.GetResult().(*ResponseWithEventID_EventID); ok {
		return x.EventID
	}
	return ""
}

func (m *ResponseWithEventID) GetError() string {
	if x, ok := m.GetResult().(*ResponseWithEventID_Error); ok {
		return x.Error
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ResponseWithEventID) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ResponseWithEventID_EventID)(nil),
		(*ResponseWithEventID_Error)(nil),
	}
}

type RequestEventByID struct {
	EventID              string   `protobuf:"bytes,1,opt,name=eventID,proto3" json:"eventID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestEventByID) Reset()         { *m = RequestEventByID{} }
func (m *RequestEventByID) String() string { return proto.CompactTextString(m) }
func (*RequestEventByID) ProtoMessage()    {}
func (*RequestEventByID) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f3e00302f51dd79, []int{4}
}

func (m *RequestEventByID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestEventByID.Unmarshal(m, b)
}
func (m *RequestEventByID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestEventByID.Marshal(b, m, deterministic)
}
func (m *RequestEventByID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestEventByID.Merge(m, src)
}
func (m *RequestEventByID) XXX_Size() int {
	return xxx_messageInfo_RequestEventByID.Size(m)
}
func (m *RequestEventByID) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestEventByID.DiscardUnknown(m)
}

var xxx_messageInfo_RequestEventByID proto.InternalMessageInfo

func (m *RequestEventByID) GetEventID() string {
	if m != nil {
		return m.EventID
	}
	return ""
}

type ResponseSuccess struct {
	// Types that are valid to be assigned to Result:
	//	*ResponseSuccess_Response
	//	*ResponseSuccess_Error
	Result               isResponseSuccess_Result `protobuf_oneof:"result"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ResponseSuccess) Reset()         { *m = ResponseSuccess{} }
func (m *ResponseSuccess) String() string { return proto.CompactTextString(m) }
func (*ResponseSuccess) ProtoMessage()    {}
func (*ResponseSuccess) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f3e00302f51dd79, []int{5}
}

func (m *ResponseSuccess) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseSuccess.Unmarshal(m, b)
}
func (m *ResponseSuccess) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseSuccess.Marshal(b, m, deterministic)
}
func (m *ResponseSuccess) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseSuccess.Merge(m, src)
}
func (m *ResponseSuccess) XXX_Size() int {
	return xxx_messageInfo_ResponseSuccess.Size(m)
}
func (m *ResponseSuccess) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseSuccess.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseSuccess proto.InternalMessageInfo

type isResponseSuccess_Result interface {
	isResponseSuccess_Result()
}

type ResponseSuccess_Response struct {
	Response string `protobuf:"bytes,1,opt,name=response,proto3,oneof"`
}

type ResponseSuccess_Error struct {
	Error string `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

func (*ResponseSuccess_Response) isResponseSuccess_Result() {}

func (*ResponseSuccess_Error) isResponseSuccess_Result() {}

func (m *ResponseSuccess) GetResult() isResponseSuccess_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *ResponseSuccess) GetResponse() string {
	if x, ok := m.GetResult().(*ResponseSuccess_Response); ok {
		return x.Response
	}
	return ""
}

func (m *ResponseSuccess) GetError() string {
	if x, ok := m.GetResult().(*ResponseSuccess_Error); ok {
		return x.Error
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ResponseSuccess) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ResponseSuccess_Response)(nil),
		(*ResponseSuccess_Error)(nil),
	}
}

func init() {
	proto.RegisterType((*Event)(nil), "Event")
	proto.RegisterType((*CreateEventRequest)(nil), "CreateEventRequest")
	proto.RegisterType((*ResponseWithEvent)(nil), "ResponseWithEvent")
	proto.RegisterType((*ResponseWithEventID)(nil), "ResponseWithEventID")
	proto.RegisterType((*RequestEventByID)(nil), "RequestEventByID")
	proto.RegisterType((*ResponseSuccess)(nil), "ResponseSuccess")
}

func init() { proto.RegisterFile("api/calendar.proto", fileDescriptor_1f3e00302f51dd79) }

var fileDescriptor_1f3e00302f51dd79 = []byte{
	// 454 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x8d, 0x69, 0x9c, 0x3a, 0xe3, 0x03, 0xe9, 0x14, 0x21, 0x63, 0xa0, 0x54, 0x3e, 0xf5, 0x40,
	0x36, 0x52, 0x0a, 0x07, 0x38, 0xb6, 0x46, 0x6d, 0x2e, 0x48, 0x38, 0x54, 0x1c, 0x23, 0x37, 0x1e,
	0x42, 0xa4, 0xc4, 0x6b, 0x76, 0xd7, 0x15, 0xfc, 0x51, 0x7e, 0x01, 0x07, 0x7e, 0x06, 0xf2, 0xac,
	0x1d, 0x51, 0x27, 0xa8, 0x55, 0x6f, 0x3b, 0x6f, 0xbe, 0xde, 0x7b, 0x1e, 0xc3, 0xd1, 0x42, 0x0e,
	0xe7, 0xe9, 0x8a, 0xf2, 0x2c, 0x55, 0xc3, 0xb4, 0x58, 0x8e, 0x9a, 0x40, 0x14, 0x4a, 0x1a, 0x19,
	0xbe, 0x5a, 0x48, 0xb9, 0x58, 0xd1, 0x88, 0xa3, 0xeb, 0xf2, 0xeb, 0xc8, 0x2c, 0xd7, 0xa4, 0x4d,
	0xba, 0x2e, 0x6c, 0x41, 0xf4, 0xdb, 0x01, 0xf7, 0xc3, 0x0d, 0xe5, 0x06, 0x9f, 0x81, 0x47, 0xd5,
	0x63, 0xb6, 0xcc, 0x02, 0xe7, 0xd8, 0x39, 0xe9, 0x27, 0xfb, 0x1c, 0x4f, 0x32, 0x7c, 0x0e, 0xfd,
	0x52, 0x93, 0x9a, 0xe5, 0xe9, 0x9a, 0x82, 0x47, 0x9c, 0xf3, 0x2a, 0xe0, 0x63, 0xba, 0x26, 0x7c,
	0x09, 0x60, 0xfb, 0x38, 0xbb, 0xc7, 0xd9, 0x3e, 0x23, 0x9c, 0x46, 0xe8, 0xe6, 0xd2, 0x50, 0xd0,
	0xe5, 0x04, 0xbf, 0xf1, 0x1d, 0x80, 0x36, 0xa9, 0x32, 0xb3, 0x8a, 0x4d, 0xe0, 0x1e, 0x3b, 0x27,
	0xfe, 0x38, 0x14, 0x96, 0xaa, 0x68, 0xa8, 0x8a, 0xcf, 0x0d, 0xd5, 0xa4, 0xcf, 0xd5, 0x55, 0x8c,
	0x6f, 0xc1, 0xa3, 0x3c, 0xb3, 0x8d, 0xbd, 0x3b, 0x1b, 0xf7, 0x29, 0xcf, 0xaa, 0x28, 0xfa, 0xe5,
	0x00, 0x9e, 0x2b, 0x4a, 0x0d, 0xb1, 0xd8, 0x84, 0xbe, 0x97, 0xa4, 0x4d, 0x8b, 0xbb, 0xb3, 0x83,
	0xbb, 0xa1, 0x1f, 0xa6, 0x96, 0xcc, 0xef, 0xdb, 0x5e, 0xec, 0xb5, 0xbc, 0xb8, 0x2d, 0xac, 0xfb,
	0x50, 0x61, 0xee, 0xfd, 0x85, 0x5d, 0xc1, 0x41, 0x42, 0xba, 0x90, 0xb9, 0xa6, 0x2f, 0x4b, 0xf3,
	0xcd, 0x7e, 0xca, 0x23, 0x70, 0x59, 0x04, 0x2b, 0xf2, 0xc7, 0x3d, 0xc1, 0xf0, 0x65, 0x27, 0xb1,
	0x30, 0x3e, 0x05, 0x97, 0x94, 0x92, 0xca, 0x0a, 0x63, 0xbc, 0x0a, 0xcf, 0x3c, 0xe8, 0x29, 0xd2,
	0xe5, 0xca, 0x44, 0x53, 0x38, 0xdc, 0x1a, 0x3b, 0x89, 0x31, 0x84, 0xfa, 0x26, 0x62, 0x6b, 0xd6,
	0x65, 0xa7, 0x39, 0x92, 0xf8, 0x1e, 0x43, 0x5f, 0xc3, 0xa0, 0x36, 0x9e, 0xe7, 0x9d, 0xfd, 0x9c,
	0xc4, 0x18, 0xb4, 0x26, 0x6e, 0xe6, 0x45, 0x9f, 0xe0, 0x71, 0x43, 0x61, 0x5a, 0xce, 0xe7, 0xa4,
	0x35, 0xbe, 0x00, 0x4f, 0xd5, 0xd0, 0x66, 0xff, 0x06, 0xb9, 0x9b, 0xc0, 0xf8, 0x8f, 0x03, 0x83,
	0x0b, 0x79, 0x5e, 0xff, 0x22, 0x53, 0x52, 0x37, 0xa4, 0xf0, 0x3d, 0xf8, 0xff, 0x5c, 0x06, 0x1e,
	0x8a, 0xed, 0x3b, 0x09, 0x9f, 0x88, 0x1d, 0x6e, 0x44, 0x1d, 0x3c, 0x05, 0xef, 0x82, 0xac, 0x1a,
	0x3c, 0x10, 0x6d, 0x71, 0x21, 0x6e, 0xb7, 0x45, 0x1d, 0x1c, 0x82, 0x7f, 0x55, 0x64, 0x9b, 0x85,
	0xf5, 0xd7, 0xf9, 0xef, 0x8e, 0x37, 0xe0, 0xc7, 0xb4, 0xa2, 0xa6, 0x7c, 0xc7, 0x9a, 0x81, 0x68,
	0x19, 0x15, 0x75, 0xae, 0x7b, 0x7c, 0x34, 0xa7, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xf6, 0x9a,
	0xea, 0x12, 0x21, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GoCalendarServerClient is the client API for GoCalendarServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GoCalendarServerClient interface {
	// CreateEvent creates a new event and stores it in the DB. It returns an event id on success or error on failure.
	CreateEvent(ctx context.Context, in *CreateEventRequest, opts ...grpc.CallOption) (*ResponseWithEventID, error)
	// GetEventByID returns an event if it exists, otherwise it returns an error.
	GetEvent(ctx context.Context, in *RequestEventByID, opts ...grpc.CallOption) (*ResponseWithEvent, error)
	// UpdateEventByID updates an existing event and returns an event id on success or an error on failure.
	UpdateEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*ResponseWithEventID, error)
	// DeleteEventByID deletes an event from the DB. It returns an event id on success or an error on failure.
	DeleteEvent(ctx context.Context, in *RequestEventByID, opts ...grpc.CallOption) (*ResponseSuccess, error)
}

type goCalendarServerClient struct {
	cc *grpc.ClientConn
}

func NewGoCalendarServerClient(cc *grpc.ClientConn) GoCalendarServerClient {
	return &goCalendarServerClient{cc}
}

func (c *goCalendarServerClient) CreateEvent(ctx context.Context, in *CreateEventRequest, opts ...grpc.CallOption) (*ResponseWithEventID, error) {
	out := new(ResponseWithEventID)
	err := c.cc.Invoke(ctx, "/GoCalendarServer/CreateEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goCalendarServerClient) GetEvent(ctx context.Context, in *RequestEventByID, opts ...grpc.CallOption) (*ResponseWithEvent, error) {
	out := new(ResponseWithEvent)
	err := c.cc.Invoke(ctx, "/GoCalendarServer/GetEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goCalendarServerClient) UpdateEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*ResponseWithEventID, error) {
	out := new(ResponseWithEventID)
	err := c.cc.Invoke(ctx, "/GoCalendarServer/UpdateEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goCalendarServerClient) DeleteEvent(ctx context.Context, in *RequestEventByID, opts ...grpc.CallOption) (*ResponseSuccess, error) {
	out := new(ResponseSuccess)
	err := c.cc.Invoke(ctx, "/GoCalendarServer/DeleteEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GoCalendarServerServer is the server API for GoCalendarServer service.
type GoCalendarServerServer interface {
	// CreateEvent creates a new event and stores it in the DB. It returns an event id on success or error on failure.
	CreateEvent(context.Context, *CreateEventRequest) (*ResponseWithEventID, error)
	// GetEventByID returns an event if it exists, otherwise it returns an error.
	GetEvent(context.Context, *RequestEventByID) (*ResponseWithEvent, error)
	// UpdateEventByID updates an existing event and returns an event id on success or an error on failure.
	UpdateEvent(context.Context, *Event) (*ResponseWithEventID, error)
	// DeleteEventByID deletes an event from the DB. It returns an event id on success or an error on failure.
	DeleteEvent(context.Context, *RequestEventByID) (*ResponseSuccess, error)
}

// UnimplementedGoCalendarServerServer can be embedded to have forward compatible implementations.
type UnimplementedGoCalendarServerServer struct {
}

func (*UnimplementedGoCalendarServerServer) CreateEvent(ctx context.Context, req *CreateEventRequest) (*ResponseWithEventID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEvent not implemented")
}
func (*UnimplementedGoCalendarServerServer) GetEvent(ctx context.Context, req *RequestEventByID) (*ResponseWithEvent, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEvent not implemented")
}
func (*UnimplementedGoCalendarServerServer) UpdateEvent(ctx context.Context, req *Event) (*ResponseWithEventID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEvent not implemented")
}
func (*UnimplementedGoCalendarServerServer) DeleteEvent(ctx context.Context, req *RequestEventByID) (*ResponseSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEvent not implemented")
}

func RegisterGoCalendarServerServer(s *grpc.Server, srv GoCalendarServerServer) {
	s.RegisterService(&_GoCalendarServer_serviceDesc, srv)
}

func _GoCalendarServer_CreateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoCalendarServerServer).CreateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GoCalendarServer/CreateEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoCalendarServerServer).CreateEvent(ctx, req.(*CreateEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoCalendarServer_GetEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestEventByID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoCalendarServerServer).GetEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GoCalendarServer/GetEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoCalendarServerServer).GetEvent(ctx, req.(*RequestEventByID))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoCalendarServer_UpdateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoCalendarServerServer).UpdateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GoCalendarServer/UpdateEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoCalendarServerServer).UpdateEvent(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoCalendarServer_DeleteEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestEventByID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoCalendarServerServer).DeleteEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GoCalendarServer/DeleteEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoCalendarServerServer).DeleteEvent(ctx, req.(*RequestEventByID))
	}
	return interceptor(ctx, in, info, handler)
}

var _GoCalendarServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "GoCalendarServer",
	HandlerType: (*GoCalendarServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEvent",
			Handler:    _GoCalendarServer_CreateEvent_Handler,
		},
		{
			MethodName: "GetEvent",
			Handler:    _GoCalendarServer_GetEvent_Handler,
		},
		{
			MethodName: "UpdateEvent",
			Handler:    _GoCalendarServer_UpdateEvent_Handler,
		},
		{
			MethodName: "DeleteEvent",
			Handler:    _GoCalendarServer_DeleteEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/calendar.proto",
}
