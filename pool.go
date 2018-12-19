package main

import (
	"sync"

	log "github.com/sirupsen/logrus"
)

type Pool struct {
	Tasks       []*Task
	concurrency int
	tasksChan   chan *Task
	wg          sync.WaitGroup
}

func NewPool(tasks []*Task, concurrency int) *Pool {
	return &Pool{
		Tasks:       tasks,
		concurrency: concurrency,
		tasksChan:   make(chan *Task),
	}
}

func (p *Pool) Run() {
	log.Debugf("Running %v task(s) at concurrency %v.", len(p.Tasks), p.concurrency)
	for i := 0; i < p.concurrency; i++ {
		go p.work()
	}

	p.wg.Add(len(p.Tasks))

	for _, task := range p.Tasks {
		p.tasksChan <- task
	}

	close(p.tasksChan)
	p.wg.Wait()
}

func (p *Pool) work() {
	for task := range p.tasksChan {
		task.Run(&p.wg)
	}
}
