package main

import (
	"html/template"
	"log"
	"net/http"
)

// StoryHandler implements Handler interface.
type StoryHandler struct {
	Title   string        `json:"title"`
	Story   []string      `json:"story"`
	Options []StoryOption `json:"options,omitempty"`
}

// StoryOption is @TODO
type StoryOption struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func (sh *StoryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/index.html"))
	if err := t.Execute(w, sh); err != nil {
		log.Fatalln("t.Execute: ", err)
	}
}

func main() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))

	http.Handle("/", &StoryHandler{
		Title: "Little Gopher",
		Story: []string{
			"Hey, how are you?",
			"It's alright!",
		},
		Options: []StoryOption{
			StoryOption{
				Text: "piece of text",
				Arc:  "prev",
			},
			StoryOption{
				Text: "piece of text",
				Arc:  "next",
			},
		},
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln("http.ListenAndServe: ", err)
	}
}
