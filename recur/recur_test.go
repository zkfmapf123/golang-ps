package recur

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func Test_recur(t *testing.T) {
	n1 := Fibo(10)

	assert.Equal(t,n1,55)
}