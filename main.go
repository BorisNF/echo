package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	port := os.Getenv("PORT")
	fmt.Printf("Staring listening at port %s", port)

	err := http.ListenAndServe(":"+port, http.HandlerFunc(echo))
	if err != nil {
		fmt.Printf("Server exited with error: %s", err)
		os.Exit(1)
	}
}

func echo(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimLeft(r.URL.Path, "/")
	words := strings.Split(path, "/")
	resp := strings.Join(words, " ")
	fmt.Fprint(w, resp)
}
