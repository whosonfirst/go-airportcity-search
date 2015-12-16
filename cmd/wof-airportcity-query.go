package main

import (
	"flag"
	"fmt"
	"github.com/blevesearch/bleve"
)

func main() {

	var db = flag.String("db", "", "The path to your search database")

	flag.Parse()
	args := flag.Args()

	index, _ := bleve.Open(*db)

	fields := make([]string, 0)
	fields = append(fields, "Latitude")
	fields = append(fields, "Longitude")
	fields = append(fields, "Names")

	for _, q := range args {

		query := bleve.NewQueryStringQuery(q)
		searchRequest := bleve.NewSearchRequest(query)
		searchRequest.Fields = fields

		searchResult, _ := index.Search(searchRequest)
		hits := searchResult.Hits

		fmt.Printf("query results for %s: %d\n", q, len(hits))

		for _, r := range hits {
			fmt.Printf("%s %s\n", r.ID, r.Fields["Names"])
			fmt.Printf("%s %.06f %.06f \n", r.ID, r.Fields["Latitude"], r.Fields["Longitude"])
		}
	}
}
