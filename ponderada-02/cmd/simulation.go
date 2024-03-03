package main

import (
	"ponderada-02/pkg/controller"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	for i := 0; i < 1; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			controller.Controller(id)
		}(i + 1)
	}
	wg.Wait()
}
