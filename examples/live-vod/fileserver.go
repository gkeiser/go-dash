package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)
	log.Fatal(http.ListenAndServe(":8000", RequestLogger(fs)))
}

func RequestLogger(targetMux http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		targetMux.ServeHTTP(w, r)

		// log request by who(IP address)
		requesterIP := r.RemoteAddr
		// for k, v := range r.Header {
		// 	log.Printf("header key %s val %s", k, v)
		// }
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
