package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)
	log.Fatal(http.ListenAndServe(":8000", RequestLogger(fs)))
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func RequestLogger(targetMux http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		setupResponse(&w, r)
		if (*r).Method == "OPTIONS" {
			return
		}
		start := time.Now()
		fmt.Printf("request started at %s - %+v \n", start, r.URL)
		targetMux.ServeHTTP(w, r)

		// log request by who(IP address)
		requesterIP := r.RemoteAddr
		for k, v := range r.Header {
			log.Printf("header key %s val %s", k, v)
		}
		log.Printf(
			"%s\t\t%s\t%s\t%s\t\t%v",
			r.Method,
			r.RequestURI,
			r.URL.Query(),
			requesterIP,
			time.Since(start),
		)
	})
}
