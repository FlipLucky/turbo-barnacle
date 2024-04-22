package main

type PageElement interface {
	GetType() string
	GetClassName() string
}

type ElementWithContent struct {
	ElementType string
	Body        string
	ClassName   string
}

func (e ElementWithContent) GetType() string {
	return e.ElementType
}

func (e ElementWithContent) GetClassName() string {
	return e.ClassName
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
	ClassName     string
}

func (e ElementWithChildren) GetType() string {
	return e.ElementType
}

func (e ElementWithChildren) GetClassName() string {
	return e.ClassName
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
