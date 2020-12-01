package func_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)
//可以有多个返回值
func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func TestFn(t *testing.T)  {
	//多个返回值
	a,b :=returnMultiValues()
	t.Log(a,b)
	//求运行时长
	tsSF :=TimeSpent(slowFn)
	t.Log(tsSF(10))
}
//封装一个timeSpent来得到程序运行时常
func TimeSpent(inner func(op int) int) func(op int) int  {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:",time.Since(start).Seconds())
		return ret
	}
}

func slowFn(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestVarParam(t *testing.T)  {
	t.Log(Sum(1,2,3,4,5))
}



//可变长参数
func Sum(ops ...int) int  {
	ret := 0
	for _,op := range ops{
		ret += op
	}
	return ret
}


//defer 函数延迟执行 通常用于释放清理资源等
func TestDefer(t *testing.T)  {
	defer Clear()
	fmt.Println("start")
	panic("Fatal error")//defer仍会执行
}

func Clear()  {
	fmt.Println("Clear resources")
}