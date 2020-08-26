package main

import (
	"go-helloworld-api/homepage"
	"go-helloworld-api/server"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	CertFile      = os.Getenv("CertFile")
	CertKey       = os.Getenv("CertKeyFile")
	ServerAddress = os.Getenv("ServerAddress")
)

var logger = log.New(os.Stdout, "go-helloworld-api", log.LstdFlags|log.Lshortfile)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/home", Logger(homepage.HomeHandler))
	srv := server.New(mux, ServerAddress)

	logger.Println("Server Starting")
	err := srv.ListenAndServeTLS(CertFile, CertKey)
	if err != nil {
		logger.Fatalf("server failed to start: %v", err)
	}
}

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		startTime := time.Now()
		next(writer, request)
		logger.Printf("Request processed in %s\n", time.Now().Sub(startTime))
	}
}
