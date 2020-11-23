package loop

import "testing"

func TestWhileLoop(t *testing.T)  {
	i := 0
	/* 等价于 while(i<5) */
	for i < 5  {
		t.Log(i)
		i++
	}
}
/**
运行结果：
=== RUN   TestWhileLoop
    TestWhileLoop: loop_test.go:8: 0
    TestWhileLoop: loop_test.go:8: 1
    TestWhileLoop: loop_test.go:8: 2
    TestWhileLoop: loop_test.go:8: 3
    TestWhileLoop: loop_test.go:8: 4
--- PASS: TestWhileLoop (0.00s)
 */
