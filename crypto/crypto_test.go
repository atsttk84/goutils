package crypto

import (
	"fmt"
	"testing"
)

type CryptoTest struct {
	Id       string
	Name     string
	Password string `crypto:"sha256"`
	Email    string `crypto:"aes_cbc"`
}

func TestEncrypt(t *testing.T) {
	input := CryptoTest{"1", "2", "3", "4"}
	fmt.Printf("before: %+v\n", input)
	Encrypt(&input)
	fmt.Printf("encrypt: %+v\n", input)
	Decrypt(&input)
	fmt.Printf("decrypt: %+v\n", input)
}
