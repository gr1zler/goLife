package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First  string `json:"Firt"`
	Second string
	Age    int
}

func main() {
	s := `[{"Firt":"Aa","Second":"Bb","Age":1},{"Firt":"Cc","Second":"Dd","Age":2}]`
	bs := []byte(s)
	p := []person{}
	err := json.Unmarshal(bs, &p)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(p)
}
