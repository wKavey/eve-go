package esi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const esiBaseURL = "https://esi.tech.ccp.is/latest/"

// APIObjectInterface allows all objects to be grabbed upon demand
type APIObjectInterface interface {
	APIInfo() *APIInfo
}

// Get returns the raw json in a.APIInfo().rawData
func Get(a APIObjectInterface, rawData interface{}, args string) error {

	url := esiBaseURL + a.APIInfo().url + args

	httpClient := http.Client{
		Timeout: time.Second * 5, // Maximum of 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, rawData)
	if err != nil {
		return fmt.Errorf(err.Error(), body)
	}

	return err
}

// ObjectLoaded returns true if the object has been created, false if not
func ObjectLoaded(a APIObjectInterface) bool {
	return a.APIInfo().initialized
}

// APIInfo keeps track of API related information
// This should be placed in all *Data structs
type APIInfo struct {
	initialized  bool
	url          string
	lastModified time.Time
}

// MakeAPIInfo returns an apiInfo type with proper defaults
func MakeAPIInfo(url string) APIInfo {
	return APIInfo{initialized: false, url: url}
}
