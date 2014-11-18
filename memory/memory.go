package memory

import (
	"github.com/syndtr/goleveldb/leveldb"
)

const (
	DBDir = "diegobot_mem"
)

func GetDatabase() *leveldb.DB {
	db, err := leveldb.OpenFile(DBDir, nil)

	if err != nil {
		panic(err)
	}
	return db
}

func GetKey(key string) string {
	db := GetDatabase()
	defer db.Close()

	data, err := db.Get([]byte(key), nil)

	if err != nil {
		panic(err)
	}
	return string(data)
}

func PutKey(key, value string) {
	db := GetDatabase()
	defer db.Close()
	err := db.Put([]byte(key), []byte(value), nil)

	if err != nil {
		panic(err)
	}
}

func DeleteKey(key string) {
	db := GetDatabase()
	defer db.Close()
	db.Delete([]byte(key), nil)
}

// TODO: Implement batch write
// TODO: implement iterator on prefix
