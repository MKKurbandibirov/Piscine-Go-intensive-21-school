// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: proto/transmitter.proto

package generated

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_transmitter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_transmitter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_transmitter_proto_rawDescGZIP(), []int{0}
}

type Hello struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId *SessionID `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
}

func (x *Hello) Reset() {
	*x = Hello{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_transmitter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hello) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hello) ProtoMessage() {}

func (x *Hello) ProtoReflect() protoreflect.Message {
	mi := &file_proto_transmitter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hello.ProtoReflect.Descriptor instead.
func (*Hello) Descriptor() ([]byte, []int) {
	return file_proto_transmitter_proto_rawDescGZIP(), []int{1}
}

func (x *Hello) GetSessionId() *SessionID {
	if x != nil {
		return x.SessionId
	}
	return nil
}

type SessionID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *SessionID) Reset() {
	*x = SessionID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_transmitter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionID) ProtoMessage() {}

func (x *SessionID) ProtoReflect() protoreflect.Message {
	mi := &file_proto_transmitter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionID.ProtoReflect.Descriptor instead.
func (*SessionID) Descriptor() ([]byte, []int) {
	return file_proto_transmitter_proto_rawDescGZIP(), []int{2}
}

func (x *SessionID) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type StatisticValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId   *SessionID           `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Frequency   float64              `protobuf:"fixed64,2,opt,name=frequency,proto3" json:"frequency,omitempty"`
	CurrentTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=current_time,json=currentTime,proto3" json:"current_time,omitempty"`
}

func (x *StatisticValue) Reset() {
	*x = StatisticValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_transmitter_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatisticValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatisticValue) ProtoMessage() {}

func (x *StatisticValue) ProtoReflect() protoreflect.Message {
	mi := &file_proto_transmitter_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatisticValue.ProtoReflect.Descriptor instead.
func (*StatisticValue) Descriptor() ([]byte, []int) {
	return file_proto_transmitter_proto_rawDescGZIP(), []int{3}
}

func (x *StatisticValue) GetSessionId() *SessionID {
	if x != nil {
		return x.SessionId
	}
	return nil
}

func (x *StatisticValue) GetFrequency() float64 {
	if x != nil {
		return x.Frequency
	}
	return 0
}

func (x *StatisticValue) GetCurrentTime() *timestamp.Timestamp {
	if x != nil {
		return x.CurrentTime
	}
	return nil
}

var File_proto_transmitter_proto protoreflect.FileDescriptor

var file_proto_transmitter_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74,
	0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x3e, 0x0a, 0x05, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x35, 0x0a, 0x0a, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x49, 0x44, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x22, 0x1b, 0x0a, 0x09, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x22, 0xa4, 0x01,
	0x0a, 0x0e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x35, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x74,
	0x65, 0x72, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x52, 0x09, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x6e, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x66, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x3d, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x32, 0x89, 0x01, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69,
	0x74, 0x74, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x08, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67,
	0x12, 0x12, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x12, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x74,
	0x65, 0x72, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x46, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12, 0x16, 0x2e, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49,
	0x44, 0x1a, 0x1b, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x30, 0x01,
	0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_transmitter_proto_rawDescOnce sync.Once
	file_proto_transmitter_proto_rawDescData = file_proto_transmitter_proto_rawDesc
)

func file_proto_transmitter_proto_rawDescGZIP() []byte {
	file_proto_transmitter_proto_rawDescOnce.Do(func() {
		file_proto_transmitter_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_transmitter_proto_rawDescData)
	})
	return file_proto_transmitter_proto_rawDescData
}

var file_proto_transmitter_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_transmitter_proto_goTypes = []interface{}{
	(*Empty)(nil),               // 0: transmitter.Empty
	(*Hello)(nil),               // 1: transmitter.Hello
	(*SessionID)(nil),           // 2: transmitter.SessionID
	(*StatisticValue)(nil),      // 3: transmitter.StatisticValue
	(*timestamp.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_proto_transmitter_proto_depIdxs = []int32{
	2, // 0: transmitter.Hello.session_id:type_name -> transmitter.SessionID
	2, // 1: transmitter.StatisticValue.session_id:type_name -> transmitter.SessionID
	4, // 2: transmitter.StatisticValue.current_time:type_name -> google.protobuf.Timestamp
	0, // 3: transmitter.Transmitter.Greeting:input_type -> transmitter.Empty
	2, // 4: transmitter.Transmitter.GetStatistics:input_type -> transmitter.SessionID
	1, // 5: transmitter.Transmitter.Greeting:output_type -> transmitter.Hello
	3, // 6: transmitter.Transmitter.GetStatistics:output_type -> transmitter.StatisticValue
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_transmitter_proto_init() }
func file_proto_transmitter_proto_init() {
	if File_proto_transmitter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_transmitter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_proto_transmitter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hello); i {
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
		file_proto_transmitter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionID); i {
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
		file_proto_transmitter_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatisticValue); i {
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
			RawDescriptor: file_proto_transmitter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_transmitter_proto_goTypes,
		DependencyIndexes: file_proto_transmitter_proto_depIdxs,
		MessageInfos:      file_proto_transmitter_proto_msgTypes,
	}.Build()
	File_proto_transmitter_proto = out.File
	file_proto_transmitter_proto_rawDesc = nil
	file_proto_transmitter_proto_goTypes = nil
	file_proto_transmitter_proto_depIdxs = nil
}
