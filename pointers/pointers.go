package main

import "fmt"

func changeVarValueByReference(var1 *int, newValue int) {
	*var1 = newValue
}

func main() {
	var var1 int = 100
	var var2 int = var1 // Copying the value of var1 to var2

	fmt.Println("var1:", var1)
	fmt.Println("var2:", var2)
	
	var1++

	fmt.Println("\nnew value of var1:", var1)
	fmt.Println("var2:", var2)

	var pointerToVar1 *int;
	pointerToVar1 = &var1 // Assigning the memory address of var1 to pointerToVar1

	fmt.Println("\nPointer to var1:", pointerToVar1)
	fmt.Println("Value at memory address of var1:", *pointerToVar1)

	changeVarValueByReference(&var1, 200)
	fmt.Println("\nvar1 after changeVarValue:", var1)
}