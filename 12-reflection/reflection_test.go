package reflection_chapter

import (
	"reflect"
	"testing"
)

// write a function walk(x interface{}, fn func(string)) which takes a struct x and calls fn for all strings fields found inside. difficulty level: recursively.
// To do this we will need to use reflection.
// Reflection in computing is the ability of a program to examine its own structure, particularly through types; it's a form of metaprogramming. It's also a great source of confusion.
// https://go.dev/blog/laws-of-reflection

// You may come across scenarios though where you want to write a function where you don't know the type at compile time.
// Go lets us get around this with the type interface{} which you can think of as just any type (in fact, in Go any is an alias for interface{}).
// So walk(x interface{}, fn func(string)) will accept any value for x.

// In short only use reflection if you really need to.

// If you want polymorphic functions, consider if you could design it around an interface (not interface{}, confusingly) so that users can use your function with multiple types if they implement whatever methods you need for your function to work.

// func TestWalk(t *testing.T) {
// 	expected := "Chris"

// 	var got []string

// 	x := struct {
// 		Name string
// 	}{expected}

// 	walk(x, func(input string) {
// 		got = append(got, input)
// 	})

// 	if len(got) != 1 {
// 		t.Errorf("wrong number of function calls, got %d want %d", len(got), 1)
// 	}

// 	// to check the string passed to fn is correct
// 	if got[0] != expected {
// 		t.Errorf("got %q, want %q", got[0], expected)
// 	}
// }

// We'll want to call our function with a struct that has a string field in it (x). Then we can spy on the function (fn) passed in to see if it is called.
// We want to store a slice of strings (got) which stores which strings were passed into fn by walk. Often in previous chapters, we have made dedicated types for this to spy on function/method invocations but in this case, we can just pass in an anonymous function for fn that closes over got.
// We use an anonymous struct with a Name field of type string to go for the simplest "happy" path.
// Finally, call walk with x and the spy and for now just check the length of got, we'll be more specific with our assertions once we've got something very basic working.

// -----------

// Our code is passing for the simple case but we know our code has a lot of shortcomings.
// We're going to be writing a number of tests where we pass in different values and checking the array of strings that fn was called with.
// We should refactor our test into a table based test to make this easier to continue testing new scenarios.

// func TestWalk(t *testing.T) {
// 	cases := []struct {
// 		Name          string
// 		Input         interface{}
// 		ExpectedCalls []string
// 	}{
// 		{
// 			"struct with one string field",
// 			struct {
// 				Name string
// 			}{"Chris"},
// 			[]string{"Chris"},
// 		},
// 	}

// 	for _, test := range cases {
// 		t.Run(test.Name, func(t *testing.T) {
// 			var got []string
// 			walk(test.Input, func(input string) {
// 				got = append(got, input)
// 			})

// 			if !reflect.DeepEqual(got, test.ExpectedCalls) {
// 				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
// 			}
// 		})
// 	}
// }

// Now we can easily add a scenario to see what happens if we have more than one string field.

// ----------

// func TestWalk(t *testing.T) {
// 	cases := []struct {
// 		Name          string
// 		Input         interface{}
// 		ExpectedCalls []string
// 	}{
// 		{
// 			"struct with one string field",
// 			struct {
// 				Name string
// 			}{"Chris"},
// 			[]string{"Chris"},
// 		},
// 		{
// 			"struct with two string fields",
// 			struct {
// 				Name string
// 				City string
// 			}{"Chris", "London"},
// 			[]string{"Chris", "London"},
// 		},
// 	}

// 	for _, test := range cases {
// 		t.Run(test.Name, func(t *testing.T) {
// 			var got []string
// 			walk(test.Input, func(input string) {
// 				got = append(got, input)
// 			})

// 			if !reflect.DeepEqual(got, test.ExpectedCalls) {
// 				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
// 			}
// 		})
// 	}
// }

// ----------

// The next shortcoming in walk is that it assumes every field is a string. Let's write a test for this scenario.

// func TestWalk(t *testing.T) {
// 	cases := []struct {
// 		Name          string
// 		Input         interface{}
// 		ExpectedCalls []string
// 	}{
// 		{
// 			"struct with one string field",
// 			struct {
// 				Name string
// 			}{"Chris"},
// 			[]string{"Chris"},
// 		},
// 		{
// 			"struct with two string fields",
// 			struct {
// 				Name string
// 				City string
// 			}{"Chris", "London"},
// 			[]string{"Chris", "London"},
// 		},
// 		{
// 			"struct with non string field",
// 			struct {
// 				Name string
// 				Age  int
// 			}{"Chris", 33},
// 			[]string{"Chris"},
// 		},
// 	}

// 	for _, test := range cases {
// 		t.Run(test.Name, func(t *testing.T) {
// 			var got []string
// 			walk(test.Input, func(input string) {
// 				got = append(got, input)
// 			})

// 			if !reflect.DeepEqual(got, test.ExpectedCalls) {
// 				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
// 			}
// 		})
// 	}
// }

// -----------

// The next scenario is what if it isn't a "flat" struct? In other words, what happens if we have a struct with some nested fields?

// func TestWalk(t *testing.T) {
// 	cases := []struct {
// 		Name          string
// 		Input         interface{}
// 		ExpectedCalls []string
// 	}{
// 		{
// 			"struct with one string field",
// 			struct {
// 				Name string
// 			}{"Chris"},
// 			[]string{"Chris"},
// 		},
// 		{
// 			"struct with two string fields",
// 			struct {
// 				Name string
// 				City string
// 			}{"Chris", "London"},
// 			[]string{"Chris", "London"},
// 		},
// 		{
// 			"struct with non string field",
// 			struct {
// 				Name string
// 				Age  int
// 			}{"Chris", 33},
// 			[]string{"Chris"},
// 		},
// 		{
// 			"nested fields",
// 			struct {
// 				Name    string
// 				Profile struct {
// 					Age  int
// 					City string
// 				}
// 			}{"Chris", struct {
// 				Age  int
// 				City string
// 			}{33, "London"}},
// 			[]string{"Chris", "London"},
// 		},
// 	}

// 	for _, test := range cases {
// 		t.Run(test.Name, func(t *testing.T) {
// 			var got []string
// 			walk(test.Input, func(input string) {
// 				got = append(got, input)
// 			})

// 			if !reflect.DeepEqual(got, test.ExpectedCalls) {
// 				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
// 			}
// 		})
// 	}
// }

// Let's just refactor this by making a known type for this scenario and reference it in the test. There is a little indirection in that some of the code for our test is outside the test but readers should be able to infer the structure of the struct by looking at the initialisation.
// Add the following type declarations somewhere in your test file

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

// Now we can add this to our cases which reads a lot clearer than before

// func TestWalk(t *testing.T) {
// 	cases := []struct {
// 		Name          string
// 		Input         interface{}
// 		ExpectedCalls []string
// 	}{
// 		{
// 			"struct with one string field",
// 			struct {
// 				Name string
// 			}{"Chris"},
// 			[]string{"Chris"},
// 		},
// 		{
// 			"struct with two string fields",
// 			struct {
// 				Name string
// 				City string
// 			}{"Chris", "London"},
// 			[]string{"Chris", "London"},
// 		},
// 		{
// 			"struct with non string field",
// 			struct {
// 				Name string
// 				Age  int
// 			}{"Chris", 33},
// 			[]string{"Chris"},
// 		},
// 		{
// 			"nested fields",
// 			Person{
// 				"Chris",
// 				Profile{
// 					33, "London",
// 				},
// 			},
// 			[]string{"Chris", "London"},
// 		},
// 	}

// 	for _, test := range cases {
// 		t.Run(test.Name, func(t *testing.T) {
// 			var got []string
// 			walk(test.Input, func(input string) {
// 				got = append(got, input)
// 			})

// 			if !reflect.DeepEqual(got, test.ExpectedCalls) {
// 				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
// 			}
// 		})
// 	}
// }

// The problem is we're only iterating on the fields on the first level of the type's hierarchy.

// -----------

// What if the value of the struct passed in is a pointer?

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Chris", 33},
			[]string{"Chris"},
		},
		{
			"nested fields",
			Person{
				"Chris",
				Profile{
					33, "London",
				},
			},
			[]string{"Chris", "London"},
		},
		{
			"pointers to things",
			&Person{
				"Chris",
				Profile{
					33, "London",
				},
			},
			[]string{"Chris", "London"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}
