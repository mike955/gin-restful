/*=============================================================================
#     FileName: config.go
#         Desc: parse config file
#       Author: ato.ye
#        Email: ato.ye@ucloud.cn
#     HomePage: http://www.ucloud.cn
#      Version: 0.0.1
#   LastChange: 2016-04-20 20:18:19
#      History:
=============================================================================*/
package ufcommon

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

var (
	configContentByte []byte
	configContentJson map[string]interface{}
)

type block struct {
	data interface{}
}

func getConfig(data map[string]interface{}) (*block, error) {
	if data == nil {
		return nil, errors.New("get config fail, config content is nil")
	}
	return &block{
		data: data,
	}, nil
}

func (b *block) getValue(key string) *block {
	m := b.getData()
	if v, ok := m[key]; ok {
		b.data = v
		return b
	}
	return nil
}

func (b *block) getData() map[string]interface{} {
	if m, ok := (b.data).(map[string]interface{}); ok {
		return m
	}
	return nil
}

// 读取配置
func LoadConfigFromFile(filename string) (err error) {
	configContentByte, err = ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	return json.Unmarshal(configContentByte, &configContentJson)
}

func LoadConfigFromData(data []byte) (err error) {
	configContentByte = data
	return json.Unmarshal(configContentByte, &configContentJson)
}

// 获取配置value,支持按层次获取，点号分割
func GetConfigByKey(keys string) (value interface{}, err error) {
	key_list := strings.Split(keys, ".")
	block, err := getConfig(configContentJson)
	if err != nil {
		return nil, err
	}
	for _, key := range key_list {
		block = block.getValue(key)
		if block == nil {
			return nil, fmt.Errorf("can not get[\"%s\"]'s value", string(keys))
		}
	}
	return block.data, nil
}

// 获取配置value,支持按层次获取，点号分割
// 返回Float64类型
func GetConfigByKeyToFloat(keys string) (value float64, err error) {
	key_list := strings.Split(keys, ".")
	block, err := getConfig(configContentJson)
	if err != nil {
		return 0.0, err
	}
	for _, key := range key_list {
		block = block.getValue(key)
		if block == nil {
			return 0.0, fmt.Errorf("can not get[\"%s\"]'s value", string(keys))
		}
	}

	value, ok := block.data.(float64)
	if !ok {
		return 0.0, fmt.Errorf("[\"%s\"]'s value is not a number", string(keys))
	}
	return value, nil
}

// 获取配置value,支持按层次获取，点号分割
// 返回Int类型
func GetConfigByKeyToInt(keys string) (value int, err error) {
	v, err := GetConfigByKeyToFloat(keys)
	if err != nil {
		return 0, err
	}
	return int(v), nil
}

// 获取配置value,支持按层次获取，点号分割
// 返回String类型
func GetConfigByKeyToString(keys string) (value string, err error) {
	key_list := strings.Split(keys, ".")
	block, err := getConfig(configContentJson)
	if err != nil {
		return "", err
	}
	for _, key := range key_list {
		block = block.getValue(key)
		if block == nil {
			return "", fmt.Errorf("can not get[\"%s\"]'s value", string(keys))
		}
	}

	value, ok := block.data.(string)
	if !ok {
		return "", fmt.Errorf("[\"%s\"]'s value is not a string", string(keys))
	}
	return value, nil
}

// 获取配置value,支持按层次获取，点号分割
// 返回Bool
func GetConfigByKeyToBool(keys string) (value interface{}, err error) {
	key_list := strings.Split(keys, ".")
	block, err := getConfig(configContentJson)
	if err != nil {
		return nil, err
	}
	for _, key := range key_list {
		block = block.getValue(key)
		if block == nil {
			return nil, fmt.Errorf("can not get[\"%s\"]'s value", string(keys))
		}
	}

	value, ok := block.data.(bool)
	if !ok {
		return "", fmt.Errorf("[\"%s\"]'s value is not a bool", string(keys))
	}
	return value, nil
}

func DumpConfigContent() {
	var pjson bytes.Buffer
	json.Indent(&pjson, configContentByte, "", "\t")
	fmt.Println(string(pjson.Bytes()))
}
