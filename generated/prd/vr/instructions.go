// Code generated by https://github.com/zheng-lan/anchor-go. DO NOT EDIT.

package vr

import (
	"bytes"
	"fmt"
	ag_spew "github.com/davecgh/go-spew/spew"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_text "github.com/gagliardetto/solana-go/text"
	ag_treeout "github.com/gagliardetto/treeout"
)

const ProgramName = "Vr"

var ProgramID ag_solanago.PublicKey = ag_solanago.MustPublicKeyFromBase58("vrTGoBuy5rYSxAfV3jaRJWHH6nN9WK4NRExGxsk1bCJ")

func SetProgramID(PublicKey ag_solanago.PublicKey) {
	ProgramID = PublicKey
	ag_solanago.RegisterInstructionDecoder(ProgramID, registryDecodeInstruction)
}

func init() {
	if !ProgramID.IsZero() {
		ag_solanago.RegisterInstructionDecoder(ProgramID, registryDecodeInstruction)
	}
}

var (
	RentProgram = ag_solanago.MustPublicKeyFromBase58("SysvarRent111111111111111111111111111111111")

	SystemProgram = ag_solanago.MustPublicKeyFromBase58("11111111111111111111111111111111")
)

var (
	Instruction_Buy = ag_binary.TypeID([8]byte{102, 6, 61, 18, 1, 218, 235, 234})

	Instruction_Claim = ag_binary.TypeID([8]byte{62, 198, 214, 193, 213, 159, 108, 210})

	Instruction_Create = ag_binary.TypeID([8]byte{24, 30, 200, 40, 5, 28, 7, 119})

	Instruction_QuoteBuy = ag_binary.TypeID([8]byte{83, 9, 231, 110, 146, 31, 40, 12})

	Instruction_QuoteSell = ag_binary.TypeID([8]byte{5, 178, 49, 206, 140, 231, 131, 145})

	Instruction_Sell = ag_binary.TypeID([8]byte{51, 230, 133, 164, 1, 127, 131, 173})
)

// InstructionIDToName returns the name of the instruction given its ID.
func InstructionIDToName(id ag_binary.TypeID) string {
	switch id {
	case Instruction_Buy:
		return "Buy"
	case Instruction_Claim:
		return "Claim"
	case Instruction_Create:
		return "Create"
	case Instruction_QuoteBuy:
		return "QuoteBuy"
	case Instruction_QuoteSell:
		return "QuoteSell"
	case Instruction_Sell:
		return "Sell"
	default:
		return ""
	}
}

type Instruction struct {
	ag_binary.BaseVariant
}

func (inst *Instruction) EncodeToTree(parent ag_treeout.Branches) {
	if enToTree, ok := inst.Impl.(ag_text.EncodableToTree); ok {
		enToTree.EncodeToTree(parent)
	} else {
		parent.Child(ag_spew.Sdump(inst))
	}
}

var InstructionImplDef = ag_binary.NewVariantDefinition(
	ag_binary.AnchorTypeIDEncoding,
	[]ag_binary.VariantType{
		{
			Name: "buy", Type: (*Buy)(nil),
		},
		{
			Name: "claim", Type: (*Claim)(nil),
		},
		{
			Name: "create", Type: (*Create)(nil),
		},
		{
			Name: "quote_buy", Type: (*QuoteBuy)(nil),
		},
		{
			Name: "quote_sell", Type: (*QuoteSell)(nil),
		},
		{
			Name: "sell", Type: (*Sell)(nil),
		},
	},
)

func (inst *Instruction) ProgramID() ag_solanago.PublicKey {
	return ProgramID
}

func (inst *Instruction) Accounts() (out []*ag_solanago.AccountMeta) {
	return inst.Impl.(ag_solanago.AccountsGettable).GetAccounts()
}

func (inst *Instruction) Data() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := ag_binary.NewBorshEncoder(buf).Encode(inst); err != nil {
		return nil, fmt.Errorf("unable to encode instruction: %w", err)
	}
	return buf.Bytes(), nil
}

func (inst *Instruction) TextEncode(encoder *ag_text.Encoder, option *ag_text.Option) error {
	return encoder.Encode(inst.Impl, option)
}

func (inst *Instruction) UnmarshalWithDecoder(decoder *ag_binary.Decoder) error {
	return inst.BaseVariant.UnmarshalBinaryVariant(decoder, InstructionImplDef)
}

func (inst *Instruction) MarshalWithEncoder(encoder *ag_binary.Encoder) error {
	err := encoder.WriteBytes(inst.TypeID.Bytes(), false)
	if err != nil {
		return fmt.Errorf("unable to write variant type: %w", err)
	}
	return encoder.Encode(inst.Impl)
}

func registryDecodeInstruction(accounts []*ag_solanago.AccountMeta, data []byte) (interface{}, error) {
	inst, err := decodeInstruction(accounts, data)
	if err != nil {
		return nil, err
	}
	return inst, nil
}

func decodeInstruction(accounts []*ag_solanago.AccountMeta, data []byte) (*Instruction, error) {
	inst := new(Instruction)
	if err := ag_binary.NewBorshDecoder(data).Decode(inst); err != nil {
		return nil, fmt.Errorf("unable to decode instruction: %w", err)
	}
	if v, ok := inst.Impl.(ag_solanago.AccountsSettable); ok {
		err := v.SetAccounts(accounts)
		if err != nil {
			return nil, fmt.Errorf("unable to set accounts for instruction: %w", err)
		}
	}
	return inst, nil
}
