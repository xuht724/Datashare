package tool

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/beego/beego/v2/server/web/context"
)

/*
获取 body 中的参数

param:
	ctx *context.Context 上下文
	body interface{} 相关格式

return:
	error
*/
func Param(ctx *context.Context, body interface{}) error {
	var error = json.Unmarshal(ctx.Input.RequestBody, &body)

	if error != nil {
		return error
	}
	return nil
}

/*
加载 JSON 文件

param:
	path string 文件路径
	t interface{} JSON 格式
*/
func LoadJson(path string, t interface{}) error {
	var data, err = ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, t)
	return err
}

/*
保存为 JSON 文件

param:
	path string 文件路径
	JSON interface{} 具体的文件数据

return:
	err error
*/
func SaveJSON(path string, JSON interface{}) error {

	// 创建对应的文件
	var file, err = os.Create(path)
	defer file.Close()

	if err != nil {
		return err
	}

	var jsonString = ToJSONString(JSON)

	// 写入配置
	_, err = file.Write([]byte(jsonString))

	return err
}

/*
判断文件是否存在

param:
	path string 文件路径

return:
	bool 如果存在则返回 true，否则 false
*/
func IsExists(path string) bool {
	var _, err = os.Stat(path)

	if err != nil {
		return os.IsExist(err)
	}

	return true
}

/*
发送 JSON 格式 body 的响应
*/
func SendJsonResponse(ctx *context.Context, data interface{}) error {
	var responseBody, err = json.Marshal(data)

	if err != nil {
		ctx.Output.SetStatus(202)
		ctx.Output.Body([]byte("ERROR"))
		return err
	}

	return ctx.Output.Body(responseBody)
}

/*
封装好的，用于 defer 来发送 JSON 格式 body 的响应
*/
func DeferSendJsonResponse(ctx *context.Context, data interface{}) {
	if err := SendJsonResponse(ctx, data); err != nil {
		panic(err)
	}
}

func ToJSONString(v interface{}) string {
	var bs, err = json.Marshal(v)

	if err != nil {
		panic(err)
	}
	var stdout bytes.Buffer
	json.Indent(&stdout, bs, "", "  ")
	return stdout.String() + "\n"
}
