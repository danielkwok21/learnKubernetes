package main

import (
	"fmt"
	"io"
	"net/http"
)

func echoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL.String(), r.Proto)
	fmt.Fprintf(w, "Host: %s\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr: %s\n\n", r.RemoteAddr)

	fmt.Fprintln(w, "Headers:")
	for name, values := range r.Header {
		for _, v := range values {
			fmt.Fprintf(w, "%s: %s\n", name, v)
		}
	}
	fmt.Fprintln(w, "")

	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintln(w, "Error reading body:", err)
		return
	}

	if len(body) == 0 {
		fmt.Fprintln(w, "(no body)")
	} else {
		fmt.Fprintln(w, "Body:")
		w.Write(body)
		fmt.Fprintln(w, "")
	}
}

func main() {
	http.HandleFunc("/", echoHandler)
	addr := ":8080"
	fmt.Println("Starting server on", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		fmt.Println("Server error:", err)
	}
}
