package main

func createDiv(children []PageElement, className string) PageElement {
	return ElementWithChildren{
		ElementType:   "div",
		ChildElements: children,
		ClassName:     className,
	}
}

type sectionConfig struct {
	Children  []PageElement
	ClassName string
}

func createSection(c sectionConfig) PageElement {
	return ElementWithChildren{
		ElementType:   "section",
		ChildElements: c.Children,
		ClassName:     c.ClassName,
	}
}

type paragraphConfig struct {
	Body      string
	ClassName string
}

func createParagraph(c paragraphConfig) PageElement {
	return ElementWithContent{
		ElementType: "p",
		Body:        c.Body,
		ClassName:   c.ClassName,
	}
}

func createHeader1(body string, className string) PageElement {
	return ElementWithContent{
		ElementType: "h1",
		Body:        body,
		ClassName:   className,
	}
}
