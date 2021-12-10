package main

import (
	"sistemas-a3/input"
	"sistemas-a3/product"
	"sistemas-a3/raw_material"
	"sistemas-a3/report"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// Raw material routes
	e.POST("/raw_material", raw_material.CreateRawMaterial)
	e.GET("/raw_material", raw_material.GetRawMaterials)
	e.GET("/raw_material/:id", raw_material.GetRawMaterial)
	e.GET("/raw_material/by_product/:id", raw_material.GetRawMaterialsByProduct)
	e.PUT("/raw_material/:id", raw_material.PutRawMaterial)
	e.DELETE("/raw_material/:id", raw_material.RemoveRawMaterial)

	// Product routes
	e.POST("/product", product.CreateProduct)
	e.GET("/product", product.GetProducts)
	e.GET("/product/:id", product.GetProduct)
	e.PUT("/product/:id", product.PutProduct)
	e.DELETE("/product/:id", product.RemoveProduct)

	// Input relation routes
	e.POST("/input", input.CreateInput)
	e.PUT("/input", input.PutInput)
	e.DELETE("/input", input.RemoveInput)

	// Report routes
	e.GET("/report/all", report.GetReportAllProducts)
	e.GET("/report/valuable", report.GetReportOnlyValuable)

	// Start server
	e.Logger.Fatal(e.Start(":6060"))
}
