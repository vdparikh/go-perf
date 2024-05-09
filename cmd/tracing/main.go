package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"runtime/trace"
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
	// Tracing Setup
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalf("failed to create trace output file: %v", err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		log.Fatalf("failed to start trace: %v", err)
	}
	defer trace.Stop()

	http.HandleFunc("/encrypt", encryptHandler)

	// http.ListenAndServe blocks until it's not listening to the port (for example if there was an error binding to the port). If it's not run in a separate goroutine, subsequent code (in this case trace.Stop()) will never get called
	go func() {
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	// Exit Signal
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan

	// Clean-up and stop tasks can be done here
}
