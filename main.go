package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/0", func(w http.ResponseWriter, r *http.Request) { io.WriteString(w, "-1") })
	http.HandleFunc("/1", func(w http.ResponseWriter, r *http.Request) { io.WriteString(w, "-1") })
	http.HandleFunc("/2", func(w http.ResponseWriter, r *http.Request) { io.WriteString(w, "-1") })

	http.HandleFunc("/3", func(w http.ResponseWriter, r *http.Request) { io.WriteString(w, "-1") })
	http.HandleFunc("/4", func(w http.ResponseWriter, r *http.Request) { io.WriteString(w, "-1") })
	http.HandleFunc("/5", func(w http.ResponseWriter, r *http.Request) { io.WriteString(w, "-1") })

	http.ListenAndServe(":3333", nil)
}
Nonsense
