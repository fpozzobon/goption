package goption

import (
	"encoding/json"
	"testing"
	"time"
)

type myStruct struct {
	Name string
	Date time.Time
}

func initStruct() myStruct {
	return myStruct{
		Name: "test",
		Date: time.Now(),
	}
}

func getEmptyStruct() myStruct {
	return myStruct{}
}

func getPtr() *myStruct {
	s := initStruct()
	return &s
}

func getNilPtr() *myStruct {
	var s myStruct
	return &s
}

func getOption() GOption[myStruct] {
	return Some(initStruct())
}

func getEmpty() GOption[myStruct] {
	return None[myStruct]()
}

func BenchmarkPointer(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ptr := getPtr()
			json.Marshal(ptr)
			nilPtr := getNilPtr()
			json.Marshal(nilPtr)
		}
	})
}

func BenchmarkOption(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			val := getOption()
			json.Marshal(val)
			empty := getEmpty()
			json.Marshal(empty)
		}
	})
}

func BenchmarkValue(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			val := initStruct()
			json.Marshal(val)
			empty := getEmptyStruct()
			json.Marshal(empty)
		}
	})
}
