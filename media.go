package wordpresser

import (
	"encoding/json"
	"net/http"
)

type Media struct {
	Id           int          `json:"id"`
	Caption      RawRendered  `json:"caption"`
	DateGmt      string       `json:"date_gmt"`
	Date         string       `json:"date"`
	Link         string       `json:"link"`
	MediaType    string       `json"media_type"`
	MimeType     string       `json:"mime_type"`
	SourceUrl    string       `json:"source_url"`
	MediaDetails MediaDetails `json:"media_details"`
	Title        RawRendered  `json:"title"`
}

type MediaDetails struct {
	File      string                 `json:"file"`
	Height    string                 `json:"height"`
	ImageMeta map[string]interface{} `json:"image_meta"`
	Sizes     map[string]interface{} `json:"sizes"`
	Width     string                 `json:"width"`
}

func (wp *Wordpresser) getMediaHandler(w http.ResponseWriter, r *http.Request) {
	media := wp.MediaStorer.Get()
	json.NewEncoder(w).Encode(media)
}

type MediaStorer interface {
	Get() []*Media
}
