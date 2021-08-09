package main
import "fmt"

type Person struct {
	name string
	age int
}

func (s *Person) setAge(age int) {
	s.age = age
}

func main() {
   
	s := Person{"Akshay", 24}
	fmt.Println("Before", s.age)
	s.setAge(25)
	fmt.Println("After",s.age)
	
	
}