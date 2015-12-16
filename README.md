# go-airportcity-search

An airport specific search engine for Airport City, using Who's On First data.

## Set up

Use the handy `Makefile` for fetching dependencies and building the `wof-airportcity-*` binaries, like this:

```
$> make deps
if test -d pkg; then rm -rf pkg; fi
go get -u "github.com/blevesearch/bleve"
go get -u "github.com/whosonfirst/go-whosonfirst-csv"
go get -u "github.com/whosonfirst/go-whosonfirst-geojson"
```

And then:

```
$> make bin
if test -d pkg; then rm -rf pkg; fi
go build -o bin/wof-airportcity-index cmd/wof-airportcity-index.go
go build -o bin/wof-airportcity-query cmd/wof-airportcity-query.go
go build -o bin/wof-airportcity-server cmd/wof-airportcity-server.go
```

## Usage

### wof-airportcity-index

```
$> ./bin/wof-airportcity-index -db test -source /usr/local/mapzen/whosonfirst-data/data/ /usr/local/mapzen/whosonfirst-data/meta/wof-campus-latest.csv
&{102544341 %!s(float64=3.95944) %!s(float64=-59.124199) [Annai Airport NAI SYAN]}
&{102544219 %!s(float64=55.063301) %!s(float64=14.7596) [Bornholm Lufthavn RNN EKRN]}
&{102544125 %!s(float64=32.325001) %!s(float64=15.061) [Misurata Airport MRA LY-MRA]}
&{102544685 %!s(float64=43.323502) %!s(float64=3.3539) [Aéroport Béziers-Vias LFMU BZR]}
&{102544735 %!s(float64=43.556301) %!s(float64=2.28918) [Aéroport Castres-Mazamet DCM LFCK]}
&{102544009 %!s(float64=54.639199) %!s(float64=25.295658) [Vilnius Airport]}
&{102544597 %!s(float64=-4.6204) %!s(float64=143.4516) [Mont de Marsan Airport MDM AYDK]}
&{102544823 %!s(float64=45.658298) %!s(float64=-0.3175) [Aéroport Cognac-Châteaubernard LFBG CNG]}
&{102544861 %!s(float64=47.986841) %!s(float64=1.769192) [Bricy Airport]}
# and so on...
```

### wof-airportcity-query

```
$> ./bin/wof-airportcity-query -db test JFK Dallas
query results for JFK: 1
102534365 [John F Kennedy Int'l Airport JFK KJFK]
102534365 40.642220 -73.787074 
query results for Dallas: 3
102528103 [Dallas Executive Airport RBD KRBD]
102528103 32.680901 -96.868202 
102525799 [Dallas Love Field DAL KDAL]
102525799 32.844590 -96.850777 
102528541 [Dallas-Fort Worth International Airport KDFW DFW]
102528541 32.895265 -97.049809 
```

### wof-airportcity-server

```
$> ./bin/wof-airportcity-server -db test
```

And then:

```
$> curl -s 'http://localhost:8080?q=TRUDEAU' | python -mjson.tool
[
    {
        "Id": 102554351,
        "Latitude": 45.463215,
        "Longitude": -73.744442,
        "Names": []
    }
]
```

_The absence of names in the response is a bug. Because Go is weird._

## Caveats

In no particular order:

* This is work in progress.
* It uses the [Bleve fulltext document index](http://www.blevesearch.com/) for all the heavy-lifting.
* It will likely be abstracted out in to a (hopefully) generic `go-whosonfirst-bleve` package.
* The indexing tool currently assumes that all WOF records of placetype `campus` are airports. This is not true.
* Needs logging.

## See also

* http://www.blevesearch.com/
* https://github.com/whosonfirst/tangram-airportcity
* https://github.com/whosonfirst/whosonfirst-data
