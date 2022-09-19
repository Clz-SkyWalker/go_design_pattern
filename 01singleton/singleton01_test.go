package main

import (
	"fmt"
	"sync"
	"testing"
)

type OnlyOne struct {
	name string
}

// 饿汉式
var Single = OnlyOne{name: "t1"}
var single1 = OnlyOne{name: "t2"}
var Single3 OnlyOne

func init() {
	Single3 = OnlyOne{name: "t3"}
}

func GetOnlyOne() OnlyOne {
	return single1
}

func TestSingle1(t *testing.T) {
	s := Single
	fmt.Println(s.name)
	s = GetOnlyOne()
	fmt.Println(s.name)
	s = Single3
	fmt.Println(s.name)
}

// 饿汉式
var mutex sync.Mutex
var single2 *OnlyOne

func GetOnlyOne2() *OnlyOne {
	mutex.Lock()
	if single2 == nil {
		single2 = new(OnlyOne)
		single2.name = "t2"
	}
	mutex.Unlock()
	return single2
}

func TestSingle2(t *testing.T) {
	s := GetOnlyOne2()
	fmt.Println(s.name)
}

func GetOnlyOne3() *OnlyOne {
	if single2 == nil {
		mutex.Lock()
		if single2 == nil {
			single2 = new(OnlyOne)
			single2.name = "t3"
		}
		mutex.Unlock()
	}
	return single2
}

func TestSingle3(t *testing.T) {
	s := GetOnlyOne3()
	fmt.Println(s.name)
}
