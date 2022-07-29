package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Stuff struct {
	things string `json:"things"`
	name   string `json:"name"`
	city   string `json:"city"`
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
	fmt.Println("Posting some things...")

	// TODO not working
	var temp Stuff
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, &temp)
	fmt.Println(temp)
	// TODO END not working

	//var someStuff Stuff
	someStuff := Stuff{things: "thing" + strconv.Itoa(count)}
	//_ = json.NewDecoder(r.Body).Decode(&someStuff)
	stuffs = append(stuffs, someStuff)
	fmt.Println(someStuff)
	fmt.Println(stuffs)

	//res, _ := json.Marshal(someStuff)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//w.Write(res)
	err = json.NewEncoder(w).Encode(someStuff) // TODO this isn't working? empty return json
	if err != nil {
		panic(err)
	}

	fmt.Println()
	count++

}

func main() {
	stuffs = append(stuffs, Stuff{things: "thing1", name: "Ron", city: "UK"})
	stuffs = append(stuffs, Stuff{things: "thing2", name: "Smith", city: "Matrix"})
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

	fmt.Println("Server Running on 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
