

all: build linter

build: clean get
	go build

clean: 
	go clean

get:
	go get -v

cover: build
	go test -coverprofile=coverage.out
	go tool cover -html=coverage.out

linter:
	gometalinter ./...

doc:
	godoc -http=:6060 &
	 xdg-open http://0.0.0.0:6060/pkg/github.com/zakyke/curlit/