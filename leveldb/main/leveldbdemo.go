package main

import (
	"github.com/syndtr/goleveldb/leveldb"
	"fmt"
	"strconv"
	"github.com/syndtr/goleveldb/leveldb/util"
)

func main() {
	db, err := leveldb.OpenFile("leveldb/path/db", nil)
	if err != nil {
		fmt.Println("create leveldb failed!")
	}
	defer db.Close()

	////不使用内存数据，直接取硬盘数据:
	//o := &opt.Options{
	//	Filter: filter.NewBloomFilter(10),
	//}
	//db, err := leveldb.OpenFile("leveldb/path/db", o)
	//defer db.Close()

	//插入
	err = db.Put([]byte("key"), []byte("value"), nil)
	if err != nil {
		fmt.Println("insert leveldb failed!")
	}
	//获取
	// 注意：返回数据结果不能修改
	data, err := db.Get([]byte("key"), nil)
	if err != nil {
		fmt.Println("get leveldb failed!")
	}
	fmt.Println(string(data))
	//删除
	err = db.Delete([]byte("key"), nil)
	if err != nil {
		fmt.Println("delete leveldb failed!")
	}

	//for 循环 插入数据
	for i := 0; i < 10; i++ {
		tempkey := "key-" + strconv.Itoa(i)
		tempvalue := "value-" + strconv.Itoa(i)
		db.Put([]byte(tempkey), []byte(tempvalue), nil)
	}

	//批量写入数据
	batch := new(leveldb.Batch)
	batch.Put([]byte("foo"), []byte("value"))
	batch.Put([]byte("bar"), []byte("another value"))
	batch.Delete([]byte("baz"))
	err = db.Write(batch, nil)

	//数据库遍历
	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()
		fmt.Printf("key = %s, value = %s \n", string(key), string(value))
	}
	iter.Release()
	err = iter.Error()
	if err != nil {
		fmt.Printf("err = %s", err)
	}

	//迭代查询
	iter = db.NewIterator(nil, nil)
	for ok := iter.Seek([]byte("key-8")); ok; ok = iter.Next() {
		key := iter.Key()
		value := iter.Value()
		fmt.Printf("partly key = %s, value = %s \n", string(key), string(value))
	}
	iter.Release()
	err = iter.Error()
	if err != nil {
		fmt.Printf("err = %s", err)
	}

	//按区间查询
	iter = db.NewIterator(&util.Range{Start: []byte("key-2"), Limit: []byte("key-6")}, nil)
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()
		fmt.Printf("range key = %s, value = %s \n", string(key), string(value))
	}
	iter.Release()
	err = iter.Error()
	if err != nil {
		fmt.Printf("err = %s", err)
	}

	//按前缀查询
	iter = db.NewIterator(util.BytesPrefix([]byte("key-")), nil)
	for iter.Next() {
		// Use key/value.
		key := iter.Key()
		value := iter.Value()
		fmt.Printf("prefix key = %s, value = %s \n", string(key), string(value))
	}
	iter.Release()
	err = iter.Error()
	if err != nil {
		fmt.Printf("err = %s", err)
	}

}
