package routes

import (
	"github.com/FlipLucky/turbo-barnacle/internal/elements"
)

type PageRoute interface {
	GetSlug() string
	GetElements() []elements.PageElement
	GetTemplate() string
	GetTitle() string
}
