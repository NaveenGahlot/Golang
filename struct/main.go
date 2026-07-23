package main

import "fmt"

// Structs in Go are used to group together related data. They are similar to classes in other programming languages, but they do not have methods or inheritance. Structs can contain fields of different types, including other structs.

// In this example, we define a Person struct with fields for first name, last name, and age. We also define a Contact struct with fields for email and phone number, and an Address struct with fields for street, city, state, and zip code. Finally, we define an Employee struct that contains a Person, Contact, and Address struct, as well as an employee ID.

// Structs can be created using the var keyword, the new keyword, or by using a struct literal. In this example, we create a Person struct using the var keyword, a Person struct using the new keyword, and a Person struct using a struct literal. We also create an Employee struct and populate its fields with data.

// Note that the lastName field in the Person struct is unexported (lowercase), which means it cannot be accessed from outside the package. This is a common practice in Go to encapsulate data and prevent direct access to certain fields.
type Person struct {
	FirstName string
	lastName  string
	Age       int
}

type Contact struct {
	Email string
	Phone string
}

type Address struct {
	Street  string
	City    string
	State   string
	ZipCode string
}

type Employee struct {
	Person_Details  Person
	Contact_Details Contact
	Address_Details Address
	EmployeeID      string
}

func main() {
	// Creating a Person struct using the var keyword
	var prince Person
	fmt.Println("Prince Person:- ", prince)
	prince.FirstName = "Prince"
	prince.lastName = "Kumar"
	prince.Age = 25
	fmt.Println("Prince Person:- ", prince)

	// Creating a Person struct using a struct literal
	person1 := Person{
		FirstName: "John",
		lastName:  "Doe",
		Age:       30,
	}
	fmt.Println("Person1:- ", person1)

	// Creating a Person struct using the new keyword
	var person2 = new(Person)
	person2.FirstName = "Jane"
	person2.lastName = "Smith"
	person2.Age = 28
	fmt.Println("Person2:- ", person2)

	// Creating an Employee struct and populating its fields
	var employee Employee
	employee.Person_Details = Person{
		FirstName: "Alice",
		lastName:  "Johnson",
		Age:       35,
	}
	employee.Contact_Details = Contact{
		Email: "alice@gmail.com",
		Phone: "123-456-7890",
	}
	employee.Address_Details = Address{
		Street:  "123 Main St",
		City:    "New York",
		State:   "NY",
		ZipCode: "10001",
	}
	employee.EmployeeID = "E001"
	fmt.Println("Employee:- ", employee)
}
