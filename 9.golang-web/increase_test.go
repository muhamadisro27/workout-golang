package golang_web

import (
	"fmt"
	"testing"
)

type Person struct {
    Name string
    Age  int
}

// Metode dengan receiver pointer
func (p *Person) IncreaseAge() {
    p.Age++
}

func TestIncrease(t *testing.T) {
    person := Person{Name: "John", Age: 30}
    person.IncreaseAge()
    person.IncreaseAge()
    fmt.Println(person.Age)
}