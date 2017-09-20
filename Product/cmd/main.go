package main

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/tengis617/oshopmicro/Product"
)

func main() {
	const addr = "postgresql://oshopadmin@localhost:26257/oshop?sslmode=disable"
	db := setupDB(addr)
	defer db.Close()

	s := Product.NewService(db)
	h := Product.MakeHTTPHandler(s)

	http.ListenAndServe(":8080", h)
}

func setupDB(addr string) *gorm.DB {
	db, err := gorm.Open("postgres", addr)
	if err != nil {
		panic(fmt.Sprintf("failed to connect: %v", err))
	}
	db.AutoMigrate(&Product.Product{})
	return db
}
