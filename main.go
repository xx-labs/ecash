package main

import (
	"CentralBank/cryptography"
	"fmt"
	"math/big"
	"time"
)

func wait(){
	time.Sleep(1*time.Second)
}

func main(){

	// generate RSA system keys
	fmt.Println("Generating RSA system keys for Central Bank...")
	d, e, N, _ := cryptography.Keygen(4096)

	// generate a coin
	fmt.Println("Customer Generating Coin...")
	wait()
	coin, _, _ := cryptography.GenerateCoin(N)

	// cast coin to big int
	c := big.NewInt(0).SetBytes(coin)

	fmt.Println("Customer Generating Blinding Coin...")
	wait()

	// generate blinding factor
	r, _ := cryptography.GenerateBlindingFactor(N)

	// create blinded coin
	b := cryptography.BlindCoin(c, r, e, N)

	// cast blinded coin to bytes
	bc := b.Bytes()

	// produce (blind) signature
	fmt.Println("Bank signs (blinded) coin and returns it...")
	wait()
	signedBlindedCoin := cryptography.Sign(bc, d, N)

	fmt.Println("Client removes blinding factor...")
	wait()
	// Client removes blinding factor from (signed) blinded coin
	signedCoin := cryptography.RemoveBlinding(signedBlindedCoin, r, N)

	fmt.Println("Client verifies signature is correct...")
	wait()

	// Verify that signature is correct
	ok := cryptography.Verify(signedCoin, e, N, coin)

	fmt.Println("Is signature valid? ",ok)
}
