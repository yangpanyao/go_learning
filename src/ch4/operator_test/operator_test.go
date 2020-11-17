package operator_test

import "testing"

func TestCompareArray(t *testing.T)  {
	a := [...]int{1,2,3,4}
	b := [...]int{1,3,4,5}
	//c := [...]int{1,2,3,4,5}
	d := [...]int{1,2,3,4}
	t.Log(a==b)
	//t.Log(a==c)
	t.Log(a==d)
}
/*
运行结果
=== RUN   TestCompareArray
    TestCompareArray: operator_test.go:10: false
    TestCompareArray: operator_test.go:12: true
--- PASS: TestCompareArray (0.00s)
 */
