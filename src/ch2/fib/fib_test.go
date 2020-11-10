package fib

import (
	"testing"
)
//斐波那契数
func TestFibList(t *testing.T)  {
	//var a int = 1
	//var b int = 1
	var (
		a int = 1
		b int = 1
	)
	t.Log("",a)
	for i :=0; i < 5; i++ {
		t.Log("",b)
		tmp := a
		a = b
		b = tmp + a
	}
}
/**
=== RUN   TestFibList
    TestFibList: fib_test.go:14:  1
    TestFibList: fib_test.go:16:  1
    TestFibList: fib_test.go:16:  2
    TestFibList: fib_test.go:16:  3
    TestFibList: fib_test.go:16:  5
    TestFibList: fib_test.go:16:  8
--- PASS: TestFibList (0.00s)
 */
/**
交换两个变量的值
 */
func TestExChange(t *testing.T)  {
	a := 1
	b := 2
	//tmp := a
	//a = b
	//b = tmp
	//t.Log(a,b)
	a,b = b,a
	t.Log(a,b)
}
/**
=== RUN   TestExChange
    TestExChange: fib_test.go:33: 2 1
--- PASS: TestExChange (0.00s)
 */