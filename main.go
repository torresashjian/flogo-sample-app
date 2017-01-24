package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/TIBCOSoftware/flogo-lib/engine"
	"github.com/TIBCOSoftware/flogo-lib/types"
	"github.com/op/go-logging"
)

func init() {
	var format = logging.MustStringFormatter(
		"%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.5s} %{color:reset} %{message}",
	)

	backend := logging.NewLogBackend(os.Stderr, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, format)
	logging.SetBackend(backendFormatter)
	logging.SetLevel(logging.INFO, "")
}

var log = logging.MustGetLogger("main")

func main() {
	flogo, err := os.Open("flogo.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	jsonParser := json.NewDecoder(flogo)
	app := &types.AppConfig{}
	jsonParser.Decode(&app)
	e, err := engine.New(app)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	e.Start()

	exitChan := setupSignalHandling()

	code := <-exitChan

	e.Stop()

	os.Exit(code)
}

func setupSignalHandling() chan int {

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	exitChan := make(chan int)
	go func() {
		for {
			s := <-signalChan
			switch s {
			// kill -SIGHUP
			case syscall.SIGHUP:
				exitChan <- 0
			// kill -SIGINT/Ctrl+c
			case syscall.SIGINT:
				exitChan <- 0
			// kill -SIGTERM
			case syscall.SIGTERM:
				exitChan <- 0
			// kill -SIGQUIT
			case syscall.SIGQUIT:
				exitChan <- 0
			default:
				log.Debug("Unknown signal.")
				exitChan <- 1
			}
		}
	}()

	return exitChan
}
