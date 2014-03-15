package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)


var dataDir = flag.String("datadir", "/var/lib/curlbin", "Directory when data reside")
var logFilePath = flag.String("logfile", "/var/log/curlbin.log", "Log file to use")
var listenAddr = flag.String("listen", ":8080", "Address to listen on")
var name = flag.String("server-name", "", "Server name (with port) to use in paste urls.")

func main() {
	flag.Parse()

	// open log first
	if *logFilePath == "-" {
		// use stdout
		log.SetOutput(os.Stdout)
	} else {
		logFile, err := os.OpenFile(*logFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			panic(err)
		}
		defer logFile.Close()

		log.SetOutput(logFile)
	}

	if _, err := os.Stat(*dataDir); os.IsNotExist(err) {
		err = os.MkdirAll(*dataDir, 0777)
		if err != nil {
			log.Fatal("failed to create datadir", err)
		}
	}

	storage := NewStorage(*dataDir)
	server := NewServer(storage, *name, *listenAddr)

	err := http.ListenAndServe(*listenAddr, server)
	if err != nil {
		log.Fatal("server error", err)
	}
}
