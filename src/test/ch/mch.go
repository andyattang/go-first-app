package ch

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	chanNumber   = 4
	workerNumber = 4
	taskNumbeer  = 100
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().Unix())
}

func Main() {
	// build workers
	wg.Add(workerNumber)
	tasks := make(chan string, chanNumber)
	for i := 0; i < workerNumber; i++ {
		go work(tasks, i+1)
	}
	// assign tasks
	for i := 0; i < taskNumbeer; i++ {
		tasks <- fmt.Sprintf("Task %d", i+1)
	}
	close(tasks)
	// wait for completing
	wg.Wait()
}

func work(tasks chan string, num int) {
	defer wg.Done()
	for {
		task, ok := <-tasks
		if !ok {
			fmt.Printf("Worker %d: shutdown\n", num)
			return
		}
		fmt.Printf("Worker %d: start - %s\n", num, task)
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		fmt.Printf("Worker %d: completed - %s\n", num, task)
	}
}
