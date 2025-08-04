package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kenya6111/go-intermediate-api/controllers"
	"github.com/kenya6111/go-intermediate-api/routers"
	"github.com/kenya6111/go-intermediate-api/services"
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
		fmt.Println(err)
	}
	defer db.Close()

	ser := services.NewMyAppService(db)
	con := controllers.NewMyAppController(ser)

	r := routers.NewRouter(con)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
