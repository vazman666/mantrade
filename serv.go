package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", sayhello)           // Устанавливаем роутер
	err := http.ListenAndServe(":8080", nil) // устанавливаем порт веб-сервера
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
func sayhello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Привет!")

	if r.Method == "POST" {
		fmt.Printf("Post, однако\n")
		io.Copy(os.Stdout, r.Body)
		fmt.Printf("\nПараметры: %v\n", r.Body)
		fmt.Printf("параметры запроса %v\n", r.URL.Query())

		decoder := json.NewDecoder(r.Body)
		fmt.Printf("Decoder %v\n", decoder)
		var params map[string]string
		decoder.Decode(&params)
		fmt.Printf("params %v\n", params)
	} else {
		fmt.Printf("Get, однако\n")
		fmt.Printf("r.URL.Query()=%v\n", r.URL.Query())
	}
}
