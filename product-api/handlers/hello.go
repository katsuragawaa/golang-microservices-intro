package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	logger *log.Logger
}

// NewHello will inject logger
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	h.logger.Println("hellu~")

	data, err := ioutil.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, "Something went wront", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(writer, "Data: %s", data)
}
