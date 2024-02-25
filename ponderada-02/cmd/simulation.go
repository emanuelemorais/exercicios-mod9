package main

import (
	"ponderada-02/pkg/controller"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			controller.Controller(i + 1)
		}()
	}
	wg.Wait()
}