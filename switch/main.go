package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	numero := 3

	fmt.Println("o numero, ", numero, " se escreve assim: ")

	switch numero {

	case 1:
		fmt.Println("um")

	case 2:
		fmt.Println("dois")
	case 3:
		fmt.Println("tres")

	}

	fmt.Println("testtando o SO")

	switch runtime.GOOS {

	case "linux":
		fallthrough
	case "mac":
		fallthrough
	case "windows":
		fmt.Println("sim", runtime.GOOS)
	default:
		fmt.Println("sei la")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("vai dormir")
	default:
		fmt.Println("trabalhar")
	}
}
