package main

type PageElement interface {
	GetType() string
}

type ElementWithContent struct {
	ElementType string
	Body        string
}

func (e ElementWithContent) GetType() string {
	return e.ElementType
}

func NewPageElementWithContent(elementType, body string) ElementWithContent {
	return ElementWithContent{
		ElementType: elementType,
		Body:        body,
	}
}

type ElementWithChildren struct {
	ElementType   string
	ChildElements []PageElement
}

func (e ElementWithChildren) GetType() string {
	return e.ElementType
}

func NewPageElementWithChildren(elementType string, childElements []PageElement) ElementWithChildren {
	return ElementWithChildren{
		ElementType:   elementType,
		ChildElements: childElements,
	}
}

type Page struct {
	Elements []PageElement
}

func NewPage(pageElements []PageElement) Page {
	return Page{
		Elements: pageElements,
	}
}
