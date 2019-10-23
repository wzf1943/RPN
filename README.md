# RPN
RPN is a web server that can parse reverse polish notation string and then write result as JSON format.

## System Diagram
RPN will use master and slave mode to make sure whole system has high avaible feature. The distribuit system algorithm will use memberlist protocal. 
<p align="center">
![avatar](https://raw.githubusercontent.com/wzf1943/RPN/master/doc/API.png)
<p align="center">
![avatar](https://raw.githubusercontent.com/wzf1943/RPN/master/doc/system.png)


## Getting Started
These instructions will get you a copy of the RPN up and running on your local machine.

### Prerequisites
golang:1.10
### Installing
```
go get github.com/gorilla/mux
go get github.com/golang-collections/collections/stack
```
## Running the tests
### Unit test
Run the go unit test
```
make test
```

## Build
### Compile Binary

```
make build
```
### Cross compile binary
```
make build-linux
```
### Clean binaries
```
make clean
```
### Docker build
```
make docker-build
```
### Docker run
```
make docker-run
``` 

## Test Endpoint
```
curl -H "Accept: application/json" -H "Content-type: application/json" -X POST -d '{"rpns":[{"input":"10 -1 +"},{"input":"10 -2 +"}]}' http://localhost:8080/parse
```


## API
### /parse
This api recieves single reverse polish notation string as input and return result as JSON format.

