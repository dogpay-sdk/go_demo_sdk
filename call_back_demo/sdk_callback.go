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
