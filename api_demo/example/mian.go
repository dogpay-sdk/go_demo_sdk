package main

import (
	sdk "gitlab.codehub.com/dogpay/go-demo/api_demo"
	"log"
)

func main() {
	dogPayClient := sdk.NewDogPayClient("http://192.168.5.17:19095", "a9d3beb06924619306a93a9d206cfd31", "22be5b0d6564fcb1de976d1c5916e87b")
	//bill := sdk.CreateBill{
	//	AppId:  "a9d3beb06924619306a93a9d206cfd31",
	//	Oid:    "sakldfjklasdjfkla",
	//	Amount: "123",
	//}

	//resultMsg, err := dogPayClient.CreateBill(&bill)
	//if err != nil {
	//	log.Println(err)
	//
	//	return
	//}
	////TODO 处理逻辑
	//
	//log.Printf("%#v", resultMsg)
	findInfo, err := dogPayClient.FindInfo(&sdk.FindInfo{
		AppId: "a9d3beb06924619306a93a9d206cfd31",
		Tid:   "e356a78752c1e2e11dc5e052ec11c625",
	})
	if err!= nil {
		log.Println(err)
		return
	}

	log.Printf("%#v", findInfo)
}
