package db

import (
	"fmt"
	"log"
	"time"

	"github.com/boltdb/bolt"
)

var DB *bolt.DB

func ConnectToBoltDB() {
	var err error
	DB, err = bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Println("Error in connecting to BOLTDB ", err)
		return
	}

}

// InsertToBOLTDB function will insert the data to BOLTDB
func InsertToBoltDB(task string) error {
	err := DB.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("data"))
		if err != nil {
			return err
		}
		err = bucket.Put([]byte(task), []byte(task))
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

// GetFromBOLTDB will retrieve the data from BOLTDB
func GetFromBoltDB(key string) (string, error) {
	var value string
	err := DB.View(func(tx *bolt.Tx) error {

		bucket := tx.Bucket([]byte("data"))
		if bucket == nil {
			return fmt.Errorf("bucket not found")
		}
		// Retrieve data from the bucket
		value = string(bucket.Get([]byte(key)))
		return nil

	})

	if err != nil {
		return value, err
	} else {
		return value, nil
	}

}

// ArrayInsertToBoltDB function will insert array of string
func ArrayInsertToBoltDB(task []string) error {
	err := DB.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("array"))
		if err != nil {
			return err
		}
		for _, val := range task {
			err = bucket.Put([]byte(val), []byte(val))
			if err != nil {
				continue
			}
		}

		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

// ArrayGetFromBoltDB will get the data from boltdb
func ArrayGetFromBoltDB(key []string) ([]string, error) {
	var value []string
	err := DB.View(func(tx *bolt.Tx) error {

		bucket := tx.Bucket([]byte("array"))
		if bucket == nil {
			return fmt.Errorf("bucket not found")
		}
		// Retrieve data from the bucket
		for _, k := range key {
			value = append(value, string(bucket.Get([]byte(k))))

		}
		return nil

	})

	if err != nil {
		return value, err
	} else {
		return value, nil
	}

}
