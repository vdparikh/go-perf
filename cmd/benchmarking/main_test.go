package main

import (
	"crypto/rand"
	"testing"
)

func BenchmarkEncryptData(b *testing.B) {
	data := []byte("Hello, World! This is a test string for encryption.")
	key := generateKey()
	for i := 0; i < b.N; i++ {
		_, err := encryptData(data, key)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkEncryptDataSmall(b *testing.B) {
	data := []byte("Short data")
	key := generateKey()
	for i := 0; i < b.N; i++ {
		_, err := encryptData(data, key)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkEncryptDataMedium(b *testing.B) {
	data := make([]byte, 1024) // 1KB of data
	rand.Read(data)            // Fill the data with random bytes
	key := generateKey()
	for i := 0; i < b.N; i++ {
		_, err := encryptData(data, key)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkEncryptDataLarge(b *testing.B) {
	data := make([]byte, 1048576) // 1MB of data
	rand.Read(data)               // Fill the data with random bytes
	key := generateKey()
	for i := 0; i < b.N; i++ {
		_, err := encryptData(data, key)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkEncryptDataAES256(b *testing.B) {
	data := []byte("Example data for AES-256")
	key := make([]byte, 32) // 256 bits for AES-256
	rand.Read(key)          // Generate a random key
	for i := 0; i < b.N; i++ {
		_, err := encryptData(data, key)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkEncryptDataConcurrency(b *testing.B) {
	data := []byte("Concurrent data encryption")
	key := generateKey()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err := encryptData(data, key)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}
