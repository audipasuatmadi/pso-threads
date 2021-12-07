package worker

import (
	"log"
	"math/rand"
	"time"
)

type Worker struct {
	C    chan int
	Name string
}

func (w Worker) Work() {
	for range w.C {
		i := <-w.C
		rn := rand.Intn(5-1) * 100
		log.Printf("%s mulai bekerja pada pekerjaan dengan id %d\n", w.Name, i)
		time.Sleep(time.Duration(rn) * time.Millisecond)
		log.Printf("%s selesai bekerja pada pekerjaan dengan id %d\n", w.Name, i)
	}
}
