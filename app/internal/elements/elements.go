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
	SetParentUUID(ElementUUID)
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

func (e *PageElement) SetParentUUID(uuid ElementUUID) {
	e.ParentUUID = uuid
}

func (e PageElement) GetAttributes() map[string]string {
	return e.Attributes
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
