package util

import "sync"

type SyncMap[T any] struct {
	m sync.Map
}

func NewSyncMap[T any]() *SyncMap[T] {
	return &SyncMap[T]{}
}

func (s *SyncMap[T]) Load(key string) (*T, bool) {
	value, ok := s.m.Load(key)
	if !ok {
		return nil, false
	}
	v, ok := value.(T)
	if !ok {
		return nil, false
	}
	return &v, true
}

func (s *SyncMap[T]) Store(key string, value T) {
	s.m.Store(key, value)
}

func (s *SyncMap[T]) Contains(key string) bool {
	_, ok := s.Load(key)
	return ok
}

func (s *SyncMap[T]) Delete(key string) {
	s.m.Delete(key)
}
