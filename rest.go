package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

// Index controller
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	email := testDb()
	fmt.Fprintf(w, "Welcome!\n%s", email)
}

// Hello controller
func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func testDb() string {
	db, err := sql.Open("postgres", "host=0.0.0.0 user=postgres password=postgres port=5432 dbname=dvdrental sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	customerID := 1
	rows, err := db.Query("SELECT email FROM customer WHERE customer_id = $1", customerID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var email string
		if err := rows.Scan(&email); err != nil {
			log.Fatal(err)
		}
		return email
	}
	return "error"
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	log.Fatal(http.ListenAndServe(":8080", router))
}
