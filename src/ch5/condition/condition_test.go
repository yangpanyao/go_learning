package condition

import "testing"

func TestIfMultiSec(t *testing.T)  {
	if a := 1 == 1; a {
		t.Log("1==1")
	}
}

/*
运行结果：
=== RUN   TestIfMultiSec
    TestIfMultiSec: condition_test.go:7: 1==1
--- PASS: TestIfMultiSec (0.00s)
PASS
 */

func TestSwitchMultiCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0,2:
			t.Log("Even")
		case 1,3:
			t.Log("Odd")
		default:
			t.Log("it is not 0-3")

		}
	}
}
/*
运行结果
=== RUN   TestSwitchMultiCase
    TestSwitchMultiCase: condition_test.go:23: Even
    TestSwitchMultiCase: condition_test.go:25: Odd
    TestSwitchMultiCase: condition_test.go:23: Even
    TestSwitchMultiCase: condition_test.go:25: Odd
    TestSwitchMultiCase: condition_test.go:27: it is not 0-3
--- PASS: TestSwitchMultiCase (0.00s)
 */

func TestSwitchCaseCondition(t *testing.T)  {
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log("Even")
		case i%2 == 1:
			t.Log("Odd")
		default:
			t.Log("unknown")

		}
	}
}
/*
运行结果：
=== RUN   TestSwitchCaseCondition
    TestSwitchCaseCondition: condition_test.go:47: Even
    TestSwitchCaseCondition: condition_test.go:49: Odd
    TestSwitchCaseCondition: condition_test.go:47: Even
    TestSwitchCaseCondition: condition_test.go:49: Odd
    TestSwitchCaseCondition: condition_test.go:47: Even
--- PASS: TestSwitchCaseCondition (0.00s)
 */