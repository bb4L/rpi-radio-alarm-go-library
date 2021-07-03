package storage

import (
	"io/ioutil"
	"os"
	"sync"

	"github.com/bb4L/rpi-radio-alarm-go-library/constants"
	"github.com/bb4L/rpi-radio-alarm-go-library/logging"
	"github.com/bb4L/rpi-radio-alarm-go-library/types"

	"gopkg.in/yaml.v2"
)

const dataFilename = "./rpi_data.yaml"

var logger = logging.GetLogger(os.Stdout, constants.DefaultPrefix, "storagehelper")

var once sync.Once

// RpiRadioAlarmData contains all the data for the RpiRadioAlarm
type RpiRadioAlarmData struct {
	Settings Settings      `yaml:"settings"`
	Alarms   []types.Alarm `yaml:"alarms"`
	Radio    types.Radio   `yaml:"radio"`
}

// Settings contains specific settings for the program
type Settings struct {
	Port           int  `yaml:"port"`
	RunAPI         bool `yaml:"run_api"`
	RunTelegrambot bool `yaml:"run_telegram_bot"`
	RunDiscordbot  bool `yaml:"run_discord_bot"`
}

// Helper struct to sync access to the storage
type Helper struct {
	Mutex *sync.RWMutex
}

func (storageHelper *Helper) getMutex() *sync.RWMutex {
	once.Do(func() {
		storageHelper.Mutex = &sync.RWMutex{}
	})

	return storageHelper.Mutex
}

func (storageHelper *Helper) getStoredData() (RpiRadioAlarmData, error) {
	fileData, err := ioutil.ReadFile(dataFilename)

	if err != nil {
		panic(err)
	}

	var data RpiRadioAlarmData

	source := []byte(fileData)
	err = yaml.Unmarshal(source, &data)
	if err != nil {
		logger.Fatalf("error: %v", err)
	}
	return data, err
}

// SaveStoredData save the data to storage and releases the lock
func (storageHelper *Helper) SaveStoredData(data RpiRadioAlarmData) {
	outSource, err := yaml.Marshal(data)
	if err != nil {
		logger.Fatalf("error: %v", err)
	}

	ioutil.WriteFile(dataFilename, outSource, 0777)
	storageHelper.getMutex().Unlock()
}
