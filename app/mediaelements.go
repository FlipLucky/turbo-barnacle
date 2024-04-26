package main

type ImageElement struct {
	Type      string
	ClassName string
	AssetUrl  string
	AltText   string
}

func (i ImageElement) GetType() string {
	return i.Type
}

func (i ImageElement) GetClassName() string {
	return i.ClassName
}

type VideoElement struct {
	Type      string
	ClassName string
	AssetUrl  string
}
