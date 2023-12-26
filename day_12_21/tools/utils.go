package tools

import (
	"fmt"
	"reflect"
)

type P interface {
	pName() string
}

type H struct {
	name string
}

func (h H) p() {
	fmt.Println(h.name)
}

func (h H) pName() string {
	fmt.Println(h.name)
	return h.name
}

func ReflectTest() {
	//var pp P
	a := 1
	//b := "dada"
	//c := "f"
	//d := 10.22
	//e := H{"ws"}
	//e.p()
	//pp = e
	//pp.pName()
	//
	//re(pp)
	//re(e)
	//re(a)
	//re(b)
	//re(c)
	//re(d)
	//re(e)
	//re(&a)
	//f := []int{1, 2, 3}
	//g := [3]string{"dasd", "Dasda"}
	//re(f)
	//re(g)
	//re(re)
	reflectSetValue(&a)
	fmt.Println(a)

}

func re(data interface{}) {

	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)
	fmt.Printf("%v===%v===%v===%v===%v\n", t, v, t.Name(), t.Kind(), v.Kind())
	n, _ := data.(int)
	//data += 5
	var num = 5 + n
	fmt.Printf("num===%v\n", num)

	var num2 = 5 + v.Int()
	fmt.Printf("num2===%v\n", num2)

}

func reflectSetValue(data interface{}) {
	t := reflect.TypeOf(data)
	fmt.Println(t.Name(), t.Kind())
	v := reflect.ValueOf(data)
	fmt.Println(v, v.Kind())
	if v.Elem().Kind() == reflect.Int {
		v.Elem().SetInt(123)
	}
	fmt.Println(v, v.Kind())

}
