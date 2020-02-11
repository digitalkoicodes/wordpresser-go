package wordpresser

import (
	"fmt"
	"net/http"
)

type Wordpresser struct {
	Config
}

type RawRendered struct {
	Raw      string `json:"raw"`
	Rendered string `json:"rendered"`
}

func New() *Wordpresser {
	wp := &Wordpresser{}
	return wp
}

// func (wp *Wordpresser) Init(handler) {
// 	handler.HandleFunc("/wordpresser", wp.wordpresserHandler).Methods("GET")
// 	handler.HandleFunc("/wp-json/wp/v2/types", wp.getTypesHandler).Methods("GET")
// 	handler.HandleFunc("/wp-json/wp/v2/pages/{id}", wp.getPageHandler).Methods("GET")
// 	handler.HandleFunc("/wp-json/wp/v2/pages", wp.postPageHandler).Methods("POST")
// 	handler.HandleFunc("/wp-json/wp/v2/pages/autosaves", wp.postPageHandler).Methods("POST")
// 	handler.HandleFunc("/wp-json/wp/v2/media", wp.getMediaHandler).Methods("GET")
// 	handler.HandleFunc("/wp-json/wp/v2/types/{slug}", wp.getTypeHandler).Methods("GET")
// 	handler.HandleFunc("/wp-json/wp/v2/themes", wp.getThemeHandler).Methods("GET")

// }

func (wp *Wordpresser) wordpresserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Wordpresser!")
}
