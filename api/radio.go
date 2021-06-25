package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/bb4L/rpi-radio-alarm-go-library/logging"
	"github.com/bb4L/rpi-radio-alarm-go-library/types"
)

// Get the radio status
func (helper *Helper) GetRadio() (types.Radio, error) {
	url := helper.AlarmURL + "/radio"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return types.Radio{}, fmt.Errorf("could not create request for GetRadio")
	}

	res, err := helper.prepareAndDoRequest(req)
	logging.GetInfoLogger().Println(res)

	if err != nil {
		logging.GetErrorLogger().Println(err)
		return types.Radio{}, err
	}

	jsonData, _ := ioutil.ReadAll(res.Body)

	if res.StatusCode != 200 {
		return types.Radio{}, fmt.Errorf(string(jsonData[:]))
	}

	var data types.Radio
	err = json.Unmarshal(jsonData, &data)

	if err != nil {
		logging.GetErrorLogger().Println(err)
		return types.Radio{}, err
	}

	return data, nil
}

// Start the radio
func (helper *Helper) StartRadio() (types.Radio, error) {
	values := map[string]string{"switch": "on"}
	jsonData, _ := json.Marshal(values)

	url := helper.AlarmURL + "/radio"
	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonData))
	if err != nil {
		return types.Radio{}, fmt.Errorf("could not create request for StartRadio")
	}

	res, err := helper.prepareAndDoRequest(req)
	logging.GetInfoLogger().Println(res)

	if err != nil {
		logging.GetErrorLogger().Println(err)
		return types.Radio{}, err
	}

	jsonData, _ = ioutil.ReadAll(res.Body)

	if res.StatusCode != 200 {
		return types.Radio{}, fmt.Errorf(string(jsonData[:]))
	}

	var data types.Radio
	err = json.Unmarshal(jsonData, &data)

	if err != nil {
		logging.GetErrorLogger().Println(err)
		return types.Radio{}, err
	}

	return data, nil
}

// Stop the radio
func (helper *Helper) StopRadio() (types.Radio, error) {
	values := map[string]string{"switch": "off"}
	jsonData, _ := json.Marshal(values)

	url := helper.AlarmURL + "/radio"
	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonData))
	if err != nil {
		return types.Radio{}, fmt.Errorf("could not create request for StopRadio")
	}

	res, err := helper.prepareAndDoRequest(req)
	logging.GetInfoLogger().Println(res)

	if err != nil {
		logging.GetErrorLogger().Println(err)
		return types.Radio{}, err
	}

	jsonData, _ = ioutil.ReadAll(res.Body)

	if res.StatusCode != 200 {
		return types.Radio{}, fmt.Errorf(string(jsonData[:]))
	}

	var data types.Radio
	err = json.Unmarshal(jsonData, &data)

	if err != nil {
		logging.GetErrorLogger().Println(err)
		return types.Radio{}, err
	}

	return data, nil
}
