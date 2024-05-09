package main

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"net/http"
	"os"
)

func main() {
	caCert, _ := os.ReadFile("./certs/ca.crt")
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	tlsConfig := &tls.Config{
		ClientCAs:  caCertPool,
		ClientAuth: tls.RequireAndVerifyClientCert,
	}

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
