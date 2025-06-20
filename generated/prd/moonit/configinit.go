// Code generated by https://github.com/zheng-lan/anchor-go. DO NOT EDIT.

package moonit

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// ConfigInit is the `config_init` instruction.
type ConfigInit struct {
	Data *ConfigParams

	// [0] = [WRITE, SIGNER] config_authority
	//
	// [1] = [WRITE] config_account
	//
	// [2] = [] system_program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewConfigInitInstructionBuilder creates a new `ConfigInit` instruction builder.
func NewConfigInitInstructionBuilder() *ConfigInit {
	nd := &ConfigInit{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetData sets the "data" parameter.
func (inst *ConfigInit) SetData(data ConfigParams) *ConfigInit {
	inst.Data = &data
	return inst
}

// SetConfigAuthorityAccount sets the "config_authority" account.
func (inst *ConfigInit) SetConfigAuthorityAccount(configAuthority ag_solanago.PublicKey) *ConfigInit {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(configAuthority).WRITE().SIGNER()
	return inst
}

// GetConfigAuthorityAccount gets the "config_authority" account.
func (inst *ConfigInit) GetConfigAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetConfigAccountAccount sets the "config_account" account.
func (inst *ConfigInit) SetConfigAccountAccount(configAccount ag_solanago.PublicKey) *ConfigInit {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(configAccount).WRITE()
	return inst
}

// GetConfigAccountAccount gets the "config_account" account.
func (inst *ConfigInit) GetConfigAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetSystemProgramAccount sets the "system_program" account.
func (inst *ConfigInit) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *ConfigInit {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "system_program" account.
func (inst *ConfigInit) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

func (inst *ConfigInit) SetAccounts(accounts []*ag_solanago.AccountMeta) error {
	inst.AccountMetaSlice = accounts
	return inst.Validate()
}

func (inst *ConfigInit) SetRemainingAccounts(metas []*ag_solanago.AccountMeta) *ConfigInit {
	inst.AccountMetaSlice = append(inst.AccountMetaSlice[0:3], metas...)
	return inst
}

func (inst *ConfigInit) GetRemainingAccounts() []*ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3:]
}

func (inst ConfigInit) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_ConfigInit,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst ConfigInit) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *ConfigInit) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Data == nil {
			return errors.New("data parameter is not set")
		}
	}

	if len(inst.AccountMetaSlice) < 3 {
		return errors.New("accounts slice has wrong length: expected 3 accounts")
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.ConfigAuthority is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.ConfigAccount is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *ConfigInit) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("ConfigInit")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Data", *inst.Data))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=3]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("config_authority", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("         config_", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("  system_program", inst.AccountMetaSlice.Get(2)))
					})
				})
		})
}

func (obj ConfigInit) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Data` param:
	err = encoder.Encode(obj.Data)
	if err != nil {
		return err
	}
	return nil
}
func (obj *ConfigInit) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Data`:
	err = decoder.Decode(&obj.Data)
	if err != nil {
		return err
	}
	return nil
}

// NewConfigInitInstruction declares a new ConfigInit instruction with the provided parameters and accounts.
func NewConfigInitInstruction(
	// Parameters:
	data ConfigParams,
	// Accounts:
	configAuthority ag_solanago.PublicKey,
	configAccount ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *ConfigInit {
	return NewConfigInitInstructionBuilder().
		SetData(data).
		SetConfigAuthorityAccount(configAuthority).
		SetConfigAccountAccount(configAccount).
		SetSystemProgramAccount(systemProgram)
}

// NewSimpleConfigInitInstruction declares a new ConfigInit instruction with the provided parameters and accounts.
func NewSimpleConfigInitInstruction(
	// Parameters:
	data ConfigParams,
	// Accounts:
	configAuthority ag_solanago.PublicKey,
	configAccount ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *ConfigInit {
	return NewConfigInitInstructionBuilder().
		SetData(data).
		SetConfigAuthorityAccount(configAuthority).
		SetConfigAccountAccount(configAccount).
		SetSystemProgramAccount(systemProgram)
}
