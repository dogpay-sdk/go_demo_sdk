package main

import (
	sdk "gitlab.codehub.com/dogpay/go-demo/api_demo"
	"log"
)

func main() {
	dogPayClient := sdk.NewDogPayClient("http://192.168.5.17:19095", "xxxx", "xxxxx")
	getAddress := sdk.GetAddress{
		AppId: dogPayClient.GetAppId(),
		Oid:   "sakldfjklasdjfkla",
	}

	resultMsg, err := dogPayClient.GetAddress(&getAddress)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resultMsg)
	////TODO 处理逻辑
	//
	//log.Printf("%#v", resultMsg)
	findInfo, err := dogPayClient.FindInfo(&sdk.FindInfo{
		AppId: dogPayClient.GetAppId(),
	})
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("%#v", findInfo)
}
