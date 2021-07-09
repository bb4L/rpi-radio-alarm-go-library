package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/bb4L/rpi-radio-alarm-go-library/types"
)

// GetRadio returns the radio status
func (helper *Helper) GetRadio(withWritePermission bool) (types.Radio, error) {
	url := helper.AlarmURL + "/radio"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return types.Radio{}, fmt.Errorf("could not create request for GetRadio")
	}

	res, err := helper.prepareAndDoRequest(req)
	if err != nil {
		logger.Println(err)
		return types.Radio{}, err
	}

	jsonData, _ := ioutil.ReadAll(res.Body)

	if res.StatusCode != 200 {
		return types.Radio{}, fmt.Errorf(string(jsonData[:]))
	}

	var data types.Radio
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		logger.Println(err)
		return types.Radio{}, err
	}

	return data, nil
}

// SaveRadio saves the radio and returns it
func (helper *Helper) SaveRadio(radio types.Radio) (types.Radio, error) {
	if radio.Running {
		return helper.StartRadio()
	}
	return helper.StopRadio()
}

// SwitchRadio changes the radio to the state running passed as argument
func (helper *Helper) SwitchRadio(running bool) (types.Radio, error) {
	if running {
		return helper.StartRadio()
	}
	return helper.StopRadio()
}

// StartRadio starts the radio
func (helper *Helper) StartRadio() (types.Radio, error) {
	values := map[string]string{"switch": "on"}
	jsonData, _ := json.Marshal(values)

	url := helper.AlarmURL + "/radio"
	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonData))
	if err != nil {
		return types.Radio{}, fmt.Errorf("could not create request for StartRadio")
	}

	res, err := helper.prepareAndDoRequest(req)
	if err != nil {
		logger.Println(err)
		return types.Radio{}, err
	}

	jsonData, _ = ioutil.ReadAll(res.Body)

	if res.StatusCode != 200 {
		return types.Radio{}, fmt.Errorf(string(jsonData[:]))
	}

	var data types.Radio
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		logger.Println(err)
		return types.Radio{}, err
	}

	return data, nil
}

// StopRadio stops the radio
func (helper *Helper) StopRadio() (types.Radio, error) {
	values := map[string]string{"switch": "off"}
	jsonData, _ := json.Marshal(values)

	url := helper.AlarmURL + "/radio"
	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonData))
	if err != nil {
		return types.Radio{}, fmt.Errorf("could not create request for StopRadio")
	}

	res, err := helper.prepareAndDoRequest(req)
	logger.Println(res)

	if err != nil {
		logger.Println(err)
		return types.Radio{}, err
	}

	jsonData, _ = ioutil.ReadAll(res.Body)

	if res.StatusCode != 200 {
		return types.Radio{}, fmt.Errorf(string(jsonData[:]))
	}

	var data types.Radio
	err = json.Unmarshal(jsonData, &data)

	if err != nil {
		logger.Println(err)
		return types.Radio{}, err
	}

	return data, nil
}
