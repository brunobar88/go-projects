package models

import (
	"database/sql"
	"personal-projects/web-project/db"
)

type Product struct {
	Id, Amount        int
	Name, Description string
	Price             float64
}

func GetAllProducts() []Product {
	dbConnection := db.DBConnection()

	allProducts, err := dbConnection.Query("select * from products order by amount asc")

	if err != nil {
		panic(err.Error())
	}

	products := makeProductsResponse(allProducts)

	defer dbConnection.Close()
	return products
}

func makeProductsResponse(allProducts *sql.Rows) []Product {
	product := Product{}
	products := []Product{}

	for allProducts.Next() {
		var id, amount int
		var name, description string
		var price float64

		err := allProducts.Scan(&id, &name, &description, &price, &amount)

		if err != nil {
			panic(err.Error())
		}

		product.Id = id
		product.Name = name
		product.Description = description
		product.Price = price
		product.Amount = amount

		products = append(products, product)
	}

	return products
}

func CreateNewProduct(name, description string, price float64, amount int) {
	dbConnection := db.DBConnection()

	scriptInsert, err := dbConnection.Prepare("insert into products (name, description, price, amount) values($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	scriptInsert.Exec(name, description, price, amount)

	defer dbConnection.Close()
}

func DeleteProduct(id string) {
	dbConnection := db.DBConnection()

	scriptDelete, err := dbConnection.Prepare("delete from products where id = $1")

	if err != nil {
		panic(err.Error())
	}

	scriptDelete.Exec(id)
	defer dbConnection.Close()
}

func GetProduct(id string) Product {
	dbConnection := db.DBConnection()

	product, err := dbConnection.Query("select * from products where id = $1", id)
	if err != nil {
		panic(err.Error())
	}

	productForEdit := Product{}

	for product.Next() {
		var name, description string
		var price float64
		var id, amount int

		err := product.Scan(&id, &name, &description, &price, &amount)
		if err != nil {
			panic(err.Error())
		}

		productForEdit.Id = id
		productForEdit.Name = name
		productForEdit.Description = description
		productForEdit.Price = price
		productForEdit.Amount = amount
	}
	defer dbConnection.Close()
	return productForEdit
}

func UpdateProduct(id int, name, descriprion string, price float64, amount int) {
	dbConnection := db.DBConnection()

	queryUpdateProduct, err := dbConnection.Prepare("update products set name=$1, description=$2, price=$3, amount=$4 where id=$5")

	if err != nil {
		panic(err.Error())
	}

	queryUpdateProduct.Exec(name, descriprion, price, amount, id)
	defer dbConnection.Close()
}
