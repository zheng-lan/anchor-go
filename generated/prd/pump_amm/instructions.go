// Code generated by https://github.com/zheng-lan/anchor-go. DO NOT EDIT.

package pump_amm

import (
	"bytes"
	"fmt"
	ag_spew "github.com/davecgh/go-spew/spew"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_text "github.com/gagliardetto/solana-go/text"
	ag_treeout "github.com/gagliardetto/treeout"
)

const ProgramName = "PumpAmm"

var ProgramID ag_solanago.PublicKey = ag_solanago.MustPublicKeyFromBase58("pAMMBay6oceH9fJKBRHGP5D4bD4sWpmSwMn52FMfXEA")

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
	Admin = ag_solanago.MustPublicKeyFromBase58("8LWu7QM2dGR1G8nKDHthckea57bkCzXyBTAKPJUBDHo8")

	AssociatedTokenProgram = ag_solanago.MustPublicKeyFromBase58("ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL")

	EventAuthorityPDA = ag_solanago.MustPublicKeyFromBase58("GS4CU59F31iL7aR2Q8zVS8DRrcRnXX1yjQ66TqNVQnaR")

	GlobalConfigPDA = ag_solanago.MustPublicKeyFromBase58("ADyA8hdefvWN2dbGGWFotbzWxrAvLW83WG6QCVXvJKqw")

	SystemProgram = ag_solanago.MustPublicKeyFromBase58("11111111111111111111111111111111")

	Token22Program = ag_solanago.MustPublicKeyFromBase58("TokenzQdBNbLqP5VEhdkAS6EPFLC1PHnBqCXEpPxuEb")

	TokenProgram = ag_solanago.MustPublicKeyFromBase58("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA")
)

var (
	Instruction_Buy = ag_binary.TypeID([8]byte{102, 6, 61, 18, 1, 218, 235, 234})

	Instruction_CollectCoinCreatorFee = ag_binary.TypeID([8]byte{160, 57, 89, 42, 181, 139, 43, 66})

	Instruction_CreateConfig = ag_binary.TypeID([8]byte{201, 207, 243, 114, 75, 111, 47, 189})

	Instruction_CreatePool = ag_binary.TypeID([8]byte{233, 146, 209, 142, 207, 104, 64, 188})

	Instruction_Deposit = ag_binary.TypeID([8]byte{242, 35, 198, 137, 82, 225, 242, 182})

	Instruction_Disable = ag_binary.TypeID([8]byte{185, 173, 187, 90, 216, 15, 238, 233})

	Instruction_ExtendAccount = ag_binary.TypeID([8]byte{234, 102, 194, 203, 150, 72, 62, 229})

	Instruction_Sell = ag_binary.TypeID([8]byte{51, 230, 133, 164, 1, 127, 131, 173})

	// Sets Pool::coin_creator from Metaplex metadata creator or BondingCurve::creator
	Instruction_SetCoinCreator = ag_binary.TypeID([8]byte{210, 149, 128, 45, 188, 58, 78, 175})

	Instruction_UpdateAdmin = ag_binary.TypeID([8]byte{161, 176, 40, 213, 60, 184, 179, 228})

	Instruction_UpdateFeeConfig = ag_binary.TypeID([8]byte{104, 184, 103, 242, 88, 151, 107, 20})

	Instruction_Withdraw = ag_binary.TypeID([8]byte{183, 18, 70, 156, 148, 109, 161, 34})
)

// InstructionIDToName returns the name of the instruction given its ID.
func InstructionIDToName(id ag_binary.TypeID) string {
	switch id {
	case Instruction_Buy:
		return "Buy"
	case Instruction_CollectCoinCreatorFee:
		return "CollectCoinCreatorFee"
	case Instruction_CreateConfig:
		return "CreateConfig"
	case Instruction_CreatePool:
		return "CreatePool"
	case Instruction_Deposit:
		return "Deposit"
	case Instruction_Disable:
		return "Disable"
	case Instruction_ExtendAccount:
		return "ExtendAccount"
	case Instruction_Sell:
		return "Sell"
	case Instruction_SetCoinCreator:
		return "SetCoinCreator"
	case Instruction_UpdateAdmin:
		return "UpdateAdmin"
	case Instruction_UpdateFeeConfig:
		return "UpdateFeeConfig"
	case Instruction_Withdraw:
		return "Withdraw"
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
			Name: "collect_coin_creator_fee", Type: (*CollectCoinCreatorFee)(nil),
		},
		{
			Name: "create_config", Type: (*CreateConfig)(nil),
		},
		{
			Name: "create_pool", Type: (*CreatePool)(nil),
		},
		{
			Name: "deposit", Type: (*Deposit)(nil),
		},
		{
			Name: "disable", Type: (*Disable)(nil),
		},
		{
			Name: "extend_account", Type: (*ExtendAccount)(nil),
		},
		{
			Name: "sell", Type: (*Sell)(nil),
		},
		{
			Name: "set_coin_creator", Type: (*SetCoinCreator)(nil),
		},
		{
			Name: "update_admin", Type: (*UpdateAdmin)(nil),
		},
		{
			Name: "update_fee_config", Type: (*UpdateFeeConfig)(nil),
		},
		{
			Name: "withdraw", Type: (*Withdraw)(nil),
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
