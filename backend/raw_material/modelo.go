package raw_material

import (
	"fmt"

	"sistemas-a3/database"
)


func AddRawMaterial(raw *raw_material) error{
	db, error := database.Connect()
	if error != nil {
		fmt.Println("Error while connecting to the database!")
		return error
	}
	defer db.Close()

	statement, error := db.Prepare("INSERT INTO materia_prima (nome, estoque) VALUES (?, ?)")
	if error != nil {
		fmt.Println("Error preparing statement!: ",error)
		return error
	}
	defer statement.Close()

    insert, error := statement.Exec(raw.Name, raw.Inventory)
	if error != nil {
		fmt.Println("Error while executing statement!: ",error)
		return error
	}

	_ = insert

	return nil
}

func AllRawMaterials() ([]raw_material, error){
	db, error := database.Connect()
	if error != nil {
		fmt.Println("Error while connecting to the database!")
		return nil, error
	}
	defer db.Close()

	lines, error := db.Query("SELECT * FROM materia_prima")
	if error != nil {
		fmt.Println("Error while getting raw materials: ",error)
		return nil, error
	}
	defer lines.Close()

	var raw_materials []raw_material
	for lines.Next() {
		var raw raw_material

		if error := lines.Scan(&raw.ID, &raw.Name, &raw.Inventory); error != nil {
			fmt.Println("Error while scanning raw material: ",error)
			return nil, error
		}

		raw_materials = append(raw_materials, raw)
	}

	return raw_materials, nil
}

func AllRawMaterialsByProduct(id_produto uint64) ([]raw_material_input, error){
	db, error := database.Connect()
	if error != nil {
		fmt.Println("Error while connecting to the database!")
		return nil, error
	}
	defer db.Close()

	lines, error := db.Query("SELECT m.id, m.nome, m.estoque, i.quantidade FROM materia_prima m, insumo i WHERE m.id=i.id_materia_prima AND i.id_produto = ?",id_produto)
	if error != nil {
		fmt.Println("Error while getting raw materials by product: ",error)
		return nil, error
	}
	defer lines.Close()

	var raw_materials []raw_material_input
	for lines.Next() {
		var raw raw_material_input

		if error := lines.Scan(&raw.ID, &raw.Name, &raw.Inventory, &raw.Quantity); error != nil {
			fmt.Println("Error while scanning raw material by product: ",error)
			return nil, error
		}

		raw_materials = append(raw_materials, raw)
	}

	return raw_materials, nil
}

func OneRawMaterial(id uint64) (*raw_material, error){
	db, error := database.Connect()
	if error != nil {
		fmt.Println("Error while connecting to the database!")
		return nil, error
	}
	defer db.Close()

	line, error := db.Query("SELECT * FROM materia_prima WHERE id = ?", id)
	if error != nil {
		fmt.Println("Error while getting raw material: ",error)
		return nil, error
	}
	defer line.Close()

	var raw raw_material
	if line.Next() {
		if error := line.Scan(&raw.ID, &raw.Name, &raw.Inventory); error != nil {
			fmt.Println("Error while scanning raw material: ",error)
			return nil, error
		}
	}

	return &raw, nil
}

func UpdateRawMaterial(id uint64, raw *raw_material) error{
	db, error := database.Connect()
	if error != nil {
		fmt.Println("Error while connecting to the database: ",error)
		return error
	}
	defer db.Close()

	statement, error := db.Prepare("UPDATE materia_prima SET nome = ?, estoque = ? WHERE id = ?")
	if error != nil {
		fmt.Println("Error while creating statement: ",error)
		return error
	}
	defer statement.Close()

	if _, error := statement.Exec(raw.Name, raw.Inventory, id); error != nil {
		fmt.Println("Error while executing statement: ",error)
		return error
	}

	return nil
}

func DeleteRawMaterial(id uint64) error {
	db, error := database.Connect()
	if error != nil {
		fmt.Println("Error while connecting to the database: ",error)
		return error
	}
	defer db.Close()

	statement, error := db.Prepare("DELETE FROM materia_prima WHERE id = ?")
	if error != nil {
		fmt.Println("Error while preparing statement: ",error)
		return error
	}
	defer statement.Close()

	if _, error := statement.Exec(id); error != nil {
		fmt.Println("Error while deleting raw material: ",error)
		return error
	}

	return nil
}
