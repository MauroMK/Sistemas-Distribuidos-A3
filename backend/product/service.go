package product

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type product struct {
	ID    uint64 `json:"id"`
	Name  string `json:"name"`
	Value float64 `json:"value"`
}

func CreateProduct(c echo.Context) error {
    prod := new(product)

 	if err := c.Bind(prod); err != nil {
		return err
	}

    if err := AddProduct(prod); err != nil {
    	return c.JSON(http.StatusInternalServerError, "Error while adding product!")
    }

	return c.JSON(http.StatusCreated, "Product sucessfully added!")
}

func GetProducts(c echo.Context) error{
	products, error := AllProducts()

	if error != nil{
		return c.JSON(http.StatusInternalServerError, "Error while geting products!")
	}

	return c.JSON(http.StatusOK, products)
}

func GetProduct(c echo.Context) error{
	ID, error := strconv.ParseUint(c.Param("id"), 10, 64)
	if error != nil {
		return c.JSON(http.StatusInternalServerError, "Error while converting id to integer!")
	}

	prod, error := OneProduct(ID)

	if error != nil {
	    return c.JSON(http.StatusInternalServerError, "Error while getting product!")	
	}

	return c.JSON(http.StatusOK, prod)
}

func PutProduct(c echo.Context) error{
	ID, error := strconv.ParseUint(c.Param("id"), 10, 64)
	if error != nil {
		return c.JSON(http.StatusInternalServerError, "Error while converting id to integer!")
	}

    prod := new(product)

 	if err := c.Bind(prod); err != nil {
		return err
	}

    if err := UpdateProduct(ID, prod); err != nil {
    	return c.JSON(http.StatusInternalServerError, "Error while updating product!")
    }

    return c.JSON(http.StatusOK, "Product sucessfully updated!")
}

func RemoveProduct(c echo.Context) error{
	ID, error := strconv.ParseUint(c.Param("id"), 10, 64)
	if error != nil {
		return c.JSON(http.StatusInternalServerError, "Error while converting id to integer!")
	}

    if err := DeleteProduct(ID); err != nil {
    	return c.JSON(http.StatusInternalServerError, "Error while removing product!")
    }

    return c.JSON(http.StatusOK, "Product sucessfully deleted!")
}
