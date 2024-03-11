package main

import (
	"ponderada-04/pkg/controller"
	database "ponderada-04/internal/database"
	"sync"
	"log"
)

func main() {


	sensors, err := database.GetAllSensors()
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	for _, sensor := range sensors {
		wg.Add(1)
		go func(id string) {
			defer wg.Done()
			controller.Controller(id)
		}(sensor.ID)
	}	
	wg.Wait()
}
