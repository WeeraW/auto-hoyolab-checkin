package main

import (
	"fmt"
	"os"
	"time"

	"github.com/brokiem/auto-hoyolab-checkin/app/checkinop"
	"github.com/brokiem/auto-hoyolab-checkin/app/configcheckin"
	"github.com/brokiem/auto-hoyolab-checkin/app/cookiereader"
	"github.com/brokiem/auto-hoyolab-checkin/app/myconsole"
	"github.com/brokiem/auto-hoyolab-checkin/icon"
	"github.com/getlantern/systray"
	"github.com/go-co-op/gocron"
	_ "github.com/zellyn/kooky/browser/chrome"
	_ "github.com/zellyn/kooky/browser/firefox"
	_ "github.com/zellyn/kooky/browser/opera"
	_ "github.com/zellyn/kooky/browser/safari"
)

var cronJob *gocron.Scheduler

func main() {

	fmt.Println(" \nAutomatic Hoyolab Check-in (https://github.com/brokiem/auto-hoyolab-checkin) \n\n[DO NOT CLOSE THIS WINDOW]\nTo minimize or hide this window, \nclick the icon in the SYSTEM TRAY then choose \"Hide window\" button\n ")
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetTemplateIcon(icon.Data, icon.Data)
	systray.SetTitle("Automatic Hoyolab Check-in")
	systray.SetTooltip("Automatic Hoyolab Check-in")

	bShow := systray.AddMenuItem("Show window", "Show console")
	bHide := systray.AddMenuItem("Hide window", "Hide console")
	systray.AddSeparator()
	bExit := systray.AddMenuItem("Exit", "Exit the whole app")
	go func() {
		for {
			select {
			case <-bHide.ClickedCh:
				myconsole.HideConsole()
			case <-bShow.ClickedCh:
				myconsole.ShowConsole()
			case <-bExit.ClickedCh:
				systray.Quit()
			}
		}
	}()

	// init config
	err := configcheckin.ReadConfiguration()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}

	err = cookiereader.ReadCookieFromBrowser()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}
	cronJob = gocron.NewScheduler(time.UTC)
	cronJob.Every(12).Hours().Do(checkinop.RunProgram)
	cronJob.StartAsync()
}

func onExit() {
	fmt.Println("Exiting...")
}
