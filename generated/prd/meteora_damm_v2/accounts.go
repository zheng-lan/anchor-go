// Code generated by https://github.com/zheng-lan/anchor-go. DO NOT EDIT.

package meteora_damm_v2

import (
	"fmt"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
)

type ClaimFeeOperatorAccount struct {
	// operator
	Operator ag_solanago.PublicKey

	// Reserve
	_padding [128]uint8
}

var ClaimFeeOperatorAccountDiscriminator = [8]byte{166, 48, 134, 86, 34, 200, 188, 150}

func (obj ClaimFeeOperatorAccount) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(ClaimFeeOperatorAccountDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Operator` param:
	err = encoder.Encode(obj.Operator)
	if err != nil {
		return err
	}
	// Serialize `_padding` param:
	err = encoder.Encode(obj._padding)
	if err != nil {
		return err
	}
	return nil
}

func (obj *ClaimFeeOperatorAccount) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(ClaimFeeOperatorAccountDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[166 48 134 86 34 200 188 150]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Operator`:
	err = decoder.Decode(&obj.Operator)
	if err != nil {
		return err
	}
	// Deserialize `_padding`:
	err = decoder.Decode(&obj._padding)
	if err != nil {
		return err
	}
	return nil
}

type ConfigAccount struct {
	// Vault config key
	VaultConfigKey ag_solanago.PublicKey

	// Only pool_creator_authority can use the current config to initialize new pool. When it's Pubkey::default, it's a public config.
	PoolCreatorAuthority ag_solanago.PublicKey

	// Pool fee
	PoolFees PoolFeesConfig

	// Activation type
	ActivationType uint8

	// Collect fee mode
	CollectFeeMode uint8

	// Config type mode, 0 for static, 1 for dynamic
	ConfigType uint8

	// padding 0
	_padding_0 [5]uint8

	// config index
	Index uint64

	// sqrt min price
	SqrtMinPrice ag_binary.Uint128

	// sqrt max price
	SqrtMaxPrice ag_binary.Uint128

	// Fee curve point
	// Padding for further use
	_padding_1 [10]uint64
}

var ConfigAccountDiscriminator = [8]byte{155, 12, 170, 224, 30, 250, 204, 130}

func (obj ConfigAccount) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(ConfigAccountDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `VaultConfigKey` param:
	err = encoder.Encode(obj.VaultConfigKey)
	if err != nil {
		return err
	}
	// Serialize `PoolCreatorAuthority` param:
	err = encoder.Encode(obj.PoolCreatorAuthority)
	if err != nil {
		return err
	}
	// Serialize `PoolFees` param:
	err = encoder.Encode(obj.PoolFees)
	if err != nil {
		return err
	}
	// Serialize `ActivationType` param:
	err = encoder.Encode(obj.ActivationType)
	if err != nil {
		return err
	}
	// Serialize `CollectFeeMode` param:
	err = encoder.Encode(obj.CollectFeeMode)
	if err != nil {
		return err
	}
	// Serialize `ConfigType` param:
	err = encoder.Encode(obj.ConfigType)
	if err != nil {
		return err
	}
	// Serialize `_padding_0` param:
	err = encoder.Encode(obj._padding_0)
	if err != nil {
		return err
	}
	// Serialize `Index` param:
	err = encoder.Encode(obj.Index)
	if err != nil {
		return err
	}
	// Serialize `SqrtMinPrice` param:
	err = encoder.Encode(obj.SqrtMinPrice)
	if err != nil {
		return err
	}
	// Serialize `SqrtMaxPrice` param:
	err = encoder.Encode(obj.SqrtMaxPrice)
	if err != nil {
		return err
	}
	// Serialize `_padding_1` param:
	err = encoder.Encode(obj._padding_1)
	if err != nil {
		return err
	}
	return nil
}

func (obj *ConfigAccount) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(ConfigAccountDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[155 12 170 224 30 250 204 130]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `VaultConfigKey`:
	err = decoder.Decode(&obj.VaultConfigKey)
	if err != nil {
		return err
	}
	// Deserialize `PoolCreatorAuthority`:
	err = decoder.Decode(&obj.PoolCreatorAuthority)
	if err != nil {
		return err
	}
	// Deserialize `PoolFees`:
	err = decoder.Decode(&obj.PoolFees)
	if err != nil {
		return err
	}
	// Deserialize `ActivationType`:
	err = decoder.Decode(&obj.ActivationType)
	if err != nil {
		return err
	}
	// Deserialize `CollectFeeMode`:
	err = decoder.Decode(&obj.CollectFeeMode)
	if err != nil {
		return err
	}
	// Deserialize `ConfigType`:
	err = decoder.Decode(&obj.ConfigType)
	if err != nil {
		return err
	}
	// Deserialize `_padding_0`:
	err = decoder.Decode(&obj._padding_0)
	if err != nil {
		return err
	}
	// Deserialize `Index`:
	err = decoder.Decode(&obj.Index)
	if err != nil {
		return err
	}
	// Deserialize `SqrtMinPrice`:
	err = decoder.Decode(&obj.SqrtMinPrice)
	if err != nil {
		return err
	}
	// Deserialize `SqrtMaxPrice`:
	err = decoder.Decode(&obj.SqrtMaxPrice)
	if err != nil {
		return err
	}
	// Deserialize `_padding_1`:
	err = decoder.Decode(&obj._padding_1)
	if err != nil {
		return err
	}
	return nil
}

type PoolAccount struct {
	// Pool fee
	PoolFees PoolFeesStruct

	// token a mint
	TokenAMint ag_solanago.PublicKey

	// token b mint
	TokenBMint ag_solanago.PublicKey

	// token a vault
	TokenAVault ag_solanago.PublicKey

	// token b vault
	TokenBVault ag_solanago.PublicKey

	// Whitelisted vault to be able to buy pool before activation_point
	WhitelistedVault ag_solanago.PublicKey

	// partner
	Partner ag_solanago.PublicKey

	// liquidity share
	Liquidity ag_binary.Uint128

	// padding, previous reserve amount, be careful to use that field
	_padding ag_binary.Uint128

	// protocol a fee
	ProtocolAFee uint64

	// protocol b fee
	ProtocolBFee uint64

	// partner a fee
	PartnerAFee uint64

	// partner b fee
	PartnerBFee uint64

	// min price
	SqrtMinPrice ag_binary.Uint128

	// max price
	SqrtMaxPrice ag_binary.Uint128

	// current price
	SqrtPrice ag_binary.Uint128

	// Activation point, can be slot or timestamp
	ActivationPoint uint64

	// Activation type, 0 means by slot, 1 means by timestamp
	ActivationType uint8

	// pool status, 0: enable, 1 disable
	PoolStatus uint8

	// token a flag
	TokenAFlag uint8

	// token b flag
	TokenBFlag uint8

	// 0 is collect fee in both token, 1 only collect fee in token a, 2 only collect fee in token b
	CollectFeeMode uint8

	// pool type
	PoolType uint8

	// padding
	_padding_0 [2]uint8

	// cumulative
	FeeAPerLiquidity [32]uint8

	// cumulative
	FeeBPerLiquidity       [32]uint8
	PermanentLockLiquidity ag_binary.Uint128

	// metrics
	Metrics PoolMetrics

	// Padding for further use
	_padding_1 [10]uint64

	// Farming reward information
	RewardInfos [2]RewardInfo
}

var PoolAccountDiscriminator = [8]byte{241, 154, 109, 4, 17, 177, 109, 188}

func (obj PoolAccount) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(PoolAccountDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `PoolFees` param:
	err = encoder.Encode(obj.PoolFees)
	if err != nil {
		return err
	}
	// Serialize `TokenAMint` param:
	err = encoder.Encode(obj.TokenAMint)
	if err != nil {
		return err
	}
	// Serialize `TokenBMint` param:
	err = encoder.Encode(obj.TokenBMint)
	if err != nil {
		return err
	}
	// Serialize `TokenAVault` param:
	err = encoder.Encode(obj.TokenAVault)
	if err != nil {
		return err
	}
	// Serialize `TokenBVault` param:
	err = encoder.Encode(obj.TokenBVault)
	if err != nil {
		return err
	}
	// Serialize `WhitelistedVault` param:
	err = encoder.Encode(obj.WhitelistedVault)
	if err != nil {
		return err
	}
	// Serialize `Partner` param:
	err = encoder.Encode(obj.Partner)
	if err != nil {
		return err
	}
	// Serialize `Liquidity` param:
	err = encoder.Encode(obj.Liquidity)
	if err != nil {
		return err
	}
	// Serialize `_padding` param:
	err = encoder.Encode(obj._padding)
	if err != nil {
		return err
	}
	// Serialize `ProtocolAFee` param:
	err = encoder.Encode(obj.ProtocolAFee)
	if err != nil {
		return err
	}
	// Serialize `ProtocolBFee` param:
	err = encoder.Encode(obj.ProtocolBFee)
	if err != nil {
		return err
	}
	// Serialize `PartnerAFee` param:
	err = encoder.Encode(obj.PartnerAFee)
	if err != nil {
		return err
	}
	// Serialize `PartnerBFee` param:
	err = encoder.Encode(obj.PartnerBFee)
	if err != nil {
		return err
	}
	// Serialize `SqrtMinPrice` param:
	err = encoder.Encode(obj.SqrtMinPrice)
	if err != nil {
		return err
	}
	// Serialize `SqrtMaxPrice` param:
	err = encoder.Encode(obj.SqrtMaxPrice)
	if err != nil {
		return err
	}
	// Serialize `SqrtPrice` param:
	err = encoder.Encode(obj.SqrtPrice)
	if err != nil {
		return err
	}
	// Serialize `ActivationPoint` param:
	err = encoder.Encode(obj.ActivationPoint)
	if err != nil {
		return err
	}
	// Serialize `ActivationType` param:
	err = encoder.Encode(obj.ActivationType)
	if err != nil {
		return err
	}
	// Serialize `PoolStatus` param:
	err = encoder.Encode(obj.PoolStatus)
	if err != nil {
		return err
	}
	// Serialize `TokenAFlag` param:
	err = encoder.Encode(obj.TokenAFlag)
	if err != nil {
		return err
	}
	// Serialize `TokenBFlag` param:
	err = encoder.Encode(obj.TokenBFlag)
	if err != nil {
		return err
	}
	// Serialize `CollectFeeMode` param:
	err = encoder.Encode(obj.CollectFeeMode)
	if err != nil {
		return err
	}
	// Serialize `PoolType` param:
	err = encoder.Encode(obj.PoolType)
	if err != nil {
		return err
	}
	// Serialize `_padding_0` param:
	err = encoder.Encode(obj._padding_0)
	if err != nil {
		return err
	}
	// Serialize `FeeAPerLiquidity` param:
	err = encoder.Encode(obj.FeeAPerLiquidity)
	if err != nil {
		return err
	}
	// Serialize `FeeBPerLiquidity` param:
	err = encoder.Encode(obj.FeeBPerLiquidity)
	if err != nil {
		return err
	}
	// Serialize `PermanentLockLiquidity` param:
	err = encoder.Encode(obj.PermanentLockLiquidity)
	if err != nil {
		return err
	}
	// Serialize `Metrics` param:
	err = encoder.Encode(obj.Metrics)
	if err != nil {
		return err
	}
	// Serialize `_padding_1` param:
	err = encoder.Encode(obj._padding_1)
	if err != nil {
		return err
	}
	// Serialize `RewardInfos` param:
	err = encoder.Encode(obj.RewardInfos)
	if err != nil {
		return err
	}
	return nil
}

func (obj *PoolAccount) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(PoolAccountDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[241 154 109 4 17 177 109 188]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `PoolFees`:
	err = decoder.Decode(&obj.PoolFees)
	if err != nil {
		return err
	}
	// Deserialize `TokenAMint`:
	err = decoder.Decode(&obj.TokenAMint)
	if err != nil {
		return err
	}
	// Deserialize `TokenBMint`:
	err = decoder.Decode(&obj.TokenBMint)
	if err != nil {
		return err
	}
	// Deserialize `TokenAVault`:
	err = decoder.Decode(&obj.TokenAVault)
	if err != nil {
		return err
	}
	// Deserialize `TokenBVault`:
	err = decoder.Decode(&obj.TokenBVault)
	if err != nil {
		return err
	}
	// Deserialize `WhitelistedVault`:
	err = decoder.Decode(&obj.WhitelistedVault)
	if err != nil {
		return err
	}
	// Deserialize `Partner`:
	err = decoder.Decode(&obj.Partner)
	if err != nil {
		return err
	}
	// Deserialize `Liquidity`:
	err = decoder.Decode(&obj.Liquidity)
	if err != nil {
		return err
	}
	// Deserialize `_padding`:
	err = decoder.Decode(&obj._padding)
	if err != nil {
		return err
	}
	// Deserialize `ProtocolAFee`:
	err = decoder.Decode(&obj.ProtocolAFee)
	if err != nil {
		return err
	}
	// Deserialize `ProtocolBFee`:
	err = decoder.Decode(&obj.ProtocolBFee)
	if err != nil {
		return err
	}
	// Deserialize `PartnerAFee`:
	err = decoder.Decode(&obj.PartnerAFee)
	if err != nil {
		return err
	}
	// Deserialize `PartnerBFee`:
	err = decoder.Decode(&obj.PartnerBFee)
	if err != nil {
		return err
	}
	// Deserialize `SqrtMinPrice`:
	err = decoder.Decode(&obj.SqrtMinPrice)
	if err != nil {
		return err
	}
	// Deserialize `SqrtMaxPrice`:
	err = decoder.Decode(&obj.SqrtMaxPrice)
	if err != nil {
		return err
	}
	// Deserialize `SqrtPrice`:
	err = decoder.Decode(&obj.SqrtPrice)
	if err != nil {
		return err
	}
	// Deserialize `ActivationPoint`:
	err = decoder.Decode(&obj.ActivationPoint)
	if err != nil {
		return err
	}
	// Deserialize `ActivationType`:
	err = decoder.Decode(&obj.ActivationType)
	if err != nil {
		return err
	}
	// Deserialize `PoolStatus`:
	err = decoder.Decode(&obj.PoolStatus)
	if err != nil {
		return err
	}
	// Deserialize `TokenAFlag`:
	err = decoder.Decode(&obj.TokenAFlag)
	if err != nil {
		return err
	}
	// Deserialize `TokenBFlag`:
	err = decoder.Decode(&obj.TokenBFlag)
	if err != nil {
		return err
	}
	// Deserialize `CollectFeeMode`:
	err = decoder.Decode(&obj.CollectFeeMode)
	if err != nil {
		return err
	}
	// Deserialize `PoolType`:
	err = decoder.Decode(&obj.PoolType)
	if err != nil {
		return err
	}
	// Deserialize `_padding_0`:
	err = decoder.Decode(&obj._padding_0)
	if err != nil {
		return err
	}
	// Deserialize `FeeAPerLiquidity`:
	err = decoder.Decode(&obj.FeeAPerLiquidity)
	if err != nil {
		return err
	}
	// Deserialize `FeeBPerLiquidity`:
	err = decoder.Decode(&obj.FeeBPerLiquidity)
	if err != nil {
		return err
	}
	// Deserialize `PermanentLockLiquidity`:
	err = decoder.Decode(&obj.PermanentLockLiquidity)
	if err != nil {
		return err
	}
	// Deserialize `Metrics`:
	err = decoder.Decode(&obj.Metrics)
	if err != nil {
		return err
	}
	// Deserialize `_padding_1`:
	err = decoder.Decode(&obj._padding_1)
	if err != nil {
		return err
	}
	// Deserialize `RewardInfos`:
	err = decoder.Decode(&obj.RewardInfos)
	if err != nil {
		return err
	}
	return nil
}

type PositionAccount struct {
	Pool ag_solanago.PublicKey

	// nft mint
	NftMint ag_solanago.PublicKey

	// fee a checkpoint
	FeeAPerTokenCheckpoint [32]uint8

	// fee b checkpoint
	FeeBPerTokenCheckpoint [32]uint8

	// fee a pending
	FeeAPending uint64

	// fee b pending
	FeeBPending uint64

	// unlock liquidity
	UnlockedLiquidity ag_binary.Uint128

	// vesting liquidity
	VestedLiquidity ag_binary.Uint128

	// permanent locked liquidity
	PermanentLockedLiquidity ag_binary.Uint128

	// metrics
	Metrics PositionMetrics

	// Farming reward information
	RewardInfos [2]UserRewardInfo

	// padding for future usage
	padding [6]ag_binary.Uint128
}

var PositionAccountDiscriminator = [8]byte{170, 188, 143, 228, 122, 64, 247, 208}

func (obj PositionAccount) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(PositionAccountDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Pool` param:
	err = encoder.Encode(obj.Pool)
	if err != nil {
		return err
	}
	// Serialize `NftMint` param:
	err = encoder.Encode(obj.NftMint)
	if err != nil {
		return err
	}
	// Serialize `FeeAPerTokenCheckpoint` param:
	err = encoder.Encode(obj.FeeAPerTokenCheckpoint)
	if err != nil {
		return err
	}
	// Serialize `FeeBPerTokenCheckpoint` param:
	err = encoder.Encode(obj.FeeBPerTokenCheckpoint)
	if err != nil {
		return err
	}
	// Serialize `FeeAPending` param:
	err = encoder.Encode(obj.FeeAPending)
	if err != nil {
		return err
	}
	// Serialize `FeeBPending` param:
	err = encoder.Encode(obj.FeeBPending)
	if err != nil {
		return err
	}
	// Serialize `UnlockedLiquidity` param:
	err = encoder.Encode(obj.UnlockedLiquidity)
	if err != nil {
		return err
	}
	// Serialize `VestedLiquidity` param:
	err = encoder.Encode(obj.VestedLiquidity)
	if err != nil {
		return err
	}
	// Serialize `PermanentLockedLiquidity` param:
	err = encoder.Encode(obj.PermanentLockedLiquidity)
	if err != nil {
		return err
	}
	// Serialize `Metrics` param:
	err = encoder.Encode(obj.Metrics)
	if err != nil {
		return err
	}
	// Serialize `RewardInfos` param:
	err = encoder.Encode(obj.RewardInfos)
	if err != nil {
		return err
	}
	// Serialize `padding` param:
	err = encoder.Encode(obj.padding)
	if err != nil {
		return err
	}
	return nil
}

func (obj *PositionAccount) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(PositionAccountDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[170 188 143 228 122 64 247 208]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Pool`:
	err = decoder.Decode(&obj.Pool)
	if err != nil {
		return err
	}
	// Deserialize `NftMint`:
	err = decoder.Decode(&obj.NftMint)
	if err != nil {
		return err
	}
	// Deserialize `FeeAPerTokenCheckpoint`:
	err = decoder.Decode(&obj.FeeAPerTokenCheckpoint)
	if err != nil {
		return err
	}
	// Deserialize `FeeBPerTokenCheckpoint`:
	err = decoder.Decode(&obj.FeeBPerTokenCheckpoint)
	if err != nil {
		return err
	}
	// Deserialize `FeeAPending`:
	err = decoder.Decode(&obj.FeeAPending)
	if err != nil {
		return err
	}
	// Deserialize `FeeBPending`:
	err = decoder.Decode(&obj.FeeBPending)
	if err != nil {
		return err
	}
	// Deserialize `UnlockedLiquidity`:
	err = decoder.Decode(&obj.UnlockedLiquidity)
	if err != nil {
		return err
	}
	// Deserialize `VestedLiquidity`:
	err = decoder.Decode(&obj.VestedLiquidity)
	if err != nil {
		return err
	}
	// Deserialize `PermanentLockedLiquidity`:
	err = decoder.Decode(&obj.PermanentLockedLiquidity)
	if err != nil {
		return err
	}
	// Deserialize `Metrics`:
	err = decoder.Decode(&obj.Metrics)
	if err != nil {
		return err
	}
	// Deserialize `RewardInfos`:
	err = decoder.Decode(&obj.RewardInfos)
	if err != nil {
		return err
	}
	// Deserialize `padding`:
	err = decoder.Decode(&obj.padding)
	if err != nil {
		return err
	}
	return nil
}

type TokenBadgeAccount struct {
	// token mint
	TokenMint ag_solanago.PublicKey

	// Reserve
	_padding [128]uint8
}

var TokenBadgeAccountDiscriminator = [8]byte{116, 219, 204, 229, 249, 116, 255, 150}

func (obj TokenBadgeAccount) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(TokenBadgeAccountDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `TokenMint` param:
	err = encoder.Encode(obj.TokenMint)
	if err != nil {
		return err
	}
	// Serialize `_padding` param:
	err = encoder.Encode(obj._padding)
	if err != nil {
		return err
	}
	return nil
}

func (obj *TokenBadgeAccount) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(TokenBadgeAccountDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[116 219 204 229 249 116 255 150]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `TokenMint`:
	err = decoder.Decode(&obj.TokenMint)
	if err != nil {
		return err
	}
	// Deserialize `_padding`:
	err = decoder.Decode(&obj._padding)
	if err != nil {
		return err
	}
	return nil
}

type VestingAccount struct {
	Position               ag_solanago.PublicKey
	CliffPoint             uint64
	PeriodFrequency        uint64
	CliffUnlockLiquidity   ag_binary.Uint128
	LiquidityPerPeriod     ag_binary.Uint128
	TotalReleasedLiquidity ag_binary.Uint128
	NumberOfPeriod         uint16
	padding                [14]uint8
	padding2               [4]ag_binary.Uint128
}

var VestingAccountDiscriminator = [8]byte{100, 149, 66, 138, 95, 200, 128, 241}

func (obj VestingAccount) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(VestingAccountDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Position` param:
	err = encoder.Encode(obj.Position)
	if err != nil {
		return err
	}
	// Serialize `CliffPoint` param:
	err = encoder.Encode(obj.CliffPoint)
	if err != nil {
		return err
	}
	// Serialize `PeriodFrequency` param:
	err = encoder.Encode(obj.PeriodFrequency)
	if err != nil {
		return err
	}
	// Serialize `CliffUnlockLiquidity` param:
	err = encoder.Encode(obj.CliffUnlockLiquidity)
	if err != nil {
		return err
	}
	// Serialize `LiquidityPerPeriod` param:
	err = encoder.Encode(obj.LiquidityPerPeriod)
	if err != nil {
		return err
	}
	// Serialize `TotalReleasedLiquidity` param:
	err = encoder.Encode(obj.TotalReleasedLiquidity)
	if err != nil {
		return err
	}
	// Serialize `NumberOfPeriod` param:
	err = encoder.Encode(obj.NumberOfPeriod)
	if err != nil {
		return err
	}
	// Serialize `padding` param:
	err = encoder.Encode(obj.padding)
	if err != nil {
		return err
	}
	// Serialize `padding2` param:
	err = encoder.Encode(obj.padding2)
	if err != nil {
		return err
	}
	return nil
}

func (obj *VestingAccount) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(VestingAccountDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[100 149 66 138 95 200 128 241]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Position`:
	err = decoder.Decode(&obj.Position)
	if err != nil {
		return err
	}
	// Deserialize `CliffPoint`:
	err = decoder.Decode(&obj.CliffPoint)
	if err != nil {
		return err
	}
	// Deserialize `PeriodFrequency`:
	err = decoder.Decode(&obj.PeriodFrequency)
	if err != nil {
		return err
	}
	// Deserialize `CliffUnlockLiquidity`:
	err = decoder.Decode(&obj.CliffUnlockLiquidity)
	if err != nil {
		return err
	}
	// Deserialize `LiquidityPerPeriod`:
	err = decoder.Decode(&obj.LiquidityPerPeriod)
	if err != nil {
		return err
	}
	// Deserialize `TotalReleasedLiquidity`:
	err = decoder.Decode(&obj.TotalReleasedLiquidity)
	if err != nil {
		return err
	}
	// Deserialize `NumberOfPeriod`:
	err = decoder.Decode(&obj.NumberOfPeriod)
	if err != nil {
		return err
	}
	// Deserialize `padding`:
	err = decoder.Decode(&obj.padding)
	if err != nil {
		return err
	}
	// Deserialize `padding2`:
	err = decoder.Decode(&obj.padding2)
	if err != nil {
		return err
	}
	return nil
}
