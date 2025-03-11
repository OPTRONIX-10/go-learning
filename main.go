package main

import (
	"fmt"
	"strings"
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

func returnTwoStrings(name string)(string, string){
	s := strings.ToUpper(name)
	names := strings.Split(s, " ")
	var nameSlice []string
	for _,val :=range names{
		nameSlice = append(nameSlice, val[:1])
	}
	if len(nameSlice) > 1{
		return nameSlice[0], nameSlice[1]
	}
	return nameSlice[0], "_"
}
func main(){
	greeting("john")
	sample([]string{"john","doe","jane"}, greeting)
	name := returnName("john", "doe")
	f1,f2 := returnTwoStrings("john doe")
	
	fmt.Println(name, f1, f2)
	//scope of pacakge
	sampleNames("john")


}