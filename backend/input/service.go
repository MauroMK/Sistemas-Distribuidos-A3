package input

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type input struct {
	Product    uint64 `json:"product"`
	Material  uint64 `json:"material"`
	Quantity uint64 `json:"quantity"`
}

func CreateInput(c echo.Context) error {
    inp := new(input)

 	if err := c.Bind(inp); err != nil {
		return err
	}

    if err := AddInput(inp); err != nil {
    	return c.JSON(http.StatusInternalServerError, "Error while adding input relation!")
    }

	return c.JSON(http.StatusCreated, "Input relation sucessfully added!")
}

func PutInput(c echo.Context) error{
    inp := new(input)

 	if err := c.Bind(inp); err != nil {
		return err
	}

    if err := UpdateInput(inp); err != nil {
    	return c.JSON(http.StatusInternalServerError, "Error while updating input relation!")
    }

    return c.JSON(http.StatusOK, "Input relation sucessfully updated!")
}

func RemoveInput(c echo.Context) error{
    inp := new(input)

 	if err := c.Bind(inp); err != nil {
		return err
	}

    if err := DeleteInput(inp); err != nil {
    	return c.JSON(http.StatusInternalServerError, "Error while removing input relation!")
    }

    return c.JSON(http.StatusOK, "Input relation sucessfully deleted!")
}
