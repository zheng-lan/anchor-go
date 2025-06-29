// Code generated by https://github.com/zheng-lan/anchor-go. DO NOT EDIT.

package pump_curve

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Creates a new coin and bonding curve.
type Create struct {
	Name    *string
	Symbol  *string
	Uri     *string
	Creator *ag_solanago.PublicKey

	// [0] = [WRITE, SIGNER] mint
	//
	// [1] = [] mint_authority
	//
	// [2] = [WRITE] bonding_curve
	//
	// [3] = [WRITE] associated_bonding_curve
	//
	// [4] = [] global
	//
	// [5] = [] mpl_token_metadata
	//
	// [6] = [WRITE] metadata
	//
	// [7] = [WRITE, SIGNER] user
	//
	// [8] = [] system_program
	//
	// [9] = [] token_program
	//
	// [10] = [] associated_token_program
	//
	// [11] = [] rent
	//
	// [12] = [] event_authority
	//
	// [13] = [] program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewCreateInstructionBuilder creates a new `Create` instruction builder.
func NewCreateInstructionBuilder() *Create {
	nd := &Create{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 14),
	}
	nd.AccountMetaSlice[1] = ag_solanago.Meta(MintAuthorityPDA)
	nd.AccountMetaSlice[4] = ag_solanago.Meta(GlobalPDA)
	nd.AccountMetaSlice[5] = ag_solanago.Meta(TokenMetadataProgram)
	nd.AccountMetaSlice[8] = ag_solanago.Meta(SystemProgram)
	nd.AccountMetaSlice[9] = ag_solanago.Meta(TokenProgram)
	nd.AccountMetaSlice[10] = ag_solanago.Meta(AssociatedTokenProgram)
	nd.AccountMetaSlice[11] = ag_solanago.Meta(RentProgram)
	nd.AccountMetaSlice[12] = ag_solanago.Meta(EventAuthorityPDA)
	nd.AccountMetaSlice[13] = ag_solanago.Meta(ProgramID)
	return nd
}

// SetName sets the "name" parameter.
func (inst *Create) SetName(name string) *Create {
	inst.Name = &name
	return inst
}

// SetSymbol sets the "symbol" parameter.
func (inst *Create) SetSymbol(symbol string) *Create {
	inst.Symbol = &symbol
	return inst
}

// SetUri sets the "uri" parameter.
func (inst *Create) SetUri(uri string) *Create {
	inst.Uri = &uri
	return inst
}

// SetCreator sets the "creator" parameter.
func (inst *Create) SetCreator(creator ag_solanago.PublicKey) *Create {
	inst.Creator = &creator
	return inst
}

// SetMintAccount sets the "mint" account.
func (inst *Create) SetMintAccount(mint ag_solanago.PublicKey) *Create {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(mint).WRITE().SIGNER()
	return inst
}

// GetMintAccount gets the "mint" account.
func (inst *Create) GetMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetMintAuthorityAccount sets the "mint_authority" account.
func (inst *Create) SetMintAuthorityAccount(mintAuthority ag_solanago.PublicKey) *Create {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(mintAuthority)
	return inst
}

// GetMintAuthorityAccount gets the "mint_authority" account.
func (inst *Create) GetMintAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetBondingCurveAccount sets the "bonding_curve" account.
func (inst *Create) SetBondingCurveAccount(bondingCurve ag_solanago.PublicKey) *Create {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(bondingCurve).WRITE()
	return inst
}

// GetBondingCurveAccount gets the "bonding_curve" account.
func (inst *Create) GetBondingCurveAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetAssociatedBondingCurveAccount sets the "associated_bonding_curve" account.
func (inst *Create) SetAssociatedBondingCurveAccount(associatedBondingCurve ag_solanago.PublicKey) *Create {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(associatedBondingCurve).WRITE()
	return inst
}

// GetAssociatedBondingCurveAccount gets the "associated_bonding_curve" account.
func (inst *Create) GetAssociatedBondingCurveAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetGlobalAccount sets the "global" account.
func (inst *Create) SetGlobalAccount(global ag_solanago.PublicKey) *Create {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(global)
	return inst
}

// GetGlobalAccount gets the "global" account.
func (inst *Create) GetGlobalAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetMplTokenMetadataAccount sets the "mpl_token_metadata" account.
func (inst *Create) SetMplTokenMetadataAccount(mplTokenMetadata ag_solanago.PublicKey) *Create {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(mplTokenMetadata)
	return inst
}

// GetMplTokenMetadataAccount gets the "mpl_token_metadata" account.
func (inst *Create) GetMplTokenMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetMetadataAccount sets the "metadata" account.
func (inst *Create) SetMetadataAccount(metadata ag_solanago.PublicKey) *Create {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(metadata).WRITE()
	return inst
}

// GetMetadataAccount gets the "metadata" account.
func (inst *Create) GetMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetUserAccount sets the "user" account.
func (inst *Create) SetUserAccount(user ag_solanago.PublicKey) *Create {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(user).WRITE().SIGNER()
	return inst
}

// GetUserAccount gets the "user" account.
func (inst *Create) GetUserAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetSystemProgramAccount sets the "system_program" account.
func (inst *Create) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *Create {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "system_program" account.
func (inst *Create) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetTokenProgramAccount sets the "token_program" account.
func (inst *Create) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *Create {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "token_program" account.
func (inst *Create) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetAssociatedTokenProgramAccount sets the "associated_token_program" account.
func (inst *Create) SetAssociatedTokenProgramAccount(associatedTokenProgram ag_solanago.PublicKey) *Create {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(associatedTokenProgram)
	return inst
}

// GetAssociatedTokenProgramAccount gets the "associated_token_program" account.
func (inst *Create) GetAssociatedTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetRentAccount sets the "rent" account.
func (inst *Create) SetRentAccount(rent ag_solanago.PublicKey) *Create {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
func (inst *Create) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetEventAuthorityAccount sets the "event_authority" account.
func (inst *Create) SetEventAuthorityAccount(eventAuthority ag_solanago.PublicKey) *Create {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(eventAuthority)
	return inst
}

// GetEventAuthorityAccount gets the "event_authority" account.
func (inst *Create) GetEventAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

// SetProgramAccount sets the "program" account.
func (inst *Create) SetProgramAccount(program ag_solanago.PublicKey) *Create {
	inst.AccountMetaSlice[13] = ag_solanago.Meta(program)
	return inst
}

// GetProgramAccount gets the "program" account.
func (inst *Create) GetProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(13)
}

func (inst *Create) SetAccounts(accounts []*ag_solanago.AccountMeta) error {
	inst.AccountMetaSlice = accounts
	return inst.Validate()
}

func (inst *Create) SetRemainingAccounts(metas []*ag_solanago.AccountMeta) *Create {
	inst.AccountMetaSlice = append(inst.AccountMetaSlice[0:14], metas...)
	return inst
}

func (inst *Create) GetRemainingAccounts() []*ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[14:]
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
		if inst.Name == nil {
			return errors.New("name parameter is not set")
		}
		if inst.Symbol == nil {
			return errors.New("symbol parameter is not set")
		}
		if inst.Uri == nil {
			return errors.New("uri parameter is not set")
		}
		if inst.Creator == nil {
			return errors.New("creator parameter is not set")
		}
	}

	if len(inst.AccountMetaSlice) < 14 {
		return errors.New("accounts slice has wrong length: expected 14 accounts")
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Mint is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.MintAuthority is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.BondingCurve is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.AssociatedBondingCurve is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Global is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.MplTokenMetadata is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.Metadata is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.User is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.AssociatedTokenProgram is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.Rent is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.EventAuthority is not set")
		}
		if inst.AccountMetaSlice[13] == nil {
			return errors.New("accounts.Program is not set")
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
					instructionBranch.Child("Params[len=4]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("   Name", *inst.Name))
						paramsBranch.Child(ag_format.Param(" Symbol", *inst.Symbol))
						paramsBranch.Child(ag_format.Param("    Uri", *inst.Uri))
						paramsBranch.Child(ag_format.Param("Creator", *inst.Creator))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=14]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                    mint", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("          mint_authority", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("           bonding_curve", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("associated_bonding_curve", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("                  global", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("      mpl_token_metadata", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("                metadata", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("                    user", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("          system_program", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("           token_program", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("associated_token_program", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("                    rent", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("         event_authority", inst.AccountMetaSlice.Get(12)))
						accountsBranch.Child(ag_format.Meta("                 program", inst.AccountMetaSlice.Get(13)))
					})
				})
		})
}

func (obj Create) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Name` param:
	err = encoder.Encode(obj.Name)
	if err != nil {
		return err
	}
	// Serialize `Symbol` param:
	err = encoder.Encode(obj.Symbol)
	if err != nil {
		return err
	}
	// Serialize `Uri` param:
	err = encoder.Encode(obj.Uri)
	if err != nil {
		return err
	}
	// Serialize `Creator` param:
	err = encoder.Encode(obj.Creator)
	if err != nil {
		return err
	}
	return nil
}
func (obj *Create) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Name`:
	err = decoder.Decode(&obj.Name)
	if err != nil {
		return err
	}
	// Deserialize `Symbol`:
	err = decoder.Decode(&obj.Symbol)
	if err != nil {
		return err
	}
	// Deserialize `Uri`:
	err = decoder.Decode(&obj.Uri)
	if err != nil {
		return err
	}
	// Deserialize `Creator`:
	err = decoder.Decode(&obj.Creator)
	if err != nil {
		return err
	}
	return nil
}

// NewCreateInstruction declares a new Create instruction with the provided parameters and accounts.
func NewCreateInstruction(
	// Parameters:
	name string,
	symbol string,
	uri string,
	creator ag_solanago.PublicKey,
	// Accounts:
	mint ag_solanago.PublicKey,
	bondingCurve ag_solanago.PublicKey,
	associatedBondingCurve ag_solanago.PublicKey,
	metadata ag_solanago.PublicKey,
	user ag_solanago.PublicKey) *Create {
	return NewCreateInstructionBuilder().
		SetName(name).
		SetSymbol(symbol).
		SetUri(uri).
		SetCreator(creator).
		SetMintAccount(mint).
		SetBondingCurveAccount(bondingCurve).
		SetAssociatedBondingCurveAccount(associatedBondingCurve).
		SetMetadataAccount(metadata).
		SetUserAccount(user)
}

// NewSimpleCreateInstruction declares a new Create instruction with the provided parameters and accounts.
func NewSimpleCreateInstruction(
	// Parameters:
	name string,
	symbol string,
	uri string,
	creator ag_solanago.PublicKey,
	// Accounts:
	mint ag_solanago.PublicKey,
	associatedBondingCurve ag_solanago.PublicKey,
	metadata ag_solanago.PublicKey,
	user ag_solanago.PublicKey) *Create {
	bondingCurve := MustFindBondingCurveAddress(mint)
	return NewCreateInstructionBuilder().
		SetName(name).
		SetSymbol(symbol).
		SetUri(uri).
		SetCreator(creator).
		SetMintAccount(mint).
		SetBondingCurveAccount(bondingCurve).
		SetAssociatedBondingCurveAccount(associatedBondingCurve).
		SetMetadataAccount(metadata).
		SetUserAccount(user)
}
