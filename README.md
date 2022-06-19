# herror 是一个简单处理错误的包
## 介绍
go语言的错误处理的方式一直让人吐槽，每写一行代码就要 if err != nil 一下,非常的痛苦。herror就是为了解决这个痛点，在一些不太需要关注错误的地方你可以使用它，比如快速的编写脚本。
## 快速开始
```go
package main

import "github.com/leijux/herror"

func F() error {
    return errors.New("err")
}

func main() {
    herror.HandleErr(F()).Must()
}
```
## 示例
对于返回的err我们通常使用如下方法处理
```go
err := Func()
if err != nil {
    log.Fatalln(err)
}
```
使用herror处理err
```go
HandleErr(Func()).Must()

HandleErr(Func()).Msg("err msg").Must()

//忽略err
HandleErr(Func()).Ignore()
```
当函数拥有两个返回值时，通常第二个参数是err，根据这个特性herror将在内部处理err并返回结果。
```go
result := ResultErr(Func()).Must()

result := ResultErr(Func()).Msg("err msg").Must()

result := ResultErr(Func()).Ignore()
```
如果拥有三个参数可以使用ResultsErr函数
```go
result1, result2 := ResultsErr(Func()).Must()

result1, result2 := ResultsErr(Func()).Ignore()
```
## TODO
- 增加自定义的err处理方法