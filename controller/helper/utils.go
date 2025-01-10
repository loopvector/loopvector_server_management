package helper

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

	"golang.org/x/crypto/argon2"
)

func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz"
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		index, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		result[i] = charset[index.Int64()]
	}
	return string(result)
}

func GenerateRandomUniqueServerName() string {
	// Get the current timestamp in your desired format
	timestamp := time.Now().Format("20060102_150405")

	// Append the timestamp to the displayName
	displayNameWithTimestamp := fmt.Sprintf("%s_%s", GenerateRandomString(8), timestamp)
	return displayNameWithTimestamp
}

func HashPassword(password string) string {
	salt, err := generateSalt(16) // Implement salt generation
	if err != nil {
		panic(err)
	}
	hash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)
	return encodeBytesToString(hash)
}

func decodeStringToBytes(encodedSalt string) ([]byte, error) {
	// Use base64.StdEncoding.DecodeString to decode the string
	decoded, err := base64.StdEncoding.DecodeString(encodedSalt)
	if err != nil {
		return nil, fmt.Errorf("failed to decode salt: %w", err)
	}
	return decoded, nil
}

func encodeBytesToString(decodeSalt []byte) string {
	// Use base64.RawStdEncoding.EncodeToString to decode the string
	encoded := base64.RawStdEncoding.EncodeToString(decodeSalt)
	return encoded
}

func generateSalt(length int) ([]byte, error) {
	// Allocate a byte slice to store the random bytes
	salt := make([]byte, length)

	// Read random bytes into the slice
	_, err := rand.Read(salt)
	if err != nil {
		return nil, fmt.Errorf("failed to generate salt: %w", err)
	}

	// Encode the random bytes to a base64 string for easy storage
	return salt, nil
}

func VerifyPassword(password string, storedHash string) (bool, error) {
	const argon2Params = "v=19,m=65536,t=3,p=2"
	// Split the stored hash into parts: salt, parameters, and hash
	parts := strings.Split(storedHash, "$")
	if len(parts) != 4 {
		return false, errors.New("invalid stored hash format")
	}

	// Extract parts
	saltBase64 := parts[1]
	hashBase64 := parts[2]
	storedHashParams := parts[0]

	// Decode the salt and stored hash
	salt, err := base64.StdEncoding.DecodeString(saltBase64)
	if err != nil {
		return false, fmt.Errorf("failed to decode salt: %w", err)
	}

	storedHashBytes, err := base64.StdEncoding.DecodeString(hashBase64)
	if err != nil {
		return false, fmt.Errorf("failed to decode hash: %w", err)
	}

	// Match hashing parameters
	if storedHashParams != argon2Params {
		return false, errors.New("hashing parameters mismatch")
	}

	// Recompute hash
	hash := argon2.IDKey([]byte(password), salt, 3, 64*1024, 2, uint32(len(storedHashBytes)))

	// Compare the hashes in constant time
	if subtle.ConstantTimeCompare(hash, storedHashBytes) != 1 {
		return false, nil // Password does not match
	}

	return true, nil // Password matches
}

func GenerateToken() string {
	bytes := make([]byte, 16) // 16 bytes = 128 bits
	if _, err := rand.Read(bytes); err != nil {
		panic(err)
	}
	return hex.EncodeToString(bytes)
}
