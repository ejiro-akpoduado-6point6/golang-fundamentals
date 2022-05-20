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

	fmt.Println(engineer)
	fmt.Printf("%+v\n", engineer)

	fmt.Println(engineer.Project.Name)
}
