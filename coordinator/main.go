package main

import (
	ccollector "coordinator/internal/controllers/collector"
	ccoordinator "coordinator/internal/controllers/coordinator"
	"fmt"
	"os"
	"sync"
)

func main() {

	fmt.Println("Servidor TCP escuchando en el puerto 3000...")

	file, err := os.Open("assets/wiki_concatenated.txt")
 	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
 	}
 	defer file.Close()

	var wg sync.WaitGroup

	wg.Go(func() {ccoordinator.Coordinate(file)})
	wg.Go(func() {ccollector.Collect()})

	wg.Wait()
}