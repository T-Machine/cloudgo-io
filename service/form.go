package service

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func Submit(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	username := r.Form["username"][0]
	password := r.Form["password"][0]
	if len(username) == 0 || len(password) == 0 {
		log.Fatal("err")
		return
	}
	t := template.Must(template.ParseFiles("templates/table.html"))
	data := map[string]string{
		"Name": username,
		"Pass": password,
	}
	if err := t.Execute(w, data); err != nil {
		log.Fatal(err)
	}
}