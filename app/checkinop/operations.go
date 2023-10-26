package checkinop

import (
	"fmt"
	"time"

	"github.com/WeeraW/auto-hoyolab-checkin/app/cookiereader"
	"github.com/WeeraW/auto-hoyolab-checkin/app/myconsants"
	"github.com/gen2brain/beeep"
	"github.com/getlantern/systray"
)

func RunProgram() {
	var err error
	messages := []string{}

	err = cookiereader.ReadCookie()
	if err != nil {
		beeep.Notify(myconsants.AppName, fmt.Sprintf("Error! %s", err.Error()), "")
		return
	}
	for _, cookie := range cookiereader.HoyolabCookies {
		message, err := CheckinHonkaiImpact3(cookie)
		if err != nil {
			messages = append(messages, fmt.Sprintf("Error! Honkai Impact 3: %s", err.Error()))
		}
		if message != "" {
			messages = append(messages, fmt.Sprintf("Honkai Impact 3: %s", message))
		}
		time.Sleep(RandomSleepTime(1, 5))
		message, err = CheckinHonkaiStarRail(cookie)
		if err != nil {
			messages = append(messages, fmt.Sprintf("Error! Honkai Star Rail: %s", err.Error()))
		}
		if message != "" {
			messages = append(messages, fmt.Sprintf("Honkai Star Rail: %s", message))
		}
		time.Sleep(RandomSleepTime(1, 5))
		message, err = CheckinGenshinImpact(cookie)
		if err != nil {
			messages = append(messages, fmt.Sprintf("Error! Genshin Impact: %s", err.Error()))
		}
		if message != "" {
			messages = append(messages, fmt.Sprintf("Genshin Impact: %s", message))
		}
		time.Sleep(RandomSleepTime(1, 5))
	}
	for _, message := range messages {
		beeep.Notify(myconsants.AppName, message, "")
	}
	systray.SetTooltip(fmt.Sprintf("%s done at %s", myconsants.AppName, time.Now().Format("15:04:05")))
}
