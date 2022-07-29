package main

import (
	"encoding/json"
	"fmt"
)

type Company_5_7 struct {
	Name   string   `json:"name"`
	Number int      `json:"number"`
	Dep    []string `json:"-"`
	IsOk   bool     `json:"is_ok"`
	Price  float64
}

func main() {
	comp := Company_5_7{"大厦", 3, []string{"人事", "研发", "产品"}, true, 3.14}
	//j, err := json.Marshal(comp)
	j, err := json.MarshalIndent(comp, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(j))

	var comp2 Company_5_7
	err2 := json.Unmarshal(j, &comp2)
	if err2 != nil {
		return
	}
	fmt.Println(comp2)

}
