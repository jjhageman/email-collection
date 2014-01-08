package route

import (
	"github.com/jjhageman/launch-rock/db"
	//"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//var emails = email.NewEmailManager()

//const PathPrefix = "/emails/"
var DB_CONN = "user=jjhageman dbname=devstatus sslmode=disable"

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
	// initialize the DbMap
	var dbmap = db.InitDb(DB_CONN)
	defer dbmap.Db.Close()

	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/new", Register)
}
