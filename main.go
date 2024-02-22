package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/0", func(w http.ResponseWriter, r *http.Request) { io.WriteString(w, "0") })
	http.HandleFunc("/1", func(w http.ResponseWriter, r *http.Request) { io.WriteString(w, "1") })
	http.HandleFunc("/2", func(w http.ResponseWriter, r *http.Request) { io.WriteString(w, "2") })

	http.ListenAndServe(":3333", nil)
}
