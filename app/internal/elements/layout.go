package elements

import "fmt"

type LayoutElementType string

func (l LayoutElementType) String() string {
	return string(l)
}

const (
	Section   LayoutElementType = "section"
	Container LayoutElementType = "container"
	Row       LayoutElementType = "row"
	Col       LayoutElementType = "col"
)

func validateLayoutElementType(t LayoutElementType) bool {
	switch t {
	case Container, Row, Col, Section:
		return true
	default:
		return false
	}
}

type LayoutElement struct {
	Type          LayoutElementType
	ClassName     []string
	ChildElements []PageElement
}

func (l LayoutElement) GetType() string {
	return string(l.Type)
}

func (l LayoutElement) GetClassName() string {
	className := ""
	for _, className := range l.ClassName {
		className += " " + className
	}
	return className
}

func NewLayoutElement(t LayoutElementType) (*LayoutElement, error) {
	if !validateLayoutElementType(t) {
		return nil, fmt.Errorf("invalid layout element type: %s", t)
	}
	classNames := []string{
		t.String(),
	}
	return &LayoutElement{
		Type:      t,
		ClassName: classNames,
	}, nil
}

func (l *LayoutElement) AppendChild(child PageElement) {
	l.ChildElements = append(l.ChildElements, child)
}
