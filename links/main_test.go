package links_test

import (
	"strconv"
	"testing"

	"github.com/SimePel/exercises/links"
)

func TestFind(t *testing.T) {
	expected := []string{
		`Link {
  Href: "/other-page",
  Text: "A link to another page",
}
`,
		`Link {
  Href: "https://www.twitter.com/joncalhoun",
  Text: "Check me out on twitter",
}
Link {
  Href: "https://github.com/gophercises",
  Text: "Gophercises is on Github",
}
`}
	for i := 0; i < len(expected); i++ {
		got := links.Find("ex" + strconv.Itoa(i+1) + ".html")
		if got != expected[i] {
			t.Errorf("links.Find returns:\n%sexpected:\n%s", got, expected[i])
		}
	}
}
