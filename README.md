## getAcFunPage

----

## Intro

####  
Get Acfun Page on "The Hottest Today"

## Testing
```go
go test getAcFunPage/PageSave
```

## Dependence
```shell
$ go get github.com/emilsjolander/goson
```

## Usage

```shell
#running redis server
./redis-server.sh
```

```go
go run main.go
```

## Development Log

1. I make a big mistake when I declare some variable in struct IndexItem use begin with lowercase, That was suck! When I use goson to reflex IndexItem to JSON, Golang throw out **"reflect.Value.Interface: cannot return value obtained from unexported field or method"** error. It kill me a lot of time to change uppercase in every file.

