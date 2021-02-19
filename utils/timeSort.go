package utils

import "time"

type TimeSort struct {
	Slice []interface{}               //承载以任意结构体为元素构成的Slice
	By    func(a, b interface{}) bool //排序规则函数,当需要对新的结构体slice进行排序时，只需定义这个函数即可
}

func (t TimeSort) Len() int { return len(t.Slice) }

func (t TimeSort) Swap(i, j int) { t.Slice[i], t.Slice[j] = t.Slice[j], t.Slice[i] }

func (t TimeSort) Less(i, j int) bool { return t.By(t.Slice[i], t.Slice[j]) }

type FileInfo struct {
	name string    `json:"name"`
	time time.Time `json:"time"`
}

func timeBy(a, b interface{}) bool {
	return a.(FileInfo).time.After(b.(FileInfo).time)
}
