package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello Gitpod!")
	print("test generic")
}

func print[T any](arg T) {
	fmt.Println(arg)
}