package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Product struct {
	ID    int
	Name  string
	Price float64
}

func main() {
	// Connect to the Adminer database
	db, err := sql.Open("mysql", "root:ROOT@tcp(127.0.0.1:3306)/goLang")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the database")

	// Perform CRUD operations
	// Create
	createProduct(db, "Laptop", 999.99)

	// Read
	products, err := readProducts(db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Products:")
	for _, p := range products {
		fmt.Println(p)
	}

	// Update
	updateProduct(db, 1, "Gaming Laptop", 1299.99)

	// Delete
	deleteProduct(db, 2)
}

func createProduct(db *sql.DB, name string, price float64) {
	result, err := db.Exec("INSERT INTO products(name, price) VALUES(?, ?)", name, price)
	if err != nil {
		log.Fatal(err)
	}
	id, _ := result.LastInsertId()
	fmt.Printf("Product %s created with ID %d\n", name, id)
}

func readProducts(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}

func updateProduct(db *sql.DB, id int, name string, price float64) {
	_, err := db.Exec("UPDATE products SET name=?, price=? WHERE id=?", name, price, id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Product with ID %d updated\n", id)
}

func deleteProduct(db *sql.DB, id int) {
	_, err := db.Exec("DELETE FROM products WHERE id=?", id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Product with ID %d deleted\n", id)
}
