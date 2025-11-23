package lib

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
)

type AESCBC struct{}

func (aesCbc *AESCBC) Ase256Encode(plaintext string, key string, iv string, blockSize int) string {
	bKey := []byte(key)
	bIV := []byte(iv)
	bPlaintext := PKCS5Padding([]byte(plaintext), blockSize, len(plaintext))
	block, err := aes.NewCipher(bKey)
	if err != nil {
		panic(err)
	}
	ciphertext := make([]byte, len(bPlaintext))
	mode := cipher.NewCBCEncrypter(block, bIV)
	mode.CryptBlocks(ciphertext, bPlaintext)
	return hex.EncodeToString(ciphertext)
}

func (aesCbc *AESCBC) Ase256Decode(cipherText string, encKey string, iv string) (decryptedString string) {
	bKey := []byte(encKey)
	bIV := []byte(iv)
	cipherTextDecoded, err := hex.DecodeString(cipherText)
	if err != nil {
		panic(err)
	}

	block, err := aes.NewCipher(bKey)
	if err != nil {
		panic(err)
	}

	mode := cipher.NewCBCDecrypter(block, bIV)
	mode.CryptBlocks([]byte(cipherTextDecoded), []byte(cipherTextDecoded))
	return string(cipherTextDecoded)
}

func PKCS5Padding(ciphertext []byte, blockSize int, after int) []byte {
	padding := (blockSize - len(ciphertext)%blockSize)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
