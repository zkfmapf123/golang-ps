package main

import (
	"fmt"

	"github.com/zkfmapf123/collections"
)

func main(){
	n := collections.NewNode(1)
	n.ToString()

	n.Add(2)
	n.ToString()

	n.Add(3)
	n.ToString()

	n.Add(4)
	n.ToString()

	n.Add(5)
	n.ToString()
	
	fmt.Println(n.Search(1))
	fmt.Println(n.Search(2))
	fmt.Println(n.Search(10))
	fmt.Println(n.Search(20))

	n.Del(3)
	n.ToString()

	n.Add(10)
	n.ToString()
}