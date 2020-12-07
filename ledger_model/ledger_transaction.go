package ledger_model

import binary_proto "github.com/blockchain-jd-com/framework-go/binary-proto"

/*
 * Author: imuge
 * Date: 2020/5/27 下午2:37
 */

var _ binary_proto.DataContract = (*LedgerTransaction)(nil)

func init() {
	binary_proto.Cdc.RegisterContract(LedgerTransaction{})
}

type LedgerTransaction struct {
	Request TransactionRequest `refContract:"514"`
	Result  TransactionResult  `refContract:"515"`
}

func (l LedgerTransaction) ContractCode() int32 {
	return binary_proto.TX_RECORD
}

func (l LedgerTransaction) ContractName() string {
	return "LedgerTransaction"
}

func (l LedgerTransaction) Description() string {
	return ""
}
