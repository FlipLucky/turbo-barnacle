package elements

type ImageElement struct {
	Type      string
	ClassName string
	AssetUrl  string
	AltText   string
}

func NewImageElement(assetUrl, altText string) ImageElement {
	return ImageElement{
		Type:     "image",
		AssetUrl: assetUrl,
		AltText:  altText,
	}
}

func (i ImageElement) GetAssetUrl() string {
	return i.AssetUrl
}

func (i ImageElement) GetAltText() string {
	return i.AltText
}

func (i ImageElement) GetClassName() string {
	return i.ClassName
}

func (i ImageElement) GetType() string {
	return i.Type
}

type VideoElement struct {
	Type      string
	ClassName string
	AssetUrl  string
}
