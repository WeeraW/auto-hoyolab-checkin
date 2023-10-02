package myconsole

import (
	"fmt"

	"github.com/WeeraW/auto-hoyolab-checkin/app/servicelogger"
	"github.com/gonutz/w32/v2"
)

var CurrentConsole w32.HWND

func Init() {
	servicelogger.Info("Initializing console...")
	CurrentConsole = w32.GetConsoleWindow()
	if CurrentConsole == 0 {
		servicelogger.Info("No console attached.")
		return
	} else {
		servicelogger.Info(fmt.Sprintf("Console attached. at PID: %v", w32.GetCurrentProcessId()))
	}
}
