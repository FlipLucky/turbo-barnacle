package elements

type Page struct {
	Id       int
	Title    string
	Slug     string
	Template string
	Elements []PageElementInterface
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

func (p Page) GetElements() []PageElementInterface {
	return p.Elements
}

func NewPage(title, slug, template string, pageElements []PageElementInterface) Page {
	return Page{
		Title:    title,
		Slug:     slug,
		Template: template,
		Elements: pageElements,
	}
}

func (p *Page) AppendElement(element PageElementInterface) {
	p.Elements = append(p.Elements, element)
}
