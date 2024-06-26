// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: requests.proto

package wordmaster

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

type WiktionaryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Word     *RequestId         `protobuf:"bytes,1,opt,name=word,proto3" json:"word,omitempty"`
	Contents *RequestedContents `protobuf:"bytes,2,opt,name=contents,proto3" json:"contents,omitempty"`
	Source   *SourceId          `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
}

func (x *WiktionaryRequest) Reset() {
	*x = WiktionaryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_requests_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WiktionaryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WiktionaryRequest) ProtoMessage() {}

func (x *WiktionaryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_requests_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WiktionaryRequest.ProtoReflect.Descriptor instead.
func (*WiktionaryRequest) Descriptor() ([]byte, []int) {
	return file_requests_proto_rawDescGZIP(), []int{0}
}

func (x *WiktionaryRequest) GetWord() *RequestId {
	if x != nil {
		return x.Word
	}
	return nil
}

func (x *WiktionaryRequest) GetContents() *RequestedContents {
	if x != nil {
		return x.Contents
	}
	return nil
}

func (x *WiktionaryRequest) GetSource() *SourceId {
	if x != nil {
		return x.Source
	}
	return nil
}

type RequestedContents struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Definition  bool `protobuf:"varint,1,opt,name=definition,proto3" json:"definition,omitempty"`
	Examples    bool `protobuf:"varint,2,opt,name=examples,proto3" json:"examples,omitempty"`
	Etymology   bool `protobuf:"varint,3,opt,name=etymology,proto3" json:"etymology,omitempty"`
	Ipa         bool `protobuf:"varint,4,opt,name=ipa,proto3" json:"ipa,omitempty"`
	Single      bool `protobuf:"varint,5,opt,name=single,proto3" json:"single,omitempty"`
	Inflections bool `protobuf:"varint,6,opt,name=inflections,proto3" json:"inflections,omitempty"`
}

func (x *RequestedContents) Reset() {
	*x = RequestedContents{}
	if protoimpl.UnsafeEnabled {
		mi := &file_requests_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestedContents) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestedContents) ProtoMessage() {}

func (x *RequestedContents) ProtoReflect() protoreflect.Message {
	mi := &file_requests_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestedContents.ProtoReflect.Descriptor instead.
func (*RequestedContents) Descriptor() ([]byte, []int) {
	return file_requests_proto_rawDescGZIP(), []int{1}
}

func (x *RequestedContents) GetDefinition() bool {
	if x != nil {
		return x.Definition
	}
	return false
}

func (x *RequestedContents) GetExamples() bool {
	if x != nil {
		return x.Examples
	}
	return false
}

func (x *RequestedContents) GetEtymology() bool {
	if x != nil {
		return x.Etymology
	}
	return false
}

func (x *RequestedContents) GetIpa() bool {
	if x != nil {
		return x.Ipa
	}
	return false
}

func (x *RequestedContents) GetSingle() bool {
	if x != nil {
		return x.Single
	}
	return false
}

func (x *RequestedContents) GetInflections() bool {
	if x != nil {
		return x.Inflections
	}
	return false
}

var File_requests_proto protoreflect.FileDescriptor

var file_requests_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x32, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x61, 0x64, 0x72, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x5f, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x69, 0x65, 0x73,
	0x2e, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x9f, 0x02, 0x0a, 0x11, 0x57, 0x69, 0x6b, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x51, 0x0a, 0x04, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x61, 0x64,
	0x72, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x64, 0x69, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x49, 0x64, 0x52, 0x04, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x61, 0x0a, 0x08, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x45, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x61, 0x64, 0x72, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x5f, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x6d,
	0x6f, 0x6e, 0x67, 0x6f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x73, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x54,
	0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3c,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x61, 0x64, 0x72, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x5f, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x69, 0x65, 0x73, 0x2e,
	0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x52, 0x06, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x22, 0xb9, 0x01, 0x0a, 0x11, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x65, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65,
	0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a,
	0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x74, 0x79, 0x6d, 0x6f, 0x6c,
	0x6f, 0x67, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x65, 0x74, 0x79, 0x6d, 0x6f,
	0x6c, 0x6f, 0x67, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x70, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x03, 0x69, 0x70, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x69, 0x6e, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x6e, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f,
	0x67, 0x6e, 0x69, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d, 0x61, 0x64, 0x72, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x77, 0x6f, 0x72, 0x64, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_requests_proto_rawDescOnce sync.Once
	file_requests_proto_rawDescData = file_requests_proto_rawDesc
)

func file_requests_proto_rawDescGZIP() []byte {
	file_requests_proto_rawDescOnce.Do(func() {
		file_requests_proto_rawDescData = protoimpl.X.CompressGZIP(file_requests_proto_rawDescData)
	})
	return file_requests_proto_rawDescData
}

var file_requests_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_requests_proto_goTypes = []interface{}{
	(*WiktionaryRequest)(nil), // 0: com.madr.external_dictionaries.mongomodel.protobuf.WiktionaryRequest
	(*RequestedContents)(nil), // 1: com.madr.external_dictionaries.mongomodel.protobuf.RequestedContents
	(*RequestId)(nil),         // 2: com.madr.external_dictionaries.mongomodel.protobuf.RequestId
	(*SourceId)(nil),          // 3: com.madr.external_dictionaries.mongomodel.protobuf.SourceId
}
var file_requests_proto_depIdxs = []int32{
	2, // 0: com.madr.external_dictionaries.mongomodel.protobuf.WiktionaryRequest.word:type_name -> com.madr.external_dictionaries.mongomodel.protobuf.RequestId
	1, // 1: com.madr.external_dictionaries.mongomodel.protobuf.WiktionaryRequest.contents:type_name -> com.madr.external_dictionaries.mongomodel.protobuf.RequestedContents
	3, // 2: com.madr.external_dictionaries.mongomodel.protobuf.WiktionaryRequest.source:type_name -> com.madr.external_dictionaries.mongomodel.protobuf.SourceId
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_requests_proto_init() }
func file_requests_proto_init() {
	if File_requests_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_requests_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WiktionaryRequest); i {
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
		file_requests_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestedContents); i {
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
			RawDescriptor: file_requests_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_requests_proto_goTypes,
		DependencyIndexes: file_requests_proto_depIdxs,
		MessageInfos:      file_requests_proto_msgTypes,
	}.Build()
	File_requests_proto = out.File
	file_requests_proto_rawDesc = nil
	file_requests_proto_goTypes = nil
	file_requests_proto_depIdxs = nil
}
