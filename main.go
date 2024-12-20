package main

import (
	"flag"
	"log"

	"github.com/WeeraW/auto-hoyolab-checkin/app/filelogger"
	"github.com/WeeraW/auto-hoyolab-checkin/app/myconsants"
	"github.com/WeeraW/auto-hoyolab-checkin/app/myconsole"
	"github.com/WeeraW/auto-hoyolab-checkin/app/myservice"
	"github.com/WeeraW/auto-hoyolab-checkin/app/servicelogger"

	// chome is not supported for now
	// _ "github.com/browserutils/kooky/browser/chrome"
	_ "github.com/browserutils/kooky/browser/edge"
	_ "github.com/browserutils/kooky/browser/firefox"

	_ "github.com/browserutils/kooky/browser/opera"
	_ "github.com/browserutils/kooky/browser/safari"
	"github.com/kardianos/service"
)

func main() {
	serviceFlag := flag.String("service", "", "service operation (install, uninstall, update, start, stop, restart)")
	flag.Parse()

	options := make(service.KeyValue)
	options["Restart"] = "on-success"
	options["SuccessExitStatus"] = "1 2 8 SIGKILL"

	// Define service config.
	svcConfig := &service.Config{
		Name:        myconsants.AppName,
		DisplayName: myconsants.AppName,
		Description: "Automatic Hoyolab Check-in",
	}
	prg := &myservice.Program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}

	// Setup the logger.
	errs := make(chan error, 5)
	servicelogger.Logger, err = s.Logger(errs)
	if err != nil {
		log.Fatal(err)
	}
	servicelogger.LogFile, err = filelogger.NewFileLogger()
	if err != nil {
		log.Fatal(err)
	}

	myconsole.Init()

	// Handle service controls (optional).
	if len(*serviceFlag) != 0 {
		err := service.Control(s, *serviceFlag)
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	// Run the service.
	if serviceFlag == nil || *serviceFlag == "" {
		err = s.Run()
		if err != nil {
			log.Fatal(err)
		}
	}

	if *serviceFlag == "install" {
		err = s.Install()
		if err != nil {
			log.Fatal(err)
		}
	}

	if *serviceFlag == "uninstall" {
		err = s.Uninstall()
		if err != nil {
			log.Fatal(err)
		}
	}

	if *serviceFlag == "update" {
		err = s.Uninstall()
		if err != nil {
			log.Fatal(err)
		}
		err = s.Install()
		if err != nil {
			log.Fatal(err)
		}
	}

	if *serviceFlag == "start" {
		err = s.Start()
		if err != nil {
			log.Fatal(err)
		}
	}

	if *serviceFlag == "stop" {
		err = s.Stop()
		if err != nil {
			log.Fatal(err)
		}
	}

	if *serviceFlag == "restart" {
		err = s.Restart()
		if err != nil {
			log.Fatal(err)
		}
	}
}
