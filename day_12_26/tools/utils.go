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
	//reflectPrintFlag(s)
	OperaMethod(&s)

}

func OperaMethod(s interface{}) {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("Not a STRUCT")
		return
	}
	//fmt.Println(v.FieldByName("Name"))
	//fmt.Println(v.FieldByName("Age"))
	//fmt.Println(v.FieldByName("Name"))

	numMethod := t.NumMethod()
	fmt.Println(numMethod)
	method0 := t.Method(0)
	fmt.Println(method0.Name)
	values := v.MethodByName("GetInfo").Call(nil)
	fmt.Printf("%#v\n", values)
	for k, v := range values {
		fmt.Println(k, v)
	}
	fmt.Println(values)
	var params []reflect.Value
	params = append(params, reflect.ValueOf(1006))
	params = append(params, reflect.ValueOf("wsssss"))
	params = append(params, reflect.ValueOf(987))
	v.MethodByName("SetInfo").Call(params)
	values = v.MethodByName("GetInfo").Call(nil)
	fmt.Println(values)
	name := v.Elem().FieldByName("Name")
	fmt.Println(name)
	name.SetString("挖庆生")
	fmt.Println(name)

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
