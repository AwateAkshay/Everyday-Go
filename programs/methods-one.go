package main
import "fmt"

type Person struct {
	name string
	age  int
}

func (s Person) getAge() int{
    return s.age
}

func main() {
	p := Person{"Akshay", 26}
	age := p.getAge()
	fmt.Println(age)
	
}