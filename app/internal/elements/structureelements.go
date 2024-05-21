package elements

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
