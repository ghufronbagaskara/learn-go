package main

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


type Produk struct {
	ID uint `gorm:"primaryKey;autoincrement;type:serial;column:id"`
	Nama string `gorm:"type:varchar(255);column:nama"`
	Kategori string `gorm:"type:varchar(50);column:kategori"`
	Harga int `gorm:"type:int;column:harga"`
}
func (Produk) TableName() string {
	return "produk"
}

func main()  {
	// 1. connect to db with gorm
	connURI := "postgresql://postgres:password@localhost:5432/database?sslmode=disable"
	db, err := gorm.Open(postgres.Open(connURI), &gorm.Config{})
	if err != nil {
		fmt.Printf("Failed connect to database: %v\n", err)
		os.Exit(1)
	}
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	fmt.Println("Database successfully connected")

	db.AutoMigrate(&Produk{})
	fmt.Println("Table has been made")


}