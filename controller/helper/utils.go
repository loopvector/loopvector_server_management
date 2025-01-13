package helper

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/big"
	"time"
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

// func HashPassword(password string) (string, error) {
// 	saltedBytes := []byte(password) // Implement salt generation
// 	hashedBytes, err := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)
// 	if err != nil {
// 		return "", err
// 	}

// 	hash := string(hashedBytes[:])
// 	return hash, nil
// }

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

// func generateSalt(length int) ([]byte, error) {
// 	// Allocate a byte slice to store the random bytes
// 	salt := make([]byte, length)

// 	// Read random bytes into the slice
// 	_, err := rand.Read(salt)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to generate salt: %w", err)
// 	}

// 	// Encode the random bytes to a base64 string for easy storage
// 	return salt, nil
// }

func VerifyPassword(password string, storedHashedPassword string) bool {
	storedPassword, err := Decrypt(storedHashedPassword)
	if err != nil {
		panic(err)
	}

	// Compare the hashes in constant time
	if subtle.ConstantTimeCompare([]byte(password), []byte(storedPassword)) != 1 {
		return false // Password does not match
	}

	return true // Password matches
}

func GenerateToken() string {
	bytes := make([]byte, 16) // 16 bytes = 128 bits
	if _, err := rand.Read(bytes); err != nil {
		panic(err)
	}
	return hex.EncodeToString(bytes)
}
