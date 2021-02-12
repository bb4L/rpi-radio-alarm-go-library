package rpiradioalarmgolibrary

import (
	"os"

	"github.com/bb4L/rpi-radio-alarm-go-library/api"
	"github.com/bb4L/rpi-radio-alarm-go-library/constants"
	"github.com/joho/godotenv"
)

// GetHelperFromEnv returns the working helper object with the config loaded from a .env file
func GetHelperFromEnv() api.Helper {

	if constants.RpiLibraryLogger == nil {
		constants.RpiLibraryLogger = constants.GetLogger()
	}
	logger := constants.RpiLibraryLogger

	err := godotenv.Load()
	if err != nil {
		logger.Fatal("Error loading .env file")
	}

	if os.Getenv("ALARMURL") == "" {
		logger.Fatal("You have to specify ALARMURL in the .env")
	}

	helper := api.Helper{AlarmURL: os.Getenv("ALARMURL"), ExtraHeader: os.Getenv("EXTRAHEADER"), ExtreaHeaderValue: os.Getenv("EXTRAHEADERVALUE"), Logger: logger}

	err = helper.CheckHealth()
	if err != nil {
		logger.Fatalf("Health check failed with: %s", err)
	}

	return helper
}
