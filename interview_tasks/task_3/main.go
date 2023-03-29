package main

import "fmt"

/*
type X struct{}

func (x *X) error() {
	fmt.Println("here`s an error")
}

func main() {
	fmt.Println(handle())
}

func handle(value x) error {
	return x.error
}
*/

// fixed

type X struct {
	Message string
}

func (x *X) Error() string {
	return x.Message
}

func main() {
	fmt.Println(handle(&X{"here`s an error"}))
}

func handle(value *X) error {
	return value
}
