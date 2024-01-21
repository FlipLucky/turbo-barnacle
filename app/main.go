package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	serverPort := ":8080"
	db, err := sql.Open("mysql", "root:root@tcp(db:3306)/db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rootHandler := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("src/index.html"))
		tmpl.Execute(w, nil)
	}

	pingHandler := func(w http.ResponseWriter, r *http.Request) {
		err = db.Ping()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("pong")
	}

	fs := http.FileServer(http.Dir("src/style"))
	http.Handle("/src/style/", http.StripPrefix("/src/style", fs))
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/ping", pingHandler)
	log.Fatal(http.ListenAndServe(serverPort, nil))
}
