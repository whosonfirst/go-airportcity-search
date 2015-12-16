package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/blevesearch/bleve"
	"net/http"
	"strconv"
)

// Please to de-dupe in wof-airportcity-index

type WOFRecord struct {
	Id        int64
	Names     []string
	Latitude  float64
	Longitude float64
}

func search(index bleve.Index, q string) ([]*WOFRecord, error) {

	fields := make([]string, 0)
	fields = append(fields, "Latitude")
	fields = append(fields, "Longitude")
	fields = append(fields, "Names")

	query := bleve.NewQueryStringQuery(q)
	req := bleve.NewSearchRequest(query)
	req.Fields = fields

	rsp, err := index.Search(req)

	if err != nil {
		return nil, err
	}

	results := make([]*WOFRecord, 0)

	for _, r := range rsp.Hits {

		id, _ := strconv.ParseInt(r.ID, 10, 64)
		f := r.Fields

		lat := f["Latitude"].(float64)
		lon := f["Longitude"].(float64)
		names := make([]string, 0)

		/*
			for _, n := range f["Names"] {
			    fmt.Println(n)
			}
		*/

		record := &WOFRecord{Id: id, Latitude: lat, Longitude: lon, Names: names}
		results = append(results, record)
	}

	return results, nil
}

func main() {

	var host = flag.String("host", "localhost", "The hostname to listen for requests on")
	var port = flag.Int("port", 8080, "The port number to listen for requests on")
	var cors = flag.Bool("cors", false, "Enable CORS headers")
	var db = flag.String("db", "", "The path to your search database")

	flag.Parse()

	index, _ := bleve.Open(*db)

	handler := func(rsp http.ResponseWriter, req *http.Request) {

		query := req.URL.Query()

		q := query.Get("q")

		if q == "" {
			http.Error(rsp, "Missing query parameter", http.StatusInternalServerError)
		}

		results, err := search(index, q)

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusInternalServerError)
		}

		js, err := json.Marshal(results)

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusInternalServerError)
			return
		}

		if *cors {
			rsp.Header().Set("Access-Control-Allow-Origin", "*")
		}

		rsp.Header().Set("Content-Type", "application/json")
		rsp.Write(js)
	}

	endpoint := fmt.Sprintf("%s:%d", *host, *port)

	fmt.Println(endpoint)

	http.HandleFunc("/", handler)
	http.ListenAndServe(endpoint, nil)
}
