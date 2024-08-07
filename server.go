package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"

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

	if err != nil {
		fmt.Println("Invalid countries")
		return
	}
	e.GET("/", func(c echo.Context) error {
		return Render(c, http.StatusOK, ContentPage(CountriesListComponent(countries[:10])))
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
