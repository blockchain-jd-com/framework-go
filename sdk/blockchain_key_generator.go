package sdk

import (
	"github.com/blockchain-jd-com/framework-go/crypto"
	"github.com/blockchain-jd-com/framework-go/crypto/framework"
	"github.com/blockchain-jd-com/framework-go/ledger_model"
)

/*
 * Author: imuge
 * Date: 2020/5/29 下午2:13
 */

type BlockchainKeyGenerator struct {
}

func NewBlockchainKeyGenerator() *BlockchainKeyGenerator {
	return &BlockchainKeyGenerator{}
}

func (b BlockchainKeyGenerator) Generate(algorithm framework.CryptoAlgorithm) ledger_model.BlockchainKeypair {
	signFunc := crypto.GetSignatureFunction(algorithm)
	cryptoKeyPair := signFunc.GenerateKeypair()
	return ledger_model.NewBlockchainKeypair(cryptoKeyPair.PubKey, cryptoKeyPair.PrivKey)
}

func (b BlockchainKeyGenerator) GenerateWithSeed(algorithm framework.CryptoAlgorithm, seed []byte) ledger_model.BlockchainKeypair {
	signFunc := crypto.GetSignatureFunction(algorithm)
	cryptoKeyPair := signFunc.GenerateKeypair()
	return ledger_model.NewBlockchainKeypair(cryptoKeyPair.PubKey, cryptoKeyPair.PrivKey)
}