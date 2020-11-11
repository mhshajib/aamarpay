package aamarpay

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

//Implementation of Aamarpay interface
func (a *AamarpayConnection) Init(sandbox bool) {
	a.AamarpayRequestUrl = "http://secure.aamarpay.com/request.php"
	a.AamarpayPaymentUrl = "http://secure.aamarpay.com"
	if sandbox {
		a.AamarpayRequestUrl = "https://sandbox.aamarpay.com/request.php"
		a.AamarpayPaymentUrl = "https://sandbox.aamarpay.com"
	}
}

//Setting Data For Aamarpay Request
func (a *AamarpayConnection) SetData(params map[string]string) {
	for key, element := range params {
		if key == "store_id" {
			a.StoreId = element
		}
		if key == "signature_key" {
			a.SignatureKey = element
		}
		if key == "amount" {
			a.Amount = element
		}
		if key == "payment_type" {
			a.PaymentType = element
		}
		if key == "currency" {
			a.Currency = element
		}
		if key == "tran_id" {
			a.TranId = element
		}
		if key == "cus_name" {
			a.CusName = element
		}
		if key == "cus_email" {
			a.CusEmail = element
		}
		if key == "cus_add1" {
			a.CusAdd1 = element
		}
		if key == "cus_add2" {
			a.CusAdd2 = element
		}
		if key == "cus_city" {
			a.CusCity = element
		}
		if key == "cus_state" {
			a.CusState = element
		}
		if key == "cus_postcode" {
			a.CusPostcode = element
		}
		if key == "cus_country" {
			a.CusCountry = element
		}
		if key == "cus_phone" {
			a.CusPhone = element
		}
		if key == "cus_fax" {
			a.CusFax = element
		}
		if key == "ship_name" {
			a.ShipName = element
		}
		if key == "ship_add1" {
			a.ShipAdd1 = element
		}
		if key == "ship_add2" {
			a.ShipAdd2 = element
		}
		if key == "ship_city" {
			a.ShipCity = element
		}
		if key == "ship_state" {
			a.ShipState = element
		}
		if key == "ship_postcode" {
			a.ShipPostcode = element
		}
		if key == "ship_country" {
			a.ShipCountry = element
		}
		if key == "desc" {
			a.Desc = element
		}
		if key == "success_url" {
			a.SuccessUrl = element
		}
		if key == "fail_url" {
			a.FailUrl = element
		}
		if key == "cancel_url" {
			a.CancelUrl = element
		}
		if key == "opt_a" {
			a.OptA = element
		}
		if key == "opt_b" {
			a.OptB = element
		}
		if key == "opt_c" {
			a.OptC = element
		}
		if key == "opt_d" {
			a.OptD = element
		}
	}
}

//Generating Post Url
func (a *AamarpayConnection) GeneratePostUrl() (string, error) {
	data := url.Values{
		"store_id": {a.StoreId},
		"signature_key": {a.SignatureKey},
		"amount": {a.Amount},
		"payment_type": {a.PaymentType},
		"currency": {a.Currency},
		"tran_id": {a.TranId},
		"cus_name": {a.CusName},
		"cus_email": {a.CusEmail},
		"cus_add1": {a.CusAdd1},
		"cus_add2": {a.CusAdd2},
		"cus_city": {a.CusCity},
		"cus_state": {a.CusState},
		"cus_postcode": {a.CusPostcode},
		"cus_country": {a.CusCountry},
		"cus_phone": {a.CusPhone},
		"cus_fax": {a.CusFax},
		"ship_name": {a.ShipName},
		"ship_add1": {a.ShipAdd1},
		"ship_add2": {a.ShipAdd2},
		"ship_city": {a.ShipCity},
		"ship_state": {a.ShipState},
		"ship_postcode": {a.ShipPostcode},
		"ship_country": {a.ShipCountry},
		"desc": {a.Desc},
		"success_url": {a.SuccessUrl},
		"fail_url": {a.FailUrl},
		"cancel_url": {a.CancelUrl},
		"opt_a": {a.OptA},
		"opt_b": {a.OptB},
		"opt_c": {a.OptC},
		"opt_d": {a.OptD},
	}

	var transResponse string = ""

	//Setting up http client
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodPost, a.AamarpayRequestUrl, strings.NewReader(data.Encode()))

	//Setting request headers
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	//Making request
	respTransId, errRequest := client.Do(req)
	if errRequest != nil {
		return "", errRequest
	}

	//Reading Request Body
	transIdResponseBody, errTransIdResponseBody := ioutil.ReadAll(respTransId.Body)
	if errTransIdResponseBody != nil {
		return "", errTransIdResponseBody
	}

	//Parsing response data
	if errParsingResponse := json.Unmarshal(transIdResponseBody, &transResponse); errParsingResponse != nil {
		return "", errParsingResponse
	}

	//Returning final response
	return fmt.Sprintf("%s%s", a.AamarpayPaymentUrl, transResponse), nil
}

