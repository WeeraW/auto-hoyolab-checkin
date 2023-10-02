package myconsole

import (
	"github.com/WeeraW/auto-hoyolab-checkin/app/servicelogger"
	"github.com/gonutz/w32/v2"
)

func HideConsole() {
	if CurrentConsole == 0 {
		servicelogger.Warning("No console attached.")
		return // no console attached
	}
	_, consoleProcID := w32.GetWindowThreadProcessId(CurrentConsole)
	if w32.GetCurrentProcessId() == consoleProcID {
		w32.ShowWindowAsync(CurrentConsole, w32.SW_HIDE)
	}
}

func ShowConsole() {
	if CurrentConsole == 0 {
		servicelogger.Warning("No console attached.")
		return // no console attached
	}
	_, consoleProcID := w32.GetWindowThreadProcessId(CurrentConsole)
	if w32.GetCurrentProcessId() == consoleProcID {
		w32.ShowWindowAsync(CurrentConsole, w32.SW_SHOW)
	}
}

func AttachConsole() {
	if CurrentConsole != 0 {
		servicelogger.Warning("Console already attached.")
		return // already attached
	}
	CurrentConsole = w32.GetConsoleWindow()
}
