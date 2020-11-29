![aamarpay](aamarpay.png)

# GO Package For Aamarpay Payment Gateway

Use aamarpay payment gateway with go language simply using this package.  

### Installation

Install the package using
```go
$ go get github.com/mhshajib/aamarpay
```

### Usage

To use the package import it in your `*.go` code
```go
import "github.com/mhshajib/aamarpay"
```

### Example

***Change param values according to your store***

```go
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
		"store_id": 	 "<your_store_id>",
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
		"opt_a":         "Reshad",
		"opt_b": 		 "Akil",
		"opt_c":		 "Liza",
		"opt_d": 		 "Sohel",
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
```

You'll get redirect url for payment as response


### **License**
The **aamarpay** is an open-source software licensed under the [MIT License](LICENSE)