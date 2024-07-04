package main

import (
	"strings"
)

// Some companies will ask you to do the Roman Numeral Kata as part of the interview process. This chapter will show how you can tackle it with TDD.
// We are going to write a function which converts an Arabic number (numbers 0 to 9) to a Roman Numeral.
// If you haven't heard of Roman Numerals they are how the Romans wrote down numbers.
// You build them by sticking symbols together and those symbols represent numbers
// So I is "one". III is three.
// Seems easy but there's a few interesting rules. V means five, but IV is 4 (not IIII).
// MCMLXXXIV is 1984. That looks complicated and it's hard to imagine how we can write code to figure this out right from the start.
// As this book stresses, a key skill for software developers is to try and identify "thin vertical slices" of useful functionality and then iterating.
// The TDD workflow helps facilitate iterative development.
// So rather than 1984, let's start with 1.

// func ConvertToRoman(arabic int) string {
// 	return "I"
// }

// I know it feels weird just to hard-code the result but with TDD we want to stay out of "red" for as long as possible.
// It may feel like we haven't accomplished much but we've defined our API and got a test capturing one of our rules; even if the "real" code is pretty dumb.

// func ConvertToRoman(arabic int) string {
// 	if arabic == 2 {
// 		return "II"
// 	}
// 	return "I"
// }

// Yup, it still feels like we're not actually tackling the problem.

// func ConvertToRoman(arabic int) string {
// 	if arabic == 3 {
// 		return "III"
// 	}
// 	if arabic == 2 {
// 		return "II"
// 	}
// 	return "I"
// }

// OK so I'm starting to not enjoy these if statements and if you look at the code hard enough you can see that we're building a string of I based on the size of arabic.
// We "know" that for more complicated numbers we will be doing some kind of arithmetic and string concatenation.
// Let's try a refactor with these thoughts in mind, it might not be suitable for the end solution but that's OK.
// We can always throw our code away and start afresh with the tests we have to guide us.

// func ConvertToRoman(arabic int) string {
// 	var result strings.Builder

// 	for i := 0; i < arabic; i++ {
// 		result.WriteString("I")
// 	}

// 	return result.String()
// }

// You may not have used strings.Builder before
// A Builder is used to efficiently build a string using Write methods.
// It minimizes memory copying.
// Normally I wouldn't bother with such optimisations until I have an actual performance problem but the amount of code is not much larger than a "manual" appending on a string so we may as well use the faster approach.
// The code looks better to me and describes the domain as we know it right now.

// Things start getting more complicated now.
// The Romans in their wisdom thought repeating characters would become hard to read and count.
// So a rule with Roman Numerals is you can't have the same character repeated more than 3 times in a row.

// Instead you take the next highest symbol and then "subtract" by putting a symbol to the left of it.
// Not all symbols can be used as subtractors; only I (1), X (10) and C (100).

// For example 5 in Roman Numerals is V.
// To create 4 you do not do IIII, instead you do IV.

// func ConvertToRoman(arabic int) string {
// 	if arabic == 4 {
// 		return "IV"
// 	}

// 	var result strings.Builder

// 	for i := 0; i < arabic; i++ {
// 		result.WriteString("I")
// 	}

// 	return result.String()
// }

// I don't "like" that we have broken our string building pattern and I want to carry on with it.

// func ConvertToRoman(arabic int) string {
// 	var result strings.Builder

// 	for i := arabic; i > 0; i-- {
// 		if i == 4 {
// 			result.WriteString("IV")
// 			break
// 		}
// 		result.WriteString("I")
// 	}

// 	return result.String()
// }

// In order for 4 to "fit" with my current thinking I now count down from the Arabic number, adding symbols to our string as we progress. Not sure if this will work in the long run but let's see!

// func ConvertToRoman(arabic int) string {
// 	var result strings.Builder

// 	for i := arabic; i > 0; i-- {
// 		if i == 5 {
// 			result.WriteString("V")
// 			break
// 		}
// 		if i == 4 {
// 			result.WriteString("IV")
// 			break
// 		}
// 		result.WriteString("I")
// 	}

// 	return result.String()
// }

// Repetition in loops like this are usually a sign of an abstraction waiting to be called out. Short-circuiting loops can be an effective tool for readability but it could also be telling you something else.
// We are looping over our Arabic number and if we hit certain symbols we are calling break but what we are really doing is subtracting over i in a ham-fisted manner.

// func ConvertToRoman(arabic int) string {
// 	var result strings.Builder

// 	for arabic > 0 {
// 		switch {
// 		case arabic > 4:
// 			result.WriteString("V")
// 			arabic -= 5
// 		case arabic > 3:
// 			result.WriteString("IV")
// 			arabic -= 4
// 		default:
// 			result.WriteString("I")
// 			arabic--
// 		}
// 	}

// 	return result.String()
// }

// Given the signals I'm reading from our code, driven from our tests of some very basic scenarios I can see that to build a Roman Numeral I need to subtract from arabic as I apply symbols
// The for loop no longer relies on an i and instead we will keep building our string until we have subtracted enough symbols away from arabic.

// I'm pretty sure this approach will be valid for 6 (VI), 7 (VII) and 8 (VIII) too. Nonetheless add the cases in to our test suite and check.

// 9 follows the same rule as 4 in that we should subtract I from the representation of the following number. 10 is represented in Roman Numerals with X; so therefore 9 should be IX.

// We should be able to adopt the same approach as before

// func ConvertToRoman(arabic int) string {
// 	var result strings.Builder

// 	for arabic > 0 {
// 		switch {
// 		case arabic > 8:
// 			result.WriteString("IX")
// 			arabic -= 9
// 		case arabic > 4:
// 			result.WriteString("V")
// 			arabic -= 5
// 		case arabic > 3:
// 			result.WriteString("IV")
// 			arabic -= 4
// 		default:
// 			result.WriteString("I")
// 			arabic--
// 		}
// 	}

// 	return result.String()
// }

// It feels like the code is still telling us there's a refactor somewhere but it's not totally obvious to me, so let's keep going.
// I'll skip the code for this too, but add to your test cases a test for 10 which should be X and make it pass before reading on.

// If you've ever done OO programming, you'll know that you should view switch statements with a bit of suspicion.
// Usually you are capturing a concept or data inside some imperative code when in fact it could be captured in a class structure instead.

// Go isn't strictly OO but that doesn't mean we ignore the lessons OO offers entirely (as much as some would like to tell you).

// Our switch statement is describing some truths about Roman Numerals along with behaviour.

// We can refactor this by decoupling the data from the behaviour.

// type RomanNumeral struct {
// 	Value  int
// 	Symbol string
// }

// var allRomanNumerals = []RomanNumeral{
// 	{10, "X"},
// 	{9, "IX"},
// 	{5, "V"},
// 	{4, "IV"},
// 	{1, "I"},
// }

// func ConvertToRoman(arabic int) string {
// 	var result strings.Builder

// 	for _, numeral := range allRomanNumerals {
// 		for arabic >= numeral.Value {
// 			result.WriteString(numeral.Symbol)
// 			arabic -= numeral.Value
// 		}
// 	}

// 	return result.String()
// }

// This feels much better.
// We've declared some rules around the numerals as data rather than hidden in an algorithm and we can see how we just work through the Arabic number, trying to add symbols to our result if they fit.

// type RomanNumeral struct {
// 	Value  int
// 	Symbol string
// }

// var allRomanNumerals = []RomanNumeral{
// 	{50, "L"},
// 	{40, "XL"},
// 	{10, "X"},
// 	{9, "IX"},
// 	{5, "V"},
// 	{4, "IV"},
// 	{1, "I"},
// }

// func ConvertToRoman(arabic int) string {
// 	var result strings.Builder

// 	for _, numeral := range allRomanNumerals {
// 		for arabic >= numeral.Value {
// 			result.WriteString(numeral.Symbol)
// 			arabic -= numeral.Value
// 		}
// 	}

// 	return result.String()
// }

// Take the same approach for the remaining symbols, it should just be a matter of adding data to both the tests and our array of symbols.

// type RomanNumeral struct {
// 	Value  int
// 	Symbol string
// }

// var allRomanNumerals = []RomanNumeral{
// 	{1000, "M"},
// 	{900, "CM"},
// 	{500, "D"},
// 	{400, "CD"},
// 	{100, "C"},
// 	{90, "XC"},
// 	{50, "L"},
// 	{40, "XL"},
// 	{10, "X"},
// 	{9, "IX"},
// 	{5, "V"},
// 	{4, "IV"},
// 	{1, "I"},
// }

// func ConvertToRoman(arabic int) string {
// 	var result strings.Builder

// 	for _, numeral := range allRomanNumerals {
// 		for arabic >= numeral.Value {
// 			result.WriteString(numeral.Symbol)
// 			arabic -= numeral.Value
// 		}
// 	}

// 	return result.String()
// }

// I didn't change the algorithm, all I had to do was update the allRomanNumerals array.

// ----------

// We're not done yet. Next we're going to write a function that converts from a Roman Numeral to an int
// Add our new function definition

// func ConvertToArabic(roman string) int {
// 	return 0
// }

// func ConvertToArabic(roman string) int {
// 	return 1
// }

// func ConvertToArabic(roman string) int {
// 	if roman == "III" {
// 		return 3
// 	}
// 	if roman == "II" {
// 		return 2
// 	}
// 	return 1
// }

// Through the dumbness of real code that works we can start to see a pattern like before.
// We need to iterate through the input and build something, in this case a total.

// func ConvertToArabic(roman string) int {
// 	total := 0
// 	for range roman {
// 		total++
// 	}
// 	return total
// }

// func ConvertToArabic(roman string) int {
// 	var arabic = 0

// 	for _, numeral := range allRomanNumerals {
// 		for strings.HasPrefix(roman, numeral.Symbol) {
// 			arabic += numeral.Value
// 			roman = strings.TrimPrefix(roman, numeral.Symbol)
// 		}
// 	}

// 	return arabic
// }

// It is basically the algorithm of ConvertToRoman(int) implemented backwards.
// Here, we loop over the given roman numeral string:
// We look for roman numeral symbols taken from allRomanNumerals, highest to lowest, at the beginning of the string.
// If we find the prefix, we add its value to arabic and trim the prefix.
// At the end, we return the sum as the arabic number.

// The HasPrefix(s, prefix) checks whether string s starts with prefix and TrimPrefix(s, prefix) removes the prefix from s, so we can proceed with the remaining roman numeral symbols.
// It works with IV and all other test cases.

// You can implement this as a recursive function, which is more elegant (in my opinion) but might be slower.
// I'll leave this up to you and some Benchmark... tests.

// ----------

// Go has types for unsigned integers, which means they cannot be negative; so that rules out one class of bug in our code immediately.
// By adding 16, it means it is a 16 bit integer which can store a max of 65535, which is still too big but gets us closer to what we need.

// Try updating the code to use uint16 rather than int.

var allRomanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

type RomanNumeral struct {
	Value  uint16
	Symbol string
}

func ConvertToRoman(arabic uint16) string {
	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String()
}

func ConvertToArabic(roman string) uint16 {
	var arabic uint16 = 0

	for _, numeral := range allRomanNumerals {
		for strings.HasPrefix(roman, numeral.Symbol) {
			arabic += numeral.Value
			roman = strings.TrimPrefix(roman, numeral.Symbol)
		}
	}

	return arabic
}
