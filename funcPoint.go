package go_debugTools

import (
	"container/list"
	"fmt"
	"runtime"
	"sort"
	"strings"
	"time"
)

type point struct {
	FuncName    string
	Line        int
	File        string
	Time        time.Time
	TimeConsume time.Duration
}

type FuncPoint struct {
	Label string
	List  *list.List
	Count map[string][2]int
}

func (t *FuncPoint) Keep() {
	funcName, file, line, _ := runtime.Caller(1)
	sp := point{
		FuncName: runtime.FuncForPC(funcName).Name(),
		Line:     line,
		File:     file,
		Time:     time.Now(),
	}
	if t.List.Back() != nil {
		sp.TimeConsume = sp.Time.Sub(t.List.Back().Value.(point).Time)
	}
	t.List.PushBack(sp)

	ck := fmt.Sprintf("%s:%d", file, line)
	if c, ok := t.Count[ck]; ok {
		t.Count[ck] = [2]int{c[0] + 1, c[1] + int(sp.TimeConsume.Nanoseconds())}
	} else {
		t.Count[ck] = [2]int{1, int(sp.TimeConsume.Nanoseconds())}
	}
}

func (t *FuncPoint) String() string {
	ans := strings.Builder{}
	ans.WriteString(t.Label + "\n")
	for p := t.List.Front(); p != t.List.Back(); p = p.Next() {
		s := p.Value.(point)
		ans.WriteString(fmt.Sprintf("  [%s] file: %s:%d func:%s consume:%d ns \n", s.Time.Format("2006-01-02 15:04:05"), s.File, s.Line, s.FuncName, s.TimeConsume.Nanoseconds()))
	}

	cList := make([]struct {
		name     string
		count    int
		duration int
	}, 0, len(t.Count))
	for k, v := range t.Count {
		cList = append(cList, struct {
			name     string
			count    int
			duration int
		}{name: k, count: v[0], duration: v[1]})
	}
	sort.Slice(cList, func(i, j int) bool {
		return cList[i].duration > cList[j].duration
	})
	ans.WriteString("count \n")
	for _, v := range cList {
		ans.WriteString(fmt.Sprintf("  %s: sum %d sumDuration %d ns avgDuration %d ns \n", v.name, v.count, v.duration, v.duration/v.count))
	}
	return ans.String()
}

func New(label string) *FuncPoint {
	return &FuncPoint{
		Label: label,
		List:  list.New(),
		Count: make(map[string][2]int),
	}
}
