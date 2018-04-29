package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

// Adventure is piece of whole story.
// Adventure implements Handler interface.
type Adventure struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options,omitempty"`
}

// Option contains navigation fields.
type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func (a *Adventure) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/index.html"))
	if err := t.Execute(w, a); err != nil {
		log.Fatalln("t.Execute: ", err)
	}
}

func main() {
	// Read json data from file
	data, err := ioutil.ReadFile("gopher.json")
	if err != nil {
		log.Fatalln("ioutil.ReadFile: ", err)
	}

	// Parse json from file to map
	Advs := make(map[string]Adventure)
	err = json.Unmarshal(data, &Advs)
	if err != nil {
		log.Fatalln("json.Unmarshal: ", err)
	}

	// Register all handlers from map
	for name := range Advs {
		http.Handle("/"+name, &Adventure{
			Title:   Advs[name].Title,
			Story:   Advs[name].Story,
			Options: Advs[name].Options,
		})
	}

	// Static files
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))

	// Use redirect from / to /intro, because we missed / earlier
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/intro", http.StatusPermanentRedirect)
	})

	// Start the server
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln("http.ListenAndServe: ", err)
	}
}
