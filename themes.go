package wordpresser

import (
	"encoding/json"
	"net/http"
)

type Theme struct {
	ThemeSupports ThemeSupport `json:"theme_supports"`
}

type ThemeSupport struct {
	Formats          []string `json:"formats"`
	PostThumbnails   bool     `json:"post-thumbnails"`
	ResponsiveEmbeds bool     `json:"responsive-embeds"`
}

func (wp *Wordpresser) getThemeHandler(w http.ResponseWriter, r *http.Request) {
	themes := wp.ThemeStorer.Get()
	json.NewEncoder(w).Encode(themes)
}

type ThemeStorer interface {
	Get() []*Theme
}
