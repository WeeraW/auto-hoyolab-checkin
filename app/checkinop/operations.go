package checkinop

import (
	"fmt"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/getlantern/systray"
)

func RunProgram() {
	messages := []string{}
	message, err := CheckinHonkaiImpact3()
	if err != nil {
		messages = append(messages, fmt.Sprintf("Error! Honkai Impact 3: %s", err.Error()))
	}
	if message != "" {
		messages = append(messages, fmt.Sprintf("Honkai Impact 3: %s", message))
	}
	time.Sleep(RandomSleepTime(5))
	message, err = CheckinHonkaiStarRail()
	if err != nil {
		messages = append(messages, fmt.Sprintf("Error! Honkai Star Rail: %s", err.Error()))
	}
	if message != "" {
		messages = append(messages, fmt.Sprintf("Honkai Star Rail: %s", message))
	}
	time.Sleep(RandomSleepTime(5))
	message, err = CheckinGenshinImpact()
	if err != nil {
		messages = append(messages, fmt.Sprintf("Error! Genshin Impact: %s", err.Error()))
	}
	if message != "" {
		messages = append(messages, fmt.Sprintf("Genshin Impact: %s", message))
	}
	time.Sleep(RandomSleepTime(5))
	for _, message := range messages {
		beeep.Notify("Hoyolab Check-in", message, "")
	}
	systray.SetTooltip(fmt.Sprintf("Automatic Hoyolab Check-in done at %s", time.Now().Format("15:04:05")))
}
