
go-build:
	CGO_ENABLED=0 go build -o build/x-byte
start:
	alias air='$(go env GOPATH)/bin/air'
	air
clean:
	rm -rf build
	