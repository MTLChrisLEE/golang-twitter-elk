package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
)

type Data struct {
	l *log.Logger
}

func NewData(l *log.Logger) *Data {
	return &Data{l}
}

func (h *Data) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(rw, "Bad request", http.StatusNotFound)
		return
	}

	d, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(rw, "Oops", http.StatusBadRequest)
		return
	}

	h.l.Printf("Data %s\n", d)
}
