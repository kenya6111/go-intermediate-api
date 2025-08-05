package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kenya6111/go-intermediate-api/api"
)

var (
	dbUser     = "docker"
	dbPassword = "docker"
	dbDatabase = "intermediateDB"
	dbConn     = fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?parseTime=true", dbUser,
		dbPassword, dbDatabase)
)

func main() {
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()
	r := api.NewRouter(db)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
