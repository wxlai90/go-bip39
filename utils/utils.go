package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"

	"github.com/wxlai90/go-bip39/consts"
)

func GenerateSeedEntropy(size int) ([]byte, error) {
	entropy := make([]byte, size/8)
	_, err := rand.Read(entropy)

	if err != nil {
		return nil, err
	}

	return entropy, nil
}

func CalculateChecksum(entropy []byte) []byte {
	h := sha256.New()
	h.Write(entropy)

	return h.Sum(nil)[:len(entropy)/8/2]
}

func BytesToBinaryString(xb []byte) string {
	bits := ""

	for _, v := range xb {
		bits += fmt.Sprintf("%08b", v)
	}

	return bits
}

func ParseIntoMnemonic(bits string) string {
	words := make([]string, len(bits)/11)

	for i := 0; i < len(bits)-11; i += 11 {
		b := bits[i : i+11]

		if idx, err := strconv.ParseInt(b, 2, 64); err == nil {
			words[i/11] = consts.Wordlist[idx]
		}
	}

	return strings.Join(words, " ")
}
