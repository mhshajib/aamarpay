package aamarpay

//Interface of Aamarpay Payment Gateway
type IPayment interface {
	Init(sandbox bool)
	SetData(params map[string]string)
	GeneratePostUrl() (string, error)
}

//use this variable to access the implementation of this interface
//It is also for singleton pattern
var Payment IPayment

