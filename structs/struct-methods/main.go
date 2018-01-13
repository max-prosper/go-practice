package main

// Person is really good struct :)
type Person struct {
	ID   int
	Name string
}

// UpdateName won't change an original stuct
func (p Person) UpdateName(name string) {
	p.Name = name
}

// SetName will change an original stuct
func (p *Person) SetName(name string) {
	p.Name = name
}

// Account is anotger really good struct :)
type Account struct {
	ID   int
	Name string
	Person
}

func main() {
	// pers := &Person{1, "Max"}
	pers := new(Person)
	pers.SetName("Max Prosper")
	// (&pers).SetName("Max Prosper")
	// fmt.Printf("updated person: %#v\n", pers)

	var acc = Account{
		ID:   1,
		Name: "Max",
		Person: Person{
			ID:   2,
			Name: "Max Prosper",
		},
	}

	acc.SetName("Max.Prosper")
}
