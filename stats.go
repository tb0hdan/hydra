package hydra

import "sync"

type Stats struct {
	lock   sync.RWMutex
	// statistics
	statsMap map[string]int64
	//
}

func (s *Stats) setKV(key string, value int64) {
	defer s.lock.Unlock()
	s.lock.Lock()
	s.statsMap[key] = value
}

func (s *Stats) addKV(kv string, delta int64) {
	previous := s.getKV(kv)
	s.setKV(kv, delta + previous)
}

func (s *Stats) getKV(kv string) int64{
	defer s.lock.RUnlock()
	s.lock.RLock()
	return s.statsMap[kv]
}

// All the setDefaults
func (s *Stats) setDefaults() {
	s.setDefaultStarts()
	//
	s.setDefaultGetItem()
	s.setDefaultGetItemErrors()
	//
	s.setDefaultProcessItem()
	s.setDefaultProcessItemErrors()
	//
	s.setDefaultSubmitItem()
	s.setDefaultSubmitItemErrors()
}

// Starts
func (s *Stats) setDefaultStarts() {
	s.setKV("starts", 0)
}

func (s *Stats) AddStarts(delta int64) {
	s.addKV("starts", delta)
}

func (s *Stats) GetStarts() int64 {
	return s.getKV("starts")
}

// GetItem
func (s *Stats) setDefaultGetItem() {
	s.setKV("getItem", 0)
}

func (s *Stats) AddGetItem(delta int64) {
	s.addKV("getItem", delta)
}

func (s *Stats) GetGetItem() int64 {
	return s.getKV("getItem")
}

func (s *Stats) setDefaultGetItemErrors() {
	s.setKV("getItemErrors", 0)
}

func (s *Stats) AddGetItemErrors(delta int64) {
	s.addKV("getItemErrors", delta)
}

func (s *Stats) GetGetItemErrors() int64 {
	return s.getKV("getItemErrors")
}

// ProcessItem
func (s *Stats) setDefaultProcessItem() {
	s.setKV("processItem", 0)
}
func (s *Stats) AddProcessItem(delta int64) {
	s.addKV("processItem", delta)
}

func (s *Stats) GetProcessItem() int64 {
	return s.getKV("processItem")
}


func (s *Stats) setDefaultProcessItemErrors() {
	s.setKV("processItemErrors", 0)
}

func (s *Stats) AddProcessItemErrors(delta int64) {
	s.addKV("processItemErrors", delta)
}

func (s *Stats) GetProcessItemErrors() int64 {
	return s.getKV("processItemErrors")
}
// SubmitItem

func (s *Stats) setDefaultSubmitItem() {
	s.setKV("submitItem", 0)
}

func (s *Stats) AddSubmitItem(delta int64) {
	s.addKV("submitItem", delta)
}

func (s *Stats) GetSubmitItem() int64 {
	return s.getKV("submitItem")
}

func (s *Stats) setDefaultSubmitItemErrors() {
	s.setKV("submitItemErrors", 0)
}

func (s *Stats) AddSubmitItemErrors(delta int64) {
	s.addKV("submitItemErrors", delta)
}

func (s *Stats) GetSubmitItemErrors() int64 {
	return s.getKV("submitItemErrors")
}

func NewStats() *Stats {
	stats := &Stats{
		lock:     sync.RWMutex{},
		statsMap: make(map[string]int64),
	}
	stats.setDefaults()

	return stats
}
