package report

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetReportAllProducts(c echo.Context) error {
	report, err := GetFullReport(true)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error getting report")
	}

	return c.JSON(http.StatusOK, report)
}

func GetReportOnlyValuable(c echo.Context) error {
	report, err := GetFullReport(false)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error getting report")
	}

	return c.JSON(http.StatusOK, report)
}
