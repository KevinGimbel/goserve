# goserve

Static file server written in go.


### Install
1. [Install Go](https://golang.org/doc/install#install).
2. Clone the repo `git clone https://github.com/kevingimbel/goserve.git`
3. Run `go build goserve.go` from within the new directory.

If you would like to run the program from everywhere, link it into you `$PATH` variable, e.g.:
```sh
  $ sudo ln -s $(pwd)/goserve /usr/local/bin
```

### Usage
`goserve` can be used from the command line as follows:

```sh
  $ goserve [-port ""]
```
This will serve the current directory to `localhost:8000` or the specified port.

### Test it
To test the server run `goserve` from the project directory and open [localhost:8000/example](localhost:8000/example)
