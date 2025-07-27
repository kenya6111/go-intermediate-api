package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main(){
	helloHandler := func(w http.ResponseWriter, req * http.Request){
		io.WriteString(w, "Hello, wold!\n")
	}

	http.HandleFunc("/",helloHandler)

	log.Println("server start at port 8080")
	fmt.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080",nil))
}
