# goserve
> Static file server written in go.

## Installation

### From archive

Go to the [release page](https://github.com/kevingimbel/goserve/releases) and download a binary for your operating system.
`goserve` is build for Windows, MacOS and Linux with [goreleaser](https://github.com/goreleaser/goreleaser).

### From source

1. [Install Go](https://golang.org/doc/install#install).
2. Clone the repo `git clone https://github.com/kevingimbel/goserve.git`
3. Run `go build goserve.go` from within the new directory.

Place the binary somewhere in your `$PATH`, e.g.
```sh
  $ sudo ln -s $(pwd)/goserve /usr/local/bin
```

### Usage

`goserve` can be used from the command line as follows:

```sh
  $ goserve [-port ""] [-verbose]
```
This will serve the current directory to `localhost:8000` or the specified port. `-verbose` logs each request to Stdout.

See also `goserve -help`.

### Test it

To test the server run `goserve` from the project directory and open [localhost:8000/example](localhost:8000/example).
