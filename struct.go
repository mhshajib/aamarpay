package aamarpay

import "net/url"

type AamarpayConnection struct {
	StoreId string
	SignatureKey string
	Amount string
	PaymentType string
	Currency string
	TranId string
	CusName string
	CusEmail string
	CusAdd1 string
	CusAdd2 string
	CusCity string
	CusState string
	CusPostcode string
	CusCountry string
	CusPhone string
	CusFax string
	ShipName string
	ShipAdd1 string
	ShipAdd2 string
	ShipCity string
	ShipState string
	ShipPostcode string
	ShipCountry string
	Desc string
	SuccessUrl string
	FailUrl string
	CancelUrl string
	OptA string
	OptB string
	OptC string
	OptD string
	AamarpayRequestUrl string
	AamarpayPaymentUrl string
	Data url.Values
}