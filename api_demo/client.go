package sdk

import (
	"errors"
	"github.com/fatih/structs"
)

type DogPayClient struct {
	gateway string //网关
	appId   string //商户AppId编号
	appKey  string //商户Appkey
}

func NewDogPayClient(gateway string, appId string, appKey string) *DogPayClient {
	return &DogPayClient{
		gateway: gateway,
		appId:   appId,
		appKey:  appKey,
	}
}

//充币
func (u *DogPayClient) CreateBill(bill *CreateBill) (*ResultMsg, error) {
	if bill == nil {
		return nil, errors.New("参数异常")
	}
	data := structs.Map(bill)

	return dogPayPost(u.appKey, u.gateway, CREATE_BILL, data )
}


//账单查询
func (u *DogPayClient) FindInfo(info *FindInfo) (*ResultMsg, error) {
	if info == nil {
		return nil, errors.New("参数异常")
	}
	data := structs.Map(info)

	return dogPayPost(u.appKey, u.gateway, FINDINFO, data )
}
