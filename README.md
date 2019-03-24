# schedule

### Build & Run
```shell
go get -v github.com/LinJinghua/schedule
cd $GOPATH/src/github.com/LinJinghua/schedule # Linux
cd /d %GOPATH%/src/github.com/LinJinghua/schedule # Windows
go build && schedule
```

### Test
```shell
curl localhost:8080/up
curl localhost:8080/down
curl localhost:8080/get

ab -n10000 -c1000 http://localhost:8080/get
ab -n10000 -c1000 http://localhost:8080/up
ab -n10000 -c1000 http://localhost:8080/down
```

