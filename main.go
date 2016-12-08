package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/TIBCOSoftware/flogo-lib/engine"
	"github.com/op/go-logging"
)

func init() {
	//	var format = logging.MustStringFormatter(
	//		"%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.5s} %{color:reset} %{message}",
	//	)
	//
	//	backend := logging.NewLogBackend(os.Stderr, "", 0)
	//	backendFormatter := logging.NewBackendFormatter(backend, format)
	//	logging.SetBackend(backendFormatter)
	//	logging.SetLevel(logging.INFO, "")
}

//var log = logging.MustGetLogger("main")

func main() {

	// Default engine config used always
	// Config should be handled by ENV VARIABLES
	// We need a way to contribute services, why not same contribution model than activity and triggers and read from app.json
	//engineConfig := GetEngineConfig()
	// Default engine used now, this should be read from app.json
	// triggersConfig := GetTriggersConfig()

	// If logrus is use user can modify it if needed, we could add some hook for non generated code
	//logLevel, _ := logging.LogLevel(engineConfig.LogLevel)
	//logging.SetLevel(logLevel, "")

	engine := engine.New()

	EnableFlowServices(engine, engineConfig)

	engine.Start()

	exitChan := setupSignalHandling()

	code := <-exitChan

	engine.Stop()

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
