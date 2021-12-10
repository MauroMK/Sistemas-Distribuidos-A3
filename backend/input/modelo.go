package input

import (
	"fmt"

	"sistemas-a3/database"
)


func AddInput(inp *input) error{
	db, error := database.Connect()
	if error != nil {
		fmt.Println("Error while connecting to the database!")
		return error
	}
	defer db.Close()

	statement, error := db.Prepare("INSERT INTO insumo (id_produto, id_materia_prima, quantidade) VALUES (?, ?, ?)")
	if error != nil {
		fmt.Println("Error while preparing statement!: ",error)
		return error
	}
	defer statement.Close()

    insert, error := statement.Exec(inp.Product, inp.Material, inp.Quantity)
	if error != nil {
		fmt.Println("Error while executing statement!: ",error)
		return error
	}

	_ = insert

	return nil
}

func UpdateInput(inp *input) error{
	db, error := database.Connect()
	if error != nil {
		fmt.Println("Error while connecting to the database: ",error)
		return error
	}
	defer db.Close()

	statement, error := db.Prepare("UPDATE insumo SET id_produto = ?, id_materia_prima = ?, quantidade = ? WHERE id_produto = ? AND id_materia_prima = ?")
	if error != nil {
		fmt.Println("Error while creating statement: ",error)
		return error
	}
	defer statement.Close()

	if _, error := statement.Exec(inp.Product, inp.Material, inp.Quantity, inp.Product, inp.Material); error != nil {
		fmt.Println("Error while executing statement: ",error)
		return error
	}

	return nil
}

func DeleteInput(inp *input) error {
	db, error := database.Connect()
	if error != nil {
		fmt.Println("Error while connecting to the database: ",error)
		return error
	}
	defer db.Close()

	statement, error := db.Prepare("DELETE FROM insumo WHERE id_produto = ? AND id_materia_prima = ?")
	if error != nil {
		fmt.Println("Error while preparing statement: ",error)
		return error
	}
	defer statement.Close()

	if _, error := statement.Exec(inp.Product, inp.Material); error != nil {
		fmt.Println("Error while deleting input relation: ",error)
		return error
	}

	return nil
}
