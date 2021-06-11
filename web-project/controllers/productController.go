package controllers

import (
	"net/http"
	"personal-projects/web-project/models"
	"strconv"
	"text/template"
)

const redirectCode int = 301

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.GetAllProducts()

	templates.ExecuteTemplate(w, "Index", products)
}

func NewProduct(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "New-product", nil)
}

func InsertProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price, _ := strconv.ParseFloat(r.FormValue("price"), 64)
		amount, _ := strconv.Atoi(r.FormValue("amount"))

		models.CreateNewProduct(name, description, price, amount)
	}

	http.Redirect(w, r, "/", redirectCode)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")

	models.DeleteProduct(productId)

	http.Redirect(w, r, "/", redirectCode)
}

func EditProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")

	product := models.GetProduct(productId)

	templates.ExecuteTemplate(w, "Edit", product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id, _ := strconv.Atoi(r.FormValue("id"))
		name := r.FormValue("name")
		description := r.FormValue("description")
		price, _ := strconv.ParseFloat(r.FormValue("price"), 64)
		amount, _ := strconv.Atoi(r.FormValue("amount"))

		models.UpdateProduct(id, name, description, price, amount)
	}

	http.Redirect(w, r, "/", 301)
}
