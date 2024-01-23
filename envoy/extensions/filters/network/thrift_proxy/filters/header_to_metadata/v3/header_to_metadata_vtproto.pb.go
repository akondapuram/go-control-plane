//go:build vtprotobuf
// +build vtprotobuf
// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// source: envoy/extensions/filters/network/thrift_proxy/filters/header_to_metadata/v3/header_to_metadata.proto

package header_to_metadatav3

import (
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

func (m *HeaderToMetadata_KeyValuePair) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *HeaderToMetadata_KeyValuePair) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *HeaderToMetadata_KeyValuePair) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if m.Encode != 0 {
		i = encodeVarint(dAtA, i, uint64(m.Encode))
		i--
		dAtA[i] = 0x30
	}
	if m.Type != 0 {
		i = encodeVarint(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x28
	}
	if msg, ok := m.ValueType.(*HeaderToMetadata_KeyValuePair_RegexValueRewrite); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if msg, ok := m.ValueType.(*HeaderToMetadata_KeyValuePair_Value); ok {
		size, err := msg.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarint(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.MetadataNamespace) > 0 {
		i -= len(m.MetadataNamespace)
		copy(dAtA[i:], m.MetadataNamespace)
		i = encodeVarint(dAtA, i, uint64(len(m.MetadataNamespace)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *HeaderToMetadata_KeyValuePair_Value) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *HeaderToMetadata_KeyValuePair_Value) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.Value)
	copy(dAtA[i:], m.Value)
	i = encodeVarint(dAtA, i, uint64(len(m.Value)))
	i--
	dAtA[i] = 0x1a
	return len(dAtA) - i, nil
}
func (m *HeaderToMetadata_KeyValuePair_RegexValueRewrite) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *HeaderToMetadata_KeyValuePair_RegexValueRewrite) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.RegexValueRewrite != nil {
		if vtmsg, ok := interface{}(m.RegexValueRewrite).(interface {
			MarshalToSizedBufferVTStrict([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.RegexValueRewrite)
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
	return len(dAtA) - i, nil
}
func (m *HeaderToMetadata_Rule) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *HeaderToMetadata_Rule) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *HeaderToMetadata_Rule) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if m.Remove {
		i--
		if m.Remove {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if m.OnMissing != nil {
		size, err := m.OnMissing.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x1a
	}
	if m.OnPresent != nil {
		size, err := m.OnPresent.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Header) > 0 {
		i -= len(m.Header)
		copy(dAtA[i:], m.Header)
		i = encodeVarint(dAtA, i, uint64(len(m.Header)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *HeaderToMetadata) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *HeaderToMetadata) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *HeaderToMetadata) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if len(m.RequestRules) > 0 {
		for iNdEx := len(m.RequestRules) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.RequestRules[iNdEx].MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0xa
		}
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
func (m *HeaderToMetadata_KeyValuePair) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MetadataNamespace)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	if vtmsg, ok := m.ValueType.(interface{ SizeVT() int }); ok {
		n += vtmsg.SizeVT()
	}
	if m.Type != 0 {
		n += 1 + sov(uint64(m.Type))
	}
	if m.Encode != 0 {
		n += 1 + sov(uint64(m.Encode))
	}
	n += len(m.unknownFields)
	return n
}

func (m *HeaderToMetadata_KeyValuePair_Value) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Value)
	n += 1 + l + sov(uint64(l))
	return n
}
func (m *HeaderToMetadata_KeyValuePair_RegexValueRewrite) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RegexValueRewrite != nil {
		if size, ok := interface{}(m.RegexValueRewrite).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.RegexValueRewrite)
		}
		n += 1 + l + sov(uint64(l))
	}
	return n
}
func (m *HeaderToMetadata_Rule) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Header)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	if m.OnPresent != nil {
		l = m.OnPresent.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	if m.OnMissing != nil {
		l = m.OnMissing.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	if m.Remove {
		n += 2
	}
	n += len(m.unknownFields)
	return n
}

func (m *HeaderToMetadata) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.RequestRules) > 0 {
		for _, e := range m.RequestRules {
			l = e.SizeVT()
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
