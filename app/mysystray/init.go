package mysystray

import (
	"fmt"
	"os"
	"runtime/debug"
	"time"

	"fyne.io/systray"
	"github.com/WeeraW/auto-hoyolab-checkin/app/checkinop"
	"github.com/WeeraW/auto-hoyolab-checkin/app/configcheckin"
	"github.com/WeeraW/auto-hoyolab-checkin/app/cookiereader"
	"github.com/WeeraW/auto-hoyolab-checkin/app/myconsants"
	"github.com/WeeraW/auto-hoyolab-checkin/app/myconsole"
	"github.com/WeeraW/auto-hoyolab-checkin/app/mynotify"
	"github.com/WeeraW/auto-hoyolab-checkin/app/servicelogger"
	"github.com/WeeraW/auto-hoyolab-checkin/icon"
	"github.com/gen2brain/beeep"
	"github.com/go-co-op/gocron"
)

var cronJob *gocron.Scheduler

func Init() {
	servicelogger.Info("Starting systray...")
	// if service.Interactive() {
	// 	servicelogger.Info("Running in terminal.")

	// 	systray.Run(onReady, onExit)
	// } else {
	// 	servicelogger.Info("Running under service manager.")
	// 	systray.Register(onReady, onExit)
	// }
	systray.Run(onReady, onExit)
}

func onReady() {
	onPanic()
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
	configMenu := systray.AddMenuItem("Configuration", "Edit configuration")
	messageConfigMenu := configMenu.AddSubMenuItem("Message mode", "Message mode")
	btnSetMessageToVerbose := messageConfigMenu.AddSubMenuItem("Verbose", "messages will be shown every action")
	btnSetMessageToSummary := messageConfigMenu.AddSubMenuItem("Summary", "messages will be shown only checking in result")
	btnSetMessageToSilent := messageConfigMenu.AddSubMenuItem("Silent", "messages will be shown when error is occurred")

	bExit := systray.AddMenuItem("Exit", "Exit the whole app")
	go func() {
		for {
			select {
			case <-bHide.ClickedCh:
				myconsole.HideConsole()
			case <-bShow.ClickedCh:
				myconsole.ShowConsole()
			case <-bLoadCookieBrowser.ClickedCh:
				err := cookiereader.ReadCookieFromBrowser()
				if err != nil {
					servicelogger.Error(err)
					mynotify.NotifyError(err.Error())
				}
			case <-bLoadCookieFile.ClickedCh:
				cookiereader.ReadCookiesFromFile()
			case <-bRetry.ClickedCh:
				checkinop.RunProgram()
			case <-btnSetMessageToVerbose.ClickedCh:
				configcheckin.SetMessageMode(configcheckin.VerboseMode)
			case <-btnSetMessageToSummary.ClickedCh:
				configcheckin.SetMessageMode(configcheckin.SummaryMode)
			case <-btnSetMessageToSilent.ClickedCh:
				configcheckin.SetMessageMode(configcheckin.SilentMode)
			case <-bExit.ClickedCh:
				systray.Quit()
			}
		}
	}()

	// init config
	err := configcheckin.ReadConfiguration()
	if err != nil {
		servicelogger.Error(err)
		mynotify.NotifyError(fmt.Sprintf("%s.\n Program will exit", err.Error()))
		os.Exit(1)
		return
	}
	err = cookiereader.ReadCookie()
	if err != nil {
		servicelogger.Error(err)
		mynotify.NotifyError(err.Error())
	}
	initCronJob()
}

func initCronJob() {
	servicelogger.LogFile.RotateLog()
	cronJob = gocron.NewScheduler(time.UTC)
	cronJob.Every(8).Hours().Do(checkinop.RunProgram)
	// rotate log file
	cronJob.Every(1).Day().At("00:00").Do(func() {
		servicelogger.LogFile.RotateLog()
	})
	cronJob.StartAsync()
}

func onExit() {
	msg := fmt.Sprintf("%s exiting...", myconsants.AppName)
	servicelogger.Info(msg)
	mynotify.Notify(msg)
	os.Exit(0)
}

func onPanic() {
	if r := recover(); r != nil {
		stackTrace := fmt.Sprintf("%s", r)
		beeep.Alert(myconsants.AppName, fmt.Sprintf("Panic! when initialize %s", stackTrace), "")
		servicelogger.Fatal(debug.Stack())
	}
}
