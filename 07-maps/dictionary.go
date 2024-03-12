package maps_dictionary

// make a custom type
type Dictionary map[string]string

// The way to handle this scenario in Go is to return a second argument which is an Error type.
// Errors can be converted to a string with the .Error() method, which we do when passing it to the assertion.
// We are also protecting assertStrings with if to ensure we don't call .Error() on nil.

// we store the error in a global variable to reuse
// var ErrNotFound = errors.New("could not find the word you were looking for")

// make a method on that type
func (d Dictionary) Search(word string) (string, error) {
	// We are using an interesting property of the map lookup
	// It can return 2 values
	// The second value is a boolean which indicates if the key was found successfully
	// This property allows us to differentiate between a word that doesn't exist
	// and a word that just doesn't have a definition
	definition, ok := d[word]

	// if the key doesn't exist
	if !ok {
		return "", ErrNotFound
	}

	// if the key does exist
	return definition, nil
}

// var ErrWordExists = errors.New("cannot add word because it already exists")

func (d Dictionary) Add(word string, definition string) error {
	// Map will not throw an error if the value already exists.
	// Instead, they will go ahead and overwrite the value with the newly provided value.
	// This can be convenient in practice, but makes our function name less than accurate.
	// Add should not modify existing values.
	// It should only add new words to our dictionary.

	// _, ok := d[word]

	// if ok {
	// 	return ErrWordExists
	// }

	// d[word] = definition
	// return nil

	_, err := d.Search(word)

	// Having a switch like this provides an extra safety net,
	// in case Search returns an error other than ErrNotFound
	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

// As our error usage grows we can make a few modifications
// We made the errors constant;
// this required us to create our own DictionaryErr type which implements the error interface
const (
	ErrNotFound   = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
	// We could reuse ErrNotFound and not add a new error.
	// However, it is often better to have a precise error for when an update fails.
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	// Go has a built-in function delete that works on maps.
	// It takes two arguments. The first is the map and the second is the key to be removed.
	// The delete function returns nothing, and we based our Delete method on the same notion
	// Since deleting a value that's not there has no effect, unlike our Update and Add methods, we don't need to complicate the API with errors.
	delete(d, word)
}
