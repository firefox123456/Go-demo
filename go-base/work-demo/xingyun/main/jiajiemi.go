package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

func main() {
	// 原始数据和密钥
	plaintext := []byte("0sO&651(TEm+Vm")
	key := []byte("yunti00000000000")
	iv := []byte("yunti00000000000")

	// 加密
	ciphertext, err := encrypt(plaintext, key, iv)
	if err != nil {
		fmt.Println("Error encrypting:", err)
		return
	}
	fmt.Printf("Encrypted data: %s\n", base64.StdEncoding.EncodeToString(ciphertext))

	// 解密
	decrypted, err := decrypt(ciphertext, key, iv)
	if err != nil {
		fmt.Println("Error decrypting:", err)
		return
	}
	fmt.Printf("Decrypted data: %s\n", string(decrypted))
}

func encrypt(plaintext, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 使用PKCS5Padding填充
	blockSize := block.BlockSize()
	padding := blockSize - len(plaintext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	plaintext = append(plaintext, padtext...)

	ciphertext := make([]byte, len(plaintext))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, plaintext)

	return ciphertext, nil
}

func decrypt(ciphertext, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	// 移除PKCS5Padding填充
	unpadder := func(src []byte) ([]byte, error) {
		padding := int(src[len(src)-1])
		return src[:len(src)-padding], nil
	}
	plaintext, err := unpadder(ciphertext)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
