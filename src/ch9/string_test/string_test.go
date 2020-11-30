package string_test

import (
	"testing"
	"unsafe"
)

func TestString(t *testing.T)  {
	var s string
	t.Log(s)
	s = "hello"
	t.Log(len(s))

	s = "\xE4\xB8\xA5" //可以储存任何二进制数据
	t.Log(s)

	s = "中"

	t.Log(len(s)) //是byte数

	c := []rune(s)//rune 可以取出字符串中的unicode
	t.Log(len(c))
	t.Log("rune size:", unsafe.Sizeof(c[0]))
	t.Logf("中 unicode %x",c[0])
	t.Logf("中 utf8 %x",s)

}
/*
运行结果
=== RUN   TestString
    TestString: string_test.go:10:
    TestString: string_test.go:12: 5
    TestString: string_test.go:15: 严
    TestString: string_test.go:19: 3
    TestString: string_test.go:22: 1
    TestString: string_test.go:23: rune size: 4
    TestString: string_test.go:24: 中 unicode 4e2d
    TestString: string_test.go:25: 中 utf8 e4b8ad
 */


func TestStringToRune(t *testing.T)  {
	s := "代码改变世界"
	for _,v := range s{
		t.Logf("%[1]c,%[1]x",v)
	}
}
/*
运行结果：
=== RUN   TestStringToRune
    TestStringToRune: string_test.go:45: 代,4ee3
    TestStringToRune: string_test.go:45: 码,7801
    TestStringToRune: string_test.go:45: 改,6539
    TestStringToRune: string_test.go:45: 变,53d8
    TestStringToRune: string_test.go:45: 世,4e16
    TestStringToRune: string_test.go:45: 界,754c
--- PASS: TestStringToRune (0.00s)
 */
