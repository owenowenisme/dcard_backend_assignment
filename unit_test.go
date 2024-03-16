package main

import (
    "encoding/json"
    "fmt"
    "github.com/joho/godotenv"
    "io"
    "net/http"
    "net/url"
    "os"
    "strings"
    "testing"
	"bytes"
)

var hostUrl string
var endpoint string
func TestInit(t *testing.T) {
	go main()
	if err := godotenv.Load(); err != nil {
		fmt.Println("Init function called")
	}

	fmt.Println("Init function called")
	hostUrl = os.Getenv("HOST_URL")
    endpoint = os.Getenv("ENDPOINT")
	t.Logf("HOST_URL: %v", hostUrl)
    t.Logf("ENDPOINT: %v", endpoint)
}
func TestInvalidCreateAd(t *testing.T) {
	data, err := os.ReadFile("testData/createAdInvalidTestData.json")
	if err != nil {
		t.Errorf("Failed to read test.json: %v", err)
		return
	}
	var testData []Ad
	err = json.Unmarshal(data, &testData)
	if err != nil {
		t.Errorf("Failed to unmarshal test.json: %v", err)
		return
	}
	for i:= range testData {
		testData[i] = InitAds()
	}
	err = json.Unmarshal(data, &testData)
	if err != nil {
		t.Errorf("Failed to unmarshal test.json: %v", err)
		return
	}
	for _, test := range testData {
		// Use test here
		requestBody, _ := json.MarshalIndent(test, "", "    ")
		res, err := http.Post(hostUrl+endpoint, "application/json", io.Reader(strings.NewReader(string(requestBody))))
		if err != nil {
			t.Errorf("Failed to send request: %v %v", hostUrl+endpoint,err)
			continue
		}
		body, _ := io.ReadAll(res.Body)

		if res.StatusCode != http.StatusCreated{
			t.Logf("Failed to create ad: {Code:%v, Body:%v}", res.Status, string(body))
			continue
		}

		t.Log("Response status code:", res.StatusCode)
		t.Log("RequestBody :", test)
		fmt.Println("---------------------------------")
	}
}
func TestCreateAd(t *testing.T) {
	data, err := os.ReadFile("testData/createAdTestData.json")
	if err != nil {
		t.Errorf("Failed to read test.json: %v", err)
		return
	}
	var testData []Ad
	err = json.Unmarshal(data, &testData)
	if err != nil {
		t.Errorf("Failed to unmarshal test.json: %v", err)
		return
	}
	for _, test := range testData {
		// Use test here
		requestBody, _ := json.MarshalIndent(test, "", "    ")
		res, err := http.Post(hostUrl+endpoint, "application/json", io.Reader(strings.NewReader(string(requestBody))))
        if err != nil {
			t.Errorf("Failed to send request: %v %v", hostUrl+endpoint,err)
			continue
		}
        body, _ := io.ReadAll(res.Body)

		if res.StatusCode != http.StatusCreated {
			t.Logf("Failed to create ad: %v %v ", res.Status, string(body))

			continue
		}

		// t.Log("Response status code:", res.StatusCode)
		// t.Log("RequestBody :", test)
		// fmt.Println("---------------------------------")
	}
}
func TestRetrieveAd(t *testing.T) {
	endpoint := "/api/v1/ad"
	data, err := os.ReadFile("testData/retrieveAdTestData.json")

	if err != nil {
		t.Errorf("Failed to read test.json: %v", err)
		return
	}
	var testData [][]string
	err = json.Unmarshal(data, &testData)
	if err != nil {
		t.Errorf("Failed to unmarshal test.json: %v", err)
		return
	}
    paramNames := []string{"offset", "limit", "age", "gender","country","platform"}

	for _, test := range testData {
        base, _ := url.Parse(hostUrl + endpoint)
        params := url.Values{}
        for i, paramName := range paramNames {
            if test[i] != "" {
                params.Add(paramName, test[i])
            }
        }
        base.RawQuery = params.Encode()
        fmt.Printf("Request URL:%v\n", base.String())
		res, err := http.Get(base.String())
		if err != nil {
			t.Errorf("Failed to send request: %v", hostUrl+endpoint)
			continue
		}
		if res.StatusCode != http.StatusOK {
			t.Errorf("Failed to retrieve ad: %v", res.Status)
			t.Log("Response status code:", res.StatusCode)
			continue
		}
		body, _ := io.ReadAll(res.Body)
		t.Log("Response status code:", res.StatusCode)
		var prettyJSON bytes.Buffer
		err = json.Indent(&prettyJSON, body, "", "\t")
		if err != nil {
			t.Errorf("Failed to format JSON: %v", err)
			continue
		}
		t.Log("Response body:", prettyJSON.String())

	}
}
