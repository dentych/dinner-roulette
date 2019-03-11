package main

import (
	"fmt"
	"github.com/dentych/dinner-dash/security"
)

func main() {
	str, err := security.GenerateSession()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(str)

}