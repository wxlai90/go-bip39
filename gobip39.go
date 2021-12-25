package gobip39

import (
	"log"

	"github.com/wxlai90/go-bip39/utils"
)

func GenerateMnemonic(size int) string {
	return generate(size)
}

func generate(size int) string {
	entropy, err := utils.GenerateSeedEntropy(size)
	if err != nil {
		log.Panicln(err)
	}

	checksum := utils.CalculateChecksum(entropy)
	entropy = append(entropy, checksum...)
	bits := utils.BytesToBinaryString(entropy)

	mnemonic := utils.ParseIntoMnemonic(bits)

	return mnemonic
}
