package main

type ErrorType string

const (
	returnError ErrorType = "Return Error"
	logicError            = "Logic Error"
	paramError            = "Parameter Error"
	noErr                 = "No Error"
)

type err struct {
	//Defines a return error, logic error or param ordering error.
	errorType ErrorType

	//The value expected considering the errType
	expected string

	//What was recieved
	recieved string
}

func NewErr(errorType ErrorType, expected, recieved string) *err {
	return &err{errorType, expected, recieved}
}
