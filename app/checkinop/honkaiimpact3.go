package checkinop

import (
	"github.com/WeeraW/auto-hoyolab-checkin/app/configcheckin"
	"github.com/WeeraW/auto-hoyolab-checkin/app/cookiereader"
)

func CheckinHonkaiImpact3(cookie cookiereader.CheckInCookieV2) (message string, err error) {
	if !configcheckin.ConfigData.HonkaiImpact3.Enable {
		return "", nil
	}

	return DoCheckIn(cookie, configcheckin.ConfigData.HonkaiImpact3)
}
