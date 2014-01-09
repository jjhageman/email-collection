package route

import (
	"github.com/coopernurse/gorp"
	"github.com/jjhageman/launch-rock/db"
	"github.com/jjhageman/launch-rock/email"
	//"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//var emails = email.NewEmailManager()

//const PathPrefix = "/emails/"
var DB_CONN = "user=jjhageman dbname=devstatus sslmode=disable"
var dbmap *gorp.DbMap

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		log.Print("got a post")
		address := r.PostFormValue("email")
		email, err1 := email.NewEmail(address)
		if err1 != nil {
			log.Println("Error:", err1)
			fmt.Fprintf(w, "Sorry, something went wrong %q", err1)
		} else {
			err2 := dbmap.Insert(email)
			if err2 != nil {
				log.Println("Error:", err2)
				fmt.Fprintf(w, "Sorry, something went wrong %q", err2)

			} else {
				log.Println("New email inserted:", address)
				fmt.Fprintf(w, "Hi there, I love %s!", address)
			}

		}
	} else {
		log.Print("invalid method used to access /new")
		http.Redirect(w, r, "/", 302)
	}
}

func RegisterHandlers() {
	// initialize the DbMap
	dbmap = db.InitDb(DB_CONN)
	//defer dbmap.Db.Close()

	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/new", Register)
}
