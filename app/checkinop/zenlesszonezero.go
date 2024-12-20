package checkinop

import (
	"github.com/WeeraW/auto-hoyolab-checkin/app/configcheckin"
	"github.com/WeeraW/auto-hoyolab-checkin/app/cookiereader"
)

func CheckinZenLessZoneZero(cookie cookiereader.CheckInCookieV2) (message string, err error) {
	if !configcheckin.ConfigData.ZenlessZoneZero.Enable || !cookie.IsEligbleForCheckinZZZ() {
		return "", nil
	}

	return DoCheckIn(cookie, configcheckin.ConfigData.ZenlessZoneZero)
}
