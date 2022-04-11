package cryptography

import "math/big"


func GenerateBlindingFactor(N *big.Int) (*big.Int, *big.Int) {

	modulusSize := N.BitLen()

	for {

		x, _ := GenerateRandomValue(modulusSize)
		r := big.NewInt(0).SetBytes(x)

		gcd := big.NewInt(0).GCD(nil, nil, r, N)

		// GCD of r and N should be one!  This is checks if big.Int(1) & GCD are equal
		if big.NewInt(1).Cmp(gcd) == 0{

			inverse := big.NewInt(0).ModInverse(r, N)

			if inverse != nil{
				return r, inverse
			}
		}
	}
}

// BlindCoin receives a coin "c", a random blinding factor "r", an RSA public key "e", and an RSA modulus "N"
// raises the blinding factor to the RSA pub key, and multiplies that result with the coin value
// returns a blinded coin
func BlindCoin(c *big.Int, r *big.Int, e *big.Int, N *big.Int)*big.Int{

	blinding := big.NewInt(0).Exp(r, e, N)

	bc := ModMul(c, blinding, N)

	return bc
}

func GetBlindingInverse(r, N *big.Int) *big.Int {
	return big.NewInt(0).ModInverse(r, N)
}

func RemoveBlinding(bc *big.Int, r *big.Int, N *big.Int) *big.Int{

	inverse := GetBlindingInverse(r, N)

	coin := ModMul(bc, inverse, N)

	return coin
}