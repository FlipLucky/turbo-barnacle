package main

import (
	"log"

	"github.com/FlipLucky/turbo-barnacle/internal/db"
	"github.com/FlipLucky/turbo-barnacle/internal/elements"
	"github.com/FlipLucky/turbo-barnacle/internal/routes"
)

func main() {
	r := routes.NewRouter()
	r.Static("/assets", "src")

	db, err := db.OpenDb()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.TestConnection()
	if err != nil {
		log.Fatal(err)
	}

	results, err := db.Query("SELECT * FROM pages")
	if err != nil {
		log.Fatal(err)
	}
	pageCollection := []elements.Page{}

	for results.Next() {
		var page elements.Page

		err = results.Scan(&page.Id, &page.Title, &page.Slug, &page.Template)
		if err != nil {
			log.Fatal(err)
		}
		pageCollection = append(pageCollection, page)
	}

	bodyText := elements.CreateContentElement(elements.BodyText)
	bodyText.ClassNames = append(bodyText.ClassNames, "body-text")
	bodyText.Content = "The Body of Henk"

	subheaderText := elements.CreateContentElement(elements.SubHeaderText)
	subheaderText.ClassNames = append(subheaderText.ClassNames, "subheader-text")
	subheaderText.Content = "The Neck of Henk"

	headerText := elements.CreateContentElement(elements.HeaderText)
	headerText.ClassNames = append(headerText.ClassNames, "header-text")
	headerText.Content = "The Head of Henk"

	tb := elements.CreateBlockElement(elements.TextBlock)
	tb.ClassNames = append(tb.ClassNames, "textblock")
	tb.AddChildElements([]elements.PageElementInterface{headerText, subheaderText, bodyText})

	page := elements.NewPage(
		"Homepage",
		"/",
		"index",
		[]elements.PageElementInterface{tb},
	)

	for _, page = range pageCollection {
		routes.NewPageRoute(r, page)
	}
	routes.NewPageRoute(r, page)

	r.Logger.Fatal(r.Start(":8080"))
}
