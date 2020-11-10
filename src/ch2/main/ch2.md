测试程序
1. 测试文件以`_test`结尾：`xxx_test.go` 
2. 测试方法名以`Test`开头：`func TestXXX(t *testing.T) {...}`

变量赋值

1. 赋值可以自动进行类型推断
2. 在一个复制语句中可以对多个变量同时进行赋值

常量的定义
快速设置连续的值
```
const (
	Monday = iota + 1
	Tuesday
	Wednesday
)

const (
	Readable = 1<<iota
	Writeable
	Executable
)
```