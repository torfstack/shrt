package service

import (
	"errors"
	"shrt/pkg/util"
)

var (
	ErrKeyNotFound = errors.New("key not found in in-mem store")
)

type memStore struct {
	store *util.SyncMap[string]
}

var _ Store = &memStore{}

func NewInMemoryStore() Store {
	m := util.NewSyncMap[string]()
	return &memStore{m}
}

func (m *memStore) Save(key string, value string) error {
	m.store.Store(key, value)
	return nil
}

func (m *memStore) Load(key string) (string, error) {
	value, ok := m.store.Load(key)
	if !ok {
		return "", ErrKeyNotFound
	}
	return *value, nil
}

func (m *memStore) Contains(key string) (bool, error) {
	return m.store.Contains(key), nil
}

func (m *memStore) Delete(key string) error {
	m.store.Delete(key)
	return nil
}
