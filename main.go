package main

import (
	"fmt"
	// "sort"
	// "strings"
)


func main(){
	// numbers, variables, strings

	// var word string = "Hello World"
	// newName := "Hello World"
	// var age  int = 20
	// newAge := 20
	// var price float32 = 20.5
	// var newPrice float64 = 22220000.322222
	// fmt.Println(word,newName,age,newAge, price,newPrice)

	// printing methods

	// age := 20.555
	// name := "John"
	// fmt.Print("My age is",age)
	// fmt.Print("My name is",name)
	// fmt.Println("My age is",age)
	// fmt.Println("My name is",name)
	// fmt.Printf("My age is %f  and my name is %s\n",age,name)
	// var word = fmt.Sprintf("My age is %f  and my name is %s",age,name)
	// fmt.Println(word)

	//arrays, alices

	// var ages [6]int = [6]int{1,2,3,4,5,6}
	// var newAges = [6]int{1,2,3,4,5}
	// newAges[2] = 10
	// var scores = []int{1,2,3,4,5,6}
	// scores = append(scores, 10)
	// range1 := scores[2:5]
	// fmt.Println(ages,newAges,scores,range1,len(ages),len(newAges))

	//Standard packagaes

	// greeetings:= "helllow good morning"
	// fmt.Println(strings.Contains(greeetings,"morning"))
	// fmt.Println(strings.ToUpper(greeetings))
	// ages := []int{20,43,11,65,12}
	// sort.Ints(ages)
	// fmt.Println(ages)

	// loops

	// x:= 0
	// for x<5{
	// 	fmt.Print(x)
	// 	x++
	// }
	// for i :=0; i<5;i++{
	// 	fmt.Print(i)
	// }
	// ages:= []int{20,43,11,65,12}
	// for index, value :=range ages{
	// 	fmt.Println(index,value)
	// }
	// for _,value :=range ages{
	// 	fmt.Println(value)
	// }

	// booleans and conditions

	age :=45
	if age <20{
		fmt.Println("You are young")
	}else if age <40{
		fmt.Println("You are old")
	}else{
		fmt.Println("You are very old")
	}



}