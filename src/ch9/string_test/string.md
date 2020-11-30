字符串
与其他主要编程语⾔的差异
1. string 是数据类型，不是引⽤或指针类型
2. string 是只读的 byte slice，len 函数可以它所包含的 byte 数
3. string 的 byte 数组可以存放任何数据

Unicode UTF8
1. Unicode 是⼀种字符集（code point）
2. UTF8 是 unicode 的存储实现 （转换为字节序列的规则）

编码与存储
字符 : “中”
Unicode : 0x4E2D
UTF-8 : 0xE4B8AD
string/[]byte : [0xE4,0xB8,0xAD]

常⽤字符串函数
1. strings 包 (https://golang.org/pkg/strings/)
2. strconv 包 (https://golang.org/pkg/strconv/)

示例
```go
package string_test

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringFn(t *testing.T)  {
	s := "A,B,C"
	//字符串分割
	parts := strings.Split(s,",")
	for _,part := range parts{
		t.Log(part)
	}
	//字符串连接
	t.Log(strings.Join(parts,"-"))
}
/*
=== RUN   TestStringFn
    TestStringFn: string_func_test.go:14: A
    TestStringFn: string_func_test.go:14: B
    TestStringFn: string_func_test.go:14: C
    TestStringFn: string_func_test.go:17: A-B-C
--- PASS: TestStringFn (0.00s)
 */
//类型转换
func TestConv(t *testing.T)  {
	s := strconv.Itoa(10)
	t.Log("str"+s)
	if i,err := strconv.Atoi("10"); err == nil {
		t.Log(10 + i)
	}
}
/*
=== RUN   TestConv
    TestConv: string_func_test.go:30: str10
    TestConv: string_func_test.go:32: 20
--- PASS: TestConv (0.00s)
 */

```




