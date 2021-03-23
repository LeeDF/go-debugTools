package funcPoint

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
	funcPoint.Keep()
	time.Sleep(time.Millisecond)
	Func2(funcPoint)
	funcPoint.Keep()
	time.Sleep(time.Millisecond)
	Func3(funcPoint)
	funcPoint.Keep()
}

func Func2(funcPoint *FuncPoint) {
	funcPoint.Keep()
	Func3(funcPoint)
	time.Sleep(time.Millisecond * 2)
	Func3(funcPoint)
	funcPoint.Keep()
}
func Func3(funcPoint *FuncPoint) {
	time.Sleep(time.Millisecond * 3)
	funcPoint.Keep()
}