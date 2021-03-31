package main

import "fmt"

var version = "0.0.3"

func GetVersion() string {
	return version
}

func main() {
	fmt.Println(GetVersion())
}
