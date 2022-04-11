package cryptography

import "testing"

func TestRSA(t *testing.T) {

	d, e, N, err := Keygen(512)

	if err != nil{
		t.Errorf("Keygen(): Error generating the RSA system keys")
	}

	msg := []byte("Super Mario")

	signature := Sign(msg, d, N)

	ok := Verify(signature, e, N, msg)

	if !ok {
		t.Errorf("Error verifying the produced RSA signature")
	}

}