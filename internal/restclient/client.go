/*
* Copyright (c) Siemens 2021
* Licensed under the MIT license
* See LICENSE file in the top-level directory
*/

package restclient

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	v1 "onboardservice/api/siemens_iedge_dmapi_v1"

	"google.golang.org/protobuf/encoding/protojson"
)

// ErrResponse implements error response when EdgeCoreClient get an error
type ErrResponse struct {
	Errors []struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"errors"`
}

// EdgeCoreRuntimeRest  API interface
type EdgeCoreRuntimeRest interface {
	Activate(configuration *v1.DeviceConfiguration) (bool, error)
	Onboarded() (bool, bool)
}

// EdgeCoreClient implements EdgeCoreRuntimeRest interface.
type EdgeCoreClient struct {
	activationURL   string
	onboardCheckURL string
	client          *http.Client
}

// NewClient Creates a new EdgeCoreClient with given IEM base url.
func NewClient(baseurl string) *EdgeCoreClient {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	cli := &http.Client{Transport: tr}
	return &EdgeCoreClient{
		activationURL:   baseurl + "/v1/activate",
		onboardCheckURL: baseurl + "/v1",
		client:          cli,
	}
}

// Activate allows the device activation to EdgeCoreRuntime by delivering onboarding configuration.
func (hc *EdgeCoreClient) Activate(configuration *v1.DeviceConfiguration) (bool, error) {
	jsonContent := protojson.MarshalOptions{Multiline: true, UseProtoNames: true}.Format(configuration)
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	//Multipart message , "files" section as documented in EdgeCoreRuntime API
	part1, _ := writer.CreateFormFile("files", "deviceconfig.json")
	part1.Write([]byte(jsonContent + "\n"))

	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return false, err

	}

	req, err := http.NewRequest("POST", hc.activationURL, payload)

	if err != nil {
		fmt.Println(err)
		return false, err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	//Make the Request to API
	res, err := hc.client.Do(req)
	if err == nil {
		defer res.Body.Close()
	} else {
		log.Println(req.URL, "failed", err)
		return false, err
	}
	//Parse response.
	body, err := io.ReadAll(res.Body)
	fmt.Println(string(body))
	if res.StatusCode >= 200 && res.StatusCode < 400 {
		log.Println("onboard succeed")
		return true, nil
	}

	if err == nil {
		// Prepare a error message containing the response from /activate
		err = errors.New(string(body))
	}

	log.Println(req.URL, ":onboard failed:", err)

	return false, err
}

// GetHttpErrorMessage consumes http response body stream and creates proper error message
func GetHttpErrorMessage(response *http.Response) string {
	var errResponse ErrResponse
	var errMessage string

	body, _ := io.ReadAll(response.Body)
	if err := json.Unmarshal(body, &errResponse); err != nil {
		errMessage = "Unexpected error, error response can not be parsed"
	} else {
		errMessage = errResponse.Errors[0].Message
	}

	return fmt.Sprintf("HTTP Response: %d - %s", response.StatusCode, errMessage)
}

// Onboarded checks the onboarding status of the device and edge accessibility of edge iot core runtime
func (hc *EdgeCoreClient) Onboarded() (isOnboarded bool, isContainerAccessible bool) {
	response, err := hc.client.Get(hc.onboardCheckURL)
	if err != nil {
		log.Println("Failed to check onboard status: ", err)
		isOnboarded, isContainerAccessible = false, false
	} else {
		defer response.Body.Close()

		switch {
		case response.StatusCode >= 200 && response.StatusCode < 400:
			isOnboarded, isContainerAccessible = true, true
			log.Printf("HTTP Response: %d - Onboarded check successful", response.StatusCode)

		case response.StatusCode == 531:
			isOnboarded, isContainerAccessible = false, true
			log.Printf("HTTP Response: %d - Device not onboarded", response.StatusCode)

		default:
			isOnboarded, isContainerAccessible = false, false
			log.Println(GetHttpErrorMessage(response))
		}
	}

	return isOnboarded, isContainerAccessible
}
