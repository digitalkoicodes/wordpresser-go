package wordpresser

type Config struct {
	PageStorer  PageStorer
	TypeStorer  TypeStorer
	MediaStorer MediaStorer
	ThemeStorer ThemeStorer
}
