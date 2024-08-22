package main

import "tests-introduction/addresses"

func main() {
	addressType := addresses.GetAddressType("Rua dos Bobos")
	println(addressType)
}