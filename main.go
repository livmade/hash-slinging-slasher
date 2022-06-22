package main

//Login HTML

import (
	"errors"
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

	http.HandleFunc("/", handleLandingPage)
	http.HandleFunc("/login", loginPage)

	fmt.Println("Starting server on port 3000...")
	http.ListenAndServe(":3000", nil)
}

func handleLandingPage(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("hash-slinging-session-id")
	if errors.Is(err, http.ErrNoCookie) {
		w.Header().Set("Location", "/login")
		w.WriteHeader(302)
		return
	}

	if cookie.Value != "super secret value" {
		// check session find user
	}

	w.WriteHeader(200)
	w.Write([]byte("cat"))
}

func loginPage(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		t.Execute(w, nil)
		return
		// server login webpage with form
	} else if r.Method == "POST" {
		// handle login form, session creation and redirect back on success with cookie
		// if not success show them the door.
		r.ParseForm()
	} else {
		// what ???? invalid method so show error message or the door.
		err=
	}
}
