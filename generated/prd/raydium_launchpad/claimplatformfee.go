// Code generated by https://github.com/zheng-lan/anchor-go. DO NOT EDIT.

package raydium_launchpad

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Claim platform fee
// # Arguments
//
// * `ctx` - The context of accounts
//
type ClaimPlatformFee struct {

	// [0] = [WRITE, SIGNER] platform_fee_wallet
	// ··········· Only the wallet stored in platform_config can collect platform fees
	//
	// [1] = [] authority
	// ··········· PDA that acts as the authority for pool vault and mint operations
	// ··········· Generated using AUTH_SEED
	//
	// [2] = [WRITE] pool_state
	// ··········· Account that stores the pool's state and parameters
	// ··········· PDA generated using POOL_SEED and both token mints
	//
	// [3] = [] platform_config
	// ··········· The platform config account
	//
	// [4] = [WRITE] quote_vault
	//
	// [5] = [WRITE] recipient_token_account
	// ··········· The address that receives the collected quote token fees
	//
	// [6] = [] quote_mint
	// ··········· The mint of quote token vault
	//
	// [7] = [] token_program
	// ··········· SPL program for input token transfers
	//
	// [8] = [] system_program
	// ··········· Required for account creation
	//
	// [9] = [] associated_token_program
	// ··········· Required for associated token program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewClaimPlatformFeeInstructionBuilder creates a new `ClaimPlatformFee` instruction builder.
func NewClaimPlatformFeeInstructionBuilder() *ClaimPlatformFee {
	nd := &ClaimPlatformFee{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 10),
	}
	nd.AccountMetaSlice[1] = ag_solanago.Meta(AuthorityPDA)
	nd.AccountMetaSlice[7] = ag_solanago.Meta(TokenProgram)
	nd.AccountMetaSlice[8] = ag_solanago.Meta(SystemProgram)
	nd.AccountMetaSlice[9] = ag_solanago.Meta(AssociatedTokenProgram)
	return nd
}

// SetPlatformFeeWalletAccount sets the "platform_fee_wallet" account.
// Only the wallet stored in platform_config can collect platform fees
func (inst *ClaimPlatformFee) SetPlatformFeeWalletAccount(platformFeeWallet ag_solanago.PublicKey) *ClaimPlatformFee {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(platformFeeWallet).WRITE().SIGNER()
	return inst
}

// GetPlatformFeeWalletAccount gets the "platform_fee_wallet" account.
// Only the wallet stored in platform_config can collect platform fees
func (inst *ClaimPlatformFee) GetPlatformFeeWalletAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetAuthorityAccount sets the "authority" account.
// PDA that acts as the authority for pool vault and mint operations
// Generated using AUTH_SEED
func (inst *ClaimPlatformFee) SetAuthorityAccount(authority ag_solanago.PublicKey) *ClaimPlatformFee {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(authority)
	return inst
}

// GetAuthorityAccount gets the "authority" account.
// PDA that acts as the authority for pool vault and mint operations
// Generated using AUTH_SEED
func (inst *ClaimPlatformFee) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetPoolStateAccount sets the "pool_state" account.
// Account that stores the pool's state and parameters
// PDA generated using POOL_SEED and both token mints
func (inst *ClaimPlatformFee) SetPoolStateAccount(poolState ag_solanago.PublicKey) *ClaimPlatformFee {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(poolState).WRITE()
	return inst
}

// GetPoolStateAccount gets the "pool_state" account.
// Account that stores the pool's state and parameters
// PDA generated using POOL_SEED and both token mints
func (inst *ClaimPlatformFee) GetPoolStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetPlatformConfigAccount sets the "platform_config" account.
// The platform config account
func (inst *ClaimPlatformFee) SetPlatformConfigAccount(platformConfig ag_solanago.PublicKey) *ClaimPlatformFee {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(platformConfig)
	return inst
}

// GetPlatformConfigAccount gets the "platform_config" account.
// The platform config account
func (inst *ClaimPlatformFee) GetPlatformConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetQuoteVaultAccount sets the "quote_vault" account.
func (inst *ClaimPlatformFee) SetQuoteVaultAccount(quoteVault ag_solanago.PublicKey) *ClaimPlatformFee {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(quoteVault).WRITE()
	return inst
}

// GetQuoteVaultAccount gets the "quote_vault" account.
func (inst *ClaimPlatformFee) GetQuoteVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetRecipientTokenAccountAccount sets the "recipient_token_account" account.
// The address that receives the collected quote token fees
func (inst *ClaimPlatformFee) SetRecipientTokenAccountAccount(recipientTokenAccount ag_solanago.PublicKey) *ClaimPlatformFee {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(recipientTokenAccount).WRITE()
	return inst
}

// GetRecipientTokenAccountAccount gets the "recipient_token_account" account.
// The address that receives the collected quote token fees
func (inst *ClaimPlatformFee) GetRecipientTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetQuoteMintAccount sets the "quote_mint" account.
// The mint of quote token vault
func (inst *ClaimPlatformFee) SetQuoteMintAccount(quoteMint ag_solanago.PublicKey) *ClaimPlatformFee {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(quoteMint)
	return inst
}

// GetQuoteMintAccount gets the "quote_mint" account.
// The mint of quote token vault
func (inst *ClaimPlatformFee) GetQuoteMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetTokenProgramAccount sets the "token_program" account.
// SPL program for input token transfers
func (inst *ClaimPlatformFee) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *ClaimPlatformFee {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "token_program" account.
// SPL program for input token transfers
func (inst *ClaimPlatformFee) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetSystemProgramAccount sets the "system_program" account.
// Required for account creation
func (inst *ClaimPlatformFee) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *ClaimPlatformFee {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "system_program" account.
// Required for account creation
func (inst *ClaimPlatformFee) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetAssociatedTokenProgramAccount sets the "associated_token_program" account.
// Required for associated token program
func (inst *ClaimPlatformFee) SetAssociatedTokenProgramAccount(associatedTokenProgram ag_solanago.PublicKey) *ClaimPlatformFee {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(associatedTokenProgram)
	return inst
}

// GetAssociatedTokenProgramAccount gets the "associated_token_program" account.
// Required for associated token program
func (inst *ClaimPlatformFee) GetAssociatedTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

func (inst *ClaimPlatformFee) SetAccounts(accounts []*ag_solanago.AccountMeta) error {
	inst.AccountMetaSlice = accounts
	return inst.Validate()
}

func (inst *ClaimPlatformFee) SetRemainingAccounts(metas []*ag_solanago.AccountMeta) *ClaimPlatformFee {
	inst.AccountMetaSlice = append(inst.AccountMetaSlice[0:10], metas...)
	return inst
}

func (inst *ClaimPlatformFee) GetRemainingAccounts() []*ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[10:]
}

func (inst ClaimPlatformFee) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_ClaimPlatformFee,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst ClaimPlatformFee) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *ClaimPlatformFee) Validate() error {
	if len(inst.AccountMetaSlice) < 10 {
		return errors.New("accounts slice has wrong length: expected 10 accounts")
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.PlatformFeeWallet is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.PoolState is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.PlatformConfig is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.QuoteVault is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.RecipientTokenAccount is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.QuoteMint is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.AssociatedTokenProgram is not set")
		}
	}
	return nil
}

func (inst *ClaimPlatformFee) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("ClaimPlatformFee")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=10]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("     platform_fee_wallet", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("               authority", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("              pool_state", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("         platform_config", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("             quote_vault", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("        recipient_token_", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("              quote_mint", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("           token_program", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("          system_program", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("associated_token_program", inst.AccountMetaSlice.Get(9)))
					})
				})
		})
}

func (obj ClaimPlatformFee) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *ClaimPlatformFee) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewClaimPlatformFeeInstruction declares a new ClaimPlatformFee instruction with the provided parameters and accounts.
func NewClaimPlatformFeeInstruction(
	// Accounts:
	platformFeeWallet ag_solanago.PublicKey,
	poolState ag_solanago.PublicKey,
	platformConfig ag_solanago.PublicKey,
	quoteVault ag_solanago.PublicKey,
	recipientTokenAccount ag_solanago.PublicKey,
	quoteMint ag_solanago.PublicKey) *ClaimPlatformFee {
	return NewClaimPlatformFeeInstructionBuilder().
		SetPlatformFeeWalletAccount(platformFeeWallet).
		SetPoolStateAccount(poolState).
		SetPlatformConfigAccount(platformConfig).
		SetQuoteVaultAccount(quoteVault).
		SetRecipientTokenAccountAccount(recipientTokenAccount).
		SetQuoteMintAccount(quoteMint)
}

// NewSimpleClaimPlatformFeeInstruction declares a new ClaimPlatformFee instruction with the provided parameters and accounts.
func NewSimpleClaimPlatformFeeInstruction(
	// Accounts:
	platformFeeWallet ag_solanago.PublicKey,
	poolState ag_solanago.PublicKey,
	platformConfig ag_solanago.PublicKey,
	quoteVault ag_solanago.PublicKey,
	recipientTokenAccount ag_solanago.PublicKey,
	quoteMint ag_solanago.PublicKey) *ClaimPlatformFee {
	return NewClaimPlatformFeeInstructionBuilder().
		SetPlatformFeeWalletAccount(platformFeeWallet).
		SetPoolStateAccount(poolState).
		SetPlatformConfigAccount(platformConfig).
		SetQuoteVaultAccount(quoteVault).
		SetRecipientTokenAccountAccount(recipientTokenAccount).
		SetQuoteMintAccount(quoteMint)
}
