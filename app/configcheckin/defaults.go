package configcheckin

var DefaultConfigStruct = Config{
	AutoHideWindow: true,
	GenshinImpact: CheckinConfig{
		GameName: "Genshin Impact",
		Enable:   true,
		ActId:    "e202102251931481",
		SignUrl:  "https://sg-hk4e-api.hoyolab.com/event/sol/sign",
		InfoUrl:  "https://sg-hk4e-api.hoyolab.com/event/sol/info",
	},
	HonkaiImpact3: CheckinConfig{
		GameName: "Honkai Impact 3",
		Enable:   true,
		ActId:    "e202110291205111",
		SignUrl:  "https://sg-public-api.hoyolab.com/event/mani/sign",
		InfoUrl:  "https://sg-public-api.hoyolab.com/event/mani/info",
	},
	HonkaiStarRail: CheckinConfig{
		GameName: "Honkai Star Rail",
		Enable:   true,
		ActId:    "e202303301540311",
		SignUrl:  "https://sg-public-api.hoyolab.com/event/luna/os/sign",
		InfoUrl:  "https://sg-public-api.hoyolab.com/event/luna/os/info",
	},
}
