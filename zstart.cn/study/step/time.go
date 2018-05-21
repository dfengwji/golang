// time
package step

import (
	"fmt"
	"time"
)

var week time.Duration

func StudyTime() {
	testDate()
}

func testTime() {
	t := time.Now()
	fmt.Println(t.Format(time.ANSIC))
}

func testDate() {
	t := time.Now()
	fmt.Println(t)
	fmt.Printf("%02d,%02d.%04d\n", t.Day(), t.Month(), t.Year())
	t = time.Now().UTC()
	fmt.Println(t)
	fmt.Println(time.Now())
	week = 60 * 60 * 24 * 7 * 1e9
	weekFromNow := t.Add(week)
	fmt.Println(weekFromNow)
	s := t.Format("20171107")
	fmt.Println(t, "=>", s)
}
