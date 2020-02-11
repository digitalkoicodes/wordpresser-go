package wordpresser

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Page struct {
	Date          string      `json:"date"`
	DateGmt       string      `json:"date_gmt"`
	Guid          string      `json:"guid"`
	Id            int         `json:"id"`
	Link          string      `json:"link"`
	Modified      string      `json:"modified"`
	ModifiedGmt   string      `json:"modified_gmt"`
	Slug          string      `json:"slug"`
	Status        string      `json:"status"`
	Type          string      `json:"type"`
	Password      string      `json:"password"`
	Parent        int         `json:"parent"`
	Title         RawRendered `json:"title"`
	Content       RawRendered `json:"content"`
	Author        int         `json:"author"`
	Excerpt       RawRendered `json:"excerpt"`
	FeaturedMedia int         `json:"featured_media"`
	CommentStatus string      `json:"comment_status"`
	PingStatus    string      `json:"ping_status"`
	MenuOrder     int         `json:"menu_order"`
	Meta          string      `json:"meta"`
	Template      string      `json:"template"`
}

func (wp *Wordpresser) GetPageHandler(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	page := wp.PageStorer.Get(id)
	err := json.NewEncoder(w).Encode(page)
	if err != nil {
		return err
	}
	return nil
}

func (wp *Wordpresser) PostPageHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var page *Page
	err := decoder.Decode(&page)
	if err != nil {
		panic(err)
	}
	page = wp.PageStorer.Post(page)
	json.NewEncoder(w).Encode(page)
}

type PageStorer interface {
	Get(id int) *Page
	Post(data *Page) *Page
}
