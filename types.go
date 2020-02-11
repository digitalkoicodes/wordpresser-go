package wordpresser

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Type struct {
	Id       int         `json:"id"`
	Name     string      `json:"name"`
	RestBase string      `json:"rest_base"`
	Slug     string      `json:"slug"`
	Supports TypeSupport `json:"supports"`
	Viewable bool        `json:"viewable"`
}

type TypeSupport struct {
	Author         bool `json:"author"`
	Comments       bool `json:"comments"`
	CustomFields   bool `json:"custom-fields"`
	Discussion     bool `json:"discussion"`
	Editor         bool `json:"editor"`
	Excerpt        bool `json:"excerpt"`
	PageAttributes bool `json:"page-attributes"`
	Revisions      bool `json:"revisions"`
	Thumbnail      bool `json:"thumbnail"`
	Title          bool `json:"title"`
}

func (wp Wordpresser) GetTypesHandler(w http.ResponseWriter, r *http.Request) error {
	types := wp.TypeStorer.GetTypes()
	json.NewEncoder(w).Encode(types)
	return nil
}

func (wp Wordpresser) GetTypeHandler(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	typeObject := wp.TypeStorer.Get(vars["slug"])
	json.NewEncoder(w).Encode(typeObject)
	return nil
}

type TypeStorer interface {
	GetTypes() map[string]Type
	Get(slug string) *Type
}
