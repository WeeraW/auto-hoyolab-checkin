package myconsole

import "github.com/gonutz/w32/v2"

var CurrentConsole w32.HWND

func HideConsole() {
	console := w32.GetConsoleWindow()
	if console == 0 {
		return // no console attached
	}
	_, consoleProcID := w32.GetWindowThreadProcessId(console)
	if w32.GetCurrentProcessId() == consoleProcID {
		w32.ShowWindowAsync(console, w32.SW_HIDE)
	}
}

func ShowConsole() {
	console := w32.GetConsoleWindow()
	if console == 0 {
		return // no console attached
	}
	_, consoleProcID := w32.GetWindowThreadProcessId(console)
	if w32.GetCurrentProcessId() == consoleProcID {
		w32.ShowWindowAsync(console, w32.SW_SHOW)
	}
}
