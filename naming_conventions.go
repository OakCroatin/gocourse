package main

import "fmt"

type Employee struct {
	firstName string
	lastName  string
	age       int
}

func namingConvention() {
	//PascalCase - for Strucs, enums,
	//snake_case - for constants, variables, and file names.
	//UPPERCASE - Constants
	//mixedCase - variables
	// Be consistent in nameing conventions.
	//lowercase - package names.
	const MAXRETRIES = 5
	var employeeId = "1001"
	fmt.Println("Employee ID ", employeeId)
}
