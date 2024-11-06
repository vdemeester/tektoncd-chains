// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v5.26.1
// source: envoy/config/core/v3/http_service.proto

package corev3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// HTTP service configuration.
type HttpService struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The service's HTTP URI. For example:
	//
	// .. code-block:: yaml
	//
	//	http_uri:
	//	  uri: https://www.myserviceapi.com/v1/data
	//	  cluster: www.myserviceapi.com|443
	HttpUri *HttpUri `protobuf:"bytes,1,opt,name=http_uri,json=httpUri,proto3" json:"http_uri,omitempty"`
	// Specifies a list of HTTP headers that should be added to each request
	// handled by this virtual host.
	RequestHeadersToAdd []*HeaderValueOption `protobuf:"bytes,2,rep,name=request_headers_to_add,json=requestHeadersToAdd,proto3" json:"request_headers_to_add,omitempty"`
}

func (x *HttpService) Reset() {
	*x = HttpService{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_core_v3_http_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HttpService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpService) ProtoMessage() {}

func (x *HttpService) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_core_v3_http_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpService.ProtoReflect.Descriptor instead.
func (*HttpService) Descriptor() ([]byte, []int) {
	return file_envoy_config_core_v3_http_service_proto_rawDescGZIP(), []int{0}
}

func (x *HttpService) GetHttpUri() *HttpUri {
	if x != nil {
		return x.HttpUri
	}
	return nil
}

func (x *HttpService) GetRequestHeadersToAdd() []*HeaderValueOption {
	if x != nil {
		return x.RequestHeadersToAdd
	}
	return nil
}

var File_envoy_config_core_v3_http_service_proto protoreflect.FileDescriptor

var file_envoy_config_core_v3_http_service_proto_rawDesc = []byte{
	0x0a, 0x27, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x1a,
	0x1f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x23, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x75, 0x72, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb0, 0x01,
	0x0a, 0x0b, 0x48, 0x74, 0x74, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a,
	0x08, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x55, 0x72, 0x69, 0x52, 0x07,
	0x68, 0x74, 0x74, 0x70, 0x55, 0x72, 0x69, 0x12, 0x67, 0x0a, 0x16, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x74, 0x6f, 0x5f, 0x61, 0x64,
	0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x09, 0xfa, 0x42, 0x06, 0x92, 0x01, 0x03, 0x10, 0xe8, 0x07, 0x52, 0x13, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x54, 0x6f, 0x41, 0x64, 0x64,
	0x42, 0x84, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x0a, 0x22, 0x69, 0x6f, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x42,
	0x10, 0x48, 0x74, 0x74, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33,
	0x3b, 0x63, 0x6f, 0x72, 0x65, 0x76, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_config_core_v3_http_service_proto_rawDescOnce sync.Once
	file_envoy_config_core_v3_http_service_proto_rawDescData = file_envoy_config_core_v3_http_service_proto_rawDesc
)

func file_envoy_config_core_v3_http_service_proto_rawDescGZIP() []byte {
	file_envoy_config_core_v3_http_service_proto_rawDescOnce.Do(func() {
		file_envoy_config_core_v3_http_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_config_core_v3_http_service_proto_rawDescData)
	})
	return file_envoy_config_core_v3_http_service_proto_rawDescData
}

var file_envoy_config_core_v3_http_service_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_config_core_v3_http_service_proto_goTypes = []interface{}{
	(*HttpService)(nil),       // 0: envoy.config.core.v3.HttpService
	(*HttpUri)(nil),           // 1: envoy.config.core.v3.HttpUri
	(*HeaderValueOption)(nil), // 2: envoy.config.core.v3.HeaderValueOption
}
var file_envoy_config_core_v3_http_service_proto_depIdxs = []int32{
	1, // 0: envoy.config.core.v3.HttpService.http_uri:type_name -> envoy.config.core.v3.HttpUri
	2, // 1: envoy.config.core.v3.HttpService.request_headers_to_add:type_name -> envoy.config.core.v3.HeaderValueOption
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_envoy_config_core_v3_http_service_proto_init() }
func file_envoy_config_core_v3_http_service_proto_init() {
	if File_envoy_config_core_v3_http_service_proto != nil {
		return
	}
	file_envoy_config_core_v3_base_proto_init()
	file_envoy_config_core_v3_http_uri_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_envoy_config_core_v3_http_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HttpService); i {
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
			RawDescriptor: file_envoy_config_core_v3_http_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_config_core_v3_http_service_proto_goTypes,
		DependencyIndexes: file_envoy_config_core_v3_http_service_proto_depIdxs,
		MessageInfos:      file_envoy_config_core_v3_http_service_proto_msgTypes,
	}.Build()
	File_envoy_config_core_v3_http_service_proto = out.File
	file_envoy_config_core_v3_http_service_proto_rawDesc = nil
	file_envoy_config_core_v3_http_service_proto_goTypes = nil
	file_envoy_config_core_v3_http_service_proto_depIdxs = nil
}
