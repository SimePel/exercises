package links

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/net/html"
)

// Find finds all links in file and returns them in the format:
//
//	Link {
//		Href: "/path",
//		Text: "Text inside <a> tag",
//	}
//
func Find(file string) string {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalln("ioutil.ReadFile: ", err)
	}

	doc, err := html.Parse(bytes.NewReader(data))
	if err != nil {
		log.Fatalln("html.Parse: ", err)
	}

	var f func(*html.Node) string
	var s string
	f = func(n *html.Node) string {
		if n.Type == html.ElementNode && n.Data == "a" {
			s += fmt.Sprintln("Link {")
			for _, a := range n.Attr {
				if a.Key == "href" {
					s += fmt.Sprintf("  Href: \"%s\",\n", a.Val)
					s += fmt.Sprintf("  Text: \"%s\",\n", n.FirstChild.Data)
					break
				}
			}
			s += fmt.Sprintln("}")
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
		return s
	}
	return f(doc)
}
