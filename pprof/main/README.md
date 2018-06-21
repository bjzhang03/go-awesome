```bash
#编译文件
go build -o mytest fibonacciexample.go
#生成需要的分析文件
./mytest -cpuprofile=mytest.prof
#使用tool对生成的文件进行分析
go tool pprof mytest mytest.prof
```
