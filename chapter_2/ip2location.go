package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// 查询IP地址归属地 Go示例代码
func main() {
	url := "https://ipapi.api.bdymkt.com/ip2location/retrieve"
	body := "{\n  \"ip\": \"12.4.12.4\"\n}"
	req, err := http.NewRequest("POST", url, strings.NewReader(body))

	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	req.Header.Add("X-Bce-Signature", "AppCode/9342a6e19426476f9819b532893b37cc")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status: ", resp.Status)
	fmt.Println("Response headers: ", resp.Header)
	respBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Response body: ", string(respBody))
}
