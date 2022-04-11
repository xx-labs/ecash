package cryptography

import (
	"golang.org/x/crypto/sha3"
	"math/big"
)


func HashToRSAModulus(h sha3.ShakeHash, data []byte, N *big.Int) []byte{

	// To avoid RSA verification failures, we set this output size to be size(N) * 0.75
	size := N.BitLen()*3/4

	// allocate memory to the digest object
	digest := make([]byte, size/8)

	// clean the hash object to ensure no leftover residues are plugged into the hash function
	h.Reset()

	// produce the XOF(Extendable Output Function) hash digest
	h.Write(data)
	_, err := h.Read(digest)

	if err != nil {
		return nil
	}
	return digest
}