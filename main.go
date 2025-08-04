package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/kenya6111/go-intermediate-api/controllers"
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
	// // dbUser := "docker"
	// // dbPassword := "docker"
	// // dbDatabase := "intermediateDB"
	// // dbConn := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	// articleID := 1
	// const sqlStr = `select * from articles where article_id = ?;`
	// row := db.QueryRow(sqlStr, articleID)
	// if err := row.Err(); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// var article models.Article
	// var createdTime sql.NullTime

	// err = row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// if createdTime.Valid {
	// 	article.CreatedAt = createdTime.Time
	// }
	// fmt.Printf("%+v\n", article)

	// articleForInsert := models.Article{
	// 	Title:    "insert test",
	// 	Contents: "Can I insert data correctly?",
	// 	UserName: "saki",
	// }
	// const sqlStrForInsert = `
	// insert into articles (title, contents, username, nice, created_at) values
	// (?, ?, ?, 0, now());
	// `
	// result, err := db.Exec(sqlStrForInsert, articleForInsert.Title, articleForInsert.Contents, articleForInsert.UserName)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// // 結果を確認
	// fmt.Println(result.LastInsertId())
	// fmt.Println(result.RowsAffected())

	// // トランザクションの開始
	// tx, err := db.Begin()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// // 現在のいいね数を取得するクエリを実行する
	// article_id := 1
	// const sqlGetNice = `select nice from articles where article_id = ?`

	// rowNice := tx.QueryRow(sqlGetNice, article_id)
	// if err := rowNice.Err(); err != nil {
	// 	fmt.Println(err)
	// 	tx.Rollback()
	// 	return
	// }
	// // 変数 nicenum に現在のいいね数を読み込む
	// var nicenum int
	// err = rowNice.Scan(&nicenum)
	// if err != nil {
	// 	fmt.Println(err)
	// 	tx.Rollback()
	// 	return
	// }
	// // いいね数を+1 する更新処理を行う
	// const sqlUpdateNice = `update articles set nice = ? where article_id = ?`
	// _, err = tx.Exec(sqlUpdateNice, nicenum+1, article_id)
	// if err != nil {
	// 	fmt.Println(err)
	// 	tx.Rollback()
	// 	return
	// }
	// // コミットして処理内容を確定させる
	// tx.Commit()

	ser := services.NewMyAppService(db)
	con := controllers.NewMyAppController(ser)

	r := mux.NewRouter()
	r.HandleFunc("/hello", con.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/article", con.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", con.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]}", con.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", con.PostNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", con.PostCommentHandler).Methods(http.MethodPost)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
