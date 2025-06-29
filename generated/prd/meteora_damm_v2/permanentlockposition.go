// Code generated by https://github.com/zheng-lan/anchor-go. DO NOT EDIT.

package meteora_damm_v2

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// PermanentLockPosition is the `permanent_lock_position` instruction.
type PermanentLockPosition struct {
	PermanentLockLiquidity *ag_binary.Uint128

	// [0] = [WRITE] pool
	//
	// [1] = [WRITE] position
	//
	// [2] = [] position_nft_account
	// ··········· The token account for nft
	//
	// [3] = [SIGNER] owner
	// ··········· owner of position
	//
	// [4] = [] event_authority
	//
	// [5] = [] program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewPermanentLockPositionInstructionBuilder creates a new `PermanentLockPosition` instruction builder.
func NewPermanentLockPositionInstructionBuilder() *PermanentLockPosition {
	nd := &PermanentLockPosition{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 6),
	}
	nd.AccountMetaSlice[4] = ag_solanago.Meta(EventAuthorityPDA)
	nd.AccountMetaSlice[5] = ag_solanago.Meta(ProgramID)
	return nd
}

// SetPermanentLockLiquidity sets the "permanent_lock_liquidity" parameter.
func (inst *PermanentLockPosition) SetPermanentLockLiquidity(permanent_lock_liquidity ag_binary.Uint128) *PermanentLockPosition {
	inst.PermanentLockLiquidity = &permanent_lock_liquidity
	return inst
}

// SetPoolAccount sets the "pool" account.
func (inst *PermanentLockPosition) SetPoolAccount(pool ag_solanago.PublicKey) *PermanentLockPosition {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(pool).WRITE()
	return inst
}

// GetPoolAccount gets the "pool" account.
func (inst *PermanentLockPosition) GetPoolAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetPositionAccount sets the "position" account.
func (inst *PermanentLockPosition) SetPositionAccount(position ag_solanago.PublicKey) *PermanentLockPosition {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(position).WRITE()
	return inst
}

// GetPositionAccount gets the "position" account.
func (inst *PermanentLockPosition) GetPositionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetPositionNftAccountAccount sets the "position_nft_account" account.
// The token account for nft
func (inst *PermanentLockPosition) SetPositionNftAccountAccount(positionNftAccount ag_solanago.PublicKey) *PermanentLockPosition {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(positionNftAccount)
	return inst
}

// GetPositionNftAccountAccount gets the "position_nft_account" account.
// The token account for nft
func (inst *PermanentLockPosition) GetPositionNftAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetOwnerAccount sets the "owner" account.
// owner of position
func (inst *PermanentLockPosition) SetOwnerAccount(owner ag_solanago.PublicKey) *PermanentLockPosition {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(owner).SIGNER()
	return inst
}

// GetOwnerAccount gets the "owner" account.
// owner of position
func (inst *PermanentLockPosition) GetOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetEventAuthorityAccount sets the "event_authority" account.
func (inst *PermanentLockPosition) SetEventAuthorityAccount(eventAuthority ag_solanago.PublicKey) *PermanentLockPosition {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(eventAuthority)
	return inst
}

// GetEventAuthorityAccount gets the "event_authority" account.
func (inst *PermanentLockPosition) GetEventAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetProgramAccount sets the "program" account.
func (inst *PermanentLockPosition) SetProgramAccount(program ag_solanago.PublicKey) *PermanentLockPosition {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(program)
	return inst
}

// GetProgramAccount gets the "program" account.
func (inst *PermanentLockPosition) GetProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

func (inst *PermanentLockPosition) SetAccounts(accounts []*ag_solanago.AccountMeta) error {
	inst.AccountMetaSlice = accounts
	return inst.Validate()
}

func (inst *PermanentLockPosition) SetRemainingAccounts(metas []*ag_solanago.AccountMeta) *PermanentLockPosition {
	inst.AccountMetaSlice = append(inst.AccountMetaSlice[0:6], metas...)
	return inst
}

func (inst *PermanentLockPosition) GetRemainingAccounts() []*ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[6:]
}

func (inst PermanentLockPosition) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_PermanentLockPosition,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst PermanentLockPosition) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *PermanentLockPosition) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.PermanentLockLiquidity == nil {
			return errors.New("permanentLockLiquidity parameter is not set")
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
			return errors.New("accounts.Position is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.PositionNftAccount is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Owner is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.EventAuthority is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.Program is not set")
		}
	}
	return nil
}

func (inst *PermanentLockPosition) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("PermanentLockPosition")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("  PermanentLockLiquidity", *inst.PermanentLockLiquidity))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=6]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("           pool", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("       position", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("  position_nft_", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("          owner", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("event_authority", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("        program", inst.AccountMetaSlice.Get(5)))
					})
				})
		})
}

func (obj PermanentLockPosition) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `PermanentLockLiquidity` param:
	err = encoder.Encode(obj.PermanentLockLiquidity)
	if err != nil {
		return err
	}
	return nil
}
func (obj *PermanentLockPosition) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `PermanentLockLiquidity`:
	err = decoder.Decode(&obj.PermanentLockLiquidity)
	if err != nil {
		return err
	}
	return nil
}

// NewPermanentLockPositionInstruction declares a new PermanentLockPosition instruction with the provided parameters and accounts.
func NewPermanentLockPositionInstruction(
	// Parameters:
	permanent_lock_liquidity ag_binary.Uint128,
	// Accounts:
	pool ag_solanago.PublicKey,
	position ag_solanago.PublicKey,
	positionNftAccount ag_solanago.PublicKey,
	owner ag_solanago.PublicKey) *PermanentLockPosition {
	return NewPermanentLockPositionInstructionBuilder().
		SetPermanentLockLiquidity(permanent_lock_liquidity).
		SetPoolAccount(pool).
		SetPositionAccount(position).
		SetPositionNftAccountAccount(positionNftAccount).
		SetOwnerAccount(owner)
}

// NewSimplePermanentLockPositionInstruction declares a new PermanentLockPosition instruction with the provided parameters and accounts.
func NewSimplePermanentLockPositionInstruction(
	// Parameters:
	permanent_lock_liquidity ag_binary.Uint128,
	// Accounts:
	pool ag_solanago.PublicKey,
	position ag_solanago.PublicKey,
	positionNftAccount ag_solanago.PublicKey,
	owner ag_solanago.PublicKey) *PermanentLockPosition {
	return NewPermanentLockPositionInstructionBuilder().
		SetPermanentLockLiquidity(permanent_lock_liquidity).
		SetPoolAccount(pool).
		SetPositionAccount(position).
		SetPositionNftAccountAccount(positionNftAccount).
		SetOwnerAccount(owner)
}
