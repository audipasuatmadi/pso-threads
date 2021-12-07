package main

import (
	"fmt"
	"log"

	"github.com/audipasuatmadi/pso-threads/go-worker-pools/internal/worker"
)

func initializeWorkers(name string, c chan int) {
	worker := worker.Worker{Name: name, C: c}
	worker.Work()
}

func main() {
	log.Printf("hello ser!")

	c := make(chan int, 10)

	// rekrut 10 pegawai
	for i := 0; i < 10; i++ {
		go initializeWorkers(fmt.Sprintf("Pegawai %d", i), c)
	}

	// beri 100 pekerjaan
	for i := 0; i < 100; i++ {
		c <- i
	}
}
