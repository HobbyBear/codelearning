package main

import (
	"net/http"
	"os"
	"syscall"
)

func main() {

	http.HandleFunc("/tradition", func(writer http.ResponseWriter, request *http.Request) {

		f, _ := os.Open("./testmmap.txt")
		buf := make([]byte, 1024)
		n, _ := f.Read(buf)
		writer.Write(buf[:n])
	})

	http.HandleFunc("/mmap", func(writer http.ResponseWriter, request *http.Request) {
		f, _ := os.Open("./testmmap.txt")
		data, err := syscall.Mmap(int(f.Fd()), 0, 5, syscall.PROT_READ, syscall.MAP_SHARED)
		if err != nil {
			panic(err)
		}
		writer.Write(data)
	})
	http.ListenAndServe(":8080", http.DefaultServeMux)
}
