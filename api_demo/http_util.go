package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gitlab.codehub.com/dogpay/go-demo/util"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"time"
)

func dogPayPost(appKey, gateway, path string, param map[string]interface{}) (*ResultMsg, error) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	for k, v := range param {
		writer.WriteField(k, fmt.Sprintf("%s", v))
	}
	bytes, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	paramData := util.GenerateMd5Sign(string(bytes), appKey)
	writer.WriteField("sign", paramData)
	err = writer.Close()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	client := &http.Client{
		Timeout: time.Second * 5,
	}
	req, err := http.NewRequest("POST", gateway+path, payload)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	result := new(ResultMsg)
	err = json.Unmarshal(bodyText, &result)
	if err != nil {
		return nil, err
	}
	return result, err

}
