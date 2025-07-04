// Code generated by https://github.com/zheng-lan/anchor-go. DO NOT EDIT.

package orca_whirlpool

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Initializes a fee_tier account usable by Whirlpools in a WhirlpoolConfig space.
//
// ### Authority
// - "fee_authority" - Set authority in the WhirlpoolConfig
//
// ### Parameters
// - `tick_spacing` - The tick-spacing that this fee-tier suggests the default_fee_rate for.
// - `default_fee_rate` - The default fee rate that a pool will use if the pool uses this
// fee tier during initialization.
//
// #### Special Errors
// - `InvalidTickSpacing` - If the provided tick_spacing is 0.
// - `FeeRateMaxExceeded` - If the provided default_fee_rate exceeds MAX_FEE_RATE.
type InitializeFeeTier struct {
	TickSpacing    *uint16
	DefaultFeeRate *uint16

	// [0] = [] config
	//
	// [1] = [WRITE] fee_tier
	//
	// [2] = [WRITE, SIGNER] funder
	//
	// [3] = [SIGNER] fee_authority
	//
	// [4] = [] system_program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewInitializeFeeTierInstructionBuilder creates a new `InitializeFeeTier` instruction builder.
func NewInitializeFeeTierInstructionBuilder() *InitializeFeeTier {
	nd := &InitializeFeeTier{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 5),
	}
	return nd
}

// SetTickSpacing sets the "tick_spacing" parameter.
func (inst *InitializeFeeTier) SetTickSpacing(tick_spacing uint16) *InitializeFeeTier {
	inst.TickSpacing = &tick_spacing
	return inst
}

// SetDefaultFeeRate sets the "default_fee_rate" parameter.
func (inst *InitializeFeeTier) SetDefaultFeeRate(default_fee_rate uint16) *InitializeFeeTier {
	inst.DefaultFeeRate = &default_fee_rate
	return inst
}

// SetConfigAccount sets the "config" account.
func (inst *InitializeFeeTier) SetConfigAccount(config ag_solanago.PublicKey) *InitializeFeeTier {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(config)
	return inst
}

// GetConfigAccount gets the "config" account.
func (inst *InitializeFeeTier) GetConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetFeeTierAccount sets the "fee_tier" account.
func (inst *InitializeFeeTier) SetFeeTierAccount(feeTier ag_solanago.PublicKey) *InitializeFeeTier {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(feeTier).WRITE()
	return inst
}

// GetFeeTierAccount gets the "fee_tier" account.
func (inst *InitializeFeeTier) GetFeeTierAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetFunderAccount sets the "funder" account.
func (inst *InitializeFeeTier) SetFunderAccount(funder ag_solanago.PublicKey) *InitializeFeeTier {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(funder).WRITE().SIGNER()
	return inst
}

// GetFunderAccount gets the "funder" account.
func (inst *InitializeFeeTier) GetFunderAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetFeeAuthorityAccount sets the "fee_authority" account.
func (inst *InitializeFeeTier) SetFeeAuthorityAccount(feeAuthority ag_solanago.PublicKey) *InitializeFeeTier {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(feeAuthority).SIGNER()
	return inst
}

// GetFeeAuthorityAccount gets the "fee_authority" account.
func (inst *InitializeFeeTier) GetFeeAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetSystemProgramAccount sets the "system_program" account.
func (inst *InitializeFeeTier) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *InitializeFeeTier {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "system_program" account.
func (inst *InitializeFeeTier) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

func (inst *InitializeFeeTier) SetAccounts(accounts []*ag_solanago.AccountMeta) error {
	inst.AccountMetaSlice = accounts
	return inst.Validate()
}

func (inst *InitializeFeeTier) SetRemainingAccounts(metas []*ag_solanago.AccountMeta) *InitializeFeeTier {
	inst.AccountMetaSlice = append(inst.AccountMetaSlice[0:5], metas...)
	return inst
}

func (inst *InitializeFeeTier) GetRemainingAccounts() []*ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[5:]
}

func (inst InitializeFeeTier) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_InitializeFeeTier,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst InitializeFeeTier) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *InitializeFeeTier) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.TickSpacing == nil {
			return errors.New("tickSpacing parameter is not set")
		}
		if inst.DefaultFeeRate == nil {
			return errors.New("defaultFeeRate parameter is not set")
		}
	}

	if len(inst.AccountMetaSlice) < 5 {
		return errors.New("accounts slice has wrong length: expected 5 accounts")
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Config is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.FeeTier is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Funder is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.FeeAuthority is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *InitializeFeeTier) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("InitializeFeeTier")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("     TickSpacing", *inst.TickSpacing))
						paramsBranch.Child(ag_format.Param("  DefaultFeeRate", *inst.DefaultFeeRate))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=5]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("        config", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("      fee_tier", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("        funder", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta(" fee_authority", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("system_program", inst.AccountMetaSlice.Get(4)))
					})
				})
		})
}

func (obj InitializeFeeTier) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `TickSpacing` param:
	err = encoder.Encode(obj.TickSpacing)
	if err != nil {
		return err
	}
	// Serialize `DefaultFeeRate` param:
	err = encoder.Encode(obj.DefaultFeeRate)
	if err != nil {
		return err
	}
	return nil
}
func (obj *InitializeFeeTier) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `TickSpacing`:
	err = decoder.Decode(&obj.TickSpacing)
	if err != nil {
		return err
	}
	// Deserialize `DefaultFeeRate`:
	err = decoder.Decode(&obj.DefaultFeeRate)
	if err != nil {
		return err
	}
	return nil
}

// NewInitializeFeeTierInstruction declares a new InitializeFeeTier instruction with the provided parameters and accounts.
func NewInitializeFeeTierInstruction(
	// Parameters:
	tick_spacing uint16,
	default_fee_rate uint16,
	// Accounts:
	config ag_solanago.PublicKey,
	feeTier ag_solanago.PublicKey,
	funder ag_solanago.PublicKey,
	feeAuthority ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *InitializeFeeTier {
	return NewInitializeFeeTierInstructionBuilder().
		SetTickSpacing(tick_spacing).
		SetDefaultFeeRate(default_fee_rate).
		SetConfigAccount(config).
		SetFeeTierAccount(feeTier).
		SetFunderAccount(funder).
		SetFeeAuthorityAccount(feeAuthority).
		SetSystemProgramAccount(systemProgram)
}

// NewSimpleInitializeFeeTierInstruction declares a new InitializeFeeTier instruction with the provided parameters and accounts.
func NewSimpleInitializeFeeTierInstruction(
	// Parameters:
	tick_spacing uint16,
	default_fee_rate uint16,
	// Accounts:
	config ag_solanago.PublicKey,
	feeTier ag_solanago.PublicKey,
	funder ag_solanago.PublicKey,
	feeAuthority ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *InitializeFeeTier {
	return NewInitializeFeeTierInstructionBuilder().
		SetTickSpacing(tick_spacing).
		SetDefaultFeeRate(default_fee_rate).
		SetConfigAccount(config).
		SetFeeTierAccount(feeTier).
		SetFunderAccount(funder).
		SetFeeAuthorityAccount(feeAuthority).
		SetSystemProgramAccount(systemProgram)
}
