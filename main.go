package main

import (
	"fmt"
	"math/rand"
	"time"

	log "github.com/sirupsen/logrus"
)

func main() {
	tasks := []*Task{
		NewTask(func() error {
			duration := time.Duration(rand.Intn(2000)) * time.Millisecond
			time.Sleep(duration)
			fmt.Printf("%v - Helloo - %v - \n", time.Now().UTC(), duration)
			return nil
		}),
	}

	p := NewPool(tasks, int(2))
	p.Run()

	var numErrors int
	for _, task := range p.Tasks {
		if task.Err != nil {
			log.Error(task.Err)
			numErrors++
		}

		if numErrors >= 10 {
			log.Error("Too many errors")
			break
		}
	}
}
