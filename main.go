package main

import (
	"fmt"
	"log/syslog"
	"net/http"
	"os"

	log "github.com/Sirupsen/logrus"
	logrus_syslog "github.com/Sirupsen/logrus/hooks/syslog"
	"github.com/nicholasjackson/example/global"
	"github.com/nicholasjackson/example/handlers"
)

func main() {
	config := os.Args[1]
	rootfolder := os.Args[2]

	global.LoadConfig(config, rootfolder)

	setupHandlers()
}

func setupHandlers() {
	logger := log.New()
	hook, err := logrus_syslog.NewSyslogHook("udp", global.Config.SysLogIP, syslog.LOG_DEBUG, "")
	if err != nil {
		log.Fatal("Unable to create sylogger")
	}

	logger.Hooks.Add(hook)

	http.Handle("/", handlers.GetRouter(logger))

	fmt.Println("Listening for connections on port", 8001)
	http.ListenAndServe(fmt.Sprintf(":%v", 8001), nil)
}
