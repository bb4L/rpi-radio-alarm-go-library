package api

import (
	"fmt"
	"net/http"
)

// Helper struct to define the api settings being used
type Helper struct {
	AlarmURL          string
	ExtraHeader       string
	ExtreaHeaderValue string
}

// CheckHealth checks wether the api is reachable and returns 200 on the health endpoint
func (helper *Helper) CheckHealth() error {
	url := helper.AlarmURL + "/health"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("could not create request for health")
	}

	res, err := helper.prepareAndDoRequest(req)
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return fmt.Errorf("could not get health, request has status code %d", res.StatusCode)
	}

	return nil
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
