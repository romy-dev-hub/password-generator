package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func generatePassword(length int) string {
	lower := "abcdefghijklmnopqrstuvwxyz"
	upper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers := "0123456789"
	symbols := "!@#$%^&*()-_=+[]{}<>?"

	all := lower + upper + numbers + symbols
	password := ""

	for i := 0; i < length; i++ {
		index, _ := rand.Int(rand.Reader, big.NewInt(int64(len(all))))
		password += string(all[index.Int64()])
	}

	return password
}

func main() {
	var length int

	fmt.Print("enter password length: ")
	fmt.Scanln(&length)

	if length <= 0 {
		fmt.Println("length must be greater than 0 !")
		return
	}

	password := generatePassword(length)
	fmt.Println("here's your password: ")
	fmt.Println(password)
}
