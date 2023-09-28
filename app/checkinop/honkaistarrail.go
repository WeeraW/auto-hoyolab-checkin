package checkinop

import "github.com/brokiem/auto-hoyolab-checkin/app/configcheckin"

func CheckinHonkaiStarRail() (message string, err error) {
	if !configcheckin.ConfigData.HonkaiStarRail.Enable {
		return "", nil
	}

	return DoCheckIn(configcheckin.ConfigData.HonkaiStarRail)
}
