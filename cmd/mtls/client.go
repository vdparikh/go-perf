package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	caCert, _ := os.ReadFile("./certs/ca.crt")
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	cert, _ := tls.LoadX509KeyPair("./certs/client.crt", "./certs/client.key")
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
