package main

import (
	"fmt"
	"log"

	"github.com/wxlai90/go-bip39/utils"
)

func main() {
	entropy, err := utils.GenerateSeedEntropy()
	if err != nil {
		log.Panicln(err)
	}

	checksum := utils.CalculateChecksum(entropy)
	entropy = append(entropy, checksum...)
	bits := utils.BytesToBinaryString(entropy)

	mnemonic := utils.ParseIntoMnemonic(bits)

	fmt.Println(mnemonic)
}
