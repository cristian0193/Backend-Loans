package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"io"

	"golang.org/x/crypto/pbkdf2"
)

const saltlen = 12
const keylen = 32
const iterations = 10002

func Encrypt(email string) string {

	header := make([]byte, saltlen+aes.BlockSize)

	salt := header[:saltlen]
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		panic(err)
	}

	iv := header[saltlen : aes.BlockSize+saltlen]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	key := pbkdf2.Key([]byte(GetStrEnv("PASSWORD_KEY")), salt, iterations, keylen, md5.New)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	ciphertext := make([]byte, len(header)+len(email))
	copy(ciphertext, header)

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize+saltlen:], []byte(email))
	return base64Encode(ciphertext)
}

func Decrypt(encrypted string) string {
	a, err := base64Decode([]byte(encrypted))
	if err != nil {
		panic(err)
	}
	ciphertext := a
	salt := ciphertext[:saltlen]
	iv := ciphertext[saltlen : aes.BlockSize+saltlen]
	key := pbkdf2.Key([]byte(GetStrEnv("PASSWORD_KEY")), salt, iterations, keylen, md5.New)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	if len(ciphertext) < aes.BlockSize {
		return ""
	}

	decrypted := ciphertext[saltlen+aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(decrypted, decrypted)

	return string(decrypted)
}

func base64Encode(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}

func base64Decode(src []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(src))
}
