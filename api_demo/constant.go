package sdk

//充值订单
const CREATE_BILL = "/v1/create_bill"

//提币订单
const WITHDRAW = "/v1/withdraw"

//账户查询
const ACCOUNTS_INFO = "/v1/accounts_info"

//充值订单
const FINDINFO = "/v1/findinfo"

type ResultMsg struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type CreateBill struct {
	AppId  string `json:"appId" structs:"appId"`   // appId   ；参与签名
	Oid    string `json:"oid" structs:"oid"`       // 客户自身平台订单号（唯一） ；参与签名
	Amount string `json:"amount" structs:"amount"` // 充值数量(usdt) ；参与签名
}

type FindInfo struct {
	AppId  string `json:"appId" structs:"appId"`   // appId   ；参与签名
	Tid    string `json:"tid" structs:"tid"`       // 平台订单号（唯一） ；参与签名
}
