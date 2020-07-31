package main

import "fmt"

// Specific interfaces
type Interface1 interface {
	Method1()
}

type Interface2 interface {
	Method2()
}

// Composition of interfaces
type ComposedInterface interface {
	Interface1
	Interface2
}

// Struct that implements both Interface 1 and Interface2
type Implementation struct{}

func (Implementation) Method1() {
	fmt.Println("Called Implementation.Method1")
}

func (Implementation) Method2() {
	fmt.Println("Called Implementation.Method2")
}

// This method expects an object that implements
// both Interface 1 and Interface,
// via ComposedInterface
func CallBoth(o ComposedInterface) {
	// Here we can call both Method1 and Method2,
	// because ComposedInterface is composed of
	// both Interface 1 and Interface2
	o.Method1()
	o.Method2()
	// We can also call methods that expect Interface1
	CallMethod1(o)
	// And also methods that expect Interface2
	CallMethod2(o)
}

func CallMethod1(o Interface1) {
	o.Method1()
}

func CallMethod2(o Interface2) {
	o.Method2()
}

// Main function
func main() {
	i := Implementation{}
	// Call function that expects a ComposedInterface
	CallBoth(i)
}

// Output:
//
// Called Implementation.Method1
// Called Implementation.Method2
// Called Implementation.Method1
// Called Implementation.Method2
