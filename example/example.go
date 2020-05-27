package example

import (
	"github.com/sudiksa17/Encryption/encryption"
)

func example() {
	key := "key"
	text := "text"
	//For AES and AES with base64 encoded
	//pass key and text that you want to encrypt
	encryptedData := encryption.AESEncrypt(key, text)
	//then for decrypting the data, pass the string
	//that is returned from the above function, it is the encrypted data and pass key along wiht it
	encryption.AESDecrypt(key, encryptedData)

	//Same for AES with base64 encoded
	encryptedData64 := encryption.AES64Encrypt(key, text)
	encryption.AES64Decrypt(key, encryptedData64)

	//For RSA Algorithm you just need to pass the text string, Encrypt function
	//will return encryptedData as well as a private key that will be used for decryption
	rsaEncryptedData, privateKey := encryption.RSAEncrypt(text)
	encryption.RSADecrypt(rsaEncryptedData, privateKey)

}
