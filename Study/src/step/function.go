// function
package step

import (
	"fmt"
)

func StudyFunc() {
	testPanic()
}

func testFun1() {
	x := 3
	y := 4
	z := 5
	max_xy := max(x, y)
	max_yz := max(y, z)
	fmt.Printf("max(%d,%d) = %d\n", x, y, max_xy)
	fmt.Printf("max(%d,%d) = %d\n", z, y, max_yz)
}

func testFun2() {
	var fs = [4]func(){}
	for i := 0; i < 4; i++ {
		defer fmt.Printf("defer i = ", i)
		defer func() {
			fmt.Printf("defer closure i = ", i)
		}()
		fs[i] = func() {
			fmt.Println("closure i = ", i)
		}
	}
	for _, f := range fs {
		f()
	}
}

func testFun() {
	f := closure(10)
	fmt.Println(f(1))
	fmt.Println(f(2))
}

func testDefer() {
	fmt.Printf("a")
	defer fmt.Println("b")
	defer fmt.Println("c")
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func testPanic() {
	testa()
	testb()
	testc()
}

func testa() {
	fmt.Println("func a")
}

func testb() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover in b")
		}
	}()
	panic("panic in b")
}

func testc() {
	fmt.Println("func c")
}

func closure(x int) func(int) int {
	fmt.Printf("%p\n", &x)
	return func(y int) int {
		fmt.Printf("%p\n", &x)
		return x + y
	}
}
