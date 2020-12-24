package tools

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"

	"github.com/bayu-aditya/myfacilities-backend/lib/variable"
)

// PasswordFromPlainText cryptography using AES 256 CBC
func PasswordFromPlainText(plain string) string {
	cryptoVar := variable.Project.Crypto

	key := cryptoVar.Key
	iv := cryptoVar.IV
	return aes256(plain, key, iv, aes.BlockSize)
}

func aes256(plaintext string, key string, iv string, blockSize int) string {
	bKey := []byte(key)
	bIV := []byte(iv)
	bPlaintext := pKCS5Padding([]byte(plaintext), blockSize, len(plaintext))
	block, _ := aes.NewCipher(bKey)
	ciphertext := make([]byte, len(bPlaintext))
	mode := cipher.NewCBCEncrypter(block, bIV)
	mode.CryptBlocks(ciphertext, bPlaintext)
	return hex.EncodeToString(ciphertext)
}

func pKCS5Padding(ciphertext []byte, blockSize int, after int) []byte {
	padding := (blockSize - len(ciphertext)%blockSize)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
