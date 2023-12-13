package main

import (
	stdLog "log"
	"os"

	"github.com/Isotere/dotenv"
	"github.com/Isotere/logs/logger"
	"github.com/pkg/errors"
)

const (
	successExitCode = 0
	failExitCode    = 1
)

func main() {
	err := dotenv.Load()
	if err != nil {
		stdLog.Fatal(errors.Wrap(err, "cannot load enviroment configuration"))
	}

	log, err := logger.New(dotenv.GetString("APP_LEVEL", logger.LogLevelDevel))
	if err != nil {
		stdLog.Fatal(errors.Wrap(err, "cannot create logger instance"))
	}
	defer log.Close()

	os.Exit(successExitCode)
}
