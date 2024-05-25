package routes

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewRouter() *echo.Echo {
	router := echo.New()
	router.Renderer = newTemplate()
	return router
}

func NewPageRoute(r *echo.Echo, p PageRoute) {
	r.GET(p.GetSlug(), func(c echo.Context) error {
		return c.Render(http.StatusOK, p.GetTemplate(), p)
	})
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
