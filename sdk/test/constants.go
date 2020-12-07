package test

import (
	"github.com/blockchain-jd-com/framework-go/crypto"
	"github.com/blockchain-jd-com/framework-go/ledger_model"
	"github.com/blockchain-jd-com/framework-go/utils/base58"
	"io/ioutil"
	"os"
)

/*
 * Author: imuge
 * Date: 2020/5/29 下午3:21
 */

var (
	// 网关服务IP
	GATEWAY_HOST = "localhost"
	// 网关服务端口
	GATEWAY_PORT = 8080
	SECURE       = false

	// 区块链上已存在的有操作权限的用户公私钥信息
	NODE_PRIVITE_KEY = crypto.DecodePrivKey(string(MustLoadFile("nodes/peer0/config/keys/jd.priv")), base58.MustDecode(string(MustLoadFile("nodes/peer0/config/keys/jd.pwd"))))
	NODE_PUBLIC_KEY  = crypto.DecodePubKey(string(MustLoadFile("nodes/peer0/config/keys/jd.pub")))
	NODE_KEY         = ledger_model.NewBlockchainKeypair(NODE_PUBLIC_KEY, NODE_PRIVITE_KEY)
)

func MustLoadFile(fileName string) []byte {
	file, _ := os.Open(fileName)
	bytes, _ := ioutil.ReadAll(file)

	return bytes
}

func LoadFile(fileName string) ([]byte, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	bytes, err := ioutil.ReadAll(file)

	return bytes, err
}
