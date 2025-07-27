package handlers

import (
	"fmt"
	"io"
	"net/http"
)


func HelloHandler (w http.ResponseWriter, req * http.Request){
		io.WriteString(w, "Hello, world!\n")
}

func PostArticleHandler (w http.ResponseWriter, req * http.Request){
		io.WriteString(w, "post article!\n")
}

func ArticleListHandler (w http.ResponseWriter, req * http.Request){
		io.WriteString(w, "get article list!\n")
}

func ArticleDetailHandler (w http.ResponseWriter, req * http.Request){
		articleId := 1
		resString := fmt.Sprintf("ArticleNo.%d\n",articleId)
		io.WriteString(w, resString)
}

func PostNiceHandler(w http.ResponseWriter, req * http.Request){
		io.WriteString(w, "post nice!\n")
}

func PostCommentHandler(w http.ResponseWriter, req * http.Request){
		io.WriteString(w, "post comment!\n")
}
