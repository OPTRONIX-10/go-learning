package main

import (
	"fmt"
	
)


func main(){

	myBill := newBill("Chris")
	myBill.updateTip(10)
	myBill.addItem("Onion Soup", 4.50)
	myBill.addItem("Meatballs", 8.50)
	myBill.addItem("Pasta", 12.75)
	fmt.Println(myBill.format())
 
}