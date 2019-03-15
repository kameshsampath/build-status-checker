package helpers

import (
	"os"

	log "github.com/sirupsen/logrus"
)

//SetLogLevel sets the log level logrus Logger
func SetLogLevel(lvl string) {
	logL, err := log.ParseLevel(lvl)
	if err == nil {
		log.SetLevel(logL)
	}
}

//HomeDir gets the user home directory
func HomeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
