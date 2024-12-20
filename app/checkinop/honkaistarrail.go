package checkinop

import (
	"github.com/WeeraW/auto-hoyolab-checkin/app/configcheckin"
	"github.com/WeeraW/auto-hoyolab-checkin/app/cookiereader"
)

func CheckinHonkaiStarRail(cookie cookiereader.CheckInCookieV2) (message string, err error) {
	if !configcheckin.ConfigData.HonkaiStarRail.Enable {
		return "", nil
	}

	return DoCheckIn(cookie, configcheckin.ConfigData.HonkaiStarRail)
}
