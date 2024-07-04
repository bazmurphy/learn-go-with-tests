package main

import "math"

// TYPES ----------

// define our own type called Rectangle which encapsulates the concept
// create a simple type using a struct
// a struct is just a named collection of fields where you can store data
type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

// METHODS ----------

// we add methods to our types
// The syntax for declaring methods is almost the same as functions and that's because they're so similar.
// The only difference is the syntax of the method receiver func (receiverName ReceiverType) MethodName(args)
// When your method is called on a variable of that type, you get your reference to its data via the receiverName variable.
// In many other programming languages this is done implicitly and you access the receiver via `this`.
// It is a convention in Go to have the receiver variable be the first letter of the type.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

// INTERFACES ----------

// We tell Go what a Shape is using an interface declaration
// We create a new type just like we did with Rectangle and Circle but this time it is an interface rather than a struct

// This is quite different to interfaces in most other programming languages.
// Normally you have to write code to say My type Foo implements interface Bar.

// But in our case

// Rectangle has a method called Area that returns a float64 so it satisfies the Shape interface
// Circle has a method called Area that returns a float64 so it satisfies the Shape interface
// string does not have such a method, so it doesn't satisfy the interface
// etc.

// In Go interface resolution is implicit. If the type you pass in matches what the interface is asking for, it will compile.

type Shape interface {
	Area() float64
}

// ----------

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// replaced by the type/method/interface pattern above
// func Area(rectangle Rectangle) float64 {
// 	return rectangle.Width * rectangle.Height
// }
