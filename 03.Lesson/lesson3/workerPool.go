package lesson3

import (
	"fmt"
	"math/rand"
	"time"
)

type workers struct {
	id       int
	duration time.Duration
}

func Worker(id int, pool chan workers) {
	rand.Seed(time.Now().UTC().UnixNano())
	second := time.Duration(1 + rand.Intn(4-1))
	time.Sleep(time.Second * second)
	pool <- workers{id, second}
}

func WorkerPool(countWorkers int) {
	pool := make(chan workers)
	for id := 0; id < countWorkers; id++ {
		go Worker(id, pool)
	}

	for i := 0; i < countWorkers; i++ {
		var worker workers
		worker = <-pool
		fmt.Println(
			"Worker", worker.id,
			"- time", time.Second*worker.duration,
		)
	}
}
