package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/bb4L/rpi-radio-alarm-go-library/types"
)

// Helper A struct to make the api accessible
type Helper struct {
	AlarmURL          string
	ExtraHeader       string
	ExtreaHeaderValue string
	Logger            *log.Logger
}

// CheckHealth checks wether the api is reachable and returns 200 on the health endpoint
func (helper *Helper) CheckHealth() error {
	url := helper.AlarmURL + "/health"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("Could not create request for health")
	}

	res, err := helper.prepareAndDoRequest(req)
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return fmt.Errorf("Could not get health, request hat status code %d", res.StatusCode)
	}

	return nil
}

func (helper *Helper) GetAlarms() ([]types.Alarm, error) {
	url := helper.AlarmURL + "/alarm"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Could not create request for getAlarms")
	}

	res, err := helper.prepareAndDoRequest(req)
	helper.Logger.Println(res)

	if err != nil {
		helper.Logger.Println(err)
	}
	// log.Println(res.body)

	jsonData, _ := ioutil.ReadAll(res.Body)

	var data []types.Alarm
	err = json.Unmarshal(jsonData, &data)

	if err != nil {
		helper.Logger.Println(err)
	}

	// TODO: implement
	return data, nil
}

func (helper *Helper) prepareAndDoRequest(req *http.Request) (*http.Response, error) {
	req = helper.addHeadersToRequest(req)
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (helper *Helper) addHeadersToRequest(req *http.Request) *http.Request {
	if helper.ExtraHeader != "" && helper.ExtreaHeaderValue != "" {
		req.Header.Set(helper.ExtraHeader, helper.ExtreaHeaderValue)
	}

	return req
}
