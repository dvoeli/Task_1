package main

import "fmt"

// We can also declare variables outside of the main function but this requires the notation var + name + type
var i int = 42

// We can also create so-called var blocks, which are just used when wanting to declare several variables simultaneously.
var (
	first_name string = "Eli"
	last_name string = "Dvoretsky"
	age int = 16
)

// In GoLang a letter that is capitalized is not another variable. Capitalization also has a meaning in Go.
//If a letter is not capitalized, it is confined to the package and cannot be used in another package.
//If it is capitalized, it is globally available.

func main() {
	// In GoLang one can declare variables using the := operator.
	//This allows one to declare the variable without explicitly saying what type the variable is.
	// All variables declared in Go inside a function block are only usable in that block.
	i := "Hello World"
	fmt.Println(i)
	fmt.Printf("%v, %T", i, i)

	// In the below block I initialise the new variables j and k.
	//I then assign the value of j, 42, which is an integer to k, turning it into a float32 value using float32().
	var j int = 42
	var k float32
	k = float32(j)
	fmt.Printf("%v, %T", k, k)


}

