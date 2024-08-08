package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	. "yvladov.com/countries-api/components"
	. "yvladov.com/countries-api/models"
)

type Template struct {
	templates *template.Template
}

func Countries() ([]Country, error) {
	jsonFile, err := os.Open("data.json")

	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var countries []Country

	json.Unmarshal(byteValue, &countries)

	return countries, nil
}

func main() {
	e := echo.New()
	countries, err := Countries()
	regions := []string{"Africa", "Americas", "Asia", "Europe", "Oceania"}

	if err != nil {
		fmt.Println("Invalid countries")
		return
	}
	e.GET("/", func(c echo.Context) error {
		size := 12
		page := 1

		start := (page - 1) * size

		end := start + size

		return Render(c, http.StatusOK, ContentPage(CountriesListComponent(countries[start:end], 2, -1)))
	})

	e.GET("/countries", func(c echo.Context) error {
		size := 12
		page, err := strconv.Atoi(c.QueryParam("page"))
		if err != nil {
			page = 1
		}

		filter := func(c Country) bool {
			return true
		}

		region, err := strconv.Atoi(c.QueryParam("region"))

		if err == nil {
			filter = func(c Country) bool {
				return c.Region == regions[region]
			}
		}
		filteredCountries := Filter(countries, filter)

		start := (page - 1) * size
		if start > len(filteredCountries) {
			return c.NoContent(http.StatusNotFound)
		}

		end := start + size

		if end > len(filteredCountries) {
			end = len(filteredCountries)
		}

		return Render(c, http.StatusOK, LoadedCountries(Filter(countries, filter)[start:end], page+1, region))
	})

	e.Static("/styles", "public/styles")
	e.Static("/images", "public/images")
	e.Logger.Fatal(e.Start(":1323"))
}

func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := t.Render(ctx.Request().Context(), buf); err != nil {
		return err
	}

	return ctx.HTML(statusCode, buf.String())
}
