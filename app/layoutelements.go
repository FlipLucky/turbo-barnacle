package main

import "fmt"

type LayoutElementType string

func (l LayoutElementType) String() string {
	return string(l)
}

const (
	Container LayoutElementType = "container"
	Row       LayoutElementType = "row"
	Col       LayoutElementType = "col"
)

func validateLayoutElementType(t LayoutElementType) bool {
	switch t {
	case Container, Row, Col:
		return true
	default:
		return false
	}
}

type LayoutElement struct {
	Type          LayoutElementType
	ClassName     string
	ChildElements []PageElement
}

func (l LayoutElement) GetType() string {
	return string(l.Type)
}

func (l LayoutElement) GetClassName() string {
	return l.ClassName
}

func NewLayoutElement(t LayoutElementType) (*LayoutElement, error) {
	if !validateLayoutElementType(t) {
		return nil, fmt.Errorf("invalid layout element type: %s", t)
	}
	return &LayoutElement{
		Type:      t,
		ClassName: t.String(),
	}, nil
}

func (l *LayoutElement) appendChild(child PageElement) {
	l.ChildElements = append(l.ChildElements, child)
}
