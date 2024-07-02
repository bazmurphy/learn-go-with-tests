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

// func walk(x interface{}, fn func(input string)) {
// 	val := reflect.ValueOf(x)

// 	if val.Kind() == reflect.Pointer {
// 		val = val.Elem()
// 	}

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

// ----------

// Let's encapsulate the responsibility of extracting the reflect.Value from a given interface{} into a function.

// func walk(x interface{}, fn func(input string)) {
// 	val := getValue(x)

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

// func getValue(x interface{}) reflect.Value {
// 	val := reflect.ValueOf(x)

// 	if val.Kind() == reflect.Pointer {
// 		val = val.Elem()
// 	}

// 	return val
// }

// This actually adds more code but I feel the abstraction level is right.
// Get the reflect.Value of x so I can inspect it, I don't care how.
// Iterate over the fields, doing whatever needs to be done depending on its type.

// ----------

// Next, we need to cover slices.

// func walk(x interface{}, fn func(input string)) {
// 	val := getValue(x)

// 	if val.Kind() == reflect.Slice {
// 		for i := 0; i < val.Len(); i++ {
// 			walk(val.Index(i).Interface(), fn)
// 		}
// 		return
// 	}

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

// func getValue(x interface{}) reflect.Value {
// 	val := reflect.ValueOf(x)

// 	if val.Kind() == reflect.Pointer {
// 		val = val.Elem()
// 	}

// 	return val
// }

// ----------

// If you think a little abstractly, we want to call walk on either
// - Each field in a struct
// - Each thing in a slice
// Our code at the moment does this but doesn't reflect it very well. We just have a check at the start to see if it's a slice (with a return to stop the rest of the code executing) and if it's not we just assume it's a struct.
// Let's rework the code so instead we check the type first and then do our work.

// func walk(x interface{}, fn func(input string)) {
// 	val := getValue(x)

// 	switch val.Kind() {
// 	case reflect.Struct:
// 		for i := 0; i < val.NumField(); i++ {
// 			walk(val.Field(i).Interface(), fn)
// 		}
// 	case reflect.Slice:
// 		for i := 0; i < val.Len(); i++ {
// 			walk(val.Index(i).Interface(), fn)
// 		}
// 	case reflect.String:
// 		fn(val.String())
// 	}
// }

// func getValue(x interface{}) reflect.Value {
// 	val := reflect.ValueOf(x)

// 	if val.Kind() == reflect.Pointer {
// 		val = val.Elem()
// 	}

// 	return val
// }

// Looking much better! If it's a struct or a slice we iterate over its values calling walk on each one. Otherwise, if it's a reflect.String we can call fn.
// Still, to me it feels like it could be better. There's repetition of the operation of iterating over fields/values and then calling walk but conceptually they're the same.

// func walk(x interface{}, fn func(input string)) {
// 	val := getValue(x)

// 	numberOfValues := 0
// 	var getField func(int) reflect.Value

// 	switch val.Kind() {
// 	case reflect.String:
// 		fn(val.String())
// 	case reflect.Struct:
// 		numberOfValues = val.NumField()
// 		getField = val.Field
// 	case reflect.Slice:
// 		numberOfValues = val.Len()
// 		getField = val.Index
// 	}

// 	for i := 0; i < numberOfValues; i++ {
// 		walk(getField(i).Interface(), fn)
// 	}
// }

// func getValue(x interface{}) reflect.Value {
// 	val := reflect.ValueOf(x)

// 	if val.Kind() == reflect.Pointer {
// 		val = val.Elem()
// 	}

// 	return val
// }

// If the value is a reflect.String then we just call fn like normal.
// Otherwise, our switch will extract out two things depending on the type
// - How many fields there are
// - How to extract the Value (Field or Index)
// Once we've determined those things we can iterate through numberOfValues calling walk with the result of the getField function.
// Now we've done this, handling arrays should be trivial.

// func walk(x interface{}, fn func(input string)) {
// 	val := getValue(x)

// 	numberOfValues := 0
// 	var getField func(int) reflect.Value

// 	switch val.Kind() {
// 	case reflect.String:
// 		fn(val.String())
// 	case reflect.Struct:
// 		numberOfValues = val.NumField()
// 		getField = val.Field
// 	case reflect.Slice, reflect.Array:
// 		numberOfValues = val.Len()
// 		getField = val.Index
// 	}

// 	for i := 0; i < numberOfValues; i++ {
// 		walk(getField(i).Interface(), fn)
// 	}
// }

// func getValue(x interface{}) reflect.Value {
// 	val := reflect.ValueOf(x)

// 	if val.Kind() == reflect.Pointer {
// 		val = val.Elem()
// 	}

// 	return val
// }

// Again if you think a little abstractly you can see that map is very similar to struct, it's just the keys are unknown at compile time.

// func walk(x interface{}, fn func(input string)) {
// 	val := getValue(x)

// 	numberOfValues := 0

// 	var getField func(int) reflect.Value

// 	switch val.Kind() {
// 	case reflect.String:
// 		fn(val.String())
// 	case reflect.Struct:
// 		numberOfValues = val.NumField()
// 		getField = val.Field
// 	case reflect.Slice, reflect.Array:
// 		numberOfValues = val.Len()
// 		getField = val.Index
// 	case reflect.Map:
// 		for _, key := range val.MapKeys() {
// 			walk(val.MapIndex(key).Interface(), fn)
// 		}
// 	}

// 	for i := 0; i < numberOfValues; i++ {
// 		walk(getField(i).Interface(), fn)
// 	}
// }

// func getValue(x interface{}) reflect.Value {
// 	val := reflect.ValueOf(x)

// 	if val.Kind() == reflect.Pointer {
// 		val = val.Elem()
// 	}

// 	return val
// }

// However, by design you cannot get values out of a map by index. It's only done by key, so that breaks our abstraction, darn.
// How do you feel right now? It felt like maybe a nice abstraction at the time but now the code feels a little wonky.
// This is OK! Refactoring is a journey and sometimes we will make mistakes. A major point of TDD is it gives us the freedom to try these things out.
// By taking small steps backed by tests this is in no way an irreversible situation. Let's just put it back to how it was before the refactor.

// func walk(x interface{}, fn func(input string)) {
// 	val := getValue(x)

// 	walkValue := func(value reflect.Value) {
// 		walk(value.Interface(), fn)
// 	}

// 	switch val.Kind() {
// 	case reflect.String:
// 		fn(val.String())
// 	case reflect.Struct:
// 		for i := 0; i < val.NumField(); i++ {
// 			walkValue(val.Field(i))
// 		}
// 	case reflect.Slice, reflect.Array:
// 		for i := 0; i < val.Len(); i++ {
// 			walkValue(val.Index(i))
// 		}
// 	case reflect.Map:
// 		for _, key := range val.MapKeys() {
// 			walk(val.MapIndex(key).Interface(), fn)
// 		}
// 	}
// }

// func getValue(x interface{}) reflect.Value {
// 	val := reflect.ValueOf(x)

// 	if val.Kind() == reflect.Pointer {
// 		val = val.Elem()
// 	}

// 	return val
// }

// We've introduced walkValue which DRYs up the calls to walk inside our switch so that they only have to extract out the reflect.Values from val.

// ----------

// Remember that maps in Go do not guarantee order. So your tests will sometimes fail because we assert that the calls to fn are done in a particular order.

// ----------

// We can iterate through all values sent through channel until it was closed with Recv()

// func walk(x interface{}, fn func(input string)) {
// 	val := getValue(x)

// 	walkValue := func(value reflect.Value) {
// 		walk(value.Interface(), fn)
// 	}

// 	switch val.Kind() {
// 	case reflect.String:
// 		fn(val.String())
// 	case reflect.Struct:
// 		for i := 0; i < val.NumField(); i++ {
// 			walkValue(val.Field(i))
// 		}
// 	case reflect.Slice, reflect.Array:
// 		for i := 0; i < val.Len(); i++ {
// 			walkValue(val.Index(i))
// 		}
// 	case reflect.Map:
// 		for _, key := range val.MapKeys() {
// 			walk(val.MapIndex(key).Interface(), fn)
// 		}
// 	case reflect.Chan:
// 		for {
// 			if v, ok := val.Recv(); ok {
// 				walkValue(v)
// 			} else {
// 				break
// 			}
// 		}
// 	}
// }

// func getValue(x interface{}) reflect.Value {
// 	val := reflect.ValueOf(x)

// 	if val.Kind() == reflect.Pointer {
// 		val = val.Elem()
// 	}

// 	return val
// }

// The next type we want to handle is func.

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walk(val.MapIndex(key).Interface(), fn)
		}
	case reflect.Chan:
		for {
			if v, ok := val.Recv(); ok {
				walkValue(v)
			} else {
				break
			}
		}
	case reflect.Func:
		valFnResult := val.Call(nil)
		for _, res := range valFnResult {
			walkValue(res)
		}
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	return val
}

// Non zero-argument functions do not seem to make a lot of sense in this scenario. But we should allow for arbitrary return values.
