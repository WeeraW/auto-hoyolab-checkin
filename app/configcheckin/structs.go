package configcheckin

import "encoding/json"

type CheckinConfig struct {
	Enable bool   `json:"Enable"`
	ActId  string `json:"ActId"`
}
type Config struct {
	AutoHideWindow bool          `json:"AutoHideWindow"`
	GenshinImpact  CheckinConfig `json:"GenshinImpact"`
	HonkaiImpact3  CheckinConfig `json:"HonkaiImpact3"`
	HonkaiStarRail CheckinConfig `json:"HonkaiStarRail"`
}

// NewConfig creates a new Config from the given JSON data.
func (Config) NewConfig(configJSONRawData []byte) (result Config, err error) {
	result = Config{}
	err = json.Unmarshal(configJSONRawData, &result)
	return result, err
}

// NewDefaultConfig creates a new Config from the default JSON data.
func (Config) NewDefaultConfig() (result Config, err error) {
	result = Config{}
	err = json.Unmarshal([]byte(DefaultConfig), &result)
	return result, err
}
