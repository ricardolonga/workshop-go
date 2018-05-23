package worker

import (
	"fmt"
	"sync"
	"testing"

	"github.com/ricardolonga/workshop-go/domain"
)

func TestWorker_Do(t *testing.T) {
	worker := New()

	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		export := &domain.Export{Filename: fmt.Sprintf("arquivo_%d.csv", i)}
		wg.Add(1)
		go worker.Do(export, wg)
	}
	wg.Wait()
}
