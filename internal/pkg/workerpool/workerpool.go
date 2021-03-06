package workerpool

// go:generate mockgen -source=workerpool-application/internal/pkg/workerpool.go -destination=./../../mocks/workerpool_mock.go -package=mocks WorkerPoolInterface
import (
	"log"
	"workerpool-application/internal/clients"
	"workerpool-application/internal/services"
)

type WorkerResp struct {
	Endpoint string
	Result   string
}

type WorkerPoolInterface interface {
	DoWork(endpoints <-chan string, responses chan<- WorkerResp)
}

type WorkerPool struct {
	client clients.ClientInterface
	hasher services.HasherInterface
}

func NewWorkerPool(client clients.ClientInterface, hasher services.HasherInterface) WorkerPoolInterface {
	return &WorkerPool{client: client, hasher: hasher}
}

func (w *WorkerPool) DoWork(endpoints <-chan string, responses chan<- WorkerResp) {
	for endpoint := range endpoints {
		body, err := w.client.DoReq(endpoint)
		if err != nil {
			responses <- WorkerResp{Endpoint: endpoint, Result: err.Error()}
			continue
		}
		log.Println("request completed", body)
		result := w.hasher.HashInput(body)
		responses <- WorkerResp{Endpoint: endpoint, Result: result}
	}
}
