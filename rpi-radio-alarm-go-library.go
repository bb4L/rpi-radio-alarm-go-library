// Package rpiradioalarmgolibrary contains main helper functions
package rpiradioalarmgolibrary

import (
	"os"

	"github.com/bb4L/rpi-radio-alarm-go-library/api"
	"github.com/bb4L/rpi-radio-alarm-go-library/logging"
	"github.com/joho/godotenv"
)

var logger = logging.GetLogger("main", os.Stdout)

// GetHelperFromEnv Returns the working helper object with the config loaded from a .env file
func GetHelperFromEnv() api.Helper {

	err := godotenv.Load()
	if err != nil {
		logger.Fatal("error loading .env file")
	}

	if os.Getenv("ALARMURL") == "" {
		logger.Fatal("you have to specify ALARMURL in the .env")
	}

	helper := api.Helper{AlarmURL: os.Getenv("ALARMURL"), ExtraHeader: os.Getenv("EXTRAHEADER"), ExtreaHeaderValue: os.Getenv("EXTRAHEADERVALUE")}

	err = helper.CheckHealth()
	if err != nil {
		logger.Fatalf("health check failed with: %s", err)
	}

	return helper
}
