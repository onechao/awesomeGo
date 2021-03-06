package main

import (
	"fmt"
	"sync"
)

func doWork(id int, w worker) {
	for n := range w.in { //判断chan是否已close
		fmt.Printf("worker %d received %c \n",
			id, n)
		//go func() { done <- true }() //等一次不需要go
		//done <- true //大小写分开
		w.done()
	}
}

type worker struct {
	in chan int
	//done chan bool
	//wg *sync.WaitGroup
	done func()
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		//done: make(chan bool),
		//wg: wg,
		done: func() {
			wg.Done()
		},
	}
	go doWork(id, w)
	return w
}

func chanDemo() {
	var wg sync.WaitGroup

	//chan是blocking的，go func有点类似异步线程

	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}
	//每个block收完后，必须有接收，否则不能发送第二次↓

	wg.Add(20)

	//第一次未处理时，不能再发送第二次，必须先收。
	for i, worker := range workers {
		worker.in <- 'a' + i
	}

	//分开等
	//for _, worker := range workers {
	//	<-worker.done
	//}

	for i, worker := range workers {
		worker.in <- 'A' + i
	}
	//分开等
	//for _, worker := range workers {
	//	<-worker.done
	//}
	wg.Wait()
}

func main() {
	chanDemo()
}
