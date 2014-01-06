package main

import (
	"../route"
	"net/http"
)

func main() {
	route.RegisterHandlers()
	http.ListenAndServe(":8080", nil)
}
