package kvstore

import (
	"github.com/dgraph-io/badger/v3"
)

// KVStore represents a key-value store backed by an LSM tree (using BadgerDB).
type KVStore struct {
	db *badger.DB
}

// NewKVStore initializes a new key-value store instance.
func NewKVStore(dir string) (*KVStore, error) {
	opts := badger.DefaultOptions(dir).WithLogger(nil) // disable logging for brevity
	db, err := badger.Open(opts)
	if err != nil {
		return nil, err
	}
	return &KVStore{db: db}, nil
}

// Get retrieves the value for a given key.
func (store *KVStore) Get(key string) (string, error) {
	var result string
	err := store.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		val, err := item.ValueCopy(nil)
		if err != nil {
			return err
		}
		result = string(val)
		return nil
	})
	return result, err
}

// Put stores a value for a given key.
func (store *KVStore) Put(key, value string) error {
	return store.db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(key), []byte(value))
	})
}

// Close shuts down the key-value store.
func (store *KVStore) Close() error {
	return store.db.Close()
}
