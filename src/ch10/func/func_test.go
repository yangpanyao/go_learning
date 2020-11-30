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
	a,b :=returnMultiValues()
	t.Log(a,b)

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