# Server

## Requirements

* [Golang 1.19.x](https://golang.org/doc/install)
* [golangci-lint latest version](https://github.com/golangci/golangci-lint#install)
* [GNU Make 4.3.x](https://www.gnu.org/software/make/)
* [gofumpt](https://github.com/mvdan/gofumpt) -> `go install mvdan.cc/gofumpt@latest`

### Editor/IDE setup

See [gopls](https://pkg.go.dev/golang.org/x/tools/gopls#section-readme) (Go language server)

## Commands

```
make
make test
make lint
make build
make fmt
```

## See all possible linters

```
golangci-lint help linters
```

## Create test coverage

```
go test -coverprofile=cover.txt
```
```
go tool cover -html=cover.txt -o cover.html
```

## Golang documentation

* [Tour of Go](https://go.dev/tour/welcome/1)
* [FAQ](https://go.dev/doc/faq)
* [Effective Go](https://go.dev/doc/effective_go)
* [Tutorials and examples](https://go.dev/doc/)
* [Blog](https://go.dev/blog/)
* [Go Tools](https://pkg.go.dev/cmd/go)
