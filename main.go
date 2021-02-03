package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const JOBS_COUNT = 5

func main() {
	var writers, arrSize, iterCount int
	flag.IntVar(&writers, "writers", 0, "Количество пишущих горутин")
	flag.IntVar(&arrSize, "arr-size", 0, "Размер массива")
	flag.IntVar(&iterCount, "iter-count", 0, "Количество итераций")
	flag.Parse()
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < writers; i++ {
		go writer(i, arrSize, iterCount)
	}
	fmt.Scanln()
}

func writer(id, arrSize, iterCount int) {
	jobs := make(chan []int, 2)
	results := make(chan string, 2)
	for i := 0; i < JOBS_COUNT; i++ {
		t := time.Now()
		go job(id, t, jobs, results)
	}
	for i := 0; i < iterCount; i++ {
		arr := make([]int, 0, arrSize)
		for j := 0; j < arrSize; j++ {
			arr = append(arr, rand.Intn(1000))
		}
		jobs <- arr
	}
	close(jobs)
	for i := 0; i < iterCount; i++ {
		fmt.Print(<-results)
	}
}

func job(id int, t time.Time, jobs <-chan []int, results chan<- string) {
	for w := range jobs {
		sort.Ints(w)
		res := fmt.Sprintf("%d %s %d %d %d\n", id, t.Format(time.StampNano), w[0], w[len(w)/2], w[len(w)-1])
		results <- res
	}
}
