# herror 是一个简单处理错误的包
## 介绍
go语言的错误处理的方式一直让人吐槽，每写一行代码就要 if err != nil 一下,非常的痛苦。herror就是为了解决这个痛点，我们可以在一些不太关注错误的地方使用它，比如快速的编写脚本。
## 快速开始
```go
package main

import "github.com/leijux/herror"

func F() ( int , error ) {
    return 1,errors.New("err")
}

func main() {
    herror.Try(F()).Must()
}
```
## 示例
通常对于返回的err，我们使用如下方法处理。
```go
err := Func()
if err != nil {
    log.Println(err)
}
```
使用herror处理err
```go
//如果err != nil，则抛出panic
HandleErr(Func()).Must()

HandleErr(Func()).Msg("err msg").Must()

//仅记录err
HandleErr(Func()).Ignore()
```
当调用的函数拥有两个返回值时，第二个返回值通常是err，根据这个特点，herror在内部处理err并返回结果。
```go
result := ResultErr(Func()).Must()

result := ResultErr(Func()).Msg("err msg").Must()

result := ResultErr(Func()).Ignore()
```
如果调用的函数拥有三个返回值，可以使用ResultsErr函数。
```go
r1, r2 := ResultsErr(Func()).Must()

r1, r2 := ResultsErr(Func()).Ignore()
```
## TODO
- 增加自定义的err处理方法
- 添加英文描述