package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func timeHandler(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(time.RFC1123)
	w.Write([]byte("The time is " + tm))
}

func getStuff(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting Stuff...")

	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotImplemented)
		return
	}

	m := make(map[string]string)
	m["things"] = "Stuff and Things"

	res, _ := json.Marshal(m)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func main() {
	r := mux.NewRouter()

	fileServer := http.FileServer(http.Dir("./public"))
	r.Handle("/", fileServer)
	r.PathPrefix("/assets/").Handler(fileServer)
	r.PathPrefix("/pages/").Handler(fileServer)

	redirectHandler := http.RedirectHandler("https://www.deegrant.com", 307)
	r.Handle("/deegrant", redirectHandler)

	r.HandleFunc("/time", timeHandler)

	r.HandleFunc("/getStuff", getStuff)

	fmt.Println("Server Running on 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
