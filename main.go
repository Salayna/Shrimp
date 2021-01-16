package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args[1]
	if arg == "create-project" {
		fmt.Println("Waiting for your implementation")
		return
	}
	fmt.Println(arg)
}
