package sdk

//获取充值地址
const GETADDRESS = "/v1/getAddress"

//提币订单
const WITHDRAW = "/v1/withdraw"

//账户查询
const ACCOUNTS_INFO = "/v1/accounts_info"

//流水日志查询（时间倒叙100条）
const RECHARGERECORD = "/v1/rechargeRecord"

type ResultMsg struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type GetAddress struct {
	AppId string `json:"appId" structs:"appId"` // appId   ；参与签名
	Oid   string `json:"uId" structs:"uId"`     // 客户自身平台订单号（唯一） ；参与签名
}

type FindInfo struct {
	AppId string `json:"appId" structs:"appId"` // appId   ；参与签名
}
