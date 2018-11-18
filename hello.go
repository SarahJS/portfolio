package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// log the request
		fmt.Fprintf(os.Stdout, "%s %s @ %v\n", r.Method, r.RequestURI, time.Now())

		// output our response
		fmt.Fprintf(w, "portfolio service is up\n")
	})
	http.ListenAndServe(":8998", nil)
}
