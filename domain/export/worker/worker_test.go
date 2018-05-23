package worker

import (
	"fmt"
	"sync"
	"testing"

	"time"

	"github.com/ricardolonga/workshop-go/domain"
)

func TestWorker_Do(t *testing.T) {
	worker := New()

	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)

		export := &domain.Export{Filename: fmt.Sprintf("arquivo_%d.csv", i)}
		go worker.Do(export, wg)

		time.Sleep(100 * time.Millisecond)
	}
	wg.Wait()
}
