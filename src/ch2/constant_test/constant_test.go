package constant_test

import "testing"

//go语言的常量 可以快速的设置连续的值
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

func TestConstantTry(t *testing.T)  {
	t.Log(Monday,Tuesday,Wednesday)
}
/**
运行结果：
=== RUN   TestConstant
    TestConstant: constant_test.go:12: 1 2 3
--- PASS: TestConstant (0.00s)
 */


func TestConstantTry1(t *testing.T)  {
	//a := 7 // 0111
	//a := 8 // 1000
	a := 9 //1001
	//比较a与比较值的比特位的值是否相等
	t.Log(Readable,Writeable,Executable)
	t.Log(a&Readable,a&Writeable,a&Executable)
	t.Log(a&Readable == Readable,a&Writeable ==Writeable, a&Executable == Executable)
}
/**
运行结果
=== RUN   TestConstantTry1
    TestConstantTry1: constant_test.go:34: 1 2 4
    TestConstantTry1: constant_test.go:35: 1 0 0
    TestConstantTry1: constant_test.go:36: true false false
--- PASS: TestConstantTry1 (0.00s)
 */