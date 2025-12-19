package services

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go.etcd.io/bbolt"
	"iris/types"
	"time"
)

var DB *bbolt.DB

func GetDB() *bbolt.DB {
	if DB == nil {
		db, err := bbolt.Open("./storage.db", 0600, nil)
		if err != nil {
			panic(err)
		}
		DB = db
	}
	return DB
}

type Persistable interface {
	Entity() *types.Entity
}

type Storage[T Persistable] struct {
	db        *bbolt.DB
	namespace string
}

func NewStorage[T Persistable](db *bbolt.DB, namespace string) *Storage[T] {
	err := db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(namespace))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return &Storage[T]{db: db, namespace: namespace}
}

func (s *Storage[T]) Create(object T) (int, error) {
	var id int
	err := s.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(s.namespace))
		if b == nil {
			return fmt.Errorf("bucket %s not found", s.namespace)
		}

		// get next sequence
		seq, err := b.NextSequence()
		if err != nil {
			return err
		}
		id = int(seq)

		object.Entity().ID = id
		object.Entity().CreatedAt = time.Now()
		object.Entity().UpdatedAt = time.Now()
		object.Entity().DeletedAt = time.Time{}

		// marshal object to JSON
		data, err := json.Marshal(object)
		if err != nil {
			return err
		}

		// store using sequence as key
		key := itob(id) // convert int -> []byte
		return b.Put(key, data)
	})
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *Storage[T]) Update(object T) error {
	return s.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(s.namespace))
		if b == nil {
			return fmt.Errorf("bucket %s not found", s.namespace)
		}

		// Get the key
		id := itob(object.Entity().ID)

		// Make sure it exists
		if b.Get(id) == nil {
			return fmt.Errorf("object with id %d not found", object.Entity().ID)
		}

		// Update timestamp
		object.Entity().UpdatedAt = time.Now()

		// Marshal object
		data, err := json.Marshal(object)
		if err != nil {
			return err
		}

		// Store it
		return b.Put(id, data)
	})
}

func (s *Storage[T]) Get(id int) (T, error) {
	var obj T
	err := s.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(s.namespace))
		if b == nil {
			return fmt.Errorf("bucket %s not found", s.namespace)
		}

		data := b.Get(itob(id))
		if data == nil {
			return fmt.Errorf("object with id %d not found", id)
		}

		if err := json.Unmarshal(data, &obj); err != nil {
			return err
		}

		// filter deleted
		if !obj.Entity().DeletedAt.IsZero() {
			return fmt.Errorf("object with id %d has been deleted", id)
		}

		return nil
	})
	return obj, err
}

func (s *Storage[T]) GetAll() ([]T, error) {
	var objs []T
	err := s.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(s.namespace))
		if b == nil {
			return fmt.Errorf("bucket %s not found", s.namespace)
		}

		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			var obj T
			if err := json.Unmarshal(v, &obj); err != nil {
				return err
			}

			// skip deleted
			if !obj.Entity().DeletedAt.IsZero() {
				continue
			}

			objs = append(objs, obj)
		}

		return nil
	})
	return objs, err
}

func (s *Storage[T]) Delete(object T) error {
	return s.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(s.namespace))
		if b == nil {
			return fmt.Errorf("bucket %s not found", s.namespace)
		}

		id := itob(object.Entity().ID)

		data := b.Get(id)
		if data == nil {
			return fmt.Errorf("object with id %d not found", object.Entity().ID)
		}

		// mark as deleted
		now := time.Now()
		object.Entity().DeletedAt = now
		object.Entity().UpdatedAt = now

		newData, err := json.Marshal(object)
		if err != nil {
			return err
		}

		return b.Put(id, newData)
	})
}

// helper: int -> []byte (lil endian)
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(b []byte) int {
	return int(binary.LittleEndian.Uint64(b))
}
