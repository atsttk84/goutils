package crypto

import (
	"fmt"

	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"
)

//https://go.dev/src/crypto/cipher/example_test.go
func EncryptAesCbc(src string) (dest string) {
	key, _ := hex.DecodeString("6368616e676520746869732070617373")
	plaintext := paddingPkcs7([]byte(src))

	if len(plaintext)%aes.BlockSize != 0 {
		panic("plaintext is not a multiple of the block size")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	dest = fmt.Sprintf("%x", ciphertext)
	return
}

func DecryptAesCbc(src string) (dest string) {
	key, _ := hex.DecodeString("6368616e676520746869732070617373")
	ciphertext, _ := hex.DecodeString(src)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	if len(ciphertext)%aes.BlockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	mode.CryptBlocks(ciphertext, ciphertext)

	dest = fmt.Sprintf("%s", ciphertext)
	return
}

func paddingPkcs7(b []byte) (d []byte) {
	padSize := aes.BlockSize - (len(b) % aes.BlockSize)
	pad := bytes.Repeat([]byte{byte(padSize)}, padSize)
	d = append(b, pad...)
	return
}
