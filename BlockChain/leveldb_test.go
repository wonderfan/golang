package lang

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"os"
	"testing"
)

func TestDB(t *testing.T) {
	fmt.Println(">>> start the database")
	db, err := leveldb.OpenFile("stateLeveldb", nil)
	if err == nil {
		defer db.Close()
	}
	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		key := iter.Key()
		//value := iter.Value()
		fmt.Println(string(key))
	}
	iter.Release()
	err = iter.Error()
	fmt.Println(err == nil)
}

func TestMain(m *testing.M) {
	fmt.Println(">>> setup the test suite")
	exitCode := m.Run()
	os.Exit(exitCode)
}
