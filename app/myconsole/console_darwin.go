//go:build darwin

package myconsole

import (
	"fmt"
	"os"

	"github.com/WeeraW/auto-hoyolab-checkin/app/servicelogger"
)

func platformInit() {
	// macOS-specific console initialization
	servicelogger.Info("Initializing console...")

	// Hide console by redirecting output to /dev/null
	devNull, err := os.Open(os.DevNull)
	if err != nil {
		servicelogger.Error(fmt.Sprintf("Failed to open /dev/null: %v", err))
		return
	}
	os.Stdout = devNull
	os.Stderr = devNull
}
