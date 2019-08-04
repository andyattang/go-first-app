package runner

import (
	"log"
	runner "test/runner/run"
	"time"
)

const timeout = 1 * time.Second

func Main() {
	log.Println("Starting work")
	r := runner.New(timeout)
	r.Add(createTask(), createTask(), createTask())

	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrInterrupt:
			log.Println("Terminating due to interrupt.")
			goto exit
		case runner.ErrTimeout:
			log.Println("Terminating due to timeout.")
			goto exit
		}
	}

exit:
	time.Sleep(time.Duration(10) * time.Second)

	log.Println("End work")
}

func createTask() func(int) {
	return func(id int) {
		log.Printf("Processor - Task #%d started", id)
		time.Sleep(time.Duration(id+2) * time.Second)
		log.Printf("Processor - Task #%d ended", id)
	}
}
