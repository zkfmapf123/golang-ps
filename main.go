package main

import (
	"fmt"

	"github.com/zkfmapf123/collections"
)

func main(){
	age := collections.HashMap()
	age.Set("leedonggyu",29)
	age.Set("limjeahyock",28)

	fmt.Println(age)

	v1, err1:= age.Get("leedonggyu")

	fmt.Println(v1 ,err1)


	v2, err2 := age.Get("err")
	fmt.Println(v2, err2)
}