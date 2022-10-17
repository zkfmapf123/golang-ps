package search

import (
	"errors"
	"testing"

	"github.com/magiconair/properties/assert"
)

func Test_linear_search(t *testing.T){
	
	_, e1 := FindNumByLinearSearch(42)
	_, e2 := FindNumByLinearSearch(100)

	assert.Equal(t,e1,nil)
	assert.Equal(t, e2, errors.New("Not Exists number"))	
}

func Test_search_hashMap(t *testing.T){
	
	FindNumByHashMap()

}

