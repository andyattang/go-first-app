package web

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)
type User struct{
	username string
	password string
}

func Main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("call / handler")
		fmt.Fprint(w, "hello")
	})
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		log.Println("call /login handler")
		if r.Method == "GET" {
			t, err := template.ParseFiles("src/web/hello.html")
			if err != nil {
				log.Println(err)
				fmt.Fprintln(w, "Err: "+err.Error())
				return
			}
			t.Execute(w, nil)
		} else {
			r.ParseForm()
			un := r.Form["username"]
			pwd := r.Form["password"]

			fmt.Printf("username:%v, password:%v\n", un, pwd)

			t, err := template.ParseFiles("src/web/userinfo.html")
			if err != nil {
				log.Println(err)
				fmt.Fprintln(w, "Err: "+err.Error())
				return
			}
			t.Execute(w, &User{username:un[0], password:pwd[0]})
		}

	})
	log.Println("Starting http server...")
	http.ListenAndServe(":88", nil)
}
