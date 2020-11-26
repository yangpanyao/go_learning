Map 是一种无序的键值对的集合。Map 最重要的一点是通过 key 来快速检索数据，key 类似于索引，指向数据的值。

Map 是一种集合，所以我们可以像迭代数组和切片那样迭代它。不过，Map 是无序的，我们无法决定它的返回顺序，这是因为 Map 是使用 hash 表来实现的。

Map 声明
```

m := map[string]int{"one": 1, "two": 2, "three": 3}
m1 := map[string]int{}
m1["one"] = 1
m2 := make(map[string]int, 10 /*Initial Capacity*/)//第二个参数为map的容量

```

示例：
```go
package map_test

import "testing"

func TestInitMap(t *testing.T) {
	m1 := map[string]int{"one":1,"two":2,"three":3}
	t.Log(m1["one"])
	t.Logf("len m1= %d",len(m1))
	m2 := map[int]int{}
	m2[4] = 1
	t.Logf("len m2= %d",len(m2))

	m3 := make(map[int]int,10)
	t.Logf("len m3= %d",len(m3))

}
/*
运行结果：
=== RUN   TestInitMap
    TestInitMap: map_test.go:7: 1
    TestInitMap: map_test.go:8: len m1= 3
    TestInitMap: map_test.go:11: len m2= 1
    TestInitMap: map_test.go:14: len m3= 0
--- PASS: TestInitMap (0.00s)
 */
```

map在访问的 Key 不存在时，仍会返回零值，不能通过返回 nil 来判断元素是否存在
示例：
```go
package map_test

import "testing"

func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[2] = 1

	if v,ok := m1[3]; ok {
		t.Logf("key 3 value is %d",v)
	}else {
		t.Log("key 3 is not exist")
	}

}
/*
运行结果：
=== RUN   TestAccessNotExistingKey
    TestAccessNotExistingKey: map_test.go:21: 0
    TestAccessNotExistingKey: map_test.go:27: key 3 is not exist
--- PASS: TestAccessNotExistingKey (0.00s)
 */
```


遍历map
示例：
```go
package map_test

import "testing"

//遍历map
func TestMapTravel(t *testing.T) {
	m1 := map[string]int{"one":1,"two":2,"three":3}
	for k,v := range m1{
		t.Log(k,v)
	}
}

/*
运行结果：
=== RUN   TestMapTravel
    TestMapTravel: map_test.go:36: one 1
    TestMapTravel: map_test.go:36: two 2
    TestMapTravel: map_test.go:36: three 3
--- PASS: TestMapTravel (0.00s)
 */
```
