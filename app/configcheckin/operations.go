package configcheckin

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

var ConfigData Config

func ReadConfiguration() error {
	fmt.Println("Reading configuration...")
	if _, err := os.Stat("config.json"); err == nil {
	} else {
		fmt.Println("Configuration file not found, creating new one...")
		configMap, _ := Config{}.NewDefaultConfig()
		jsonByte, _ := json.MarshalIndent(configMap, "", " ")

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
	fmt.Println("Configuration loaded!")
	return nil
}
