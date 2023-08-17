package db

import "os"

const (
	TestDbPath = "test.db"
)

func Cleanup(path string) {
	_ = os.Remove(path)
}
