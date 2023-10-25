package creational

import "sync"

type safesingleton struct {
	count int
	sync.RWMutex
}

var safeinstance safesingleton

func GetSafeInstance() *safesingleton {
	return &safeinstance
}

func (s *safesingleton) AddOne() int {
	s.Lock()
	defer s.Unlock()
	s.count++
	return s.count
}

func (s *safesingleton) GetCount() int {
	s.RLock()
	defer s.RUnlock()
	return s.count
}
