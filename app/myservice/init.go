package myservice

import (
	"time"

	"fyne.io/systray"
	"github.com/WeeraW/auto-hoyolab-checkin/app/mysystray"
	"github.com/WeeraW/auto-hoyolab-checkin/app/servicelogger"
	"github.com/gonutz/w32/v2"
	"github.com/kardianos/service"
)

// Experimental service manager
type Program struct {
	exit chan struct{}
}

func (p *Program) Start(s service.Service) error {
	servicelogger.Info("Starting service...")
	if service.Interactive() {
		servicelogger.Info("Program Running in terminal.")
		if w32.GetConsoleWindow() != 0 {
			servicelogger.Infof("\nAutomatic Hoyolab Check-in (https://github.com/WeeraW/auto-hoyolab-checkin) \n\n[DO NOT CLOSE THIS WINDOW]\nTo minimize or hide this window, \nclick the icon in the SYSTEM TRAY then choose \"Hide window\" button")
			servicelogger.Infof("Press Ctrl+C to exit...\n\n\n")
		}
	} else {
		servicelogger.Info("\nProgram Running under service manager.")
	}
	p.exit = make(chan struct{})
	go p.run()
	return nil
}

func (p *Program) run() error {
	servicelogger.Info("Initializing systray...")
	mysystray.Init()
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-p.exit:
			systray.Quit()
			ticker.Stop()
			return nil
		case <-ticker.C:
			// do work
		}
	}
}

func (p *Program) Stop(s service.Service) error {
	servicelogger.Info("Stopping service...")
	close(p.exit)
	return nil
}
