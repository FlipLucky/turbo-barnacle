package elements

import (
	"strings"

	"github.com/google/uuid"
)

type ElementUUID string

func createElementUUID() ElementUUID {
	return ElementUUID(uuid.NewString())
}

type PageElementInterface interface {
	GetType() string
	GetClassNames() string
	GetElementUUID() ElementUUID
	GetParentUUID() ElementUUID
	GetAttributes() map[string]string
}

type PageElement struct {
	Type        string
	ClassNames  []string
	ElementUUID ElementUUID
	ParentUUID  ElementUUID
	Attributes  map[string]string
}

func (e PageElement) GetType() string {
	return e.Type
}
func (e PageElement) GetClassNames() string {
	return strings.Join(e.ClassNames, " ")
}
func (e PageElement) GetElementUUID() ElementUUID {
	return e.ElementUUID
}
func (e PageElement) GetParentUUID() ElementUUID {
	return e.ParentUUID
}
func (e PageElement) GetAttributes() map[string]string {
	return e.Attributes
}

type StructureElementInterface interface {
	PageElement
	GetChildElements() []PageElement
}

type StructureElement struct {
	PageElement   PageElement
	ChildElements []PageElementInterface
}

func (e StructureElement) GetChildElements() []PageElementInterface {
	return e.ChildElements
}

type BlockElementInterface interface {
	PageElement
	GetChildElements() []PageElement
}

type BlockElement struct {
	PageElement
	ChildElements []PageElementInterface
}

func (e BlockElement) GetChildElements() []PageElementInterface {
	return e.ChildElements
}

type ContentElementInterface interface {
	PageElement
	GetContent() string
}

type ContentElement struct {
	PageElement
	Content string
}

func (e ContentElement) GetContent() string {
	return e.Content
}

func CreateContentElement() ContentElement {
	return ContentElement{
		PageElement: PageElement{
			Type:        "p",
			ClassNames:  []string{"red"},
			ElementUUID: createElementUUID(),
			Attributes:  map[string]string{},
		},
		Content: "Henk",
	}
}

// type ContentElementCollection struct {
// 	Collection map[ElementUUID]PageElement
// }
//
// func (c *ContentElementCollection) AddElement(e PageElement) {
// 	c.Collection[e.GetElementUUID()] = e
// }
//
// func CreateContentElementCollection() *ContentElementCollection {
// 	c := new(ContentElementCollection)
// 	c.Collection = make(map[ElementUUID]PageElement)
// 	return c
// }
//
// func ElementSeeder() {
// 	m := CreateContentElementCollection()
// 	p2 := CreateTextElement("p")
// 	m.AddElement(p2)
// }
//
// type ElementCollection map[ElementUUID]PageElement
//
// func (c ElementCollection) AddElement(e PageElement) {
// 	c[e.GetElementUUID()] = e
// }
