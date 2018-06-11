package main

import "fmt"

//Person struct is a representation of a person
type Person struct {
	firstName string
	lastName  string
	contact   ContactInfo
}

//ContactInfo info for a person
type ContactInfo struct {
	email   string
	zipcode int
}

func createNewPerson(first string, second string) Person {
	return Person{first, second, ContactInfo{"anujva@gmail.com", 800016}}
}

func (person Person) print() {
	fmt.Println(person.firstName + " " + person.lastName)
	fmt.Println(person.contact)
}

func (person *Person) newFirstName(first string) {
	person.firstName = first
}

func main() {
	alex := createNewPerson("Alex", "Anderson")
	fmt.Println(alex)
	alex.print()
	alex.newFirstName("Jimmy")
	fmt.Printf("%+v", alex)
	fmt.Println()
}
