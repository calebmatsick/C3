/*
  Package for encrypting client-server traffic
*/

package security

import (
	// Standard
	"crypto/aes"
	"crypto/cipher"
	crand "crypto/rand"
	"fmt"
	"io"
	mrand "math/rand"
	"time"
)

const (
	seedString = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	letterIdxBits = 6
	letterIdxMask = 1<<letterIdxBits - 1
	letterIdxMax = 63 /letterIdxBits
	keyLength = 32
)

var (
	src = mrand.NewSource(time.Now().UnixNano())
)

// Encrypts traffic, returns byte array
func Encrypt(input string) []byte {
	mes := []byte(input)
	key := randKey()

	c, err := aes.NewCipher(key)

	if err != nil {
		fmt.Println(err)
	}

	gcm, err := cipher.NewGCM(c)

	if err != nil {
		fmt.Println(err)
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(crand.Reader, nonce); err != nil {
		fmt.Println(err)
	}

	mes = gcm.Seal(nonce, nonce, mes, nil)

	return mes
}


// Decrypts traffic, returns plaintext
func Decrypt(ciphertext []byte) string {
	key := []byte("passphrasewhichneedstobe32bytes!")

	c, err := aes.NewCipher(key)

	if err != nil {
		fmt.Println(err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err)
	}

	nonceSize := gcm.NonceSize()

	if len(ciphertext) < nonceSize {
		fmt.Println(err)
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	decCmd, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		fmt.Println(err)
	}

	return string(decCmd)
}


// Generate random key
func randKey() []byte {
	b := make([]byte, keyLength)
	for i, cache, remain := keyLength-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(seedString) {
			b[i] = seedString[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return b
}