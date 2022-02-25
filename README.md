## 一、说明
- 该SDK是对 dogpay api接口文档的Go版封装，如有疑问，可咨询官网客服

## 二、使用步骤

### 2.1 获取`sdk`

go get gitlab.codehub.com/dogpay/go-demo


```go
package main

import (
	sdk "gitlab.codehub.com/dogpay/go-demo"
	"log"
)
//创建订单
func main() {
	dogPayClient := sdk.NewDogPayClient("http://127.0.0.1:19095", "a9d3beb06924619306a93a9d206cfd31", "22be5b0d6564fcb1de976d1c5916e87b")
	bill := sdk.CreateBill{
		AppId:  "a9d3beb06924619306a93a9d206cfd31",
		Oid:    "sakldfjklasdjfkla",
		Amount: "123",
	}

	resultMsg, err := dogPayClient.CreateBill(&bill)
	if err!= nil {
		log.Println(err)

		return
	}
	//TODO 处理逻辑

	log.Printf("%#v",resultMsg)
}

```

## 三、关于回调

```go
package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/recharge", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != "POST" {
			return
		}
		log.Println("访问充币回调接口")
		bytes, err := ioutil.ReadAll(request.Body)
		defer request.Body.Close()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(string(bytes))
		//todo 校验逻辑
		//......
		writer.Write([]byte("SUCCESS"))
	})
	http.HandleFunc("/withdraw", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != "POST" {
			return
		}
		log.Println("访问提币回调接口")
		bytes, err := ioutil.ReadAll(request.Body)
		defer request.Body.Close()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(string(bytes))
		//todo 校验逻辑
		//......
		writer.Write([]byte("SUCCESS"))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

```
