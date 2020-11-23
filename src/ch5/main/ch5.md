条件和循环
1. 循环
Go语言仅支持循环关键字for
```
for (i := 7; i <= 9; i++)
```
Go语言for语句的程序体不需要 `()`

示例
```go
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

```

2. if条件语句

```
if 布尔表达式 {
   /* 在布尔表达式为 true 时执行 */
}

if 布尔表达式 {
   /* 在布尔表达式为 true 时执行 */
} else {
   /* 在布尔表达式为 false 时执行 */
}

if 布尔表达式1 {
   /* 在布尔表达式1为 true 时执行 */
} else if 布尔表达式2{
   /* 在布尔表达式2为 true 时执行 */
} else {
   /* 在两个布尔表达式为 false 时执行 */
}

```

if条件与其他主要编程语言的差异
1. 表达式结果必须为布尔值
2. 支持变量复制

```
if var declaration; condition {
    //code to be executed if condition is true
}
```

示例：
```go
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
```

3. switch 条件

```
switch os := runtime.GOOS; os {
    case "darwin":
        fmt.Println("OS .X")
        //break
    case "linux":
        fmt.Println("linux")
    default:
        //freebsd,openbsd,
        //plan9,windows...
        fmt.Printf("%s.",os)

}
```

```
switch {
    case 0 <= Num && Num <= 3:
        fmt.Printf("0-3");
    case 4 <= Num && Num <= 6:
        fmt.Printf("0-3");
    case 7 <= Num && Num <= 9:
        fmt.Printf("0-3");
}
```
switch 条件与其他主要编程语言的差异

1. 条件表达式不限制为常量或者整数
2. 单个case中可以出现多个结果选项，使用逗号分割；
3. 与C语言等规则相反，Go语言不需要声明break来明确退出一个case
4. 可以不设定switch之后的条件表达式，在此种情况下，整个switch 结构与if... else...的逻辑作用相同

示例：
```
func TestSwitchMultiCase(t *testing.T) {
	for i:=0; i<5;i++ {
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
```

