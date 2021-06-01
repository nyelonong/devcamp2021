package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

var db *sql.DB

func GetDB() *sql.DB {
	var err error

	if db == nil {
		connStr := "user=postgres password=docker dbname=devcamp sslmode=disable"
		db, err = sql.Open("postgres", connStr)
		if err != nil {
			panic(err)
		}
	}

	return db
}

func main() {
	http.HandleFunc("/", handleHomePage)
	http.HandleFunc("/shop", handleFirstShop)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleHomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, the site is running :)")
}

func handleFirstShop(w http.ResponseWriter, r *http.Request) {
	firstShop, err := GetFirstShop()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, firstShop)
}

func GetFirstShop() (string, error) {
	var shop string

	db := GetDB()
	err := db.QueryRow("SELECT name FROM SHOP LIMIT 1").Scan(&shop)
	if err != nil {
		log.Println(err)
	}
	return shop, err
}
