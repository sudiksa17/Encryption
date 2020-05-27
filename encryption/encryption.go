package encryption

import (
	"Encryption/algos/aes"
	rsaAlgo "Encryption/algos/rsa"
	"crypto/rsa"
)

func AESEncrypt(key, text string) string {
	return aes.Encrypt(key, text)
}

func AESDecrypt(key, encryptedData string) string {
	return aes.Decrypt(key, encryptedData)
}

func AES64Encrypt(key, text string) string {
	return aes.Encrypt64(key, text)
}

func AES64Decrypt(key, encryptedData string) string {
	return aes.Decrypt64(key, encryptedData)
}

func RSAEncrypt(text string) (string, *rsa.PrivateKey) {
	return rsaAlgo.Encrypt(text)
}

func RSADecrypt(encryptedData string, privateKey *rsa.PrivateKey) string {
	return rsaAlgo.Decrypt(encryptedData, privateKey)
}
