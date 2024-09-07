package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "index", nil)
}

func create(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/create.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "create", nil)
}

func save_article(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	anons := r.FormValue("anons")
	full_text := r.FormValue("full_text")
	connStr := "host='localhost' port=5432 user='postgres2' password='David2004' dbname='ddd' sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	insert, err := db.Query("INSERT INTO articles (title, anons, full_text) VALUES($1, $2, $3)", title, anons, full_text)
	if err != nil {
		log.Fatal(err.Error())
	}
	// insert = insert
	defer insert.Close()
	// fmt.Println(insert.RowsAffected())
	// fmt.Println(insert.RowsAffected())
	// sel, err := db.Query(fmt.Sprintf("SELECT title FROM articles"))
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	// defer insert.Close()
	// fmt.Println(sel)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func HandleFunc() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/create", create)
	http.HandleFunc("/save_article", save_article)
	http.ListenAndServe(":8085", nil)
}

func main() {
	// connStr := "user=postgres2 password=David2004 dbname=ddd sslmode=disable"
	// db, err := sql.Open("postgres", connStr)
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	// defer db.Close()

	// insert, err := db.Query(fmt.Sprintf("CREATE TABLE articles (title VARCHAR(200) NOT NULL default 'No title',anons VARCHAR(200) NOT NULL default 'No anons', full_text VARCHAR(500) NOT NULL default 'No text')"))
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	// defer insert.Close()

	HandleFunc()
}
