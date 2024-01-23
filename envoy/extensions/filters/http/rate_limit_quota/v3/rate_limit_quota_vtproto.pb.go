//go:build vtprotobuf
// +build vtprotobuf
// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// source: envoy/extensions/filters/http/rate_limit_quota/v3/rate_limit_quota.proto

package rate_limit_quotav3

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

func (m *RateLimitQuotaFilterConfig) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *RateLimitQuotaFilterConfig) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *RateLimitQuotaFilterConfig) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if len(m.RequestHeadersToAddWhenNotEnforced) > 0 {
		for iNdEx := len(m.RequestHeadersToAddWhenNotEnforced) - 1; iNdEx >= 0; iNdEx-- {
			if vtmsg, ok := interface{}(m.RequestHeadersToAddWhenNotEnforced[iNdEx]).(interface {
				MarshalToSizedBufferVTStrict([]byte) (int, error)
			}); ok {
				size, err := vtmsg.MarshalToSizedBufferVTStrict(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarint(dAtA, i, uint64(size))
			} else {
				encoded, err := proto.Marshal(m.RequestHeadersToAddWhenNotEnforced[iNdEx])
				if err != nil {
					return 0, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = encodeVarint(dAtA, i, uint64(len(encoded)))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if m.FilterEnforced != nil {
		if vtmsg, ok := interface{}(m.FilterEnforced).(interface {
			MarshalToSizedBufferVTStrict([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.FilterEnforced)
			if err != nil {
				return 0, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = encodeVarint(dAtA, i, uint64(len(encoded)))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.FilterEnabled != nil {
		if vtmsg, ok := interface{}(m.FilterEnabled).(interface {
			MarshalToSizedBufferVTStrict([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.FilterEnabled)
			if err != nil {
				return 0, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = encodeVarint(dAtA, i, uint64(len(encoded)))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.BucketMatchers != nil {
		if vtmsg, ok := interface{}(m.BucketMatchers).(interface {
			MarshalToSizedBufferVTStrict([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.BucketMatchers)
			if err != nil {
				return 0, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = encodeVarint(dAtA, i, uint64(len(encoded)))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Domain) > 0 {
		i -= len(m.Domain)
		copy(dAtA[i:], m.Domain)
		i = encodeVarint(dAtA, i, uint64(len(m.Domain)))
		i--
		dAtA[i] = 0x12
	}
	if m.RlqsServer != nil {
		if vtmsg, ok := interface{}(m.RlqsServer).(interface {
			MarshalToSizedBufferVTStrict([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.RlqsServer)
			if err != nil {
				return 0, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = encodeVarint(dAtA, i, uint64(len(encoded)))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RateLimitQuotaOverride) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *RateLimitQuotaOverride) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *RateLimitQuotaOverride) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if m.BucketMatchers != nil {
		if vtmsg, ok := interface{}(m.BucketMatchers).(interface {
			MarshalToSizedBufferVTStrict([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.BucketMatchers)
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
	if len(m.Domain) > 0 {
		i -= len(m.Domain)
		copy(dAtA[i:], m.Domain)
		i = encodeVarint(dAtA, i, uint64(len(m.Domain)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RateLimitQuotaBucketSettings_NoAssignmentBehavior) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *RateLimitQuotaBucketSettings_NoAssignmentBehavior) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *RateLimitQuotaBucketSettings_NoAssignmentBehavior) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if msg, ok := m.NoAssignmentBehavior.(*RateLimitQuotaBucketSettings_NoAssignmentBehavior_FallbackRateLimit); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	return len(dAtA) - i, nil
}

func (m *RateLimitQuotaBucketSettings_NoAssignmentBehavior_FallbackRateLimit) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *RateLimitQuotaBucketSettings_NoAssignmentBehavior_FallbackRateLimit) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.FallbackRateLimit != nil {
		if vtmsg, ok := interface{}(m.FallbackRateLimit).(interface {
			MarshalToSizedBufferVTStrict([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.FallbackRateLimit)
			if err != nil {
				return 0, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = encodeVarint(dAtA, i, uint64(len(encoded)))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *RateLimitQuotaBucketSettings_ExpiredAssignmentBehavior_ReuseLastAssignment) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *RateLimitQuotaBucketSettings_ExpiredAssignmentBehavior_ReuseLastAssignment) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *RateLimitQuotaBucketSettings_ExpiredAssignmentBehavior_ReuseLastAssignment) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	return len(dAtA) - i, nil
}

func (m *RateLimitQuotaBucketSettings_ExpiredAssignmentBehavior) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *RateLimitQuotaBucketSettings_ExpiredAssignmentBehavior) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *RateLimitQuotaBucketSettings_ExpiredAssignmentBehavior) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if msg, ok := m.ExpiredAssignmentBehavior.(*RateLimitQuotaBucketSettings_ExpiredAssignmentBehavior_ReuseLastAssignment_); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if msg, ok := m.ExpiredAssignmentBehavior.(*RateLimitQuotaBucketSettings_ExpiredAssignmentBehavior_FallbackRateLimit); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if m.ExpiredAssignmentBehaviorTimeout != nil {
		size, err := (*durationpb.Duration)(m.ExpiredAssignmentBehaviorTimeout).MarshalToSizedBufferVTStrict(dAtA[:i])
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

func (m *RateLimitQuotaBucketSettings_ExpiredAssignmentBehavior_FallbackRateLimit) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *RateLimitQuotaBucketSettings_ExpiredAssignmentBehavior_FallbackRateLimit) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.FallbackRateLimit != nil {
		if vtmsg, ok := interface{}(m.FallbackRateLimit).(interface {
			MarshalToSizedBufferVTStrict([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.FallbackRateLimit)
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
	return len(dAtA) - i, nil
}
func (m *RateLimitQuotaBucketSettings_ExpiredAssignmentBehavior_ReuseLastAssignment_) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *RateLimitQuotaBucketSettings_ExpiredAssignmentBehavior_ReuseLastAssignment_) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.ReuseLastAssignment != nil {
		size, err := m.ReuseLastAssignment.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func (m *RateLimitQuotaBucketSettings_DenyResponseSettings) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *RateLimitQuotaBucketSettings_DenyResponseSettings) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *RateLimitQuotaBucketSettings_DenyResponseSettings) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if len(m.ResponseHeadersToAdd) > 0 {
		for iNdEx := len(m.ResponseHeadersToAdd) - 1; iNdEx >= 0; iNdEx-- {
			if vtmsg, ok := interface{}(m.ResponseHeadersToAdd[iNdEx]).(interface {
				MarshalToSizedBufferVTStrict([]byte) (int, error)
			}); ok {
				size, err := vtmsg.MarshalToSizedBufferVTStrict(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarint(dAtA, i, uint64(size))
			} else {
				encoded, err := proto.Marshal(m.ResponseHeadersToAdd[iNdEx])
				if err != nil {
					return 0, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = encodeVarint(dAtA, i, uint64(len(encoded)))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.GrpcStatus != nil {
		if vtmsg, ok := interface{}(m.GrpcStatus).(interface {
			MarshalToSizedBufferVTStrict([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.GrpcStatus)
			if err != nil {
				return 0, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = encodeVarint(dAtA, i, uint64(len(encoded)))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.HttpBody != nil {
		size, err := (*wrapperspb.BytesValue)(m.HttpBody).MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x12
	}
	if m.HttpStatus != nil {
		if vtmsg, ok := interface{}(m.HttpStatus).(interface {
			MarshalToSizedBufferVTStrict([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.HttpStatus)
			if err != nil {
				return 0, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = encodeVarint(dAtA, i, uint64(len(encoded)))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RateLimitQuotaBucketSettings_BucketIdBuilder_ValueBuilder) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *RateLimitQuotaBucketSettings_BucketIdBuilder_ValueBuilder) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *RateLimitQuotaBucketSettings_BucketIdBuilder_ValueBuilder) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if msg, ok := m.ValueSpecifier.(*RateLimitQuotaBucketSettings_BucketIdBuilder_ValueBuilder_CustomValue); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if msg, ok := m.ValueSpecifier.(*RateLimitQuotaBucketSettings_BucketIdBuilder_ValueBuilder_StringValue); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	return len(dAtA) - i, nil
}

func (m *RateLimitQuotaBucketSettings_BucketIdBuilder_ValueBuilder_StringValue) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *RateLimitQuotaBucketSettings_BucketIdBuilder_ValueBuilder_StringValue) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.StringValue)
	copy(dAtA[i:], m.StringValue)
	i = encodeVarint(dAtA, i, uint64(len(m.StringValue)))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}
func (m *RateLimitQuotaBucketSettings_BucketIdBuilder_ValueBuilder_CustomValue) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *RateLimitQuotaBucketSettings_BucketIdBuilder_ValueBuilder_CustomValue) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.CustomValue != nil {
		if vtmsg, ok := interface{}(m.CustomValue).(interface {
			MarshalToSizedBufferVTStrict([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.CustomValue)
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
	return len(dAtA) - i, nil
}
func (m *RateLimitQuotaBucketSettings_BucketIdBuilder) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *RateLimitQuotaBucketSettings_BucketIdBuilder) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *RateLimitQuotaBucketSettings_BucketIdBuilder) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if len(m.BucketIdBuilder) > 0 {
		for k := range m.BucketIdBuilder {
			v := m.BucketIdBuilder[k]
			baseI := i
			size, err := v.MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarint(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarint(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *RateLimitQuotaBucketSettings) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *RateLimitQuotaBucketSettings) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *RateLimitQuotaBucketSettings) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if m.ExpiredAssignmentBehavior != nil {
		size, err := m.ExpiredAssignmentBehavior.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x2a
	}
	if m.NoAssignmentBehavior != nil {
		size, err := m.NoAssignmentBehavior.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x22
	}
	if m.DenyResponseSettings != nil {
		size, err := m.DenyResponseSettings.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x1a
	}
	if m.ReportingInterval != nil {
		size, err := (*durationpb.Duration)(m.ReportingInterval).MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x12
	}
	if m.BucketIdBuilder != nil {
		size, err := m.BucketIdBuilder.MarshalToSizedBufferVTStrict(dAtA[:i])
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
func (m *RateLimitQuotaFilterConfig) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RlqsServer != nil {
		if size, ok := interface{}(m.RlqsServer).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.RlqsServer)
		}
		n += 1 + l + sov(uint64(l))
	}
	l = len(m.Domain)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	if m.BucketMatchers != nil {
		if size, ok := interface{}(m.BucketMatchers).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.BucketMatchers)
		}
		n += 1 + l + sov(uint64(l))
	}
	if m.FilterEnabled != nil {
		if size, ok := interface{}(m.FilterEnabled).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.FilterEnabled)
		}
		n += 1 + l + sov(uint64(l))
	}
	if m.FilterEnforced != nil {
		if size, ok := interface{}(m.FilterEnforced).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.FilterEnforced)
		}
		n += 1 + l + sov(uint64(l))
	}
	if len(m.RequestHeadersToAddWhenNotEnforced) > 0 {
		for _, e := range m.RequestHeadersToAddWhenNotEnforced {
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

func (m *RateLimitQuotaOverride) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Domain)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	if m.BucketMatchers != nil {
		if size, ok := interface{}(m.BucketMatchers).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.BucketMatchers)
		}
		n += 1 + l + sov(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *RateLimitQuotaBucketSettings_NoAssignmentBehavior) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if vtmsg, ok := m.NoAssignmentBehavior.(interface{ SizeVT() int }); ok {
		n += vtmsg.SizeVT()
	}
	n += len(m.unknownFields)
	return n
}

func (m *RateLimitQuotaBucketSettings_NoAssignmentBehavior_FallbackRateLimit) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FallbackRateLimit != nil {
		if size, ok := interface{}(m.FallbackRateLimit).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.FallbackRateLimit)
		}
		n += 1 + l + sov(uint64(l))
	}
	return n
}
func (m *RateLimitQuotaBucketSettings_ExpiredAssignmentBehavior_ReuseLastAssignment) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += len(m.unknownFields)
	return n
}

func (m *RateLimitQuotaBucketSettings_ExpiredAssignmentBehavior) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ExpiredAssignmentBehaviorTimeout != nil {
		l = (*durationpb.Duration)(m.ExpiredAssignmentBehaviorTimeout).SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	if vtmsg, ok := m.ExpiredAssignmentBehavior.(interface{ SizeVT() int }); ok {
		n += vtmsg.SizeVT()
	}
	n += len(m.unknownFields)
	return n
}

func (m *RateLimitQuotaBucketSettings_ExpiredAssignmentBehavior_FallbackRateLimit) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FallbackRateLimit != nil {
		if size, ok := interface{}(m.FallbackRateLimit).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.FallbackRateLimit)
		}
		n += 1 + l + sov(uint64(l))
	}
	return n
}
func (m *RateLimitQuotaBucketSettings_ExpiredAssignmentBehavior_ReuseLastAssignment_) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ReuseLastAssignment != nil {
		l = m.ReuseLastAssignment.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	return n
}
func (m *RateLimitQuotaBucketSettings_DenyResponseSettings) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.HttpStatus != nil {
		if size, ok := interface{}(m.HttpStatus).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.HttpStatus)
		}
		n += 1 + l + sov(uint64(l))
	}
	if m.HttpBody != nil {
		l = (*wrapperspb.BytesValue)(m.HttpBody).SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	if m.GrpcStatus != nil {
		if size, ok := interface{}(m.GrpcStatus).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.GrpcStatus)
		}
		n += 1 + l + sov(uint64(l))
	}
	if len(m.ResponseHeadersToAdd) > 0 {
		for _, e := range m.ResponseHeadersToAdd {
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

func (m *RateLimitQuotaBucketSettings_BucketIdBuilder_ValueBuilder) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if vtmsg, ok := m.ValueSpecifier.(interface{ SizeVT() int }); ok {
		n += vtmsg.SizeVT()
	}
	n += len(m.unknownFields)
	return n
}

func (m *RateLimitQuotaBucketSettings_BucketIdBuilder_ValueBuilder_StringValue) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.StringValue)
	n += 1 + l + sov(uint64(l))
	return n
}
func (m *RateLimitQuotaBucketSettings_BucketIdBuilder_ValueBuilder_CustomValue) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CustomValue != nil {
		if size, ok := interface{}(m.CustomValue).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.CustomValue)
		}
		n += 1 + l + sov(uint64(l))
	}
	return n
}
func (m *RateLimitQuotaBucketSettings_BucketIdBuilder) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.BucketIdBuilder) > 0 {
		for k, v := range m.BucketIdBuilder {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.SizeVT()
			}
			l += 1 + sov(uint64(l))
			mapEntrySize := 1 + len(k) + sov(uint64(len(k))) + l
			n += mapEntrySize + 1 + sov(uint64(mapEntrySize))
		}
	}
	n += len(m.unknownFields)
	return n
}

func (m *RateLimitQuotaBucketSettings) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BucketIdBuilder != nil {
		l = m.BucketIdBuilder.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	if m.ReportingInterval != nil {
		l = (*durationpb.Duration)(m.ReportingInterval).SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	if m.DenyResponseSettings != nil {
		l = m.DenyResponseSettings.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	if m.NoAssignmentBehavior != nil {
		l = m.NoAssignmentBehavior.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	if m.ExpiredAssignmentBehavior != nil {
		l = m.ExpiredAssignmentBehavior.SizeVT()
		n += 1 + l + sov(uint64(l))
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
