// Code generated by https://github.com/zheng-lan/anchor-go. DO NOT EDIT.

package orca_whirlpool

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Transfer a locked position to to a different token account.
//
// ### Authority
// - `position_authority` - The authority that owns the position token.
type TransferLockedPosition struct {

	// [0] = [SIGNER] position_authority
	//
	// [1] = [WRITE] receiver
	//
	// [2] = [] position
	//
	// [3] = [] position_mint
	//
	// [4] = [WRITE] position_token_account
	//
	// [5] = [WRITE] destination_token_account
	//
	// [6] = [WRITE] lock_config
	//
	// [7] = [] token2022_program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewTransferLockedPositionInstructionBuilder creates a new `TransferLockedPosition` instruction builder.
func NewTransferLockedPositionInstructionBuilder() *TransferLockedPosition {
	nd := &TransferLockedPosition{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 8),
	}
	return nd
}

// SetPositionAuthorityAccount sets the "position_authority" account.
func (inst *TransferLockedPosition) SetPositionAuthorityAccount(positionAuthority ag_solanago.PublicKey) *TransferLockedPosition {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(positionAuthority).SIGNER()
	return inst
}

// GetPositionAuthorityAccount gets the "position_authority" account.
func (inst *TransferLockedPosition) GetPositionAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetReceiverAccount sets the "receiver" account.
func (inst *TransferLockedPosition) SetReceiverAccount(receiver ag_solanago.PublicKey) *TransferLockedPosition {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(receiver).WRITE()
	return inst
}

// GetReceiverAccount gets the "receiver" account.
func (inst *TransferLockedPosition) GetReceiverAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetPositionAccount sets the "position" account.
func (inst *TransferLockedPosition) SetPositionAccount(position ag_solanago.PublicKey) *TransferLockedPosition {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(position)
	return inst
}

// GetPositionAccount gets the "position" account.
func (inst *TransferLockedPosition) GetPositionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetPositionMintAccount sets the "position_mint" account.
func (inst *TransferLockedPosition) SetPositionMintAccount(positionMint ag_solanago.PublicKey) *TransferLockedPosition {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(positionMint)
	return inst
}

// GetPositionMintAccount gets the "position_mint" account.
func (inst *TransferLockedPosition) GetPositionMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetPositionTokenAccountAccount sets the "position_token_account" account.
func (inst *TransferLockedPosition) SetPositionTokenAccountAccount(positionTokenAccount ag_solanago.PublicKey) *TransferLockedPosition {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(positionTokenAccount).WRITE()
	return inst
}

// GetPositionTokenAccountAccount gets the "position_token_account" account.
func (inst *TransferLockedPosition) GetPositionTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetDestinationTokenAccountAccount sets the "destination_token_account" account.
func (inst *TransferLockedPosition) SetDestinationTokenAccountAccount(destinationTokenAccount ag_solanago.PublicKey) *TransferLockedPosition {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(destinationTokenAccount).WRITE()
	return inst
}

// GetDestinationTokenAccountAccount gets the "destination_token_account" account.
func (inst *TransferLockedPosition) GetDestinationTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetLockConfigAccount sets the "lock_config" account.
func (inst *TransferLockedPosition) SetLockConfigAccount(lockConfig ag_solanago.PublicKey) *TransferLockedPosition {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(lockConfig).WRITE()
	return inst
}

// GetLockConfigAccount gets the "lock_config" account.
func (inst *TransferLockedPosition) GetLockConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetToken2022ProgramAccount sets the "token2022_program" account.
func (inst *TransferLockedPosition) SetToken2022ProgramAccount(token2022Program ag_solanago.PublicKey) *TransferLockedPosition {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(token2022Program)
	return inst
}

// GetToken2022ProgramAccount gets the "token2022_program" account.
func (inst *TransferLockedPosition) GetToken2022ProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

func (inst *TransferLockedPosition) SetAccounts(accounts []*ag_solanago.AccountMeta) error {
	inst.AccountMetaSlice = accounts
	return inst.Validate()
}

func (inst *TransferLockedPosition) SetRemainingAccounts(metas []*ag_solanago.AccountMeta) *TransferLockedPosition {
	inst.AccountMetaSlice = append(inst.AccountMetaSlice[0:8], metas...)
	return inst
}

func (inst *TransferLockedPosition) GetRemainingAccounts() []*ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[8:]
}

func (inst TransferLockedPosition) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_TransferLockedPosition,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst TransferLockedPosition) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *TransferLockedPosition) Validate() error {
	if len(inst.AccountMetaSlice) < 8 {
		return errors.New("accounts slice has wrong length: expected 8 accounts")
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.PositionAuthority is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Receiver is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Position is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.PositionMint is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.PositionTokenAccount is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.DestinationTokenAccount is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.LockConfig is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.Token2022Program is not set")
		}
	}
	return nil
}

func (inst *TransferLockedPosition) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("TransferLockedPosition")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=8]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("position_authority", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("          receiver", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("          position", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("     position_mint", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("   position_token_", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("destination_token_", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("       lock_config", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta(" token2022_program", inst.AccountMetaSlice.Get(7)))
					})
				})
		})
}

func (obj TransferLockedPosition) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *TransferLockedPosition) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewTransferLockedPositionInstruction declares a new TransferLockedPosition instruction with the provided parameters and accounts.
func NewTransferLockedPositionInstruction(
	// Accounts:
	positionAuthority ag_solanago.PublicKey,
	receiver ag_solanago.PublicKey,
	position ag_solanago.PublicKey,
	positionMint ag_solanago.PublicKey,
	positionTokenAccount ag_solanago.PublicKey,
	destinationTokenAccount ag_solanago.PublicKey,
	lockConfig ag_solanago.PublicKey,
	token2022Program ag_solanago.PublicKey) *TransferLockedPosition {
	return NewTransferLockedPositionInstructionBuilder().
		SetPositionAuthorityAccount(positionAuthority).
		SetReceiverAccount(receiver).
		SetPositionAccount(position).
		SetPositionMintAccount(positionMint).
		SetPositionTokenAccountAccount(positionTokenAccount).
		SetDestinationTokenAccountAccount(destinationTokenAccount).
		SetLockConfigAccount(lockConfig).
		SetToken2022ProgramAccount(token2022Program)
}

// NewSimpleTransferLockedPositionInstruction declares a new TransferLockedPosition instruction with the provided parameters and accounts.
func NewSimpleTransferLockedPositionInstruction(
	// Accounts:
	positionAuthority ag_solanago.PublicKey,
	receiver ag_solanago.PublicKey,
	position ag_solanago.PublicKey,
	positionMint ag_solanago.PublicKey,
	positionTokenAccount ag_solanago.PublicKey,
	destinationTokenAccount ag_solanago.PublicKey,
	lockConfig ag_solanago.PublicKey,
	token2022Program ag_solanago.PublicKey) *TransferLockedPosition {
	return NewTransferLockedPositionInstructionBuilder().
		SetPositionAuthorityAccount(positionAuthority).
		SetReceiverAccount(receiver).
		SetPositionAccount(position).
		SetPositionMintAccount(positionMint).
		SetPositionTokenAccountAccount(positionTokenAccount).
		SetDestinationTokenAccountAccount(destinationTokenAccount).
		SetLockConfigAccount(lockConfig).
		SetToken2022ProgramAccount(token2022Program)
}
