package main

import (
	"fmt"

	"cryptit.com/decrypt"
	"cryptit.com/encrypt"
)

func main() {
	fmt.Println(decrypt.Nimbus(encrypt.Nimbus("Prajwal")))
}
