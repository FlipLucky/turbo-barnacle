package main

import (
	"github.com/FlipLucky/turbo-barnacle/internal/elements"
	"github.com/FlipLucky/turbo-barnacle/internal/routes"
)

func main() {
	r := routes.NewRouter()
	r.Static("/assets", "src")
	p := elements.CreateContentElement()
	page := elements.NewPage(
		"Homepage",
		"/",
		"index",
		[]elements.PageElementInterface{p},
	)
	routes.NewPageRoute(r, page)

	r.Logger.Fatal(r.Start(":8080"))
}
