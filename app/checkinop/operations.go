package checkinop

import (
	"fmt"
	"runtime/debug"
	"time"

	"fyne.io/systray"
	"github.com/WeeraW/auto-hoyolab-checkin/app/configcheckin"
	"github.com/WeeraW/auto-hoyolab-checkin/app/cookiereader"
	"github.com/WeeraW/auto-hoyolab-checkin/app/myconsants"
	"github.com/WeeraW/auto-hoyolab-checkin/app/servicelogger"
	"github.com/gen2brain/beeep"
)

type Message struct {
	Message string
	IsError bool
}

func RunProgram() {
	defer onPanic()
	var err error
	messages := []Message{}

	if len(cookiereader.HoyolabCookies) == 0 {
		beeep.Notify(myconsants.AppName, "No cookie found!", "")
		return
	}
	if err != nil {
		beeep.Notify(myconsants.AppName, fmt.Sprintf("Error! %s", err.Error()), "")
		return
	}
	for _, cookie := range cookiereader.HoyolabCookies {
		message, err := CheckinHonkaiImpact3(cookie)
		if err != nil {
			messages = append(messages, Message{Message: fmt.Sprintf("Error! Honkai Impact 3: %s", err.Error()), IsError: true})
		}
		if message != "" {
			messages = append(messages, Message{Message: fmt.Sprintf("Honkai Impact 3: %s", message), IsError: false})
		}
		time.Sleep(RandomSleepTime(1, 5))
		message, err = CheckinHonkaiStarRail(cookie)
		if err != nil {
			messages = append(messages, Message{Message: fmt.Sprintf("Error! Honkai Star Rail: %s", err.Error()), IsError: true})
		}
		if message != "" {
			messages = append(messages, Message{Message: fmt.Sprintf("Honkai Star Rail: %s", message), IsError: false})
		}
		time.Sleep(RandomSleepTime(1, 5))
		message, err = CheckinGenshinImpact(cookie)
		if err != nil {
			messages = append(messages, Message{Message: fmt.Sprintf("Error! Genshin Impact: %s", err.Error()), IsError: true})
		}
		if message != "" {
			messages = append(messages, Message{Message: fmt.Sprintf("Genshin Impact: %s", message), IsError: false})
		}
		time.Sleep(RandomSleepTime(1, 5))
		message, err = CheckinZenLessZoneZero(cookie)
		if err != nil {
			messages = append(messages, Message{Message: fmt.Sprintf("Error! ZenLess Zone Zero: %s", err.Error()), IsError: true})
		}
		if message != "" {
			messages = append(messages, Message{Message: fmt.Sprintf("ZenLess Zone Zero: %s", message), IsError: false})
		}
	}
	displayMessage(messages)
	systray.SetTooltip(fmt.Sprintf("%s done at %s", myconsants.AppName, time.Now().Format("15:04:05")))
}

func onPanic() {
	if r := recover(); r != nil {
		stackTrace := fmt.Sprintf("%s", r)
		beeep.Notify(myconsants.AppName, fmt.Sprintf("Panic! %s", stackTrace), "")
		servicelogger.Fatal(string(debug.Stack()))
	}
}

func displayMessage(message []Message) {
	switch configcheckin.ConfigData.MessageMode {
	case configcheckin.VerboseMode:
		for _, msg := range message {
			beeep.Notify(myconsants.AppName, msg.Message, "")
		}
	case configcheckin.SummaryMode:
		summaryMessage := ""
		for _, msg := range message {
			if msg.IsError {
				beeep.Alert(myconsants.AppName, msg.Message, "")
			} else {
				summaryMessage += msg.Message + "\n"
			}
		}
		if summaryMessage != "" {
			beeep.Notify(myconsants.AppName, summaryMessage, "")
		}
	case configcheckin.SilentMode:
		for _, msg := range message {
			if msg.IsError {
				beeep.Alert(myconsants.AppName, msg.Message, "")
			}
		}
	}
}
