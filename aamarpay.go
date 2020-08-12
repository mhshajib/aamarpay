package aamarpay

import (
	"fmt"
	"github.com/andelf/go-curl"
	"strings"
)

//	Real implementation of Aamarpay interface
type AamarpayConnection struct {}

func generateUrlString(params map[string]string) string {
	fieldsString := ""
	for key, element := range params {
		fieldsString += fmt.Sprintf("%s=%s&", key, element)
	}
	return strings.TrimRight(fieldsString, "&")
}

//Generating Post Url
func (a *AamarpayConnection) GeneratePostUrl(params map[string]string, sandbox bool) (string, error) {
	var result string
	requestUrl := "http://secure.aamarpay.com/request.php"
	if sandbox {
		requestUrl = "https://sandbox.aamarpay.com/request.php"
	}

	//Initializing Curl
	easy := curl.EasyInit()
	defer easy.Cleanup()

	//Generating url string from array map
	fieldsString := generateUrlString(params)

	//Setting Required Options
	easy.Setopt(curl.OPT_VERBOSE, true)
	easy.Setopt(curl.OPT_URL, requestUrl)
	easy.Setopt(curl.OPT_POST, len(params))
	easy.Setopt(curl.OPT_POSTFIELDS, fieldsString)
	easy.Setopt(curl.OPT_SSL_VERIFYPEER, false)

	// Callback function for curl post call
	makeCall := func (buf []byte, data interface{}) bool {
		respUrl := strings.ReplaceAll(string(buf), "\"", "")
		respUrl = strings.ReplaceAll(respUrl, "\\", "")
		result = requestUrl+respUrl
		return true
	}

	easy.Setopt(curl.OPT_WRITEFUNCTION, makeCall)

	//Executing Curl Request
	if err := easy.Perform(); err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return result, err
	}

	return result, nil
}

