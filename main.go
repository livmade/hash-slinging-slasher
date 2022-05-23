package main

//Test HTML

import (
	"html/template"
	"net/http"
	//"log"
	//"fmt"
	//"os"
)

var t *template.Template

type usrID struct {
	Name string
}

func main() {
	var err error
	t, err = template.ParseFiles("login.gohtml")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", loginPage)
	http.ListenAndServe(":3000", nil)
}

func loginPage(w http.ResponseWriter, r *http.Request) {

	uN := usrID{"Olivia Smith"}
	err := t.Execute(w, uN)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
