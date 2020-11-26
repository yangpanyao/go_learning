Map 的 value 可以是⼀个⽅法
与 Go 的 Dock type 接⼝⽅式⼀起，可以⽅便的实现单⼀⽅法对象的⼯⼚模式

示例：
```go
package map_ext

import "testing"
//map 的value 可以是一个方法
func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int {return op}
	m[2] = func(op int) int {return op*op}
	m[3] = func(op int) int {return op*op*op}
	t.Log(m[1](2),m[2](2),m[3](2))

}
/**
运行结果：
=== RUN   TestMapWithFunValue
    TestMapWithFunValue: map_ext_test.go:10: 2 4 8
--- PASS: TestMapWithFunValue (0.00s)
PASS
 */
```


Go 的内置集合中没有 Set 实现， 可以 map[type]bool
1. 元素的唯⼀性
2. 基本操作
1) 添加元素
2) 判断元素是否存在
3) 删除元素
4) 元素个数

```go
package map_ext
import "testing"
//Go 的内置集合中没有 Set 实现， 可以 使用 map[type]bool 实现
func TestMapForSet(t *testing.T)  {
	mySet := map[int]bool{}
	mySet[1] = true

	n := 3
	if mySet[n] {
		t.Logf("%d is exist",n)
	} else {
		t.Logf("%d is not exist", n)
	}
	mySet[3] = true
	delete(mySet,1)
	t.Log(len(mySet))

}
/*
运行结果：
=== RUN   TestMapForSet
    TestMapForSet: map_ext_test.go:29: 3 is not exist
    TestMapForSet: map_ext_test.go:33: 1
--- PASS: TestMapForSet (0.00s)
 */

```