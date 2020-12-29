package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"fmt"
	"log"
)

func main() {
	// 生成rsa的密钥
	lowLevelKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("Generate key failed! error := %s", err.Error())
	}

	origData := "hello world!"
	//对数据进行加密的操作
	res, err := rsa.EncryptPKCS1v15(rand.Reader, &lowLevelKey.PublicKey, []byte(origData))
	if err != nil {
		log.Fatalf("Encrypt data failed! Error := %s", err.Error())
	}
	fmt.Printf("%v\n", res)
	//对数据进行解密的操作
	plain, err := rsa.DecryptPKCS1v15(rand.Reader, lowLevelKey, res)
	if err != nil {
		log.Fatalf("Decrypt data failed! Error := %s", err.Error())
	}
	fmt.Printf("%s\n", plain)

	// 对数据执行hash的操作
	hashed := sha512.Sum512([]byte(origData))
	v15, err := rsa.SignPKCS1v15(rand.Reader, lowLevelKey, crypto.SHA512, hashed[:])
	if err != nil {
		log.Fatalf("Sign data failed! error := %s", err.Error())
	}
	fmt.Printf("signature := %v\n", v15)
	//验证数据的签名
	err = rsa.VerifyPKCS1v15(&lowLevelKey.PublicKey, crypto.SHA512, hashed[:], v15)
	if err != nil {
		log.Fatalf("Verify pkcs1v15 failed! error := %s", err.Error())
	}
}
