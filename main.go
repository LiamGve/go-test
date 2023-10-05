package main

import "fmt"
import "rsc.io/quote"
import "net/http"

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, quote.Go(), r.URL.Path)
	})

	http.ListenAndServe(":80", nil)
}
