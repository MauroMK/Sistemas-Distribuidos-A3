package main

import (
	"sistemas-a3/backend-reports/report"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Routes
	e.GET("/report/all", report.GetReportAllProducts)
	e.GET("/report/valuable", report.GetReportOnlyValuable)

	// Start server
	e.Logger.Fatal(e.Start(":9001"))
}
