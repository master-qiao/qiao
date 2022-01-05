package util

import "io/ioutil"

//InitConfigFile 初始化配置文件信息
func InitConfigFile(configFilePath string, out interface{}) error {
	content, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return err
	}
	err = FromByteJSON(content, out)
	if err != nil {
		return err
	}
	return nil
}