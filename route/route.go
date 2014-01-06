package route

import (
	//"../email"
	//"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//var emails = email.NewEmailManager()

//const PathPrefix = "/emails/"

var chttp = http.NewServeMux()

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		log.Print("got a post")
		fmt.Fprintf(w, "Hi there, I love %s!", r.PostFormValue("email"))
	} else {
		log.Print("invalid method used to access /new")
		http.Redirect(w, r, "/", 302)
	}
}

func RegisterHandlers() {
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/new", Register)
}
