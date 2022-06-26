package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func nameParam(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	fmt.Fprintf(w, "Hello, %s!", v["PARAM"])
}

func bad(w http.ResponseWriter, r *http.Request) {
	// w.WriteHeader(http.StatusInternalServerError)
	status := http.StatusText(http.StatusInternalServerError)
	http.Error(w, status, http.StatusInternalServerError)
}

func data(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("data body read: %s", err.Error())
		http.Error(w, "can not read request body", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "I got message:\n%s", body)
}

func headers(w http.ResponseWriter, r *http.Request) {
	sum := 0
	for _, v := range []string{"a", "b"} {
		head := r.Header.Get(v)
		n, err := strconv.Atoi(head)
		if err != nil {
			log.Printf("header read: %s", err.Error())
			status := fmt.Sprintf("header \"%s\" is not a number %s", v, head)
			http.Error(w, status, http.StatusBadRequest)
			return
		}
		sum += n
	}
	w.Header().Add("a+b", strconv.Itoa(sum))
}
