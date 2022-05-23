package main

//Login HTML

import (
	"fmt"
	"html/template"
	"net/http"
	//"log"
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
	http.HandleFunc("/login", loginPage)
	http.ListenAndServe(":3000", nil)

}

func loginPage(w http.ResponseWriter, r *http.Request) {
	uN := usrID{"username"}
	r.ParseForm()
	// logic part of log in
	fmt.Println("username:", r.Form["username"])
	fmt.Println("password:", r.Form["password"])

	err := t.Execute(w, uN)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	
	t.Execute(w, struct{ Success bool }{true})
}

func userInput() {

} //This takes in user input
