package product

import (
	"fmt"

	"sistemas-a3/database"
)


func AddProduct(prod *product) error{
	db, error := database.Connect()
	if error != nil {
		fmt.Println("Error while connecting to the database!")
		return error
	}
	defer db.Close()

	statement, error := db.Prepare("INSERT INTO produto (nome, valor) VALUES (?, ?)")
	if error != nil {
		fmt.Println("Error preparing statement!: ",error)
		return error
	}
	defer statement.Close()

    insert, error := statement.Exec(prod.Name, prod.Value)
	if error != nil {
		fmt.Println("Error while executing statement!: ",error)
		return error
	}

	_ = insert

	return nil
}

func AllProducts() ([]product, error){
	db, error := database.Connect()
	if error != nil {
		fmt.Println("Error while connecting to the database!")
		return nil, error
	}
	defer db.Close()

	lines, error := db.Query("SELECT * FROM produto")
	if error != nil {
		fmt.Println("Error while getting products: ",error)
		return nil, error
	}
	defer lines.Close()

	var products []product
	for lines.Next() {
		var prod product

		if error := lines.Scan(&prod.ID, &prod.Name, &prod.Value); error != nil {
			fmt.Println("Error while scanning product: ",error)
			return nil, error
		}

		products = append(products, prod)
	}

	return products, nil
}

func OneProduct(id uint64) (*product, error){
	db, error := database.Connect()
	if error != nil {
		fmt.Println("Error while connecting to the database!")
		return nil, error
	}
	defer db.Close()

	line, error := db.Query("SELECT * FROM produto WHERE id = ?", id)
	if error != nil {
		fmt.Println("Error while getting product: ",error)
		return nil, error
	}
	defer line.Close()

	var prod product
	if line.Next() {
		if error := line.Scan(&prod.ID, &prod.Name, &prod.Value); error != nil {
			fmt.Println("Error while scanning product: ",error)
			return nil, error
		}
	}

	return &prod, nil
}

func UpdateProduct(id uint64, prod *product) error{
	db, error := database.Connect()
	if error != nil {
		fmt.Println("Error while connecting to the database: ",error)
		return error
	}
	defer db.Close()

	statement, error := db.Prepare("UPDATE produto SET nome = ?, valor = ? WHERE id = ?")
	if error != nil {
		fmt.Println("Error while creating statement: ",error)
		return error
	}
	defer statement.Close()

	if _, error := statement.Exec(prod.Name, prod.Value, id); error != nil {
		fmt.Println("Error while executing statement: ",error)
		return error
	}

	return nil
}

func DeleteProduct(id uint64) error {
	db, error := database.Connect()
	if error != nil {
		fmt.Println("Error while connecting to the database: ",error)
		return error
	}
	defer db.Close()

	statement, error := db.Prepare("DELETE FROM produto WHERE id = ?")
	if error != nil {
		fmt.Println("Error while preparing statement: ",error)
		return error
	}
	defer statement.Close()

	if _, error := statement.Exec(id); error != nil {
		fmt.Println("Error while deleting product: ",error)
		return error
	}

	return nil
}
