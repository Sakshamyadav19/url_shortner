package main

import (
	"fmt"
	"net/http"
	"url-shortner/handler"
)

func main(){
	http.HandleFunc("/shorten",handler.ShortenUrl)
	http.HandleFunc("/redirect/",handler.RedirectHandler)

	err:=http.ListenAndServe(":8080",nil)

	if err!=nil{
		fmt.Println("Error while running server",err)
	}

}