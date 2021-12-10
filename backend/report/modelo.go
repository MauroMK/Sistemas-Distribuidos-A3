package report

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"sistemas-a3/database"
)



type ProductReport struct {
	ID       uint64
	Name     string
	Quantity uint64
}

type MaterialRest struct {
	RawMaterialId uint64
	Name          string
	Quantity      uint64
}

type raw_material struct {
	ID        uint64 `json:"id"`
	Name      string `json:"name"`
	Inventory uint64 `json:"inventory"`
}

type productWithInput struct {
	ID         uint64  `json:"id"`
	Name       string  `json:"name"`
	Value      float64 `json:"value"`
	InputsId   string  `json:"inputs_id"`
	InputsName string  `json:"inputs_name"`
	InputsQty  string  `json:"inputs_qty"`
}

func GetFullReport(all_products bool) ([]ProductReport, error) {
	db, error := database.Connect()
	if error != nil {
		fmt.Println("Error while connecting to the database!")
		return nil, error
	}
	defer db.Close()

	products, err := GetAllProductsWithInputs(db)
	if err != nil {
		fmt.Println("Error while getting products: ", err)
		return nil, err
	}

	rawMaterials, err := GetAllRawMaterials(db)
	if err != nil {
		fmt.Println("Error while getting products: ", err)
		return nil, err
	}

	fmt.Println("Products: ", products)
	fmt.Println("Raw Materials: ", rawMaterials)

	var products_to_be_produced []ProductReport
	for _, product := range products {
		var inventoryBackup[]uint64
		for _, material := range rawMaterials {
			inventoryBackup = append(inventoryBackup, material.Inventory)
		}
		// Just do the magic
		inputsIdsList := []uint64{}
		for _, id := range strings.Split(product.InputsId, ",") {
			parsedId, err := strconv.ParseUint(id, 10, 64)
			if err != nil {
				fmt.Println("Error converting ID")
				continue
			}
			inputsIdsList = append(inputsIdsList, parsedId)
		}

		inputsNames := []string{}
		for _, name := range strings.Split(product.InputsName, ",") {
			inputsNames = append(inputsNames, name)
		}

		inputsQtyList := []uint64{}
		for _, qty := range strings.Split(product.InputsQty, ",") {
			parsedQty, err := strconv.ParseUint(qty, 10, 64)
			if err != nil {
				fmt.Println("Error converting qty")
				continue
			}
			inputsQtyList = append(inputsQtyList, parsedQty)
		}

		quantityByInputId := map[uint64]uint64{}
		for index := range inputsIdsList {
			for index2, material := range rawMaterials {
				if(material.ID == inputsIdsList[index]){
					if material.Inventory >= inputsQtyList[index] {
						rest := material.Inventory % inputsQtyList[index]
						quantityOfItems := (material.Inventory - rest) / inputsQtyList[index]
						quantityByInputId[inputsIdsList[index]] = quantityOfItems
						if(!all_products){
							rawMaterials[index2].Inventory = rest
						}
					} else {
						quantityByInputId[inputsIdsList[index]] = 0
					}
				}
			}
		}

		minorQuantity := uint64(9999999)
		hasMaterialsEnough := true
		for _, quantityByInputId := range quantityByInputId {
			if quantityByInputId <= 0 {
				hasMaterialsEnough = false
			} else {
				if quantityByInputId < minorQuantity {
					minorQuantity = quantityByInputId
				}
			}
		}

		if hasMaterialsEnough {
			products_to_be_produced = append(products_to_be_produced, ProductReport{
				ID:       product.ID,
				Name:     product.Name,
				Quantity: minorQuantity,
			})
		} else {
			fmt.Println("Vai habilitar o backup!", inventoryBackup)
			//rawMaterials = rawMaterialsBackup
			for index := range rawMaterials {
				rawMaterials[index].Inventory = inventoryBackup[index]
			}
		}
	}

	fmt.Println("These products will be produced: ", products_to_be_produced)
	return products_to_be_produced, nil
}

func GetAllProductsWithInputs(db *sql.DB) ([]productWithInput, error) {
	lines, error := db.Query("select p.id, p.nome, p.valor, GROUP_CONCAT(i.id_materia_prima) materias_prima, GROUP_CONCAT(mp.nome) materia_prima_nome, GROUP_CONCAT(i.quantidade) quantidades from produto p left join insumo i on (p.id = i.id_produto) left join materia_prima mp on (i.id_materia_prima = mp.id) group by p.id order by p.valor desc")
	if error != nil {
		fmt.Println("Error while getting product: ", error)
		return nil, error
	}
	defer lines.Close()

	var products []productWithInput

	for lines.Next() {
		var prod productWithInput

		if error := lines.Scan(&prod.ID, &prod.Name, &prod.Value, &prod.InputsId, &prod.InputsName, &prod.InputsQty); error != nil {
			fmt.Println("Error while scanning product: ", error)
			return nil, error
		}

		products = append(products, prod)
	}

	return products, nil
}

func GetAllRawMaterials(db *sql.DB) ([]raw_material, error) {
	lines, error := db.Query("SELECT * FROM materia_prima")
	if error != nil {
		fmt.Println("Error while getting raw materials: ", error)
		return nil, error
	}
	defer lines.Close()

	var raw_materials []raw_material
	for lines.Next() {
		var raw raw_material

		if error := lines.Scan(&raw.ID, &raw.Name, &raw.Inventory); error != nil {
			fmt.Println("Error while scanning raw material: ", error)
			return nil, error
		}

		raw_materials = append(raw_materials, raw)
	}

	return raw_materials, nil
}
