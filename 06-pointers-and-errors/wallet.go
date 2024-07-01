package pointers_and_errors

import (
	"errors"
	"fmt"
)

type Bitcoin int

// Implement "Stringer" on Bitcoin

//	type Stringer interface {
//		String() string
//	}

// This interface is defined in the fmt package
// and lets you define how your type is printed when used with the %s format string in prints
// The syntax for creating a method on a type declaration is the same as it is on a struct.
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// In Go if a symbol (variables, types, functions et al)
// starts with a lowercase symbol then it is private
// outside the package it's defined in.

// In our case we want our methods to be able to manipulate this value, but no one else

// Remember we can access the internal balance field in the struct using the "receiver" variable.

type Wallet struct {
	// Go lets you create new types from existing ones
	balance Bitcoin
}

// (!) In Go, when you call a function or a method the arguments are copied.

// When calling func (w Wallet) Deposit(amount int) the w is a copy of whatever we called the method from.

// When you create a value - like a wallet, it is stored somewhere in memory.
// You can find out what the address of that bit of memory with &myVal.

// address of balance in Deposit is 0xc00000a340
// address of balance in test is 0xc00000a338

// The addresses of the two balances are different

// When we change the value of the balance inside the code,
// we are working on a copy of what came from the test.
// Therefore the balance in the test is unchanged.

// We can fix this with pointers
// Pointers let us point to some values and then let us change them.
// So rather than taking a copy of the whole Wallet,
// we instead take a pointer to that wallet so that we can change the original values within it.

// FROM:	func (w Wallet) Deposit(amount int) {
// TO:		func (w *Wallet) Deposit(amount int) {

// FROM: 	func (w Wallet) Balance() int {
// TO:		func (w *Wallet) Balance() int {

// The difference is the receiver type is *Wallet rather than Wallet which you can read as "a pointer to a wallet".

func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
	// BEFORE using a pointer * : address of balance in Deposit is 0xc00000a340
	// AFTER using a pointer * : address of balance in Deposit is 0xc00000a338
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	// But we didn't dereference the pointer in the function and seemingly addressed the object directly

	// return (*w).balance

	// The makers of Go deemed this notation cumbersome, so the language permits us to write w.balance, without an explicit dereference
	// These pointers to structs even have their own name: struct pointers and they are automatically dereferenced

	return w.balance
}

// In Go, errors are values, so we can refactor it out into a variable and have a single source of truth for it.
// The var keyword allows us to define values global to the package.
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// What should happen if you try to Withdraw more than is left in the account? For now, our requirement is to assume there is not an overdraft facility.
// How do we signal a problem when using Withdraw?
// In Go, if you want to indicate an error it is idiomatic for your function to return an err for the caller to check and act on.

func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		// errors.New creates a new error with a message of your choosing
		return ErrInsufficientFunds
	}

	w.balance -= amount

	// we have to return a nil error if successful
	return nil
}
