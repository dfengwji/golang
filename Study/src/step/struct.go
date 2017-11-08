// one
package step

import (
	"fmt"
)

type person struct {
	name    string
	age     int
	contact struct {
		phone, city string
	}
}

type human struct {
	name string
	sex  int
}

type teacher struct {
	human
	age int
}

func older(one, two person) (person, int) {
	if one.age > two.age {
		return one, one.age - two.age
	}
	return two, two.age - one.age
}

func StudyStruct() {
	testStruct2()
}

func testStruct() {
	var p person
	p.name = "alsex"
	p.age = 19
	fmt.Printf("first person name = %s and age = %d", p.name, p.age)

	pp := person{name: "tss", age: 58}
	fmt.Printf("\nsecond person name = %s and age = %d", pp.name, pp.age)

	olderp, diff := older(p, pp)
	fmt.Printf("\nthe older is %s,the age diff is %d", olderp.name, diff)
	updateAge(&p)
}

func testStruct2() {
	a := &person{
		name: "dde",
		age:  8,
	}
	updateAge(a)

	b := teacher{age: 10, human: human{name: "jon", sex: 24}}
	fmt.Println(b)
}

func testStruct3() {
	a := &struct {
		name string
		age  int
	}{
		name: "gee",
		age:  55,
	}
	fmt.Println(a)

	b := person{name: "ddd", age: 45}
	b.contact.phone = "2333333"
	b.contact.city = "cdu"
	fmt.Println(b)
}

func updateAge(p *person) {
	p.age = 13
	fmt.Println("update", p)
}
