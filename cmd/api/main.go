package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "0.0.1"

type config struct {
	port int
	env  string
}

type AppStatus struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

type application struct {
	config           config
	logger           *log.Logger
	dsn              string
	DB               *sql.DB
	maxFileSize      int
	allowedFileTypes []string
}

func main() {

	var cfg config

	flag.IntVar(&cfg.port, "port", 8080, "server port to listen on")
	flag.StringVar(&cfg.env, "env", "development", "application environment (development|production)")
	flag.StringVar(&cfg.env, "dsn", "", "database connection string")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{
		config:           cfg,
		logger:           logger,
		maxFileSize:      256 << 20,
		allowedFileTypes: []string{"audio/wave", "application/ogg"},
	}

	db, err := sql.Open("mysql", app.dsn)

	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	app.DB = db

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Println("Starting Server On Port: ", cfg.port)

	err = srv.ListenAndServe()

	if err != nil {
		log.Println(err)
	}

}
