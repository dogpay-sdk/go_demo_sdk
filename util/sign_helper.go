package util

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"sort"
	"strconv"
)

/**
MD5签名
*/
func GenerateMd5Sign(jsonStr string, key string) string {
	return signMd5(generateClearTextSign(jsonStr, key))
}

/**
校验MD5签名
*/
func VerifyMd5Sign(jsonStr string, rawSign string, key string) bool {
	return signMd5(generateClearTextSign(jsonStr, key)) == rawSign
}

/**
 * 根据json串生成明文签名
 *
 * @param jsonStr
 * @return
 */
func generateClearTextSign(jsonStr string, key string) string {
	var tempMap map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &tempMap)
	if err != nil {
		log.Println(err)
		return ""
	}

	//不参与签名的字段
	var ignoreKeys = make(map[string]interface{})
	ignoreKeys["sign"] = 1
	ignoreKeys["sign_type"] = 1
	ignoreKeys["signType"] = 1
	ignoreKeys["gateway_url"] = 1

	var keys []string
	for k := range tempMap {
		//去掉无需签名的字段
		_, has := ignoreKeys[k]
		if !has {
			keys = append(keys, k)
		}
	}

	sort.Strings(keys)
	var text string
	for i := range keys {
		text += keys[i] + "=" + strval(tempMap[keys[i]]) + "&"
	}

	clearText := text + "key=" + key
	//log.Println("签名明文:", clearText)
	return clearText
}

// Strval 获取变量的字符串值
// 浮点型 3.0将会转换成字符串3, "3"
// 非数值或字符类型的变量将会被转换成JSON格式字符串
func strval(value interface{}) string {
	var key string
	if value == nil {
		return key
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}

	return key
}

/**
MD5加密
*/
func signMd5(str string) string {
	m := md5.New()
	_, err := io.WriteString(m, str)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	arr := m.Sum(nil)
	return fmt.Sprintf("%x", arr)
}
