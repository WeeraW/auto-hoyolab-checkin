package configcheckin

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/WeeraW/auto-hoyolab-checkin/app/servicelogger"
)

var ConfigData Config

func ReadConfiguration() error {
	servicelogger.Info("Reading configuration...")
	if _, err := os.Stat("config.json"); err == nil {
	} else {
		fmt.Println("Configuration file not found, creating new one...")
		configMap, err := Config{}.NewDefaultConfig()
		if err != nil {
			return err
		}
		jsonByte, err := json.MarshalIndent(configMap, "", " ")
		if err != nil {
			return err
		}
		_ = os.WriteFile("config.json", jsonByte, 0644)
	}

	jsonFile, _ := os.Open("config.json")
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	var result Config
	err = json.Unmarshal(byteValue, &result)
	if err != nil {
		return err
	}
	ConfigData = result
	servicelogger.Info("Configuration loaded!")
	return nil
}
