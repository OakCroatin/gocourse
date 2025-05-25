package basics

import "fmt"

var middleName = "Count"

func main() {
	//var age int
	//var name string = "Jane"
	//examples of different variable declaration styles.
	//Here a string is declared without naming it string
	//var name1 = "John"
	//You can declare variables by simply using := as seen below
	//count := 10
	//store := "Mary's Autobody"
	firstName()
	fmt.Println(middleName)
}

func firstName() {
	firstName := "Michael"
	fmt.Println(firstName)
}
