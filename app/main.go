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

	secondRouteHandler := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("src/templates/page-2.html"))
		tmpl.Execute(w, nil)
	}

	pingHandler := func(w http.ResponseWriter, r *http.Request) {
		err = db.Ping()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("pong")
	}

	// styling
	styleFileHandler := http.FileServer(http.Dir("src/style"))
	http.Handle("/src/style/", http.StripPrefix("/src/style", styleFileHandler))
	// js if needed
	jsFileHandler := http.FileServer(http.Dir("src/js"))
	http.Handle("src/js.", http.StripPrefix("/src/js", jsFileHandler))

	// routes
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/page-2", secondRouteHandler)
	http.HandleFunc("/ping", pingHandler)

	// serving the router
	log.Fatal(http.ListenAndServe(serverPort, nil))
}
