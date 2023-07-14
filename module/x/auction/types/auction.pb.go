// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: auction/v1/auction.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// AuctionStatus represents the status of an Auction.
// An Auction can be either in progress or finished.
type AuctionStatus int32

const (
	// The Auction status is unspecified. Probaly haven't started yet.
	AuctionStatus_AUCTION_STATUS_UNSPECIFIED AuctionStatus = 0
	// The Auction is still ongoing.
	AuctionStatus_AUCTION_STATUS_IN_PROGRESS AuctionStatus = 1
	// The Auction has ended.
	AuctionStatus_AUCTION_STATUS_FINISH AuctionStatus = 2
)

var AuctionStatus_name = map[int32]string{
	0: "AUCTION_STATUS_UNSPECIFIED",
	1: "AUCTION_STATUS_IN_PROGRESS",
	2: "AUCTION_STATUS_FINISH",
}

var AuctionStatus_value = map[string]int32{
	"AUCTION_STATUS_UNSPECIFIED": 0,
	"AUCTION_STATUS_IN_PROGRESS": 1,
	"AUCTION_STATUS_FINISH":      2,
}

func (x AuctionStatus) String() string {
	return proto.EnumName(AuctionStatus_name, int32(x))
}

func (AuctionStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_efe336ece9e41ddd, []int{0}
}

// AuctionPeriod represents a period of auctions.
// Each AuctionPeriod has a unique identifier and a starting block height.
// An AuctionPeriod can have multiple Auctions.
type AuctionPeriod struct {
	Id               uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	StartBlockHeight uint64 `protobuf:"varint,2,opt,name=start_block_height,json=startBlockHeight,proto3" json:"start_block_height,omitempty"`
}

func (m *AuctionPeriod) Reset()         { *m = AuctionPeriod{} }
func (m *AuctionPeriod) String() string { return proto.CompactTextString(m) }
func (*AuctionPeriod) ProtoMessage()    {}
func (*AuctionPeriod) Descriptor() ([]byte, []int) {
	return fileDescriptor_efe336ece9e41ddd, []int{0}
}
func (m *AuctionPeriod) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AuctionPeriod) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuctionPeriod.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AuctionPeriod) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuctionPeriod.Merge(m, src)
}
func (m *AuctionPeriod) XXX_Size() int {
	return m.Size()
}
func (m *AuctionPeriod) XXX_DiscardUnknown() {
	xxx_messageInfo_AuctionPeriod.DiscardUnknown(m)
}

var xxx_messageInfo_AuctionPeriod proto.InternalMessageInfo

func (m *AuctionPeriod) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *AuctionPeriod) GetStartBlockHeight() uint64 {
	if m != nil {
		return m.StartBlockHeight
	}
	return 0
}

// Auction represents a single auction.
// An Auction has a unique identifier relative to its Auction Period Id , an amount being auctioned, a status, and a highest bid.
type Auction struct {
	Id              uint64        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AuctionAmount   *types.Coin   `protobuf:"bytes,2,opt,name=auction_amount,json=auctionAmount,proto3" json:"auction_amount,omitempty"`
	Status          AuctionStatus `protobuf:"varint,3,opt,name=status,proto3,enum=auction.v1.AuctionStatus" json:"status,omitempty"`
	HighestBid      *Bid          `protobuf:"bytes,4,opt,name=highest_bid,json=highestBid,proto3" json:"highest_bid,omitempty"`
	AuctionPeriodId uint64        `protobuf:"varint,5,opt,name=auction_period_id,json=auctionPeriodId,proto3" json:"auction_period_id,omitempty"`
}

func (m *Auction) Reset()         { *m = Auction{} }
func (m *Auction) String() string { return proto.CompactTextString(m) }
func (*Auction) ProtoMessage()    {}
func (*Auction) Descriptor() ([]byte, []int) {
	return fileDescriptor_efe336ece9e41ddd, []int{1}
}
func (m *Auction) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Auction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Auction.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Auction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Auction.Merge(m, src)
}
func (m *Auction) XXX_Size() int {
	return m.Size()
}
func (m *Auction) XXX_DiscardUnknown() {
	xxx_messageInfo_Auction.DiscardUnknown(m)
}

var xxx_messageInfo_Auction proto.InternalMessageInfo

func (m *Auction) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Auction) GetAuctionAmount() *types.Coin {
	if m != nil {
		return m.AuctionAmount
	}
	return nil
}

func (m *Auction) GetStatus() AuctionStatus {
	if m != nil {
		return m.Status
	}
	return AuctionStatus_AUCTION_STATUS_UNSPECIFIED
}

func (m *Auction) GetHighestBid() *Bid {
	if m != nil {
		return m.HighestBid
	}
	return nil
}

func (m *Auction) GetAuctionPeriodId() uint64 {
	if m != nil {
		return m.AuctionPeriodId
	}
	return 0
}

// Bid represents a bid on an Auction.
// A Bid includes the identifier of the Auction, the amount of the bid, and the address of the bidder.
type Bid struct {
	AuctionId     uint64      `protobuf:"varint,1,opt,name=auction_id,json=auctionId,proto3" json:"auction_id,omitempty"`
	BidAmount     *types.Coin `protobuf:"bytes,2,opt,name=bid_amount,json=bidAmount,proto3" json:"bid_amount,omitempty"`
	BidderAddress string      `protobuf:"bytes,3,opt,name=bidder_address,json=bidderAddress,proto3" json:"bidder_address,omitempty"`
}

func (m *Bid) Reset()         { *m = Bid{} }
func (m *Bid) String() string { return proto.CompactTextString(m) }
func (*Bid) ProtoMessage()    {}
func (*Bid) Descriptor() ([]byte, []int) {
	return fileDescriptor_efe336ece9e41ddd, []int{2}
}
func (m *Bid) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Bid) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Bid.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Bid) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bid.Merge(m, src)
}
func (m *Bid) XXX_Size() int {
	return m.Size()
}
func (m *Bid) XXX_DiscardUnknown() {
	xxx_messageInfo_Bid.DiscardUnknown(m)
}

var xxx_messageInfo_Bid proto.InternalMessageInfo

func (m *Bid) GetAuctionId() uint64 {
	if m != nil {
		return m.AuctionId
	}
	return 0
}

func (m *Bid) GetBidAmount() *types.Coin {
	if m != nil {
		return m.BidAmount
	}
	return nil
}

func (m *Bid) GetBidderAddress() string {
	if m != nil {
		return m.BidderAddress
	}
	return ""
}

func init() {
	proto.RegisterEnum("auction.v1.AuctionStatus", AuctionStatus_name, AuctionStatus_value)
	proto.RegisterType((*AuctionPeriod)(nil), "auction.v1.AuctionPeriod")
	proto.RegisterType((*Auction)(nil), "auction.v1.Auction")
	proto.RegisterType((*Bid)(nil), "auction.v1.Bid")
}

func init() { proto.RegisterFile("auction/v1/auction.proto", fileDescriptor_efe336ece9e41ddd) }

var fileDescriptor_efe336ece9e41ddd = []byte{
	// 461 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x8d, 0xd3, 0x52, 0x94, 0xa9, 0x92, 0x86, 0x95, 0x90, 0x9c, 0x4a, 0x58, 0x55, 0x24, 0xa4,
	0xaa, 0x02, 0x9b, 0x94, 0x0b, 0x37, 0xb0, 0x43, 0xda, 0xfa, 0x40, 0x1a, 0xd9, 0xc9, 0x01, 0x2e,
	0x2b, 0xdb, 0xbb, 0x8a, 0x07, 0x9a, 0x6c, 0xe4, 0x5d, 0x47, 0xf4, 0x03, 0xb8, 0xf3, 0x59, 0x1c,
	0x7b, 0xe4, 0x88, 0x92, 0x2b, 0x1f, 0x81, 0xb2, 0xde, 0x94, 0xaa, 0x70, 0xe8, 0xcd, 0xfb, 0xde,
	0x9b, 0x99, 0xf7, 0xc6, 0x03, 0x76, 0x52, 0x66, 0x0a, 0xc5, 0xdc, 0x5b, 0xf6, 0x3c, 0xf3, 0xe9,
	0x2e, 0x0a, 0xa1, 0x04, 0x81, 0xed, 0x73, 0xd9, 0x3b, 0x74, 0x32, 0x21, 0x67, 0x42, 0x7a, 0x69,
	0x22, 0xb9, 0xb7, 0xec, 0xa5, 0x5c, 0x25, 0x3d, 0x2f, 0x13, 0x68, 0xb4, 0xdd, 0x0f, 0xd0, 0xf4,
	0x2b, 0xf5, 0x88, 0x17, 0x28, 0x18, 0x69, 0x41, 0x1d, 0x99, 0x6d, 0x1d, 0x59, 0xc7, 0xbb, 0x51,
	0x1d, 0x19, 0x79, 0x01, 0x44, 0xaa, 0xa4, 0x50, 0x34, 0xbd, 0x12, 0xd9, 0x17, 0x9a, 0x73, 0x9c,
	0xe6, 0xca, 0xae, 0x6b, 0xbe, 0xad, 0x99, 0x60, 0x43, 0x5c, 0x68, 0xbc, 0xfb, 0xdb, 0x82, 0xc7,
	0xa6, 0xdf, 0x3f, 0x9d, 0xde, 0x41, 0xcb, 0x18, 0xa3, 0xc9, 0x4c, 0x94, 0xf3, 0xaa, 0xcb, 0xfe,
	0x69, 0xc7, 0xad, 0x3c, 0xba, 0x1b, 0x8f, 0xae, 0xf1, 0xe8, 0xf6, 0x05, 0xce, 0xa3, 0xa6, 0x29,
	0xf0, 0xb5, 0x9e, 0xf4, 0x60, 0x4f, 0xaa, 0x44, 0x95, 0xd2, 0xde, 0x39, 0xb2, 0x8e, 0x5b, 0xa7,
	0x1d, 0xf7, 0x6f, 0x52, 0xd7, 0x8c, 0x8d, 0xb5, 0x20, 0x32, 0x42, 0xf2, 0x0a, 0xf6, 0x73, 0x9c,
	0xe6, 0x5c, 0x2a, 0x9a, 0x22, 0xb3, 0x77, 0xf5, 0xc4, 0x83, 0xbb, 0x75, 0x01, 0xb2, 0x08, 0x8c,
	0x26, 0x40, 0x46, 0x4e, 0xe0, 0xc9, 0xd6, 0xe6, 0x42, 0xaf, 0x84, 0x22, 0xb3, 0x1f, 0xe9, 0x14,
	0x07, 0xc9, 0xdd, 0x55, 0x85, 0xac, 0xfb, 0xcd, 0x82, 0x9d, 0x4d, 0xcd, 0x33, 0xd8, 0xee, 0x9c,
	0xde, 0x46, 0x6e, 0x18, 0x24, 0x64, 0xe4, 0x0d, 0x40, 0x8a, 0xec, 0xc1, 0xa9, 0x1b, 0x29, 0x32,
	0x93, 0xf8, 0x39, 0xb4, 0x52, 0x64, 0x8c, 0x17, 0x34, 0x61, 0xac, 0xe0, 0xb2, 0x4a, 0xde, 0x88,
	0x9a, 0x15, 0xea, 0x57, 0xe0, 0xc9, 0xe7, 0xdb, 0xbf, 0x58, 0xc5, 0x27, 0x0e, 0x1c, 0xfa, 0x93,
	0xfe, 0x38, 0xbc, 0x1c, 0xd2, 0x78, 0xec, 0x8f, 0x27, 0x31, 0x9d, 0x0c, 0xe3, 0xd1, 0xa0, 0x1f,
	0x9e, 0x85, 0x83, 0xf7, 0xed, 0xda, 0x7f, 0xf8, 0x70, 0x48, 0x47, 0xd1, 0xe5, 0x79, 0x34, 0x88,
	0xe3, 0xb6, 0x45, 0x3a, 0xf0, 0xf4, 0x1e, 0x7f, 0x16, 0x0e, 0xc3, 0xf8, 0xa2, 0x5d, 0x0f, 0x3e,
	0xfe, 0x58, 0x39, 0xd6, 0xcd, 0xca, 0xb1, 0x7e, 0xad, 0x1c, 0xeb, 0xfb, 0xda, 0xa9, 0xdd, 0xac,
	0x9d, 0xda, 0xcf, 0xb5, 0x53, 0xfb, 0xf4, 0x76, 0x8a, 0x2a, 0x2f, 0x53, 0x37, 0x13, 0x33, 0xef,
	0xbc, 0x48, 0x96, 0xa8, 0xae, 0x5f, 0x06, 0x05, 0xb2, 0x29, 0xbf, 0xff, 0x9c, 0x09, 0x56, 0x5e,
	0x71, 0xef, 0xeb, 0xf6, 0x70, 0x3d, 0x75, 0xbd, 0xe0, 0x32, 0xdd, 0xd3, 0x37, 0xf9, 0xfa, 0x4f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x02, 0xbc, 0xc4, 0x76, 0xdb, 0x02, 0x00, 0x00,
}

func (m *AuctionPeriod) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuctionPeriod) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuctionPeriod) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.StartBlockHeight != 0 {
		i = encodeVarintAuction(dAtA, i, uint64(m.StartBlockHeight))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintAuction(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Auction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Auction) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Auction) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AuctionPeriodId != 0 {
		i = encodeVarintAuction(dAtA, i, uint64(m.AuctionPeriodId))
		i--
		dAtA[i] = 0x28
	}
	if m.HighestBid != nil {
		{
			size, err := m.HighestBid.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAuction(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.Status != 0 {
		i = encodeVarintAuction(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x18
	}
	if m.AuctionAmount != nil {
		{
			size, err := m.AuctionAmount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAuction(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintAuction(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Bid) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Bid) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Bid) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BidderAddress) > 0 {
		i -= len(m.BidderAddress)
		copy(dAtA[i:], m.BidderAddress)
		i = encodeVarintAuction(dAtA, i, uint64(len(m.BidderAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if m.BidAmount != nil {
		{
			size, err := m.BidAmount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAuction(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.AuctionId != 0 {
		i = encodeVarintAuction(dAtA, i, uint64(m.AuctionId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintAuction(dAtA []byte, offset int, v uint64) int {
	offset -= sovAuction(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AuctionPeriod) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovAuction(uint64(m.Id))
	}
	if m.StartBlockHeight != 0 {
		n += 1 + sovAuction(uint64(m.StartBlockHeight))
	}
	return n
}

func (m *Auction) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovAuction(uint64(m.Id))
	}
	if m.AuctionAmount != nil {
		l = m.AuctionAmount.Size()
		n += 1 + l + sovAuction(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovAuction(uint64(m.Status))
	}
	if m.HighestBid != nil {
		l = m.HighestBid.Size()
		n += 1 + l + sovAuction(uint64(l))
	}
	if m.AuctionPeriodId != 0 {
		n += 1 + sovAuction(uint64(m.AuctionPeriodId))
	}
	return n
}

func (m *Bid) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AuctionId != 0 {
		n += 1 + sovAuction(uint64(m.AuctionId))
	}
	if m.BidAmount != nil {
		l = m.BidAmount.Size()
		n += 1 + l + sovAuction(uint64(l))
	}
	l = len(m.BidderAddress)
	if l > 0 {
		n += 1 + l + sovAuction(uint64(l))
	}
	return n
}

func sovAuction(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAuction(x uint64) (n int) {
	return sovAuction(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AuctionPeriod) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuction
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AuctionPeriod: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuctionPeriod: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartBlockHeight", wireType)
			}
			m.StartBlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartBlockHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAuction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuction
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Auction) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuction
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Auction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Auction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuctionAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AuctionAmount == nil {
				m.AuctionAmount = &types.Coin{}
			}
			if err := m.AuctionAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= AuctionStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HighestBid", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.HighestBid == nil {
				m.HighestBid = &Bid{}
			}
			if err := m.HighestBid.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuctionPeriodId", wireType)
			}
			m.AuctionPeriodId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AuctionPeriodId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAuction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuction
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Bid) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuction
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Bid: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Bid: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuctionId", wireType)
			}
			m.AuctionId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AuctionId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BidAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BidAmount == nil {
				m.BidAmount = &types.Coin{}
			}
			if err := m.BidAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BidderAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BidderAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuction
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipAuction(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAuction
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAuction
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAuction
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthAuction
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAuction
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAuction
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAuction        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAuction          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAuction = fmt.Errorf("proto: unexpected end of group")
)
