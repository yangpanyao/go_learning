Go 语言切片是对数组的抽象。
Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go中提供了一种灵活，功能强悍的内置类型切片("动态数组"),与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。

切片的声明
```

var s0 []int
s0 = append(s0, 1)
s := []int{}
s1 := []int{1, 2, 3}
s2 := make([]int, 2, 4) 
/*[]type, len, cap
其中len个元素会被初始化为默认零值，未初始化元素不可以访问
 */

```

```go
package slice_test

import "testing"

func TestSliceInit(t *testing.T)  {
	var s0  []int
	t.Log(len(s0),cap(s0))
	s0 = append(s0,1)
	t.Log(len(s0),cap(s0))

	s1 := []int{1,2,3,4}
	t.Log(len(s1),cap(s1))

	s2 := make([]int,3,5)
	t.Log(len(s2),cap(s2))
	s2 = append(s2,1)
	t.Log(len(s2),cap(s2))
}
/*
运行结果
=== RUN   TestSliceInit
    TestSliceInit: slice_test.go:7: 0 0
    TestSliceInit: slice_test.go:9: 1 1
    TestSliceInit: slice_test.go:12: 4 4
    TestSliceInit: slice_test.go:15: 3 5
    TestSliceInit: slice_test.go:17: 4 5
--- PASS: TestSliceInit (0.00s)
PASS
 */
```

切片的容量 会随着 切片元素的增长以 原有容量乘2的速度增长

```go
package slice_test
import "testing"
func TestSliceGrowing(t *testing.T)  {
	var s []int
	for i := 0; i < 10; i++{
		s = append(s,i)
		t.Log(len(s),cap(s))
	}
}
/*
运行结果
=== RUN   TestSliceGrowing
    TestSliceGrowing: slice_test.go:24: 1 1
    TestSliceGrowing: slice_test.go:24: 2 2
    TestSliceGrowing: slice_test.go:24: 3 4
    TestSliceGrowing: slice_test.go:24: 4 4
    TestSliceGrowing: slice_test.go:24: 5 8
    TestSliceGrowing: slice_test.go:24: 6 8
    TestSliceGrowing: slice_test.go:24: 7 8
    TestSliceGrowing: slice_test.go:24: 8 8
    TestSliceGrowing: slice_test.go:24: 9 16
    TestSliceGrowing: slice_test.go:24: 10 16
--- PASS: TestSliceGrowing (0.00s)
PASS
 */

```


切片采用共享储存结构 当 新截取的切片内部发生改变时，被截取的切片内部也会随之改变

```go
package slice_test
import "testing"
func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan","Feb","Mar","Apr","May","Jun","Jul","Aug","Sep","Oct","Nov","Dec"}

	Q2 := year[3:6]
	t.Log(Q2,len(Q2),cap(Q2))
	summer := year[5:8]
	t.Log(summer,len(summer),cap(summer))
	summer[0] = "unknown"
	t.Log(Q2)
	t.Log(year)
}
/*
运行结果
=== RUN   TestSliceShareMemory
    TestSliceShareMemory: slice_test.go:62: [Apr May Jun] 3 9
    TestSliceShareMemory: slice_test.go:64: [Jun Jul Aug] 3 7
    TestSliceShareMemory: slice_test.go:66: [Apr May unknown]
    TestSliceShareMemory: slice_test.go:67: [Jan Feb Mar Apr May unknown Jul Aug Sep Oct Nov Dec]
--- PASS: TestSliceShareMemory (0.00s)
 */
```

