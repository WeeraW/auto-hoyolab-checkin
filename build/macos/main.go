//go:build darwin
// +build darwin

package main

import (
	"log"

	"github.com/WeeraW/auto-hoyolab-checkin/app/filelogger"
	"github.com/WeeraW/auto-hoyolab-checkin/app/myconsole"
	"github.com/WeeraW/auto-hoyolab-checkin/app/myservice"
	"github.com/WeeraW/auto-hoyolab-checkin/app/servicelogger"

	// chome is not supported for now
	// _ "github.com/browserutils/kooky/browser/chrome"
	_ "github.com/browserutils/kooky/browser/edge"
	_ "github.com/browserutils/kooky/browser/firefox"

	_ "github.com/browserutils/kooky/browser/opera"
	_ "github.com/browserutils/kooky/browser/safari"
)

func main() {
	var err error

	servicelogger.Init()
	servicelogger.LogFile, err = filelogger.NewFileLogger()
	if err != nil {
		log.Fatal(err)
	}

	myconsole.Init()
	myservice.Init()
}
