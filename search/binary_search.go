package search

import (
	"errors"
	"fmt"
)

var arr = []int{5,12,36,39,40,42,29,61,80,95,101,222,317,358,804,999}

/*
	linear O(n)
*/
func FindNumByLinearSearch(num int) (int, error){

	for _, v := range arr {
		if v == num {
			return num, nil
		}
	}

	return 0, errors.New("Not Exists number")
}

/*
	use HashMap O(1)
*/
func FindNumByHashMap(){
	m := newHashMap()
	addArr(*m, arr)
	m.printAll()
}

type HashMap map[int]int

func newHashMap() *HashMap{
	hashMap := HashMap{}
	return &hashMap
}

func addArr(m HashMap, list []int){
	
	len := len(list)
	for _,v :=range list {
		hash := m.hashTable(v, len)
		
		m[hash]= v
	}
}

func (m *HashMap) hashTable(key int, len int) int {
	return key % len
}

func (m *HashMap) printAll(){

	for key, elem := range *m {
		fmt.Println(key, elem)
	}
}	


