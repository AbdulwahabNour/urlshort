package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/AbdulwahabNour/urlshort/urlShort"
)

func main() {
	ymlName := flag.String("yml", "urls.yaml", "Path to YAML file containing" )
	flag.Parse()
	mux := defaultMux()
 
	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlShort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback

	yaml , err:= urlShort.YamlFileHandler(*ymlName)
    if err != nil  {
		log.Fatalf("Can't create Yaml redirect URL provider. \n%v", err)
	}
// 	yaml := `
// - path: /urlshort
//   url: https://github.com/gophercises/urlshort
// - path: /urlshort-final
//   url: https://github.com/gophercises/urlshort/tree/solution
// `

	yamlHandler, err := urlShort.YAMLHandler(yaml, mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
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

 