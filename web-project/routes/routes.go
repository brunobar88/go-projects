package routes

import (
	"net/http"
	"personal-projects/web-project/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.NewProduct)
	http.HandleFunc("/insert", controllers.InsertProduct)
	http.HandleFunc("/delete", controllers.DeleteProduct)
	http.HandleFunc("/edit", controllers.EditProduct)
	http.HandleFunc("/update", controllers.UpdateProduct)
}
