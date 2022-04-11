package cryptography

import "math/big"

// InModulus receives a value "x" and checks if it is inside the RSA modulus "N"
// Returns true if value is within the range, false if not
func InModulus(x *big.Int, N *big.Int) bool{

	if x.Cmp(N) != -1 {
		return false
	}

	return true
}

// ModMul receives values "x" and "y", to be multiplied, and the corresponding modulo "N" to 
// This function is a combination of two big.Int library functions and simplifies code production and readability
func ModMul(x, y, N *big.Int)*big.Int {

	xy := big.NewInt(0).Mul(x, y)

	result := big.NewInt(0).Mod(xy, N)

	return result
}