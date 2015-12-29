prep:
	if test -d pkg; then rm -rf pkg; fi

self:   prep

deps:   self
	go get -u "github.com/blevesearch/bleve"
	go get -u "github.com/whosonfirst/go-whosonfirst-csv"
	go get -u "github.com/whosonfirst/go-whosonfirst-geojson"
fmt:
	go fmt cmd/*.go

bin:	self
	@GOPATH=$(shell pwd) \
	go build -o bin/wof-airportcity-index cmd/wof-airportcity-index.go
	@GOPATH=$(shell pwd) \
	go build -o bin/wof-airportcity-query cmd/wof-airportcity-query.go
	@GOPATH=$(shell pwd) \
	go build -o bin/wof-airportcity-server cmd/wof-airportcity-server.go
