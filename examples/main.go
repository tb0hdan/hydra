package main

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/tb0hdan/hydra"
)

type MyWorkerNode struct {
	logger hydra.Logger
}

func (mn *MyWorkerNode) Process(ctx context.Context, item interface{}) (interface{}, error) {
	defer ctx.Done()

	log.Println("Processing item", item)
	time.Sleep(1 * time.Second)
	return item, nil
}

func (mn *MyWorkerNode) GetItem(ctx context.Context) (interface{}, error) {
	defer ctx.Done()

	item := "xXx"
	log.Println("Getting item", item)
	return item, nil
}

func (mn *MyWorkerNode) SubmitResult(ctx context.Context, result interface{}) error {
	defer ctx.Done()

	log.Println("Submitting result", result)
	return nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 60 * time.Second)
	defer cancel()

	logger := log.New()
	mn := &MyWorkerNode{logger: logger}
	worker := hydra.New(ctx, 32, mn, logger)
	go func(w *hydra.Worker){
		time.Sleep(30 * time.Second)
		stats := w.Stats()
		logger.Println(stats)
		w.Stop()
	}(worker)
	worker.Run()
}
