package fn_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)


func returnMultiValues() (int,int){
	return rand.Intn(10),rand.Intn(20)
}

func timeSpend(inner func(op int) int) func (op int) int {
	return func(n int) int{
		start:= time.Now()
		ret:=inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return  ret
	}
}


func slowFun(op int) int{
	time.Sleep(time.Second*1)
	return op
}

func TestFn(t *testing.T){
	a,_ :=returnMultiValues()
	t.Log(a)
	tsSF:=timeSpend(slowFun)
	t.Log(tsSF(1))
}

func Sum(ops ...int) int{
	ret :=0
	for _,op:=range ops{
		ret += op
	}
	return ret
}

func TestVarParam(t *testing.T){
	t.Log(Sum(1,2,3,4))
	t.Log(Sum(1,2,3,4,5))
}


func Clear(){
	fmt.Println("Clear Resource")
}


func TestDefer(t *testing.T){
	defer Clear()
	fmt.Println("start")
	panic("err")
	fmt.Println("End")


}