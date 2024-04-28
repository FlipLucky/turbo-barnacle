package elements

type Page struct {
	Title    string
	Slug     string
	Template string
	Elements []PageElement
}

func (p Page) GetTitle() string {
	return p.Title
}

func (p Page) GetSlug() string {
	return p.Slug
}

func (p Page) GetTemplate() string {
	return p.Template
}

func (p Page) GetElements() []PageElement {
	return p.Elements
}

func NewPage(title, slug, template string, pageElements []PageElement) Page {
	return Page{
		Title:    title,
		Slug:     slug,
		Template: template,
		Elements: pageElements,
	}
}

func (p *Page) AppendElement(element PageElement) {
	p.Elements = append(p.Elements, element)
}
