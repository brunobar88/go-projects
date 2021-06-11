package main

import (
	"net/http"
	"personal-projects/web-project/routes"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
