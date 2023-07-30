package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
)

func HiHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Hello, world!\n"))
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hi", HiHandler)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", 9443),
		Handler: mux,
		TLSConfig: &tls.Config{
			GetCertificate: func(*tls.ClientHelloInfo) (*tls.Certificate, error) {
				// Always get latest certs/srv.crt and certs/srv.key
				// ex: keeping certificates file somewhere in global location where created certificates updated and this closure function can refer that
				cert, err := tls.LoadX509KeyPair("certs/srv.crt", "certs/srv.key")
				if err != nil {
					return nil, err
				}
				return &cert, nil
			},
		},
	}

	// run server on port "8443"
	log.Fatal(srv.ListenAndServeTLS("", ""))
}
