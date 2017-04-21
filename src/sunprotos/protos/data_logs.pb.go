// Code generated by protoc-gen-go.
// source: data_logs.proto
// DO NOT EDIT!

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of ActivityType from enums.proto

// Ignoring public import of BadgeType from enums.proto

// Ignoring public import of CameraInterpolation from enums.proto

// Ignoring public import of CameraTarget from enums.proto

// Ignoring public import of Costume from enums.proto

// Ignoring public import of EncounterType from enums.proto

// Ignoring public import of Filter from enums.proto

// Ignoring public import of Form from enums.proto

// Ignoring public import of Gender from enums.proto

// Ignoring public import of HoloIapItemCategory from enums.proto

// Ignoring public import of ItemCategory from enums.proto

// Ignoring public import of ItemEffect from enums.proto

// Ignoring public import of NotificationState from enums.proto

// Ignoring public import of Platform from enums.proto

// Ignoring public import of PokemonFamilyId from enums.proto

// Ignoring public import of PokemonId from enums.proto

// Ignoring public import of PokemonMove from enums.proto

// Ignoring public import of PokemonMovementType from enums.proto

// Ignoring public import of PokemonRarity from enums.proto

// Ignoring public import of PokemonType from enums.proto

// Ignoring public import of QuestType from enums.proto

// Ignoring public import of Slot from enums.proto

// Ignoring public import of TeamColor from enums.proto

// Ignoring public import of TutorialState from enums.proto

// Ignoring public import of AssetDigestEntry from data.proto

// Ignoring public import of BackgroundToken from data.proto

// Ignoring public import of BuddyPokemon from data.proto

// Ignoring public import of ClientVersion from data.proto

// Ignoring public import of DownloadUrlEntry from data.proto

// Ignoring public import of PlayerBadge from data.proto

// Ignoring public import of PlayerData from data.proto

// Ignoring public import of PokedexEntry from data.proto

// Ignoring public import of PokemonData from data.proto

// Ignoring public import of PokemonDisplay from data.proto

// Ignoring public import of RedeemPasscodeReward from data.proto

// Ignoring public import of ItemAward from inventory_item.proto

// Ignoring public import of ItemData from inventory_item.proto

// Ignoring public import of ItemId from inventory_item.proto

// Ignoring public import of ItemType from inventory_item.proto

type BuddyPokemonLogEntry_Result int32

const (
	BuddyPokemonLogEntry_UNSET       BuddyPokemonLogEntry_Result = 0
	BuddyPokemonLogEntry_CANDY_FOUND BuddyPokemonLogEntry_Result = 1
)

var BuddyPokemonLogEntry_Result_name = map[int32]string{
	0: "UNSET",
	1: "CANDY_FOUND",
}
var BuddyPokemonLogEntry_Result_value = map[string]int32{
	"UNSET":       0,
	"CANDY_FOUND": 1,
}

func (x BuddyPokemonLogEntry_Result) String() string {
	return proto.EnumName(BuddyPokemonLogEntry_Result_name, int32(x))
}
func (BuddyPokemonLogEntry_Result) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor6, []int{1, 0}
}

type CatchPokemonLogEntry_Result int32

const (
	CatchPokemonLogEntry_UNSET            CatchPokemonLogEntry_Result = 0
	CatchPokemonLogEntry_POKEMON_CAPTURED CatchPokemonLogEntry_Result = 1
	CatchPokemonLogEntry_POKEMON_FLED     CatchPokemonLogEntry_Result = 2
	CatchPokemonLogEntry_POKEMON_HATCHED  CatchPokemonLogEntry_Result = 3
)

var CatchPokemonLogEntry_Result_name = map[int32]string{
	0: "UNSET",
	1: "POKEMON_CAPTURED",
	2: "POKEMON_FLED",
	3: "POKEMON_HATCHED",
}
var CatchPokemonLogEntry_Result_value = map[string]int32{
	"UNSET":            0,
	"POKEMON_CAPTURED": 1,
	"POKEMON_FLED":     2,
	"POKEMON_HATCHED":  3,
}

func (x CatchPokemonLogEntry_Result) String() string {
	return proto.EnumName(CatchPokemonLogEntry_Result_name, int32(x))
}
func (CatchPokemonLogEntry_Result) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor6, []int{2, 0}
}

type FortSearchLogEntry_Result int32

const (
	FortSearchLogEntry_UNSET   FortSearchLogEntry_Result = 0
	FortSearchLogEntry_SUCCESS FortSearchLogEntry_Result = 1
)

var FortSearchLogEntry_Result_name = map[int32]string{
	0: "UNSET",
	1: "SUCCESS",
}
var FortSearchLogEntry_Result_value = map[string]int32{
	"UNSET":   0,
	"SUCCESS": 1,
}

func (x FortSearchLogEntry_Result) String() string {
	return proto.EnumName(FortSearchLogEntry_Result_name, int32(x))
}
func (FortSearchLogEntry_Result) EnumDescriptor() ([]byte, []int) { return fileDescriptor6, []int{3, 0} }

type ActionLogEntry struct {
	TimestampMs int64 `protobuf:"varint,1,opt,name=timestamp_ms,json=timestampMs" json:"timestamp_ms,omitempty"`
	Sfida       bool  `protobuf:"varint,2,opt,name=sfida" json:"sfida,omitempty"`
	// Types that are valid to be assigned to Action:
	//	*ActionLogEntry_CatchPokemon
	//	*ActionLogEntry_FortSearch
	//	*ActionLogEntry_BuddyPokemon
	Action isActionLogEntry_Action `protobuf_oneof:"Action"`
}

func (m *ActionLogEntry) Reset()                    { *m = ActionLogEntry{} }
func (m *ActionLogEntry) String() string            { return proto.CompactTextString(m) }
func (*ActionLogEntry) ProtoMessage()               {}
func (*ActionLogEntry) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

type isActionLogEntry_Action interface {
	isActionLogEntry_Action()
}

type ActionLogEntry_CatchPokemon struct {
	CatchPokemon *CatchPokemonLogEntry `protobuf:"bytes,3,opt,name=catch_pokemon,json=catchPokemon,oneof"`
}
type ActionLogEntry_FortSearch struct {
	FortSearch *FortSearchLogEntry `protobuf:"bytes,4,opt,name=fort_search,json=fortSearch,oneof"`
}
type ActionLogEntry_BuddyPokemon struct {
	BuddyPokemon *BuddyPokemonLogEntry `protobuf:"bytes,5,opt,name=buddy_pokemon,json=buddyPokemon,oneof"`
}

func (*ActionLogEntry_CatchPokemon) isActionLogEntry_Action() {}
func (*ActionLogEntry_FortSearch) isActionLogEntry_Action()   {}
func (*ActionLogEntry_BuddyPokemon) isActionLogEntry_Action() {}

func (m *ActionLogEntry) GetAction() isActionLogEntry_Action {
	if m != nil {
		return m.Action
	}
	return nil
}

func (m *ActionLogEntry) GetTimestampMs() int64 {
	if m != nil {
		return m.TimestampMs
	}
	return 0
}

func (m *ActionLogEntry) GetSfida() bool {
	if m != nil {
		return m.Sfida
	}
	return false
}

func (m *ActionLogEntry) GetCatchPokemon() *CatchPokemonLogEntry {
	if x, ok := m.GetAction().(*ActionLogEntry_CatchPokemon); ok {
		return x.CatchPokemon
	}
	return nil
}

func (m *ActionLogEntry) GetFortSearch() *FortSearchLogEntry {
	if x, ok := m.GetAction().(*ActionLogEntry_FortSearch); ok {
		return x.FortSearch
	}
	return nil
}

func (m *ActionLogEntry) GetBuddyPokemon() *BuddyPokemonLogEntry {
	if x, ok := m.GetAction().(*ActionLogEntry_BuddyPokemon); ok {
		return x.BuddyPokemon
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ActionLogEntry) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ActionLogEntry_OneofMarshaler, _ActionLogEntry_OneofUnmarshaler, _ActionLogEntry_OneofSizer, []interface{}{
		(*ActionLogEntry_CatchPokemon)(nil),
		(*ActionLogEntry_FortSearch)(nil),
		(*ActionLogEntry_BuddyPokemon)(nil),
	}
}

func _ActionLogEntry_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ActionLogEntry)
	// Action
	switch x := m.Action.(type) {
	case *ActionLogEntry_CatchPokemon:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CatchPokemon); err != nil {
			return err
		}
	case *ActionLogEntry_FortSearch:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.FortSearch); err != nil {
			return err
		}
	case *ActionLogEntry_BuddyPokemon:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.BuddyPokemon); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ActionLogEntry.Action has unexpected type %T", x)
	}
	return nil
}

func _ActionLogEntry_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ActionLogEntry)
	switch tag {
	case 3: // Action.catch_pokemon
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CatchPokemonLogEntry)
		err := b.DecodeMessage(msg)
		m.Action = &ActionLogEntry_CatchPokemon{msg}
		return true, err
	case 4: // Action.fort_search
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(FortSearchLogEntry)
		err := b.DecodeMessage(msg)
		m.Action = &ActionLogEntry_FortSearch{msg}
		return true, err
	case 5: // Action.buddy_pokemon
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(BuddyPokemonLogEntry)
		err := b.DecodeMessage(msg)
		m.Action = &ActionLogEntry_BuddyPokemon{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ActionLogEntry_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ActionLogEntry)
	// Action
	switch x := m.Action.(type) {
	case *ActionLogEntry_CatchPokemon:
		s := proto.Size(x.CatchPokemon)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ActionLogEntry_FortSearch:
		s := proto.Size(x.FortSearch)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ActionLogEntry_BuddyPokemon:
		s := proto.Size(x.BuddyPokemon)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type BuddyPokemonLogEntry struct {
	Result         BuddyPokemonLogEntry_Result `protobuf:"varint,1,opt,name=result,enum=SUNProtos.Data.Logs.BuddyPokemonLogEntry_Result" json:"result,omitempty"`
	PokemonId      PokemonId                   `protobuf:"varint,2,opt,name=pokemon_id,json=pokemonId,enum=SUNProtos.Enums.PokemonId" json:"pokemon_id,omitempty"`
	Amount         int32                       `protobuf:"varint,3,opt,name=amount" json:"amount,omitempty"`
	PokemonDisplay *PokemonDisplay             `protobuf:"bytes,4,opt,name=pokemon_display,json=pokemonDisplay" json:"pokemon_display,omitempty"`
	Pokemon        uint64                      `protobuf:"fixed64,5,opt,name=pokemon" json:"pokemon,omitempty"`
}

func (m *BuddyPokemonLogEntry) Reset()                    { *m = BuddyPokemonLogEntry{} }
func (m *BuddyPokemonLogEntry) String() string            { return proto.CompactTextString(m) }
func (*BuddyPokemonLogEntry) ProtoMessage()               {}
func (*BuddyPokemonLogEntry) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{1} }

func (m *BuddyPokemonLogEntry) GetResult() BuddyPokemonLogEntry_Result {
	if m != nil {
		return m.Result
	}
	return BuddyPokemonLogEntry_UNSET
}

func (m *BuddyPokemonLogEntry) GetPokemonId() PokemonId {
	if m != nil {
		return m.PokemonId
	}
	return PokemonId_MISSINGNO
}

func (m *BuddyPokemonLogEntry) GetAmount() int32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *BuddyPokemonLogEntry) GetPokemonDisplay() *PokemonDisplay {
	if m != nil {
		return m.PokemonDisplay
	}
	return nil
}

func (m *BuddyPokemonLogEntry) GetPokemon() uint64 {
	if m != nil {
		return m.Pokemon
	}
	return 0
}

type CatchPokemonLogEntry struct {
	Result         CatchPokemonLogEntry_Result `protobuf:"varint,1,opt,name=result,enum=SUNProtos.Data.Logs.CatchPokemonLogEntry_Result" json:"result,omitempty"`
	PokemonId      PokemonId                   `protobuf:"varint,2,opt,name=pokemon_id,json=pokemonId,enum=SUNProtos.Enums.PokemonId" json:"pokemon_id,omitempty"`
	CombatPoints   int32                       `protobuf:"varint,3,opt,name=combat_points,json=combatPoints" json:"combat_points,omitempty"`
	PokemonDataId  uint64                      `protobuf:"fixed64,4,opt,name=pokemon_data_id,json=pokemonDataId" json:"pokemon_data_id,omitempty"`
	PokemonDisplay *PokemonDisplay             `protobuf:"bytes,5,opt,name=pokemon_display,json=pokemonDisplay" json:"pokemon_display,omitempty"`
}

func (m *CatchPokemonLogEntry) Reset()                    { *m = CatchPokemonLogEntry{} }
func (m *CatchPokemonLogEntry) String() string            { return proto.CompactTextString(m) }
func (*CatchPokemonLogEntry) ProtoMessage()               {}
func (*CatchPokemonLogEntry) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{2} }

func (m *CatchPokemonLogEntry) GetResult() CatchPokemonLogEntry_Result {
	if m != nil {
		return m.Result
	}
	return CatchPokemonLogEntry_UNSET
}

func (m *CatchPokemonLogEntry) GetPokemonId() PokemonId {
	if m != nil {
		return m.PokemonId
	}
	return PokemonId_MISSINGNO
}

func (m *CatchPokemonLogEntry) GetCombatPoints() int32 {
	if m != nil {
		return m.CombatPoints
	}
	return 0
}

func (m *CatchPokemonLogEntry) GetPokemonDataId() uint64 {
	if m != nil {
		return m.PokemonDataId
	}
	return 0
}

func (m *CatchPokemonLogEntry) GetPokemonDisplay() *PokemonDisplay {
	if m != nil {
		return m.PokemonDisplay
	}
	return nil
}

type FortSearchLogEntry struct {
	Result      FortSearchLogEntry_Result `protobuf:"varint,1,opt,name=result,enum=SUNProtos.Data.Logs.FortSearchLogEntry_Result" json:"result,omitempty"`
	FortId      string                    `protobuf:"bytes,2,opt,name=fort_id,json=fortId" json:"fort_id,omitempty"`
	Items       []*ItemData               `protobuf:"bytes,3,rep,name=items" json:"items,omitempty"`
	Eggs        int32                     `protobuf:"varint,4,opt,name=eggs" json:"eggs,omitempty"`
	PokemonEggs []*PokemonData            `protobuf:"bytes,5,rep,name=pokemon_eggs,json=pokemonEggs" json:"pokemon_eggs,omitempty"`
}

func (m *FortSearchLogEntry) Reset()                    { *m = FortSearchLogEntry{} }
func (m *FortSearchLogEntry) String() string            { return proto.CompactTextString(m) }
func (*FortSearchLogEntry) ProtoMessage()               {}
func (*FortSearchLogEntry) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{3} }

func (m *FortSearchLogEntry) GetResult() FortSearchLogEntry_Result {
	if m != nil {
		return m.Result
	}
	return FortSearchLogEntry_UNSET
}

func (m *FortSearchLogEntry) GetFortId() string {
	if m != nil {
		return m.FortId
	}
	return ""
}

func (m *FortSearchLogEntry) GetItems() []*ItemData {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *FortSearchLogEntry) GetEggs() int32 {
	if m != nil {
		return m.Eggs
	}
	return 0
}

func (m *FortSearchLogEntry) GetPokemonEggs() []*PokemonData {
	if m != nil {
		return m.PokemonEggs
	}
	return nil
}

func init() {
	proto.RegisterType((*ActionLogEntry)(nil), "SUNProtos.Data.Logs.ActionLogEntry")
	proto.RegisterType((*BuddyPokemonLogEntry)(nil), "SUNProtos.Data.Logs.BuddyPokemonLogEntry")
	proto.RegisterType((*CatchPokemonLogEntry)(nil), "SUNProtos.Data.Logs.CatchPokemonLogEntry")
	proto.RegisterType((*FortSearchLogEntry)(nil), "SUNProtos.Data.Logs.FortSearchLogEntry")
	proto.RegisterEnum("SUNProtos.Data.Logs.BuddyPokemonLogEntry_Result", BuddyPokemonLogEntry_Result_name, BuddyPokemonLogEntry_Result_value)
	proto.RegisterEnum("SUNProtos.Data.Logs.CatchPokemonLogEntry_Result", CatchPokemonLogEntry_Result_name, CatchPokemonLogEntry_Result_value)
	proto.RegisterEnum("SUNProtos.Data.Logs.FortSearchLogEntry_Result", FortSearchLogEntry_Result_name, FortSearchLogEntry_Result_value)
}

func init() { proto.RegisterFile("data_logs.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 622 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xc1, 0x6e, 0xd3, 0x4c,
	0x10, 0xc7, 0xeb, 0xa4, 0x76, 0xdb, 0x71, 0x9a, 0x58, 0xdb, 0xe8, 0xfb, 0xa2, 0x22, 0xa1, 0x60,
	0x10, 0x84, 0x8b, 0x85, 0xc2, 0x05, 0x2e, 0x48, 0x69, 0xec, 0x90, 0x40, 0x9b, 0x5a, 0x9b, 0xe6,
	0x00, 0x17, 0x6b, 0x13, 0xbb, 0xae, 0x45, 0xed, 0xb5, 0xbc, 0x1b, 0xa4, 0x9c, 0x79, 0x16, 0xde,
	0x82, 0x1b, 0x2f, 0x86, 0x76, 0x6d, 0x57, 0xa6, 0x18, 0x14, 0x21, 0x2e, 0xd1, 0xce, 0x3f, 0x33,
	0xff, 0x9d, 0xfc, 0x36, 0x33, 0xd0, 0xf1, 0x09, 0x27, 0xde, 0x2d, 0x0d, 0x99, 0x95, 0x66, 0x94,
	0x53, 0x74, 0xb2, 0x58, 0xce, 0x5d, 0x71, 0x62, 0x96, 0x4d, 0x38, 0xb1, 0xce, 0x69, 0xc8, 0x4e,
	0xf5, 0x20, 0xd9, 0xc4, 0x45, 0xc6, 0x29, 0x88, 0x92, 0xe2, 0xdc, 0x8d, 0x92, 0xcf, 0x41, 0xc2,
	0x69, 0xb6, 0xf5, 0x22, 0x1e, 0xc4, 0xb9, 0x6a, 0x7e, 0x6f, 0x40, 0x7b, 0xb4, 0xe6, 0x11, 0x4d,
	0xce, 0x69, 0xe8, 0x24, 0x3c, 0xdb, 0xa2, 0x47, 0xd0, 0xe2, 0x51, 0x1c, 0x30, 0x4e, 0xe2, 0xd4,
	0x8b, 0x59, 0x4f, 0xe9, 0x2b, 0x83, 0x26, 0xd6, 0xef, 0xb4, 0x0b, 0x86, 0xba, 0xa0, 0xb2, 0xeb,
	0xc8, 0x27, 0xbd, 0x46, 0x5f, 0x19, 0x1c, 0xe2, 0x3c, 0x40, 0x2e, 0x1c, 0xaf, 0x09, 0x5f, 0xdf,
	0x78, 0x29, 0xfd, 0x14, 0xc4, 0x34, 0xe9, 0x35, 0xfb, 0xca, 0x40, 0x1f, 0x3e, 0xb7, 0x6a, 0xfa,
	0xb4, 0xc6, 0x22, 0xd3, 0xcd, 0x13, 0xcb, 0xab, 0xa7, 0x7b, 0xb8, 0xb5, 0xae, 0xe8, 0xe8, 0x1d,
	0xe8, 0xd7, 0x34, 0xe3, 0x1e, 0x0b, 0x48, 0xb6, 0xbe, 0xe9, 0xed, 0x4b, 0xbf, 0x67, 0xb5, 0x7e,
	0x13, 0x9a, 0xf1, 0x85, 0x4c, 0xab, 0xb8, 0xc1, 0xf5, 0x9d, 0x2a, 0xba, 0x5b, 0x6d, 0x7c, 0x7f,
	0x7b, 0xd7, 0x9d, 0xfa, 0x87, 0xee, 0xce, 0x44, 0x66, 0x4d, 0x77, 0xab, 0x8a, 0x7e, 0x76, 0x08,
	0x5a, 0x8e, 0xce, 0xfc, 0xd6, 0x80, 0x6e, 0x5d, 0x09, 0x9a, 0x82, 0x96, 0x05, 0x6c, 0x73, 0xcb,
	0x25, 0xc5, 0xf6, 0xf0, 0xc5, 0xce, 0xb7, 0x59, 0x58, 0xd6, 0xe1, 0xa2, 0x1e, 0xbd, 0x06, 0x28,
	0x1a, 0xf7, 0x22, 0x5f, 0x72, 0x6f, 0x0f, 0x4f, 0x2b, 0x6e, 0x8e, 0x7c, 0xf6, 0xc2, 0x64, 0xe6,
	0xe3, 0xa3, 0xb4, 0x3c, 0xa2, 0xff, 0x40, 0x23, 0x31, 0xdd, 0x24, 0x5c, 0x3e, 0x88, 0x8a, 0x8b,
	0x08, 0xbd, 0x85, 0x4e, 0x69, 0xe9, 0x47, 0x2c, 0xbd, 0x25, 0xdb, 0x82, 0xf0, 0xc3, 0xfb, 0x5d,
	0x16, 0xb6, 0x76, 0x9e, 0x85, 0xdb, 0xe9, 0x4f, 0x31, 0xea, 0xc1, 0x41, 0x15, 0xaa, 0x86, 0xcb,
	0xd0, 0x7c, 0x02, 0x5a, 0xfe, 0x3b, 0xd0, 0x11, 0xa8, 0xcb, 0xf9, 0xc2, 0xb9, 0x32, 0xf6, 0x50,
	0x07, 0xf4, 0xf1, 0x68, 0x6e, 0x7f, 0xf0, 0x26, 0x97, 0xcb, 0xb9, 0x6d, 0x28, 0xe6, 0x97, 0x26,
	0x74, 0xeb, 0xfe, 0x0f, 0x3b, 0xe2, 0xab, 0x2b, 0xfd, 0x87, 0xf8, 0x1e, 0xc3, 0xf1, 0x9a, 0xc6,
	0x2b, 0xc2, 0xbd, 0x94, 0x46, 0x09, 0x67, 0x05, 0xc5, 0x56, 0x2e, 0xba, 0x52, 0x43, 0x4f, 0x2b,
	0x2c, 0xc5, 0x98, 0x46, 0xbe, 0x64, 0xa9, 0xe1, 0xe3, 0x92, 0x15, 0xe1, 0x64, 0xe6, 0xd7, 0x31,
	0x57, 0xff, 0x86, 0xb9, 0xe9, 0xd6, 0x91, 0xed, 0x82, 0xe1, 0x5e, 0xbe, 0x77, 0x2e, 0x2e, 0xe7,
	0xde, 0x78, 0xe4, 0x5e, 0x2d, 0xb1, 0x63, 0x1b, 0x0a, 0x32, 0xa0, 0x55, 0xaa, 0x93, 0x73, 0xc7,
	0x36, 0x1a, 0xe8, 0x04, 0x3a, 0xa5, 0x32, 0x1d, 0x5d, 0x8d, 0xa7, 0x8e, 0x6d, 0x34, 0xcd, 0xaf,
	0x0d, 0x40, 0xbf, 0x4e, 0x11, 0x9a, 0xdc, 0x7b, 0x03, 0x6b, 0xc7, 0xf1, 0xbb, 0xff, 0x02, 0xff,
	0xc3, 0x81, 0x9c, 0xe5, 0x02, 0xff, 0x11, 0xd6, 0x44, 0x38, 0xf3, 0xd1, 0x2b, 0x50, 0xc5, 0x42,
	0x12, 0x5c, 0x9b, 0x03, 0x7d, 0x68, 0x56, 0xfc, 0x67, 0xe5, 0xca, 0xb2, 0x66, 0x62, 0x65, 0x89,
	0x0f, 0x71, 0x25, 0xce, 0x0b, 0x10, 0x82, 0xfd, 0x20, 0x0c, 0x99, 0x24, 0xad, 0x62, 0x79, 0x46,
	0x6f, 0xa0, 0x55, 0x02, 0x96, 0xdf, 0xa9, 0xd2, 0xf4, 0xc1, 0xef, 0xe8, 0x0a, 0x37, 0xbd, 0x28,
	0x70, 0xc2, 0x90, 0x99, 0xfd, 0x3a, 0xae, 0x3a, 0x1c, 0x2c, 0x96, 0xe3, 0xb1, 0xb3, 0x58, 0x18,
	0xca, 0xd9, 0xe1, 0x47, 0x4d, 0xee, 0x4e, 0xe6, 0xee, 0xb9, 0x8a, 0xdb, 0x58, 0xe5, 0xd1, 0xcb,
	0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa2, 0xeb, 0x16, 0x8f, 0xa0, 0x05, 0x00, 0x00,
}