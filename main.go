package main

import "fmt"

//Structs

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	janella := person{
		firstName: "Janella",
		lastName:  "Mina",
		contactInfo: contactInfo{
			email:   "5701ja.mi@gmail.com",
			zipCode: 1810,
		},
	}

	// & points to the address in memory
	janellaPointer := &janella
	janellaPointer.updateName("Jami")
	janella.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// * is the value the memory address is pointing at
func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}
