package elements

type BlockElementType string

const (
	TextBlock BlockElementType = "textblock"
)

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

func CreateBlockElement(t BlockElementType) *BlockElement {
	return &BlockElement{
		PageElement: PageElement{
			Type:        string(t),
			ElementUUID: createElementUUID(),
			Attributes:  map[string]string{},
		},
	}
}

func (e *BlockElement) AddChildElements(collection []PageElementInterface) {
	for _, element := range collection {
		element.SetParentUUID(e.ElementUUID)
		e.ChildElements = append(e.ChildElements, element)
	}
}
