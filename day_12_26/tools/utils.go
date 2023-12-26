package tools

import (
	"fmt"
	"reflect"
)

func init() {

}

func ReflectStructTest() {
	s := Student{1001, "wss", 26}
	fmt.Println(s.getInfo())
	s.setInfo(1002, "ws", 18)
	fmt.Println(s.getInfo())
	s.printS()
	fmt.Printf("%#v\n", s)
	reflectPrintFlag(s)

}

func reflectPrintFlag(s Student) {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("Not a STRUCT")
		return
	}

	filed0 := t.Field(0)
	fmt.Println(filed0.Name)
	fmt.Println(filed0.Tag.Get("json"))
	fmt.Println(filed0.Type)

	filedName, ok := t.FieldByName("Name")
	if ok {
		fmt.Println(filedName.Name)
		fmt.Println(filedName.Tag.Get("json"))
		fmt.Println(filedName.Type)
	}

	fmt.Println(t.NumField())

	fmt.Println(v.FieldByName("Name"))
	fmt.Println(v.FieldByName("Age"))
	fmt.Println(v.FieldByName("Id"))
}
