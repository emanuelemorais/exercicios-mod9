package main

import (
	"sync"
	"github.com/emanuelemorais/exercicios-mod9/ponderada-01/pkg/controller"

)

func main() {

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			controller.Controller()
		}()
	}
	wg.Wait()
}