package main

import "fmt"

func fullName(firstName string, secondName string) (string) {
	fullName := firstName +" "+ secondName
	return fullName
}

func main()  {	
	fullName := fullName("Akshay", "Awate")
	fmt.Println(fullName)
}