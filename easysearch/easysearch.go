package main

import (
	"fmt"
	"github.com/blevesearch/bleve"
)

func main() {
	// Open an index if it already exists.
	index, err := bleve.Open("example.bleve")
	if err != nil {
		mapping := bleve.NewIndexMapping();
		index, err = bleve.New("example.bleve", mapping)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	data := struct {
		Name string
	} {
		Name: "text",
	}

	// Index some data.
	index.Index("id", data);

	// Search for some text
	query := bleve.NewMatchQuery("text")
	search := bleve.NewSearchRequest(query);
	searchResults, err := index.Search(search)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(searchResults)
}
