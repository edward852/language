
# 声明
主要代码来自[Dave Cheney](https://dave.cheney.net)，原文在[这里](https://dave.cheney.net/high-performance-go-workshop/gophercon-2019.html)，代码在[这里](https://github.com/davecheney/high-performance-go-workshop/tree/master/examples/words)。    

# 前置准备
- go
- chrome
- graphviz

# pprof
## cpu
```shell
go tool pprof -http=:8080 cpu.pprof
```

## memory
```shell
go tool pprof -http=:8080 mem.pprof
```

# trace
```shell
go tool trace -http=:8080 trace.out
```
