package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func RandomHex() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
