//go:build darwin
// +build darwin

package myservice

import (
	"github.com/WeeraW/auto-hoyolab-checkin/app/mysystray"
	"github.com/WeeraW/auto-hoyolab-checkin/app/servicelogger"
)

func platformInit() {
	// init systray
	servicelogger.Info("Initializing systray...")
	mysystray.Init()
}
