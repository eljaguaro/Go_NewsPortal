package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

type Article struct {
	Id                     uint16
	Title, Anons, FullText string
}

var posts = []Article{}
var showPost = Article{}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	db, err := sql.Open("postgres", fmt.Sprintf("host='localhost' port=5432 user='postgres2' password='David2004' dbname='ddd' sslmode=disable"))
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	res, err := db.Query("SELECT * FROM articles")
	if err != nil {
		log.Fatal(err.Error())
	}
	posts = []Article{}

	for res.Next() {
		var post Article
		err = res.Scan(&post.Id, &post.Title, &post.Anons, &post.FullText)
		if err != nil {
			log.Fatal(err.Error())
		}
		posts = append(posts, post)
	}
	t.ExecuteTemplate(w, "index", posts)
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

	if title == "" || anons == "" || full_text == "" {
		fmt.Fprintf(w, "Не все данные заполнены")
	} else {

		db, err := sql.Open("postgres", fmt.Sprintf("host='localhost' port=5432 user='postgres2' password='David2004' dbname='ddd' sslmode=disable"))
		if err != nil {
			log.Fatal(err.Error())
		}
		defer db.Close()

		insert, err := db.Query("INSERT INTO articles (title, anons, full_text) VALUES($1, $2, $3)", title, anons, full_text)
		if err != nil {
			log.Fatal(err.Error())
		}
		defer insert.Close()
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func show_post(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	t, err := template.ParseFiles("templates/show.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	db, err := sql.Open("postgres", fmt.Sprintf("host='localhost' port=5432 user='postgres2' password='David2004' dbname='ddd' sslmode=disable"))
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	res, err := db.Query("SELECT * FROM articles WHERE id = 1")
	if err != nil {
		log.Fatal(err.Error())
	}
	showPost = Article{}

	for res.Next() {
		var post Article
		err = res.Scan(&post.Id, &post.Title, &post.Anons, &post.FullText)
		if err != nil {
			log.Fatal(err.Error())
		}
		showPost = post
	}
	t.ExecuteTemplate(w, "show", showPost)
}

func HandleFunc() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/", index).Methods("GET")
	rtr.HandleFunc("/create", create).Methods("GET")
	rtr.HandleFunc("/save_article", save_article).Methods("POST")
	rtr.HandleFunc("/post/{id:[0-9]+}", show_post).Methods("GET")

	http.Handle("/", rtr)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
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
