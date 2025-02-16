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

//  gorm many2many
type Penduduk struct {
	ID uint `gorm:"primaryKey;autoincrement;type:serial;column:id"`
	Alamat []Alamat `gorm:"many2many:penduduk_alamat"`
}

type Alamat struct {
	ID uint `gorm:"primaryKey;autoincrement;type:serial;column:id"`
	AlamatLengkap string `gorm:"columng:alamat_lengkap"`
}

func main()  {
	// 1. connect to db with gorm
	connURI := "postgresql://postgres:password@localhost:5432/database?sslmode=disable"
	db, err := gorm.Open(postgres.Open(connURI), &gorm.Config{
		SkipDefaultTransaction: true,  // add this to do transaction
	})
	if err != nil {
		fmt.Printf("Failed connect to database: %v\n", err)
		os.Exit(1)
	}
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	fmt.Println("Database successfully connected")

	db.AutoMigrate(&Produk{})
	fmt.Println("Table has been made")

	// // 2. adding data using create 
	// produk := Produk{Nama: "Stik PS", Kategori: "Perabot", Harga: 70098}
	// result := db.Create(&produk)
	// if result.Error != nil {
	// 	fmt.Printf("Failed to create table: %v\n", result.Error)
	// 	os.Exit(1)
	// }
	// fmt.Println("Successfully adding data to table")
	
	// // 3. select data
	// var selectedProduk Produk
	// result = db.First(&selectedProduk, 2)
	// if result.Error != nil {
	// 	fmt.Printf("Failed to select data: %v\n", result.Error)
	// 	os.Exit(1)
	// }
	// fmt.Println(selectedProduk)


	// // 4. select data more than 1
	// var produkSlice []Produk
	// result = db.Find(&produkSlice, []uint{1,4}) // select a few data with index
	// if result.Error != nil {
	// 	fmt.Printf("Failed to catch a few data : %v \n", result.Error)
	// 	os.Exit(1)
	// }
	// fmt.Println(produkSlice)

	// // 4. another form for search, using where
	// result = db.Where(map[string]interface{}{"id":6}).Find(&produkSlice)
	// if result.Error != nil {
	// 	fmt.Printf("Failed to catch a few data : %v \n", result.Error)
	// 	os.Exit(1)
	// }
	// fmt.Println(produkSlice)


	// // 5. update data
	// result = db.Model(&Produk{ID:1}).Updates(&Produk{Nama: "Samsung a5", Kategori: "Smartphone", Harga: 120000000})
	// if result.Error != nil {
	// 	fmt.Printf("Failed to catch a few data : %v \n", result.Error)
	// 	os.Exit(1)
	// }


	// // 6. delete from db
	// result = db.Delete(Produk{ID:1})
	// if result.Error != nil {
	// 	fmt.Printf("Failed to delete data : %v \n", result.Error)
	// 	os.Exit(1)
	// }

	// // 7. transaction 
	// db.Transaction(func(tx *gorm.DB) error {
	// 	result = tx.Delete(&Produk{ID: 1})
	// 	if result.Error != nil {
	// 		fmt.Printf("Failed do the transaction : %v \n", result.Error)
	// 		return result.Error	// return error disini akan auto me return rollback	
	// 	} 
	// 	fmt.Println("Transaction success")
	// 	return nil
	// })

	// adding many2many relation data to table
	penduduk := Penduduk{
		Alamat: []Alamat{
			{
				AlamatLengkap: "Kota Probolinggo",
			},
			{
				AlamatLengkap: "Kota Malang",
			},
		},
	}

	db.Create(&penduduk)

	// var penduduk Penduduk
	// db.Preload("Alamat").First(&penduduk, 1)
	// fmt.Println(penduduk)


}