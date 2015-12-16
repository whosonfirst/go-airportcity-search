# go-airportcity-search

An airport specific search engine for Airport City, using Who's On First data.

## Usage

### wof-airportcity-index

```
$> ./bin/wof-airportcity-index -db test -source /usr/local/mapzen/whosonfirst-data/data/ /usr/local/mapzen/whosonfirst-data/meta/wof-campus-latest.csv
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

_The absence of names in the response is a bug._

## Caveats

In no particular order:

* This is work in progress
* It uses the [Bleve fulltext document index](http://www.blevesearch.com/) for all the heavy-lifting
* It will likely be abstracted out in to a (hopefully) generic `go-whosonfirst-bleve` package

## See also

* http://www.blevesearch.com/
* https://github.com/whosonfirst/tangram-airportcity
