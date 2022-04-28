package main

import (
	"fmt"
	"net/http"
	"strings"
	"text/template"
)

func main() {
	http.HandleFunc("/", Sayhello)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9056", nil)
	if err != nil {
		panic(err)
	}

}

func Sayhello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Welcome to login Form")

}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		fmt.Println("method:", r.Method)
		t, _ := template.ParseFiles("login.html")
		t.Execute(w, nil)
	} else {

		r.ParseForm()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}

}
