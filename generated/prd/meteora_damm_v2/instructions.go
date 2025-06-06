// Code generated by https://github.com/zheng-lan/anchor-go. DO NOT EDIT.

package meteora_damm_v2

import (
	"bytes"
	"fmt"
	ag_spew "github.com/davecgh/go-spew/spew"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_text "github.com/gagliardetto/solana-go/text"
	ag_treeout "github.com/gagliardetto/treeout"
)

const ProgramName = "MeteoraDammV2"

var ProgramID ag_solanago.PublicKey = ag_solanago.MustPublicKeyFromBase58("cpamdpZCGKUy5JxQXB4dcpGPiikHawvSWAd6mEn1sGG")

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
	EventAuthorityPDA = ag_solanago.MustPublicKeyFromBase58("3rmHSu74h1ZcmAisVcWerTCiRDQbUrBKmcwptYGjHfet")

	PoolAuthorityPDA = ag_solanago.MustPublicKeyFromBase58("HLnpSz9h2S4hiLQ43rnSD9XkcUThA7B8hQMKmDaiTLcC")

	SystemProgram = ag_solanago.MustPublicKeyFromBase58("11111111111111111111111111111111")

	Token22Program = ag_solanago.MustPublicKeyFromBase58("TokenzQdBNbLqP5VEhdkAS6EPFLC1PHnBqCXEpPxuEb")
)

var (
	Instruction_AddLiquidity = ag_binary.TypeID([8]byte{181, 157, 89, 67, 143, 182, 52, 72})

	Instruction_ClaimPartnerFee = ag_binary.TypeID([8]byte{97, 206, 39, 105, 94, 94, 126, 148})

	Instruction_ClaimPositionFee = ag_binary.TypeID([8]byte{180, 38, 154, 17, 133, 33, 162, 211})

	Instruction_ClaimProtocolFee = ag_binary.TypeID([8]byte{165, 228, 133, 48, 99, 249, 255, 33})

	Instruction_ClaimReward = ag_binary.TypeID([8]byte{149, 95, 181, 242, 94, 90, 158, 162})

	Instruction_CloseClaimFeeOperator = ag_binary.TypeID([8]byte{38, 134, 82, 216, 95, 124, 17, 99})

	Instruction_CloseConfig = ag_binary.TypeID([8]byte{145, 9, 72, 157, 95, 125, 61, 85})

	Instruction_ClosePosition = ag_binary.TypeID([8]byte{123, 134, 81, 0, 49, 68, 98, 98})

	Instruction_CreateClaimFeeOperator = ag_binary.TypeID([8]byte{169, 62, 207, 107, 58, 187, 162, 109})

	// ADMIN FUNCTIONS /////
	Instruction_CreateConfig = ag_binary.TypeID([8]byte{201, 207, 243, 114, 75, 111, 47, 189})

	Instruction_CreateDynamicConfig = ag_binary.TypeID([8]byte{81, 251, 122, 78, 66, 57, 208, 82})

	Instruction_CreatePosition = ag_binary.TypeID([8]byte{48, 215, 197, 153, 96, 203, 180, 133})

	Instruction_CreateTokenBadge = ag_binary.TypeID([8]byte{88, 206, 0, 91, 60, 175, 151, 118})

	Instruction_FundReward = ag_binary.TypeID([8]byte{188, 50, 249, 165, 93, 151, 38, 63})

	Instruction_InitializeCustomizablePool = ag_binary.TypeID([8]byte{20, 161, 241, 24, 189, 221, 180, 2})

	// USER FUNCTIONS ////
	Instruction_InitializePool = ag_binary.TypeID([8]byte{95, 180, 10, 172, 84, 174, 232, 40})

	Instruction_InitializePoolWithDynamicConfig = ag_binary.TypeID([8]byte{149, 82, 72, 197, 253, 252, 68, 15})

	Instruction_InitializeReward = ag_binary.TypeID([8]byte{95, 135, 192, 196, 242, 129, 230, 68})

	Instruction_LockPosition = ag_binary.TypeID([8]byte{227, 62, 2, 252, 247, 10, 171, 185})

	Instruction_PermanentLockPosition = ag_binary.TypeID([8]byte{165, 176, 125, 6, 231, 171, 186, 213})

	Instruction_RefreshVesting = ag_binary.TypeID([8]byte{9, 94, 216, 14, 116, 204, 247, 0})

	Instruction_RemoveAllLiquidity = ag_binary.TypeID([8]byte{10, 51, 61, 35, 112, 105, 24, 85})

	Instruction_RemoveLiquidity = ag_binary.TypeID([8]byte{80, 85, 209, 72, 24, 206, 177, 108})

	Instruction_SetPoolStatus = ag_binary.TypeID([8]byte{112, 87, 135, 223, 83, 204, 132, 53})

	Instruction_Swap = ag_binary.TypeID([8]byte{248, 198, 158, 145, 225, 117, 135, 200})

	Instruction_UpdateRewardDuration = ag_binary.TypeID([8]byte{138, 174, 196, 169, 213, 235, 254, 107})

	Instruction_UpdateRewardFunder = ag_binary.TypeID([8]byte{211, 28, 48, 32, 215, 160, 35, 23})

	Instruction_WithdrawIneligibleReward = ag_binary.TypeID([8]byte{148, 206, 42, 195, 247, 49, 103, 8})
)

// InstructionIDToName returns the name of the instruction given its ID.
func InstructionIDToName(id ag_binary.TypeID) string {
	switch id {
	case Instruction_AddLiquidity:
		return "AddLiquidity"
	case Instruction_ClaimPartnerFee:
		return "ClaimPartnerFee"
	case Instruction_ClaimPositionFee:
		return "ClaimPositionFee"
	case Instruction_ClaimProtocolFee:
		return "ClaimProtocolFee"
	case Instruction_ClaimReward:
		return "ClaimReward"
	case Instruction_CloseClaimFeeOperator:
		return "CloseClaimFeeOperator"
	case Instruction_CloseConfig:
		return "CloseConfig"
	case Instruction_ClosePosition:
		return "ClosePosition"
	case Instruction_CreateClaimFeeOperator:
		return "CreateClaimFeeOperator"
	case Instruction_CreateConfig:
		return "CreateConfig"
	case Instruction_CreateDynamicConfig:
		return "CreateDynamicConfig"
	case Instruction_CreatePosition:
		return "CreatePosition"
	case Instruction_CreateTokenBadge:
		return "CreateTokenBadge"
	case Instruction_FundReward:
		return "FundReward"
	case Instruction_InitializeCustomizablePool:
		return "InitializeCustomizablePool"
	case Instruction_InitializePool:
		return "InitializePool"
	case Instruction_InitializePoolWithDynamicConfig:
		return "InitializePoolWithDynamicConfig"
	case Instruction_InitializeReward:
		return "InitializeReward"
	case Instruction_LockPosition:
		return "LockPosition"
	case Instruction_PermanentLockPosition:
		return "PermanentLockPosition"
	case Instruction_RefreshVesting:
		return "RefreshVesting"
	case Instruction_RemoveAllLiquidity:
		return "RemoveAllLiquidity"
	case Instruction_RemoveLiquidity:
		return "RemoveLiquidity"
	case Instruction_SetPoolStatus:
		return "SetPoolStatus"
	case Instruction_Swap:
		return "Swap"
	case Instruction_UpdateRewardDuration:
		return "UpdateRewardDuration"
	case Instruction_UpdateRewardFunder:
		return "UpdateRewardFunder"
	case Instruction_WithdrawIneligibleReward:
		return "WithdrawIneligibleReward"
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
			Name: "add_liquidity", Type: (*AddLiquidity)(nil),
		},
		{
			Name: "claim_partner_fee", Type: (*ClaimPartnerFee)(nil),
		},
		{
			Name: "claim_position_fee", Type: (*ClaimPositionFee)(nil),
		},
		{
			Name: "claim_protocol_fee", Type: (*ClaimProtocolFee)(nil),
		},
		{
			Name: "claim_reward", Type: (*ClaimReward)(nil),
		},
		{
			Name: "close_claim_fee_operator", Type: (*CloseClaimFeeOperator)(nil),
		},
		{
			Name: "close_config", Type: (*CloseConfig)(nil),
		},
		{
			Name: "close_position", Type: (*ClosePosition)(nil),
		},
		{
			Name: "create_claim_fee_operator", Type: (*CreateClaimFeeOperator)(nil),
		},
		{
			Name: "create_config", Type: (*CreateConfig)(nil),
		},
		{
			Name: "create_dynamic_config", Type: (*CreateDynamicConfig)(nil),
		},
		{
			Name: "create_position", Type: (*CreatePosition)(nil),
		},
		{
			Name: "create_token_badge", Type: (*CreateTokenBadge)(nil),
		},
		{
			Name: "fund_reward", Type: (*FundReward)(nil),
		},
		{
			Name: "initialize_customizable_pool", Type: (*InitializeCustomizablePool)(nil),
		},
		{
			Name: "initialize_pool", Type: (*InitializePool)(nil),
		},
		{
			Name: "initialize_pool_with_dynamic_config", Type: (*InitializePoolWithDynamicConfig)(nil),
		},
		{
			Name: "initialize_reward", Type: (*InitializeReward)(nil),
		},
		{
			Name: "lock_position", Type: (*LockPosition)(nil),
		},
		{
			Name: "permanent_lock_position", Type: (*PermanentLockPosition)(nil),
		},
		{
			Name: "refresh_vesting", Type: (*RefreshVesting)(nil),
		},
		{
			Name: "remove_all_liquidity", Type: (*RemoveAllLiquidity)(nil),
		},
		{
			Name: "remove_liquidity", Type: (*RemoveLiquidity)(nil),
		},
		{
			Name: "set_pool_status", Type: (*SetPoolStatus)(nil),
		},
		{
			Name: "swap", Type: (*Swap)(nil),
		},
		{
			Name: "update_reward_duration", Type: (*UpdateRewardDuration)(nil),
		},
		{
			Name: "update_reward_funder", Type: (*UpdateRewardFunder)(nil),
		},
		{
			Name: "withdraw_ineligible_reward", Type: (*WithdrawIneligibleReward)(nil),
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
