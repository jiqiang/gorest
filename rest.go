package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

// Staff model
type Staff struct {
	id        int
	firstName string
	lastName  string
	addressID int
	email     string
	storeID   int
	active    bool
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "host=0.0.0.0 user=postgres password=postgres port=5432 dbname=dvdrental sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
}

func staffsIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	rows, err := db.Query("SELECT staff_id, first_name, last_name, address_id, email, store_id, active FROM staff")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	defer rows.Close()

	staffs := make([]*Staff, 0)
	for rows.Next() {
		staff := new(Staff)
		err := rows.Scan(&staff.id, &staff.firstName, &staff.lastName, &staff.addressID, &staff.email, &staff.storeID, &staff.active)
		if err != nil {
			log.Fatal(err)
		}
		staffs = append(staffs, staff)
	}
	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, staff := range staffs {
		fmt.Fprintf(w, "%d, %s, %s, %d, %s, %d, %t\n", staff.id, staff.firstName, staff.lastName, staff.addressID, staff.email, staff.storeID, staff.active)
	}
}

func main() {
	router := httprouter.New()
	router.GET("/staffs", staffsIndex)
	log.Fatal(http.ListenAndServe(":8080", router))
}
