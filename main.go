package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Stuff struct {
	Things string `json:"things"`
	Name   string `json:"name"`
	City   string `json:"city"`
}

var stuffs []Stuff
var count int

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

func postStuff(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var stuff Stuff
	_ = json.NewDecoder(r.Body).Decode(&stuff)
	stuffs = append(stuffs, stuff)

	json.NewEncoder(w).Encode(stuff)
}

func formSubmit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var stuff Stuff
	_ = json.NewDecoder(r.Body).Decode(&stuff)
	stuff.Things = "thing" + strconv.Itoa(count)
	count++
	stuffs = append(stuffs, stuff)

	fmt.Println(stuffs)
	json.NewEncoder(w).Encode(stuffs)
}

func main() {
	stuffs = append(stuffs, Stuff{Things: "thing1", Name: "Ron", City: "UK"})
	stuffs = append(stuffs, Stuff{Things: "thing2", Name: "Smith", City: "Matrix"})
	count = 3

	r := mux.NewRouter()

	fileServer := http.FileServer(http.Dir("./public"))
	r.Handle("/", fileServer)
	r.PathPrefix("/assets/").Handler(fileServer)
	r.PathPrefix("/pages/").Handler(fileServer)

	redirectHandler := http.RedirectHandler("https://www.deegrant.com", 307)
	r.Handle("/deegrant", redirectHandler)

	r.HandleFunc("/time", timeHandler)
	r.HandleFunc("/getStuff", getStuff)
	r.HandleFunc("/postStuff", postStuff)
	r.HandleFunc("/formSubmit", formSubmit)

	fmt.Println("Server Running on 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
