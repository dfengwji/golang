// method
package step

import (
	"fmt"
)

type Role struct {
	Name string
}

func StudyMethod() {
	a := Role{}
	a.setName("jonh")
	fmt.Println(a.Name)
}

func (this Role) setName(n string) {
	this.Name = n
	fmt.Println("set role name", n)
}
