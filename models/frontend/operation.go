package frontend

import (
	"encoding/json"
	"errors"
	"fmt"
	"goProject/models/blockchain"
	"goProject/models/goFastdfs"
	"goProject/models/tool"
	"goProject/models/user"
	"io/ioutil"
	"os"
	"time"

	"github.com/beego/beego/v2/server/web/context"
)

/*
前端将需要上传的文件提交给后端
*/
func UploadFile(ctx *context.Context) {
	var body = UploadToDfsResult{} // 返回请求体

	// 响应
	defer (func(ctx *context.Context) {
		responseBody, err := json.Marshal(body)

		if err != nil {
			ctx.Output.SetStatus(202)
			ctx.Output.Body([]byte("ERROR"))
			panic(err)
		}
		ctx.Output.Body(responseBody)
		fmt.Printf("Upload file response: %v\n", tool.ToJSONString(body))
	})(ctx)

	// 接收文件
	ctx.Request.ParseForm()
	file, handler, _ := ctx.Request.FormFile("file")
	var filestream, _ = ioutil.ReadAll(file)

	// 转存到本地
	var path = user.Conf.FilePath + handler.Filename
	ioutil.WriteFile(path, filestream, 0777)

	defer os.Remove(path) // 后端记得删除文件

	// 上传到分布式储存
	var uploadResult = goFastdfs.UploadFile(path, user.Conf.DfsURL)

	// 验证上传是否成功，并处理失败的情况
	if uploadResult.Md5 == "" {
		body.Status = "Failed"
		body.Message = "Fail to upload to dfs"
		body.Md5 = ""

		var res = fmt.Sprintf("%+v", uploadResult)
		panic(res)
	} else {
		body.Md5 = uploadResult.Md5
	}

	// 加载本地配置，但为了保证唯一性，考虑不加载
	// tool.LoadJson(user.UserFile, user.UserInfo)

	var newData = user.Data{
		Md5:      uploadResult.Md5,
		Filename: handler.Filename,
		Time:     fmt.Sprint(time.Now().Unix()),
	}

	// 判断文件是否是更新
	if index := user.UserInfo.HasExist(newData); index != -1 {
		user.UserInfo.Data[index] = newData
		body.Message = "REPEATED"
	} else {
		user.UserInfo.Data = append(user.UserInfo.Data, newData)
		body.Message = "SUCCESS"
	}

	// 保存到本地文件
	err := tool.SaveJSON(user.UserFile, user.UserInfo)

	// 返回响应
	if err != nil {
		body.Status = "Failed"
		body.Message = "Fail to save user information"
	} else {
		body.Status = "OK"
	}
}

/*
获取当前账户在存储服务器上存储了的数据集的值
*/
func GetDataNum(ctx *context.Context) {
	tool.LoadJson(user.UserFile, &user.UserInfo)

	var body = GetDataNumResponseBody{
		DataNum: uint(len(user.UserInfo.Data)),
	}

	var responseBody, err = json.Marshal(body)
	if err != nil {
		panic(err)
	}

	ctx.Output.Body(responseBody)
	fmt.Printf("Get data num response: %v", tool.ToJSONString(body))
}

/*
获取当前账户在存储服务器上存储了的数据集的列表
*/
func GetDataList(ctx *context.Context) {
	tool.LoadJson(user.UserFile, &user.UserInfo)

	var body = GetDataListResponseBody{
		State: true,
		Data:  DataList(user.UserInfo.Data),
	}

	if err := tool.SendJsonResponse(ctx, body); err != nil {
		panic(err)
	}

	fmt.Printf("Get data list response: %v", tool.ToJSONString(body))
}

/*
获取文件下载链接
*/
func GetDownloadURL(ctx *context.Context) {

	// 验证下载请求
	var requestBody = GetDownloadRequestBody{}
	if err := tool.Param(ctx, &requestBody); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Get download url request: %v", tool.ToJSONString(requestBody))

	var responseBody = GetDownloadResponseBody{}

	if verifyResult := blockchain.VerifyDownloadRequest(user.Conf.BlockchainURL, requestBody.Address, requestBody.Target, requestBody.IP, requestBody.Type); !verifyResult.State {
		responseBody.Filename = ""
		responseBody.Md5 = ""
		responseBody.URL = ""
	} else {
		var url, filename = goFastdfs.GetDownloadURL(user.Conf.DfsURL, requestBody.Id)
		responseBody.Filename = filename
		responseBody.URL = url
		responseBody.Md5 = requestBody.Id
	}

	tool.SendJsonResponse(ctx, responseBody)
	fmt.Printf("Get download url response: %v", tool.ToJSONString(responseBody))

	// 进行日志上传
	if !(requestBody.Type == 1) {
		blockchain.AddLog(user.Conf.BlockchainURL, requestBody.Address, requestBody.Target, requestBody.IP, fmt.Sprint((time.Now().Unix())))
	}

}

/*
下载文件
*/
func Download(ctx *context.Context) {
	// 验证下载请求
	var requestBody = GetDownloadRequestBody{}
	tool.Param(ctx, &requestBody)
	fmt.Printf("Get download url request: %v", tool.ToJSONString(requestBody))

	// 验证下载连接诶
	if verifyResult := blockchain.VerifyDownloadRequest(user.Conf.BlockchainURL, requestBody.Address, requestBody.Target, requestBody.IP, requestBody.Type); !verifyResult.State {
		ctx.Output.SetStatus(403)
		ctx.Output.Body([]byte("Failed to verify"))
		return
	}

	// 下载到本地
	if status, fileURL := goFastdfs.DownloadFile(user.Conf.DfsURL, requestBody.Id, user.Conf.FilePath); !status {
		ctx.Output.SetStatus(404)
		ctx.Output.Body([]byte("Download failure"))
		return
	} else {

		ctx.Output.Download(fileURL)
		defer os.Remove(fileURL)

		// 进行日志上传
		if !(requestBody.Type == 1) {
			blockchain.AddLog(user.Conf.BlockchainURL, requestBody.Address, requestBody.Target, requestBody.IP, fmt.Sprint((time.Now().Unix())))
		}

	}

}

/*
删除某个文件
*/
func Delete(ctx *context.Context) {

	// 解析 request body
	var body = DeleteRequestBody{}
	var err = tool.Param(ctx, &body)
	if err != nil {
		panic(err)
	}

	// 向 dfs 进行 delete 请求
	var response = goFastdfs.DeleteOneFile(user.Conf.DfsURL, body.Md5)

	defer tool.DeferSendJsonResponse(ctx, response)

	// 如果 dfs 删除失败，报错并响应
	if response.Status != "ok" {
		response.Message = "fail to delete file in dfs"
		panic(errors.New("fail to delete file in dfs"))
	}

	// 删除 datalist 上的数据
	// 如果删除失败则报错并响应
	if err := user.UserInfo.Delete(body.Md5); err != nil {
		response.Status = "error"
		response.Message = err.Error()
		panic(err)
	}

	// 保存 user 的数据
	if err := tool.SaveJSON(user.UserFile, user.UserInfo); err != nil {
		panic(err)
	}

}
