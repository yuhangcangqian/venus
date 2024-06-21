// Code generated by github.com/filecoin-project/venus/venus-devtool/state-type-gen. DO NOT EDIT.
package types

import (
	"github.com/filecoin-project/venus/venus-shared/actors/types"
)

const (
	EthAddressLength = types.EthAddressLength
	EthBloomSize     = types.EthBloomSize
	EthHashLength    = types.EthHashLength
)

var (
	EmptyEthBloom     = types.EmptyEthBloom
	EmptyEthHash      = types.EmptyEthHash
	EmptyEthInt       = types.EmptyEthInt
	EmptyEthNonce     = types.EmptyEthNonce
	EmptyRootHash     = types.EmptyRootHash
	EmptyUncleHash    = types.EmptyUncleHash
	ErrInvalidAddress = types.ErrInvalidAddress
	EthBigIntZero     = types.EthBigIntZero
	FullEthBloom      = types.FullEthBloom
)

type (
	EthAddress                     = types.EthAddress
	EthAddressList                 = types.EthAddressList
	EthBigInt                      = types.EthBigInt
	EthBlock                       = types.EthBlock
	EthBlockNumberOrHash           = types.EthBlockNumberOrHash
	EthBytes                       = types.EthBytes
	EthCall                        = types.EthCall
	EthCallTraceAction             = types.EthCallTraceAction
	EthCallTraceResult             = types.EthCallTraceResult
	EthCreateTraceAction           = types.EthCreateTraceAction
	EthCreateTraceResult           = types.EthCreateTraceResult
	EthEstimateGasParams           = types.EthEstimateGasParams
	EthFeeHistory                  = types.EthFeeHistory
	EthFeeHistoryParams            = types.EthFeeHistoryParams
	EthFilterID                    = types.EthFilterID
	EthFilterResult                = types.EthFilterResult
	EthFilterSpec                  = types.EthFilterSpec
	EthHash                        = types.EthHash
	EthHashList                    = types.EthHashList
	EthLog                         = types.EthLog
	EthNonce                       = types.EthNonce
	EthSubscribeParams             = types.EthSubscribeParams
	EthSubscriptionID              = types.EthSubscriptionID
	EthSubscriptionParams          = types.EthSubscriptionParams
	EthSubscriptionResponse        = types.EthSubscriptionResponse
	EthSyncingResult               = types.EthSyncingResult
	EthTopicSpec                   = types.EthTopicSpec
	EthTrace                       = types.EthTrace
	EthTraceBlock                  = types.EthTraceBlock
	EthTraceReplayBlockTransaction = types.EthTraceReplayBlockTransaction
	EthTraceTransaction            = types.EthTraceTransaction
	EthTxReceipt                   = types.EthTxReceipt
	EthUint64                      = types.EthUint64
)

var (
	CastEthAddress                        = types.CastEthAddress
	DecodeHexString                       = types.DecodeHexString
	DecodeHexStringTrimSpace              = types.DecodeHexStringTrimSpace
	EthAddressFromActorID                 = types.EthAddressFromActorID
	EthAddressFromFilecoinAddress         = types.EthAddressFromFilecoinAddress
	EthAddressFromPubKey                  = types.EthAddressFromPubKey
	EthBloomSet                           = types.EthBloomSet
	EthHashFromCid                        = types.EthHashFromCid
	EthHashFromTxBytes                    = types.EthHashFromTxBytes
	EthUint64FromBytes                    = types.EthUint64FromBytes
	EthUint64FromHex                      = types.EthUint64FromHex
	GetContractEthAddressFromCode         = types.GetContractEthAddressFromCode
	IsEthAddress                          = types.IsEthAddress
	NewEthBlock                           = types.NewEthBlock
	NewEthBlockNumberOrHashFromHexString  = types.NewEthBlockNumberOrHashFromHexString
	NewEthBlockNumberOrHashFromNumber     = types.NewEthBlockNumberOrHashFromNumber
	NewEthBlockNumberOrHashFromPredefined = types.NewEthBlockNumberOrHashFromPredefined
	ParseEthAddress                       = types.ParseEthAddress
	ParseEthHash                          = types.ParseEthHash
	SetEip155ChainID                      = types.SetEip155ChainID
	TryEthAddressFromFilecoinAddress      = types.TryEthAddressFromFilecoinAddress
)
