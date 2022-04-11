package cryptography

import (
	"math/big"
	"testing"
)

func TestGenerateBlindingFactor(t *testing.T) {

	_, _, N, _ := Keygen(128)

	ctr := 0
	iterations := 100000

	one := big.NewInt(1)

	for i:=0; i<iterations; i++{

		r, _ := GenerateBlindingFactor(N)

		inv := GetBlindingInverse(r, N)

		result := ModMul(r, inv, N)

		if one.Cmp(result) != 0{
			break
		}
		ctr++
	}

	if ctr!= iterations {
		t.Errorf("Error in the generation of blinding factors (or inverse)")
	}
}
