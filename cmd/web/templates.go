package main

import (
	"github.com/Ueetim/snippetbox/internal/models"
)

// holding structure for any dynamic data to be passed into the templates
type templateData struct {
	Snippet 	*models.Snippet
	Snippets	[]*models.Snippet
}