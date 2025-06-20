// Code generated by https://github.com/zheng-lan/anchor-go. DO NOT EDIT.

package orca_whirlpool

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Sets the fee rate for a Whirlpool.
// Fee rate is represented as hundredths of a basis point.
// Only the current fee authority has permission to invoke this instruction.
//
// ### Authority
// - "fee_authority" - Set authority that can modify pool fees in the WhirlpoolConfig
//
// ### Parameters
// - `fee_rate` - The rate that the pool will use to calculate fees going onwards.
//
// #### Special Errors
// - `FeeRateMaxExceeded` - If the provided fee_rate exceeds MAX_FEE_RATE.
type SetFeeRate struct {
	FeeRate *uint16

	// [0] = [] whirlpools_config
	//
	// [1] = [WRITE] whirlpool
	//
	// [2] = [SIGNER] fee_authority
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewSetFeeRateInstructionBuilder creates a new `SetFeeRate` instruction builder.
func NewSetFeeRateInstructionBuilder() *SetFeeRate {
	nd := &SetFeeRate{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetFeeRate sets the "fee_rate" parameter.
func (inst *SetFeeRate) SetFeeRate(fee_rate uint16) *SetFeeRate {
	inst.FeeRate = &fee_rate
	return inst
}

// SetWhirlpoolsConfigAccount sets the "whirlpools_config" account.
func (inst *SetFeeRate) SetWhirlpoolsConfigAccount(whirlpoolsConfig ag_solanago.PublicKey) *SetFeeRate {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(whirlpoolsConfig)
	return inst
}

// GetWhirlpoolsConfigAccount gets the "whirlpools_config" account.
func (inst *SetFeeRate) GetWhirlpoolsConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetWhirlpoolAccount sets the "whirlpool" account.
func (inst *SetFeeRate) SetWhirlpoolAccount(whirlpool ag_solanago.PublicKey) *SetFeeRate {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(whirlpool).WRITE()
	return inst
}

// GetWhirlpoolAccount gets the "whirlpool" account.
func (inst *SetFeeRate) GetWhirlpoolAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetFeeAuthorityAccount sets the "fee_authority" account.
func (inst *SetFeeRate) SetFeeAuthorityAccount(feeAuthority ag_solanago.PublicKey) *SetFeeRate {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(feeAuthority).SIGNER()
	return inst
}

// GetFeeAuthorityAccount gets the "fee_authority" account.
func (inst *SetFeeRate) GetFeeAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

func (inst *SetFeeRate) SetAccounts(accounts []*ag_solanago.AccountMeta) error {
	inst.AccountMetaSlice = accounts
	return inst.Validate()
}

func (inst *SetFeeRate) SetRemainingAccounts(metas []*ag_solanago.AccountMeta) *SetFeeRate {
	inst.AccountMetaSlice = append(inst.AccountMetaSlice[0:3], metas...)
	return inst
}

func (inst *SetFeeRate) GetRemainingAccounts() []*ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3:]
}

func (inst SetFeeRate) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_SetFeeRate,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst SetFeeRate) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *SetFeeRate) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.FeeRate == nil {
			return errors.New("feeRate parameter is not set")
		}
	}

	if len(inst.AccountMetaSlice) < 3 {
		return errors.New("accounts slice has wrong length: expected 3 accounts")
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.WhirlpoolsConfig is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Whirlpool is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.FeeAuthority is not set")
		}
	}
	return nil
}

func (inst *SetFeeRate) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("SetFeeRate")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param(" FeeRate", *inst.FeeRate))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=3]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("whirlpools_config", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("        whirlpool", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("    fee_authority", inst.AccountMetaSlice.Get(2)))
					})
				})
		})
}

func (obj SetFeeRate) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `FeeRate` param:
	err = encoder.Encode(obj.FeeRate)
	if err != nil {
		return err
	}
	return nil
}
func (obj *SetFeeRate) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `FeeRate`:
	err = decoder.Decode(&obj.FeeRate)
	if err != nil {
		return err
	}
	return nil
}

// NewSetFeeRateInstruction declares a new SetFeeRate instruction with the provided parameters and accounts.
func NewSetFeeRateInstruction(
	// Parameters:
	fee_rate uint16,
	// Accounts:
	whirlpoolsConfig ag_solanago.PublicKey,
	whirlpool ag_solanago.PublicKey,
	feeAuthority ag_solanago.PublicKey) *SetFeeRate {
	return NewSetFeeRateInstructionBuilder().
		SetFeeRate(fee_rate).
		SetWhirlpoolsConfigAccount(whirlpoolsConfig).
		SetWhirlpoolAccount(whirlpool).
		SetFeeAuthorityAccount(feeAuthority)
}

// NewSimpleSetFeeRateInstruction declares a new SetFeeRate instruction with the provided parameters and accounts.
func NewSimpleSetFeeRateInstruction(
	// Parameters:
	fee_rate uint16,
	// Accounts:
	whirlpoolsConfig ag_solanago.PublicKey,
	whirlpool ag_solanago.PublicKey,
	feeAuthority ag_solanago.PublicKey) *SetFeeRate {
	return NewSetFeeRateInstructionBuilder().
		SetFeeRate(fee_rate).
		SetWhirlpoolsConfigAccount(whirlpoolsConfig).
		SetWhirlpoolAccount(whirlpool).
		SetFeeAuthorityAccount(feeAuthority)
}
