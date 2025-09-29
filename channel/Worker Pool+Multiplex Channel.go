package main

import (
	"fmt"
	"math/rand"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- string) {
	for job := range jobs {
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		results <- fmt.Sprintf("Worker %d finished job %d", id, job)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	jobs := make(chan int, 5)
	results := make(chan string, 5)

	// สร้าง worker 2 ตัว
	for i := 1; i <= 2; i++ {
		go worker(i, jobs, results)
	}

	// ส่งงานเข้าไป
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// ใช้ select รับผลลัพธ์
	for i := 0; i < 5; i++ {
		select {
		case res := <-results:
			fmt.Println(res)
		}
	}
}
