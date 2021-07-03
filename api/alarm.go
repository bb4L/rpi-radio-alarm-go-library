// Package api contains the data / helper functions to interact with the api
package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/bb4L/rpi-radio-alarm-go-library/constants"
	"github.com/bb4L/rpi-radio-alarm-go-library/logging"
	"github.com/bb4L/rpi-radio-alarm-go-library/types"
)

var logger = logging.GetLogger(os.Stdout, constants.DefaultPrefix, "alarm")

// GetAlarms gets all alarms
func (helper *Helper) GetAlarms(withWritePermission bool) ([]types.Alarm, error) {
	url := helper.AlarmURL + "/alarm"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create request for GetAlarms")
	}

	res, err := helper.prepareAndDoRequest(req)
	if err != nil {
		logger.Println(err)
		return nil, err
	}

	jsonData, _ := ioutil.ReadAll(res.Body)

	if res.StatusCode != 200 {
		return nil, fmt.Errorf(string(jsonData[:]))
	}

	var data []types.Alarm
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		logger.Println(err)
		return nil, err
	}

	return data, nil
}

// GetAlarm gets a specific alarm by index
func (helper *Helper) GetAlarm(idx int, withWritePermission bool) (types.Alarm, error) {
	url := helper.AlarmURL + "/alarm/" + strconv.Itoa(idx)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return types.Alarm{}, fmt.Errorf("could not create request for GetAlarm")
	}

	res, err := helper.prepareAndDoRequest(req)
	if err != nil {
		logger.Println(err)
		return types.Alarm{}, err
	}

	jsonData, _ := ioutil.ReadAll(res.Body)

	if res.StatusCode != 200 {
		return types.Alarm{}, fmt.Errorf(string(jsonData[:]))
	}

	var data types.Alarm
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		logger.Println(err)
		return types.Alarm{}, err
	}

	return data, nil
}

// SaveAlarm saves the given alarm on the given index
func (helper *Helper) SaveAlarm(idx int, alarm types.Alarm) (types.Alarm, error) {
	return helper.ChangeAlarm(alarm, idx)
}

// ChangeAlarm changes the alarm on the given index with the data of the passed instance
func (helper *Helper) ChangeAlarm(alarm types.Alarm, idx int) (types.Alarm, error) {
	url := helper.AlarmURL + "/alarm/" + strconv.Itoa(idx)

	byteAlarm, marshalError := json.Marshal(alarm)
	if marshalError != nil {
		return types.Alarm{}, fmt.Errorf("could not marshal alarm")
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(byteAlarm))
	if err != nil {
		return types.Alarm{}, fmt.Errorf("could not create request for ChangeAlarm")
	}

	res, err := helper.prepareAndDoRequest(req)
	if err != nil {
		logger.Println(err)
		return types.Alarm{}, err
	}

	jsonData, _ := ioutil.ReadAll(res.Body)

	if res.StatusCode != 200 {
		return types.Alarm{}, fmt.Errorf(string(jsonData[:]))
	}

	var data types.Alarm
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		logger.Println(err)
		return types.Alarm{}, fmt.Errorf("could not unmarshal result")
	}

	return data, nil
}

// AddAlarm adds the given alarm
func (helper *Helper) AddAlarm(alarm types.Alarm) ([]types.Alarm, error) {
	url := helper.AlarmURL + "/alarm"

	jsonData, err := json.Marshal(alarm)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonData))
	if err != nil {
		return nil, fmt.Errorf("could not create request for AddAlarm")
	}

	res, err := helper.prepareAndDoRequest(req)
	if err != nil {
		logger.Println(err)
		return nil, err
	}

	jsonData, _ = ioutil.ReadAll(res.Body)

	if res.StatusCode != 200 {
		return nil, fmt.Errorf(string(jsonData[:]))
	}

	var data []types.Alarm
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		logger.Println(err)
		return nil, err
	}

	return data, nil
}

// DeleteAlarm deletes the alarm with the given index
func (helper *Helper) DeleteAlarm(idx int) ([]types.Alarm, error) {
	url := helper.AlarmURL + "/alarm/" + strconv.Itoa(idx)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create request for AddAlarm")
	}

	res, err := helper.prepareAndDoRequest(req)
	if err != nil {
		logger.Println(err)
		return nil, err
	}

	jsonData, _ := ioutil.ReadAll(res.Body)

	if res.StatusCode != 200 {
		return nil, fmt.Errorf(string(jsonData[:]))
	}

	var data []types.Alarm
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		logger.Println(err)
		return nil, err
	}

	return data, nil
}
