package blockchain

import (
	"fmt"
	"goProject/models/tool"

	"github.com/beego/beego/v2/client/httplib"
)

/*
验证下载请求是否合法

param:
	blockchainURL string
	address string
	target int
	ip string

return:
	VerifyResult {
		state: true/false
	}
*/
func VerifyDownloadRequest(blockchainURL string, address string, target int, ip string, dataType uint) VerifyResult {
	var body = VerifyRequestBody{
		Address: address,
		Target:  target,
		Ip:      ip,
		Type:    dataType,
	}

	var response VerifyResult
	var request = httplib.Post(blockchainURL + "/verify")
	request.JSONBody(body)
	request.ToJSON(&response)

	fmt.Printf("Verify result: %v", tool.ToJSONString(response))
	return response
}

/*
向链端添加日志

param:
	blockchainURL string 链端 URL
	address string 链端参数
	target int 链端参数
	ip string 链端参数
	time string 链端参数

returns:
	type AddLogResult struct {
		State   bool   `json:"state"`
		Message string `json:"message"`
	}
*/
func AddLog(blockchainURL string, address string, target int, ip string, time string) AddLogResult {
	var response AddLogResult
	var request = httplib.Post(blockchainURL + "/log/public")
	var body = AddLogRequestBody{
		Address: address,
		Target:  target,
		Ip:      ip,
		Time:    time,
	}

	request.JSONBody(body)
	request.ToJSON(&response)

	fmt.Printf("Add log request: %v", tool.ToJSONString(body))
	fmt.Printf("Add log result: %v", tool.ToJSONString(response))
	return response
}

/*
获取链端的日志
*/
func GetLogs(blockchainURL string, serialNum int) GetLogsResult {
	var response GetLogsResult
	var request = httplib.Get(blockchainURL + "/getLogs")
	var body = GetLogsRequestBody{
		SerialNum: serialNum,
	}

	request.JSONBody(body)
	request.ToJSON(&response)

	fmt.Printf("Get logs result: %v", tool.ToJSONString(response))
	return response
}
