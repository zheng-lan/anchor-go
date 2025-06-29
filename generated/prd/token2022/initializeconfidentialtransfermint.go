// Code generated by https://github.com/zheng-lan/anchor-go. DO NOT EDIT.

package token2022

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// InitializeConfidentialTransferMint is the `initialize_confidential_transfer_mint` instruction.
type InitializeConfidentialTransferMint struct {

	// [0] = [WRITE] mint
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewInitializeConfidentialTransferMintInstructionBuilder creates a new `InitializeConfidentialTransferMint` instruction builder.
func NewInitializeConfidentialTransferMintInstructionBuilder() *InitializeConfidentialTransferMint {
	nd := &InitializeConfidentialTransferMint{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 1),
	}
	return nd
}

// SetMintAccount sets the "mint" account.
func (inst *InitializeConfidentialTransferMint) SetMintAccount(mint ag_solanago.PublicKey) *InitializeConfidentialTransferMint {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(mint).WRITE()
	return inst
}

// GetMintAccount gets the "mint" account.
func (inst *InitializeConfidentialTransferMint) GetMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

func (inst *InitializeConfidentialTransferMint) SetAccounts(accounts []*ag_solanago.AccountMeta) error {
	inst.AccountMetaSlice = accounts
	return inst.Validate()
}

func (inst *InitializeConfidentialTransferMint) SetRemainingAccounts(metas []*ag_solanago.AccountMeta) *InitializeConfidentialTransferMint {
	inst.AccountMetaSlice = append(inst.AccountMetaSlice[0:1], metas...)
	return inst
}

func (inst *InitializeConfidentialTransferMint) GetRemainingAccounts() []*ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1:]
}

func (inst InitializeConfidentialTransferMint) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_InitializeConfidentialTransferMint),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst InitializeConfidentialTransferMint) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *InitializeConfidentialTransferMint) Validate() error {
	if len(inst.AccountMetaSlice) < 1 {
		return errors.New("accounts slice has wrong length: expected 1 accounts")
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Mint is not set")
		}
	}
	return nil
}

func (inst *InitializeConfidentialTransferMint) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("InitializeConfidentialTransferMint")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=1]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("mint", inst.AccountMetaSlice.Get(0)))
					})
				})
		})
}

func (obj InitializeConfidentialTransferMint) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *InitializeConfidentialTransferMint) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewInitializeConfidentialTransferMintInstruction declares a new InitializeConfidentialTransferMint instruction with the provided parameters and accounts.
func NewInitializeConfidentialTransferMintInstruction(
	// Accounts:
	mint ag_solanago.PublicKey) *InitializeConfidentialTransferMint {
	return NewInitializeConfidentialTransferMintInstructionBuilder().
		SetMintAccount(mint)
}

// NewSimpleInitializeConfidentialTransferMintInstruction declares a new InitializeConfidentialTransferMint instruction with the provided parameters and accounts.
func NewSimpleInitializeConfidentialTransferMintInstruction(
	// Accounts:
	mint ag_solanago.PublicKey) *InitializeConfidentialTransferMint {
	return NewInitializeConfidentialTransferMintInstructionBuilder().
		SetMintAccount(mint)
}
