package mybase

import (
	"fmt"

	"github.com/google/uuid"
	"go.etcd.io/bbolt"
)

const (
	defaultDBName = "default"
)

type MyBase struct {
	db *bbolt.DB
}

type Collection struct {
	*bbolt.Bucket
}

type M map[string]string

func New() (*MyBase, error) {
	dbName := fmt.Sprintf("%s.mb", defaultDBName)
	db, err := bbolt.Open(dbName, 0666, nil)
	if err != nil {
		return nil, err
	}
	return &MyBase{db: db}, nil
}

func (mb *MyBase) CreateCollection(name string) (*Collection, error) {
	tx, err := mb.db.Begin(true)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	bucket, err := tx.CreateBucketIfNotExists([]byte(name))
	if err != nil {
		return nil, err
	}

	return &Collection{Bucket: bucket}, nil
}

func (mb *MyBase) Insert(collectionName string, data M) (uuid.UUID, error) {
	id := uuid.New()
	tx, err := mb.db.Begin(true)
	if err != nil {
		return id, nil
	}

	bucket, err := tx.CreateBucketIfNotExists([]byte(collectionName))
	if err != nil {
		return id, err
	}

	for k, v := range data {
		if err := bucket.Put([]byte(k), []byte(v)); err != nil {
			return id, err
		}
	}
	if err := bucket.Put([]byte("id"), []byte(id.String())); err != nil {
		return id, err
	}

	return id, tx.Commit()
}

// get https://localhost:8081/users?eq.name=(example@123.com)
// func (mb *MyBase) Select(collectionName string, query string) (M, error) {
// 	tx, err := mb.db.Begin(false)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer tx.Rollback()

// 	bucket := tx.Bucket([]byte(collectionName))
// 	if bucket == nil {
// 		return nil, fmt.Errorf("Bucket (%s) is not foud ", collectionName)
// 	}
// }
