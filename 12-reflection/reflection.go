package reflection_chapter

import "reflect"

// func walk(x interface{}, fn func(input string)) {
// 	fn("I still can't believe South Korea beat Germany 2-0 to put them last in their group")
// }

// func walk(x interface{}, fn func(input string)) {
// 	val := reflect.ValueOf(x)
// 	field := val.Field(0)
// 	fn(field.String())
// }

// This code is very unsafe and very naive, but remember: our goal when we are in "red" (the tests failing) is to write the smallest amount of code possible. We then write more tests to address our concerns.
// We need to use reflection to have a look at x and try and look at its properties.
// The reflect package has a function ValueOf which returns us a Value of a given variable. This has ways for us to inspect a value, including its fields which we use on the next line.
// We then make some very optimistic assumptions about the value passed in:
// We look at the first and only field. However, there may be no fields at all, which would cause a panic.
// We then call String(), which returns the underlying value as a string. However, this would be wrong if the field was something other than a string.

// -----------

// val has a method NumField which returns the number of fields in the value. This lets us iterate over the fields and call fn which passes our test.

// func walk(x interface{}, fn func(input string)) {
// 	val := reflect.ValueOf(x)

// 	for i := 0; i < val.NumField(); i++ {
// 		field := val.Field(i)
// 		fn(field.String())
// 	}
// }

// ----------

// We need to check that the type of the field is a string.
// We can do that by checking its Kind.

// func walk(x interface{}, fn func(input string)) {
// 	val := reflect.ValueOf(x)

// 	for i := 0; i < val.NumField(); i++ {
// 		field := val.Field(i)

// 		if field.Kind() == reflect.String {
// 			fn(field.String())
// 		}
// 	}
// }

// The `Kind()` method in this context is part of Go's reflection package. It returns the specific kind of type that a `reflect.Value` represents.
// 1. `field := val.Field(i)` gets the i-th field of the struct value.
// 2. `field.Kind()` returns the kind of this field.
// 3. `reflect.String` is a constant in the `reflect` package representing the string kind.
// 4. The condition `if field.Kind() == reflect.String` checks if the current field is of type string.
// The `Kind()` method is crucial for type checking in reflection. It allows you to determine the underlying type of a `reflect.Value`, which is especially useful when working with interfaces or when you don't know the exact type at compile time.
// In Go, `Kind` represents the specific kind of type, as opposed to the type itself. For example, `int`, `int32`, and `int64` are different types, but they all have the kind `reflect.Int`. Similarly, all string types have the kind `reflect.String`.
// The possible kinds include basic types (like Bool, Int, Float64, String), as well as composite types (like Array, Slice, Map, Struct), and special types (like Interface, Func).
// In this code, `Kind()` is used to identify which fields of the struct are strings, so that the provided function can be applied only to string fields.

// func walk(x interface{}, fn func(input string)) {
// 	val := reflect.ValueOf(x)

// 	for i := 0; i < val.NumField(); i++ {
// 		field := val.Field(i)

// 		if field.Kind() == reflect.String {
// 			fn(field.String())
// 		}

// 		if field.Kind() == reflect.Struct {
// 			walk(field.Interface(), fn)
// 		}
// 	}
// }

// The solution is quite simple, we again inspect its Kind and if it happens to be a struct we just call walk again on that inner struct.

// ------------

// When you're doing a comparison on the same value more than once generally refactoring into a switch will improve readability and make your code easier to extend.
// What if the value of the struct passed in is a pointer?

// func walk(x interface{}, fn func(input string)) {
// 	val := reflect.ValueOf(x)

// 	for i := 0; i < val.NumField(); i++ {
// 		field := val.Field(i)

// 		switch field.Kind() {
// 		case reflect.String:
// 			fn(field.String())
// 		case reflect.Struct:
// 			walk(field.Interface(), fn)
// 		}
// 	}
// }

// When you're doing a comparison on the same value more than once generally refactoring into a switch will improve readability and make your code easier to extend.

// ----------

// You can't use NumField on a pointer Value, we need to extract the underlying value before we can do that by using Elem().

func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Struct:
			walk(field.Interface(), fn)
		}
	}
}

// ----------

// Let's encapsulate the responsibility of extracting the reflect.Value from a given interface{} into a function.
