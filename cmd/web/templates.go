package main

import (
	"html/template"
	"path/filepath"

	"github.com/Ueetim/snippetbox/internal/models"
)

// holding structure for any dynamic data to be passed into the templates
type templateData struct {
	Snippet 	*models.Snippet
	Snippets	[]*models.Snippet
}

func newTemplateCache() (map[string]*template.Template, error) {
	// map acts as a cache for templates
	cache := map[string]*template.Template{}

	// get a slice of all filepaths matching the pattern
	pages, err := filepath.Glob("./ui/html/pages/*.tmpl")
	if err != nil {
		return nil, err
	}

	// loop through page filepaths one by one
	for _, page := range pages {
		// extract the file name
		name := filepath.Base(page)

		// parse base template file
		ts, err := template.ParseFiles("./ui/html/base.tmpl")
		if err != nil {
			return nil, err
		}

		// add partials to the template set
		ts, err = ts.ParseGlob("./ui/html/partials/*.tmpl")
		if err != nil {
			return nil, err
		}

		// add the page template
		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		// add complete set to the map, with name as key
		cache[name] = ts
	}
	return cache, nil
}