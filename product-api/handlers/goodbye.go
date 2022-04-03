package handlers

import (
	"log"
	"net/http"
)

type Goodbye struct {
	logger *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (goodbye *Goodbye) ServeHTTP(writer http.ResponseWriter, _ *http.Request) {
	writer.Write([]byte("Bye!"))
}
