package clitaskmanager

import (
	"encoding/binary"
	"fmt"
	"os"
	"time"

	"github.com/boltdb/bolt"
)

var db *bolt.DB

var TABLE_TODOS []byte = []byte("todos")

func InitailizeDB() {
	var err error
	db, err = bolt.Open("tasks.db", 0600, &bolt.Options{Timeout: 2 * time.Second})
	if err != nil {
		panic(err)
	}
}

func FinalizeDB() {
	db.Close()
}

func AddTaskInDB(task string) {
	err := db.Update(func(tx *bolt.Tx) error {
		todos, err := tx.CreateBucketIfNotExists(TABLE_TODOS)
		if err != nil {
			return err
		}

		taskId, err := todos.NextSequence()
		if err != nil {
			return err
		}

		binaryTaskId := make([]byte, 8)
		binary.BigEndian.PutUint64(binaryTaskId, taskId)
		if err := todos.Put(binaryTaskId, []byte(task)); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("Added \"%s\" to your task list.\n", task)
}

func ListTasksFromDB() {
	err := db.View(func(tx *bolt.Tx) error {
		todos := tx.Bucket(TABLE_TODOS)
		if todos == nil {
			return nil // Don't output anything
		}

		fmt.Println("You have the following tasks:")
		taskCount := 1
		cursor := todos.Cursor()
		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			fmt.Printf("%d. %s\n", taskCount, string(v))
			taskCount++
		}

		return nil
	})

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func CompleteTaskInDB(taskIds []int) {
	err := db.Update(func(tx *bolt.Tx) error {
		todos := tx.Bucket(TABLE_TODOS)
		if todos == nil {
			return fmt.Errorf("Bucket %s does not exist in DB.", string(TABLE_TODOS))
		}

		var keys, vals [][]byte
		cursor := todos.Cursor()
		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			keys = append(keys, k)
			vals = append(vals, v)
		}

		for _, taskId := range taskIds {
			if taskId <= 0 || taskId > len(vals) {
				continue
			}

			fmt.Printf("You have completed the \"%s\" task.\n", string(vals[taskId - 1]))
			todos.Delete(keys[taskId - 1])
		}

		return nil
	})

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
