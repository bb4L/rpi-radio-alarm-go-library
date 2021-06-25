// Package rpiradioalarmgolibrary contains main helper functions
package rpiradioalarmgolibrary

import (
	"os"

	"github.com/bb4L/rpi-radio-alarm-go-library/api"
	"github.com/bb4L/rpi-radio-alarm-go-library/logging"
	"github.com/joho/godotenv"
)

// Returns the working helper object with the config loaded from a .env file
func GetHelperFromEnv() api.Helper {

	err := godotenv.Load()
	if err != nil {
		logging.GetFatalLogger().Fatal("error loading .env file")
	}

	if os.Getenv("ALARMURL") == "" {
		logging.GetFatalLogger().Fatal("you have to specify ALARMURL in the .env")
	}

	helper := api.Helper{AlarmURL: os.Getenv("ALARMURL"), ExtraHeader: os.Getenv("EXTRAHEADER"), ExtreaHeaderValue: os.Getenv("EXTRAHEADERVALUE")}

	err = helper.CheckHealth()
	if err != nil {
		logging.GetFatalLogger().Fatalf("health check failed with: %s", err)
	}

	return helper
}
