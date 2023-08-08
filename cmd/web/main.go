package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// for application-wide dependencies
type application struct {
	errorLog *log.Logger
	infoLog *log.Logger
}

func main() {

	// flag for server start
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// for logging info msgs
	infoLog := log.New(os.Stdout, "INFO\t", log.Ltime)

	// for logging error msgs
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// init a new instance of application struct
	app := &application{
		errorLog: errorLog,
		infoLog: infoLog,
	}

	srv := &http.Server {
		Addr: *addr,
		ErrorLog: errorLog,
		Handler: app.routes(),
	}

	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}