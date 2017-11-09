package step

import "fmt"

func StudyPointer() {
	test2()
}

func testPointer() {
	var i1 = 5
	fmt.Printf("an integer:%d,its location in memory:%p\n", i1, &i1)
	var intP *int
	intP = &i1
	fmt.Printf("the value at memory location %p is %d\n", intP, *intP)
}

func test2()  {
	s := "good bye"
	var p = &s
	fmt.Printf("here is the string s : %s\n",s)
	fmt.Printf("here is the string *p : %s\n",*p)
	*p = "ciao"
	fmt.Printf("here is the pointer p : %p\n",p)
	fmt.Printf("here is the string *p : %s\n",*p)
	fmt.Printf("here is the string s : %s\n",s)
}

func testCrash()  {
	var p *int = nil
	*p = 0
}
