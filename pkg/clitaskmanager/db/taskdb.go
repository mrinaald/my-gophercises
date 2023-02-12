package db

import (
	"encoding/binary"
	"fmt"
	"time"

	"github.com/boltdb/bolt"
)

var db *bolt.DB

var taskBucket []byte = []byte("tasks")

type Task struct {
	Key  int
	Task string
}

func InitailizeDB(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 2 * time.Second})
	if err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})
}

func CloseDB() {
	db.Close()
}

func AddTask(task string) (int, error) {
	var taskId int
	err := db.Update(func(tx *bolt.Tx) error {
		todos := tx.Bucket(taskBucket)

		id64, err := todos.NextSequence()
		if err != nil {
			return err
		}

		taskId = int(id64)
		return todos.Put(itob(taskId), []byte(task))
	})

	if err != nil {
		return -1, err
	}

	return taskId, nil

}

func GetTaskList() ([]Task, error) {
	var tasks []Task
	err := db.View(func(tx *bolt.Tx) error {
		todos := tx.Bucket(taskBucket)
		if todos == nil {
			return fmt.Errorf("Bucket %s does not exist in DB.", string(taskBucket))
		}

		cursor := todos.Cursor()
		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			tasks = append(tasks, Task{btoi(k), string(v)})
		}

		return nil
	})

	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func CompleteTask(taskIds []int) error {
	return db.Update(func(tx *bolt.Tx) error {
		todos := tx.Bucket(taskBucket)
		if todos == nil {
			return fmt.Errorf("Bucket %s does not exist in DB.", string(taskBucket))
		}

		for _, taskId := range taskIds {
			if err := todos.Delete(itob(taskId)); err != nil {
				return err
			}
		}

		return nil
	})
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
