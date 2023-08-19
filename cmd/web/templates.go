package main

import (
	"html/template"
	"path/filepath"
	"time"

	"github.com/Ueetim/snippetbox/internal/models"
)

// holding structure for any dynamic data to be passed into the templates
type templateData struct {
	CurrentYear	int
	Snippet 	*models.Snippet
	Snippets	[]*models.Snippet
	Form		any
	Flash		string
}

// return a nicely formatted time.Time object
func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

// keep track of custom functions globally
var functions = template.FuncMap{
	"humanDate": humanDate,
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

		// parse base template file. FUncMap must be registered before parsing
		ts, err := template.New(name).Funcs(functions).ParseFiles("./ui/html/base.tmpl")
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