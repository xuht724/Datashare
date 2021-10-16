package goFastdfs

import (
	"fmt"
	"goProject/models/tool"
	"reflect"

	"github.com/beego/beego/v2/client/httplib"
)

/*
获取文件统计信息

param:
	hostURL string: 储存服务器 URL

return:
	response Status

返回实例:
	[
		message: ""
		status: "ok"
		data: [
			{date: 20210915 fileCount: 1 totalSize: 584},
			{date: all fileCount: 1 totalSize: 584}
		]
	]
*/
func GetStatus(hostURL string) Status {
	var response Status
	var request = httplib.Get(hostURL + "/stat")
	request.ToJSON(&response)
	fmt.Println(reflect.ValueOf(response).Kind())
	return response
}

/*
文件上传函数

param:
	filename string: 上传的文件名
	scene string: 场景
	postPath string: 自定义路径
	hostURL string: 储存服务器 URL

return:
	response UploadResult

返回实例:
	[
		retmsg:
		retcode: 0
		path:  /group1/default/20210915/11/25/2/login.html
		size:  584
		domain:  http://localhost:8080
		scene:  default
		mtime:  1.631676311e+09
		scenes:  default
		src:  /group1/default/20210915/11/25/2/login.html
		url:  http://localhost:8080/group1/default/20210915/11/25/2/login.html?name=login.html&download=1
		md5:  6054b52e6981f9960fcf334b0ddb72e9
	]
*/
func UploadFile(filename string, hostURL string) UploadResult {
	// fmt.Println(filename, scene, postPath, hostURL)
	var request = httplib.Post(hostURL + "/upload")
	request.PostFile("file", filename)
	request.Param("output", "json")

	var uploadResult UploadResult
	request.ToJSON(&uploadResult)

	return uploadResult
}

/*
删除一个文件

param:
	hostURL string: 服务器地址
	md5 string: 文件 md5

return:
	response DeleteResult

返回实例:
	{
		Data: ""
		Status: "ok"
		Message: "remove success"
	}
*/
func DeleteOneFile(hostURL string, md5 string) DeleteResult {
	var response DeleteResult
	var url = hostURL + "/delete?md5=" + string(md5)
	var request = httplib.Delete(url)
	fmt.Println(url)
	request.ToJSON(&response)
	return response
}

/*
获取文件信息

param:
	hostURL string: 服务器地址
	md5 string: 文件 md5

return:
	response FileInfoResult

返回实例:
	{
		Data: {
			Name: "login.html"
			Rename: ""
			Path: "files/default/20210915/17/51/3"
			Md5: "6054b52e6981f9960fcf334b0ddb72e9"
			Size: 584
			Peers: [
				"http://172.18.42.27:8080"
			]
			Scene: "default"
			TimeStamp: 1631699507
			Offset: -1
		}
		Status: "ok"
		Message: ""
	}
*/
func GetFileInfo(hostURL string, md5 string) FileInfoResult {
	var response FileInfoResult
	var url = hostURL + "/get_file_info?md5=" + md5
	var request = httplib.Get(url)
	fmt.Println(url)
	request.ToJSON(&response)
	return response
}

/*
获取文件夹信息

param:
	hostURL string: 服务器地址
	dirname string: 文件夹名

return:
	response DirInfoResult

返回实例:
	{
		Data: [
			{
				Name: "_big"
				Md5: "78ecb468e7f5e66ec081be9f3fc7863a"
				Path: ""
				Size: 0
				Mtime: 1631698634
				IsDir: true
			},
			{
				Name: "_tmp"
				Md5: "c495b986a2bd8d318c92825ed3563ec9"
				Path: ""
				Size: 0
				Mtime: 1631698641
				IsDir: true
			},
			{
				Name: "default"
				Md5: "25e702e295075eb9b80be634aa88b301"
				Path: ""
				Size: 0
				Mtime: 1631698641
				IsDir: true
			}
		],
		Status: "ok"
		Message: ""
	}
*/
func GetDirInfo(hostURL string, dirname string) DirInfoResult {
	var response DirInfoResult
	var url = hostURL + "/?dir=" + dirname
	var request = httplib.Get(url)
	fmt.Println(url)
	request.ToJSON(&response)
	return response
}

/*
获取下载链接（分布式服务器的）

param:
	hostURL string 分布式储存服务器的地址
	md5 string 文件的 md5

return:
	usl string 文件的下载 url
	filename string 文件名
*/
func GetDownloadURL(hostURL string, md5 string) (url string, filename string) {
	var response = GetFileInfo(hostURL, md5)
	if response.Status == "fail" {
		return "no such file", ""
	}

	fmt.Printf("Get download url [from gofastdfs] response: %v", tool.ToJSONString(response))

	if response.Data.Md5 == "" {
		return "", response.Data.Name
	}

	return hostURL + "/" + response.Data.Path[6:] + "/" + response.Data.Name + "?name=" + response.Data.Name + "&download=1", response.Data.Name

	// "http://localhost:8080/group1/default/20210915/17/51/3/login.html?name=login.html&download=1"
}

/*
下载文件函数

param:
	hosrURL string: 服务器地址
	md5 string: 文件 md5
	savePath string: 储存路径
*/
func DownloadFile(hostURL string, md5 string, path string) (downloadStatus bool, fileURL string) {
	var downloadURL, filename = GetDownloadURL(hostURL, md5)
	var request = httplib.Get(downloadURL)
	fileURL = path + filename
	var error = request.ToFile(fileURL)

	return error != nil, fileURL
}
