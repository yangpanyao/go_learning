package array_test

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr[1] = 2
	arr1 := [3]int{1,2,3}
	arr2 := [...]int{1,2,3,4,5}
	arr3 := [2][2]int{{1,2},{3,4}}
	t.Log(arr)
	t.Log(arr1)
	t.Log(arr2)
	t.Log(arr3)
}
/*
运行结果
=== RUN   TestArrayInit
    TestArrayInit: array_test.go:11: [0 2 0]
    TestArrayInit: array_test.go:12: [1 2 3]
    TestArrayInit: array_test.go:13: [1 2 3 4 5]
    TestArrayInit: array_test.go:14: [[1 2] [3 4]]
--- PASS: TestArrayInit (0.00s)
 */

func TestArrayTravel(t *testing.T)  {
	arr := [...]int{1,2,3,4,5}
	//使用for对数组进行遍历
	//第一种
	for i := 0; i < len(arr); i++ {
		t.Log(arr[i])
	}
	//第二种
	//也可使用 for idx,e := range arr {} 的结构对数组进行遍历 其中 idx为数组的索引 e为索引对应的值
	for key, val := range arr{
		t.Log(key,val)
	}
}
/*
运行结果
=== RUN   TestArrayTravel
    TestArrayTravel: array_test.go:31: 1
    TestArrayTravel: array_test.go:31: 2
    TestArrayTravel: array_test.go:31: 3
    TestArrayTravel: array_test.go:31: 4
    TestArrayTravel: array_test.go:31: 5
    TestArrayTravel: array_test.go:36: 0 1
    TestArrayTravel: array_test.go:36: 1 2
    TestArrayTravel: array_test.go:36: 2 3
    TestArrayTravel: array_test.go:36: 3 4
    TestArrayTravel: array_test.go:36: 4 5
--- PASS: TestArrayTravel (0.00s)
 */

//数组截取
func TestArraySection(t *testing.T)  {
	arr := [...]int{1,2,3,4,5}
	arrSec := arr[3:]
	t.Log(arrSec)
}

/*
运行结果
=== RUN   TestArraySection
    TestArraySection: array_test.go:58: [4 5]
--- PASS: TestArraySection (0.00s)
PASS
 */

