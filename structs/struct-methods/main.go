package main

import "fmt"

type Person struct {
	Id   int
	Name string
}

// Won't change an original stuct
func (p Person) UpdateName(name string) {
	p.Name = name
}

// Will change an original stuct
func (p *Person) SetName(name string) {
	p.Name = name
}

type Account struct {
	Id   int
	Name string
	Person
}

func main() {
	// pers := &Person{1, "Max"}
	pers := new(Person)
	pers.SetName("Max Prosper")
	// (&pers).SetName("Max Prosper")
	fmt.Printf("updated person: %#v\n", pers)

	var acc Account = Account{
		Id:   1,
		Name: "Max",
		Person: Person{
			Id:   2,
			Name: "Max Prosper",
		},
	}

	acc.SetName("Max.Prosper")
}
