package rest

import (
	"fmt"
	"net/http"

	"goji.io"
	"goji.io/pat"
)

const local = "localhost:8000"

func StarHttpService() {
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/ping"), ping)

	http.ListenAndServe(local, mux)
}

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ping")
	fmt.Fprintf(w, "pong\n")
}
