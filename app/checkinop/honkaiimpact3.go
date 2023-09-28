package checkinop

import "github.com/brokiem/auto-hoyolab-checkin/app/configcheckin"

func CheckinHonkaiImpact3() (message string, err error) {
	if !configcheckin.ConfigData.HonkaiImpact3.Enable {
		return "", nil
	}

	return DoCheckIn(configcheckin.ConfigData.HonkaiImpact3)
}
