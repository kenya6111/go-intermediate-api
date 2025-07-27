package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kenya6111/go-intermediate-api/handlers"
)

func main(){
	r := mux.NewRouter()
	r.HandleFunc("/hello",handlers.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/article",handlers.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list",handlers.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/1",handlers.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice",handlers.PostNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment",handlers.PostCommentHandler).Methods(http.MethodPost)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080",r))
	// http.ListenAndServe 関数の第二引数というのは、実は「サーバーの中で使うルータを指定する」部分なのです
}
