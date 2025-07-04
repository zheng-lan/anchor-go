// Code generated by https://github.com/zheng-lan/anchor-go. DO NOT EDIT.

package meteora_pools

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Deposit tokens to the pool in a balanced ratio.
type AddBalanceLiquidity struct {
	PoolTokenAmount     *uint64
	MaximumTokenAAmount *uint64
	MaximumTokenBAmount *uint64

	// [0] = [WRITE] pool
	// ··········· Pool account (PDA)
	//
	// [1] = [WRITE] lp_mint
	// ··········· LP token mint of the pool
	//
	// [2] = [WRITE] user_pool_lp
	// ··········· user pool lp token account. lp will be burned from this account upon success liquidity removal.
	//
	// [3] = [WRITE] a_vault_lp
	// ··········· LP token account of vault A. Used to receive/burn the vault LP upon deposit/withdraw from the vault.
	//
	// [4] = [WRITE] b_vault_lp
	// ··········· LP token account of vault B. Used to receive/burn the vault LP upon deposit/withdraw from the vault.
	//
	// [5] = [WRITE] a_vault
	// ··········· Vault account for token a. token a of the pool will be deposit / withdraw from this vault account.
	//
	// [6] = [WRITE] b_vault
	// ··········· Vault account for token b. token b of the pool will be deposit / withdraw from this vault account.
	//
	// [7] = [WRITE] a_vault_lp_mint
	// ··········· LP token mint of vault a
	//
	// [8] = [WRITE] b_vault_lp_mint
	// ··········· LP token mint of vault b
	//
	// [9] = [WRITE] a_token_vault
	// ··········· Token vault account of vault A
	//
	// [10] = [WRITE] b_token_vault
	// ··········· Token vault account of vault B
	//
	// [11] = [WRITE] user_a_token
	// ··········· User token A account. Token will be transfer from this account if it is add liquidity operation. Else, token will be transfer into this account.
	//
	// [12] = [WRITE] user_b_token
	// ··········· User token B account. Token will be transfer from this account if it is add liquidity operation. Else, token will be transfer into this account.
	//
	// [13] = [SIGNER] user
	// ··········· User account. Must be owner of user_a_token, and user_b_token.
	//
	// [14] = [] vault_program
	// ··········· Vault program. the pool will deposit/withdraw liquidity from the vault.
	//
	// [15] = [] token_program
	// ··········· Token program.
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewAddBalanceLiquidityInstructionBuilder creates a new `AddBalanceLiquidity` instruction builder.
func NewAddBalanceLiquidityInstructionBuilder() *AddBalanceLiquidity {
	nd := &AddBalanceLiquidity{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 16),
	}
	return nd
}

// SetPoolTokenAmount sets the "pool_token_amount" parameter.
func (inst *AddBalanceLiquidity) SetPoolTokenAmount(pool_token_amount uint64) *AddBalanceLiquidity {
	inst.PoolTokenAmount = &pool_token_amount
	return inst
}

// SetMaximumTokenAAmount sets the "maximum_token_a_amount" parameter.
func (inst *AddBalanceLiquidity) SetMaximumTokenAAmount(maximum_token_a_amount uint64) *AddBalanceLiquidity {
	inst.MaximumTokenAAmount = &maximum_token_a_amount
	return inst
}

// SetMaximumTokenBAmount sets the "maximum_token_b_amount" parameter.
func (inst *AddBalanceLiquidity) SetMaximumTokenBAmount(maximum_token_b_amount uint64) *AddBalanceLiquidity {
	inst.MaximumTokenBAmount = &maximum_token_b_amount
	return inst
}

// SetPoolAccount sets the "pool" account.
// Pool account (PDA)
func (inst *AddBalanceLiquidity) SetPoolAccount(pool ag_solanago.PublicKey) *AddBalanceLiquidity {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(pool).WRITE()
	return inst
}

// GetPoolAccount gets the "pool" account.
// Pool account (PDA)
func (inst *AddBalanceLiquidity) GetPoolAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetLpMintAccount sets the "lp_mint" account.
// LP token mint of the pool
func (inst *AddBalanceLiquidity) SetLpMintAccount(lpMint ag_solanago.PublicKey) *AddBalanceLiquidity {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(lpMint).WRITE()
	return inst
}

// GetLpMintAccount gets the "lp_mint" account.
// LP token mint of the pool
func (inst *AddBalanceLiquidity) GetLpMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetUserPoolLpAccount sets the "user_pool_lp" account.
// user pool lp token account. lp will be burned from this account upon success liquidity removal.
func (inst *AddBalanceLiquidity) SetUserPoolLpAccount(userPoolLp ag_solanago.PublicKey) *AddBalanceLiquidity {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(userPoolLp).WRITE()
	return inst
}

// GetUserPoolLpAccount gets the "user_pool_lp" account.
// user pool lp token account. lp will be burned from this account upon success liquidity removal.
func (inst *AddBalanceLiquidity) GetUserPoolLpAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetAVaultLpAccount sets the "a_vault_lp" account.
// LP token account of vault A. Used to receive/burn the vault LP upon deposit/withdraw from the vault.
func (inst *AddBalanceLiquidity) SetAVaultLpAccount(aVaultLp ag_solanago.PublicKey) *AddBalanceLiquidity {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(aVaultLp).WRITE()
	return inst
}

// GetAVaultLpAccount gets the "a_vault_lp" account.
// LP token account of vault A. Used to receive/burn the vault LP upon deposit/withdraw from the vault.
func (inst *AddBalanceLiquidity) GetAVaultLpAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetBVaultLpAccount sets the "b_vault_lp" account.
// LP token account of vault B. Used to receive/burn the vault LP upon deposit/withdraw from the vault.
func (inst *AddBalanceLiquidity) SetBVaultLpAccount(bVaultLp ag_solanago.PublicKey) *AddBalanceLiquidity {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(bVaultLp).WRITE()
	return inst
}

// GetBVaultLpAccount gets the "b_vault_lp" account.
// LP token account of vault B. Used to receive/burn the vault LP upon deposit/withdraw from the vault.
func (inst *AddBalanceLiquidity) GetBVaultLpAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetAVaultAccount sets the "a_vault" account.
// Vault account for token a. token a of the pool will be deposit / withdraw from this vault account.
func (inst *AddBalanceLiquidity) SetAVaultAccount(aVault ag_solanago.PublicKey) *AddBalanceLiquidity {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(aVault).WRITE()
	return inst
}

// GetAVaultAccount gets the "a_vault" account.
// Vault account for token a. token a of the pool will be deposit / withdraw from this vault account.
func (inst *AddBalanceLiquidity) GetAVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetBVaultAccount sets the "b_vault" account.
// Vault account for token b. token b of the pool will be deposit / withdraw from this vault account.
func (inst *AddBalanceLiquidity) SetBVaultAccount(bVault ag_solanago.PublicKey) *AddBalanceLiquidity {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(bVault).WRITE()
	return inst
}

// GetBVaultAccount gets the "b_vault" account.
// Vault account for token b. token b of the pool will be deposit / withdraw from this vault account.
func (inst *AddBalanceLiquidity) GetBVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetAVaultLpMintAccount sets the "a_vault_lp_mint" account.
// LP token mint of vault a
func (inst *AddBalanceLiquidity) SetAVaultLpMintAccount(aVaultLpMint ag_solanago.PublicKey) *AddBalanceLiquidity {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(aVaultLpMint).WRITE()
	return inst
}

// GetAVaultLpMintAccount gets the "a_vault_lp_mint" account.
// LP token mint of vault a
func (inst *AddBalanceLiquidity) GetAVaultLpMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetBVaultLpMintAccount sets the "b_vault_lp_mint" account.
// LP token mint of vault b
func (inst *AddBalanceLiquidity) SetBVaultLpMintAccount(bVaultLpMint ag_solanago.PublicKey) *AddBalanceLiquidity {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(bVaultLpMint).WRITE()
	return inst
}

// GetBVaultLpMintAccount gets the "b_vault_lp_mint" account.
// LP token mint of vault b
func (inst *AddBalanceLiquidity) GetBVaultLpMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetATokenVaultAccount sets the "a_token_vault" account.
// Token vault account of vault A
func (inst *AddBalanceLiquidity) SetATokenVaultAccount(aTokenVault ag_solanago.PublicKey) *AddBalanceLiquidity {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(aTokenVault).WRITE()
	return inst
}

// GetATokenVaultAccount gets the "a_token_vault" account.
// Token vault account of vault A
func (inst *AddBalanceLiquidity) GetATokenVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetBTokenVaultAccount sets the "b_token_vault" account.
// Token vault account of vault B
func (inst *AddBalanceLiquidity) SetBTokenVaultAccount(bTokenVault ag_solanago.PublicKey) *AddBalanceLiquidity {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(bTokenVault).WRITE()
	return inst
}

// GetBTokenVaultAccount gets the "b_token_vault" account.
// Token vault account of vault B
func (inst *AddBalanceLiquidity) GetBTokenVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetUserATokenAccount sets the "user_a_token" account.
// User token A account. Token will be transfer from this account if it is add liquidity operation. Else, token will be transfer into this account.
func (inst *AddBalanceLiquidity) SetUserATokenAccount(userAToken ag_solanago.PublicKey) *AddBalanceLiquidity {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(userAToken).WRITE()
	return inst
}

// GetUserATokenAccount gets the "user_a_token" account.
// User token A account. Token will be transfer from this account if it is add liquidity operation. Else, token will be transfer into this account.
func (inst *AddBalanceLiquidity) GetUserATokenAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetUserBTokenAccount sets the "user_b_token" account.
// User token B account. Token will be transfer from this account if it is add liquidity operation. Else, token will be transfer into this account.
func (inst *AddBalanceLiquidity) SetUserBTokenAccount(userBToken ag_solanago.PublicKey) *AddBalanceLiquidity {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(userBToken).WRITE()
	return inst
}

// GetUserBTokenAccount gets the "user_b_token" account.
// User token B account. Token will be transfer from this account if it is add liquidity operation. Else, token will be transfer into this account.
func (inst *AddBalanceLiquidity) GetUserBTokenAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

// SetUserAccount sets the "user" account.
// User account. Must be owner of user_a_token, and user_b_token.
func (inst *AddBalanceLiquidity) SetUserAccount(user ag_solanago.PublicKey) *AddBalanceLiquidity {
	inst.AccountMetaSlice[13] = ag_solanago.Meta(user).SIGNER()
	return inst
}

// GetUserAccount gets the "user" account.
// User account. Must be owner of user_a_token, and user_b_token.
func (inst *AddBalanceLiquidity) GetUserAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(13)
}

// SetVaultProgramAccount sets the "vault_program" account.
// Vault program. the pool will deposit/withdraw liquidity from the vault.
func (inst *AddBalanceLiquidity) SetVaultProgramAccount(vaultProgram ag_solanago.PublicKey) *AddBalanceLiquidity {
	inst.AccountMetaSlice[14] = ag_solanago.Meta(vaultProgram)
	return inst
}

// GetVaultProgramAccount gets the "vault_program" account.
// Vault program. the pool will deposit/withdraw liquidity from the vault.
func (inst *AddBalanceLiquidity) GetVaultProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(14)
}

// SetTokenProgramAccount sets the "token_program" account.
// Token program.
func (inst *AddBalanceLiquidity) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *AddBalanceLiquidity {
	inst.AccountMetaSlice[15] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "token_program" account.
// Token program.
func (inst *AddBalanceLiquidity) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(15)
}

func (inst *AddBalanceLiquidity) SetAccounts(accounts []*ag_solanago.AccountMeta) error {
	inst.AccountMetaSlice = accounts
	return inst.Validate()
}

func (inst *AddBalanceLiquidity) SetRemainingAccounts(metas []*ag_solanago.AccountMeta) *AddBalanceLiquidity {
	inst.AccountMetaSlice = append(inst.AccountMetaSlice[0:16], metas...)
	return inst
}

func (inst *AddBalanceLiquidity) GetRemainingAccounts() []*ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[16:]
}

func (inst AddBalanceLiquidity) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_AddBalanceLiquidity,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst AddBalanceLiquidity) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *AddBalanceLiquidity) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.PoolTokenAmount == nil {
			return errors.New("poolTokenAmount parameter is not set")
		}
		if inst.MaximumTokenAAmount == nil {
			return errors.New("maximumTokenAAmount parameter is not set")
		}
		if inst.MaximumTokenBAmount == nil {
			return errors.New("maximumTokenBAmount parameter is not set")
		}
	}

	if len(inst.AccountMetaSlice) < 16 {
		return errors.New("accounts slice has wrong length: expected 16 accounts")
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Pool is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.LpMint is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.UserPoolLp is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.AVaultLp is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.BVaultLp is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.AVault is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.BVault is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.AVaultLpMint is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.BVaultLpMint is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.ATokenVault is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.BTokenVault is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.UserAToken is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.UserBToken is not set")
		}
		if inst.AccountMetaSlice[13] == nil {
			return errors.New("accounts.User is not set")
		}
		if inst.AccountMetaSlice[14] == nil {
			return errors.New("accounts.VaultProgram is not set")
		}
		if inst.AccountMetaSlice[15] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
	}
	return nil
}

func (inst *AddBalanceLiquidity) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("AddBalanceLiquidity")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=3]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("       PoolTokenAmount", *inst.PoolTokenAmount))
						paramsBranch.Child(ag_format.Param("   MaximumTokenAAmount", *inst.MaximumTokenAAmount))
						paramsBranch.Child(ag_format.Param("   MaximumTokenBAmount", *inst.MaximumTokenBAmount))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=16]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("           pool", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("        lp_mint", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("   user_pool_lp", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("     a_vault_lp", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("     b_vault_lp", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("        a_vault", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("        b_vault", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("a_vault_lp_mint", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("b_vault_lp_mint", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("  a_token_vault", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("  b_token_vault", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("   user_a_token", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("   user_b_token", inst.AccountMetaSlice.Get(12)))
						accountsBranch.Child(ag_format.Meta("           user", inst.AccountMetaSlice.Get(13)))
						accountsBranch.Child(ag_format.Meta("  vault_program", inst.AccountMetaSlice.Get(14)))
						accountsBranch.Child(ag_format.Meta("  token_program", inst.AccountMetaSlice.Get(15)))
					})
				})
		})
}

func (obj AddBalanceLiquidity) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `PoolTokenAmount` param:
	err = encoder.Encode(obj.PoolTokenAmount)
	if err != nil {
		return err
	}
	// Serialize `MaximumTokenAAmount` param:
	err = encoder.Encode(obj.MaximumTokenAAmount)
	if err != nil {
		return err
	}
	// Serialize `MaximumTokenBAmount` param:
	err = encoder.Encode(obj.MaximumTokenBAmount)
	if err != nil {
		return err
	}
	return nil
}
func (obj *AddBalanceLiquidity) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `PoolTokenAmount`:
	err = decoder.Decode(&obj.PoolTokenAmount)
	if err != nil {
		return err
	}
	// Deserialize `MaximumTokenAAmount`:
	err = decoder.Decode(&obj.MaximumTokenAAmount)
	if err != nil {
		return err
	}
	// Deserialize `MaximumTokenBAmount`:
	err = decoder.Decode(&obj.MaximumTokenBAmount)
	if err != nil {
		return err
	}
	return nil
}

// NewAddBalanceLiquidityInstruction declares a new AddBalanceLiquidity instruction with the provided parameters and accounts.
func NewAddBalanceLiquidityInstruction(
	// Parameters:
	pool_token_amount uint64,
	maximum_token_a_amount uint64,
	maximum_token_b_amount uint64,
	// Accounts:
	pool ag_solanago.PublicKey,
	lpMint ag_solanago.PublicKey,
	userPoolLp ag_solanago.PublicKey,
	aVaultLp ag_solanago.PublicKey,
	bVaultLp ag_solanago.PublicKey,
	aVault ag_solanago.PublicKey,
	bVault ag_solanago.PublicKey,
	aVaultLpMint ag_solanago.PublicKey,
	bVaultLpMint ag_solanago.PublicKey,
	aTokenVault ag_solanago.PublicKey,
	bTokenVault ag_solanago.PublicKey,
	userAToken ag_solanago.PublicKey,
	userBToken ag_solanago.PublicKey,
	user ag_solanago.PublicKey,
	vaultProgram ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey) *AddBalanceLiquidity {
	return NewAddBalanceLiquidityInstructionBuilder().
		SetPoolTokenAmount(pool_token_amount).
		SetMaximumTokenAAmount(maximum_token_a_amount).
		SetMaximumTokenBAmount(maximum_token_b_amount).
		SetPoolAccount(pool).
		SetLpMintAccount(lpMint).
		SetUserPoolLpAccount(userPoolLp).
		SetAVaultLpAccount(aVaultLp).
		SetBVaultLpAccount(bVaultLp).
		SetAVaultAccount(aVault).
		SetBVaultAccount(bVault).
		SetAVaultLpMintAccount(aVaultLpMint).
		SetBVaultLpMintAccount(bVaultLpMint).
		SetATokenVaultAccount(aTokenVault).
		SetBTokenVaultAccount(bTokenVault).
		SetUserATokenAccount(userAToken).
		SetUserBTokenAccount(userBToken).
		SetUserAccount(user).
		SetVaultProgramAccount(vaultProgram).
		SetTokenProgramAccount(tokenProgram)
}

// NewSimpleAddBalanceLiquidityInstruction declares a new AddBalanceLiquidity instruction with the provided parameters and accounts.
func NewSimpleAddBalanceLiquidityInstruction(
	// Parameters:
	pool_token_amount uint64,
	maximum_token_a_amount uint64,
	maximum_token_b_amount uint64,
	// Accounts:
	pool ag_solanago.PublicKey,
	lpMint ag_solanago.PublicKey,
	userPoolLp ag_solanago.PublicKey,
	aVaultLp ag_solanago.PublicKey,
	bVaultLp ag_solanago.PublicKey,
	aVault ag_solanago.PublicKey,
	bVault ag_solanago.PublicKey,
	aVaultLpMint ag_solanago.PublicKey,
	bVaultLpMint ag_solanago.PublicKey,
	aTokenVault ag_solanago.PublicKey,
	bTokenVault ag_solanago.PublicKey,
	userAToken ag_solanago.PublicKey,
	userBToken ag_solanago.PublicKey,
	user ag_solanago.PublicKey,
	vaultProgram ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey) *AddBalanceLiquidity {
	return NewAddBalanceLiquidityInstructionBuilder().
		SetPoolTokenAmount(pool_token_amount).
		SetMaximumTokenAAmount(maximum_token_a_amount).
		SetMaximumTokenBAmount(maximum_token_b_amount).
		SetPoolAccount(pool).
		SetLpMintAccount(lpMint).
		SetUserPoolLpAccount(userPoolLp).
		SetAVaultLpAccount(aVaultLp).
		SetBVaultLpAccount(bVaultLp).
		SetAVaultAccount(aVault).
		SetBVaultAccount(bVault).
		SetAVaultLpMintAccount(aVaultLpMint).
		SetBVaultLpMintAccount(bVaultLpMint).
		SetATokenVaultAccount(aTokenVault).
		SetBTokenVaultAccount(bTokenVault).
		SetUserATokenAccount(userAToken).
		SetUserBTokenAccount(userBToken).
		SetUserAccount(user).
		SetVaultProgramAccount(vaultProgram).
		SetTokenProgramAccount(tokenProgram)
}
