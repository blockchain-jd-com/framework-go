package ledger_model

const (
	// 操作类型名称 TODO 不全
	OperationTypeUserRegister                    = "UserRegisterOperation"           // 用户注册
	OperationTypeDataAccountRegisterOperation    = "DataAccountRegisterOperation"    // 数据账户注册
	OperationTypeDataAccountKVSetOperation       = "DataAccountKVSetOperation"       // KV写入
	OperationTypeEventAccountRegisterOperation   = "EventAccountRegisterOperation"   // 事件账户注册
	OperationTypeEventPublishOperation           = "EventPublishOperation"           // 事件发布
	OperationTypeParticipantRegisterOperation    = "ParticipantRegisterOperation"    // 参与方注册
	OperationTypeParticipantStateUpdateOperation = "ParticipantStateUpdateOperation" // 参与方变更
	OperationTypeContractCodeDeployOperation     = "ContractCodeDeployOperation"     // 合约部署
	OperationTypeContractEventSendOperation      = "ContractEventSendOperation"      // 合约调用
	OperationTypeRolesConfigureOperation         = "RolesConfigureOperation"         // 角色配置
	OperationTypeUserAuthorizeOperation          = "UserAuthorizeOperation"          // 用户授权
)
