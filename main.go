package main

import (
	"log"
	"net/http"
	"os"
)

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, os.Getenv("REDIRECT_TARGET"), http.StatusMovedPermanently)
}

func main() {
	addr, ok := os.LookupEnv("REDIRECT_SOURCE")
	if !ok {
		addr = ":80"
	}

	target, ok := os.LookupEnv("REDIRECT_TARGET")
	if !ok {
		log.Fatal("missing REDIRECT_TARGET env var")
	}

	log.Println("redirecting from " + addr + " to " + target)

	if err := http.ListenAndServe(addr, http.HandlerFunc(redirect)); err != nil {
		log.Fatal(err)
	}
}
