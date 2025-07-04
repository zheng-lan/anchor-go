// Code generated by https://github.com/zheng-lan/anchor-go. DO NOT EDIT.

package vr

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Create is the `create` instruction.
type Create struct {
	Params *CreateParams

	// [0] = [WRITE, SIGNER] payer
	//
	// [1] = [SIGNER] owner
	//
	// [2] = [SIGNER] token_wallet_authority
	//
	// [3] = [] mint_a
	//
	// [4] = [] mint_b
	//
	// [5] = [WRITE] token_wallet_b
	// ··········· Token wallet that funds the pool with token B.
	//
	// [6] = [WRITE] pool
	//
	// [7] = [WRITE] vault_a
	//
	// [8] = [WRITE] vault_b
	//
	// [9] = [] token_program_a
	// ··········· Required programs and sysvars
	//
	// [10] = [] token_program_b
	//
	// [11] = [] system_program
	//
	// [12] = [] rent
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewCreateInstructionBuilder creates a new `Create` instruction builder.
func NewCreateInstructionBuilder() *Create {
	nd := &Create{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 13),
	}
	nd.AccountMetaSlice[11] = ag_solanago.Meta(SystemProgram)
	nd.AccountMetaSlice[12] = ag_solanago.Meta(RentProgram)
	return nd
}

// SetParams sets the "params" parameter.
func (inst *Create) SetParams(params CreateParams) *Create {
	inst.Params = &params
	return inst
}

// SetPayerAccount sets the "payer" account.
func (inst *Create) SetPayerAccount(payer ag_solanago.PublicKey) *Create {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(payer).WRITE().SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
func (inst *Create) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetOwnerAccount sets the "owner" account.
func (inst *Create) SetOwnerAccount(owner ag_solanago.PublicKey) *Create {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(owner).SIGNER()
	return inst
}

// GetOwnerAccount gets the "owner" account.
func (inst *Create) GetOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetTokenWalletAuthorityAccount sets the "token_wallet_authority" account.
func (inst *Create) SetTokenWalletAuthorityAccount(tokenWalletAuthority ag_solanago.PublicKey) *Create {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(tokenWalletAuthority).SIGNER()
	return inst
}

// GetTokenWalletAuthorityAccount gets the "token_wallet_authority" account.
func (inst *Create) GetTokenWalletAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetMintAAccount sets the "mint_a" account.
func (inst *Create) SetMintAAccount(mintA ag_solanago.PublicKey) *Create {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(mintA)
	return inst
}

// GetMintAAccount gets the "mint_a" account.
func (inst *Create) GetMintAAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetMintBAccount sets the "mint_b" account.
func (inst *Create) SetMintBAccount(mintB ag_solanago.PublicKey) *Create {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(mintB)
	return inst
}

// GetMintBAccount gets the "mint_b" account.
func (inst *Create) GetMintBAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetTokenWalletBAccount sets the "token_wallet_b" account.
// Token wallet that funds the pool with token B.
func (inst *Create) SetTokenWalletBAccount(tokenWalletB ag_solanago.PublicKey) *Create {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(tokenWalletB).WRITE()
	return inst
}

// GetTokenWalletBAccount gets the "token_wallet_b" account.
// Token wallet that funds the pool with token B.
func (inst *Create) GetTokenWalletBAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetPoolAccount sets the "pool" account.
func (inst *Create) SetPoolAccount(pool ag_solanago.PublicKey) *Create {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(pool).WRITE()
	return inst
}

// GetPoolAccount gets the "pool" account.
func (inst *Create) GetPoolAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetVaultAAccount sets the "vault_a" account.
func (inst *Create) SetVaultAAccount(vaultA ag_solanago.PublicKey) *Create {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(vaultA).WRITE()
	return inst
}

// GetVaultAAccount gets the "vault_a" account.
func (inst *Create) GetVaultAAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetVaultBAccount sets the "vault_b" account.
func (inst *Create) SetVaultBAccount(vaultB ag_solanago.PublicKey) *Create {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(vaultB).WRITE()
	return inst
}

// GetVaultBAccount gets the "vault_b" account.
func (inst *Create) GetVaultBAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetTokenProgramAAccount sets the "token_program_a" account.
// Required programs and sysvars
func (inst *Create) SetTokenProgramAAccount(tokenProgramA ag_solanago.PublicKey) *Create {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(tokenProgramA)
	return inst
}

// GetTokenProgramAAccount gets the "token_program_a" account.
// Required programs and sysvars
func (inst *Create) GetTokenProgramAAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetTokenProgramBAccount sets the "token_program_b" account.
func (inst *Create) SetTokenProgramBAccount(tokenProgramB ag_solanago.PublicKey) *Create {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(tokenProgramB)
	return inst
}

// GetTokenProgramBAccount gets the "token_program_b" account.
func (inst *Create) GetTokenProgramBAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetSystemProgramAccount sets the "system_program" account.
func (inst *Create) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *Create {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "system_program" account.
func (inst *Create) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetRentAccount sets the "rent" account.
func (inst *Create) SetRentAccount(rent ag_solanago.PublicKey) *Create {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
func (inst *Create) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

func (inst *Create) SetAccounts(accounts []*ag_solanago.AccountMeta) error {
	inst.AccountMetaSlice = accounts
	return inst.Validate()
}

func (inst *Create) SetRemainingAccounts(metas []*ag_solanago.AccountMeta) *Create {
	inst.AccountMetaSlice = append(inst.AccountMetaSlice[0:13], metas...)
	return inst
}

func (inst *Create) GetRemainingAccounts() []*ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[13:]
}

func (inst Create) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_Create,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst Create) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *Create) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Params == nil {
			return errors.New("params parameter is not set")
		}
	}

	if len(inst.AccountMetaSlice) < 13 {
		return errors.New("accounts slice has wrong length: expected 13 accounts")
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Owner is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.TokenWalletAuthority is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.MintA is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.MintB is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.TokenWalletB is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.Pool is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.VaultA is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.VaultB is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.TokenProgramA is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.TokenProgramB is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.Rent is not set")
		}
	}
	return nil
}

func (inst *Create) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("Create")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Params", *inst.Params))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=13]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                 payer", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("                 owner", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("token_wallet_authority", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("                mint_a", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("                mint_b", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("        token_wallet_b", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("                  pool", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("               vault_a", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("               vault_b", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("       token_program_a", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("       token_program_b", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("        system_program", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("                  rent", inst.AccountMetaSlice.Get(12)))
					})
				})
		})
}

func (obj Create) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Params` param:
	err = encoder.Encode(obj.Params)
	if err != nil {
		return err
	}
	return nil
}
func (obj *Create) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Params`:
	err = decoder.Decode(&obj.Params)
	if err != nil {
		return err
	}
	return nil
}

// NewCreateInstruction declares a new Create instruction with the provided parameters and accounts.
func NewCreateInstruction(
	// Parameters:
	params CreateParams,
	// Accounts:
	payer ag_solanago.PublicKey,
	owner ag_solanago.PublicKey,
	tokenWalletAuthority ag_solanago.PublicKey,
	mintA ag_solanago.PublicKey,
	mintB ag_solanago.PublicKey,
	tokenWalletB ag_solanago.PublicKey,
	pool ag_solanago.PublicKey,
	vaultA ag_solanago.PublicKey,
	vaultB ag_solanago.PublicKey,
	tokenProgramA ag_solanago.PublicKey,
	tokenProgramB ag_solanago.PublicKey) *Create {
	return NewCreateInstructionBuilder().
		SetParams(params).
		SetPayerAccount(payer).
		SetOwnerAccount(owner).
		SetTokenWalletAuthorityAccount(tokenWalletAuthority).
		SetMintAAccount(mintA).
		SetMintBAccount(mintB).
		SetTokenWalletBAccount(tokenWalletB).
		SetPoolAccount(pool).
		SetVaultAAccount(vaultA).
		SetVaultBAccount(vaultB).
		SetTokenProgramAAccount(tokenProgramA).
		SetTokenProgramBAccount(tokenProgramB)
}

// NewSimpleCreateInstruction declares a new Create instruction with the provided parameters and accounts.
func NewSimpleCreateInstruction(
	// Parameters:
	params CreateParams,
	// Accounts:
	payer ag_solanago.PublicKey,
	owner ag_solanago.PublicKey,
	tokenWalletAuthority ag_solanago.PublicKey,
	mintA ag_solanago.PublicKey,
	mintB ag_solanago.PublicKey,
	tokenWalletB ag_solanago.PublicKey,
	tokenProgramA ag_solanago.PublicKey,
	tokenProgramB ag_solanago.PublicKey) *Create {
	pool := MustFindPoolAddress(owner, mintA, mintB)
	vaultA := MustFindVaultAAddress(pool, mintA)
	vaultB := MustFindVaultBAddress(pool, mintB)
	return NewCreateInstructionBuilder().
		SetParams(params).
		SetPayerAccount(payer).
		SetOwnerAccount(owner).
		SetTokenWalletAuthorityAccount(tokenWalletAuthority).
		SetMintAAccount(mintA).
		SetMintBAccount(mintB).
		SetTokenWalletBAccount(tokenWalletB).
		SetPoolAccount(pool).
		SetVaultAAccount(vaultA).
		SetVaultBAccount(vaultB).
		SetTokenProgramAAccount(tokenProgramA).
		SetTokenProgramBAccount(tokenProgramB)
}
