package decryption

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func decryptCipherText(ciphertext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	nonceSize := aesGCM.NonceSize()
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}

func decode(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

func Decrypt(text, key string) string {
	ciphertext := decode(text)
	mKey := decode(key)
	plaintext, _ := decryptCipherText(ciphertext, mKey)
	return string(plaintext)
}
