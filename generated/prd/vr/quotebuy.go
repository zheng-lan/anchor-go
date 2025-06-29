// Code generated by https://github.com/zheng-lan/anchor-go. DO NOT EDIT.

package vr

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// QuoteBuy is the `quote_buy` instruction.
type QuoteBuy struct {
	Params *SwapParams

	// [0] = [] pool
	//
	// [1] = [] owner
	//
	// [2] = [] user
	//
	// [3] = [] mint_a
	//
	// [4] = [] mint_b
	//
	// [5] = [] program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewQuoteBuyInstructionBuilder creates a new `QuoteBuy` instruction builder.
func NewQuoteBuyInstructionBuilder() *QuoteBuy {
	nd := &QuoteBuy{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 6),
	}
	nd.AccountMetaSlice[5] = ag_solanago.Meta(ProgramID)
	return nd
}

// SetParams sets the "params" parameter.
func (inst *QuoteBuy) SetParams(params SwapParams) *QuoteBuy {
	inst.Params = &params
	return inst
}

// SetPoolAccount sets the "pool" account.
func (inst *QuoteBuy) SetPoolAccount(pool ag_solanago.PublicKey) *QuoteBuy {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(pool)
	return inst
}

// GetPoolAccount gets the "pool" account.
func (inst *QuoteBuy) GetPoolAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetOwnerAccount sets the "owner" account.
func (inst *QuoteBuy) SetOwnerAccount(owner ag_solanago.PublicKey) *QuoteBuy {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(owner)
	return inst
}

// GetOwnerAccount gets the "owner" account.
func (inst *QuoteBuy) GetOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetUserAccount sets the "user" account.
func (inst *QuoteBuy) SetUserAccount(user ag_solanago.PublicKey) *QuoteBuy {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(user)
	return inst
}

// GetUserAccount gets the "user" account.
func (inst *QuoteBuy) GetUserAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetMintAAccount sets the "mint_a" account.
func (inst *QuoteBuy) SetMintAAccount(mintA ag_solanago.PublicKey) *QuoteBuy {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(mintA)
	return inst
}

// GetMintAAccount gets the "mint_a" account.
func (inst *QuoteBuy) GetMintAAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetMintBAccount sets the "mint_b" account.
func (inst *QuoteBuy) SetMintBAccount(mintB ag_solanago.PublicKey) *QuoteBuy {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(mintB)
	return inst
}

// GetMintBAccount gets the "mint_b" account.
func (inst *QuoteBuy) GetMintBAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetProgramAccount sets the "program" account.
func (inst *QuoteBuy) SetProgramAccount(program ag_solanago.PublicKey) *QuoteBuy {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(program)
	return inst
}

// GetProgramAccount gets the "program" account.
func (inst *QuoteBuy) GetProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

func (inst *QuoteBuy) SetAccounts(accounts []*ag_solanago.AccountMeta) error {
	inst.AccountMetaSlice = accounts
	return inst.Validate()
}

func (inst *QuoteBuy) SetRemainingAccounts(metas []*ag_solanago.AccountMeta) *QuoteBuy {
	inst.AccountMetaSlice = append(inst.AccountMetaSlice[0:6], metas...)
	return inst
}

func (inst *QuoteBuy) GetRemainingAccounts() []*ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[6:]
}

func (inst QuoteBuy) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_QuoteBuy,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst QuoteBuy) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *QuoteBuy) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Params == nil {
			return errors.New("params parameter is not set")
		}
	}

	if len(inst.AccountMetaSlice) < 6 {
		return errors.New("accounts slice has wrong length: expected 6 accounts")
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Pool is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Owner is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.User is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.MintA is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.MintB is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.Program is not set")
		}
	}
	return nil
}

func (inst *QuoteBuy) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("QuoteBuy")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Params", *inst.Params))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=6]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("   pool", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("  owner", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("   user", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta(" mint_a", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta(" mint_b", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("program", inst.AccountMetaSlice.Get(5)))
					})
				})
		})
}

func (obj QuoteBuy) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Params` param:
	err = encoder.Encode(obj.Params)
	if err != nil {
		return err
	}
	return nil
}
func (obj *QuoteBuy) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Params`:
	err = decoder.Decode(&obj.Params)
	if err != nil {
		return err
	}
	return nil
}

// NewQuoteBuyInstruction declares a new QuoteBuy instruction with the provided parameters and accounts.
func NewQuoteBuyInstruction(
	// Parameters:
	params SwapParams,
	// Accounts:
	pool ag_solanago.PublicKey,
	owner ag_solanago.PublicKey,
	user ag_solanago.PublicKey,
	mintA ag_solanago.PublicKey,
	mintB ag_solanago.PublicKey) *QuoteBuy {
	return NewQuoteBuyInstructionBuilder().
		SetParams(params).
		SetPoolAccount(pool).
		SetOwnerAccount(owner).
		SetUserAccount(user).
		SetMintAAccount(mintA).
		SetMintBAccount(mintB)
}

// NewSimpleQuoteBuyInstruction declares a new QuoteBuy instruction with the provided parameters and accounts.
func NewSimpleQuoteBuyInstruction(
	// Parameters:
	params SwapParams,
	owner ag_solanago.PublicKey,
	user ag_solanago.PublicKey,
	mintA ag_solanago.PublicKey,
	mintB ag_solanago.PublicKey) *QuoteBuy {
	pool := MustFindPoolAddress(owner, mintA, mintB)
	return NewQuoteBuyInstructionBuilder().
		SetParams(params).
		SetPoolAccount(pool).
		SetOwnerAccount(owner).
		SetUserAccount(user).
		SetMintAAccount(mintA).
		SetMintBAccount(mintB)
}
