package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// 3. struct for saving data
type Product struct {
	ID uint
	Nama string
	Kategori string
	Harga string
}

func main()  {
	// 1. connect to db
	connURI  := "postgresql://postgres:password@localhost:5432/postgres?sslmode=disable"
	db, err := sql.Open("pgx", connURI)
	if err != nil {
		fmt.Printf("Error connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close() // last. defer to close db

	// 2. ping to db
	err = db.Ping()
	if err != nil {
		fmt.Printf("Error connect to database: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Succesfully connect to database")

	// 4. creating table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS produk(
			id SERIAL PRIMARY KEY,
			nama VARCHAR(255),
			kategori VARCHAR(50),
			harga INT
		)
	`)
	if err != nil {
		fmt.Printf("Failed to create table: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Table successfully created")
	
	// 5. insert data
	_, err = db.Exec(`INSERT INTO produk (nama, kategori, harga) VALUES($1, $2, $3)`, "Kertas A4", "Kertas", 35000)
	if err != nil {
		fmt.Printf("Failed to insert data: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Data inserted successfully")


	// 6. read from db
	row := db.QueryRow(`SELECT id, nama, kategori, harga FROM produk WHERE id = $1`, 1)  // queryRow func only take 1 row
	if row == nil {
		fmt.Println("Failed to reading from db")
	}

	// 7. assign data that we write to var with our struct 
	var product1 Product 
	err = row.Scan(&product1.ID, &product1.Nama, &product1.Kategori, &product1.Harga)
	if err != nil {
		fmt.Println("Failed to catch data from db")
		os.Exit(1)
	}

	fmt.Println(product1)


	// 8. select a few data from db
	rows, err := db.Query(`SELECT id, nama, kategori, harga FROM produk`)
	if rows == nil || err != nil {
		fmt.Printf("Failed to retrieve data: %v\n", err)
		os.Exit(1)
	}
	
	var productSlice []Product
	for rows.Next() {
		var eachProduct Product
		err = rows.Scan(&eachProduct.ID, &eachProduct.Nama, &eachProduct.Kategori, &eachProduct.Harga)
		if err != nil {
			fmt.Printf("Failed to retrieve data: %v\n", err)
			os.Exit(1)
		}

		productSlice = append(productSlice, eachProduct)
	}

	fmt.Println(productSlice)


	// 9. update data from db
	_, err = db.Exec(`UPDATE produk SET nama = $1, kategori = $2, harga = $3 WHERE id=$4`, "Buku Cerita", "Buku", "100000", 1 )
	if err != nil {
		fmt.Printf("Failed to update data: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Succesfully updated the data")


	// 10. deleting data  from db
	_, err = db.Exec(`DELETE FROM produk WHERE id=$1`, 3)
	if err != nil {
		fmt.Printf("Failed to delete data: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Succesfully deleted the data")
}