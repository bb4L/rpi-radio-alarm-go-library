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
func (helper *Helper) GetRadio() (types.Radio, error) {
	url := helper.AlarmURL + "/radio"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return types.Radio{}, fmt.Errorf("Could not create request for GetRadio")
	}

	res, err := helper.prepareAndDoRequest(req)
	helper.Logger.Println(res)

	if err != nil {
		helper.Logger.Println(err)
		return types.Radio{}, err
	}

	jsonData, _ := ioutil.ReadAll(res.Body)

	if res.StatusCode != 200 {
		return types.Radio{}, fmt.Errorf(string(jsonData[:]))
	}

	var data types.Radio
	err = json.Unmarshal(jsonData, &data)

	if err != nil {
		helper.Logger.Println(err)
		return types.Radio{}, err
	}

	return data, nil
}

// StartRadio starts the radio
func (helper *Helper) StartRadio() (types.Radio, error) {
	// TODO: bug not working
	values := map[string]string{"switch": "on"}
	jsonData, err := json.Marshal(values)

	url := helper.AlarmURL + "/radio"
	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonData))
	if err != nil {
		return types.Radio{}, fmt.Errorf("Could not create request for StartRadio")
	}

	res, err := helper.prepareAndDoRequest(req)
	helper.Logger.Println(res)

	if err != nil {
		helper.Logger.Println(err)
		return types.Radio{}, err
	}

	jsonData, _ = ioutil.ReadAll(res.Body)

	if res.StatusCode != 200 {
		return types.Radio{}, fmt.Errorf(string(jsonData[:]))
	}

	var data types.Radio
	err = json.Unmarshal(jsonData, &data)

	if err != nil {
		helper.Logger.Println(err)
		return types.Radio{}, err
	}

	return data, nil
}

// StopRadio stops the radio
func (helper *Helper) StopRadio() (types.Radio, error) {
	values := map[string]string{"switch": "off"}
	jsonData, err := json.Marshal(values)

	url := helper.AlarmURL + "/radio"
	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonData))
	if err != nil {
		return types.Radio{}, fmt.Errorf("Could not create request for StopRadio")
	}

	res, err := helper.prepareAndDoRequest(req)
	helper.Logger.Println(res)

	if err != nil {
		helper.Logger.Println(err)
		return types.Radio{}, err
	}

	jsonData, _ = ioutil.ReadAll(res.Body)

	if res.StatusCode != 200 {
		return types.Radio{}, fmt.Errorf(string(jsonData[:]))
	}

	var data types.Radio
	err = json.Unmarshal(jsonData, &data)

	if err != nil {
		helper.Logger.Println(err)
		return types.Radio{}, err
	}

	return data, nil
}
