// two
package step

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var cn chan string

func StudyGoroutine() {
	testSelect3()
}

func StudyChannel() {
	testChan()
}

func testGo1() {
	cn = make(chan string)
	go pingPong()
	for i := 0; i < 10; i++ {
		cn <- fmt.Sprintf("From main:Hello,#%d", i)
		fmt.Println(<-cn)
	}
}

func testGo() {
	c := make(chan bool, 1)
	go func() {
		fmt.Println("go go go!!")
		<-c
	}()
	c <- true
}

func testGo2() {
	c := make(chan bool)
	go func() {
		fmt.Println("go go go!!")
		c <- true
		close(c)
	}()
	for v := range c {
		fmt.Println(v)
	}
}

func testGo3() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	c := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go add(c, i)
	}
	for i := 0; i < 10; i++ {
		<-c
	}
}

func testGo4() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go add1(&wg, i)
	}
	wg.Wait()
}

func testSelect() {
	c1, c2 := make(chan int), make(chan string)
	o := make(chan bool, 2)
	go func() {
		for {
			select {
			case v, ok := <-c1:
				if !ok {
					o <- true
					break
				}
				fmt.Println("c1", v)
			case v, ok := <-c2:
				if !ok {
					o <- true
					break
				}
				fmt.Println("c2", v)
			}
		}
	}()
	c1 <- 1
	c2 <- "hi"
	c1 <- 3
	c2 <- "jon"

	close(c1)
	close(c2)

	for i := 0; i < 2; i++ {
		<-o
	}
}

func testSelect2() {
	c := make(chan int)
	go func() {
		for v := range c {
			fmt.Println(v)
		}
	}()
	for {
		select {
		case c <- 0:
		case c <- 1:
		}
	}
}

func testSelect3() {
	c := make(chan bool)
	select {
	case v := <-c:
		fmt.Println(v)
	case <-time.After(3 * time.Second):
		fmt.Println("time out!!!")
	}
}

func add1(wg *sync.WaitGroup, index int) {
	a := 1
	for i := 0; i < 1000000000; i++ {
		a += i
	}
	fmt.Println(index, a)
	wg.Done()
}

func add(c chan bool, index int) {
	a := 1
	for i := 0; i < 1000000000; i++ {
		a += i
	}
	fmt.Println(index, a)
	c <- true
}

func pingPong() {
	i := 0
	for {
		fmt.Println(<-cn)
		cn <- fmt.Sprintf("From pingPong:Hi,#%d", i)
		i++
	}
}

func testChan() {
	c := make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3
	close(c)
	for val := range c {
		fmt.Println(val)
		if len(c) <= 0 {
			break
		}
	}
}

func loop(c chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println("loop-", i)
	}
	c <- 0
}
