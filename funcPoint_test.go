package go_debugTools

import (
	"testing"
	"time"
)


func TestFuncPoint_Piling(t *testing.T) {
	point := New("test")
	Func1(point)
	t.Log(point.String())
}

func Func1(funcPoint *FuncPoint) {
	time.Sleep(time.Millisecond)
	Func2(funcPoint)
	time.Sleep(time.Millisecond)
	Func3(funcPoint)
}

func Func2(funcPoint *FuncPoint) {
	Func3(funcPoint)
	time.Sleep(time.Millisecond * 2)
	Func3(funcPoint)
}
func Func3(funcPoint *FuncPoint) {
	time.Sleep(time.Millisecond * 3)

}