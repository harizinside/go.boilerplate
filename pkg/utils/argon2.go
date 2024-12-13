package utils

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

// Argon2 parameters
var (
	time    uint32 = 1         // Number of iterations
	memory  uint32 = 64 * 1024 // Memory usage in KiB (64MB)
	threads uint8  = 4         // Number of threads
	keyLen  uint32 = 32        // Length of the hash (in bytes)
	saltLen        = 16        // Salt length
)

// generateSalt creates a cryptographically secure salt.
func generateSalt(length int) ([]byte, error) {
	salt := make([]byte, length)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

// HashPassword hashes a plain-text password using Argon2.
func HashPassword(password string) (string, error) {
	salt, err := generateSalt(saltLen)
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, time, memory, threads, keyLen)

	// Encode hash, salt, and parameters into a single string
	encodedHash := fmt.Sprintf(
		"$argon2id$v=%d$t=%d$m=%d$p=%d$%s$%s",
		argon2.Version,
		time,
		memory,
		threads,
		base64.RawStdEncoding.EncodeToString(salt),
		base64.RawStdEncoding.EncodeToString(hash),
	)
	return encodedHash, nil
}

// VerifyPassword compares a plain-text password against a stored hash.
func VerifyPassword(password, encodedHash string) (bool, error) {
	// Trim leading/trailing spaces and normalize the format
	encodedHash = strings.TrimSpace(encodedHash)

	// Ensure correct format by splitting into parts
	parts := strings.Split(encodedHash, "$")
	if len(parts) != 8 {
		return false, errors.New("invalid hash format")
	}

	// Extract parameters
	var version int
	_, err := fmt.Sscanf(parts[2], "v=%d", &version)
	if err != nil || version != argon2.Version {
		return false, errors.New("invalid Argon2 version")
	}

	var t, m uint32
	var p uint8
	_, err = fmt.Sscanf(parts[3], "t=%d", &t)
	if err != nil {
		return false, err
	}
	_, err = fmt.Sscanf(parts[4], "m=%d", &m)
	if err != nil {
		return false, err
	}
	_, err = fmt.Sscanf(parts[5], "p=%d", &p)
	if err != nil {
		return false, err
	}

	// Decode salt and hash
	salt, err := base64.RawStdEncoding.DecodeString(parts[6])
	if err != nil {
		return false, errors.New("invalid salt encoding")
	}
	expectedHash, err := base64.RawStdEncoding.DecodeString(parts[7])
	if err != nil {
		return false, errors.New("invalid hash encoding")
	}

	// Generate hash using the same parameters and salt
	hash := argon2.IDKey([]byte(password), salt, t, m, p, keyLen)

	// Compare the generated hash and stored hash
	if !equalSlices(hash, expectedHash) {
		return false, nil
	}
	return true, nil
}

// Compare two byte slices securely
func equalSlices(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	var result byte
	for i := range a {
		result |= a[i] ^ b[i]
	}
	return result == 0
}
