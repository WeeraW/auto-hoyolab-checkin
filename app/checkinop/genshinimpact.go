package checkinop

import "github.com/brokiem/auto-hoyolab-checkin/app/configcheckin"

func CheckinGenshinImpact() (message string, err error) {
	if !configcheckin.ConfigData.GenshinImpact.Enable {
		return "", nil
	}

	return DoCheckIn(configcheckin.ConfigData.GenshinImpact)
}
