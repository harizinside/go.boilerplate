package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

// Argon2 Parameters
var (
	Memory      uint32 = 64 * 1024
	Iterations  uint32 = 3
	Parallelism uint8  = 2
	SaltLength  int    = 16
	KeyLength   uint32 = 32
)

// HashPassword hashes the plain text password using Argon2
func HashPassword(password string) (string, error) {
	// Generate a random salt
	salt := make([]byte, SaltLength)
	if _, err := rand.Read(salt); err != nil {
		return "", fmt.Errorf("failed to generate salt: %v", err)
	}

	// Hash the password
	hash := argon2.IDKey([]byte(password), salt, Iterations, Memory, Parallelism, KeyLength)

	// Encode the salt and hash in a single string
	encodedHash := fmt.Sprintf("%s$%s", base64.RawStdEncoding.EncodeToString(salt), base64.RawStdEncoding.EncodeToString(hash))
	return encodedHash, nil
}

// VerifyPassword compares the plain text password with the stored Argon2 hash
func VerifyPassword(password, encodedHash string) (bool, error) {
	// Split the hash into salt and key
	parts := strings.Split(encodedHash, "$")
	if len(parts) != 2 {
		return false, fmt.Errorf("invalid hash format")
	}

	// Decode the salt
	salt, err := base64.RawStdEncoding.DecodeString(parts[0])
	if err != nil {
		return false, fmt.Errorf("failed to decode salt: %v", err)
	}

	// Decode the hash
	expectedHash, err := base64.RawStdEncoding.DecodeString(parts[1])
	if err != nil {
		return false, fmt.Errorf("failed to decode hash: %v", err)
	}

	// Hash the input password using the same parameters and salt
	hash := argon2.IDKey([]byte(password), salt, Iterations, Memory, Parallelism, uint32(len(expectedHash)))

	// Compare the hashes
	if !equalSlices(hash, expectedHash) {
		return false, nil
	}
	return true, nil
}

func equalSlices(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
