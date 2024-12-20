package checkinop

import (
	"github.com/WeeraW/auto-hoyolab-checkin/app/configcheckin"
	"github.com/WeeraW/auto-hoyolab-checkin/app/cookiereader"
)

func CheckinGenshinImpact(cookie cookiereader.CheckInCookieV2) (message string, err error) {
	if !configcheckin.ConfigData.GenshinImpact.Enable {
		return "", nil
	}

	return DoCheckIn(cookie, configcheckin.ConfigData.GenshinImpact)
}
