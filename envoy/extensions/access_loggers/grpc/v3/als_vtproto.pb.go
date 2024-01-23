//go:build vtprotobuf
// +build vtprotobuf
// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// source: envoy/extensions/access_loggers/grpc/v3/als.proto

package grpcv3

import (
	durationpb "github.com/planetscale/vtprotobuf/types/known/durationpb"
	wrapperspb "github.com/planetscale/vtprotobuf/types/known/wrapperspb"
	proto "google.golang.org/protobuf/proto"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	bits "math/bits"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func (m *HttpGrpcAccessLogConfig) MarshalVTStrict() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVTStrict(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HttpGrpcAccessLogConfig) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *HttpGrpcAccessLogConfig) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if len(m.AdditionalResponseTrailersToLog) > 0 {
		for iNdEx := len(m.AdditionalResponseTrailersToLog) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AdditionalResponseTrailersToLog[iNdEx])
			copy(dAtA[i:], m.AdditionalResponseTrailersToLog[iNdEx])
			i = encodeVarint(dAtA, i, uint64(len(m.AdditionalResponseTrailersToLog[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.AdditionalResponseHeadersToLog) > 0 {
		for iNdEx := len(m.AdditionalResponseHeadersToLog) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AdditionalResponseHeadersToLog[iNdEx])
			copy(dAtA[i:], m.AdditionalResponseHeadersToLog[iNdEx])
			i = encodeVarint(dAtA, i, uint64(len(m.AdditionalResponseHeadersToLog[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.AdditionalRequestHeadersToLog) > 0 {
		for iNdEx := len(m.AdditionalRequestHeadersToLog) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AdditionalRequestHeadersToLog[iNdEx])
			copy(dAtA[i:], m.AdditionalRequestHeadersToLog[iNdEx])
			i = encodeVarint(dAtA, i, uint64(len(m.AdditionalRequestHeadersToLog[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if m.CommonConfig != nil {
		size, err := m.CommonConfig.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TcpGrpcAccessLogConfig) MarshalVTStrict() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVTStrict(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TcpGrpcAccessLogConfig) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *TcpGrpcAccessLogConfig) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if m.CommonConfig != nil {
		size, err := m.CommonConfig.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CommonGrpcAccessLogConfig) MarshalVTStrict() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVTStrict(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CommonGrpcAccessLogConfig) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *CommonGrpcAccessLogConfig) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if len(m.CustomTags) > 0 {
		for iNdEx := len(m.CustomTags) - 1; iNdEx >= 0; iNdEx-- {
			if vtmsg, ok := interface{}(m.CustomTags[iNdEx]).(interface {
				MarshalToSizedBufferVTStrict([]byte) (int, error)
			}); ok {
				size, err := vtmsg.MarshalToSizedBufferVTStrict(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarint(dAtA, i, uint64(size))
			} else {
				encoded, err := proto.Marshal(m.CustomTags[iNdEx])
				if err != nil {
					return 0, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = encodeVarint(dAtA, i, uint64(len(encoded)))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if m.GrpcStreamRetryPolicy != nil {
		if vtmsg, ok := interface{}(m.GrpcStreamRetryPolicy).(interface {
			MarshalToSizedBufferVTStrict([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.GrpcStreamRetryPolicy)
			if err != nil {
				return 0, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = encodeVarint(dAtA, i, uint64(len(encoded)))
		}
		i--
		dAtA[i] = 0x3a
	}
	if m.TransportApiVersion != 0 {
		i = encodeVarint(dAtA, i, uint64(m.TransportApiVersion))
		i--
		dAtA[i] = 0x30
	}
	if len(m.FilterStateObjectsToLog) > 0 {
		for iNdEx := len(m.FilterStateObjectsToLog) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.FilterStateObjectsToLog[iNdEx])
			copy(dAtA[i:], m.FilterStateObjectsToLog[iNdEx])
			i = encodeVarint(dAtA, i, uint64(len(m.FilterStateObjectsToLog[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.BufferSizeBytes != nil {
		size, err := (*wrapperspb.UInt32Value)(m.BufferSizeBytes).MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x22
	}
	if m.BufferFlushInterval != nil {
		size, err := (*durationpb.Duration)(m.BufferFlushInterval).MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x1a
	}
	if m.GrpcService != nil {
		if vtmsg, ok := interface{}(m.GrpcService).(interface {
			MarshalToSizedBufferVTStrict([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.GrpcService)
			if err != nil {
				return 0, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = encodeVarint(dAtA, i, uint64(len(encoded)))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.LogName) > 0 {
		i -= len(m.LogName)
		copy(dAtA[i:], m.LogName)
		i = encodeVarint(dAtA, i, uint64(len(m.LogName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarint(dAtA []byte, offset int, v uint64) int {
	offset -= sov(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *HttpGrpcAccessLogConfig) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CommonConfig != nil {
		l = m.CommonConfig.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	if len(m.AdditionalRequestHeadersToLog) > 0 {
		for _, s := range m.AdditionalRequestHeadersToLog {
			l = len(s)
			n += 1 + l + sov(uint64(l))
		}
	}
	if len(m.AdditionalResponseHeadersToLog) > 0 {
		for _, s := range m.AdditionalResponseHeadersToLog {
			l = len(s)
			n += 1 + l + sov(uint64(l))
		}
	}
	if len(m.AdditionalResponseTrailersToLog) > 0 {
		for _, s := range m.AdditionalResponseTrailersToLog {
			l = len(s)
			n += 1 + l + sov(uint64(l))
		}
	}
	n += len(m.unknownFields)
	return n
}

func (m *TcpGrpcAccessLogConfig) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CommonConfig != nil {
		l = m.CommonConfig.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *CommonGrpcAccessLogConfig) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.LogName)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	if m.GrpcService != nil {
		if size, ok := interface{}(m.GrpcService).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.GrpcService)
		}
		n += 1 + l + sov(uint64(l))
	}
	if m.BufferFlushInterval != nil {
		l = (*durationpb.Duration)(m.BufferFlushInterval).SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	if m.BufferSizeBytes != nil {
		l = (*wrapperspb.UInt32Value)(m.BufferSizeBytes).SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	if len(m.FilterStateObjectsToLog) > 0 {
		for _, s := range m.FilterStateObjectsToLog {
			l = len(s)
			n += 1 + l + sov(uint64(l))
		}
	}
	if m.TransportApiVersion != 0 {
		n += 1 + sov(uint64(m.TransportApiVersion))
	}
	if m.GrpcStreamRetryPolicy != nil {
		if size, ok := interface{}(m.GrpcStreamRetryPolicy).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.GrpcStreamRetryPolicy)
		}
		n += 1 + l + sov(uint64(l))
	}
	if len(m.CustomTags) > 0 {
		for _, e := range m.CustomTags {
			if size, ok := interface{}(e).(interface {
				SizeVT() int
			}); ok {
				l = size.SizeVT()
			} else {
				l = proto.Size(e)
			}
			n += 1 + l + sov(uint64(l))
		}
	}
	n += len(m.unknownFields)
	return n
}

func sov(x uint64) (n int) {
	return (bits.Len64(x|1) + 6) / 7
}
func soz(x uint64) (n int) {
	return sov(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
