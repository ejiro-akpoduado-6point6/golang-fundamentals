package main

import "fmt"

type Engineer struct {
	firstName string
	lastName  string
	Age       int
	Project   Project
}

type Project struct {
	Name         string
	Priority     string
	Technologies []string
}

func (e Engineer) Print() {
	fmt.Println("Engineer Information")
	fmt.Printf("Name: %s\n", e.firstName)
	fmt.Printf("Name: %s\n", e.lastName)
	fmt.Printf("Name: %d\n", e.Age)
	fmt.Printf("Current Project: %s\n", e.Project.Name)
}

func (e *Engineer) PrintAge(age int) {
	 e.Age = age
	fmt.Println("Engineer Information")
	fmt.Printf("First Name: %s\n", e.firstName)
	fmt.Printf("Last Name: %s\n", e.lastName)
	fmt.Printf("Current age: %d\n", e.Age)
	fmt.Printf("Current Project: %s\n", e.Project.Name)
	fmt.Printf("Updated Age: %d\n", e.Age)
}

// Using a pointer receiver rather than a value receiver

func (e *Engineer) AddOneYear() {
	e.Age += 1
}

// a method that returns a value (string)
func (e *Engineer) GetProjectPriority() string {
	return e.Project.Priority
}



func main() {
	fmt.Println("Structs Tutorial")

	engineer := Engineer{
		firstName: "Ejiro",
		lastName:  "Akpoduado",
		Age:       27,
		Project: Project{
			Name: "Beginners guide to go",
			Priority: "High",
			Technologies: []string{"Go", "JavaScript"},
		},
	}

	// fmt.Println(engineer)
	// fmt.Printf("%+v\n", engineer)

	// fmt.Println(engineer.Project.Name)

	// engineer.Print()
	// engineer.Print()
	engineer.AddOneYear()
	engineer.Print()
	engineer.PrintAge(34)
	fmt.Println(engineer.GetProjectPriority())




}
