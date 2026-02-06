package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	// config port+addr
	var (
		defaultAddr = "0.0.0.0" //all interface for docker future
		defaultPort = "17013"
	)

	addr := defaultAddr
	port := defaultPort

	// port override for env also for docker but fix in future
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}
	listen := addr + ":" + port

	// http handler basic implementation - only get allowed
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, err := w.Write([]byte("hola mundo"))
		if err != nil {
			log.Printf("Response write error: %v", err)
		}
	})

	http.Handle("/", handler)

	//inint server
	log.Printf("server started on http://%s ...", listen)
	if err := http.ListenAndServe(listen, nil); err != nil {
		log.Fatalf("ListenAndServe: %v", err)
	}
}
