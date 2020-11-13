package main

import (
	"fmt"
	"github.com/codihuston/gorilla-mux-http/api/v1/products"
	// "os"
)

func main() {
	// a := App{}

	p := products.Product{ID: 1}
	// p := api.v1.Product{ID: 1}

	// a.Initialize(
	// 	os.Getenv("APP_DB_USERNAME"),
	// 	os.Getenv("APP_DB_PASSWORD"),
	// 	os.Getenv("APP_DB_NAME"),
	// )

	// a.Run(":8010")

	fmt.Println(p)
}
