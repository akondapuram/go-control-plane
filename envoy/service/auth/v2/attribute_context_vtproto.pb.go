//go:build vtprotobuf
// +build vtprotobuf
// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// source: envoy/service/auth/v2/attribute_context.proto

package authv2

import (
	timestamppb "github.com/planetscale/vtprotobuf/types/known/timestamppb"
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

func (m *AttributeContext_Peer) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *AttributeContext_Peer) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *AttributeContext_Peer) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if len(m.Certificate) > 0 {
		i -= len(m.Certificate)
		copy(dAtA[i:], m.Certificate)
		i = encodeVarint(dAtA, i, uint64(len(m.Certificate)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Principal) > 0 {
		i -= len(m.Principal)
		copy(dAtA[i:], m.Principal)
		i = encodeVarint(dAtA, i, uint64(len(m.Principal)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Labels) > 0 {
		for k := range m.Labels {
			v := m.Labels[k]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarint(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarint(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarint(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Service) > 0 {
		i -= len(m.Service)
		copy(dAtA[i:], m.Service)
		i = encodeVarint(dAtA, i, uint64(len(m.Service)))
		i--
		dAtA[i] = 0x12
	}
	if m.Address != nil {
		if vtmsg, ok := interface{}(m.Address).(interface {
			MarshalToSizedBufferVTStrict([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.Address)
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

func (m *AttributeContext_Request) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *AttributeContext_Request) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *AttributeContext_Request) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if m.Http != nil {
		size, err := m.Http.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x12
	}
	if m.Time != nil {
		size, err := (*timestamppb.Timestamp)(m.Time).MarshalToSizedBufferVTStrict(dAtA[:i])
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

func (m *AttributeContext_HttpRequest) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *AttributeContext_HttpRequest) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *AttributeContext_HttpRequest) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if len(m.Body) > 0 {
		i -= len(m.Body)
		copy(dAtA[i:], m.Body)
		i = encodeVarint(dAtA, i, uint64(len(m.Body)))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.Protocol) > 0 {
		i -= len(m.Protocol)
		copy(dAtA[i:], m.Protocol)
		i = encodeVarint(dAtA, i, uint64(len(m.Protocol)))
		i--
		dAtA[i] = 0x52
	}
	if m.Size != 0 {
		i = encodeVarint(dAtA, i, uint64(m.Size))
		i--
		dAtA[i] = 0x48
	}
	if len(m.Fragment) > 0 {
		i -= len(m.Fragment)
		copy(dAtA[i:], m.Fragment)
		i = encodeVarint(dAtA, i, uint64(len(m.Fragment)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Query) > 0 {
		i -= len(m.Query)
		copy(dAtA[i:], m.Query)
		i = encodeVarint(dAtA, i, uint64(len(m.Query)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Scheme) > 0 {
		i -= len(m.Scheme)
		copy(dAtA[i:], m.Scheme)
		i = encodeVarint(dAtA, i, uint64(len(m.Scheme)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Host) > 0 {
		i -= len(m.Host)
		copy(dAtA[i:], m.Host)
		i = encodeVarint(dAtA, i, uint64(len(m.Host)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Path) > 0 {
		i -= len(m.Path)
		copy(dAtA[i:], m.Path)
		i = encodeVarint(dAtA, i, uint64(len(m.Path)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Headers) > 0 {
		for k := range m.Headers {
			v := m.Headers[k]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarint(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarint(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarint(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Method) > 0 {
		i -= len(m.Method)
		copy(dAtA[i:], m.Method)
		i = encodeVarint(dAtA, i, uint64(len(m.Method)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarint(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AttributeContext) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *AttributeContext) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *AttributeContext) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if m.MetadataContext != nil {
		if vtmsg, ok := interface{}(m.MetadataContext).(interface {
			MarshalToSizedBufferVTStrict([]byte) (int, error)
		}); ok {
			size, err := vtmsg.MarshalToSizedBufferVTStrict(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
		} else {
			encoded, err := proto.Marshal(m.MetadataContext)
			if err != nil {
				return 0, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = encodeVarint(dAtA, i, uint64(len(encoded)))
		}
		i--
		dAtA[i] = 0x5a
	}
	if len(m.ContextExtensions) > 0 {
		for k := range m.ContextExtensions {
			v := m.ContextExtensions[k]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarint(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarint(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarint(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x52
		}
	}
	if m.Request != nil {
		size, err := m.Request.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x22
	}
	if m.Destination != nil {
		size, err := m.Destination.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x12
	}
	if m.Source != nil {
		size, err := m.Source.MarshalToSizedBufferVTStrict(dAtA[:i])
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
func (m *AttributeContext_Peer) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Address != nil {
		if size, ok := interface{}(m.Address).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.Address)
		}
		n += 1 + l + sov(uint64(l))
	}
	l = len(m.Service)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	if len(m.Labels) > 0 {
		for k, v := range m.Labels {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sov(uint64(len(k))) + 1 + len(v) + sov(uint64(len(v)))
			n += mapEntrySize + 1 + sov(uint64(mapEntrySize))
		}
	}
	l = len(m.Principal)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	l = len(m.Certificate)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *AttributeContext_Request) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Time != nil {
		l = (*timestamppb.Timestamp)(m.Time).SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	if m.Http != nil {
		l = m.Http.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *AttributeContext_HttpRequest) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	l = len(m.Method)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	if len(m.Headers) > 0 {
		for k, v := range m.Headers {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sov(uint64(len(k))) + 1 + len(v) + sov(uint64(len(v)))
			n += mapEntrySize + 1 + sov(uint64(mapEntrySize))
		}
	}
	l = len(m.Path)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	l = len(m.Host)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	l = len(m.Scheme)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	l = len(m.Query)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	l = len(m.Fragment)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	if m.Size != 0 {
		n += 1 + sov(uint64(m.Size))
	}
	l = len(m.Protocol)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	l = len(m.Body)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *AttributeContext) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Source != nil {
		l = m.Source.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	if m.Destination != nil {
		l = m.Destination.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	if m.Request != nil {
		l = m.Request.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	if len(m.ContextExtensions) > 0 {
		for k, v := range m.ContextExtensions {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sov(uint64(len(k))) + 1 + len(v) + sov(uint64(len(v)))
			n += mapEntrySize + 1 + sov(uint64(mapEntrySize))
		}
	}
	if m.MetadataContext != nil {
		if size, ok := interface{}(m.MetadataContext).(interface {
			SizeVT() int
		}); ok {
			l = size.SizeVT()
		} else {
			l = proto.Size(m.MetadataContext)
		}
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
