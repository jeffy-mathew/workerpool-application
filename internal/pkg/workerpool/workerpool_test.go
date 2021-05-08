package workerpool

import (
	"errors"
	"sync"
	"testing"
	"workerpool-application/internal/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestDoWork(t *testing.T) {
	t.Run("should respond with hash when request is success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		clientMock := mocks.NewMockClientInterface(ctrl)
		hasherMock := mocks.NewMockHasherInterface(ctrl)
		mockWg := &sync.WaitGroup{}
		mockWg.Add(1)
		clientMock.EXPECT().DoReq("http://google.com").Return([]byte("response"), nil)
		hasherMock.EXPECT().HashInput([]byte("response")).Return("new hash").Do(func(arg0 interface{}) {
			defer mockWg.Done()
		})
		workerPool := NewWorkerPool(clientMock, hasherMock)
		endpointsChan := make(chan string)
		resultsChan := make(chan WorkerResp)
		go workerPool.DoWork(endpointsChan, resultsChan)
		wg := &sync.WaitGroup{}
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			endpointsChan <- "http://google.com"
		}(wg)
		expectedResult := WorkerResp{
			Endpoint: "http://google.com",
			Result:   "new hash",
		}
		assert.Equal(t, expectedResult, <-resultsChan)
		wg.Wait()
		mockWg.Wait()
	})
	t.Run("should respond with error when request is failed", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		clientMock := mocks.NewMockClientInterface(ctrl)
		hasherMock := mocks.NewMockHasherInterface(ctrl)
		mockWg := &sync.WaitGroup{}
		mockWg.Add(1)
		clientMock.EXPECT().DoReq("http://google.com").Return(nil, errors.New("request failed")).Do(func(arg0 interface{}) {
			defer mockWg.Done()
		})
		workerPool := NewWorkerPool(clientMock, hasherMock)
		endpointsChan := make(chan string)
		resultsChan := make(chan WorkerResp)
		go workerPool.DoWork(endpointsChan, resultsChan)
		wg := &sync.WaitGroup{}
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			endpointsChan <- "http://google.com"
		}(wg)
		expectedResult := WorkerResp{
			Endpoint: "http://google.com",
			Result:   "request failed",
		}
		assert.Equal(t, expectedResult, <-resultsChan)
		wg.Wait()
		mockWg.Wait()
	})
}
