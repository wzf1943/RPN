# RPN
RPN is a web server that can parse reverse polish notation string and then write result as JSON format.

## System Diagram 
---
![avatar](https://raw.githubusercontent.com/wzf1943/RPN/master/doc/API.png)
---
![avatar](https://raw.githubusercontent.com/wzf1943/RPN/master/doc/system.png)


## Getting Started
These instructions will get you a copy of the RPN up and running on your local machine.

### Prerequisites
golang:1.10
### Installing
```
go get github.com/gorilla/mux
go get github.com/golang-collections/collections/stack
go get github.com/jessevdk/go-flags
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
```
curl -H "Accept: application/json" -H "Content-type: application/json" -X GET http://localhost:8080/health
```

## API
### /parse
This api recieves single reverse polish notation string as input and return result as JSON format.
#### Input
```json
{
	"rpns": [{
			"input": "10 -1 +"
		},
		{
			"input": "10 -2 +"
		}
	]
}
```
#### Output
```
{
    "rpn": [
        {
            "name": "10 -1 +",
            "result": 9
        },
        {
            "name": "10 -2 +",
            "result": 8
        }
    ]
}
```
### /Health
This api helps load balancer check the server's healty situation
#### Input
N/A
#### Output
```json
{
    "status": "OK",
    "code": 200
}
```