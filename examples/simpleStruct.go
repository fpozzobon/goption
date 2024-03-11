package main

import (
	"encoding/json"
	"fmt"
	"github.com/fpozzobon/goption"
	"time"
)

type myStruct struct {
	Name string
	Date time.Time
}

func main() {
	myOption := getOption()

	res, err := json.Marshal(myOption)
	if err != nil {
		panic(err)
	}
	resStr := string(res)
	fmt.Println(resStr)
}

func getOption() goption.GOption[myStruct] {
	s := myStruct{
		Name: "test",
	}
	return goption.Some(s)
}
