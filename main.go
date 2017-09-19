package main

import (
	"fmt"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/jinzhu/gorm"
	"github.com/tengis617/oshopmicro/Product"
)

func main() {
	const addr = "postgresql://maxroach@localhost:26257/bank?sslmode=disable"
	db := setupDB(addr)
	svc := Product.NewService(db)

	createProductHandler := httptransport.NewServer()
}

func setupDB(addr string) *gorm.DB {
	db, err := gorm.Open("postgres", addr)
	if err != nil {
		panic(fmt.Sprintf("failed to connect: %v", err))
	}
	db.AutoMigrate(&Product.Product{})
	return db
}
