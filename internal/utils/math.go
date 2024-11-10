package utils

import (
	"crypto/rand"
	"fmt"
)

func GenShorten(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	result := make([]byte, length)

	_, err := rand.Read(result)
	if err != nil {
		return "", fmt.Errorf("failed to generate random bytes: %v", err)
	}

	for i := 0; i < length; i++ {
		result[i] = charset[int(result[i])%len(charset)]
	}

	return string(result), nil
}
