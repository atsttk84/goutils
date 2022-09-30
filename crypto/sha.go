package crypto

import (
	"crypto/sha256"
	"fmt"
)

// https://pkg.go.dev/crypto/sha256
func SumSha256(s string) (sha string) {
	sum := sha256.Sum256([]byte(s))
	sha = fmt.Sprintf("%x", sum)
	return
}
