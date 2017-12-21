package digest

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
	"math"
	"strconv"
)

const NUM_ZEROES = 6

type nonceError struct {
	e string
}

func (n nonceError) Error() string {
	return n.e
}

func hash (s string) string {
	sum := sha256.Sum256([]byte(s))
	return hex.EncodeToString(sum[:])
}

func FindNonce (s string) (string, error) {
	nonce := float64(0)
	count := 0
	nonceIsValid := false
	maxFloat := math.Pow(2, 32)
	var strNonce string
	for ; nonce < maxFloat && !nonceIsValid; nonce++ {
		count++
		strNonce = strconv.FormatFloat(nonce, 'f', 0, 32)
		nonceIsValid = ValidNonce(strNonce, s)
	}

	if ValidNonce(strNonce, s) {
		return strNonce, nil
	} else {
		return "", error(nonceError{"Error, unable to find nonce."})
	}
}

func ValidNonce (nonce string, message string) bool {
	messageNonceHash := hash(message + nonce)
	splitSum := strings.Split(messageNonceHash, "")
	for char := range splitSum[:NUM_ZEROES] {
		if splitSum[char] != "0" {
			return false
		}
	}
	return true
}