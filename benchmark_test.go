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

func getPtr() *myStruct {
	s := initStruct()
	return &s
}

func getOption() GOption[myStruct] {
	return Some(initStruct())
}

func BenchmarkPointer(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ptr := getPtr()
			json.Marshal(ptr)
		}
	})
}

func BenchmarkOption(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			val := getOption()
			json.Marshal(val)
		}
	})
}

func BenchmarkValue(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			val := initStruct()
			json.Marshal(val)
		}
	})
}
