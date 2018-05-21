// three
package step

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

func getDemo(addr string) {
	res, err := http.Get(addr)
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}

// This example fetches several URLs concurrently,
// using a WaitGroup to block until all the fetches are complete.
func StudyWaitGroup() {
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.baidu.com/",
	}
	for _, url := range urls {
		// Increment the WaitGroup counter.
		wg.Add(1)
		// Launch a goroutine to fetch the URL.
		go func(url string) {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()
			// Fetch the URL.
			getDemo(url)
		}(url)
	}
	// Wait for all HTTP fetches to complete.
	wg.Wait()
	fmt.Println("--------------------group wait over ---------------------------")
}

func StudyOnce() {
	testOnce1()
}

func testOnce2() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody)
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
	fmt.Println("once over -------------------------------")
	// Output:
	// Only once
}

func testOnce1() {
	var once sync.Once
	onceBody := func() {
		time.Sleep(3e9)
		fmt.Println("Only once")
	}
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		j := i
		go func(int) {
			once.Do(onceBody)
			fmt.Println(j)
			done <- true
		}(j)
	}
	//给一部分时间保证能够输出完整【方法一】
	//for i := 0; i < 10; i++ {
	//    <-done
	//}

	//给一部分时间保证能够输出完整【方法二】
	<-done
	time.Sleep(3e9)
}
