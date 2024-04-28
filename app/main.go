package main

import (
	"fmt"

	"github.com/FlipLucky/turbo-barnacle/internal/elements"
	"github.com/FlipLucky/turbo-barnacle/internal/routes"
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
	r.Static("/assets", "src")

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

	col2, err := elements.NewLayoutElement(elements.Col)
	if err != nil {
		fmt.Printf("Error creating layout element: %s", err)
	}

	img := elements.NewImageElement(
		"/assets/images/synthwave.jpg",
		"synthwave",
	)

	col.AppendChild(p2)
	col2.AppendChild(img)
	row.AppendChild(col)
	row.AppendChild(col2)
	section.AppendChild(row)

	page := elements.NewPage(
		"Homepage",
		"/",
		"index",
		[]elements.PageElement{
			section,
		},
	)
	routes.NewPageRoute(r, page)

	r.Logger.Fatal(r.Start(":8080"))
}
