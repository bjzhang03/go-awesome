####cgo调用C++函数实现

大概的流程为：
* 将C++中的某个函数转换成C的形式，变成端口的形式暴露出去
* 将声明放在.h文件中，将函数体封装在动态库或者静态库中，供go调用
* 在go中通过cgo LDFLAGS指明Lib所在的地址，通过cgo CFLAGS指令include所在的地址，即可调用

在这里我走了一些弯路，供参考：
* 忽略#cgo LDFLAGS #cgo CFLAGS以为是被注释的部分，其实很重要
* c.h文件中忽视#ifdef __cplusplus该指令指明内部使用c++的方式进行编译
* 忽略test.go文件中编译制导指令　-lstdc++ 很关键

最后运行的命令

```bash
./build.sh
go build test.go
```
