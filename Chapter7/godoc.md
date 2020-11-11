1. Open up a code repository that has a package or code that you wish to generate documentation from
2. Create a package for the Go program and install it on your localhost
3. Move the binary package to either your $GOROOT or $GOPATH
4. Install godoc - `go get -v  golang.org/x/tools/cmd/godoc`
5. To run Godoc on localhost - `godoc -http=:6060`