package aamarpay

//Interface of Aamarpay Payment Gateway
type IPayment interface {
	GeneratePostUrl(params map[string]string, sandbox bool) (string, error)
}

//use this variable to access the implementation of this interface
//It is also for singleton pattern
var Payment IPayment

