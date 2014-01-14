package main

import (
	"github.com/jjhageman/launch-rock/route"
	"net/http"
	"os"
)

func main() {
	route.RegisterHandlers()
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}
