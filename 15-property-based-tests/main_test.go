package main

import (
	"fmt"
	"testing"
	"testing/quick"
)

// func TestRomanNumberals(t *testing.T) {
// 	got := ConvertToRoman(1)
// 	want := "I"

// 	if got != want {
// 		t.Errorf("got %q, want %q", got, want)
// 	}
// }

// Now use that uneasy feeling to write a new test to force us to write slightly less dumb code.

// We can use subtests to nicely group our tests

// func TestRomanNumerals(t *testing.T) {
// 	t.Run("1 gets converted to I", func(t *testing.T) {
// 		got := ConvertToRoman(1)
// 		want := "I"

// 		if got != want {
// 			t.Errorf("got %q, want %q", got, want)
// 		}
// 	})

// 	t.Run("2 gets converted to II", func(t *testing.T) {
// 		got := ConvertToRoman(2)
// 		want := "II"

// 		if got != want {
// 			t.Errorf("got %q, want %q", got, want)
// 		}
// 	})
// }

// So we need to write more tests to drive us forward.

// We have some repetition in our tests. When you're testing something which feels like it's a matter of "given input X, we expect Y" you should probably use table based tests.

// func TestRomanNumerals(t *testing.T) {
// 	cases := []struct {
// 		Description string
// 		Arabic      int
// 		Want        string
// 	}{
// 		{"1 gets converted to I", 1, "I"},
// 		{"2 gets converted to II", 2, "II"},
// 	}

// 	for _, test := range cases {
// 		t.Run(test.Description, func(t *testing.T) {
// 			got := ConvertToRoman(test.Arabic)
// 			if got != test.Want {
// 				t.Errorf("got %q, want %q", got, test.Want)
// 			}
// 		})
// 	}
// }

// We can now easily add more cases without having to write any more test boilerplate.

// func TestRomanNumerals(t *testing.T) {
// 	cases := []struct {
// 		Description string
// 		Arabic      int
// 		Want        string
// 	}{
// 		{"1 gets converted to I", 1, "I"},
// 		{"2 gets converted to II", 2, "II"},
// 		{"3 gets converted to III", 3, "III"},
// 	}

// 	for _, test := range cases {
// 		t.Run(test.Description, func(t *testing.T) {
// 			got := ConvertToRoman(test.Arabic)
// 			if got != test.Want {
// 				t.Errorf("got %q, want %q", got, test.Want)
// 			}
// 		})
// 	}
// }

// func TestRomanNumerals(t *testing.T) {
// 	cases := []struct {
// 		Description string
// 		Arabic      int
// 		Want        string
// 	}{
// 		{"1 gets converted to I", 1, "I"},
// 		{"2 gets converted to II", 2, "II"},
// 		{"3 gets converted to III", 3, "III"},
// 		{"4 gets converted to IV (can't repeat more than 3 times)", 4, "IV"},
// 	}

// 	for _, test := range cases {
// 		t.Run(test.Description, func(t *testing.T) {
// 			got := ConvertToRoman(test.Arabic)
// 			if got != test.Want {
// 				t.Errorf("got %q, want %q", got, test.Want)
// 			}
// 		})
// 	}
// }

// func TestRomanNumerals(t *testing.T) {
// 	cases := []struct {
// 		Description string
// 		Arabic      int
// 		Want        string
// 	}{
// 		{"1 gets converted to I", 1, "I"},
// 		{"2 gets converted to II", 2, "II"},
// 		{"3 gets converted to III", 3, "III"},
// 		{"4 gets converted to IV (can't repeat more than 3 times)", 4, "IV"},
// 		{"5 gets converted to V", 5, "V"},
// 	}

// 	for _, test := range cases {
// 		t.Run(test.Description, func(t *testing.T) {
// 			got := ConvertToRoman(test.Arabic)
// 			if got != test.Want {
// 				t.Errorf("got %q, want %q", got, test.Want)
// 			}
// 		})
// 	}
// }

// func TestRomanNumerals(t *testing.T) {
// 	cases := []struct {
// 		Description string
// 		Arabic      int
// 		Want        string
// 	}{
// 		{"1 gets converted to I", 1, "I"},
// 		{"2 gets converted to II", 2, "II"},
// 		{"3 gets converted to III", 3, "III"},
// 		{"4 gets converted to IV (can't repeat more than 3 times)", 4, "IV"},
// 		{"5 gets converted to V", 5, "V"},
// 		{"6 gets converted to VI", 6, "VI"},
// 		{"7 gets converted to VII", 7, "VII"},
// 		{"8 gets converted to VIII", 8, "VIII"},
// 	}

// 	for _, test := range cases {
// 		t.Run(test.Description, func(t *testing.T) {
// 			got := ConvertToRoman(test.Arabic)
// 			if got != test.Want {
// 				t.Errorf("got %q, want %q", got, test.Want)
// 			}
// 		})
// 	}
// }

// func TestRomanNumerals(t *testing.T) {
// 	cases := []struct {
// 		Description string
// 		Arabic      int
// 		Want        string
// 	}{
// 		{"1 gets converted to I", 1, "I"},
// 		{"2 gets converted to II", 2, "II"},
// 		{"3 gets converted to III", 3, "III"},
// 		{"4 gets converted to IV (can't repeat more than 3 times)", 4, "IV"},
// 		{"5 gets converted to V", 5, "V"},
// 		{"6 gets converted to VI", 6, "VI"},
// 		{"7 gets converted to VII", 7, "VII"},
// 		{"8 gets converted to VIII", 8, "VIII"},
// 		{"9 gets converted to IX", 9, "IX"},
// 	}

// 	for _, test := range cases {
// 		t.Run(test.Description, func(t *testing.T) {
// 			got := ConvertToRoman(test.Arabic)
// 			if got != test.Want {
// 				t.Errorf("got %q, want %q", got, test.Want)
// 			}
// 		})
// 	}
// }

// func TestRomanNumerals(t *testing.T) {
// 	cases := []struct {
// 		Description string
// 		Arabic      int
// 		Want        string
// 	}{
// 		{"1 gets converted to I", 1, "I"},
// 		{"2 gets converted to II", 2, "II"},
// 		{"3 gets converted to III", 3, "III"},
// 		{"4 gets converted to IV (can't repeat more than 3 times)", 4, "IV"},
// 		{"5 gets converted to V", 5, "V"},
// 		{"6 gets converted to VI", 6, "VI"},
// 		{"7 gets converted to VII", 7, "VII"},
// 		{"8 gets converted to VIII", 8, "VIII"},
// 		{"9 gets converted to IX", 9, "IX"},
// 		{"10 gets converted to X", 10, "X"},
// 		{"14 gets converted to XIV", 14, "XIV"},
// 		{"18 gets converted to XVIII", 18, "XVIII"},
// 		{"20 gets converted to XX", 20, "XX"},
// 		{"39 gets converted to XXXIX", 39, "XXXIX"},
// 	}

// 	for _, test := range cases {
// 		t.Run(test.Description, func(t *testing.T) {
// 			got := ConvertToRoman(test.Arabic)
// 			if got != test.Want {
// 				t.Errorf("got %q, want %q", got, test.Want)
// 			}
// 		})
// 	}
// }

// // Does this abstraction work for bigger numbers? Extend the test suite so it works for the Roman number for 50 which is L.

// func TestRomanNumerals(t *testing.T) {
// 	cases := []struct {
// 		Description string
// 		Arabic      int
// 		Want        string
// 	}{
// 		{"1 gets converted to I", 1, "I"},
// 		{"2 gets converted to II", 2, "II"},
// 		{"3 gets converted to III", 3, "III"},
// 		{"4 gets converted to IV (can't repeat more than 3 times)", 4, "IV"},
// 		{"5 gets converted to V", 5, "V"},
// 		{"6 gets converted to VI", 6, "VI"},
// 		{"7 gets converted to VII", 7, "VII"},
// 		{"8 gets converted to VIII", 8, "VIII"},
// 		{"9 gets converted to IX", 9, "IX"},
// 		{"10 gets converted to X", 10, "X"},
// 		{"14 gets converted to XIV", 14, "XIV"},
// 		{"18 gets converted to XVIII", 18, "XVIII"},
// 		{"20 gets converted to XX", 20, "XX"},
// 		{"39 gets converted to XXXIX", 39, "XXXIX"},
// 		{"40 gets converted to XL", 40, "XL"},
// 		{"47 gets converted to XLVII", 47, "XLVII"},
// 		{"49 gets converted to XLIX", 49, "XLIX"},
// 		{"50 gets converted to L", 50, "L"},
// 	}

// 	for _, test := range cases {
// 		t.Run(test.Description, func(t *testing.T) {
// 			got := ConvertToRoman(test.Arabic)
// 			if got != test.Want {
// 				t.Errorf("got %q, want %q", got, test.Want)
// 			}
// 		})
// 	}
// }

// I removed description as I felt the data described enough of the information.
// I added a few other edge cases I found just to give me a little more confidence.
// With table based tests this is very cheap to do.

// func TestRomanNumerals(t *testing.T) {
// 	cases := []struct {
// 		Arabic int
// 		Roman  string
// 	}{
// 		{Arabic: 1, Roman: "I"},
// 		{Arabic: 2, Roman: "II"},
// 		{Arabic: 3, Roman: "III"},
// 		{Arabic: 4, Roman: "IV"},
// 		{Arabic: 5, Roman: "V"},
// 		{Arabic: 6, Roman: "VI"},
// 		{Arabic: 7, Roman: "VII"},
// 		{Arabic: 8, Roman: "VIII"},
// 		{Arabic: 9, Roman: "IX"},
// 		{Arabic: 10, Roman: "X"},
// 		{Arabic: 14, Roman: "XIV"},
// 		{Arabic: 18, Roman: "XVIII"},
// 		{Arabic: 20, Roman: "XX"},
// 		{Arabic: 39, Roman: "XXXIX"},
// 		{Arabic: 40, Roman: "XL"},
// 		{Arabic: 47, Roman: "XLVII"},
// 		{Arabic: 49, Roman: "XLIX"},
// 		{Arabic: 50, Roman: "L"},
// 		{Arabic: 100, Roman: "C"},
// 		{Arabic: 90, Roman: "XC"},
// 		{Arabic: 400, Roman: "CD"},
// 		{Arabic: 500, Roman: "D"},
// 		{Arabic: 900, Roman: "CM"},
// 		{Arabic: 1000, Roman: "M"},
// 		{Arabic: 1984, Roman: "MCMLXXXIV"},
// 		{Arabic: 3999, Roman: "MMMCMXCIX"},
// 		{Arabic: 2014, Roman: "MMXIV"},
// 		{Arabic: 1006, Roman: "MVI"},
// 		{Arabic: 798, Roman: "DCCXCVIII"},
// 	}
// 	for _, test := range cases {
// 		t.Run(fmt.Sprintf("%d gets converted to %q", test.Arabic, test.Roman), func(t *testing.T) {
// 			got := ConvertToRoman(test.Arabic)
// 			if got != test.Roman {
// 				t.Errorf("got %q, want %q", got, test.Roman)
// 			}
// 		})
// 	}
// }

// We can re-use our test cases here with a little refactoring

// Move the cases variable outside of the test as a package variable in a var block.

// var (
// 	cases = []struct {
// 		Arabic int
// 		Roman  string
// 	}{
// 		{Arabic: 1, Roman: "I"},
// 		{Arabic: 2, Roman: "II"},
// 		{Arabic: 3, Roman: "III"},
// 		{Arabic: 4, Roman: "IV"},
// 		{Arabic: 5, Roman: "V"},
// 		{Arabic: 6, Roman: "VI"},
// 		{Arabic: 7, Roman: "VII"},
// 		{Arabic: 8, Roman: "VIII"},
// 		{Arabic: 9, Roman: "IX"},
// 		{Arabic: 10, Roman: "X"},
// 		{Arabic: 14, Roman: "XIV"},
// 		{Arabic: 18, Roman: "XVIII"},
// 		{Arabic: 20, Roman: "XX"},
// 		{Arabic: 39, Roman: "XXXIX"},
// 		{Arabic: 40, Roman: "XL"},
// 		{Arabic: 47, Roman: "XLVII"},
// 		{Arabic: 49, Roman: "XLIX"},
// 		{Arabic: 50, Roman: "L"},
// 		{Arabic: 100, Roman: "C"},
// 		{Arabic: 90, Roman: "XC"},
// 		{Arabic: 400, Roman: "CD"},
// 		{Arabic: 500, Roman: "D"},
// 		{Arabic: 900, Roman: "CM"},
// 		{Arabic: 1000, Roman: "M"},
// 		{Arabic: 1984, Roman: "MCMLXXXIV"},
// 		{Arabic: 3999, Roman: "MMMCMXCIX"},
// 		{Arabic: 2014, Roman: "MMXIV"},
// 		{Arabic: 1006, Roman: "MVI"},
// 		{Arabic: 798, Roman: "DCCXCVIII"},
// 	}
// )

// func TestConvertingToArabic(t *testing.T) {
// 	for _, test := range cases[:1] {
// 		t.Run(fmt.Sprintf("%q gets converted to %d", test.Roman, test.Arabic), func(t *testing.T) {
// 			got := ConvertToArabic(test.Roman)
// 			if got != test.Arabic {
// 				t.Errorf("got %d, want %d", got, test.Arabic)
// 			}
// 		})
// 	}
// }

// Next, change the slice index in our test to move to the next test case (e.g. cases[:2]). Make it pass yourself with the dumbest code you can think of, continue writing dumb code (best book ever right?) for the third case too.

// func TestConvertingToArabic(t *testing.T) {
// 	for _, test := range cases[:2] {
// 		t.Run(fmt.Sprintf("%q gets converted to %d", test.Roman, test.Arabic), func(t *testing.T) {
// 			got := ConvertToArabic(test.Roman)
// 			if got != test.Arabic {
// 				t.Errorf("got %d, want %d", got, test.Arabic)
// 			}
// 		})
// 	}
// }

// Next we move to cases[:4] (IV) which now fails because it gets 2 back as that's the length of the string.

// func TestConvertingToArabic(t *testing.T) {
// 	for _, test := range cases[:4] {
// 		t.Run(fmt.Sprintf("%q gets converted to %d", test.Roman, test.Arabic), func(t *testing.T) {
// 			got := ConvertToArabic(test.Roman)
// 			if got != test.Arabic {
// 				t.Errorf("got %d, want %d", got, test.Arabic)
// 			}
// 		})
// 	}
// }

// Now that we have our functions to convert an arabic number into a roman numeral and back, we can take our tests a step further:

// There have been a few rules in the domain of Roman Numerals that we have worked with in this chapter
// - Can't have more than 3 consecutive symbols
// - Only I (1), X (10) and C (100) can be "subtractors"
// - Taking the result of ConvertToRoman(N) and passing it to ConvertToArabic should return us N

// The tests we have written so far can be described as "example" based tests where we provide examples for the tooling to verify.

// What if we could take these rules that we know about our domain and somehow exercise them against our code?

// Property based tests help you do this by throwing random data at your code and verifying the rules you describe always hold true.
// A lot of people think property based tests are mainly about random data but they would be mistaken.
// The real challenge about property based tests is having a good understanding of your domain so you can write these properties.

// func TestPropertiesOfConversion(t *testing.T) {
// 	assertion := func(arabic int) bool {
// 		roman := ConvertToRoman(arabic)
// 		fromRoman := ConvertToArabic(roman)
// 		return fromRoman == arabic
// 	}

// 	if err := quick.Check(assertion, nil); err != nil {
// 		t.Error("failed checks", err)
// 	}
// }

// Our first test will check that if we transform a number into Roman, when we use our other function to convert it back to a number that we get what we originally had.

// - Given random number (e.g 4).
// - Call ConvertToRoman with random number (should return IV if 4).
// - Take the result of above and pass it to ConvertToArabic.
// - The above should give us our original input (4).

// This feels like a good test to build us confidence because it should break if there's a bug in either.
// The only way it could pass is if they have the same kind of bug; which isn't impossible but feels unlikely.

// We're using the testing/quick package from the standard library

// Reading from the bottom, we provide quick.Check a function that it will run against a number of random inputs, if the function returns false it will be seen as failing the check.

// Our assertion function above takes a random number and runs our functions to test the property.

// -----------

// Try running it; your computer may hang for a while, so kill it when you're bored :)
// What's going on? Try adding the following to the assertion code.

// func TestPropertiesOfConversion(t *testing.T) {
// 	assertion := func(arabic int) bool {
// 		if arabic < 0 || arabic > 3999 {
// 			log.Println(arabic)
// 			return true
// 		}
// 		roman := ConvertToRoman(arabic)
// 		fromRoman := ConvertToArabic(roman)
// 		return fromRoman == arabic
// 	}

// 	if err := quick.Check(assertion, nil); err != nil {
// 		t.Error("failed checks", err)
// 	}
// }

// Just running this very simple property has exposed a flaw in our implementation. We used int as our input but:
// - You can't do negative numbers with Roman Numerals
// - Given our rule of a max of 3 consecutive symbols we can't represent a value greater than 3999 (well, kinda) and int has a much higher maximum value than 3999.

// This is great! We've been forced to think more deeply about our domain which is a real strength of property based tests.
// Clearly int is not a great type. What if we tried something a little more appropriate?

// ----------

// I updated assertion in the test to give a bit more visibility.

var (
	cases = []struct {
		Arabic uint16
		Roman  string
	}{
		{Arabic: 1, Roman: "I"},
		{Arabic: 2, Roman: "II"},
		{Arabic: 3, Roman: "III"},
		{Arabic: 4, Roman: "IV"},
		{Arabic: 5, Roman: "V"},
		{Arabic: 6, Roman: "VI"},
		{Arabic: 7, Roman: "VII"},
		{Arabic: 8, Roman: "VIII"},
		{Arabic: 9, Roman: "IX"},
		{Arabic: 10, Roman: "X"},
		{Arabic: 14, Roman: "XIV"},
		{Arabic: 18, Roman: "XVIII"},
		{Arabic: 20, Roman: "XX"},
		{Arabic: 39, Roman: "XXXIX"},
		{Arabic: 40, Roman: "XL"},
		{Arabic: 47, Roman: "XLVII"},
		{Arabic: 49, Roman: "XLIX"},
		{Arabic: 50, Roman: "L"},
		{Arabic: 100, Roman: "C"},
		{Arabic: 90, Roman: "XC"},
		{Arabic: 400, Roman: "CD"},
		{Arabic: 500, Roman: "D"},
		{Arabic: 900, Roman: "CM"},
		{Arabic: 1000, Roman: "M"},
		{Arabic: 1984, Roman: "MCMLXXXIV"},
		{Arabic: 3999, Roman: "MMMCMXCIX"},
		{Arabic: 2014, Roman: "MMXIV"},
		{Arabic: 1006, Roman: "MVI"},
		{Arabic: 798, Roman: "DCCXCVIII"},
	}
)

func TestConvertingToArabic(t *testing.T) {
	for _, test := range cases[:4] {
		t.Run(fmt.Sprintf("%q gets converted to %d", test.Roman, test.Arabic), func(t *testing.T) {
			got := ConvertToArabic(test.Roman)
			if got != test.Arabic {
				t.Errorf("got %d, want %d", got, test.Arabic)
			}
		})
	}
}

// func TestPropertiesOfConversion(t *testing.T) {
// 	assertion := func(arabic uint16) bool {
// 		if arabic > 3999 {
// 			return true
// 		}
// 		t.Log("testing", arabic)
// 		roman := ConvertToRoman(arabic)
// 		fromRoman := ConvertToArabic(roman)
// 		return fromRoman == arabic
// 	}

// 	if err := quick.Check(assertion, nil); err != nil {
// 		t.Error("failed checks", err)
// 	}
// }

// Notice that now we are logging the input using the log method from the testing framework.
// Make sure you run the go test command with the flag -v to print the additional output (go test -v).

// If you run the test they now actually run and you can see what is being tested.
// You can run multiple times to see our code stands up well to the various values!
// This gives me a lot of confidence that our code is working how we want.

// The default number of runs quick.Check performs is 100 but you can change that with a config.

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			return true
		}
		t.Log("testing", arabic)
		roman := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)
		return fromRoman == arabic
	}

	if err := quick.Check(assertion, &quick.Config{
		MaxCount: 1000,
	}); err != nil {
		t.Error("failed checks", err)
	}
}
