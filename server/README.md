# Server
Implement RESTful API for client.

## 1.Prerequiste

#### Download `go` from: https://golang.org

#### Add `go` to PATH if installer did not add it , type "go" in terminal to check if it works.

####  For MacOs and Linux, type:
    ./configure.sh
#### for the first time.

#### For Windows, current working directory should be add to `GOPATH` if `GOPATH` is determined by Environment Variables. If not, try:
    go env -w GOPATH=/your/current/working/dir:$GOPATH


## 2. Build and Run
#### build and run
    make run

#### build
    make build
    
The executable binary file will be generated in ./bin/

#### clean
    make clean

#### For Windows, try to compile and run through go command directly according to Makefile

