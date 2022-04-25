// Code generated by protoc-gen-go. DO NOT EDIT.
// source: game_s2c.proto

package msg

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

//所有的服务端->客户端的协议
type S2CMsgType int32

const (
	S2CMsgType_ROOM_INFO_RES   S2CMsgType = 0
	S2CMsgType_START_MATCH_RES S2CMsgType = 1
	S2CMsgType_START_GAME      S2CMsgType = 2
	S2CMsgType_END_GAME        S2CMsgType = 3
	S2CMsgType_SETTLE_CARDS    S2CMsgType = 4
)

var S2CMsgType_name = map[int32]string{
	0: "ROOM_INFO_RES",
	1: "START_MATCH_RES",
	2: "START_GAME",
	3: "END_GAME",
	4: "SETTLE_CARDS",
}

var S2CMsgType_value = map[string]int32{
	"ROOM_INFO_RES":   0,
	"START_MATCH_RES": 1,
	"START_GAME":      2,
	"END_GAME":        3,
	"SETTLE_CARDS":    4,
}

func (x S2CMsgType) String() string {
	return proto.EnumName(S2CMsgType_name, int32(x))
}

func (S2CMsgType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7a66076066de807d, []int{0}
}

//房间信息，断线重连也用该结构体
type S2CRoomInfo struct {
	Bottom               int64                  `protobuf:"varint,1,opt,name=bottom,proto3" json:"bottom"`
	TableStatus          int32                  `protobuf:"varint,2,opt,name=tableStatus,proto3" json:"tableStatus"`
	UserArr              []*S2CUserInfo         `protobuf:"bytes,3,rep,name=userArr,proto3" json:"userArr"`
	SelfCards            *S2CUserEndInfo        `protobuf:"bytes,4,opt,name=selfCards,proto3" json:"selfCards"`
	Ticker               int32                  `protobuf:"varint,5,opt,name=ticker,proto3" json:"ticker"`
	SpareArr             []*S2CCardsAndCardType `protobuf:"bytes,6,rep,name=spareArr,proto3" json:"spareArr"`
	SpecialType          int32                  `protobuf:"varint,7,opt,name=specialType,proto3" json:"specialType"`
	RoomId               int64                  `protobuf:"varint,8,opt,name=roomId,proto3" json:"roomId"`
	MinLimit             int64                  `protobuf:"varint,9,opt,name=MinLimit,proto3" json:"MinLimit"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *S2CRoomInfo) Reset()         { *m = S2CRoomInfo{} }
func (m *S2CRoomInfo) String() string { return proto.CompactTextString(m) }
func (*S2CRoomInfo) ProtoMessage()    {}
func (*S2CRoomInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a66076066de807d, []int{0}
}

func (m *S2CRoomInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S2CRoomInfo.Unmarshal(m, b)
}
func (m *S2CRoomInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S2CRoomInfo.Marshal(b, m, deterministic)
}
func (m *S2CRoomInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2CRoomInfo.Merge(m, src)
}
func (m *S2CRoomInfo) XXX_Size() int {
	return xxx_messageInfo_S2CRoomInfo.Size(m)
}
func (m *S2CRoomInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_S2CRoomInfo.DiscardUnknown(m)
}

var xxx_messageInfo_S2CRoomInfo proto.InternalMessageInfo

func (m *S2CRoomInfo) GetBottom() int64 {
	if m != nil {
		return m.Bottom
	}
	return 0
}

func (m *S2CRoomInfo) GetTableStatus() int32 {
	if m != nil {
		return m.TableStatus
	}
	return 0
}

func (m *S2CRoomInfo) GetUserArr() []*S2CUserInfo {
	if m != nil {
		return m.UserArr
	}
	return nil
}

func (m *S2CRoomInfo) GetSelfCards() *S2CUserEndInfo {
	if m != nil {
		return m.SelfCards
	}
	return nil
}

func (m *S2CRoomInfo) GetTicker() int32 {
	if m != nil {
		return m.Ticker
	}
	return 0
}

func (m *S2CRoomInfo) GetSpareArr() []*S2CCardsAndCardType {
	if m != nil {
		return m.SpareArr
	}
	return nil
}

func (m *S2CRoomInfo) GetSpecialType() int32 {
	if m != nil {
		return m.SpecialType
	}
	return 0
}

func (m *S2CRoomInfo) GetRoomId() int64 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

func (m *S2CRoomInfo) GetMinLimit() int64 {
	if m != nil {
		return m.MinLimit
	}
	return 0
}

//玩家确定牌型
type S2CSettleCards struct {
	Uid                  int64    `protobuf:"varint,1,opt,name=uid,proto3" json:"uid"`
	ChairId              int32    `protobuf:"varint,2,opt,name=chairId,proto3" json:"chairId"`
	SpecialType          int32    `protobuf:"varint,3,opt,name=specialType,proto3" json:"specialType"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *S2CSettleCards) Reset()         { *m = S2CSettleCards{} }
func (m *S2CSettleCards) String() string { return proto.CompactTextString(m) }
func (*S2CSettleCards) ProtoMessage()    {}
func (*S2CSettleCards) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a66076066de807d, []int{1}
}

func (m *S2CSettleCards) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S2CSettleCards.Unmarshal(m, b)
}
func (m *S2CSettleCards) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S2CSettleCards.Marshal(b, m, deterministic)
}
func (m *S2CSettleCards) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2CSettleCards.Merge(m, src)
}
func (m *S2CSettleCards) XXX_Size() int {
	return xxx_messageInfo_S2CSettleCards.Size(m)
}
func (m *S2CSettleCards) XXX_DiscardUnknown() {
	xxx_messageInfo_S2CSettleCards.DiscardUnknown(m)
}

var xxx_messageInfo_S2CSettleCards proto.InternalMessageInfo

func (m *S2CSettleCards) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *S2CSettleCards) GetChairId() int32 {
	if m != nil {
		return m.ChairId
	}
	return 0
}

func (m *S2CSettleCards) GetSpecialType() int32 {
	if m != nil {
		return m.SpecialType
	}
	return 0
}

//开始比赛
type S2CStartGame struct {
	SpecialType          int32                  `protobuf:"varint,1,opt,name=specialType,proto3" json:"specialType"`
	UserArr              []*S2CUserInfo         `protobuf:"bytes,2,rep,name=userArr,proto3" json:"userArr"`
	Ticker               int32                  `protobuf:"varint,3,opt,name=ticker,proto3" json:"ticker"`
	SpareArr             []*S2CCardsAndCardType `protobuf:"bytes,4,rep,name=spareArr,proto3" json:"spareArr"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *S2CStartGame) Reset()         { *m = S2CStartGame{} }
func (m *S2CStartGame) String() string { return proto.CompactTextString(m) }
func (*S2CStartGame) ProtoMessage()    {}
func (*S2CStartGame) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a66076066de807d, []int{2}
}

func (m *S2CStartGame) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S2CStartGame.Unmarshal(m, b)
}
func (m *S2CStartGame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S2CStartGame.Marshal(b, m, deterministic)
}
func (m *S2CStartGame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2CStartGame.Merge(m, src)
}
func (m *S2CStartGame) XXX_Size() int {
	return xxx_messageInfo_S2CStartGame.Size(m)
}
func (m *S2CStartGame) XXX_DiscardUnknown() {
	xxx_messageInfo_S2CStartGame.DiscardUnknown(m)
}

var xxx_messageInfo_S2CStartGame proto.InternalMessageInfo

func (m *S2CStartGame) GetSpecialType() int32 {
	if m != nil {
		return m.SpecialType
	}
	return 0
}

func (m *S2CStartGame) GetUserArr() []*S2CUserInfo {
	if m != nil {
		return m.UserArr
	}
	return nil
}

func (m *S2CStartGame) GetTicker() int32 {
	if m != nil {
		return m.Ticker
	}
	return 0
}

func (m *S2CStartGame) GetSpareArr() []*S2CCardsAndCardType {
	if m != nil {
		return m.SpareArr
	}
	return nil
}

//三墩牌型和牌
type S2CCardsAndCardType struct {
	HeadCards            []byte   `protobuf:"bytes,1,opt,name=headCards,proto3" json:"headCards"`
	MidCards             []byte   `protobuf:"bytes,2,opt,name=midCards,proto3" json:"midCards"`
	TailCards            []byte   `protobuf:"bytes,3,opt,name=tailCards,proto3" json:"tailCards"`
	HeadType             int32    `protobuf:"varint,4,opt,name=headType,proto3" json:"headType"`
	MidType              int32    `protobuf:"varint,5,opt,name=midType,proto3" json:"midType"`
	TailType             int32    `protobuf:"varint,6,opt,name=tailType,proto3" json:"tailType"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *S2CCardsAndCardType) Reset()         { *m = S2CCardsAndCardType{} }
func (m *S2CCardsAndCardType) String() string { return proto.CompactTextString(m) }
func (*S2CCardsAndCardType) ProtoMessage()    {}
func (*S2CCardsAndCardType) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a66076066de807d, []int{3}
}

func (m *S2CCardsAndCardType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S2CCardsAndCardType.Unmarshal(m, b)
}
func (m *S2CCardsAndCardType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S2CCardsAndCardType.Marshal(b, m, deterministic)
}
func (m *S2CCardsAndCardType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2CCardsAndCardType.Merge(m, src)
}
func (m *S2CCardsAndCardType) XXX_Size() int {
	return xxx_messageInfo_S2CCardsAndCardType.Size(m)
}
func (m *S2CCardsAndCardType) XXX_DiscardUnknown() {
	xxx_messageInfo_S2CCardsAndCardType.DiscardUnknown(m)
}

var xxx_messageInfo_S2CCardsAndCardType proto.InternalMessageInfo

func (m *S2CCardsAndCardType) GetHeadCards() []byte {
	if m != nil {
		return m.HeadCards
	}
	return nil
}

func (m *S2CCardsAndCardType) GetMidCards() []byte {
	if m != nil {
		return m.MidCards
	}
	return nil
}

func (m *S2CCardsAndCardType) GetTailCards() []byte {
	if m != nil {
		return m.TailCards
	}
	return nil
}

func (m *S2CCardsAndCardType) GetHeadType() int32 {
	if m != nil {
		return m.HeadType
	}
	return 0
}

func (m *S2CCardsAndCardType) GetMidType() int32 {
	if m != nil {
		return m.MidType
	}
	return 0
}

func (m *S2CCardsAndCardType) GetTailType() int32 {
	if m != nil {
		return m.TailType
	}
	return 0
}

//结束比赛
type S2CEndGame struct {
	HitArr               []*S2CHitRob      `protobuf:"bytes,1,rep,name=hitArr,proto3" json:"hitArr"`
	HomeRunUid           int64             `protobuf:"varint,2,opt,name=homeRunUid,proto3" json:"homeRunUid"`
	UserArr              []*S2CUserEndInfo `protobuf:"bytes,3,rep,name=userArr,proto3" json:"userArr"`
	HomeRunChairId       int32             `protobuf:"varint,4,opt,name=homeRunChairId,proto3" json:"homeRunChairId"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *S2CEndGame) Reset()         { *m = S2CEndGame{} }
func (m *S2CEndGame) String() string { return proto.CompactTextString(m) }
func (*S2CEndGame) ProtoMessage()    {}
func (*S2CEndGame) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a66076066de807d, []int{4}
}

func (m *S2CEndGame) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S2CEndGame.Unmarshal(m, b)
}
func (m *S2CEndGame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S2CEndGame.Marshal(b, m, deterministic)
}
func (m *S2CEndGame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2CEndGame.Merge(m, src)
}
func (m *S2CEndGame) XXX_Size() int {
	return xxx_messageInfo_S2CEndGame.Size(m)
}
func (m *S2CEndGame) XXX_DiscardUnknown() {
	xxx_messageInfo_S2CEndGame.DiscardUnknown(m)
}

var xxx_messageInfo_S2CEndGame proto.InternalMessageInfo

func (m *S2CEndGame) GetHitArr() []*S2CHitRob {
	if m != nil {
		return m.HitArr
	}
	return nil
}

func (m *S2CEndGame) GetHomeRunUid() int64 {
	if m != nil {
		return m.HomeRunUid
	}
	return 0
}

func (m *S2CEndGame) GetUserArr() []*S2CUserEndInfo {
	if m != nil {
		return m.UserArr
	}
	return nil
}

func (m *S2CEndGame) GetHomeRunChairId() int32 {
	if m != nil {
		return m.HomeRunChairId
	}
	return 0
}

//打抢
type S2CHitRob struct {
	HitUid               int64    `protobuf:"varint,1,opt,name=hitUid,proto3" json:"hitUid"`
	BeHitUid             int64    `protobuf:"varint,2,opt,name=beHitUid,proto3" json:"beHitUid"`
	HitScore             int64    `protobuf:"varint,3,opt,name=hitScore,proto3" json:"hitScore"`
	HitChairId           int32    `protobuf:"varint,4,opt,name=hitChairId,proto3" json:"hitChairId"`
	BeHitChairId         int32    `protobuf:"varint,5,opt,name=beHitChairId,proto3" json:"beHitChairId"`
	HitHeadScore         int64    `protobuf:"varint,6,opt,name=hitHeadScore,proto3" json:"hitHeadScore"`
	HitMidScore          int64    `protobuf:"varint,7,opt,name=hitMidScore,proto3" json:"hitMidScore"`
	HitTailScore         int64    `protobuf:"varint,8,opt,name=hitTailScore,proto3" json:"hitTailScore"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *S2CHitRob) Reset()         { *m = S2CHitRob{} }
func (m *S2CHitRob) String() string { return proto.CompactTextString(m) }
func (*S2CHitRob) ProtoMessage()    {}
func (*S2CHitRob) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a66076066de807d, []int{5}
}

func (m *S2CHitRob) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S2CHitRob.Unmarshal(m, b)
}
func (m *S2CHitRob) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S2CHitRob.Marshal(b, m, deterministic)
}
func (m *S2CHitRob) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2CHitRob.Merge(m, src)
}
func (m *S2CHitRob) XXX_Size() int {
	return xxx_messageInfo_S2CHitRob.Size(m)
}
func (m *S2CHitRob) XXX_DiscardUnknown() {
	xxx_messageInfo_S2CHitRob.DiscardUnknown(m)
}

var xxx_messageInfo_S2CHitRob proto.InternalMessageInfo

func (m *S2CHitRob) GetHitUid() int64 {
	if m != nil {
		return m.HitUid
	}
	return 0
}

func (m *S2CHitRob) GetBeHitUid() int64 {
	if m != nil {
		return m.BeHitUid
	}
	return 0
}

func (m *S2CHitRob) GetHitScore() int64 {
	if m != nil {
		return m.HitScore
	}
	return 0
}

func (m *S2CHitRob) GetHitChairId() int32 {
	if m != nil {
		return m.HitChairId
	}
	return 0
}

func (m *S2CHitRob) GetBeHitChairId() int32 {
	if m != nil {
		return m.BeHitChairId
	}
	return 0
}

func (m *S2CHitRob) GetHitHeadScore() int64 {
	if m != nil {
		return m.HitHeadScore
	}
	return 0
}

func (m *S2CHitRob) GetHitMidScore() int64 {
	if m != nil {
		return m.HitMidScore
	}
	return 0
}

func (m *S2CHitRob) GetHitTailScore() int64 {
	if m != nil {
		return m.HitTailScore
	}
	return 0
}

//玩家结束的信息
type S2CUserEndInfo struct {
	Uid                  int64    `protobuf:"varint,1,opt,name=uid,proto3" json:"uid"`
	HeadCards            []byte   `protobuf:"bytes,2,opt,name=headCards,proto3" json:"headCards"`
	MidCards             []byte   `protobuf:"bytes,3,opt,name=midCards,proto3" json:"midCards"`
	TailCards            []byte   `protobuf:"bytes,4,opt,name=tailCards,proto3" json:"tailCards"`
	HeadType             int32    `protobuf:"varint,5,opt,name=headType,proto3" json:"headType"`
	MidType              int32    `protobuf:"varint,6,opt,name=midType,proto3" json:"midType"`
	TailType             int32    `protobuf:"varint,7,opt,name=tailType,proto3" json:"tailType"`
	SpecialType          int32    `protobuf:"varint,8,opt,name=specialType,proto3" json:"specialType"`
	TotalWin             int64    `protobuf:"varint,11,opt,name=totalWin,proto3" json:"totalWin"`
	HeadWin              int64    `protobuf:"varint,12,opt,name=headWin,proto3" json:"headWin"`
	MidWin               int64    `protobuf:"varint,13,opt,name=midWin,proto3" json:"midWin"`
	TailWin              int64    `protobuf:"varint,14,opt,name=tailWin,proto3" json:"tailWin"`
	HeadPlus             int64    `protobuf:"varint,15,opt,name=headPlus,proto3" json:"headPlus"`
	MidPlus              int64    `protobuf:"varint,16,opt,name=midPlus,proto3" json:"midPlus"`
	TailPlus             int64    `protobuf:"varint,17,opt,name=tailPlus,proto3" json:"tailPlus"`
	TotalPlus            int64    `protobuf:"varint,18,opt,name=totalPlus,proto3" json:"totalPlus"`
	WinScore             int64    `protobuf:"varint,19,opt,name=winScore,proto3" json:"winScore"`
	ChairId              int32    `protobuf:"varint,20,opt,name=chairId,proto3" json:"chairId"`
	HomeRunHeadPlus      int64    `protobuf:"varint,21,opt,name=homeRunHeadPlus,proto3" json:"homeRunHeadPlus"`
	HomeRunMidPlus       int64    `protobuf:"varint,22,opt,name=homeRunMidPlus,proto3" json:"homeRunMidPlus"`
	HomeRunTailPlus      int64    `protobuf:"varint,23,opt,name=homeRunTailPlus,proto3" json:"homeRunTailPlus"`
	HomeRunTotalPlus     int64    `protobuf:"varint,24,opt,name=homeRunTotalPlus,proto3" json:"homeRunTotalPlus"`
	FinalScore           int64    `protobuf:"varint,25,opt,name=finalScore,proto3" json:"finalScore"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *S2CUserEndInfo) Reset()         { *m = S2CUserEndInfo{} }
func (m *S2CUserEndInfo) String() string { return proto.CompactTextString(m) }
func (*S2CUserEndInfo) ProtoMessage()    {}
func (*S2CUserEndInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a66076066de807d, []int{6}
}

func (m *S2CUserEndInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S2CUserEndInfo.Unmarshal(m, b)
}
func (m *S2CUserEndInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S2CUserEndInfo.Marshal(b, m, deterministic)
}
func (m *S2CUserEndInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2CUserEndInfo.Merge(m, src)
}
func (m *S2CUserEndInfo) XXX_Size() int {
	return xxx_messageInfo_S2CUserEndInfo.Size(m)
}
func (m *S2CUserEndInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_S2CUserEndInfo.DiscardUnknown(m)
}

var xxx_messageInfo_S2CUserEndInfo proto.InternalMessageInfo

func (m *S2CUserEndInfo) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *S2CUserEndInfo) GetHeadCards() []byte {
	if m != nil {
		return m.HeadCards
	}
	return nil
}

func (m *S2CUserEndInfo) GetMidCards() []byte {
	if m != nil {
		return m.MidCards
	}
	return nil
}

func (m *S2CUserEndInfo) GetTailCards() []byte {
	if m != nil {
		return m.TailCards
	}
	return nil
}

func (m *S2CUserEndInfo) GetHeadType() int32 {
	if m != nil {
		return m.HeadType
	}
	return 0
}

func (m *S2CUserEndInfo) GetMidType() int32 {
	if m != nil {
		return m.MidType
	}
	return 0
}

func (m *S2CUserEndInfo) GetTailType() int32 {
	if m != nil {
		return m.TailType
	}
	return 0
}

func (m *S2CUserEndInfo) GetSpecialType() int32 {
	if m != nil {
		return m.SpecialType
	}
	return 0
}

func (m *S2CUserEndInfo) GetTotalWin() int64 {
	if m != nil {
		return m.TotalWin
	}
	return 0
}

func (m *S2CUserEndInfo) GetHeadWin() int64 {
	if m != nil {
		return m.HeadWin
	}
	return 0
}

func (m *S2CUserEndInfo) GetMidWin() int64 {
	if m != nil {
		return m.MidWin
	}
	return 0
}

func (m *S2CUserEndInfo) GetTailWin() int64 {
	if m != nil {
		return m.TailWin
	}
	return 0
}

func (m *S2CUserEndInfo) GetHeadPlus() int64 {
	if m != nil {
		return m.HeadPlus
	}
	return 0
}

func (m *S2CUserEndInfo) GetMidPlus() int64 {
	if m != nil {
		return m.MidPlus
	}
	return 0
}

func (m *S2CUserEndInfo) GetTailPlus() int64 {
	if m != nil {
		return m.TailPlus
	}
	return 0
}

func (m *S2CUserEndInfo) GetTotalPlus() int64 {
	if m != nil {
		return m.TotalPlus
	}
	return 0
}

func (m *S2CUserEndInfo) GetWinScore() int64 {
	if m != nil {
		return m.WinScore
	}
	return 0
}

func (m *S2CUserEndInfo) GetChairId() int32 {
	if m != nil {
		return m.ChairId
	}
	return 0
}

func (m *S2CUserEndInfo) GetHomeRunHeadPlus() int64 {
	if m != nil {
		return m.HomeRunHeadPlus
	}
	return 0
}

func (m *S2CUserEndInfo) GetHomeRunMidPlus() int64 {
	if m != nil {
		return m.HomeRunMidPlus
	}
	return 0
}

func (m *S2CUserEndInfo) GetHomeRunTailPlus() int64 {
	if m != nil {
		return m.HomeRunTailPlus
	}
	return 0
}

func (m *S2CUserEndInfo) GetHomeRunTotalPlus() int64 {
	if m != nil {
		return m.HomeRunTotalPlus
	}
	return 0
}

func (m *S2CUserEndInfo) GetFinalScore() int64 {
	if m != nil {
		return m.FinalScore
	}
	return 0
}

//用户基本信息
type S2CUserInfo struct {
	Uid                  int64    `protobuf:"varint,1,opt,name=uid,proto3" json:"uid"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	Head                 string   `protobuf:"bytes,3,opt,name=head,proto3" json:"head"`
	Amount               int64    `protobuf:"varint,4,opt,name=amount,proto3" json:"amount"`
	ChairId              int32    `protobuf:"varint,5,opt,name=chairId,proto3" json:"chairId"`
	Status               int32    `protobuf:"varint,6,opt,name=status,proto3" json:"status"`
	IsSettleCards        bool     `protobuf:"varint,7,opt,name=isSettleCards,proto3" json:"isSettleCards"`
	SpecialType          int32    `protobuf:"varint,8,opt,name=specialType,proto3" json:"specialType"`
	Ip                   string   `protobuf:"bytes,9,opt,name=ip,proto3" json:"ip"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *S2CUserInfo) Reset()         { *m = S2CUserInfo{} }
func (m *S2CUserInfo) String() string { return proto.CompactTextString(m) }
func (*S2CUserInfo) ProtoMessage()    {}
func (*S2CUserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a66076066de807d, []int{7}
}

func (m *S2CUserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S2CUserInfo.Unmarshal(m, b)
}
func (m *S2CUserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S2CUserInfo.Marshal(b, m, deterministic)
}
func (m *S2CUserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2CUserInfo.Merge(m, src)
}
func (m *S2CUserInfo) XXX_Size() int {
	return xxx_messageInfo_S2CUserInfo.Size(m)
}
func (m *S2CUserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_S2CUserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_S2CUserInfo proto.InternalMessageInfo

func (m *S2CUserInfo) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *S2CUserInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *S2CUserInfo) GetHead() string {
	if m != nil {
		return m.Head
	}
	return ""
}

func (m *S2CUserInfo) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *S2CUserInfo) GetChairId() int32 {
	if m != nil {
		return m.ChairId
	}
	return 0
}

func (m *S2CUserInfo) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *S2CUserInfo) GetIsSettleCards() bool {
	if m != nil {
		return m.IsSettleCards
	}
	return false
}

func (m *S2CUserInfo) GetSpecialType() int32 {
	if m != nil {
		return m.SpecialType
	}
	return 0
}

func (m *S2CUserInfo) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

//玩家金额
type S2CUserScore struct {
	Score                int64    `protobuf:"varint,1,opt,name=score,proto3" json:"score"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *S2CUserScore) Reset()         { *m = S2CUserScore{} }
func (m *S2CUserScore) String() string { return proto.CompactTextString(m) }
func (*S2CUserScore) ProtoMessage()    {}
func (*S2CUserScore) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a66076066de807d, []int{8}
}

func (m *S2CUserScore) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S2CUserScore.Unmarshal(m, b)
}
func (m *S2CUserScore) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S2CUserScore.Marshal(b, m, deterministic)
}
func (m *S2CUserScore) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2CUserScore.Merge(m, src)
}
func (m *S2CUserScore) XXX_Size() int {
	return xxx_messageInfo_S2CUserScore.Size(m)
}
func (m *S2CUserScore) XXX_DiscardUnknown() {
	xxx_messageInfo_S2CUserScore.DiscardUnknown(m)
}

var xxx_messageInfo_S2CUserScore proto.InternalMessageInfo

func (m *S2CUserScore) GetScore() int64 {
	if m != nil {
		return m.Score
	}
	return 0
}

func init() {
	proto.RegisterEnum("msg.S2CMsgType", S2CMsgType_name, S2CMsgType_value)
	proto.RegisterType((*S2CRoomInfo)(nil), "msg.S2CRoomInfo")
	proto.RegisterType((*S2CSettleCards)(nil), "msg.S2CSettleCards")
	proto.RegisterType((*S2CStartGame)(nil), "msg.S2CStartGame")
	proto.RegisterType((*S2CCardsAndCardType)(nil), "msg.S2CCardsAndCardType")
	proto.RegisterType((*S2CEndGame)(nil), "msg.S2CEndGame")
	proto.RegisterType((*S2CHitRob)(nil), "msg.S2CHitRob")
	proto.RegisterType((*S2CUserEndInfo)(nil), "msg.S2CUserEndInfo")
	proto.RegisterType((*S2CUserInfo)(nil), "msg.S2CUserInfo")
	proto.RegisterType((*S2CUserScore)(nil), "msg.S2cUserScore")
}

func init() { proto.RegisterFile("game_s2c.proto", fileDescriptor_7a66076066de807d) }

var fileDescriptor_7a66076066de807d = []byte{
	// 926 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x56, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0xc6, 0x76, 0x7e, 0x4f, 0xd2, 0x34, 0x3b, 0x5d, 0x8a, 0x41, 0x08, 0x45, 0xd6, 0x6a, 0x15,
	0x55, 0xa2, 0x12, 0x81, 0x17, 0x88, 0xb2, 0x61, 0x53, 0x69, 0xb3, 0x45, 0xe3, 0x54, 0xdc, 0x11,
	0x39, 0xb1, 0xdb, 0x8c, 0x88, 0xed, 0xc8, 0x9e, 0x08, 0xf1, 0x08, 0xbc, 0x0a, 0x3c, 0x04, 0x37,
	0x5c, 0xf1, 0x26, 0xbc, 0x05, 0x3a, 0x67, 0x66, 0x1c, 0x3b, 0x4d, 0xbb, 0xbd, 0x6a, 0xbe, 0xef,
	0x9c, 0x33, 0x73, 0xce, 0x37, 0xdf, 0x78, 0x0a, 0xbd, 0x87, 0x20, 0x8e, 0x96, 0xf9, 0x68, 0x7d,
	0xbd, 0xcb, 0x52, 0x99, 0x32, 0x27, 0xce, 0x1f, 0xbc, 0x7f, 0x6d, 0xe8, 0xf8, 0xa3, 0x09, 0x4f,
	0xd3, 0xf8, 0x26, 0xb9, 0x4f, 0xd9, 0x25, 0x34, 0x56, 0xa9, 0x94, 0x69, 0xec, 0x5a, 0x03, 0x6b,
	0xe8, 0x70, 0x8d, 0xd8, 0x00, 0x3a, 0x32, 0x58, 0x6d, 0x23, 0x5f, 0x06, 0x72, 0x9f, 0xbb, 0xf6,
	0xc0, 0x1a, 0xd6, 0x79, 0x99, 0x62, 0x57, 0xd0, 0xdc, 0xe7, 0x51, 0x36, 0xce, 0x32, 0xd7, 0x19,
	0x38, 0xc3, 0xce, 0xa8, 0x7f, 0x1d, 0xe7, 0x0f, 0xd7, 0xfe, 0x68, 0x72, 0x97, 0x47, 0x19, 0x2e,
	0xce, 0x4d, 0x02, 0xfb, 0x0e, 0xda, 0x79, 0xb4, 0xbd, 0x9f, 0x04, 0x59, 0x98, 0xbb, 0xb5, 0x81,
	0x35, 0xec, 0x8c, 0x2e, 0xca, 0xd9, 0xd3, 0x24, 0xa4, 0x82, 0x43, 0x16, 0x36, 0x26, 0xc5, 0xfa,
	0xd7, 0x28, 0x73, 0xeb, 0xb4, 0xb7, 0x46, 0xec, 0x07, 0x68, 0xe5, 0xbb, 0x20, 0x8b, 0x70, 0xdf,
	0x06, 0xed, 0xeb, 0x9a, 0x95, 0xa8, 0x70, 0x9c, 0x84, 0xf8, 0x77, 0xf1, 0xfb, 0x2e, 0xe2, 0x45,
	0x26, 0x8e, 0x93, 0xef, 0xa2, 0xb5, 0x08, 0xb6, 0x18, 0x70, 0x9b, 0x6a, 0x9c, 0x12, 0x85, 0xfb,
	0x65, 0x28, 0x4a, 0xe8, 0xb6, 0x94, 0x10, 0x0a, 0xb1, 0xaf, 0xa0, 0x35, 0x17, 0xc9, 0x07, 0x11,
	0x0b, 0xe9, 0xb6, 0x29, 0x52, 0x60, 0xef, 0x17, 0xe8, 0xf9, 0xa3, 0x89, 0x1f, 0x49, 0xb9, 0x8d,
	0x54, 0xd7, 0x7d, 0x70, 0xf6, 0x22, 0xd4, 0x5a, 0xe2, 0x4f, 0xe6, 0x42, 0x73, 0xbd, 0x09, 0x44,
	0x76, 0x13, 0x6a, 0x11, 0x0d, 0x3c, 0xee, 0xc9, 0x79, 0xd4, 0x93, 0xf7, 0xa7, 0x05, 0x5d, 0xdc,
	0x40, 0x06, 0x99, 0x7c, 0x1f, 0xc4, 0xd1, 0x71, 0x89, 0xf5, 0x78, 0x8c, 0xd2, 0xa9, 0xd8, 0x9f,
	0x3a, 0x95, 0x83, 0xc4, 0xce, 0x93, 0x12, 0xd7, 0x5e, 0x2a, 0xb1, 0xf7, 0xb7, 0x05, 0x17, 0x27,
	0x32, 0xd8, 0xd7, 0xd0, 0xde, 0x44, 0x41, 0xa8, 0xce, 0x1e, 0x3b, 0xee, 0xf2, 0x03, 0x81, 0xf2,
	0xc6, 0x42, 0x07, 0x6d, 0x0a, 0x16, 0x18, 0x2b, 0x65, 0x20, 0xb6, 0x2a, 0xe8, 0xa8, 0xca, 0x82,
	0xc0, 0x4a, 0x5c, 0x86, 0x84, 0xa8, 0x51, 0xff, 0x05, 0x46, 0xd1, 0x63, 0xa1, 0x42, 0xca, 0x3d,
	0x06, 0x62, 0x15, 0x2e, 0x41, 0xa1, 0x86, 0xaa, 0x32, 0xd8, 0xfb, 0xcb, 0x02, 0xf0, 0x47, 0x93,
	0x69, 0x12, 0x92, 0xd8, 0x6f, 0xa1, 0xb1, 0x11, 0x12, 0x45, 0xb0, 0x48, 0x84, 0x9e, 0x11, 0x61,
	0x26, 0x24, 0x4f, 0x57, 0x5c, 0x47, 0xd9, 0x37, 0x00, 0x9b, 0x34, 0x8e, 0xf8, 0x3e, 0xb9, 0x13,
	0xea, 0x90, 0x1d, 0x5e, 0x62, 0xd8, 0xb7, 0xc7, 0x17, 0xe5, 0xa4, 0xf5, 0x8b, 0x53, 0x79, 0x0b,
	0x3d, 0x5d, 0x3c, 0xd1, 0xbe, 0x51, 0xd3, 0x1d, 0xb1, 0xde, 0x1f, 0x36, 0xb4, 0x8b, 0x66, 0xf0,
	0x2c, 0x37, 0x42, 0xde, 0x15, 0xde, 0xd3, 0x08, 0xe7, 0x5d, 0x45, 0x33, 0x15, 0x51, 0xad, 0x15,
	0x98, 0x14, 0x14, 0xd2, 0x5f, 0xa7, 0x99, 0x72, 0x9f, 0xc3, 0x0b, 0x4c, 0x43, 0x09, 0x59, 0xed,
	0xa0, 0xc4, 0x30, 0x0f, 0xba, 0xb4, 0x8e, 0xc9, 0x50, 0x32, 0x57, 0x38, 0xcc, 0xd9, 0x08, 0x39,
	0x8b, 0x82, 0x50, 0xed, 0xd1, 0xa0, 0x3d, 0x2a, 0x1c, 0x3a, 0x7a, 0x23, 0xe4, 0x5c, 0xe8, 0x94,
	0x26, 0xa5, 0x94, 0x29, 0xbd, 0xca, 0x22, 0x10, 0x5b, 0x95, 0xd2, 0x2a, 0x56, 0x29, 0x38, 0xef,
	0x9f, 0x3a, 0xdd, 0xc4, 0x92, 0x9e, 0x27, 0x6e, 0x62, 0xc5, 0x88, 0xf6, 0x73, 0x46, 0x74, 0x9e,
	0x33, 0x62, 0xed, 0x39, 0x23, 0xd6, 0x9f, 0x36, 0x62, 0xe3, 0x69, 0x23, 0x36, 0xab, 0x46, 0x3c,
	0xbe, 0xe6, 0xad, 0xc7, 0xd7, 0x1c, 0xab, 0x53, 0x19, 0x6c, 0x7f, 0x16, 0x89, 0xdb, 0x51, 0x47,
	0x67, 0x30, 0xee, 0x89, 0xfb, 0x63, 0xa8, 0x4b, 0x21, 0x03, 0xd1, 0x24, 0xb1, 0xa0, 0xc0, 0x99,
	0x32, 0x89, 0x42, 0x58, 0x81, 0x7b, 0x63, 0xa0, 0xa7, 0x2a, 0x34, 0x34, 0xb3, 0xfd, 0xb4, 0xdd,
	0xe7, 0xee, 0xb9, 0xb6, 0x88, 0xc6, 0x7a, 0x36, 0x0a, 0xf5, 0x55, 0x95, 0x86, 0x66, 0x36, 0x0a,
	0xbd, 0xd2, 0xdd, 0x69, 0x4c, 0x5a, 0x62, 0xa7, 0x14, 0x64, 0x14, 0x3c, 0x10, 0x58, 0xf9, 0x9b,
	0x48, 0xd4, 0x41, 0x5f, 0xa8, 0x4a, 0x83, 0xcb, 0x5f, 0xd2, 0xd7, 0xd5, 0x2f, 0xe9, 0x10, 0xce,
	0xf5, 0xe5, 0x98, 0x99, 0x66, 0x3f, 0xa7, 0xe2, 0x63, 0xba, 0x74, 0xb9, 0xe6, 0xba, 0xf5, 0x4b,
	0x4a, 0x3c, 0x62, 0x4b, 0x2b, 0x2e, 0xcc, 0x20, 0x5f, 0x54, 0x56, 0x34, 0x34, 0xbb, 0x82, 0xbe,
	0xa1, 0x8a, 0xb1, 0x5c, 0x4a, 0x7d, 0xc4, 0xe3, 0xa5, 0xba, 0x17, 0x49, 0xa0, 0x8d, 0xfc, 0xa5,
	0xfa, 0x52, 0x1c, 0x18, 0xef, 0x3f, 0x8b, 0x1e, 0x67, 0xf3, 0xa5, 0x3e, 0xe1, 0x61, 0x06, 0xb5,
	0x24, 0x88, 0x23, 0xb2, 0x6f, 0x9b, 0xd3, 0x6f, 0xe4, 0xf0, 0x4c, 0xc8, 0xb5, 0x6d, 0x4e, 0xbf,
	0xf1, 0xa4, 0x83, 0x38, 0xdd, 0x27, 0x92, 0xec, 0xea, 0x70, 0x8d, 0xca, 0x1a, 0xd6, 0xab, 0x1a,
	0x5e, 0x42, 0x23, 0x57, 0x6f, 0xbd, 0x32, 0xaa, 0x46, 0xec, 0x0d, 0x9c, 0x89, 0xbc, 0xf4, 0xc4,
	0x91, 0x59, 0x5b, 0xbc, 0x4a, 0xbe, 0xc0, 0xb1, 0x3d, 0xb0, 0xc5, 0x8e, 0x5e, 0xd0, 0x36, 0xb7,
	0xc5, 0xce, 0x7b, 0x83, 0x4f, 0xdb, 0x1a, 0x47, 0x55, 0xa7, 0xfb, 0x1a, 0xea, 0x39, 0xc9, 0xa2,
	0xa6, 0x55, 0xe0, 0x2a, 0xa4, 0x2f, 0xf2, 0x3c, 0x7f, 0xa0, 0x35, 0x5e, 0xc1, 0x19, 0xbf, 0xbd,
	0x9d, 0x2f, 0x6f, 0x3e, 0xfe, 0x78, 0xbb, 0xe4, 0x53, 0xbf, 0xff, 0x19, 0xbb, 0x80, 0x73, 0x7f,
	0x31, 0xe6, 0x8b, 0xe5, 0x7c, 0xbc, 0x98, 0xcc, 0x88, 0xb4, 0x58, 0x0f, 0x40, 0x91, 0xef, 0xc7,
	0xf3, 0x69, 0xdf, 0x66, 0x5d, 0x68, 0x4d, 0x3f, 0xbe, 0x53, 0xc8, 0x61, 0x7d, 0xe8, 0xfa, 0xd3,
	0xc5, 0xe2, 0xc3, 0x74, 0x39, 0x19, 0xf3, 0x77, 0x7e, 0xbf, 0xb6, 0x6a, 0xd0, 0x3f, 0x48, 0xdf,
	0xff, 0x1f, 0x00, 0x00, 0xff, 0xff, 0x97, 0x3a, 0x2d, 0x28, 0x32, 0x09, 0x00, 0x00,
}