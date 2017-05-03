package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

var db *sql.DB

// Index controller
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var firstName string
	err := db.QueryRow("SELECT first_name FROM customer WHERE customer_id = $1", 1).Scan(&firstName)
	if err != nil {
		fmt.Fprintf(w, "Database error\n")
	} else {
		fmt.Fprintf(w, "Welcome %s!\n", firstName)
	}
}

// Hello controller
func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func main() {
	var err error
	db, err = sql.Open("postgres", "host=0.0.0.0 user=postgres password=postgres port=5432 dbname=dvdrental sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxIdleConns(100)

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	log.Fatal(http.ListenAndServe(":8080", router))
}
