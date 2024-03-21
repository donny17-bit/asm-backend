package helper

import (
	"crypto/rand"
	"math/big"
)

func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// Generate cryptographically secure random bytes
	randBytes := make([]byte, length)
	for i := 0; i < length; i++ {
		// Generate a random index within the charset length
		randIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			panic(err) // Handle error if any
		}
		// Use the random index to select a character from the charset
		randBytes[i] = charset[randIndex.Int64()]
	}

	// Convert bytes to string and return
	return string(randBytes)
}
