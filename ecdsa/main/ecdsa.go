package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha512"
	"fmt"
)

/*ecdsa的签名和验证*/
func main() {
	key, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	if err != nil {
		fmt.Errorf("Generate key failed! error:= %s", err.Error())
	}
	msg := "hello world!"
	hash := sha512.Sum384([]byte(msg))

	sign, err := ecdsa.SignASN1(rand.Reader, key, hash[:])
	if err != nil {
		fmt.Errorf("Sign Error! error:= %s", err.Error())
	}

	fmt.Printf("signature := %x\n", sign)

	valid := ecdsa.VerifyASN1(&key.PublicKey, hash[:], sign)
	fmt.Printf("signature verified! %t", valid)

}
