package main

import (
	"github.com/FlipLucky/turbo-barnacle/internal/elements"
	"github.com/FlipLucky/turbo-barnacle/internal/routes"
)

func main() {
	r := routes.NewRouter()
	r.Static("/assets", "src")
	p := elements.CreateContentElement(elements.BodyText)
	p.Content = "henk"
	p.ClassNames = append(p.ClassNames, "body-text")
	tb := elements.CreateBlockElement(elements.TextBlock)
	tb.ClassNames = append(tb.ClassNames, "textblock")
	tb.AddChildElements([]elements.PageElementInterface{p})
	page := elements.NewPage(
		"Homepage",
		"/",
		"index",
		[]elements.PageElementInterface{tb},
	)
	routes.NewPageRoute(r, page)

	r.Logger.Fatal(r.Start(":8080"))
}
