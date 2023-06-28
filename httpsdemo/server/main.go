package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
)

func main() {

	s := &http.Server{
		Addr: ":443",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, World!\n")
		}),
		TLSConfig: &tls.Config{
			ClientAuth: tls.RequireAndVerifyClientCert,
		},
	}

	fmt.Println(s.ListenAndServeTLS("./final_csr.crt",
		"./final_private.key"))
}
