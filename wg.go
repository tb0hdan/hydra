package hydra

import "sync"

type MyWG struct {
	wg *sync.WaitGroup
	wgLen int
}

func (mg *MyWG) Len() int{
	return mg.wgLen
}

func (mg *MyWG) Add(delta int) {
	mg.wgLen += delta
	mg.wg.Add(delta)
}

func (mg *MyWG) Done() {
	if mg.wgLen > 0 {
		mg.wgLen--
	}

	mg.wg.Done()
}

func (mg *MyWG) Wait() {
	mg.wg.Wait()
}
