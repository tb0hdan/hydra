package hydra

import (
	"context"
	"sync"
)

type Worker struct {
	ctx context.Context
	stopChan chan struct{}
	exitFlag bool
	workerCount int
	workerNode WorkerNode
	mg *MyWG
	stats *Stats
	logger Logger
}

func (w *Worker) checkIfWorkerNeeded(workerID int) {
	if w.exitFlag {
		w.logger.Debugln("Exit requested, not adding new workers...")

		return
	}

	if w.mg.Len() < w.workerCount {
		// start new one
		w.logger.Debugln("Adding new worker...", workerID)
		go w.workerFlow(workerID)
	}
}

func (w *Worker) workerFlow(workerID int) {
	w.logger.Debugf("Starting worker: %d\n", workerID)
	// Increase running worker count
	w.mg.Add(1)
	w.stats.AddStarts(1)

	// on function exit decrease worker amount count
	// and start new one
	defer func(){
		w.mg.Done()
		w.checkIfWorkerNeeded(workerID)
	}()

	// get item
	w.stats.AddGetItem(1)
	item, err := w.workerNode.GetItem(w.ctx)
	if err != nil {
		w.logger.Debugf("Could not get item: %v", err)
		w.stats.AddGetItemErrors(1)
		return
	}

	// process item
	w.stats.AddProcessItem(1)
	result, err := w.workerNode.Process(w.ctx, item)
	if err != nil {
		w.logger.Debugf("Could not process item: %v", err)
		w.stats.AddProcessItemErrors(1)
		return
	}
	// submit result
	w.stats.AddSubmitItem(1)
	err = w.workerNode.SubmitResult(w.ctx, result)
	if err != nil {
		w.logger.Debugf("Could not submit result: %v", err)
		w.stats.AddSubmitItemErrors(1)
		return
	}
}

func (w *Worker) Run() {
	// start workers, future starts will be event-based
	for i := 0; i <= w.workerCount; i++ {
		go w.workerFlow(i)
	}
	//
	select {
	case <-w.stopChan:
		w.logger.Println("Stop requested, exiting...")
		break
	case <-w.ctx.Done():
		w.logger.Println("Ctx.Done(), exiting...")
		break
	}
	w.exitFlag = true
}

func (w *Worker) Stop() {
	w.stopChan <- struct{}{}
}

func (w *Worker) Stats() *Stats{
	return w.stats
}

func New(ctx context.Context, workerCount int, workerNode WorkerNode, logger Logger) *Worker {
	return &Worker{
		ctx:      ctx,
		stopChan: make(chan struct{}),
		workerCount: workerCount,
		workerNode: workerNode,
		mg: &MyWG{wg: &sync.WaitGroup{}},
		stats: NewStats(),
		logger: logger,
	}
}




