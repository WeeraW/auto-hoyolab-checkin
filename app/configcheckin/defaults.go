package configcheckin

var DefaultConfigStruct = Config{
	AutoHideWindow: true,
	LogToFile:      true,
	MessageMode:    SummaryMode,
	GenshinImpact: CheckinConfig{
		GameName:      "Genshin Impact",
		Enable:        true,
		ActId:         "e202102251931481",
		SignReqMethod: POST,
		SignUrl:       "https://sg-hk4e-api.hoyolab.com/event/sol/sign",
		InfoReqMethod: GET,
		InfoUrl:       "https://sg-hk4e-api.hoyolab.com/event/sol/info",
	},
	HonkaiImpact3: CheckinConfig{
		GameName:      "Honkai Impact 3",
		Enable:        true,
		ActId:         "e202110291205111",
		SignReqMethod: POST,
		SignUrl:       "https://sg-public-api.hoyolab.com/event/mani/sign",
		InfoReqMethod: GET,
		InfoUrl:       "https://sg-public-api.hoyolab.com/event/mani/info",
	},
	HonkaiStarRail: CheckinConfig{
		GameName:      "Honkai Star Rail",
		Enable:        true,
		ActId:         "e202303301540311",
		SignReqMethod: POST,
		SignUrl:       "https://sg-public-api.hoyolab.com/event/luna/os/sign",
		InfoReqMethod: GET,
		InfoUrl:       "https://sg-public-api.hoyolab.com/event/luna/os/info",
	},
	ZenlessZoneZero: CheckinConfig{
		GameName:      "Zenless Zone Zero",
		Enable:        true,
		ActId:         "e202406031448091",
		SignReqMethod: POST,
		SignUrl:       "https://sg-public-api.hoyolab.com/event/luna/zzz/os/sign",
		InfoReqMethod: GET,
		InfoUrl:       "https://sg-public-api.hoyolab.com/event/luna/zzz/os/info",
	},
}
