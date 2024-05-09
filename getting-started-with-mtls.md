# Getting Started with mTLS and Go

Creating a mutual TLS (mTLS) setup in Go is a good way to secure communication between clients and servers by ensuring both parties authenticate each other with their certificates. This doc provides a detailed example of setting up mTLS in Go, along with instructions on how to generate the necessary certificates and keys. 

Mutual TLS (mTLS) is an extension of TLS (Transport Layer Security) that provides two-way authentication. While TLS typically involves the client verifying the identity of the server, mTLS adds an additional layer where the server also verifies the identity of the client. This is particularly useful for secure server-to-server communications or for client-server applications where both parties need to confirm the other’s identity.



## Prerequisites
- Go installed on your machine.
- OpenSSL for generating certificates (or any other tool that can generate SSL certificates).

## Source Code
For this project, you will have the following files:
- `cmd/mtls/server.go` — Contains the code for the mTLS server.
- `cmd/mtls/client.go` — Contains the code for the mTLS client.


## Generate Certificates
First, you need to generate CA, server, and client certificates. Open your terminal and execute the following commands:

```bash
cd certs/

# Generate the CA key and certificate 
openssl req -new -x509 -days 365 -nodes -newkey rsa:2048 -keyout ca.key -out ca.crt -subj "/CN=Example CA" -extensions v3_ca -config <(printf "[req]\ndistinguished_name=req\n[ v3_ca ]\nbasicConstraints = critical,CA:true")

# Create the Server Key, CSR, and Certificate
openssl req -new -nodes -keyout server.key -out server.csr -config openssl.cnf
openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 365 -extensions req_ext -extfile openssl.cnf

# Create the Client Key and CSR
openssl req -new -nodes -keyout client.key -out client.csr -config openssl.cnf
openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out client.crt -days 365 -extensions req_ext -extfile openssl.cnf
```

These commands create keys and certificates for the CA, server, and client, which are necessary for mTLS.

If the private key is encrypted, decrypt it using OpenSSL before using it in your Go application:
```bash
openssl rsa -in server.key -out server_decrypted.key
```

You can verify your certificates using 
```bash
openssl verify -CAfile ca.crt server.crt
openssl verify -CAfile ca.crt client.crt
```
These commands should return `OK` if everything is set up correctly.


## Server Code

The server is set up to listen on HTTPS with mTLS. It requires clients to present a certificate that is signed by a trusted CA (in this case, ca.crt). Here's how the server is configured:

```go
tlsConfig := &tls.Config{
    ClientCAs:  caCertPool,
    ClientAuth: tls.RequireAndVerifyClientCert,
}

server := &http.Server{
    Addr:      ":8443",
    TLSConfig: tlsConfig,
}
```

- ClientCAs: Specifies the pool of CAs that server trusts. Only client certificates signed by these CAs are accepted.
- ClientAuth: This field is set to RequireAndVerifyClientCert, which mandates that the client must provide a valid certificate for the connection to proceed.

`server.go`
```go
package main

import (
    "crypto/tls"
    "crypto/x509"
    "io/ioutil"
    "log"
    "net/http"
)

func main() {
    caCert, _ := ioutil.ReadFile("certs/ca.crt")
    caCertPool := x509.NewCertPool()
    caCertPool.AppendCertsFromPEM(caCert)

    tlsConfig := &tls.Config{
        ClientCAs:  caCertPool,
        ClientAuth: tls.RequireAndVerifyClientCert,
    }
    tlsConfig.BuildNameToCertificate()

    server := &http.Server{
        Addr:      ":8443",
        TLSConfig: tlsConfig,
    }

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("Hello, world!"))
    })

    log.Fatal(server.ListenAndServeTLS("certs/server.crt", "certs/server.key"))
}
```

### Running the Server
Execute the following command:
```bash
go run server.go
```

## Client Code
The client is configured to securely connect to the server using TLS and also presents its own certificate for server-side verification:

```go
tlsConfig := &tls.Config{
    RootCAs:      caCertPool,
    Certificates: []tls.Certificate{cert},
}

client := &http.Client{
    Transport: &http.Transport{
        TLSClientConfig: tlsConfig,
    },
}
```
- RootCAs: Contains the CA certificates that the client trusts. For successful connection, the server's certificate must be signed by one of these CAs.
- Certificates: This is the client's certificate and corresponding private key, which are presented to the server during the TLS handshake.


`client.go`
```go
package main

import (
    "crypto/tls"
    "crypto/x509"
    "io/ioutil"
    "log"
    "net/http"
)

func main() {
    caCert, _ := ioutil.ReadFile("ca.crt")
    caCertPool := x509.NewCertPool()
    caCertPool.AppendCertsFromPEM(caCert)

    cert, _ := tls.LoadX509KeyPair("client.crt", "client.key")
    client := &http.Client{
        Transport: &http.Transport{
            TLSClientConfig: &tls.Config{
                RootCAs:      caCertPool,
                Certificates: []tls.Certificate{cert},
            },
        },
    }

    resp, err := client.Get("https://localhost:8443")
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)
    log.Println(string(body))
}
```

### Running the Client
In another terminal, execute the following command:
```bash
go run client.go
```

## Expected Output
The client should receive a "Hello, world!" response from the server, confirming that the mTLS connection is established and working.

