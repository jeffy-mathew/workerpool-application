package app

import (
	"crypto/md5"
	"flag"
	"log"
	"net/http"
	"strconv"
	"time"
	"workerpool-application/internal/clients"
	"workerpool-application/internal/pkg/workerpool"
	"workerpool-application/internal/services"
)

type App interface {
	Run()
}

type WorkerPoolApp struct {
	workerPool workerpool.WorkerPoolInterface
}

func NewApp() App {
	client := clients.NewClient(&http.Client{Timeout: 15 * time.Second})
	hasher := services.NewHasher(md5.New())
	workerpool := workerpool.NewWorkerPool(client, hasher)
	return &WorkerPoolApp{
		workerPool: workerpool,
	}
}

func (app *WorkerPoolApp) Run() {
	startTime := time.Now()
	parallelFlag := flag.String("parallel", "5", "number of concurrent requests to be made - default is 5")
	flag.Parse()
	endpoints := flag.Args()
	if len(endpoints) == 0 {
		log.Fatal("FATAL: endpoints cannot be empty")
	}
	// convert parallel flag to int
	parallel, err := strconv.Atoi(*parallelFlag)
	if err != nil {
		log.Fatal("FATAL: invalid value for parallel flag", parallelFlag)
	}
	// start workers
	jobs := make(chan string, len(endpoints))
	results := make(chan workerpool.WorkerResp, len(endpoints))

	for w := 0; w < parallel; w++ {
		go app.workerPool.DoWork(jobs, results)
	}

	for j := 0; j < len(endpoints); j++ {
		jobs <- endpoints[j]
	}
	close(jobs)

	for a := 1; a <= len(endpoints); a++ {
		result := <-results
		log.Printf("%s - %s", result.Endpoint, result.Result[0:32])
	}
	close(results)
	log.Println("INFO: time elaspesed ", time.Since(startTime))
}
