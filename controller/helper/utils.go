package helper

import (
	"crypto/rand"
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
