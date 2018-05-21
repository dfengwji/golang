// reflection
package step

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

type Manager struct {
	User
	Title string
}

func (this User) Hello(n string) {
	fmt.Println("hello ", n, ",my name is ", this.Name)
}

func Info(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("Type:", t.Name())

	if k := t.Kind(); k != reflect.Struct {
		fmt.Printf("not match!!!")
		return
	}

	v := reflect.ValueOf(o)
	fmt.Println("fields:")

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s:%v = %v\n", f.Name, f.Type, val)
	}

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%6s:%v\n", m.Name, m.Type)
	}
}

func Set(o interface{}) {
	v := reflect.ValueOf(o)
	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("set not match")
	} else {
		v = v.Elem()
	}
	f := v.FieldByName("Name")
	if !f.IsValid() {
		fmt.Println("bad")
		return
	}
	if f.Kind() == reflect.String {
		f.SetString("lucy")
	}
}

func StudyReflect() {
	u := User{1, "jon", 12}
	Info(u)
	Set(&u)
	fmt.Println(u)

	f := reflect.ValueOf(u)
	mv := f.MethodByName("Hello")
	args := []reflect.Value{reflect.ValueOf("robot")}
	mv.Call(args)

	m := Manager{User: User{1, "Sum", 45}, Title: "leader"}
	t := reflect.TypeOf(m)
	fmt.Printf("%#v\n", t.Field(1))
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0: 1}))

	x := 123
	v := reflect.ValueOf(&x)
	v.Elem().SetInt(999)
	fmt.Println(x)

}
