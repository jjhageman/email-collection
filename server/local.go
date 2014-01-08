package main

import (
	//"github.com/jjhageman/launch-rock/db"
	"github.com/jjhageman/launch-rock/route"
	"net/http"
)

func main() {
	route.RegisterHandlers()
	http.ListenAndServe(":8080", nil)
}
