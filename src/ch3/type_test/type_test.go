package type_test

import (
	"testing"
)

//Go语言不允许隐式类型转换  别名和原有类型也不能进行隐式转换
type MyInt int64//别名
func TestImplicit(t *testing.T)  {
	var a int32 = 1
	var b int64
	b = int64(a)
	var c  MyInt
	c = MyInt(b)
	t.Log(a,b,c)
}

func TestPoint(t *testing.T)  {
	a := 1
	aPtr := &a //获取a在内存中的指针
	t.Log(a,aPtr)
	t.Logf("%T,%T",a,aPtr)

}
/**
运行结果
=== RUN   TestPoint
    TestPoint: type_test.go:20: 1 0xc0000a6198
    TestPoint: type_test.go:21: int,*int
--- PASS: TestPoint (0.00s)
PASS
 */

//在go语言中 字符串默认的初始化值为空字符串
func TestString(t *testing.T)  {
	var s string
	t.Log("*"+s+"*")
	t.Log(len(s))
}
/**
=== RUN   TestString
TestString: type_test.go:35: **
--- PASS: TestString (0.00s)

 */
