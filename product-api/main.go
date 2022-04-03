package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc(
		"/",
		func(rw http.ResponseWriter, r *http.Request) {
			log.Println("Hey")
			d, err := ioutil.ReadAll(r.Body)

			if err != nil {
				http.Error(
					rw,
					"some error",
					http.StatusBadRequest,
				)

				// Same as above
				// rw.WriteHeader(http.StatusBadRequest)
				// rw.Write([]byte("Some error"))

				return
			}

			fmt.Fprintf(rw, "Your data %s", d)
		},
	)

	http.HandleFunc(
		"/goodbye",
		func(http.ResponseWriter, *http.Request) {
			log.Println("Bye")
		},
	)

	http.ListenAndServe(":9090", nil)
}
