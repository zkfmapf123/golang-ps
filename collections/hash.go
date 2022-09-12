package collections

import (
	"errors"
	"hash/fnv"
)

/*
	Hash-Map

	use LinkedList || Tree

	* LinkedList -> 데이터 저장이 많지 않다면 O(n)
	* Tree -> 데이터 저장이 많다면 O(logN)
	* 편의상 Generic 안씀

	Average
		- Search O(1) ***
		- Insert O(1)
		- Delete O(1)

	Worst
		- Search O(1)
		- Insert O(n)
		- Delete O(n)

	** More
	- Go 에서는 Separate Chaining 방법을 사용 (use Array)
	- 한 Bucket에는 8개의 슬롯이 존재
	- Go HashTable 슬롯 사용이 6.5개 이상이라면, 해시테이블 사이즈를 두배로 늘리는 Growing 과정을 수행

	** Implementation
	- Use Compact Hash Function
*/

const (
	BUCKET_COUNT = 10
)

type bucket struct {
	k string
	v int
}

type hashMap struct {
	b []bucket
}

func HashMap() *hashMap{
	return &hashMap{
		b : make([]bucket, BUCKET_COUNT, BUCKET_COUNT),
	}
}

func (h *hashMap) Set(k string, v int){
	hashNum := hashTable(k, len(h.b))
	b := bucket{
		k : k, 
		v : v,
	}

	h.b[hashNum] = b
}

func (h *hashMap) Get(k string) (int, error){
	hashNum := hashTable(k, len(h.b))
	num := h.b[hashNum]

	if num.k == "" {
		return 0, errors.New("not exists value")
	}

	return num.v, nil
}

/*
* @todo more Search -> Hash Function
*/
func hashTable(k string, len int) int {
	h := fnv.New32a()
	h.Write([]byte(k))
	return int(h.Sum32()) % len
}


