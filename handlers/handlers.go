package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kenya6111/go-intermediate-api/models"
)


func HelloHandler (w http.ResponseWriter, req * http.Request){
		io.WriteString(w, "Hello, world!\n")
}

func PostArticleHandler (w http.ResponseWriter, req * http.Request){
		io.WriteString(w, "post article!\n")
		
		length, err := strconv.Atoi(req.Header.Get("Content-Length"))// Go ではこの書き方はちょっと古い・微妙。以下のように io.ReadAll(req.Body) を使うのが今の主流：
		if err != nil {
			http.Error(w, "cannot get content length\n", http.StatusBadRequest)
			return
		}
		reqBodybuffer := make([]byte, length)
		
		if _,err := req.Body.Read(reqBodybuffer); ! errors.Is(err, io.EOF) {
			http.Error(w, "fail to get request body \n", http.StatusBadRequest)
			return
		}
		defer req.Body.Close()
		fmt.Println("⭐️")
		fmt.Println(reqBodybuffer)
		
		var reqArticle models.Article
		if err := json.Unmarshal(reqBodybuffer, &reqArticle); err != nil {//マーシャルするの逆、解体。json（バイトスライス）からgoに解体する
			http.Error(w, "fail to decode json\n", http.StatusBadRequest)
			return
		}
		fmt.Println("👹")
		fmt.Println(reqArticle)
		article :=reqArticle

		jsonData,err := json.Marshal(article) //整列させる→goをJsonに整列させる→マーシャルする（goからバイト配列）
		if err != nil{
			http.Error(w, "failed to encode json \n",http.StatusInternalServerError)
		}
		
		w.Write(jsonData)
}

func ArticleListHandler (w http.ResponseWriter, req * http.Request){
		queryMap := req.URL.Query()

		var page int
		if p,ok := queryMap["page"]; ok && len(p)>0{
			var err error
			page ,err = strconv.Atoi(p[0])
			if err != nil{
				http.Error(w, "Invalid query parameter",http.StatusBadRequest)
				return
			}
		} else {
			page = 1
		}

		articleList := []models.Article{models.Article1,models.Article2}
		jsonData, err := json.Marshal(articleList)
		if err != nil {
			errMsg := fmt.Sprintf("fail to encode json (page %d)\n", page)
			http.Error(w, errMsg, http.StatusInternalServerError)
			return
		}

		w.Write(jsonData)
}

func ArticleDetailHandler (w http.ResponseWriter, req * http.Request){
		articleID,err  := strconv.Atoi(mux.Vars(req)["id"])
		if err != nil{
			http.Error(w,"Invalid query parameter",http.StatusBadRequest)
			return
		}

		article := models.Article1
		jsonData, err := json.Marshal(article)
		if err != nil {
			errMsg := fmt.Sprintf("fail to encode json (articleID %d)\n", articleID)
			http.Error(w, errMsg, http.StatusInternalServerError)
			return
		}

		w.Write(jsonData)
}

func PostNiceHandler(w http.ResponseWriter, req * http.Request){
	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func PostCommentHandler(w http.ResponseWriter, req * http.Request){
	comment := models.Comment1
	jsonData, err := json.Marshal(comment)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}
