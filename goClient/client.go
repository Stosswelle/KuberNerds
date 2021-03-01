package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const requestAmount = 1500
const goRoutineCount = 1
const sleepTimeMilli = 50

func MakeRequest(url string, ch chan<- int64) {
	for i := 0; i < requestAmount; i++ {
		func() {
			start := time.Now()
			resp, err := http.Get(url)
			milliSecs := time.Since(start).Milliseconds()
			if err != nil {
				// ch <- fmt.Sprint("HTTP call failed: ", err)
				ch <- 404
			}
			defer resp.Body.Close()
			// code := resp.StatusCode
			// ch <- fmt.Sprintf("%d milliseconds elapsed with status code: %d", milliSecs, code)
			ch <- milliSecs
		}()
		time.Sleep(sleepTimeMilli * time.Millisecond)
	}
}

func main() {
	start := time.Now()
	ch := make(chan int64)
	if len(os.Args) < 3 {
		fmt.Println("usage requires 2 arguments: HOST_ADDRESS and PORT_NUMBER")
		return
	}
	var host string = os.Args[1]
	var port string = os.Args[2]
	url := fmt.Sprintf("http://%s:%s/productpage", host, port)
	fmt.Println(url)
	for i := 0; i < goRoutineCount; i++ {
		go MakeRequest(url, ch)
	}
	var milliSecs int64 = 0
	for i := 0; i < requestAmount*goRoutineCount; i++ {
		// fmt.Printf("the %dth request took ", i)
		// fmt.Println(<-ch)
		milliSecs += <-ch
	}
	fmt.Printf("%d requests are sent in total\n", requestAmount*goRoutineCount)
	fmt.Printf("%.2fs spent waiting for response\n", float64(milliSecs)/1000.0)
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
