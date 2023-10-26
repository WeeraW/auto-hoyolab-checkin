package mysystray

import (
	"fmt"
	"os"
	"time"

	"github.com/WeeraW/auto-hoyolab-checkin/app/checkinop"
	"github.com/WeeraW/auto-hoyolab-checkin/app/configcheckin"
	"github.com/WeeraW/auto-hoyolab-checkin/app/cookiereader"
	"github.com/WeeraW/auto-hoyolab-checkin/app/myconsants"
	"github.com/WeeraW/auto-hoyolab-checkin/app/myconsole"
	"github.com/WeeraW/auto-hoyolab-checkin/app/servicelogger"
	"github.com/WeeraW/auto-hoyolab-checkin/icon"
	"github.com/gen2brain/beeep"
	"github.com/getlantern/systray"
	"github.com/go-co-op/gocron"
	"github.com/kardianos/service"
)

var cronJob *gocron.Scheduler

func Init() {
	servicelogger.Info("Starting systray...")
	if service.Interactive() {
		servicelogger.Info("Running in terminal.")

		systray.Run(onReady, onExit)
	} else {
		servicelogger.Info("Running under service manager.")
		systray.Register(onReady, onExit)
	}
}

func onReady() {
	servicelogger.Info("Systray ready!")
	systray.SetTemplateIcon(icon.Data, icon.Data)
	systray.SetTitle(myconsants.AppName)
	systray.SetTooltip(myconsants.AppName)

	bShow := systray.AddMenuItem("Show window", "Show console")
	bHide := systray.AddMenuItem("Hide window", "Hide console")
	systray.AddSeparator()
	bRetry := systray.AddMenuItem("Retry now", "Retry checkin now")
	systray.AddSeparator()
	bLoadCookieBrowser := systray.AddMenuItem("Load cookie form browser", "Load cookie from browser")
	bLoadCookieFile := systray.AddMenuItem("Load cookie form file", "Load cookie from file")
	systray.AddSeparator()
	if myconsole.CurrentConsole == 0 {
		bHide.Disable()
		bShow.Disable()
	}

	bExit := systray.AddMenuItem("Exit", "Exit the whole app")
	go func() {
		for {
			select {
			case <-bHide.ClickedCh:
				myconsole.HideConsole()
			case <-bShow.ClickedCh:
				myconsole.ShowConsole()
			case <-bLoadCookieBrowser.ClickedCh:
				cookiereader.ReadCookieFromBrowser()
			case <-bLoadCookieFile.ClickedCh:
				cookiereader.ReadCookiesFromFile()
			case <-bRetry.ClickedCh:
				checkinop.RunProgram()
			case <-bExit.ClickedCh:
				systray.Quit()
			}
		}
	}()

	// init config
	err := configcheckin.ReadConfiguration()
	if err != nil {
		servicelogger.Error(err)
		beeep.Alert("Hoyolab Auto Checkin Error", fmt.Sprintf("%s.\n Program will exit", err.Error()), "")
		os.Exit(1)
		return
	}

	cronJob = gocron.NewScheduler(time.UTC)
	cronJob.Every(8).Hours().Do(checkinop.RunProgram)
	cronJob.StartAsync()
}

func onExit() {
	servicelogger.Info("Systray exiting...")
	os.Exit(0)
}
