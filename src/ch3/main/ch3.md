基本数据类型

|   |   |
| ------------ | ------------ |
| bool  | 布尔型  |
| string  | 字符串型  |
| int,int8,int16,int32,int64 |  整型 |
| unit,unit8,unit16,uint32,unit64,uintptr  | 无符号整型  |
| byte //alias uint8 | 字节类型 别名无符号8位整型  |
| rune //alias int32  |  |
| float32,float64  |  浮点型 |
| complex64,complex128  | 32 位实数和虚数,64 位实数和虚数  |

类型转换 与其他主流编程语言的差异

1. Go语言不允许隐式类型转换

错误示例：
```
package type_test

import "testing"


func TestImplicit(t *testing.T)  {
	var a int = 1
	var b int64
}
/**
运行结果
# command-line-arguments_test [command-line-arguments.test]
./type_test.go:8:4: cannot use a (type int) as type int64 in assignment
*/
```
正确示例:
```
package type_test

import "testing"

func TestImplicit(t *testing.T)  {
	var a int32 = 1
	var b int64
	b = int64(a)
	t.Log(a,b)
}
/**
运行结果
=== RUN   TestImplicit
    TestImplicit: type_test.go:9: 1 1
--- PASS: TestImplicit (0.00s)
*/
```


2. 别名和原有类型也不能进行隐式转换
错误示例：

```
package type_test

import "testing"

type MyInt int64
func TestImplicit(t *testing.T)  {
	var a int32 = 1
	var b int64
	b = int64(a)
	var c  MyInt
	c = b
	t.Log(a,b)
}
/**
运行结果
# command-line-arguments_test [command-line-arguments.test]
./type_test.go:11:4: cannot use b (type int64) as type MyInt in assignment
*/
```

正确示例：
```
package type_test

import "testing"

type MyInt int64
func TestImplicit(t *testing.T)  {
	var a int32 = 1
	var b int64
	b = int64(a)
	var c  MyInt
	c = MyInt(b)
	t.Log(a,b,c)
}

/**
运行结果：
=== RUN   TestImplicit
    TestImplicit: type_test.go:12: 1 1 1
--- PASS: TestImplicit (0.00s)
PASS
*/
```

类型的预定义值
1. math.MaxInt64
2. math.MaxFloat64
3. math.MaxUint64

指针类型 与其他主流编程语言的差异
1. 不支持指针运算
2. string是值类型， 其默认的为初始化值为空字符串，而不是nil



