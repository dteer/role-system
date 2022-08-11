package settings

import (
	"log"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/rs/zerolog"
)

func initSentry() {
	var debug bool
	if Conf.RunMode == "DEBUG" {
		debug = true
	} else {
		debug = false
	}
	err := sentry.Init(sentry.ClientOptions{
		// Either set your DSN here or set the SENTRY_DSN environment variable.
		Dsn: Conf.Sentry.Dsn,
		// Either set environment and release here or set the SENTRY_ENVIRONMENT
		// and SENTRY_RELEASE environment variables.
		Environment: "",
		Release:     "",
		// Enable printing of SDK debug messages.
		// Useful when getting started or trying to figure something out.
		Debug:            debug,
		AttachStacktrace: true,
	})
	if err != nil {
		log.Fatalf("sentry.Init Failed: %s", err)
	}
	// Flush buffered events before the program terminates.
	// Set the timeout to the maximum duration the program can afford to wait.
	defer sentry.Flush(2 * time.Second)
}

type sendSentry struct {
}

func (l *sendSentry) Run(e *zerolog.Event, level zerolog.Level, message string) {
	levelConf := Conf.Sentry.Level
	if zerolog.Level(levelConf) <= level {
		sentry.CaptureMessage(message)
	}

}

func InitLog() zerolog.Logger {
	initSentry()
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	Logger = logger.Hook(&sendSentry{})
	return Logger
}
