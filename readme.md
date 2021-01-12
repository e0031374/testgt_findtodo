# find TODO cli tool

## Prerequisites
1. install go
    + golang site to install [link](https://golang.org/doc/install)
1. check go version
    ```
    $ go version
    go version go1.15.6 linux/amd64
    ```
1. set $GOPATH
    + assuming `main.go` is in `$HOME/go/src/findTodo/main.go`
    + `export GOPATH=$HOME/go`

## Operation
1. print all files from root directory with default settings
    + `$ go run main.go`
1. print all files from specified root directory
    + `$ go run main.go [path/to/dir]`

## Flags
+ specify all flags before cli-argument for root directory to run tool
    + indicated as `[path/to/dir]`

### Flag: match type `-e`
1. Exact match 
    + Matches only `"TODO"`, not `//"TODO"`, not `abc"TODO"def`
    + `$ go run main.go -e=true [path/to/dir]`
1. Contains substring match
    + Matches all `"TODO"`, `//"TODO"`, and `abc"TODO"def`
    + `$ go run main.go -e=false [path/to/dir]`
1. Default: contains substring match

### Flag: absolute path to file `-a`
1. Specify absolute path from `/` to flagged file
    + `$ go run main.go -a=true [path/to/dir]`
1. Specify relative path to flagged file (default)
    + `$ go run main.go`

### Flag: help `-h`
1. Check usage of tool
    + `$ go run main.go -h`
