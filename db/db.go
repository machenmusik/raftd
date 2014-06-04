package db

import (
	"encoding/json"
	"sync"
)

// The key-value database.
type DB struct {
	data  map[string]string
	mutex sync.RWMutex
}

// Creates a new database.
func New() *DB {
	return &DB{
		data: make(map[string]string),
	}
}

// Retrieves the value for a given key.
func (db *DB) Get(key string) (string, bool) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	val, ok := db.data[key]
	return val, ok
}

// Dump the database as json
func (db *DB) DumpJSON() ([]byte) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	val, _ := json.Marshal(db.data)
	return val
}

// Sets the value for a given key.
func (db *DB) Put(key string, value string) {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	db.data[key] = value
}
