package main

import (
	"fmt"
	"sync"
)

type Pool struct {
	tasks chan func()
	wg    sync.WaitGroup
}

func NewPool(size int) *Pool {
	p := &Pool{
		tasks: make(chan func()),
	}
	p.wg.Add(size)
	for i := 0; i < size; i++ {
		go p.worker()
	}
	return p
}

func (p *Pool) worker() {
	defer p.wg.Done()
	for task := range p.tasks {
		fmt.Printf("worker\n")
		task()
	}
}

func (p *Pool) AddTask(task func()) {
	p.tasks <- task
}

func (p *Pool) Close() {
	close(p.tasks)
}

func (p *Pool) Wait() {
	p.wg.Wait()
}

func main() {
	fmt.Printf("main\n")
	pool := NewPool(4)
	for i := 0; i < 10; i++ {
		pool.AddTask(func() {
			fmt.Println("task is running")
		})
	}
	pool.Close()
	pool.Wait()
}

