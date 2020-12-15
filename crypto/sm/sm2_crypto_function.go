package sm

import (
	"github.com/blockchain-jd-com/framework-go/crypto/framework"
	"github.com/blockchain-jd-com/framework-go/utils/sm2"
)

/**
 * @Author: imuge
 * @Date: 2020/4/28 3:00 下午
 */

var (
	SM2_ECPOINT_SIZE           = 65
	SM2_PRIVKEY_SIZE           = 32
	SM2_SIGNATUREDIGEST_SIZE   = 64
	SM2_HASHDIGEST_SIZE        = 32
	SM2_PUBKEY_LENGTH          = framework.ALGORYTHM_CODE_SIZE + framework.KEY_TYPE_BYTES + SM2_ECPOINT_SIZE
	SM2_PRIVKEY_LENGTH         = framework.ALGORYTHM_CODE_SIZE + framework.KEY_TYPE_BYTES + SM2_PRIVKEY_SIZE
	SM2_SIGNATUREDIGEST_LENGTH = framework.ALGORYTHM_CODE_SIZE + SM2_SIGNATUREDIGEST_SIZE
)

var _ framework.AsymmetricEncryptionFunction = (*SM2CryptoFunction)(nil)
var _ framework.SignatureFunction = (*SM2CryptoFunction)(nil)

type SM2CryptoFunction struct {
}

func (S SM2CryptoFunction) Encrypt(pubKey framework.PubKey, data []byte) framework.AsymmetricCiphertext {
	rawPubKeyBytes := pubKey.GetRawKeyBytes()

	// 验证原始公钥长度为65字节
	if len(rawPubKeyBytes) != SM2_ECPOINT_SIZE {
		panic("This key has wrong format!")
	}

	// 验证密钥数据的算法标识对应SM2算法
	if pubKey.GetAlgorithm() != S.GetAlgorithm().Code {
		panic("The is not sm2 public key!")
	}

	// 调用SM2加密算法计算密文
	return framework.NewAsymmetricCiphertext(S.GetAlgorithm(), sm2.Encrypt(sm2.BytesToPubKey(rawPubKeyBytes), data))
}

func (S SM2CryptoFunction) Decrypt(privKey framework.PrivKey, ciphertext framework.AsymmetricCiphertext) []byte {
	rawPrivKeyBytes := privKey.GetRawKeyBytes()
	rawCiphertextBytes := ciphertext.GetRawCiphertext()

	// 验证原始私钥长度为32字节
	if len(rawPrivKeyBytes) != SM2_PRIVKEY_SIZE {
		panic("This key has wrong format!")
	}

	// 验证密钥数据的算法标识对应SM2算法
	if privKey.GetAlgorithm() != S.GetAlgorithm().Code {
		panic("This key is not SM2 private key!")
	}

	// 验证密文数据的算法标识对应SM2算法，并且密文符合长度要求
	if ciphertext.GetAlgorithm() != S.GetAlgorithm().Code || len(rawCiphertextBytes) < SM2_ECPOINT_SIZE+SM2_HASHDIGEST_SIZE {
		panic("This is not SM2 ciphertext!")
	}

	// 调用SM2解密算法得到明文结果
	return sm2.Decrypt(sm2.BytesToPrivKey(rawPrivKeyBytes), rawCiphertextBytes)
}

func (S SM2CryptoFunction) SupportCiphertext(ciphertextBytes []byte) bool {
	// 验证输入字节数组长度>=算法标识长度+椭圆曲线点长度+哈希长度，字节数组的算法标识对应SM2算法
	return len(ciphertextBytes) >= framework.ALGORYTHM_CODE_SIZE+SM2_ECPOINT_SIZE+SM2_HASHDIGEST_SIZE && S.GetAlgorithm().Match(ciphertextBytes, 0)
}

func (S SM2CryptoFunction) ParseCiphertext(ciphertextBytes []byte) framework.AsymmetricCiphertext {
	if S.SupportCiphertext(ciphertextBytes) {
		return framework.ParseAsymmetricCiphertext(ciphertextBytes)
	} else {
		panic("ciphertextBytes are invalid!")
	}
}

func (S SM2CryptoFunction) Sign(privKey framework.PrivKey, data []byte) framework.SignatureDigest {
	rawPrivKeyBytes := privKey.GetRawKeyBytes()

	// 验证原始私钥长度为256比特，即32字节
	if len(rawPrivKeyBytes) != SM2_PRIVKEY_SIZE {
		panic("This key has wrong format!")
	}

	// 验证密钥数据的算法标识对应SM2签名算法
	if privKey.GetAlgorithm() != S.GetAlgorithm().Code {
		panic("This key is not SM2 private key!")
	}

	// 调用SM2签名算法计算签名结果
	return framework.NewSignatureDigest(S.GetAlgorithm(), sm2.Sign(sm2.BytesToPrivKey(rawPrivKeyBytes), data))
}

func (S SM2CryptoFunction) Verify(pubKey framework.PubKey, data []byte, digest framework.SignatureDigest) bool {
	rawPubKeyBytes := pubKey.GetRawKeyBytes()
	rawDigestBytes := digest.GetRawDigest()

	// 验证原始公钥长度为520比特，即65字节
	if len(rawPubKeyBytes) != SM2_ECPOINT_SIZE {
		panic("This key has wrong format!")
	}

	// 验证密钥数据的算法标识对应SM2签名算法
	if pubKey.GetAlgorithm() != S.GetAlgorithm().Code {
		panic("This key is not SM2 public key!")
	}

	// 验证签名数据的算法标识对应SM2签名算法，并且原始签名长度为64字节
	if digest.GetAlgorithm() != S.GetAlgorithm().Code || len(rawDigestBytes) != SM2_SIGNATUREDIGEST_SIZE {
		panic("This is not SM2 signature digest!")
	}

	// 调用SM2验签算法验证签名结果
	return sm2.Verify(sm2.BytesToPubKey(rawPubKeyBytes), data, rawDigestBytes)
}

func (S SM2CryptoFunction) RetrievePubKey(privKey framework.PrivKey) framework.PubKey {
	return framework.NewPubKey(S.GetAlgorithm(), sm2.PubKeyToBytes(sm2.RetrievePubKey(sm2.BytesToPrivKey(privKey.GetRawKeyBytes()))))
}

func (S SM2CryptoFunction) SupportPrivKey(privKeyBytes []byte) bool {
	// 验证输入字节数组长度=算法标识长度+密钥类型长度+密钥长度，密钥数据的算法标识对应SM2算法，并且密钥类型是私钥
	return len(privKeyBytes) == SM2_PRIVKEY_LENGTH && S.GetAlgorithm().Match(privKeyBytes, 0) && privKeyBytes[framework.ALGORYTHM_CODE_SIZE] == framework.PRIVATE.Code
}

func (S SM2CryptoFunction) ParsePrivKey(privKeyBytes []byte) framework.PrivKey {
	if S.SupportPrivKey(privKeyBytes) {
		return framework.ParsePrivKey(privKeyBytes)
	} else {
		panic("privKeyBytes are invalid!")
	}
}

func (S SM2CryptoFunction) SupportPubKey(pubKeyBytes []byte) bool {
	// 验证输入字节数组长度=算法标识长度+密钥类型长度+椭圆曲线点长度，密钥数据的算法标识对应SM2算法，并且密钥类型是公钥
	return len(pubKeyBytes) == SM2_PUBKEY_LENGTH && S.GetAlgorithm().Match(pubKeyBytes, 0) && pubKeyBytes[framework.ALGORYTHM_CODE_SIZE] == framework.PUBLIC.Code
}

func (S SM2CryptoFunction) ParsePubKey(pubKeyBytes []byte) framework.PubKey {
	if S.SupportPubKey(pubKeyBytes) {
		return framework.ParsePubKey(pubKeyBytes)
	} else {
		panic("pubKeyBytes are invalid!")
	}
}

func (S SM2CryptoFunction) SupportDigest(digestBytes []byte) bool {
	// 验证输入字节数组长度=算法标识长度+签名长度，字节数组的算法标识对应SM2算法
	return len(digestBytes) == SM2_SIGNATUREDIGEST_LENGTH && S.GetAlgorithm().Match(digestBytes, 0)
}

func (S SM2CryptoFunction) ParseDigest(digestBytes []byte) framework.SignatureDigest {
	if S.SupportDigest(digestBytes) {
		return framework.ParseSignatureDigest(digestBytes)
	} else {
		panic("digestBytes are invalid!")
	}
}

func (S SM2CryptoFunction) GenerateKeypair() framework.AsymmetricKeypair {
	priv, pub := sm2.GenerateKeyPair()
	return framework.NewAsymmetricKeypair(framework.NewPubKey(S.GetAlgorithm(), sm2.PubKeyToBytes(pub)), framework.NewPrivKey(S.GetAlgorithm(), sm2.PrivKeyToBytes(priv)))
}

func (S SM2CryptoFunction) GenerateKeypairWithSeed(seed []byte) (keypair framework.AsymmetricKeypair, err error) {
	defer func() {
		r := recover()
		if r != nil {
			err = r.(error)
			return
		}
	}()
	priv, pub := sm2.GenerateKeyPairWithSeed(seed)
	keypair = framework.NewAsymmetricKeypair(framework.NewPubKey(S.GetAlgorithm(), sm2.PubKeyToBytes(pub)), framework.NewPrivKey(S.GetAlgorithm(), sm2.PrivKeyToBytes(priv)))

	return
}

func (S SM2CryptoFunction) GetAlgorithm() framework.CryptoAlgorithm {
	return SM2_ALGORITHM
}
