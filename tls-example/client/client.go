package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// Read the key pair to create certificate
	// cert, err := tls.LoadX509KeyPair("client.crt", "client.key")	//Uncomment if you want enable mTLS validation
	// if err != nil {	//Uncomment if you want enable mTLS validation
	// 	log.Fatal(err)	//Uncomment if you want enable mTLS validation
	// }	//Uncomment if you want enable mTLS validation

	// Create a CA certificate pool and add cert.pem to it
	caCert, err := ioutil.ReadFile("rootCA.crt")
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Create a HTTPS client and supply the created CA pool and certificate
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: caCertPool,
				// Certificates: []tls.Certificate{cert},	//Uncomment if you want enable mTLS validation
			},
		},
	}

	// Request /hello via the created HTTPS client over port 8443 via GET
	r, err := client.Get("https://localhost:8443/hello")
	if err != nil {
		log.Fatal(err)
	}

	// Read the response body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Print the response body to stdout
	fmt.Printf("%s\n", body)
}
