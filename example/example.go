package example

import (
	"Encryption/algos/aes"
	"Encryption/algos/rsa"
)

func example() {
	key := "key"
	text := "text"
	//For AES and AES with base64 encoded
	//pass key and text that you want to encrypt
	encryptedData := aes.Encrypt(key, text)
	//then for decrypting the data, pass the string
	//that is returned from the above function, it is the encrypted data and pass key along wiht it
	aes.Decrypt(key, encryptedData)

	//Same for AES with base64 encoded
	encryptedData64 := aes.Encrypt64(key, text)
	aes.Decrypt64(key, encryptedData64)

	//For RSA Algorithm you just need to pass the text string, Encrypt function
	//will return encryptedData as well as a private key that will be used for decryption
	rsaEncryptedData, privateKey := rsa.Encrypt(text)
	rsa.Decrypt(rsaEncryptedData, privateKey)

}
