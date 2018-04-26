package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var yamlFile string

func init() {
	flag.StringVar(&yamlFile, "yaml", "", "set the location of yaml file")
}

func main() {
	flag.Parse()
	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/gh":         "https://github.com",
		"/yaml-godoc": "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := MapHandler(pathsToUrls, mux)

	yaml, err := ioutil.ReadFile(yamlFile)
	if err != nil {
		log.Fatalln("ReadFile: ", err)
	}

	yamlHandler, err := YAMLHandler(yaml, mapHandler)
	if err != nil {
		log.Fatalln("YAMLHandler: ", err)
	}

	log.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
