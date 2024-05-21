package elements

type ContentElementType string

const (
	BodyText      ContentElementType = "bodytext"
	SubHeaderText ContentElementType = "subheadertext"
	HeaderText    ContentElementType = "headertext"
)

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

func CreateContentElement(t ContentElementType) *ContentElement {
	return &ContentElement{
		PageElement: PageElement{
			Type:        string(t),
			ElementUUID: createElementUUID(),
			Attributes:  map[string]string{},
		},
	}
}
