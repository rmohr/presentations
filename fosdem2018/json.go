package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	obj := &User{ID: "123", Name: "John", Age: 23}
	data, _ := json.MarshalIndent(obj, "  ", "")
	fmt.Println(string(data))
}
