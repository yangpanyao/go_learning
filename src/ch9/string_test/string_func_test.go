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
