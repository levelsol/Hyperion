package main

import (
	"Hyperion/core/method"
	"Hyperion/core/method/methods"
)

func main() {
	registerMethod()
}

func registerMethod() {
	method.RegisterMethod(methods.Join{})
	method.RegisterMethod(methods.Ping{})
	method.RegisterMethod(methods.MOTD{})
}