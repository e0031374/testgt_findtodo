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

## Flag: match type `-e`
1. Exact match 
    + Matches only `"TODO"`, not `//"TODO"`, not `abc"TODO"def`
    + `$ go run main.go -e=true`
1. Contains substring match
    + Matches all `"TODO"`, `//"TODO"`, and `abc"TODO"def`
    + `$ go run main.go -e=false`
1. Default: contains substring match

## Flag: path `-p`
1. Run program with `./test` as the root
    + `$ go run main.go -p="./test"`
1. Default: run with `.` current dir as root dir
    + the following are comparable
        + `$ go run main.go`
        + `$ go run main.go -p="."`

## Flag: help `-h`
1. Check usage of tool
    + `$ go run main.go -h`
