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

	tran_id := strconv.Itoa(rand.New(rand.NewSource(time.Now().UnixNano())).Int())
	params := map[string]string{
		"store_id": "<your_store_id>",
		"signature_key": "<your_signature_key>",
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
		"success_url":   "<your_success_url>",
		"fail_url":      "<your_fail_url>",
		"cancel_url":    "<your_cancel_url>",
		"opt_a":         "Reshad", "opt_b": "Akil",
		"opt_c": "Liza", "opt_d": "Sohel",
	}

	//Generating url for post request
	response, err := aamarpay.Payment.GeneratePostUrl(params, true)
	if err != nil {
		panic(err)
	}

	//Url Generated Successfully. Printing url.
	fmt.Println(response)
}

