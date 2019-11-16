package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

// being used inside the person struct
type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	garrett := person{firstName: "Garrett", lastName: "Halstein"}

	fmt.Println(garrett)

	fmt.Printf("%+v", garrett)

	logan := person{
		firstName: "Logan",
		lastName:  "Schwartz",
		contactInfo: contactInfo{
			email:   "logan@gmail.com",
			zipCode: 11746,
		},
	}
	fmt.Println()
	logan.print()

}

func (ptrToP *person) updateName(newFirstName string) {
	(*ptrToP).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
