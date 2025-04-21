package queue

import (
  "time"
  bolt "go.etcd.io/bbolt"
  "github.com/google/uuid"
)

const bucketMessages = "messages"

type Store struct {
  db *bolt.DB
}

func NewStore(path string) (*Store, error) {
  db, err := bolt.Open(path, 0600, &bolt.Options{Timeout: 1 * time.Second})
  if err != nil {
    return nil, err
  }
  // crear bucket si no existe
  db.Update(func(tx *bolt.Tx) error {
    _, e := tx.CreateBucketIfNotExists([]byte(bucketMessages))
    return e
  })
  return &Store{db: db}, nil
}

// Encola un mensaje y devuelve su id
func (s *Store) Enqueue(queue string, data []byte) (string, error) {
  id := uuid.New().String()
  key := []byte(queue + "|" + id)
  return id, s.db.Update(func(tx *bolt.Tx) error {
    b := tx.Bucket([]byte(bucketMessages))
    return b.Put(key, data)
  })
}

// Devuelve el siguiente mensaje sin borrarlo
func (s *Store) Peek(queue string) (id string, data []byte, err error) {
  prefix := []byte(queue + "|")
  err = s.db.View(func(tx *bolt.Tx) error {
    c := tx.Bucket([]byte(bucketMessages)).Cursor()
    for k, v := c.Seek(prefix); k != nil && len(k) >= len(prefix) && string(k[:len(prefix)]) == string(prefix); k, v = c.Next() {
      id = string(k[len(prefix):])
      data = append([]byte(nil), v...)
      return nil
    }
    return bolt.ErrBucketNotFound
  })
  return
}

// Quita un mensaje ya acked
func (s *Store) Ack(queue, id string) error {
  key := []byte(queue + "|" + id)
  return s.db.Update(func(tx *bolt.Tx) error {
    return tx.Bucket([]byte(bucketMessages)).Delete(key)
  })
}