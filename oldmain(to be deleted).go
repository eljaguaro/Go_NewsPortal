// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"

// 	_ "github.com/go-sql-driver/mysql"
// )

// type User struct {
// 	Name string `json: "name"`
// 	Age  uint16 `json: "age"`
// }

// func main() {
// 	fmt.Println("Работа с MySQL")
// 	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/")
// 	if err != nil {
// 		log.Fatal(err.Error)
// 	}
// 	defer db.Close()

// 	// insert, err := db.Query("INSERT INTO `users` (`name`, `age`) VALUES('ALEX', 35)")
// 	// if err != nil {
// 	// 	log.Fatal(err.Error)
// 	// }
// 	// defer insert.Close()

// 	res, err := db.Query("SELECT `name`, `age`,  FROM `users`")
// 	if err != nil {
// 		log.Fatal(err.Error)
// 	}
// 	for res.Next() {
// 		var user User
// 		err = res.Scan(&user.Name, &user.Age)
// 	}
// 	fmt.Println("Подключено к MySQL")
// }
