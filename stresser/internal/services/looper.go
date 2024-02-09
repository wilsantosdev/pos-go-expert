package services

import (
	"fmt"
	"stresser/internal/domain/entity"
	"stresser/internal/domain/interfaces"
	"sync"
)

type looper struct {
	requester interfaces.Requester
}

func NewLooper(requester interfaces.Requester) *looper {
	return &looper{
		requester: requester,
	}
}

func (l *looper) Loop(concurrent int, requests int, url string) (entity.Report, error) {
	wg := sync.WaitGroup{}
	exeChannel := make(chan int, concurrent)
	report := entity.NewReport()
	defer close(exeChannel)
	for i := 0; i < requests; i++ {
		exeChannel <- i
		wg.Add(1)
		go func() {
			defer wg.Done()
			requestResponse, err := l.requester.Request(url)
			if err != nil {
				fmt.Println(err)
			}
			report.AddResponse(*requestResponse)
			<-exeChannel
		}()

	}
	wg.Wait()

	return *report, nil
}
