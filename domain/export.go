package domain

import "sync"

type Export struct {
	Filename string
}

type Worker interface {
	Do(export *Export, wg *sync.WaitGroup)
}
