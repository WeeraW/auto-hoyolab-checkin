package configcheckin

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/WeeraW/auto-hoyolab-checkin/app/servicelogger"
)

var ConfigData Config

func ReadConfiguration() error {
	servicelogger.Info("Reading configuration...")
	// log current directory
	dir, err := os.Getwd()
	if err != nil {
		servicelogger.Error(fmt.Sprintf("Failed to get current directory: %v", err))
	} else {
		servicelogger.Infof("Current directory: %s", dir)
	}

	// if current directory is not the same as the directory of the executable, change to the executable's directory
	exePath, err := os.Executable()
	if err != nil {
		servicelogger.Error(fmt.Sprintf("Failed to get executable path: %v", err))
	} else {
		exeDir := filepath.Dir(exePath)
		if dir != exeDir {
			servicelogger.Infof("Changing directory to executable's directory: %s", exeDir)
			err = os.Chdir(exeDir)
			if err != nil {
				servicelogger.Error(fmt.Sprintf("Failed to change directory: %v", err))
			}
		}
	}

	if _, err := os.Stat("config.json"); err == nil {
		fmt.Println("Configuration file found, loading...")
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

	servicelogger.LogToFile = ConfigData.LogToFile

	servicelogger.Info("Configuration loaded!")
	servicelogger.Infof("Configuration: %s", ConfigData.Inspector())
	return nil
}

func SaveConfiguration() error {
	servicelogger.Info("Saving configuration...")
	jsonByte, err := json.MarshalIndent(ConfigData, "", " ")
	if err != nil {
		return err
	}
	_ = os.WriteFile("config.json", jsonByte, 0644)
	servicelogger.Info("Configuration saved!")
	return nil
}

func SetMessageMode(mode MessageMode) {
	ConfigData.MessageMode = mode
	_ = SaveConfiguration()
}
