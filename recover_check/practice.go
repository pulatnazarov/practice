package recover_check

import "fmt"

func Recover() {
	ALang := "GO Language"
	entry(&ALang, nil)
	fmt.Printf("Return successfully from the main function")
}

// This function is created to handle
// the panic occurs in entry function,
// but it does not handle the panic
// occurred in the entry function
// because it called in the normal
// function
func handlePanic() {

	if a := recover(); a != nil {
		fmt.Println("RECOVER", a)
	}
}

// Function
func entry(lang *string, aName *string) {

	// Normal function
	defer handlePanic()

	// When the value of lang
	// is nil it will panic
	if lang == nil {
		panic("Error: Language cannot be nil")
	}

	// When the value of aName
	// is nil it will panic
	if aName == nil {
		panic("Error: Author name cannot be nil")
	}

	fmt.Printf("Author Language: %s \n Author Name: %s\n", *lang, *aName)
	fmt.Printf("Return successfully from the entry function ")
}
