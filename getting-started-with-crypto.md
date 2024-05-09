# Getting Started with Crypto and Go
This doc will walk you through basic crypto operation in Go

The full code can be found in `cmd/crypto/main.go`


## Prerequisites
Before you begin, ensure you have the following installed:
- Go. This example requires Go 1.13 or later due to the use of certain cryptographic features.
- Docker
- VS Code

## Running the code

Execute `go run cmd/crypto/main.go`

```bash
Original plaintext: Hello, world!

AES Encrypted: 4ff5d9ee9c6f97906474ae110f17ae43f1ece71c810e6b3e17058467e6804ec9cbafc8afa8ab0b9ced
AES Decrypted: Hello, world!

RSA Encrypted: xEGEee16baWVU1JXpofEfzqZ1GAzTIry0J2WSyMXHq32m5SxzXWuBJteLXsUqj+0xrwJi73pPDUx4OmmdUaINrczfyMOry4Cj+j9xnfUQzGYDK/gki+9J0GC958GHzhfapQS5i5xv+z5MrixGTGKKlloNgomKr3Vvew56zQAqznySjDwSlmFJKGZe166VDYCF/T7edLetGAjQM9EsSyfa0v1HqyHYRwOIkB18y2GZFNnvXRq6LoG39OlYx9sip04kCGJjxXVaSycesqQBYIYdOh3Q/qK4DHbt5i8aKU1z7lw2CEthOC5XRzYOA+pCbTZX3CM0F7v9lEmQll+1qSpjQ==
RSA Decrypted: Hello, world!

SHA-256 Checksum: 315f5bdb76d078c43b8ac0064e4a0164612b1fce77c869345bfc94c75894edd3

HMAC: df976418b446bafccf2342dfab02376e9bcd5f5915c931d9e18f269f484e8464
```

## Operations
**AES Encryption/Decryption:**
- AES (Advanced Encryption Standard): This is a symmetric-key encryption algorithm. It means the same secret key is used for both encryption and decryption. AES is a powerful algorithm known for its speed and security, making it a popular choice for bulk encryption of data. 
- This part of the code demonstrates how to encrypt and decrypt text using AES-256. It involves creating an AES cipher block, generating a nonce (random number used once), and performing encryption/decryption. Go implements AES in the crypto/aes package

```go
package main

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/hex"
    "io"
    "log"
)

func main() {
    key := []byte("the-key-has-to-be-32-bytes-long!") // 32 bytes for AES-256
    plaintext := "Hello, world!"

    ciphertext, err := encryptAES(key, plaintext)
    if err != nil {
        log.Fatal(err)
    }
    log.Println("Encrypted:", ciphertext)

    result, err := decryptAES(key, ciphertext)
    if err != nil {
        log.Fatal(err)
    }
    log.Println("Decrypted:", result)
}

func encryptAES(key []byte, plaintext string) (string, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return "", err
    }
    nonce := make([]byte, 12) // Correct size for GCM
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
    nonce, _ := hex.DecodeString(ciphertext[:24])
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
```
**RSA Encryption/Decryption:** 
- RSA (Rivest–Shamir–Adleman) is an asymmetric-key encryption algorithm. Unlike AES, RSA uses a public-key pair for encryption and decryption. The public key can be widely distributed, while the private key is kept secret. Data encrypted with the public key can only be decrypted with the corresponding private key. RSA is often used for secure key exchange due to this property
- This part of code generates RSA keys, then encrypts and decrypts text using RSA with OAEP padding.
- This is commonly used for secure data transmission.

```go
package main

import (
    "crypto/rand"
    "crypto/rsa"
    "crypto/sha256"
    "log"
)

func main() {
    privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
    if err != nil {
        log.Fatal(err)
    }
    publicKey := &privateKey.PublicKey

    plaintext := "Hello, RSA!"
    ciphertext, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, []byte(plaintext), nil)
    if err != nil {
        log.Fatal(err)
    }
    log.Println("RSA Encrypted:", ciphertext)

    decryptedText, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, ciphertext, nil)
    if err != nil {
        log.Fatal(err)
    }
    log.Println("RSA Decrypted:", string(decryptedText))
}
```

**Checksum with SHA-256:** 
- Checksum is a simple way to detect accidental data corruption during transmission or storage. It involves applying a mathematical function to the data and generating a fixed-size value. Any change in the data will likely result in a different checksum value, indicating potential data corruption. Checksums are not as secure as cryptographic hashes like HMAC since they are not designed to withstand malicious modification.
- This part of code calculates a SHA-256 hash of the input data. 
- This is useful for verifying data integrity.

```go
package main

import (
    "crypto/sha256"
    "encoding/hex"
    "log"
)

func main() {
    data := "Hello, world!"
    hash := sha256.Sum256([]byte(data))
    log.Println("SHA-256 Hash:", hex.EncodeToString(hash[:]))
}
```

**HMAC Generation:** 
- HMAC is a message authentication code based on a cryptographic hash function (like SHA-256). It uses a secret key to generate a unique tag for the data. Similar to a checksum, any change in the data or the key will invalidate the HMAC tag. HMAC provides both data integrity and authenticity verification, making it more secure than a simple checksum
- This part of code generates a Hash-based Message Authentication Code (HMAC) to verify both the data integrity and the authenticity of a message.

```go
package main

import (
    "crypto/hmac"
    "crypto/sha256"
    "encoding/hex"
    "log"
)

func main() {
    key := []byte("supersecretkey")
    message := "important message"
    mac := hmac.New(sha256.New, key)
    mac.Write([]byte(message))
    expectedMAC := mac.Sum(nil)
    log.Println("HMAC:", hex.EncodeToString(expectedMAC))
}
```


## Few more details
Below are additional topics and considerations related to cryptography in Go that can enhance your understanding and application of these principles in real-world scenarios.

### Key Management
Key management is a critical aspect of cryptography. How you generate, store, and handle cryptographic keys can significantly affect the security of your cryptographic operations.

- **Key Generation**: Always use a strong, cryptographically secure method to generate keys. Avoid hard-coding keys directly in your source code.
- **Key Storage**: Use secure storage solutions such as hardware security modules (HSM), key management services (KMS) provided by cloud providers, or dedicated secret management tools like HashiCorp Vault.
- **Key Rotation**: Implement key rotation policies to limit the lifetime of keys and reduce the risk of key compromise.

### Digital Signatures and Certificate Handling
Digital signatures are used to verify the authenticity and integrity of data. Go provides robust support for creating and verifying digital signatures using its `crypto` packages.

- **Creating a Digital Signature**: Use the `rsa`, `ecdsa`, or `ed25519` packages for signing data. These packages support different types of cryptographic algorithms suitable for digital signatures.
- **Certificate Handling**: The `crypto/x509` package in Go handles X.509 certificates, which are essential for TLS and SSL communication. Parsing, generating, and validating certificates are all supported.

### Secure Communication
Secure communication is vital for protecting data in transit. Go supports various features to secure communication channels.

- **TLS Support**: The `crypto/tls` package provides a robust implementation of TLS (Transport Layer Security). It is straightforward to set up a TLS server or client using Go to ensure encrypted and authenticated connections.
- **HTTP/2 Support**: Go's `net/http` package includes built-in support for HTTP/2, which inherently uses TLS for security, providing a more efficient and secure web communication protocol.

### Performance Considerations
While security is paramount, the performance of cryptographic operations can also impact the usability and scalability of applications.

- **Benchmarking**: Regularly benchmark cryptographic operations, especially when they are part of critical performance paths in your applications. Go's built-in `testing` package allows you to write benchmarks to measure and optimize performance.
- **Avoid Blocking**: Cryptographic operations, especially those involving disk or network I/O (like fetching keys from a remote service), can block. Use Go's concurrency features, such as goroutines and channels, to handle potentially blocking operations without slowing down your application.

### Cryptography Libraries
Beyond the standard library, several third-party libraries can enhance Go's cryptographic capabilities:

- **libsodium**: There are Go bindings for libsodium, a modern and easy-to-use crypto library that offers additional cryptographic primitives and easier APIs.
- **Google Tink**: Developed by Google, Tink offers a set of multi-language, cross-platform cryptographic APIs that handle many common pitfalls in using cryptographic primitives directly.

## Conclusion
When implementing cryptography in Go, always stay updated with the latest security practices and updates in the Go ecosystem. Cryptography is a rapidly evolving field, and what is considered secure today might not be secure tomorrow. Regularly review and update your cryptographic practices to ensure that your applications remain secure against new vulnerabilities and threats.

