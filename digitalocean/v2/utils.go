package v2

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// URL endpoint for version 2 api
const (
	DIGITALOCEAN_API_URL_PREFIX = "https://api.digitalocean.com/v2"
)

type Configuration struct {
	API_TOKEN string
}

func GetConfig() (*Configuration, error) {
	var file *os.File
	var err error
	if file, err = os.Open("config.json"); err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err1 := decoder.Decode(&configuration)
	if err1 != nil {
		return nil, err1
	} else {
		return &configuration, nil
	}
}
func HandleError(resp *http.Response) string {
	log.Printf("Status Code : %v", resp.StatusCode)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	jsonByte := []byte(string(body))
	var resFmt map[string]string

	if err := json.Unmarshal(jsonByte, &resFmt); err != nil {
		log.Println(err)
		log.Println(string(body))
	}
	return resFmt["message"]
}

/*
Function returns Request Object, it can be used to process GET and POST request
*/
func GetJsonResponse(reqType string, doEndPoint string, queryString string) ([]byte, error) {
	log.Printf("Endpoint : %v : Request Type : %v\nQueryString : %v\n", doEndPoint, reqType, queryString)
	// Get HTTP Client
	var config *Configuration
	var configErr error
	if config, configErr = GetConfig(); configErr != nil {
		return nil, configErr
	}
	authToken := config.API_TOKEN
	httpClient := &http.Client{}
	// Start creating request object
	var req *http.Request
	var err error

	if reqType == "GET" {
		req, err = http.NewRequest(reqType, doEndPoint, strings.NewReader(queryString))
	} else {
		req, err = http.NewRequest(reqType, doEndPoint, strings.NewReader(queryString))
	}
	if err != nil {
		return nil, err
	}
	// Add header to request for Authentication
	req.Header.Add("Authorization", "Bearer "+authToken)
	req.Header.Add("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if resp.StatusCode == 200 || resp.StatusCode == 202 || resp.StatusCode == 204 || resp.StatusCode == 201 {
		defer resp.Body.Close()
		jsonBody, _ := ioutil.ReadAll(resp.Body)
		return jsonBody, nil

	} else {
		return nil, errors.New(HandleError(resp))
	}
}
