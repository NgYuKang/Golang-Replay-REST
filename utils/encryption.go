package utils

import (
	"Golang-Replay-REST/configs"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

func Decrypt(encryptedByte []byte) ([]byte, error) {
	block, err := aes.NewCipher([]byte(configs.EnvEncryptKey()))
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := encryptedByte[:gcm.NonceSize()]
	cipherText := encryptedByte[gcm.NonceSize():]
	plainText, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return nil, err
	}

	return plainText, nil

}

func Encrypt(rawFile []byte) ([]byte, error) {
	block, err := aes.NewCipher([]byte(configs.EnvEncryptKey()))
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	encryptedBytes := gcm.Seal(nonce, nonce, rawFile, nil)

	return encryptedBytes, nil
}
