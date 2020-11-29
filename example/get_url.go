package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/mhshajib/aamarpay"
)

func main() {
	//Setting Aamarpay Interface
	aamarpay.Payment = &aamarpay.AamarpayConnection{}
	//Initializing Payment
	aamarpay.Payment.Init(true) //true for sandbox and false for live.

	tran_id := strconv.Itoa(rand.New(rand.NewSource(time.Now().UnixNano())).Int())
	params := map[string]string{
		"store_id":      "aamarpay",
		"signature_key": "28c78bb1f45112f5d40b956fe104645a",
		"amount":        "200",
		"payment_type":  "VISA",
		"currency":      "BDT",
		"tran_id":       tran_id,
		"cus_name":      "Mr. ABC",
		"cus_email":     "cus_mail@gmail.com",
		"cus_add1":      "Dhaka",
		"cus_add2":      "Mohakhali DOHS",
		"cus_city":      "Dhaka",
		"cus_state":     "Dhaka",
		"cus_postcode":  "1206",
		"cus_country":   "Bangladesh",
		"cus_phone":     "01711111111",
		"cus_fax":       "Not¬Applicable",
		"ship_name":     "some name",
		"ship_add1":     "House B-121,Road 21",
		"ship_add2":     "Mohakhali",
		"ship_city":     "Dhaka",
		"ship_state":    "Dhaka",
		"ship_postcode": "1212",
		"ship_country":  "Bangladesh",
		"desc":          "some university",
		"success_url":   "http://mhsajib.com/testSuccess",
		"fail_url":      "http://mhsajib.com/testFail",
		"cancel_url":    "<http://mhsajib.com/testCancel",
		"opt_a":         "Reshad",
		"opt_b":         "Akil",
		"opt_c":         "Liza",
		"opt_d":         "Sohel",
	}
	//Setting Data For Payment
	aamarpay.Payment.SetData(params)

	//Generating url for payment request
	response, err := aamarpay.Payment.GeneratePaymentUrl()
	if err != nil {
		panic(err)
	}

	//Url Generated Successfully. Printing url.
	fmt.Println(response)
}
