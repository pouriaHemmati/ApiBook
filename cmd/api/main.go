package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

type config struct {
	port int
	evn  string
}

type application struct {
	config config
	log    *log.Logger
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "server port")
	flag.StringVar(&cfg.evn, "env", "dec", "environment (dev|stage|prod)")
	flag.Parse()

	logger := log.New(os.Stdin, "", log.Ldate|log.Ltime)

	app := &application{config: cfg, log: logger}

	addr := fmt.Sprintf(":%d", cfg.port)

	ser := &http.Server{
		Addr:         addr,
		Handler:      app.route(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting %s server on %s", cfg.evn, addr)
	err := ser.ListenAndServe()
	logger.Fatal(err)
}
