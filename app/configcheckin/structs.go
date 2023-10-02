package configcheckin

import "encoding/json"

type CheckinConfig struct {
	GameName string `json:"GameName"`
	Enable   bool   `json:"Enable"`
	ActId    string `json:"ActId"`
	SignUrl  string `json:"SignUrl"`
	InfoUrl  string `json:"InfoUrl"`
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
	return DefaultConfigStruct, nil
}
