package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	_, err := bcrypt.GenerateFromPassword([]byte("john123"), 10)
	if err != nil {
		log.Fatal("Diededed")
	}

	hello := []byte{36, 50, 97, 36, 49, 48, 36, 81, 71, 81, 48, 56, 70, 118, 72, 110, 111, 89, 72, 90, 101, 113, 70, 115, 55, 53, 114, 77, 46, 48, 101, 85, 103, 85, 50, 120, 88, 53, 112, 66, 88, 102, 53, 104, 71, 97, 99, 121, 98, 115, 77, 81, 48, 98, 114, 53, 113, 116, 104, 113}

	err = bcrypt.CompareHashAndPassword(hello, []byte("john123"))
	if err != nil {
		fmt.Println("NOT SUCCESS")
		return
	}

	fmt.Println("SUCCESS")
}
