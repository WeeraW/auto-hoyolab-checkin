package configcheckin

import (
	"encoding/json"
	"fmt"
)

type SignReqMethod string

func (s SignReqMethod) String() string {
	return string(s)
}

const (
	GET  SignReqMethod = "GET"
	POST SignReqMethod = "POST"
)

type CheckinConfig struct {
	GameName      string        `json:"GameName"`
	Enable        bool          `json:"Enable"`
	ActId         string        `json:"ActId"`
	SignReqMethod SignReqMethod `json:"SignReqMethod"`
	SignUrl       string        `json:"SignUrl"`
	InfoReqMethod SignReqMethod `json:"InfoReqMethod"`
	InfoUrl       string        `json:"InfoUrl"`
}

type MessageMode string

const (
	// SilentMode is a flag to enable or disable the service message. messages will be shown when error is occurred.
	SilentMode MessageMode = "slilent"
	// VerboseMode is a flag to enable or disable the service message. messages will be shown every action.
	VerboseMode MessageMode = "verbose"
	// SummaryMode is a flag to enable or disable the service message. messages will be shown only checking in result.
	SummaryMode MessageMode = "summary"
)

type Config struct {
	AutoHideWindow  bool          `json:"AutoHideWindow" default:"true"`
	LogToFile       bool          `json:"LogToFile"`
	MessageMode     MessageMode   `json:"MessageMode"`
	GenshinImpact   CheckinConfig `json:"GenshinImpact"`
	HonkaiImpact3   CheckinConfig `json:"HonkaiImpact3"`
	HonkaiStarRail  CheckinConfig `json:"HonkaiStarRail"`
	ZenlessZoneZero CheckinConfig `json:"ZenlessZoneZero"`
}

// NewConfig creates a new Config from the given JSON data.
func (Config) NewConfig(configJSONRawData []byte) (result Config, err error) {
	result = Config{}
	err = json.Unmarshal(configJSONRawData, &result)
	return result, err
}

func (c *Config) Inspector() string {
	return fmt.Sprintf(
		"AutoHideWindow: %v\nLogToFile: %v\nMessageMode: %v\nGenshinImpact: %v\nHonkaiImpact3: %v\nHonkaiStarRail: %v\nZenlessZoneZero: %v\n",
		c.AutoHideWindow, c.LogToFile, c.MessageMode, c.GenshinImpact, c.HonkaiImpact3, c.HonkaiStarRail, c.ZenlessZoneZero,
	)
}

// NewDefaultConfig creates a new Config from the default JSON data.
func (Config) NewDefaultConfig() (result Config, err error) {
	return DefaultConfigStruct, nil
}
