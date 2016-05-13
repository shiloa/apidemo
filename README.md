## Simple demo on how to build a [very] basic RESTful API with Go

Used for tech-talk I've given. 

Note: Based on Golang version `1.3`, and assume `$GOPATH` is configured, as well as the DB connection.

```sh
$ go get github.com/shiloa/apidemo
$ cd $GOPATH/src/github.com/shiloa/apidemo
$ go run server.go
$ curl localhost:3001
```

