package step

import (
	"fmt"
	"strings"
	"strconv"
)

func StudyString() {
	testConv()
}

func testFields() {
	str := "The quick brown fox jumps over the lazy dog"
	sl := strings.Fields(str)
	fmt.Printf("Splitted in slice: %v\n", sl)
	for _, val := range sl {
		fmt.Printf("%s - ", val)
	}
	fmt.Println()
	str2 := "GO1|The ABC of Go|25"
	sl2 := strings.Split(str2, "|")
	fmt.Printf("Splitted in slice: %v\n", sl2)
	for _, val := range sl2 {
		fmt.Printf("%s - ", val)
	}
	fmt.Println()
	str3 := strings.Join(sl2, ";")
	fmt.Printf("sl2 joined by ;: %s\n", str3)
}

func testReader(){
	str := "this is test string"
	reader := strings.NewReader(str)
	fmt.Printf("string len = %d", reader.Len())
}

func testConv()  {
	var orig string = "666"
	var an int
	var newStr string
	fmt.Printf("the size of ints is:%d\n",strconv.IntSize)
	an,_ = strconv.Atoi(orig)
	fmt.Printf("the integer is : %d\n",an)
	an += 5
	newStr = strconv.Itoa(an)
	fmt.Printf("this new string is :%s\n",newStr)
}
