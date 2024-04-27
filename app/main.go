package main

import (
	"fmt"

	"github.com/FlipLucky/turbo-barnacle/internal/elements"
	"github.com/FlipLucky/turbo-barnacle/internal/elements/routes"
	"github.com/labstack/echo/v4"
)

type Data struct {
	Data elements.Page
}

func newData(data elements.Page) Data {
	return Data{
		Data: data,
	}
}

func main() {
	r := routes.NewRouter()
	r.Static("/src", "src")
	// create page --> create append child method
	// add section
	// add columns
	// fill fields
	// render

	section, err := elements.NewLayoutElement(elements.Section)
	if err != nil {
		fmt.Printf("Error creating layout element: %s", err)
	}

	p2 := elements.TextElement{
		ElementType: elements.Paragraph,
		Body:        "Hello Constant World!!",
		ClassName:   "red",
	}

	row, err := elements.NewLayoutElement(elements.Row)
	if err != nil {
		fmt.Printf("Error creating layout element: %s", err)
	}

	col, err := elements.NewLayoutElement(elements.Col)
	if err != nil {
		fmt.Printf("Error creating layout element: %s", err)
	}

	col.AppendChild(p2)
	row.AppendChild(col)
	section.AppendChild(row)

	page := elements.NewPage(
		[]elements.PageElement{
			section,
		},
	)
	data := newData(page)

	r.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", data)
	})

	r.Logger.Fatal(r.Start(":8080"))
}
