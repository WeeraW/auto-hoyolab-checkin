package mynotify

import (
	"github.com/WeeraW/auto-hoyolab-checkin/app/configcheckin"
	"github.com/WeeraW/auto-hoyolab-checkin/app/myconsants"
	"github.com/gen2brain/beeep"
)

func Notify(message string) {
	if configcheckin.ConfigData.MessageMode == configcheckin.VerboseMode {
		beeep.Notify(myconsants.AppName, message, "")
	}
}

func NotifyError(message string) {
	beeep.Notify(myconsants.AppName, message, "")
}
