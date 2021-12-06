package user

import (
	"errors"
	"goProject/models/tool"
	"os"
)

var Conf = Config{}
var UserInfo = User{}

/*
初始化
*/
func Init() {
	InitConfig(ConfigFile, DefaultConfig, &Conf)
	InitConfig(UserFile, DefaultUserInfo, &UserInfo)
	os.MkdirAll(Conf.FilePath, 0777)
}

/*
获取/生成配置文件

param:
	path string 文件路径
	defaultConfig interface{} 默认的配置选项
	config interface{} 用于获取配置数据
*/
func InitConfig(path string, defaultConfig interface{}, config interface{}) error {
	// 判断文件是否存在
	if !tool.IsExists(path) { // 如果不存在，则进行初始化
		tool.SaveJSON(path, defaultConfig)
	}

	// 加载配置文件
	return tool.LoadJson(path, &config)
}

/*
判断文件是否已经存在

param:
	data Data

return:
	int
*/
func (user *User) HasExist(data Data) int {
	for index, elem := range user.Data {
		if elem.Md5 == data.Md5 {
			return index
		}
	}

	return -1
}

/*
删除对应的文件

param:
	md5 string 文件的哈希

return:
	error
*/
func (user *User) Delete(md5 string) error {
	var index = -1
	for i, elem := range user.Data {
		if elem.Md5 == md5 {
			index = i
			break
		}
	}

	if index == -1 {
		return errors.New("no such file in datalist")
	}

	user.Data = append(user.Data[:index], user.Data[index+1:]...)
	return nil
}

/*
根据类型筛选已经上传的文件

param:
	dataType uint 文件类型

return:
	[]Data
*/
func (user *User) GetDataByType(dataType uint) []Data {
	var filteredData []Data = []Data{}

	for _, elem := range user.Data {
		if elem.Type == dataType {
			filteredData = append(filteredData, elem)
		}
	}

	return filteredData
}

/*
根据类型获取对应的文件数量

param:
	dataType uint 数据类型

return:
	uint 文件数目
*/
func (user *User) GetDataNumByType(dataType uint) uint {
	var dataNum = 0

	for _, elem := range user.Data {
		if elem.Type == dataType {
			dataNum++
		}
	}

	return uint(dataNum)
}
