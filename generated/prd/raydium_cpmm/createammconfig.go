// Code generated by https://github.com/zheng-lan/anchor-go. DO NOT EDIT.

package raydium_cpmm

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// CreateAmmConfig is the `create_amm_config` instruction.
type CreateAmmConfig struct {
	Index           *uint16
	TradeFeeRate    *uint64
	ProtocolFeeRate *uint64
	FundFeeRate     *uint64
	CreatePoolFee   *uint64

	// [0] = [WRITE, SIGNER] owner
	//
	// [1] = [WRITE] amm_config
	//
	// [2] = [] system_program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewCreateAmmConfigInstructionBuilder creates a new `CreateAmmConfig` instruction builder.
func NewCreateAmmConfigInstructionBuilder() *CreateAmmConfig {
	nd := &CreateAmmConfig{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetIndex sets the "index" parameter.
func (inst *CreateAmmConfig) SetIndex(index uint16) *CreateAmmConfig {
	inst.Index = &index
	return inst
}

// SetTradeFeeRate sets the "trade_fee_rate" parameter.
func (inst *CreateAmmConfig) SetTradeFeeRate(trade_fee_rate uint64) *CreateAmmConfig {
	inst.TradeFeeRate = &trade_fee_rate
	return inst
}

// SetProtocolFeeRate sets the "protocol_fee_rate" parameter.
func (inst *CreateAmmConfig) SetProtocolFeeRate(protocol_fee_rate uint64) *CreateAmmConfig {
	inst.ProtocolFeeRate = &protocol_fee_rate
	return inst
}

// SetFundFeeRate sets the "fund_fee_rate" parameter.
func (inst *CreateAmmConfig) SetFundFeeRate(fund_fee_rate uint64) *CreateAmmConfig {
	inst.FundFeeRate = &fund_fee_rate
	return inst
}

// SetCreatePoolFee sets the "create_pool_fee" parameter.
func (inst *CreateAmmConfig) SetCreatePoolFee(create_pool_fee uint64) *CreateAmmConfig {
	inst.CreatePoolFee = &create_pool_fee
	return inst
}

// SetOwnerAccount sets the "owner" account.
func (inst *CreateAmmConfig) SetOwnerAccount(owner ag_solanago.PublicKey) *CreateAmmConfig {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(owner).WRITE().SIGNER()
	return inst
}

// GetOwnerAccount gets the "owner" account.
func (inst *CreateAmmConfig) GetOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetAmmConfigAccount sets the "amm_config" account.
func (inst *CreateAmmConfig) SetAmmConfigAccount(ammConfig ag_solanago.PublicKey) *CreateAmmConfig {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(ammConfig).WRITE()
	return inst
}

// GetAmmConfigAccount gets the "amm_config" account.
func (inst *CreateAmmConfig) GetAmmConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetSystemProgramAccount sets the "system_program" account.
func (inst *CreateAmmConfig) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *CreateAmmConfig {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "system_program" account.
func (inst *CreateAmmConfig) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

func (inst *CreateAmmConfig) SetAccounts(accounts []*ag_solanago.AccountMeta) error {
	inst.AccountMetaSlice = accounts
	return inst.Validate()
}

func (inst *CreateAmmConfig) SetRemainingAccounts(metas []*ag_solanago.AccountMeta) *CreateAmmConfig {
	inst.AccountMetaSlice = append(inst.AccountMetaSlice[0:3], metas...)
	return inst
}

func (inst *CreateAmmConfig) GetRemainingAccounts() []*ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3:]
}

func (inst CreateAmmConfig) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_CreateAmmConfig,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst CreateAmmConfig) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *CreateAmmConfig) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Index == nil {
			return errors.New("index parameter is not set")
		}
		if inst.TradeFeeRate == nil {
			return errors.New("tradeFeeRate parameter is not set")
		}
		if inst.ProtocolFeeRate == nil {
			return errors.New("protocolFeeRate parameter is not set")
		}
		if inst.FundFeeRate == nil {
			return errors.New("fundFeeRate parameter is not set")
		}
		if inst.CreatePoolFee == nil {
			return errors.New("createPoolFee parameter is not set")
		}
	}

	if len(inst.AccountMetaSlice) < 3 {
		return errors.New("accounts slice has wrong length: expected 3 accounts")
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Owner is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.AmmConfig is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *CreateAmmConfig) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("CreateAmmConfig")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=5]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("            Index", *inst.Index))
						paramsBranch.Child(ag_format.Param("     TradeFeeRate", *inst.TradeFeeRate))
						paramsBranch.Child(ag_format.Param("  ProtocolFeeRate", *inst.ProtocolFeeRate))
						paramsBranch.Child(ag_format.Param("      FundFeeRate", *inst.FundFeeRate))
						paramsBranch.Child(ag_format.Param("    CreatePoolFee", *inst.CreatePoolFee))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=3]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("         owner", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("    amm_config", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("system_program", inst.AccountMetaSlice.Get(2)))
					})
				})
		})
}

func (obj CreateAmmConfig) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Index` param:
	err = encoder.Encode(obj.Index)
	if err != nil {
		return err
	}
	// Serialize `TradeFeeRate` param:
	err = encoder.Encode(obj.TradeFeeRate)
	if err != nil {
		return err
	}
	// Serialize `ProtocolFeeRate` param:
	err = encoder.Encode(obj.ProtocolFeeRate)
	if err != nil {
		return err
	}
	// Serialize `FundFeeRate` param:
	err = encoder.Encode(obj.FundFeeRate)
	if err != nil {
		return err
	}
	// Serialize `CreatePoolFee` param:
	err = encoder.Encode(obj.CreatePoolFee)
	if err != nil {
		return err
	}
	return nil
}
func (obj *CreateAmmConfig) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Index`:
	err = decoder.Decode(&obj.Index)
	if err != nil {
		return err
	}
	// Deserialize `TradeFeeRate`:
	err = decoder.Decode(&obj.TradeFeeRate)
	if err != nil {
		return err
	}
	// Deserialize `ProtocolFeeRate`:
	err = decoder.Decode(&obj.ProtocolFeeRate)
	if err != nil {
		return err
	}
	// Deserialize `FundFeeRate`:
	err = decoder.Decode(&obj.FundFeeRate)
	if err != nil {
		return err
	}
	// Deserialize `CreatePoolFee`:
	err = decoder.Decode(&obj.CreatePoolFee)
	if err != nil {
		return err
	}
	return nil
}

// NewCreateAmmConfigInstruction declares a new CreateAmmConfig instruction with the provided parameters and accounts.
func NewCreateAmmConfigInstruction(
	// Parameters:
	index uint16,
	trade_fee_rate uint64,
	protocol_fee_rate uint64,
	fund_fee_rate uint64,
	create_pool_fee uint64,
	// Accounts:
	owner ag_solanago.PublicKey,
	ammConfig ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *CreateAmmConfig {
	return NewCreateAmmConfigInstructionBuilder().
		SetIndex(index).
		SetTradeFeeRate(trade_fee_rate).
		SetProtocolFeeRate(protocol_fee_rate).
		SetFundFeeRate(fund_fee_rate).
		SetCreatePoolFee(create_pool_fee).
		SetOwnerAccount(owner).
		SetAmmConfigAccount(ammConfig).
		SetSystemProgramAccount(systemProgram)
}

// NewSimpleCreateAmmConfigInstruction declares a new CreateAmmConfig instruction with the provided parameters and accounts.
func NewSimpleCreateAmmConfigInstruction(
	// Parameters:
	index uint16,
	trade_fee_rate uint64,
	protocol_fee_rate uint64,
	fund_fee_rate uint64,
	create_pool_fee uint64,
	// Accounts:
	owner ag_solanago.PublicKey,
	ammConfig ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *CreateAmmConfig {
	return NewCreateAmmConfigInstructionBuilder().
		SetIndex(index).
		SetTradeFeeRate(trade_fee_rate).
		SetProtocolFeeRate(protocol_fee_rate).
		SetFundFeeRate(fund_fee_rate).
		SetCreatePoolFee(create_pool_fee).
		SetOwnerAccount(owner).
		SetAmmConfigAccount(ammConfig).
		SetSystemProgramAccount(systemProgram)
}
