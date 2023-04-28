package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	http.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		db, err := sql.Open("mysql", "root:Siva@2662@tcp(127.0.0.1:3306)/mydatabase")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		rows, err := db.Query("SHOW DATABASES")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var databases []string
		for rows.Next() {
			var database string
			err := rows.Scan(&database)
			if err != nil {
				log.Fatal(err)
			}
			databases = append(databases, database)
		}
		// _, err = db.Exec("CREATE DATABASE hand")
		// if err != nil {
		// 	log.Fatal(err)
		// }

		fmt.Fprintf(w, "Databases: %v", databases)
		fmt.Println("Database created successfully")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
