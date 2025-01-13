package helper

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"loopvector_server_management/helper"

	"github.com/spf13/viper"
)

// Encrypt encrypts the given plaintext using AES-GCM with the master key.
func Encrypt(plainText string) (string, error) {
	masterKey := viper.GetString(helper.KMasterKey)
	// log.Println("encrypt using master key: ", masterKey)
	key := []byte(masterKey)
	if len(key) != 32 {
		return "", fmt.Errorf("master key must be 32 bytes long")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("failed to create cipher block: %w", err)
	}

	nonce := make([]byte, 12) // GCM nonce size is 12 bytes
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", fmt.Errorf("failed to generate nonce: %w", err)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("failed to create GCM cipher: %w", err)
	}

	cipherText := aesGCM.Seal(nil, nonce, []byte(plainText), nil)
	fullCipherText := append(nonce, cipherText...)
	return base64.StdEncoding.EncodeToString(fullCipherText), nil
}

// Decrypt decrypts the given ciphertext using AES-GCM with the master key.
func Decrypt(cipherTextBase64 string) (string, error) {
	masterKey := viper.GetString(helper.KMasterKey)
	// log.Println("decrypt using master key: ", masterKey)
	key := []byte(masterKey)
	if len(key) != 32 {
		return "", fmt.Errorf("master key must be 32 bytes long")
	}
	log.Println("cipherTextBase64: ", cipherTextBase64)
	cipherText, err := base64.StdEncoding.DecodeString(cipherTextBase64)
	if err != nil {
		return "", fmt.Errorf("failed to decode base64 ciphertext: %w", err)
	}

	if len(cipherText) < 12 {
		return "", fmt.Errorf("ciphertext too short")
	}

	nonce, cipherText := cipherText[:12], cipherText[12:]
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("failed to create cipher block: %w", err)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("failed to create GCM cipher: %w", err)
	}

	plainText, err := aesGCM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return "", fmt.Errorf("decryption failed: %w", err)
	}

	return string(plainText), nil
}
