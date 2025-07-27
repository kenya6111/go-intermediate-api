package handlers

import (
	"fmt"
	"io"
	"net/http"
)


func HelloHandler (w http.ResponseWriter, req * http.Request){
	if req.Method == http.MethodGet {
		io.WriteString(w, "Hello, world!\n")
	}else{
		http.Error(w,"Invalid method",http.StatusMethodNotAllowed)
	}
}

func PostArticleHandler (w http.ResponseWriter, req * http.Request){
	if req.Method == http.MethodPost {
		io.WriteString(w, "post article!\n")
	}else{
		http.Error(w,"Invalid method",http.StatusMethodNotAllowed)
	}
}

func ArticleListHandler (w http.ResponseWriter, req * http.Request){
	if req.Method == http.MethodGet {
		// 通常通りにレスポンスを返す
		io.WriteString(w, "get article list!\n")
	}else{
		http.Error(w,"Invalid method",http.StatusMethodNotAllowed)
	}
}

func ArticleDetailHandler (w http.ResponseWriter, req * http.Request){
	if req.Method == http.MethodGet {
		// 通常通りにレスポンスを返す
		articleId := 1
		resString := fmt.Sprintf("ArticleNo.%d\n",articleId)
		io.WriteString(w, resString)
	}else{
		http.Error(w,"Invalid method",http.StatusMethodNotAllowed)
	}
	
}

func PostNiceHandler(w http.ResponseWriter, req * http.Request){
	if req.Method == http.MethodPost {
		// 通常通りにレスポンスを返す
		io.WriteString(w, "post nice!\n")
	}else{
		http.Error(w,"Invalid method",http.StatusMethodNotAllowed)
	}
}

func PostCommentHandler(w http.ResponseWriter, req * http.Request){
	if req.Method == http.MethodPost {
		// 通常通りにレスポンスを返す
		io.WriteString(w, "post comment!\n")
	}else{
		http.Error(w,"Invalid method",http.StatusMethodNotAllowed)
	}
}
