package handlers

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
		articleId,err  := strconv.Atoi(mux.Vars(req)["id"])
		if err != nil{
			http.Error(w,"Invalid query parameter",http.StatusBadRequest)
			return
		}
		resString := fmt.Sprintf("ArticleNo.%d\n",articleId)
		io.WriteString(w, resString)
}

func PostNiceHandler(w http.ResponseWriter, req * http.Request){
		io.WriteString(w, "post nice!\n")
}

func PostCommentHandler(w http.ResponseWriter, req * http.Request){
		io.WriteString(w, "post comment!\n")
}
