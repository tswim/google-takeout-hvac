package main

import (
	"fmt"
	"os"
	"takeout/parser/model"
)

func main() {
	var dataMap = make(map[string] model.TakeOutSummary);
	dataMap["a"] = model.TakeOutSummary{};
	fmt.Println(dataMap["a"].Events);
	fmt.Println(dataMap["a"].Cycles);
	jsonFile, err := os.Open("users.json")

	if err != nil {
    	fmt.Println(err)
	}
	defer jsonFile.Close();
}