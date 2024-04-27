package routes

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Router echo.Echo

func NewRouter() *echo.Echo {
	router := echo.New()
	router.Renderer = newTemplate()
	router.Static("/src", "src")
	return router
}

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
