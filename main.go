package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func timeHandler(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(time.RFC1123)
	w.Write([]byte("The time is " + tm))
}

func main() {
	r := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./public"))
	r.Handle("/", fileServer)

	redirectHandler := http.RedirectHandler("https://www.deegrant.com", 307)
	r.Handle("/deegrant", redirectHandler)

	r.HandleFunc("/time", timeHandler)

	fmt.Println("Server Running on 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
