package controllers

import (
	"fmt"
	"net/http"

	"github.com/aggdotred/lenslocked.com/views"
)

type Galleries struct {
	NewView *views.View
}

type CreateGalleryForm struct {
	GalleryName string `schema:"gallery-name"`
}

func NewGalleries() *Galleries {
	return &Galleries{
		NewView: views.NewView("bootstrap", "galleries/new"),
	}
}

// New gallery form
// POST /galleries/new
func (g *Galleries) New(w http.ResponseWriter, r *http.Request) {
	if err := g.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

// Create is used to process create gallery form when a user tries to create a new gallery
//
// POST /galleries/new
func (g *Galleries) Create(w http.ResponseWriter, r *http.Request) {

	var form CreateGalleryForm
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	fmt.Fprintln(w, "Gallery Name is", form.GalleryName)

}
