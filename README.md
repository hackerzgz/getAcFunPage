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

2. **Same struct** must be a package to import OR It will show you:

> cannot use pageList (type []IndexItem) as type []PageSave.IndexItem in argument...

3. **References to static files** is based on the relative path calls the function file.

4. Now here is a problem which `HandleGetResp()` will call twice time in Brower. But in `curl http://localhost:9000` ,it call only once.

> Just log the requests. You will realize that your browser also requests /favicon.ico.
> See https://en.wikipedia.org/wiki/Favicon for more information.

