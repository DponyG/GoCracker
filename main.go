package main

import (
	"html/template"
	"net/http"
)

//Our only page
type IndexPage struct {
	Title string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	p := IndexPage{Title: "Password Cracker"}
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8000", nil)
}
