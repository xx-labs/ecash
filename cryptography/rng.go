package cryptography

import (
	"crypto/rand"
	"math/big"
)

func GenerateRandomValue(size int) ([]byte, error){
	value := make([]byte, size/8)

	_, err := rand.Read(value)

	if err != nil {
		return nil, err
	}

	return value, nil
}

func GenerateCoin (N *big.Int) ([]byte, []byte, error){

	preimage, _ := GenerateRandomValue(256)

	coin := HashToRSAModulus(h, preimage, N)

	return coin, preimage, nil
}
