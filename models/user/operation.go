package user

import (
	"encoding/json"
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
	if !tool.IsExists(path) { // 如果不存在

		// 创建对应的文件
		var file, err = os.Create(path)
		defer file.Close()

		if err != nil {
			return err
		}

		// 初始化配置
		data, err := json.Marshal(defaultConfig)
		if err != nil {
			return err
		}

		// 写入配置
		file.Write(data)
	}

	// 加载配置文件
	return tool.LoadJson(path, &config)
}

func (user *User) HasExist(data Data) int {
	for index, elem := range user.Data {
		if elem.Md5 == data.Md5 {
			return index
		}
	}

	return -1
}

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
