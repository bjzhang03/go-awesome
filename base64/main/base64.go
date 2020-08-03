package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	msg := "Hello, 世界"
	encoded := base64.StdEncoding.EncodeToString([]byte(msg))
	fmt.Println(encoded)
	decoded, err := base64.StdEncoding.DecodeString("CgsKBQECAwQFEAYYPw==")
	if err != nil {
		fmt.Println("decode error:", err)
		return
	}
	fmt.Println(decoded)
	// Output:
	// SGVsbG8sIOS4lueVjA==
	// Hello, 世界
}
