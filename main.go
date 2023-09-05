package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Code       int         `json:"code"`
	Data       interface{} `json:"data"`
	ErrMessage string      `json:"err_message"`
}

func main() {
	//post http请求，自定义请求头
	httpPost()
}

func httpPost() {
	//post http请求，自定义请求头
	url := "http://localhost:8080/Insert"
	token := "your-x-token"

	// 创建一个请求体
	body := make(map[string]string)
	body["tableName"] = "goods_info"
	body["columnName"] = "id,name,yj,hyj,jj,cls,stock"
	body["data"] = "(10,'cs',1,1,1,'1',1), (11,'cs',1,1,1,'1',1), (12,'cs',1,1,1,'1',1)"

	//map转[]byte
	bodyByte, _ := json.Marshal(body)

	// 创建一个请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyByte))
	if err != nil {
		fmt.Println("请求创建失败:", err)
		return
	}

	// 设置请求头中的X-Token字段
	req.Header.Set("X-Token", token)

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("请求发送失败:", err)
		return
	}

	// 解析JSON响应体
	var response Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		fmt.Println("JSON解析失败:", err)
		return
	}

	fmt.Println("响应状态码:", resp.StatusCode)
	fmt.Println("响处理结果:", response)

}
