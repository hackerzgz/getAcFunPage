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
$ go get github.com/bmizerany/pat
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
```
cannot use pageList (type []IndexItem) as type []PageSave.IndexItem in argument...
```

3. **References to static files** is based on the relative path calls the function file.

4. Now here is a problem which `HandleGetResp()` will call twice time in Brower. But in `curl http://localhost:9000` ,it call only once.
```
Just log the requests. You will realize that your browser also requests /favicon.ico.
See https://en.wikipedia.org/wiki/Favicon for more information.
```

5. Golang panic error : 
```shell
http: panic serving 127.0.0.1:53512: dial tcp :6379: socket: too many open files
goroutine 5322 [running]:
net/http.(*conn).serve.func1(0xc820f87f80)
    /usr/local/go/src/net/http/server.go:1389 +0xc1
panic(0x797240, 0xc820b12050)
    /usr/local/go/src/runtime/panic.go:426 +0x4e9
main.GetPageAndJSON(0x0, 0x0)
    /home/hackerzgz/workspace/golang/src/getAcFunPage/main.go:130 +0x20a
main.HandleGetResp(0x7f2103407500, 0xc8212fb450, 0xc8210a68c0)
    /home/hackerzgz/workspace/golang/src/getAcFunPage/main.go:82 +0x18
net/http.HandlerFunc.ServeHTTP(0x8902f0, 0x7f2103407500, 0xc8212fb450, 0xc8210a68c0)
    /usr/local/go/src/net/http/server.go:1618 +0x3a
net/http.(*ServeMux).ServeHTTP(0xc820015740, 0x7f2103407500, 0xc8212fb450, 0xc8210a68c0)
    /usr/local/go/src/net/http/server.go:1910 +0x17d
net/http.serverHandler.ServeHTTP(0xc82008a680, 0x7f2103407500, 0xc8212fb450, 0xc8210a68c0)
    /usr/local/go/src/net/http/server.go:2081 +0x19e
net/http.(*conn).serve(0xc820f87f80)
    /usr/local/go/src/net/http/server.go:1472 +0xf2e
created by net/http.(*Server).Serve
    /usr/local/go/src/net/http/server.go:2137 +0x44e
```
when i use webbench in **client 300 and time 60s**.
This Problem Slove When you set `ulimit -n 99999` . 

> Because of Linux System file descriptors limit of your operating system(ubuntu defaults to 1024 which can be a problem) which is right the redis maxclients setting.

> [-- by Stack Overflow](http://stackoverflow.com/questions/19971968/go-golang-redis-too-many-open-files-error)

6. Same in 5. It will log:
```shell
=== Get PageList Done ===
=== Get PageList Done ===
=== JSON Trans Done ===
=== JSON Trans Done ===
=== JSON Trans Done ===
=== JSON Trans Done ===
=== JSON Trans Done ===
=== Get PageList Done ===
=== JSON Trans Done ===
=== JSON Trans Done ===
=== JSON Trans Done ===
=== JSON Trans Done ===
=== JSON Trans Done ===
=== JSON Trans Done ===
=== JSON Trans Done ===
```
It show that the Efficiency of JSON Trans were so low. So I Cache PageList JSON to Redis too.

7. Acfun `404` when GetPageInfo() not Finish. It throw
```shell
2016/08/02 16:38:01 statusCode -->  200 OK
Get 2952333 PageInfo Error.
panic: runtime error: index out of range
```
So I return `-1` when Acfun was `404`.

## Benchmark

+ System   : CentOS 7.0 64bit
+ CPU      : 1 Core
+ Memory   : 1 GB
+ Bandwidth: 1Mbps

![WebBench -c 1000 -t 15](https://github.com/HackeZ/getAcFunPage/blob/master/doc/benchmark-c1k-t15.png)

WebBench -c 1000 -t 15

![WebBench -c 300 -t 60](https://github.com/HackeZ/getAcFunPage/blob/master/doc/benchmark-c300-t60.png)

WebBench -c 300 -t 60

## Demo
[Demo](http://123.207.0.81:9001/)
