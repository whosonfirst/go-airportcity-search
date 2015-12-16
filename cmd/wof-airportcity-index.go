package main

import (
	"flag"
	"fmt"
	"github.com/blevesearch/bleve"
	csv "github.com/whosonfirst/go-whosonfirst-csv"
	geojson "github.com/whosonfirst/go-whosonfirst-geojson"
	"io"
	"path"
	"strconv"
)

// Please to de-dupe in wof-airportcity-server

type WOFRecord struct {
	Id        string
	Latitude  float64
	Longitude float64
	Names     []string
}

func main() {

	var source = flag.String("source", "", "The source directory where WOF data lives")
	var db = flag.String("db", "", "The path to your search database")

	flag.Parse()
	args := flag.Args()

	mapping := bleve.NewIndexMapping()

	index, err := bleve.New(*db, mapping)

	if err != nil {
		panic(err)
	}

	for _, csv_path := range args {

		reader, reader_err := csv.NewDictReader(csv_path)

		if reader_err != nil {
			panic(reader_err)
		}

		for {
			row, err := reader.Read()

			if err == io.EOF {
				break
			}

			if err != nil {
				panic(err)
			}

			rel_path, ok := row["path"]

			if !ok {
				continue
			}

			abs_path := path.Join(*source, rel_path)
			feature, err := geojson.UnmarshalFile(abs_path)

			if err != nil {
				panic(err)
			}

			id := feature.Id()
			name := feature.Name()

			str_id := strconv.Itoa(id)

			names := make([]string, 0)
			names = append(names, name)

			body := feature.Body()

			lat, ok := body.Path("properties.geom:latitude").Data().(float64)

			if !ok {
				fmt.Printf("%s missing latitude, skipping\n", str_id)
				continue
			}

			lon, ok := body.Path("properties.geom:longitude").Data().(float64)

			if !ok {
				continue
			}

			if lat == 0.0 || lon == 0.0 {
				continue
			}

			properties, _ := body.S("properties").ChildrenMap()

			for key, details := range properties {

				if key != "wof:concordances" {
					continue
				}

				conc, _ := details.ChildrenMap()

				codes := make(map[string]bool)

				for k, v := range conc {

					code := ""

					if k == "faa:code" {
						code = v.Data().(string)
					} else if k == "iata:code" {
						code = v.Data().(string)
					} else if k == "icao:code" {
						code = v.Data().(string)
					} else {
						// pass
					}

					if code != "" {
						codes[code] = true
					}
				}

				for code, _ := range codes {
					names = append(names, code)
				}
			}

			record := &WOFRecord{Id: str_id, Names: names, Latitude: lat, Longitude: lon}
			fmt.Printf("%s\n", record)
			index.Index(str_id, record)
		}
	}

}
