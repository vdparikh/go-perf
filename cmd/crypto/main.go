package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {
	// AES Encryption/Decryption
	key := []byte("the-key-has-to-be-32-bytes-long!") // 32 bytes for AES-256
	plaintext := "Hello, world!"
	fmt.Println("Original plaintext:", plaintext)

	ciphertext, err := encryptAES(key, plaintext)
	if err != nil {
		fmt.Println("Error encrypting:", err)
		return
	}
	fmt.Println("AES Encrypted:", ciphertext)

	decrypted, err := decryptAES(key, ciphertext)
	if err != nil {
		fmt.Println("Error decrypting:", err)
		return
	}
	fmt.Println("AES Decrypted:", decrypted)

	// RSA Encryption/Decryption
	privateKey, publicKey, err := generateRSAKeys()
	if err != nil {
		fmt.Println("Error generating RSA keys:", err)
		return
	}

	rsaCiphertext, err := encryptRSA(publicKey, plaintext)
	if err != nil {
		fmt.Println("Error RSA encrypting:", err)
		return
	}
	fmt.Println("RSA Encrypted:", base64.StdEncoding.EncodeToString(rsaCiphertext))

	rsaDecrypted, err := decryptRSA(privateKey, rsaCiphertext)
	if err != nil {
		fmt.Println("Error RSA decrypting:", err)
		return
	}
	fmt.Println("RSA Decrypted:", rsaDecrypted)

	// Checksum using SHA-256
	checksum := generateChecksum(plaintext)
	fmt.Println("SHA-256 Checksum:", checksum)

	// HMAC
	message := "important message"
	hmac := generateHMAC(key, message)
	fmt.Println("HMAC:", hmac)
}

func encryptAES(key []byte, plaintext string) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	nonce := make([]byte, 12) // Correct nonce length for GCM
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	ciphertext := aesgcm.Seal(nil, nonce, []byte(plaintext), nil)
	return hex.EncodeToString(nonce) + hex.EncodeToString(ciphertext), nil
}

func decryptAES(key []byte, ciphertext string) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	nonce, _ := hex.DecodeString(ciphertext[:24]) // 12 bytes nonce, each byte encoded as 2 hex characters
	ciphertextBytes, _ := hex.DecodeString(ciphertext[24:])
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	plaintextBytes, err := aesgcm.Open(nil, nonce, ciphertextBytes, nil)
	if err != nil {
		return "", err
	}
	return string(plaintextBytes), nil
}

func generateRSAKeys() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}
	return privateKey, &privateKey.PublicKey, nil
}

func encryptRSA(publicKey *rsa.PublicKey, plaintext string) ([]byte, error) {
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, []byte(plaintext), nil)
	if err != nil {
		return nil, err
	}
	return ciphertext, nil
}

func decryptRSA(privateKey *rsa.PrivateKey, ciphertext []byte) (string, error) {
	plaintext, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, ciphertext, nil)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}

func generateChecksum(data string) string {
	hasher := sha256.New()
	hasher.Write([]byte(data))
	return hex.EncodeToString(hasher.Sum(nil))
}

func generateHMAC(key []byte, data string) string {
	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(data))
	return hex.EncodeToString(mac.Sum(nil))
}
