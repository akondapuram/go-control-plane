// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.4
// source: envoy/config/core/v3/http_uri.proto

package corev3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	duration "github.com/golang/protobuf/ptypes/duration"
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

// Envoy external URI descriptor
type HttpUri struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The HTTP server URI. It should be a full FQDN with protocol, host and path.
	//
	// Example:
	//
	// .. code-block:: yaml
	//
	//	uri: https://www.googleapis.com/oauth2/v1/certs
	Uri string `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	// Specify how “uri“ is to be fetched. Today, this requires an explicit
	// cluster, but in the future we may support dynamic cluster creation or
	// inline DNS resolution. See `issue
	// <https://github.com/envoyproxy/envoy/issues/1606>`_.
	//
	// Types that are assignable to HttpUpstreamType:
	//
	//	*HttpUri_Cluster
	HttpUpstreamType isHttpUri_HttpUpstreamType `protobuf_oneof:"http_upstream_type"`
	// Sets the maximum duration in milliseconds that a response can take to arrive upon request.
	Timeout *duration.Duration `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *HttpUri) Reset() {
	*x = HttpUri{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_core_v3_http_uri_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HttpUri) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpUri) ProtoMessage() {}

func (x *HttpUri) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_core_v3_http_uri_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpUri.ProtoReflect.Descriptor instead.
func (*HttpUri) Descriptor() ([]byte, []int) {
	return file_envoy_config_core_v3_http_uri_proto_rawDescGZIP(), []int{0}
}

func (x *HttpUri) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (m *HttpUri) GetHttpUpstreamType() isHttpUri_HttpUpstreamType {
	if m != nil {
		return m.HttpUpstreamType
	}
	return nil
}

func (x *HttpUri) GetCluster() string {
	if x, ok := x.GetHttpUpstreamType().(*HttpUri_Cluster); ok {
		return x.Cluster
	}
	return ""
}

func (x *HttpUri) GetTimeout() *duration.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

type isHttpUri_HttpUpstreamType interface {
	isHttpUri_HttpUpstreamType()
}

type HttpUri_Cluster struct {
	// A cluster is created in the Envoy "cluster_manager" config
	// section. This field specifies the cluster name.
	//
	// Example:
	//
	// .. code-block:: yaml
	//
	//	cluster: jwks_cluster
	Cluster string `protobuf:"bytes,2,opt,name=cluster,proto3,oneof"`
}

func (*HttpUri_Cluster) isHttpUri_HttpUpstreamType() {}

var File_envoy_config_core_v3_http_uri_proto protoreflect.FileDescriptor

var file_envoy_config_core_v3_http_uri_proto_rawDesc = []byte{
	0x0a, 0x23, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x75, 0x72, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70,
	0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70, 0x61,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcf, 0x01, 0x0a, 0x07, 0x48, 0x74, 0x74, 0x70, 0x55,
	0x72, 0x69, 0x12, 0x19, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12, 0x23, 0x0a,
	0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x48, 0x00, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x12, 0x47, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x12,
	0xfa, 0x42, 0x0f, 0xaa, 0x01, 0x0c, 0x08, 0x01, 0x1a, 0x06, 0x08, 0x80, 0x80, 0x80, 0x80, 0x10,
	0x32, 0x00, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x3a, 0x20, 0x9a, 0xc5, 0x88,
	0x1e, 0x1b, 0x0a, 0x19, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x55, 0x72, 0x69, 0x42, 0x19, 0x0a,
	0x12, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x42, 0x80, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06,
	0x02, 0x10, 0x02, 0x0a, 0x22, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x42, 0x0c, 0x48, 0x74, 0x74, 0x70, 0x55, 0x72, 0x69,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67,
	0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x76, 0x33, 0x3b, 0x63, 0x6f, 0x72, 0x65, 0x76, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_envoy_config_core_v3_http_uri_proto_rawDescOnce sync.Once
	file_envoy_config_core_v3_http_uri_proto_rawDescData = file_envoy_config_core_v3_http_uri_proto_rawDesc
)

func file_envoy_config_core_v3_http_uri_proto_rawDescGZIP() []byte {
	file_envoy_config_core_v3_http_uri_proto_rawDescOnce.Do(func() {
		file_envoy_config_core_v3_http_uri_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_config_core_v3_http_uri_proto_rawDescData)
	})
	return file_envoy_config_core_v3_http_uri_proto_rawDescData
}

var file_envoy_config_core_v3_http_uri_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_config_core_v3_http_uri_proto_goTypes = []interface{}{
	(*HttpUri)(nil),           // 0: envoy.config.core.v3.HttpUri
	(*duration.Duration)(nil), // 1: google.protobuf.Duration
}
var file_envoy_config_core_v3_http_uri_proto_depIdxs = []int32{
	1, // 0: envoy.config.core.v3.HttpUri.timeout:type_name -> google.protobuf.Duration
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_envoy_config_core_v3_http_uri_proto_init() }
func file_envoy_config_core_v3_http_uri_proto_init() {
	if File_envoy_config_core_v3_http_uri_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_config_core_v3_http_uri_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HttpUri); i {
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
	file_envoy_config_core_v3_http_uri_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*HttpUri_Cluster)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_config_core_v3_http_uri_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_config_core_v3_http_uri_proto_goTypes,
		DependencyIndexes: file_envoy_config_core_v3_http_uri_proto_depIdxs,
		MessageInfos:      file_envoy_config_core_v3_http_uri_proto_msgTypes,
	}.Build()
	File_envoy_config_core_v3_http_uri_proto = out.File
	file_envoy_config_core_v3_http_uri_proto_rawDesc = nil
	file_envoy_config_core_v3_http_uri_proto_goTypes = nil
	file_envoy_config_core_v3_http_uri_proto_depIdxs = nil
}
