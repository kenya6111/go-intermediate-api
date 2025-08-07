package middlewares

import (
	"fmt"
	"log"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// リクエスト情報をロギング
		fmt.Println("!!!!!!!!!!!!!!")
		log.Println(req.RequestURI, req.Method, "asasa")
		next.ServeHTTP(w, req)
		fmt.Println("?????????????")
	})
}
