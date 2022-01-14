package configmanager

import (
	"encoding/json"
	"io/ioutil"
)

type AppConfig struct {
	Port        string `json:"port"`
	LogFilePath string `json:"logFilePath"`
}

var (
	AppConf *AppConfig
)

func InitAppConfig(file string) error {
	raw, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(raw, &AppConf); err != nil {
		return err
	}
	return nil
}

func GetAppConfig() *AppConfig {
	return AppConf
}
