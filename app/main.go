package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	err := t.templates.ExecuteTemplate(w, name, data)
	if err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return nil
}

func newTemplate() *Template {
	t := template.Must(template.ParseGlob("src/templates/*.html"))

	return &Template{
		templates: t,
	}
}

type Data struct {
	Data Page
}

func newData(data Page) Data {
	return Data{
		Data: data,
	}
}

func main() {
	router := echo.New()
	router.Renderer = newTemplate()
	router.Static("/src", "src")

	p := createParagraph(paragraphConfig{Body: "Hello, World!", ClassName: "red"})

	section := createSection(sectionConfig{Children: []PageElement{p}})

	page := NewPage(
		[]PageElement{
			section,
		},
	)
	data := newData(page)

	router.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", data)
	})

	router.Logger.Fatal(router.Start(":8080"))
}
