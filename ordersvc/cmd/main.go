package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/tengis617/oshopmicro/ordersvc"
)

var (
	addr = "postgresql://root@localhost:26257/oshop?sslmode=disable"
)

func main() {
	db := setupDB(addr)
	ordersvc.NewOrderService(db)
}

func setupDB(addr string) *gorm.DB {
	db, err := gorm.Open("postgres", addr)

	if err != nil {
		log.Fatal("couldnt connect to database")
	}

	db.AutoMigrate(&ordersvc.Order{})

	return db
}
