package main

import (
	"fmt"
	"takeout/parser/model"
)

func main() {
	var dataMap = make(map[string] model.TakeOutSummary);
	dataMap["a"] = model.TakeOutSummary{};
	fmt.Println(dataMap["a"].StartTs);
}