package cryptography

import (
	"crypto/rand"
	"crypto/rsa"
	"golang.org/x/crypto/sha3"
	"math/big"
)

var h = sha3.NewCShake256(nil, nil)


func Keygen(size int) (*big.Int, *big.Int, *big.Int, error){
	privateKey, err := rsa.GenerateKey(rand.Reader, size)      // here 2048 is the number of bits for RSA

	if err!= nil {
		return nil, nil, nil, err
	}

	return privateKey.D, big.NewInt(int64(privateKey.E)), privateKey.N, nil

}

func Sign(msg []byte, d *big.Int, N *big.Int) *big.Int{
	// TODO: Strengthen this Signing module

	m := big.NewInt(0).SetBytes(msg)

	if !InModulus(m, N){
		m = big.NewInt(0).SetBytes(HashToRSAModulus(h, msg, N))
	}

	sig := big.NewInt(0).Exp(m, d, N)

	return sig
}

func Verify(signature *big.Int, e *big.Int, N *big.Int, msg []byte) bool{

	m := big.NewInt(0).SetBytes(msg)

	ok := m.Cmp(big.NewInt(0).Exp(signature, e, N))

	if ok != 0 {
		return false
	}

	return true
}