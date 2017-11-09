// ttt
package main

import (
	"fmt"
	one "step"
	"time"
)

func main() {
	one.StudyControl()
	time.Sleep(5 * time.Second)
}

func testIf() {
	x := 10
	if x > 10 {
		fmt.Print("x is greater than 10")
	} else {
		fmt.Print("x is less than 10")
	}
}

func testFor() {
	sum := 0
	for index := 0; index < 10; index++ {
		sum += index
	}
	fmt.Print("the sum = ", sum)
}

func testForLabel() {

	/*
		for {
			for i := 0; i < 10; i++ {
				if i > 3 {
					break LABEL1
				}
			}
		}
	*/
	for i := 0; i < 10; i++ {
		for {
			goto LABEL1
			fmt.Println("dddddd")
		}
		fmt.Println(i)
	}
LABEL1:
	fmt.Println("Exit for")
}

func testSwitch() {
	i := 10
	switch i {
	case 1:
		fmt.Print("i == 1")
	case 2:
		fmt.Print("i == 2")
		fallthrough
	case 8:
		fmt.Print("i == 10")
	default:
		fmt.Print("can not find i")
	}
}
