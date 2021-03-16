package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/bb4L/rpi-radio-alarm-go-library/types"
)

// GetAlarms get all alarms
func (helper *Helper) GetAlarms() ([]types.Alarm, error) {
	url := helper.AlarmURL + "/alarm"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Could not create request for GetAlarms")
	}

	res, err := helper.prepareAndDoRequest(req)
	helper.Logger.Println(res)

	if err != nil {
		helper.Logger.Println(err)
		return nil, err
	}

	jsonData, _ := ioutil.ReadAll(res.Body)

	if res.StatusCode != 200 {
		return nil, fmt.Errorf(string(jsonData[:]))
	}

	var data []types.Alarm
	err = json.Unmarshal(jsonData, &data)

	if err != nil {
		helper.Logger.Println(err)
		return nil, err
	}

	return data, nil
}

// GetAlarm get a specific alarm
func (helper *Helper) GetAlarm(idx int) (types.Alarm, error) {
	url := helper.AlarmURL + "/alarm/" + strconv.Itoa(idx)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return types.Alarm{}, fmt.Errorf("Could not create request for GetAlarm")
	}

	res, err := helper.prepareAndDoRequest(req)
	helper.Logger.Println(res)

	if err != nil {
		helper.Logger.Println(err)
		return types.Alarm{}, err
	}

	jsonData, _ := ioutil.ReadAll(res.Body)

	if res.StatusCode != 200 {
		return types.Alarm{}, fmt.Errorf(string(jsonData[:]))
	}

	var data types.Alarm
	err = json.Unmarshal(jsonData, &data)

	if err != nil {
		helper.Logger.Println(err)
		return types.Alarm{}, err
	}

	return data, nil
}

// ChangeAlarm change the alarm on the fiven index with the data of the passed instance
func (helper *Helper) ChangeAlarm(alarm types.Alarm, idx int) (types.Alarm, error) {
	url := helper.AlarmURL + "/alarm/" + strconv.Itoa(idx)

	byteAlarm, marshalError := json.Marshal(alarm)
	if marshalError != nil {
		return types.Alarm{}, fmt.Errorf("Could not marshal alarm")
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(byteAlarm))
	if err != nil {
		return types.Alarm{}, fmt.Errorf("Could not create request for ChangeAlarm")
	}

	res, err := helper.prepareAndDoRequest(req)
	helper.Logger.Println(res)

	if err != nil {
		helper.Logger.Println(err)
		return types.Alarm{}, err
	}

	jsonData, _ := ioutil.ReadAll(res.Body)

	if res.StatusCode != 200 {
		return types.Alarm{}, fmt.Errorf(string(jsonData[:]))
	}

	var data types.Alarm
	err = json.Unmarshal(jsonData, &data)

	if err != nil {
		helper.Logger.Println(err)
		return types.Alarm{}, fmt.Errorf("Could not unmarshal result")
	}

	return data, nil
}

// AddAlarm Adds the given alarm
func (helper *Helper) AddAlarm(alarm types.Alarm) ([]types.Alarm, error) {
	url := helper.AlarmURL + "/alarm"

	jsonData, err := json.Marshal(alarm)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonData))
	if err != nil {
		return nil, fmt.Errorf("Could not create request for AddAlarm")
	}

	res, err := helper.prepareAndDoRequest(req)
	helper.Logger.Println(res)
	if err != nil {
		helper.Logger.Println(err)
		return nil, err
	}

	jsonData, _ = ioutil.ReadAll(res.Body)

	if res.StatusCode != 200 {
		return nil, fmt.Errorf(string(jsonData[:]))
	}

	var data []types.Alarm
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		helper.Logger.Println(err)
		return nil, err
	}

	return data, nil
}

// DeleteAlarm Delete the alarm with the given index
func (helper *Helper) DeleteAlarm(idx int) ([]types.Alarm, error) {
	url := helper.AlarmURL + "/alarm/" + strconv.Itoa(idx)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Could not create request for AddAlarm")
	}

	res, err := helper.prepareAndDoRequest(req)
	helper.Logger.Println(res)
	if err != nil {
		helper.Logger.Println(err)
		return nil, err
	}

	jsonData, _ := ioutil.ReadAll(res.Body)

	if res.StatusCode != 200 {
		return nil, fmt.Errorf(string(jsonData[:]))
	}

	var data []types.Alarm
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		helper.Logger.Println(err)
		return nil, err
	}

	return data, nil
}
