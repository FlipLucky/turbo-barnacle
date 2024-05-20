package elements

type TextElementType string

const (
	H1        TextElementType = "h1"
	H2        TextElementType = "h2"
	H3        TextElementType = "h3"
	H4        TextElementType = "h4"
	H5        TextElementType = "h5"
	H6        TextElementType = "h6"
	Paragraph TextElementType = "p"
)

type TextElement struct {
	ElementType TextElementType
	Body        string
	ClassName   string
	ElementUUID ElementUUID
}

func (t TextElement) GetType() string {
	return string(t.ElementType)
}

func (t TextElement) GetClassName() string {
	return t.ClassName
}

func (t TextElement) GetElementUUID() ElementUUID {
	return t.ElementUUID
}

func CreateTextElement(t TextElementType) *TextElement {
	e := new(TextElement)
	e.ElementType = t
	e.ElementUUID = createElementUUID()
	return e
}
