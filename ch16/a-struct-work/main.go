package main

import (
	"fmt"
	"sync"
)

type Worker struct {
	in, out chan int
	sbw     int
	mutex   *sync.Mutex
}

func (w *Worker) readThem() {
	w.sbw++
	go func() {
		partial := 0
		for i := range w.in {
			partial += i
		}
		w.out <- partial
		w.mutex.Lock()
		w.sbw--
		if w.sbw == 0 {
			close(w.out)
		}
		w.mutex.Unlock()
	}()
}

func (w *Worker) gatherResult() int {
	total := 0
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for i := range w.out {
			total += i
		}
		wg.Done()
	}()
	wg.Wait()
	return total
}

func main() {
	mutex := &sync.Mutex{}
	in := make(chan int, 100)
	wrNum := 10
	out := make(chan int)
	wrk := Worker{in: in, out: out, mutex: mutex}
	for i := 1; i <= wrNum; i++ {
		wrk.readThem()
	}
	for i := 1; i <= 100; i++ {
		in <- i
	}
	close(in)
	res := wrk.gatherResult()
	fmt.Println("Sum =", res)
}
