package main

import (
	"fmt"
	// "sort"
	// "strings"
)


func sample (names []string, fn func(string)){
	for _,valus :=range names{
		fn(valus)
	}

}

func greeting(name string){
	fmt.Println("Hello", name)
	
}

func returnName(name1 string, name2 string)string{
	return name1 + " " + name2
}
func main(){
	greeting("john")
	sample([]string{"john","doe","jane"}, greeting)
	name := returnName("john", "doe")
	fmt.Println(name)


}