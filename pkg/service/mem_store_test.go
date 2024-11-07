package service

import "testing"

func Test_MemStoreContainsAfterSave(t *testing.T) {
	store := NewInMemoryStore()
	err := store.Save("key", "value")
	if err != nil {
		t.Error("failed to save")
	}
	ok, err := store.Contains("key")
	if err != nil {
		t.Error("failed to check for key")
	}
	if !ok {
		t.Error("key not found after save")
	}
}

func Test_MemStoreDoesNotContainAfterDelete(t *testing.T) {
	store := NewInMemoryStore()
	_ = store.Save("key", "value")
	err := store.Delete("key")
	if err != nil {
		t.Error("failed to delete")
	}
	ok, err := store.Contains("key")
	if err != nil {
		t.Error("failed to check for key")
	}
	if ok {
		t.Error("key found after delete")
	}
}

func Test_MemStoreCanLoadAfterSave(t *testing.T) {
	store := NewInMemoryStore()
	_ = store.Save("key", "value")
	value, err := store.Load("key")
	if err != nil {
		t.Error("failed to load")
	}
	if value != "value" {
		t.Error("loaded value does not match saved value")
	}
}
