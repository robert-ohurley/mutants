package main

type ErrType int

const (
	returnErr ErrType = iota + 1
	logicErr
	paramErr
	noErr
)

type err struct {
	//Defines a return error, logic error or param ordering error.
	errType ErrType

	//The value expected considering the errType
	expected string

	//What was recieved
	recieved string

	//after performing substitutions defined by the popes, a codeBody is set.
	//we can later walk the tree and execute all leaf node code
	codeBody string
}

func NewErr(errType ErrType, expected, recieved string) *err {
	return &err{errType, expected, recieved, ""}
}

func (e *err) SetCodyBody(s string) {
	e.codeBody = s
}
