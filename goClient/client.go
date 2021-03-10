package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/cheggaaa/pb/v3"
)

const requestAmount = 1500
const goRoutineCount = 1
const sleepTimeMilli = 50

func MakeRequest(url string, ch chan<- int64) {
	bar := pb.StartNew(requestAmount)
	for i := 0; i < requestAmount; i++ {
		func() {
			start := time.Now()
			resp, err := http.Get(url)
			milliSecs := time.Since(start).Milliseconds()
			if err != nil {
				ch <- milliSecs
				os.Stderr.WriteString(err.Error() + "\n")
				return
			}
			defer resp.Body.Close()
			ch <- milliSecs
		}()
		bar.Increment()
		time.Sleep(sleepTimeMilli * time.Millisecond)
	}
	bar.Finish()
}

func main() {
	start := time.Now()
	ch := make(chan int64)
	if len(os.Args) < 2 {
		os.Stderr.WriteString("error: usage requires 1 arguments: HOST_ADDRESS:PORT_NUMBER\n")
		return
	}
	var host string = os.Args[1]
	url := fmt.Sprintf("http://%s", host)
	fmt.Println(url)
	for i := 0; i < goRoutineCount; i++ {
		go MakeRequest(url, ch)
	}
	var milliSecs int64 = 0
	for i := 0; i < requestAmount*goRoutineCount; i++ {
		milliSecs += <-ch
	}
	fmt.Printf("%d requests are sent in total\n", requestAmount*goRoutineCount)
	fmt.Printf("%.2fs spent waiting for response\n", float64(milliSecs)/1000.0)
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
