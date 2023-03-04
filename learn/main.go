package main

import (
	"fmt"

	"cryptit.com/decrypt"
	"cryptit.com/encrypt"
)

func main() {
	// uuid := uuid.NewRandom()
	// fmt.Println(uuid)

	fmt.Println(decrypt.Nimbus(encrypt.Nimbus("Prajwal")))
}
