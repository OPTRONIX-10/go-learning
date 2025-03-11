package main

import (
	"fmt"
	
)

func updateString(s *string){
	*s = "Hello"
}
func main(){
 name := "Abhishek"
 newPointer := &name
 fmt.Println("Name: ", name)
 updateString(newPointer)
 fmt.Println("Name: ", name)
}