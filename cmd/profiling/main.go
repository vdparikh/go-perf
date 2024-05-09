package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"
	"net/http"
	_ "net/http/pprof"
)

func generateKey() []byte {
	return []byte("examplekey123456")
}

func encryptData(data []byte, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	ciphertext := make([]byte, aes.BlockSize+len(data))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	// +++++ PROFILING +++++
	// Intentionally inefficient: allocate new slice each iteration
	encrypted := make([]byte, len(data))
	stream.XORKeyStream(encrypted, data)
	ciphertext = append(iv, encrypted...)

	// More efficient: use the existing slice
	// stream.XORKeyStream(ciphertext[aes.BlockSize:], data)
	return hex.EncodeToString(ciphertext), nil
}

func encryptHandler(w http.ResponseWriter, r *http.Request) {
	data := []byte("Hello, World! This is a test string for encryption.") // Example data
	key := generateKey()
	encrypted, err := encryptData(data, key)
	if err != nil {
		http.Error(w, "Failed to encrypt data", http.StatusInternalServerError)
		return
	}
	w.Write([]byte(encrypted))
}

func main() {
	// Handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/encrypt", encryptHandler)
	http.ListenAndServe(":8080", nil)
}
