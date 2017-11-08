package step

import (
	"fmt"
	"sort"
)

func StudyArray() {
	testMap()
}

func testMap() {
	numbers := make(map[string]int)
	numbers["one"] = 1
	numbers["two"] = 20
	numbers["three"] = 35
	fmt.Printf("the 3rd nubmer is = %d\n", numbers["three"])
	for k, v := range numbers {
		fmt.Printf("the key = %s and value is =%d\n", k, v)
	}
	delete(numbers, "one")
	fmt.Println(numbers)
	m := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}
	s := make([]int, len(m))
	i := 0
	for k, _ := range m {
		s[i] = k
		i++
	}
	sort.Ints(s)
	fmt.Println(s)
}

func testSMap() {
	sm := make([]map[int]string, 5)
	for key := range sm {
		sm[key] = make(map[int]string, 1)
		sm[key][1] = "OK"
		fmt.Println(sm[key])
	}
	fmt.Println(sm)
}

func testArray() {
	a := [...]int{19: 1}
	var p *[20]int = &a
	fmt.Println(a)
	fmt.Println(p)
	b := [10]int{}
	b[1] = 2
	fmt.Println(b)
	c := new([10]int)
	c[1] = 2
	fmt.Println(c)
}

func testSort() {
	a := [...]int{5, 3, 6, 9, 10}
	fmt.Println(a)
	num := len(a)
	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++ {
			if a[i] < a[j] {
				tmp := a[i]
				a[i] = a[j]
				a[j] = tmp
			}
		}
	}
	fmt.Println(a)
}

func testSlice2() {
	a := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}
	fmt.Println(a)
	s1 := make([]int, 3, 100)
	fmt.Println(len(s1), cap(s1))
	s2 := a[2:5]
	fmt.Println(len(s2), cap(s2), s2)
	s3 := s2[1:8]
	fmt.Println(s3)
	s3[1] = "w"
	fmt.Println(s2, s3)
	b := []string{"x", "y", "z", "t"}
	copy(a, b)
	fmt.Println(a)
	c := b[:]
	fmt.Println(c)
}

func testAppend2() {
	s1 := make([]int, 3, 6)
	fmt.Printf("%p\n", s1)
	s1 = append(s1, 1, 2, 3)
	fmt.Printf("%v %p\n", s1, s1)
	s1 = append(s1, 1, 2, 3)
	fmt.Printf("%v %p\n", s1, s1)
}

func testSlice() {
	s := make([]int, 0)
	fmt.Println(s)
	a := testAppend(s)
	fmt.Println(s)
	fmt.Println(a)
}

func testAppend(s []int) []int {
	s = append(s, 3)
	return s
}
