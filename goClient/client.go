package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const requestAmount = 100
const goRoutineCount = 5

func MakeRequest(url string, ch chan<- string) {
	for i := 0; i < requestAmount; i++ {

		start := time.Now()
		resp, err := http.Get(url)
		secs := time.Since(start).Milliseconds()
		code := 404
		if err == nil {
			code = resp.StatusCode
		}
		ch <- fmt.Sprintf("%d milliseconds elapsed with status code: %d", secs, code)
		time.Sleep(440 * time.Millisecond)
	}
}

func main() {
	start := time.Now()
	ch := make(chan string)
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
	for i := 0; i < requestAmount*goRoutineCount; i++ {
		fmt.Printf("the %dth request took ", i)
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
