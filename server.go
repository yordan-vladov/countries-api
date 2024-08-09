package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

func CountryCodes(c []Country) map[string]*Country {
	m := make(map[string]*Country)

	for _, country := range c {
		m[country.Alpha3Code] = &country
	}

	return m
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	countries, err := Countries()
	countryCodes := CountryCodes(countries)

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

		return Render(c, http.StatusOK, ContentPage(CountriesListComponent(countries[start:end], 2, -1, "")))
	})

	e.POST("/countries", func(c echo.Context) error {
		size := 12
		// get page
		page, err := strconv.Atoi(c.FormValue("page"))
		if err != nil {
			page = 1
		}

		// get region and set region filter
		regionFilter := func(c Country) bool {
			return true
		}

		region, err := strconv.Atoi(c.FormValue("region"))

		if err != nil {
			region = -1
		}
		if region != -1 {
			regionFilter = func(c Country) bool {
				return c.Region == regions[region]
			}
		}

		searchFilter := func(c Country) bool {
			return true
		}
		search := c.FormValue("search")

		if search != "" {
			searchFilter = func(c Country) bool {
				return strings.Contains(strings.ToLower(c.Name), strings.ToLower(search))
			}
		}
		// get filtered countries
		filteredCountries := Filter(countries, regionFilter, searchFilter)

		start := (page - 1) * size
		if start > len(filteredCountries) {
			return c.NoContent(http.StatusNotFound)
		}

		end := start + size

		if end > len(filteredCountries) {
			end = len(filteredCountries)
		}

		return Render(c, http.StatusOK, LoadedCountries(filteredCountries[start:end], page+1, region, search))
	})

	e.GET("/country/:code", func(c echo.Context) error {

		var country *Country

		code := c.Param("code")

		country, ok := countryCodes[code]

		if !ok {
			country = &countries[0]
		}
		var borders []Border

		for _, borderCode := range country.Borders {
			borders = append(borders, Border{Code: borderCode, Name: countryCodes[borderCode].Name})
		}
		return Render(c, http.StatusOK, ContentPage(CountryPage(*country, borders)))
	})

	e.Static("/styles", "public/styles")
	e.Static("/images", "public/images")
	e.Static("/scripts", "public/scripts")
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
