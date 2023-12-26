package tools

import (
	"fmt"
	"strconv"
)

type Student struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (s Student) getInfo() map[string]string {
	result := make(map[string]string, 3)
	result["id"] = strconv.Itoa(s.Id)
	result["name"] = s.Name
	result["age"] = strconv.Itoa(s.Age)
	return result
}

func (s *Student) setInfo(id int, name string, age int) {
	s.Id = id
	s.Name = name
	s.Age = age
}

func (s Student) printS() {
	fmt.Println(s)
}
