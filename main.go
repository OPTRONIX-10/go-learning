package main

import (
	"fmt"
	
)


func main(){
 nameAges := map[string]int{
	"James": 32,
	"Jill": 40,
	"Jack": 23,
	"Jenny": 26,
 }
 fmt.Print(nameAges)
 for k,v := range nameAges{
	fmt.Print(k," ",v)
 }
 nameAges["James"] = 33
 fmt.Println(nameAges)
}