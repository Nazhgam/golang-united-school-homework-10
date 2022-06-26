package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

const (
	hello   = "Hello, %s!"
	message = "I got message:\n%s"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()
	router.HandleFunc("/name/{PARAM}", Greater).Methods(http.MethodGet)
	router.HandleFunc("/bad", Bad).Methods(http.MethodGet)
	router.HandleFunc("/data", IGotMessage).Methods(http.MethodPost)
	router.HandleFunc("/headers", Header).Methods(http.MethodPost)
	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}

func Greater(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	param := params["PARAM"]
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(hello, param)))
}

func Bad(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func IGotMessage(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Empty body"))
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("error: %v", err)))
		return
	}

	if body == nil || len(body) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No body set"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(message, string(body))))
}

func Header(w http.ResponseWriter, r *http.Request) {
	strA := r.Header.Get("a")
	strB := r.Header.Get("b")
	if strA == "" || strB == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No headers set Response expected to have"))
		return
	}
	a, errA := strconv.Atoi(strA)
	if errA != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("%v", errA)))
		return
	}

	b, errB := strconv.Atoi(strB)
	if errB != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("%v", errB)))
		return
	}

	sum := strconv.Itoa(a + b)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("a+b", sum)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("a+b:%s", sum)))
}
